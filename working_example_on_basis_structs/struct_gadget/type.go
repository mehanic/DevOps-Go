package main

import "fmt"

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

type TapeRecorder struct {
	Microphones int
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapeRecorder) Record() {
	fmt.Println("Recording!")
}

func (t TapeRecorder) Stop() {
	fmt.Println("Stopped!")
}

func main() {
	// Create instances of TapePlayer and TapeRecorder
	player := TapePlayer{Batteries: "AA"}
	recorder := TapeRecorder{Microphones: 2}

	// Use TapePlayer
	fmt.Println("Using TapePlayer:")
	player.Play("Song A")
	player.Stop()

	// Use TapeRecorder
	fmt.Println("\nUsing TapeRecorder:")
	recorder.Play("Song B")
	recorder.Record()
	recorder.Stop()
}

// Using TapePlayer:
// Playing Song A
// Stopped!

// Using TapeRecorder:
// Playing Song B
// Recording!
// Stopped!
