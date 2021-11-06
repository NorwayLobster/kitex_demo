package main

import (
	"context"
	"fmt"
	"log"
	"module_name_daytime/kitex_gen/api"
	"module_name_daytime/kitex_gen/api/servicenamedaytimeinthriftfile"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

func main() {
	c, err := servicenamedaytimeinthriftfile.NewClient("example", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}

	clientTime := time.Now().Unix()
	req := &api.DaytimeRequest{ClientDaytime: fmt.Sprintf("%v", clientTime)}
	resp, err := c.Daytime(context.Background(), req, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
