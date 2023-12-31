package logger

import (
	"fmt"
	"strings"
	"time"
)

const (
	TrafficTypAccess  TrafficTyp = "recv_at"
	TrafficTypRequest TrafficTyp = "sent_to"
)

type TrafficTyp string

// Traffic is provided by user when logging
type Traffic struct {
	Typ  TrafficTyp    // Typ: type of traffic, access or request
	Cmd  string        // Cmd: command
	Code int64         // Code: error code
	Msg  string        // Msg: error message if you have
	Cost time.Duration // Cost: elapse of processing
	Req  any
	Resp any
}

type TrafficEntry interface {
	// Data logs traffic
	Data(traffic *Traffic)
	// DataWith logs traffic with fields
	DataWith(traffic *Traffic, fields Fields)
	// WithFields adds fields to traffic dataLogger
	WithFields(fields Fields) TrafficEntry
	// WithTracing adds requestId to traffic dataLogger
	WithTracing(requestId string) TrafficEntry
	// WithIgnores adds ignores to traffic dataLogger
	WithIgnores(ignores ...string) TrafficEntry
}

func copyFields(fields Fields) Fields {
	if len(fields) == 0 {
		return map[string]any{}
	}
	mapCopy := make(map[string]any, len(fields))
	for k, v := range fields {
		mapCopy[k] = v
	}
	return mapCopy
}

// convertToMessage converts a Traffic to a string
func convertToMessage(tb *Traffic, separator string) string {
	if tb == nil {
		return ""
	}
	if tb.Typ == "" {
		tb.Typ = defaultFieldOccupied
	}
	if tb.Msg == "" {
		tb.Msg = defaultFieldOccupied
	}
	if tb.Cmd == "" {
		tb.Cmd = defaultFieldOccupied
	}
	return strings.Join(append([]string{
		string(tb.Typ),
		tb.Cmd,
		fmt.Sprintf("%dms", tb.Cost.Milliseconds()),
		fmt.Sprintf("%d", tb.Code),
		tb.Msg,
	}), separator)
}
