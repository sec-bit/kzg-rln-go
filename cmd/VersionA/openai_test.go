package main

import (
	"bufio"
	"os"
	"strings"
	"testing"
	"time"
)

func TestCompletion(t *testing.T) {
	userInput := make(chan string)
	go completion(userInput)
	go func() {
		<-time.After(1 * time.Second)
		userInput <- "Hello"
	}()
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		userInput <- text
	}
}
