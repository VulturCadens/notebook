package main

import (
	"fmt"
	"log"
	"log/syslog"
	"os"
)

func main() {
	logger, err := syslog.NewLogger(
		syslog.LOG_NOTICE|syslog.LOG_LOCAL0, // https://golang.org/pkg/log/syslog/#Priority
		log.Ldate|log.Lshortfile,            // https://golang.org/pkg/log/#pkg-constants
	)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	logger.Print("Lorem ipsum dolor sit amet.")

	// tail /var/log/syslog
	// journalctl -n 5
}
