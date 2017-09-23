package main

import (
	"bytes"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"net/http"
)

func main() {
	argsLength := len(os.Args)
	if argsLength  != 2 && argsLength != 3 {
		log.Fatalf("Invalid number of arguments\n%s", formatUsage())
	}

	listener, err := StartServer()
	if err != nil {
		panic(err)
	}

	port, err := getPort(listener.Addr().String())
	if err != nil {
		panic(err)
	}

	dockerNode, err := NewDockerNode(port)
	if err != nil {
		panic(err)
	}

	startRegistrationTicker(os.Args[1], dockerNode)

	panic(http.Serve(listener, nil))
}

func getPort(address string) (int, error) {
	slice := strings.Split(address, ":")

	return strconv.Atoi(slice[len(slice)-1])
}

func startRegistrationTicker(xcServerBaseUrl string, dockerNode *DockerNode) {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				_, err := Register(xcServerBaseUrl, dockerNode)
				if err != nil {
					log.Println("Error : %s", err)
				} else {
					log.Println("Registered docker node to xc server.")

				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func formatUsage() (usage string) {
	var buffer bytes.Buffer
	buffer.WriteString("\n")
	buffer.WriteString("Usage : xcagent XC-SERVER-URL [NODE-HOSTNAME]\n")
	buffer.WriteString("\n")

	return buffer.String()
}
