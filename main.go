package main

import (
	"log"
	api "module_name_daytime/kitex_gen/api/servicenamedaytimeinthriftfile"
)

func main() {
	svr := api.NewServer(new(ServiceNameDaytimeInThriftFileImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
