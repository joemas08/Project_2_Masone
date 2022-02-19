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
			log.Fatalln(err)
		}
		if strings.ToLower(userAnswer) == "yes" {
			go fortune(ch)
			ch <- ""

		} else if strings.ToLower(userAnswer) == "no" {
			os.Exit(-1)
		} else {
			fmt.Println("Answer must be YES or NO! ")
			continue
		}
	}

}
func fortune(c chan string) {
	fileasBytes, err := ioutil.ReadFile("Fortunes.txt")
	if err != nil {
		log.Fatalln("Error reading this file: ", err)

	}
	fortunes := strings.Split(string(fileasBytes), "%%")

	for {
		msg := <-c
		randomIndex := rand.Intn(len(fortunes))
		fortChosen := fortunes[randomIndex]
		fmt.Println(fortChosen)
		fmt.Println(msg)
	}
}
