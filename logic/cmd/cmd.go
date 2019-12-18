package cmd

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/spf13/cobra"
	"github.com/tinyhole/im/idl/mua/im/logic"
	eventbus2 "github.com/tinyhole/im/logic/application/eventbus"
	"github.com/tinyhole/im/logic/application/service"
	dsvc "github.com/tinyhole/im/logic/domain/service"
	"github.com/tinyhole/im/logic/infrastructure/config"
	"github.com/tinyhole/im/logic/infrastructure/driver/eventbus/nsq"
	"github.com/tinyhole/im/logic/infrastructure/driver/redis"
	"github.com/tinyhole/im/logic/infrastructure/gateway"
	"github.com/tinyhole/im/logic/infrastructure/logger"
	"github.com/tinyhole/im/logic/infrastructure/repository/sessionstate"
	"github.com/tinyhole/im/logic/infrastructure/server"
	"github.com/tinyhole/im/logic/interfaces/rpc"
	"go.uber.org/dig"
	"os"
)

var (
	RootCmd = &cobra.Command{
		Short: "run server",
		Run: func(cmd *cobra.Command, args []string) {
			Run()
		},
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "show version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("version: %d.%d.%d.%d", MAJOR, MINOR, PATCH, BUILD)
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
		os.Exit(1)
	}
}

func buildContainer() *dig.Container {
	container := dig.New()
	mustSccProvider(container, config.NewBaseConfig)
	mustSccProvider(container, config.NewConfig)
	mustSccProvider(container, logger.NewLogger)

	mustSccProvider(container, redis.NewRedisPool)
	mustSccProvider(container, sessionstate.NewSessionStateRepo)

	mustSccProvider(container, nsq.NewManager)

	mustSccProvider(container, gateway.NewAuthClient)

	mustSccProvider(container, dsvc.NewSessionService)
	mustSccProvider(container, service.NewAppService)
	mustSccProvider(container, server.NewRPCServer)
	return container
}

func Run() {
	var (
		err error
	)
	container := buildContainer()
	err = container.Invoke(func(microSvc micro.Service, appService *service.AppService, eventMgr eventbus2.Manager) {
		microSvc.Init()
		logic.RegisterLogicHandler(microSvc.Server(), rpc.NewHandler(appService))
		err := eventMgr.Run()
		if err != nil {
			return
		}
		microSvc.Run()
		eventMgr.Stop()
	})
	fmt.Printf("%v", err)
}
