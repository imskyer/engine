package engine_test

import (
	"context"
	"fmt"
	"github.com/autom8ter/engine"
	"github.com/autom8ter/engine/config"
	"github.com/autom8ter/engine/examples/examplepb/client"
	"github.com/autom8ter/util"
	"github.com/grpc-ecosystem/grpc-gateway/examples/proto/examplepb"
	"google.golang.org/grpc/grpclog"
	"os"
	"testing"
)

func init() {
	if err := os.Setenv("DEBUG", "t"); err != nil {
		grpclog.Fatalln(err.Error())
	}
}

func TestGRPC(t *testing.T) {
	var eng = engine.New("tcp", ":3002").With(
		config.WithPlugins("Plugin", "bin/example.so"),
	)
	go eng.Serve()
	var grpcCli = client.ExampleClient(":3002")
	resp, err := grpcCli.EchoBody(context.Background(), &examplepb.SimpleMessage{
		Id:  "yoyoyoyoyo",
		Num: 199,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println("GRPC RESPONSE:")
	fmt.Println(util.ToPrettyJsonString(resp))
}
