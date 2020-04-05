package app

import (
	"fmt"

	"github.com/gyan1230/2020/docker-kubernetes/server/controller"
)

func mapurls() {
	fmt.Println("url mapping...")
	router.GET("/home", controller.Home)
	router.GET("/send", controller.SendInfo)
}
