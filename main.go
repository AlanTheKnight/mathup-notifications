package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	// Loggers
	infoLog  *log.Logger
	errorLog *log.Logger
)

func main() {
	f, err := os.OpenFile("logs/info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic("⚠️ Error opening file: " + err.Error())
	}
	defer f.Close()

	infoLog = log.New(f, "INFO: ", log.Ldate|log.Ltime)
	errorLog = log.New(f, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	RetrieveTokens()

	go serve()

	// Create a new channel that keeps running until Ctrl+C is clicked.
	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel

	infoLog.Printf("🛑 Server was stopped")
	fmt.Printf("\n❌ Stop...\n")
}

func serve() {
	for {
		infoLog.Printf("📡 Server is running")
		time.Sleep(time.Minute)
	}
}
