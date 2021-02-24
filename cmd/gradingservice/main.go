package main

import (
	"fmt"

	"github.com/zechtz/distributed/grades"
	"github.com/zechtz/distributed/registry"
	"github.com/zechtz/distributed/service"
	"golang.org/x/net/context"

	stlog "log"
)

func main() {
	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var r registry.Registration
	r.ServiceName = registry.GradingService
	r.ServiceURL = serviceAddress

	ctx, err := service.Start(context.Background(), r, host, port, grades.RegisterHandlers)

	if err != nil {
		stlog.Fatal(err)
	}

	<-ctx.Done()
	fmt.Println("Shutting down grading server...")
}
