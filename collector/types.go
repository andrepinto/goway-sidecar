package collector

import "github.com/andrepinto/goway-sidecar/proto"

type HttpLogger struct {
	Context      map[string]string      `json:"context,omitempty"`
	Properties   map[string]string      `json:"properties,omitempty"`
	Base
	Data         *proto.HttpLoggerRequest `json:"data,omitempty"`
	callback     HttpLoggerCallback
}

type HttpLoggerCallback func(logger *HttpLogger)


type Base struct {
	Type      string `json:"type,omitempty"`
	Id 	  string `json:"id,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	SentAt    string `json:"sent_at,omitempty"`
	Key       string `json:"key,omitempty"`
}


type base interface {
	setId(string)
	setTimestamp(string)
}

