package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	file, _ := os.Open("/opt/terminalSplash/splashText.txt")
	fileScanner := bufio.NewScanner(file)
	numOfLines := 0
	for fileScanner.Scan() {
		numOfLines++
	}

	randomQuoteIndex := rand.Intn(numOfLines)

	counter := 0
	file, _ = os.Open("/opt/terminalSplash/splashText.txt")
	fileScanner = bufio.NewScanner(file)
	for fileScanner.Scan() {
		if counter == randomQuoteIndex {
			fmt.Println(fileScanner.Text())
			break
		} else {
			counter++
		}
	}

}
