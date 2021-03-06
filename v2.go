package lens

import (
	"context"
	"fmt"
	"strings"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/RTradeLtd/grpc/lensv2"
	"github.com/RTradeLtd/rtfs/v2"

	"github.com/RTradeLtd/Lens/v2/analyzer/images"
	"github.com/RTradeLtd/Lens/v2/analyzer/ocr"
	"github.com/RTradeLtd/Lens/v2/engine"
	"github.com/RTradeLtd/Lens/v2/source/planetary"
)

// V2 is the new Lens API, and implements the LensV2 gRPC interface directly.
type V2 struct {
	se   engine.Searcher
	ipfs rtfs.Manager

	// Analysis classes
	oc *ocr.Analyzer
	px *planetary.Extractor
	tf images.TensorflowAnalyzer

	l *zap.SugaredLogger
}

// V2Options denotes options for the V2 Lens API
type V2Options struct {
	TesseractConfigPath string

	Engine engine.Opts
}

// NewV2 instantiates a new V2 API
func NewV2(
	opts V2Options,
	ipfs rtfs.Manager,
	ia images.TensorflowAnalyzer,
	logger *zap.SugaredLogger,
) (*V2, error) {
	if logger == nil {
		logger = zap.NewNop().Sugar()
	}

	// create new engine
	se, err := engine.New(logger.Named("engine"), opts.Engine)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate search engine: %s", err.Error())
	}
	go se.Run()

	return &V2{
		se:   se,
		ipfs: ipfs,

		tf: ia,
		px: planetary.NewPlanetaryExtractor(ipfs),
		oc: ocr.NewAnalyzer(opts.TesseractConfigPath, logger.Named("ocr")),
		l:  logger.Named("service.v2"),
	}, nil
}

// NewV2WithEngine instantiates a Lens V2 service with the given engine
func NewV2WithEngine(
	opts V2Options,
	ipfs rtfs.Manager,
	ia images.TensorflowAnalyzer,
	se engine.Searcher,
	logger *zap.SugaredLogger,
) *V2 {
	if logger == nil {
		logger = zap.NewNop().Sugar()
	}

	return &V2{
		se:   se,
		ipfs: ipfs,

		tf: ia,
		px: planetary.NewPlanetaryExtractor(ipfs),
		oc: ocr.NewAnalyzer(opts.TesseractConfigPath, logger.Named("ocr")),
		l:  logger.Named("service.v2"),
	}
}

// Close releases Lens resources
func (v *V2) Close() { v.se.Close() }

// Index analyzes and stores the given object
func (v *V2) Index(ctx context.Context, req *lensv2.IndexReq) (*lensv2.IndexResp, error) {
	var l = v.l.With("request", req)
	switch req.GetType() {
	case lensv2.IndexReq_IPLD:
		break
	default:
		return nil, status.Errorf(codes.InvalidArgument,
			"invalid data type '%s' provided", req.GetType())
	}

	var hash = req.GetHash()
	var reindex = req.GetOptions().GetReindex()
	content, md, err := v.magnify(hash, magnifyOpts{
		DisplayName: req.GetDisplayName(),
		Tags:        req.GetTags(),
		Reindex:     reindex,
	})
	if err != nil {
		l.Errorw("failed to magnify document", "error", err)
		if strings.Contains(err.Error(), "failed to find content") {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Errorf(codes.FailedPrecondition,
			"failed to perform magnification for '%s': %s", hash, err.Error())
	}

	if err = v.store(hash, content, md, reindex); err != nil {
		l.Errorw("failed to store document", "error", err)
		return nil, status.Errorf(codes.Internal,
			"failed to store requested document: %s", err.Error())
	}

	l.Info("document indexed")

	return &lensv2.IndexResp{
		Doc: &lensv2.Document{
			Hash:        hash,
			DisplayName: md.DisplayName,
			MimeType:    md.MimeType,
			Category:    md.Category,
			Tags:        md.Tags,
		},
	}, nil
}

// Search executes a query against the Lens index
func (v *V2) Search(ctx context.Context, req *lensv2.SearchReq) (*lensv2.SearchResp, error) {
	var (
		err     error
		results []engine.Result
		opts    = req.GetOptions()
	)

	if req.GetQuery() == "" &&
		len(opts.GetCategories()) < 1 &&
		len(opts.GetHashes()) < 1 &&
		len(opts.GetMimeTypes()) < 1 &&
		len(opts.GetRequired()) < 1 &&
		len(opts.GetTags()) < 1 {
		return nil, status.Errorf(codes.InvalidArgument,
			"no search parameters provided")
	}

	if opts == nil {
		results, err = v.se.Search(ctx, engine.Query{Text: req.GetQuery()})
	} else {
		results, err = v.se.Search(ctx, engine.Query{
			Text:       req.GetQuery(),
			Required:   opts.GetRequired(),
			Tags:       opts.GetTags(),
			Categories: opts.GetCategories(),
			MimeTypes:  opts.GetMimeTypes(),
			Hashes:     opts.GetHashes(),
		})
	}
	if err != nil {
		v.l.Errorw("error occured on query execution",
			"error", err, "query", req)
		return nil, status.Errorf(codes.Internal,
			"error occured on query execution: %s", err.Error())
	}

	v.l.Debugw("query completed",
		"query", req, "results", len(results))
	return &lensv2.SearchResp{
		Results: func() []*lensv2.SearchResp_Result {
			var formatted = make([]*lensv2.SearchResp_Result, len(results))
			for i := 0; i < len(results); i++ {
				var r = results[i]
				formatted[i] = &lensv2.SearchResp_Result{
					Score: float32(r.Score),
					Doc: &lensv2.Document{
						Hash:        r.Hash,
						DisplayName: r.MD.DisplayName,
						MimeType:    r.MD.MimeType,
						Category:    r.MD.Category,
						Tags:        r.MD.Tags,
					},
				}
			}
			return formatted
		}(),
	}, nil
}

// Remove unindexes and deletes the requested object
func (v *V2) Remove(ctx context.Context, req *lensv2.RemoveReq) (*lensv2.RemoveResp, error) {
	if req.GetHash() == "" {
		return nil, status.Errorf(codes.InvalidArgument,
			"no hash to remove was provided")
	}

	if err := v.remove(req.GetHash()); err != nil {
		return nil, status.Errorf(codes.NotFound,
			"failed to remove requested hash: %s", err.Error())
	}

	return &lensv2.RemoveResp{}, nil
}
