package main

import (
	"bufio"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"sync"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Printf("Ctrl+c pressed, exiting...")
		os.Exit(0)
	}()

	//	feedback, err := regexp.Compile(`^(1?[0-9]) (.*?)$`)
	feedback, err := regexp.Compile(`^(1?[0-9]) ([[:ascii:]]*?)$`)
	if err != nil {
		log.Fatal(err)
	}

	fifo := os.Args[1]

	out, err := os.OpenFile(fifo, os.O_RDONLY, 0600)
	defer out.Close()
	if err != nil {
		log.Fatal(err)
	}

	mut := &sync.Mutex{}
	size := 128000 // 128kb

	for {
		scanner := bufio.NewScanner(bufio.NewReaderSize(out, size))
		for scanner.Scan() {
			groups := feedback.FindStringSubmatch(scanner.Text())
			if groups != nil && strings.TrimSpace(groups[2]) != "" {
				mut.Lock()
				log.Printf("%s: %s", groups[1], groups[2])
				mut.Unlock()
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}
