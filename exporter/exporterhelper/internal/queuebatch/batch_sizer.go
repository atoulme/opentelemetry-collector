package queuebatch

import (
	"go.opentelemetry.io/collector/exporter/exporterhelper/internal/request"
	"go.opentelemetry.io/collector/exporter/exporterhelper/internal/sizer"
)

// batchSizer is an interface that returns the size of the given element.
type batchSizer interface {
	Sizer() sizer.Sizer
	Sizeof(request.Request) int64
}

func newSizer(sizerType sizer.SizerType) batchSizer {
	switch sizerType {
	case sizer.SizerTypeBytes:
		return bytesSizer{
			sizer.Sizer{
				SizerType:     sizer.SizerTypeBytes,
				LogsSizer:     &sizer.LogsBytesSizer{},
				MetricsSizer:  &sizer.MetricsBytesSizer{},
				ProfilesSizer: &sizer.ProfilesBytesSizer{},
				TracesSizer:   &sizer.TracesBytesSizer{},
			},
		}
	case sizer.SizerTypeItems:
		return itemsSizer{
			sizer.Sizer{
				SizerType:     sizer.SizerTypeItems,
				LogsSizer:     &sizer.LogsCountSizer{},
				MetricsSizer:  &sizer.MetricsCountSizer{},
				ProfilesSizer: &sizer.ProfilesCountSizer{},
				TracesSizer:   &sizer.TracesCountSizer{},
			},
		}
	default:
		return requestsSizer{
			sizer.Sizer{
				SizerType:     sizer.SizerTypeRequests,
				LogsSizer:     &sizer.LogsCountSizer{},
				MetricsSizer:  &sizer.MetricsCountSizer{},
				ProfilesSizer: &sizer.ProfilesCountSizer{},
				TracesSizer:   &sizer.TracesCountSizer{},
			},
		}
	}
}

// requestsSizer is a batchSizer implementation that returns the size of a queue element as one request.
type requestsSizer struct {
	s sizer.Sizer
}

func (r requestsSizer) Sizer() sizer.Sizer {
	return r.s
}

func (r requestsSizer) Sizeof(_ request.Request) int64 {
	return 1
}

type itemsSizer struct {
	s sizer.Sizer
}

func (i itemsSizer) Sizer() sizer.Sizer {
	return i.s
}

func (i itemsSizer) Sizeof(r request.Request) int64 {
	return int64(r.ItemsCount())
}

type bytesSizer struct {
	s sizer.Sizer
}

func (b bytesSizer) Sizer() sizer.Sizer {
	return b.s
}

func (b bytesSizer) Sizeof(r request.Request) int64 {
	return int64(r.BytesSize())
}
