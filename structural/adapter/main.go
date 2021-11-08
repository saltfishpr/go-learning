// @file: main.go
// @date: 2021/10/28

package main

func main() {
	var audioPlayer MediaPlayer = new(AudioPlayer)
	audioPlayer.Play("mp3", "beyond the horizon.mp3")
	audioPlayer.Play("mp4", "alone.mp4")
	audioPlayer.Play("vlc", "far far away.vlc")
	audioPlayer.Play("avi", "mind me.avi")
}
