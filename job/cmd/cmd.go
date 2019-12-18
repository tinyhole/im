package cmd

import (
	"fmt"
	"github.com/micro/go-micro"
	"github.com/spf13/cobra"
	"github.com/tinyhole/im/idl/mua/im/job"
	"github.com/tinyhole/im/job/application/event"
	"github.com/tinyhole/im/job/application/service"
	dsvc "github.com/tinyhole/im/job/domain/service"
	"github.com/tinyhole/im/job/domain/util"
	"github.com/tinyhole/im/job/infrastructure/config"
	"github.com/tinyhole/im/job/infrastructure/driver/eventbus/nsq"
	"github.com/tinyhole/im/job/infrastructure/driver/mongo"
	"github.com/tinyhole/im/job/infrastructure/gateway"
	"github.com/tinyhole/im/job/infrastructure/logger"
	"github.com/tinyhole/im/job/infrastructure/repository"
	"github.com/tinyhole/im/job/infrastructure/server"
	"github.com/tinyhole/im/job/interfaces/rpc"
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
	mustSccProvider(container, mongo.NewMgoSession)
	//repository
	mustSccProvider(container, repository.NewMgoDB)
	mustSccProvider(container, repository.NewInboxRepo)

	//event bus
	mustSccProvider(container, nsq.NewManager)

	//gateway
	mustSccProvider(container, gateway.NewApClient)
	mustSccProvider(container, gateway.NewSequenceClient)
	mustSccProvider(container, gateway.NewSessionClient)
	mustSccProvider(container, gateway.NewRelationClient)
	mustSccProvider(container, util.NewObjectIDCli)
	mustSccProvider(container, util.NewInboxIDClient)
	//event handler
	mustSccProvider(container, event.NewMsgHandler)
	//service
	mustSccProvider(container, dsvc.NewJobService)
	mustSccProvider(container, event.NewEventService)
	mustSccProvider(container, service.NewAppService)
	//server
	mustSccProvider(container, server.NewRPCServer)
	return container
}

func Run() {
	container := buildContainer()
	err := container.Invoke(func(microSvc micro.Service, appService *service.AppService,
		eventService *event.EventService) {
		eventService.Run()
		microSvc.Init()
		job.RegisterJobHandler(microSvc.Server(), rpc.NewHandler(appService))
		microSvc.Run()
		eventService.Stop()
	})

	if err != nil {
		fmt.Printf("%v", err)
	}
}
