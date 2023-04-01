package main

import "fmt"

func main() {
	basicChannel()

	twoMessagesChannel()

	twoChannels()

}

func twoChannels() {
	email1 := make(chan string)
	email2 := make(chan string)

	go say("message 1", email1)
	go say("message 2", email2)

	// Channels do not respond in order
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Send for email 1", m1)
		case m2 := <-email2:
			fmt.Println("Send for email 2", m2)
		}
	}
}

func twoMessagesChannel() {
	// two message channel
	channel := make(chan string, 2)
	channel <- "MESSAGE 1"
	channel <- "MESSAGE 2"

	fmt.Printf("Length: %d\nCapacity: %d\n", len(channel), cap(channel))

	// close channel
	close(channel)

	for message := range channel {
		fmt.Println(message)
	}
}

func basicChannel() {
	channel := make(chan string)
	fmt.Println("hello")

	go say("bye", channel)

	fmt.Println(<-channel)
}

// IO channel c chan string
// Only output channel c <- chan string
// Only input
func say(text string, c chan<- string) {
	c <- text
}
