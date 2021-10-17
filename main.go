package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	f, err := os.OpenFile("logs/info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("⚠️  Error opening file: " + err.Error())
	}
	defer f.Close()

	infoLog := log.New(f, "INFO: ", log.Ldate|log.Ltime)

	go serve(infoLog)

	// Create a new channel that keeps running until Ctrl+C is clicked.
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel

	fmt.Printf("\n❌ Stop...\n")
}

func serve(logger *log.Logger) {
	for {
		logger.Printf("📡 Server is running")
		time.Sleep(time.Minute)
	}
}
