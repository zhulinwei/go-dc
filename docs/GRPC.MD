# Go-dc

## GRPC

一般情况下项目即可以作为服务端对外提供rpc调用服务，也可以作为客户端调用其他rpc服务，本项目将分别对这两种情况给出对应的实践方式

### 作为服务端对外提供服务

#### 定义协议
我们将定义的协议放在[pkg/rpc/protobuf](https://github.com/zhulinwei/go-dc/blob/master/pkg/rpc/protobuf/greeter.proto)文件夹下

```proto
syntax = "proto3";

package protobuf;

message HelloRequest {
    string name = 1;
}

message HelloReply {
    string message = 1;
}

service Greeter {
    rpc SayHello (HelloRequest) returns (HelloReply) {}
}

```

使用脚手架protoc生成协议对应的go文件
```shell
protoc --go_out=plugins=grpc:. *.proto
```

生成的go文件详见[pkg/rpc/protobuf/greeter.pb.go](https://github.com/zhulinwei/go-dc/blob/master/pkg/rpc/protobuf/greeter.pb.go)

#### 注册服务

```go
func GRPCRun(port string) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	protobuf.RegisterGreeterServer(server, &GreeterServer{})
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```

详见[pkg/rpc/server](https://github.com/zhulinwei/go-dc/blob/master/pkg/rpc/server.go)

#### 启动rpc服务

```go
func main() {
	httpPort := config.ServerConfig().HttpPort
	grpcPort := config.ServerConfig().GrpcPort

	go rpc.GRPCRun(grpcPort)
	if err := router.BuildRoute().Run(httpPort); err != nil {
		log.Fatalf("server run failed: %v", err)
	}
}

```
详见项目启动文件[cmd/main.go](https://github.com/zhulinwei/go-dc/blob/master/cmd/main.go)，这里需要注意如果rpc服务与普通web服务放在同一个项目里，则需要需要开启go程

### 作为客户端调用其他服务

#### 定义配置文件

注意需要考虑调用多个rpc服务的情况
```yaml
...
grpc:
  - name: grpc1
    addr: localhost:8080
  - name: grpc2
    addr: localhost:8081
...
```

详见[configs/config.yaml](https://github.com/zhulinwei/go-dc/blob/master/configs/config.yaml)

#### 连接rpc服务

```go

type IGrpc interface {
	Client1() *grpc.ClientConn
	Client2() *grpc.ClientConn
}

type Grpc struct {
	ClientMap map[string]*grpc.ClientConn
}

func BuildGrpc() IGrpc {
	return Grpc{
		ClientMap: grpcClientMap,
	}
}

func (grpc Grpc) Client1() *grpc.ClientConn {
	return grpc.ClientMap["grpc1"]
}

func (grpc Grpc) Client2() *grpc.ClientConn {
	return grpc.ClientMap["grpc2"]
}

func init() {
	grpcConfigs := config.ServerConfig().Grpc
	
	grpcOnce.Do(func() {
		grpcMutex.Lock()
		defer grpcMutex.Unlock()

		grpcClientMap = make(map[string]*grpc.ClientConn, len(grpcConfigs))
		for _, grpcConfig := range grpcConfigs {
			var err error
			var client *grpc.ClientConn
			if client, err = grpc.Dial(grpcConfig.Addr, grpc.WithInsecure()); err != nil {
				log.Fatal("grpc dial fail: %v", err)
				return
			}
			grpcClientMap[grpcConfig.Name] = client
		}
	})
}
```

连接详情见[pkg/rpc/client.go](https://github.com/zhulinwei/go-dc/blob/master/pkg/rpc/client.go)，完成rpc的连接后，后面可以正常使用了

### 调用服务

```go
type IGreeter interface {
	QueryGreeterFromGrpc(name string) (*greeterPb.HelloReply, error)
}

type GreeterClient struct {
	Client *grpc.ClientConn
}

func BuildGreeterService() IGreeter {
	return GreeterClient{
		Client: rpc.BuildGrpc().Client1(),
	}
}

func (greeter GreeterClient) QueryGreeterFromGrpc(name string) (*greeterPb.HelloReply, error) {
	greeterClient := greeterPb.NewGreeterClient(greeter.Client)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return greeterClient.SayHello(ctx, &greeterPb.HelloRequest{Name: name})
}

```

详见[pkg/service/greeter.go](https://github.com/zhulinwei/go-dc/blob/master/pkg/service/greeter.go)
