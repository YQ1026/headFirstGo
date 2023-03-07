package gadget

import (
	"fmt"
)

// 播放器对象以及方法
type TapePlayer struct {
	batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("playing", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("stopped!")
}

// 录音机对象以及方法
type TapeRecorder struct {
	Microphone int
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("录制歌曲", song)
}

func (t TapeRecorder) Stop() {
	fmt.Println("stopped!")
}

func (t TapeRecorder) Record() {
	fmt.Println("我开始录音了哈")
}
