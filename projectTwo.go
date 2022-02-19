package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
)

func main() {
	var userAnswer string
	ch := make(chan string)

	for {
		fmt.Println("Would you like another fortune?: YES or NO")
		_, err := fmt.Scan(&userAnswer)
		if err != nil {
			log.Fatalln("Error reading this file: ", err)
		}
		if strings.ToLower(userAnswer) == "yes" {
			go fortune(ch)
		} else if strings.ToLower(userAnswer) == "no" {
			os.Exit(-1)
		} else {
			fmt.Println("Answer must be YES or NO! ")
			continue
		}
	}

}
func fortune(c chan string) {
	fileasBytes, err := ioutil.ReadFile("Fortunes.txt") //converting file to bytes
	if err != nil {
		log.Fatalln("Error reading this file: ", err)

	}
	fortunes := strings.Split(string(fileasBytes), "%%")
	randomIndex := rand.Intn(len(fortunes))
	pick := fortunes[randomIndex]
	c <- pick
}
