package cmd

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/spf13/cobra"
	"github.com/tinyhole/im/idl/mua/im/sequence"
	"github.com/tinyhole/im/sequence/application"
	"github.com/tinyhole/im/sequence/infrastructure/config"
	"github.com/tinyhole/im/sequence/infrastructure/driver/mongo"
	"github.com/tinyhole/im/sequence/infrastructure/logger"
	"github.com/tinyhole/im/sequence/infrastructure/repository"
	"github.com/tinyhole/im/sequence/infrastructure/server"
	"github.com/tinyhole/im/sequence/interfaces/rpc"
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
			fmt.Printf("version: %d.%d.%d.%d\n", MAJOR, MINOR, PATCH, BUILD)
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
	//config
	mustSccProvider(container, config.NewBaseConfig)
	mustSccProvider(container, config.NewConfig)
	mustSccProvider(container, server.NewRPCServer)
	//db repository
	mustSccProvider(container, mongo.NewMgoSession)
	mustSccProvider(container, repository.NewMgoDB)
	mustSccProvider(container, repository.NewAutoIncrRepo)
	//logger
	mustSccProvider(container, logger.NewLogger)
	//domain service
	//application service
	mustSccProvider(container, application.NewAppService)

	return container
}

func Run() {
	var (
		err error
	)
	container := buildContainer()
	err = container.Invoke(func(microSvc micro.Service,
		appService *application.AppService, log logger.Logger) {
		microSvc.Init()
		sequence.RegisterSequenceHandler(microSvc.Server(), rpc.NewHandler(appService))
		microSvc.Run()
	})
	if err != nil {
		fmt.Printf("Invoke error [%v]", err)
	}
}
