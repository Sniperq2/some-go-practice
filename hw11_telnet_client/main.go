package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	flag "github.com/spf13/pflag"
)

var (
	timeout string
	wg      *sync.WaitGroup
)

func readerLoop(t *TelnetClient, done chan os.Signal) {
	defer wg.Done()
	for {
		select {
		case d := <-done:
			if d == syscall.SIGQUIT {
				(*t).Close()
				log.Println("Bye-bye")
				return
			}
			log.Println("Bye-bye")
			return
		default:
			err := (*t).Receive()
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func writerLoop(t *TelnetClient, done chan os.Signal) {
	defer wg.Done()
	for {
		select {
		case d := <-done:
			if d == syscall.SIGQUIT {
				(*t).Close()
				log.Println("Bye-bye")
				return
			}
			log.Println("Bye-bye")
			return
		default:
			err := (*t).Send()
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func main() {
	flag.StringVar(&timeout, "timeout", "10s", "timeout")
	flag.Parse()
	host := os.Args[2]
	port := os.Args[3]
	if len(host) == 0 {
		log.Fatalf("Please set host")
	}
	if len(port) == 0 {
		log.Fatalf("Please set port")
	}

	durationTimeout, err := time.ParseDuration(timeout)
	if err != nil {
		log.Fatalf("wrong duration")
	}

	telnetClient := NewTelnetClient(net.JoinHostPort(host, port), durationTimeout, os.Stdin, os.Stdout)
	if err := telnetClient.Connect(); err != nil {
		log.Fatalln(err)
	}

	notifySignal := make(chan os.Signal, 1)
	signal.Notify(notifySignal, syscall.SIGINT, syscall.SIGQUIT)

	wg = &sync.WaitGroup{}
	wg.Add(2)

	go readerLoop(&telnetClient, notifySignal)
	go writerLoop(&telnetClient, notifySignal)

	go func() {
		for {
			switch <-notifySignal {
			case syscall.SIGINT:
				return

			case syscall.SIGQUIT:
				return
			}
		}
	}()
	wg.Wait()

}
