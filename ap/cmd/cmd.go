package cmd

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	microSrv "github.com/micro/go-micro/server"
	"github.com/micro/go-plugins/transport/tcp"
	"github.com/spf13/cobra"
	"github.com/tinyhole/im/ap/config"
	"github.com/tinyhole/im/ap/logger"
	"github.com/tinyhole/im/ap/rpc"
	"github.com/tinyhole/im/ap/tcpserver/client"
	"github.com/tinyhole/im/ap/tcpserver/gateway"
	"github.com/tinyhole/im/ap/tcpserver/server"
	"github.com/tinyhole/im/idl/mua/im/ap"
	"go.uber.org/dig"
	"os"
)

var (
	RootCmd = &cobra.Command{
		Short: "run server",
		Long:  "run server",
		Run: func(cmd *cobra.Command, args []string) {
			run()
			return
		},
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "print version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(fmt.Sprintf("version: %d.%d.%d.%d", MAJOR, MINOR, PATCH, BUILD))
			return
		},
	}
)

func Execute() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.Execute()
}

func mustSccProvider(c *dig.Container, constructor interface{}, opts ...dig.ProvideOption) {
	var (
		err error
	)
	err = c.Provide(constructor, opts...)
	if err != nil {
		fmt.Printf("dig provider err [%v]", err)
		os.Exit(1)
	}
}

func buildContainer() *dig.Container {
	container := dig.New()
	mustSccProvider(container, config.NewBaseConfig)
	mustSccProvider(container, config.NewConfig)
	mustSccProvider(container, logger.NewLogger)
	mustSccProvider(container, gateway.NewAuthClient)
	mustSccProvider(container, gateway.NewSessionClient)
	return container
}

func run() {
	var (
		err error
	)
	container := buildContainer()
	err = container.Invoke(func(baseConf *config.BaseConfig, srvConf *config.Config,
		log logger.Logger, authClient gateway.AuthClient, sessionClient gateway.SessionClient) {

		rpcSrv := microSrv.NewServer(microSrv.Name("mua.im.ap"), microSrv.Id(fmt.Sprintf("%d", baseConf.SrvID)))

		rpcSvc := micro.NewService(micro.Server(rpcSrv), micro.Transport(tcp.NewTransport()), micro.Name("mua.im.ap"),
			micro.Registry(etcd.NewRegistry(registry.Addrs(baseConf.RegistryCenterAddr))))

		tcpSrv := server.NewAPServer(log, server.WithLocalAddr(fmt.Sprintf(":%d", srvConf.ApPort)),
			server.WithSrvID(baseConf.SrvID), server.WithAuthClient(authClient), server.WithSessionClient(sessionClient))

		tcpSrv.Start()

		rpcSrv.Init()
		ap.RegisterAPHandler(rpcSvc.Server(), rpc.NewHandler(client.APClient))
		rpcSvc.Run()

	})
	if err != nil {
		fmt.Printf("Invoke error [%v]", err)
	}
}



