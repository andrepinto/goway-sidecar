package collector


import (
	"net"
	log "github.com/sirupsen/logrus"
	"github.com/andrepinto/goway-sidecar/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/andrepinto/goway-sidecar/outputs"
	"time"
	"strconv"
)

type CollectorRpcServer struct {
	Port string
	Output outputs.Output
}

type CollectorRpcServerOptions struct {
	Port string
	Output outputs.Output
}


func NewCollectorRpcServer(options *CollectorRpcServerOptions) *CollectorRpcServer{
	if options.Port == "" {
		panic("port not found")
	}
	return &CollectorRpcServer{
		Port: options.Port,
		Output: options.Output,
	}
}


func (sv *CollectorRpcServer) Start() error {
	lis, err := net.Listen("tcp", sv.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}
	// Creates a new gRPC server
	s := grpc.NewServer()
	proto.RegisterGoWayCollectAgentServiceServer(s, sv)
	go s.Serve(lis)
	return nil
}


func (sv *CollectorRpcServer) HttpLogger(context context.Context, in *proto.HttpLoggerRequest) (*proto.HttpLoggerResponse, error) {
	log.Info(in)

	item := HttpLogger{
		Context: map[string]string{
			"service":   "sv.ServiceConfig.Id",
			"version":    "sv.ServiceConfig.Version",
			"service_id":    "sv.ServiceId",
			"environment": "sv.Env",
		},
		Data:in,
		Properties: in.Properties,
	}

	arr := []*outputs.HttpLoggerClient{}


		date, _ := DateTimeToDate(item.Data.Time.Year, item.Data.Time.Month, item.Data.Time.Day, item.Data.Time.Hour, item.Data.Time.Min, item.Data.Time.Sec)


		tr := &outputs.HttpLoggerClient{
			RequestId: item.Data.RequestId,
			Base: outputs.BaseClientRequest{
				Properties: item.Properties,
				Id: item.Key,
				Key: item.Key,
				Context: item.Context,
			},
			Data: outputs.HttpLoggerRequestClient{
				Protocol:item.Data.Protocol,
				Uri: item.Data.Uri,
				Host: item.Data.Host,
				Ip: item.Data.Ip,
				ElapsedTime: item.Data.ElapsedTime,
				BasePath: item.Data.BasePath,
				RequestHeader: item.Data.RequestHeader,
				RequestBody: item.Data.RequestBody,
				Version: item.Data.Version,
				Method: item.Data.Method,
				Time: date,
				ResponseBody: item.Data.ResponseBody,
				Tags: item.Data.Tags,
				Status: item.Data.Status,
				ServicePath: item.Data.ServicePath,
				Metadata: item.Data.Metadata,
			},
			Timestamp: item.Timestamp,
		}
		arr = append(arr, tr)




	sv.Output.Send(arr)

	return &proto.HttpLoggerResponse{}, nil
}

func DateTimeToDate(year string, month string, day string, hour string, min string, sec string) (time.Time, error){
	iy, err :=strconv.Atoi(year)
	im, err :=strconv.Atoi(month)
	id, err :=strconv.Atoi(day)
	ih, err :=strconv.Atoi(hour)
	imin, err :=strconv.Atoi(min)
	is, err :=strconv.Atoi(sec)

	if err!=nil{
		return time.Now(), err
	}

	date := time.Date(iy, time.Month(im), id, ih, imin, is, 0, time.UTC)
	return date, nil
}