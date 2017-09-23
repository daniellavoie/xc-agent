package main

import "os"

type DockerNode struct {
	Hostname string `json:"hostname"`
	Port     int `json:"port"`
}

func NewDockerNode(port int) (*DockerNode, error) {
	if(len(os.Args) == 3) {
		return &DockerNode{Hostname: os.Args[2], Port: port}, nil
	} else {
		hostname, err := os.Hostname()
		if err != nil {
			return nil, err
		}

		return &DockerNode{Hostname: hostname, Port: port}, nil
	}
}
