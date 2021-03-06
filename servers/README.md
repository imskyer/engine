# servers
--
    import "github.com/autom8ter/engine/servers"


## Usage

#### func  NewGrpcServer

```go
func NewGrpcServer(c *config.Config) driver.Server
```
NewGrpcServer creates a new GrpcServer instance.

#### type GrpcServer

```go
type GrpcServer struct {
}
```

GrpcServer wraps grpc.Server setup process.

#### func (*GrpcServer) Serve

```go
func (s *GrpcServer) Serve(lis net.Listener) error
```
Serve implements Server.Serve for starting the grpc server

#### func (*GrpcServer) Shutdown

```go
func (s *GrpcServer) Shutdown()
```
Shutdown implements Server.Shutdown for gracefully shutting down the grpc server
