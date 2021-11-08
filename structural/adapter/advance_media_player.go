// @file: advance_media_player.go
// @date: 2021/10/28

package main

import "fmt"

type AdvancedMediaPlayer interface {
	PlayVlc(filename string)
	PlayMp4(filename string)
}

type VlcPlayer struct{}

func (VlcPlayer) PlayVlc(filename string) {
	fmt.Println("Playing vlc file. Filename: ", filename)
}

func (VlcPlayer) PlayMp4(filename string) {
	// do noting
}

type Mp4Player struct{}

func (Mp4Player) PlayVlc(filename string) {
	// do noting
}

func (Mp4Player) PlayMp4(filename string) {
	fmt.Println("Playing mp4 file. Filename: ", filename)
}
