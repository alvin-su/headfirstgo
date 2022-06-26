package gadget

import "fmt"

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing ", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

type TapeRecord struct {
	Microphones int
}

func (t TapeRecord) Play(song string) {
	fmt.Println("Playing ", song)
}

func (t TapeRecord) Record() {
	fmt.Println("Recording")
}

func (t TapeRecord) Stop() {
	fmt.Println("Stopped!")
}
