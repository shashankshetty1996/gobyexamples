/*
	Channels are the pipes that connect concurrent goroutines, You can send values into channels from one goroutine and receive those values into another goroutine.

	Note:
	1. When we run the program the "ping" messages is successfully passed from one goroutine to another via our channel.

	2. By default sends and recieves blocks until both the sender and receiver are ready. This property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization.
*/

package main

import "fmt"

func main() {
	// Create a new channel with make(cha val-type). Channels are typed by the values they convey.
	message := make(chan string)

	// Send a value into a channel using the channel <- syntax. Here we send "ping" to the messages channel we made above, from a new goroutine.
	go func() { message <- "ping" }()

	// The <- channel syntax recieves a value from the channel. Here we'll receieve the "ping" message we sent above and print it out.
	msg := <-message
	fmt.Println(msg)
}
