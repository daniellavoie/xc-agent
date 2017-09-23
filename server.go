package main

import (
	"net"
	"net/http"
	"xcomponent.com/xcagent/command"
	"encoding/json"
)

func StartServer() (net.Listener, error) {
	http.HandleFunc("/info", infoEndpoint)
	http.HandleFunc("/command/run", runCommandEndpoint)
	http.HandleFunc("/command/logs", logsCommandEndpoint)

	return net.Listen("tcp", "0.0.0.0:0")
}

func infoEndpoint(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("hello world"))
}

func logsCommandEndpoint(response http.ResponseWriter, request *http.Request) {
	var logsCommand command.LogsCommand
	if request.Body == nil {
		http.Error(response, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(request.Body).Decode(&logsCommand)
	if err != nil {
		http.Error(response, err.Error(), 400)
		return
	}

	output, err := command.ExecuteLogsCommand(logsCommand)

	if err != nil {
		http.Error(response, err.Error(), 500)
	} else {
		response.Header().Set("Content-Type", "application/json")
		json.NewEncoder(response).Encode(output)
	}

}

func runCommandEndpoint(response http.ResponseWriter, request *http.Request) {
	var runCommand command.RunCommand
	if request.Body == nil {
		http.Error(response, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(request.Body).Decode(&runCommand)
	if err != nil {
		http.Error(response, err.Error(), 400)
		return
	}

	output, err := command.ExecuteRunCommand(runCommand)

	if err != nil {
		http.Error(response, err.Error(), 500)
	} else {
		response.Header().Set("Content-Type", "application/json")
		json.NewEncoder(response).Encode(output)
	}

}
