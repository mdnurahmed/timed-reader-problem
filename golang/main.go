package main

import (
	"fmt"
	"time"
)

func inputThread(input chan string) {
	var s string
	fmt.Scanf("%s\n", &s)
	input <- s
}
func main() {
	ticker := time.NewTicker(time.Second)
	input := make(chan string)
	counter := 0
	go inputThread(input)

waitforinput:
	for {
		select {
		case text := <-input:
			fmt.Printf("user typed - %s\n", text)
			fmt.Println("----------------------")
			counter = 0
			go inputThread(input)
		case <-ticker.C:
			counter++
			fmt.Printf("\nTick at %d seconds \n", counter)
			if counter == 3 {
				break waitforinput
			}
		}
	}
	fmt.Println("End of Code")

}
