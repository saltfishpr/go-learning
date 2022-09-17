// @file: adapter.go
// @date: 2021/10/28

package main

// MediaAdapter 将 AdvancedMediaPlayer 转化为 MediaPlayer
type MediaAdapter struct {
	AdvancedMediaPlayer
}

func NewMediaAdapter(audioType string) *MediaAdapter {
	switch audioType {
	case "vlc":
		return &MediaAdapter{AdvancedMediaPlayer: &VlcPlayer{}}
	case "mp4":
		return &MediaAdapter{AdvancedMediaPlayer: &Mp4Player{}}
	default:
		return nil
	}
}

func (m MediaAdapter) Play(audioType string, filename string) {
	switch audioType {
	case "vlc":
		m.PlayVlc(filename)
	case "mp4":
		m.PlayMp4(filename)
	default:
	}
}
