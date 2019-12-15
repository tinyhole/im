package cmd

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/spf13/cobra"
	"github.com/tinyhole/im/idl/mua/im/relation"
	"github.com/tinyhole/im/relation/application"
	"github.com/tinyhole/im/relation/domain/service"
	"github.com/tinyhole/im/relation/infrastructure/config"
	"github.com/tinyhole/im/relation/infrastructure/driver/mongo"
	"github.com/tinyhole/im/relation/infrastructure/gateway"
	"github.com/tinyhole/im/relation/infrastructure/logger"
	"github.com/tinyhole/im/relation/infrastructure/repository"
	"github.com/tinyhole/im/relation/infrastructure/server"
	"github.com/tinyhole/im/relation/interfaces/objconv"
	"github.com/tinyhole/im/relation/interfaces/rpc"
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
		os.Exit(1)
	}
}

func buildContainer() *dig.Container {
	container := dig.New()
	//config
	mustSccProvider(container, config.NewBaseConfig)
	mustSccProvider(container, config.NewConfig)
	//db repository
	mustSccProvider(container, mongo.NewMgoSession)
	mustSccProvider(container, repository.NewMgoDB)
	mustSccProvider(container, repository.NewPersonalRelationRepo)
	mustSccProvider(container, repository.NewGroupRelationRepo)
	mustSccProvider(container, repository.NewGroupRepo)

	//logger
	mustSccProvider(container, logger.NewLogger)

	//gateway
	mustSccProvider(container, gateway.NewSequenceClient)

	//objconv
	mustSccProvider(container, objconv.NewGroupConv)
	mustSccProvider(container, objconv.NewGroupRelationConv)
	mustSccProvider(container, objconv.NewPersonalRelationConv)

	//domain service
	mustSccProvider(container, service.NewRelationService)

	//application service
	mustSccProvider(container, application.NewAppService)

	//interface
	mustSccProvider(container, server.NewRPCServer)
	return container
}

func Run() {
	container := buildContainer()
	err := container.Invoke(func(microSvc micro.Service,
		appService *application.AppService,
		groupConv *objconv.GroupConv,
		groupRelationConv *objconv.GroupRelationConv,
		personalRelationConv *objconv.PersonalRelationConv) {
		microSvc.Init()
		relation.RegisterRelationHandler(microSvc.Server(), rpc.NewHandler(appService,
			groupConv, groupRelationConv, personalRelationConv))
		microSvc.Run()
	})
	if err != nil {
		fmt.Printf("invoke error [%v]", err)
	}
}
