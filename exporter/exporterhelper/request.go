// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package exporterhelper // import "go.opentelemetry.io/collector/exporter/exporterhelper"

import (
	"go.opentelemetry.io/collector/exporter/exporterhelper/internal/sizer"
)

type RequestSizerType = sizer.SizerType

var (
	RequestSizerTypeBytes    = sizer.SizerTypeBytes
	RequestSizerTypeItems    = sizer.SizerTypeItems
	RequestSizerTypeRequests = sizer.SizerTypeRequests
)
