package main

import (
	"fmt"
	"go-headFirst/ch11/gadget"
)

type Player interface {
	Play(song string)
	Stop()
}

func palyList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func TryOut(player Player) {
	player.Play("Test Trach")
	player.Stop()
	record, ok := player.(gadget.TapeRecord) //类型断言,取回具体类型的值
	if ok {
		record.Record()
	} else {
		fmt.Println("Player was not a TapeRecorder")
	}

}

func main() {
	var player Player = gadget.TapePlayer{}
	player2 := gadget.TapeRecord{}
	songSlice := []string{"歌曲1", "歌曲2", "歌曲3"}
	palyList(player, songSlice)

	palyList(player2, songSlice)

	TryOut(player2)
	TryOut(player)
}
