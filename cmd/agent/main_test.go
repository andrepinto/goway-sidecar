package main

import (
	"github.com/andrepinto/goway-sidecar/proto"
	"google.golang.org/grpc"

	"testing"
	"golang.org/x/net/context"
	"time"
	"fmt"
)

func Test_Main(t *testing.T){

	conn, err := grpc.Dial(":5000", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	navyGrpcClient := proto.NewGoWayCollectAgentServiceClient(conn)

	data := &proto.HttpLoggerRequest{
		Key:         "entry.ServiceName",
		Version:     "entry.Version",
		BasePath:    "entry.BasePath",
		ElapsedTime: float32(300 / time.Millisecond),
		Ip:          "entry.Ip",
		Host:        "entry.Host",
		Uri:         "entry.Uri",
		Time: &proto.DateTime{
			Year:  fmt.Sprintf("%d", 2018),
			Month: fmt.Sprintf("%d", 1),
			Day:   fmt.Sprintf("%d", 1),
			Hour:  fmt.Sprintf("%d", 1),
			Min:   fmt.Sprintf("%d", 1),
			Sec:   fmt.Sprintf("%d", 1),
		},
		Protocol: "entry.Protocol",
		Method:   "entry.Method",
		RequestId: "reqId",
		Properties: map[string]string{
			"service-gateway":   "main",
		},
		Tags:[]string{"gateway", "service-gateway", "entry.Product", "entry.Client", "entry.ServiceName"},
		Status: fmt.Sprintf("%d","entry.Status"),
		ResponseBody: []byte("entry.ResBody"),
		RequestBody: []byte("entry.ReqBody"),
		RequestHeader: []string{"requestHeaders"},
		ServicePath: "entry.ServicePath",
		Metadata: map[string]string{"product": "entry.Product", "client":"entry.Client", "version":"entry.Version"},
	}

	_, err = navyGrpcClient.HttpLogger(context.Background(),data)
}
