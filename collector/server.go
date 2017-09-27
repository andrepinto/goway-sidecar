package collector


import (
	"net"
	log "github.com/sirupsen/logrus"
	"github.com/andrepinto/goway-sidecar/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/api/support/bundler"
	"github.com/andrepinto/goway-sidecar/outputs"
)

type CollectorRpcServer struct {
	Port string
	Action *HttpLoggerAction
	htppLoggerBundler *bundler.Bundler
	Context map[string]string
}

type CollectorRpcServerOptions struct {
	Port string
	Output outputs.Output
	DelayThreshold int
	BundleCountThreshold int
	Context map[string]string
}


func NewCollectorRpcServer(options *CollectorRpcServerOptions) *CollectorRpcServer{
	if options.Port == "" {
		panic("port not found")
	}

	action := &HttpLoggerAction{}
	action.htppLoggerBundler = initHttpLoggerBundler(action.sendHttpLogger, options.DelayThreshold, options.BundleCountThreshold)
	action.Output = options.Output

	return &CollectorRpcServer{
		Port: options.Port,
		Action: action,
		Context: options.Context,
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

	item := &HttpLogger{
		Context: sv.Context,
		Data:in,
		Properties: in.Properties,
	}

	sv.Action.Fire(item)

	return &proto.HttpLoggerResponse{}, nil
}

