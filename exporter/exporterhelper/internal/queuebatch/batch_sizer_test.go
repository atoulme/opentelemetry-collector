package queuebatch

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/collector/exporter/exporterhelper/internal/requesttest"
)

func TestItemsSizer(t *testing.T) {
	sz := itemsSizer{}
	assert.EqualValues(t, 3, sz.Sizeof(&requesttest.FakeRequest{Items: 3}))
}
