package queuebatch

import (
	"go.opentelemetry.io/collector/exporter/exporterhelper/internal/request"
	"go.opentelemetry.io/collector/exporter/exporterhelper/internal/sizer"
)

// batchSizer is an interface that returns the size of the given element.
type batchSizer interface {
	Sizeof(request.Request) int64
}

func newSizer(sizerType sizer.SizerType) batchSizer {
	switch sizerType {
	case sizer.SizerTypeBytes:
		return bytesSizer{}
	case sizer.SizerTypeItems:
		return itemsSizer{}
	default:
		return requestsSizer{}
	}
}

// requestsSizer is a batchSizer implementation that returns the size of a queue element as one request.
type requestsSizer struct{}

func (r requestsSizer) Sizeof(_ request.Request) int64 {
	return 1
}

type itemsSizer struct{}

func (i itemsSizer) Sizeof(r request.Request) int64 {
	return int64(r.ItemsCount())
}

type bytesSizer struct{}

func (b bytesSizer) Sizeof(r request.Request) int64 {
	return int64(r.BytesSize())
}
