// Copyright 2019 The go-ethereum Authors
// (original work)
// Copyright 2024 The Erigon Authors
// (modifications)
// This file is part of Erigon.
//
// Erigon is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Erigon is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with Erigon. If not, see <http://www.gnu.org/licenses/>.

package rpc

import (
	"reflect"
	"strings"
	"testing"

	"github.com/erigontech/erigon/common/log/v3"
	"github.com/stretchr/testify/require"
)

func TestCallPanicLogsCorrectly(t *testing.T) {
	fn := reflect.ValueOf(func() {
		panic("simulate panic")
	})
	logger := log.New("test", "TestCallPanicLogsCorrectly")
	logBuffer := new(strings.Builder)
	logger.SetHandler(log.StreamHandler(logBuffer, log.TerminalFormat()))
	c := &callback{fn: fn, rcvr: reflect.ValueOf(nil), errPos: -1, logger: logger}
	_, err := c.call(t.Context(), "test", []reflect.Value{}, nil)
	require.Regexp(t, "method handler crashed", err)

	require.Contains(t, logBuffer.String(), "service_test.go", "Expect to have details about the panic in the log")
}
