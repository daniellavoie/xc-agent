package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func Register(managerBaseUrl string, dockerNode *DockerNode) (response *http.Response, err error) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(dockerNode)

	return http.Post(fmt.Sprintf("%s/docker-node", managerBaseUrl), "application/json; charset=utf-8", b)
}
