package tts

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEdgeTTSService_GenerateAudio(t *testing.T) {
	tts := NewEdgeTTSService()

	const format = "audio-24khz-48kbitrate-mono-mp3"

	audio, err := tts.GenerateAudio("你好，我是无道", "zh-CN-XiaoxiaoNeural", 0, 0, format)
	assert.Nil(t, err)
	assert.NotEmpty(t, audio)

	_ = os.WriteFile("test.mp3", audio, 0644)
	//t.Logf("audio: %v", audio)
}
