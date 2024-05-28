package main

import (
	"go-whatsapp-bot/api"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	wac, err := api.CreateClient()
	if err != nil {
		log.Fatalf("Failed to create client: %v\n", err)
	}

	api.ConnectClient(wac)
	wac.AddEventHandler(api.HandleEvent)

	err = wac.Connect()
	if err != nil {
		panic(err)
	}

	// Listen to Ctrl+C (you can also do something else that prevents the program from exiting)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	wac.Disconnect()
}
