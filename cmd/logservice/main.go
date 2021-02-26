package main

import (
	"fmt"
	stlog "log"

	"context"

	"github.com/zechtz/distributed/log"
	"github.com/zechtz/distributed/registry"
	"github.com/zechtz/distributed/service"
)

func main() {
	log.Run("./app.log")

	host, port := "localhost", "4000"

	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var r registry.Registration
	r.ServiceName = registry.LogService
	r.ServiceURL = serviceAddress
	r.RequiredServices = make([]registry.ServiceName, 0)
	r.ServiceUpdateURL = r.ServiceURL + "/services"
	r.HeartbeatURL = r.ServiceURL + "/heartbeat"

	ctx, err := service.Start(context.Background(), r, host, port, log.RegisterHandlers)

	if err != nil {
		stlog.Fatal(err)
	}

	<-ctx.Done()

	fmt.Println("Shutting down log service...")
}
