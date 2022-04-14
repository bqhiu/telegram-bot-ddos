package main

import (
	"bufio"
	"crypto/tls"
	"net/http"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func AddCommand(cmdName string, handleFunc func(*tgbotapi.BotAPI, tgbotapi.Update)) {
	actions[cmdName] = handleFunc
}

func ParseCommand(command string) (string, string) {
	cmdSplit := strings.Split(command[1:], " ")
	if len(cmdSplit) > 1 {
		cmdAction := cmdSplit[0]
		args := command[len(cmdAction)+2:]
		return cmdAction, args
	}
	return command[1:], ""
}

func ReadLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func HttpGet(url string) (*http.Response, error) {

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{
		Transport: transport,
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
