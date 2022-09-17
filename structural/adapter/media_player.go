// @file: media_player.go
// @date: 2021/10/28

package main

import "fmt"

type MediaPlayer interface {
	Play(audioType string, filename string)
}

type AudioPlayer struct{}

func (AudioPlayer) Play(audioType string, filename string) {
	switch audioType {
	case "mp3":
		fmt.Println("Playing mp3 file. Filename: ", filename)
	case "vlc", "mp4": // use MediaAdapter to play vlc and mp4 file
		mediaAdapter := NewMediaAdapter(audioType)
		mediaAdapter.Play(audioType, filename)
	default:
		fmt.Println("Invalid media type: ", audioType)
	}
}
