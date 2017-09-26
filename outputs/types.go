package outputs

import "time"

type BaseClientRequest struct {
	Context    map[string]string        `json:"context,omitempty"`
	Properties map[string]string        `json:"properties,omitempty"`
	Key        string            `json:"key,omitempty"`
	Id         string            `json:"id,omitempty"`
	Timestamp  string            `json:"time,omitempty"`
}

type HttpLoggerClient struct {
	RequestId string    `json:"request_id,omitempty"`
	Base      BaseClientRequest        `json:"base,omitempty"`
	Data      HttpLoggerRequestClient `json:"data,omitempty"`
	Timestamp string            `json:"time,omitempty"`
}

type HttpLoggerRequestClient struct {
	BasePath      string    `json:"base_path,omitempty"`
	ElapsedTime   float32   `json:"elapsed_time,omitempty"`
	Host          string    `json:"host,omitempty"`
	Ip            string    `json:"ip,omitempty"`
	Method        string    `json:"method,omitempty"`
	RequestBody   []byte    `json:"request_body,omitempty"`
	RequestHeader []string   `json:"request_header,omitempty"`
	ResponseBody  []byte    `json:"response_body,omitempty"`
	Version       string    `json:"version,omitempty"`
	Uri           string    `json:"uri,omitempty"`
	Protocol      string    `json:"protocol,omitempty"`
	Time          time.Time `json:"time,omitempty"`
	Tags          []string  `json:"tags,omitempty"`
	Status        string    `json:"status,omitempty"`
	ServicePath   string    `json:"service_path"`
	Metadata      map[string]string `json:"metadata"`
}
