package command

import (
	"os/exec"
	"github.com/satori/go.uuid"
	"os"
	"fmt"
	"bufio"
	"log"
	"bytes"
)

type Output struct {
	Content string `json:"content"`
}

type LogsCommand struct {
	StackName   string `json:"stackName"`
	ServiceName string `json:serviceName`
	Instance    string `json:instance`
}
type RunCommand struct {
	StackName       string `json:"stackName"`
	StackDefinition string `json:"stackDefinition"`
}

func ExecuteLogsCommand(command LogsCommand) (output *Output, err error) {
	cmd := exec.Command("docker", "logs", fmt.Sprintf("%s_%s_%s", command.StackName, command.ServiceName, command.Instance))

	var outbuf, errbuf bytes.Buffer

	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	cmd.Start()
	cmd.Wait()

	stderr := outbuf.String()

	return &Output{Content: stderr}, nil
}

func ExecuteRunCommand(command RunCommand) (output *Output, err error) {
	randomUuid := uuid.NewV4()
	tmpFolder := fmt.Sprintf("%s%s", os.TempDir(), randomUuid)

	os.MkdirAll(tmpFolder, os.ModePerm)

	defer func() {
		log.Print("Deleting temporary folder.\n")
		os.RemoveAll(tmpFolder)
	}()

	filePath := fmt.Sprintf("%s/docker-compose.yml", tmpFolder)
	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(command.StackDefinition)
	if err != err {
		return nil, err
	}
	writer.Flush()

	log.Print("Executing docker-compose.\n")

	cmd := exec.Command("docker-compose", "-p", command.StackName, "-f", filePath, "up", "-d")

	var outbuf, errbuf bytes.Buffer

	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	cmd.Start()
	cmd.Wait()

	stderr := errbuf.String()

	return &Output{Content: stderr}, nil
}
