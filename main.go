package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"vandyahmad/skyshi/config"
	"vandyahmad/skyshi/helper"
	rtr "vandyahmad/skyshi/router"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Skyshy Test By Vandy Ahmad")
	config.InitDB()

	router := gin.Default()
	router.NoRoute(func(ctx *gin.Context) {
		response := helper.ApiResponse("Failed", "Error", nil)
		ctx.JSON(404, response)
	})
	rtr.ActivityRouter(router)
	rtr.TodoRouter(router)
	go func() {
		router.Run(":3030")
	}()
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	signal := <-c
	log.Fatalf("Proses Selesai dengan signal: %v\n", signal.String())
}
