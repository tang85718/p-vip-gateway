package main

import (
	"./api"
	"./server"
	"os"
)

//type Gateway struct{}
//
//func (g *Gateway) Hello(c context.Context, req *pb.HelloReq, rsp *pb.HelloResponse) error {
//	rsp.Result = "Hello " + req.Name
//	return nil
//}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "2000"
	}

	handler := &server.GatewayHandler{}
	processor := api.NewGatewayProcessor(handler)

	server.Serve(processor, port)

	//r := registry.NewRegistry(registry.Addrs("10.0.1.200"))

	//service := micro.NewService(
	//	//micro.Registry(r),
	//	micro.Name("gateway"),
	//	micro.Version("1.0.0"),
	//	micro.RegisterTTL(time.Second*30),
	//	micro.RegisterInterval(time.Second*15),
	//	//micro.Flags(cli.BoolFlag{
	//	//	Name:  "fuck",
	//	//	Usage: "test",
	//	//}),
	//)

	//service.Init()
	//service.Init(
	//	micro.Action(func(c *cli.Context) {
	//		if c.Bool("fuck") {
	//			setupClient(service)
	//			os.Exit(0)
	//		}
	//	}),
	//)

	//pb.RegisterGateWayHandler(service.Server(), new(Gateway))
	//
	//if err := service.Run(); err != nil {
	//	fmt.Println(err)
	//}
}
