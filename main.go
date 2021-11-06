package main

import (
	"log"
	api "module_name_daytime/kitex_gen/api/echo"
)

func main() {
	svr := api.NewServer(new(EchoImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
