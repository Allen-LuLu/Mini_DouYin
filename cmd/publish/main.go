package main

import (
	publish "Mini_DouYin/kitex_gen/publish/publish"
	"log"
)

func main() {

	svr := publish.NewServer(new(PublishImpl))

	err := svr.Run()

	if err != nil {
		log.Panicf(err.Error())
	}

}
