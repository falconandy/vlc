package vlc

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"strings"
	"testing"
	"time"
)

const (
	sampleVideoPath = "sample.mkv"
	debugMode       = true
)

func TestPlayer_Start(t *testing.T) {
	p := NewPlayer(nil).SetDebugMode(debugMode)
	err := p.Start()
	assert.Nil(t, err)

	time.Sleep(time.Second * 2)

	err = p.Shutdown()
	assert.Nil(t, err)
}

func TestPlayer_Play(t *testing.T) {
	p := NewPlayer(nil).SetDebugMode(debugMode)
	err := p.Start()
	assert.Nil(t, err)

	err = p.Play(sampleVideoPath)
	assert.Nil(t, err)

	isPlaying, err := p.IsPlaying()
	assert.Nil(t, err)
	assert.True(t, isPlaying)

	time.Sleep(time.Second * 5)

	err = p.Pause()
	assert.Nil(t, err)

	isPlaying, err = p.IsPlaying()
	assert.Nil(t, err)
	assert.True(t, isPlaying)

	time.Sleep(time.Second * 2)

	err = p.Pause()
	assert.Nil(t, err)

	isPlaying, err = p.IsPlaying()
	assert.Nil(t, err)
	assert.True(t, isPlaying)

	time.Sleep(time.Second * 3)

	err = p.Stop()
	assert.Nil(t, err)

	isPlaying, err = p.IsPlaying()
	assert.Nil(t, err)
	assert.False(t, isPlaying)

	err = p.Shutdown()
	assert.Nil(t, err)
}

func TestPlayer_Seek(t *testing.T) {
	p := NewPlayer(nil).SetDebugMode(debugMode)
	err := p.Start()
	assert.Nil(t, err)

	err = p.Play(sampleVideoPath)
	assert.Nil(t, err)

	length, err := p.Length()
	assert.Nil(t, err)
	assert.True(t, length > 0)

	time.Sleep(time.Second * 2)

	for i := 0; i < 5; i++ {
		seekTo := Duration(rand.Intn(int(length) - 10))
		err := p.Seek(seekTo)
		assert.Nil(t, err)

		time.Sleep(time.Second * 2)

		position, err := p.Position()
		assert.Nil(t, err)
		assert.True(t, seekTo < position)
		assert.True(t, position < length)
	}

	err = p.Stop()
	assert.Nil(t, err)

	err = p.Shutdown()
	assert.Nil(t, err)
}

func TestPlayer_Speed(t *testing.T) {
	p := NewPlayer(nil).SetDebugMode(debugMode)
	err := p.Start()
	assert.Nil(t, err)

	err = p.Play(sampleVideoPath)
	assert.Nil(t, err)

	length, err := p.Length()
	assert.Nil(t, err)
	assert.True(t, length > 0)

	seekTo := Duration(int(length) / 2)
	err = p.Seek(seekTo)
	assert.Nil(t, err)

	err = p.SpeedFaster()
	assert.Nil(t, err)
	err = p.SpeedFaster()
	assert.Nil(t, err)
	err = p.SpeedFaster()
	assert.Nil(t, err)

	time.Sleep(time.Second * 5)

	err = p.SpeedNormal()
	assert.Nil(t, err)

	time.Sleep(time.Second * 3)

	err = p.SpeedSlower()
	assert.Nil(t, err)
	err = p.SpeedSlower()
	assert.Nil(t, err)

	time.Sleep(time.Second * 5)

	err = p.Stop()
	assert.Nil(t, err)

	err = p.Shutdown()
	assert.Nil(t, err)
}

func TestPlayer_Audio(t *testing.T) {
	p := NewPlayer(nil).SetDebugMode(debugMode)
	err := p.Start()
	assert.Nil(t, err)

	err = p.Play(sampleVideoPath)
	assert.Nil(t, err)

	length, err := p.Length()
	assert.Nil(t, err)
	assert.True(t, length > 0)

	time.Sleep(time.Second * 2)

	tracks, track, err := p.AudioTracks()
	assert.Nil(t, err)
	assert.Equal(t, 1, track)

	trackIndexes := make([]int, 0, len(tracks))
	for trackIndex := range tracks {
		trackIndexes = append(trackIndexes, trackIndex)
	}
	sort.Ints(trackIndexes)

	for i := 0; i < 3; i++ {
		seekTo := Duration(rand.Intn(int(length) - 10))
		err := p.Seek(seekTo)
		assert.Nil(t, err)

		track := -1
		if len(trackIndexes) > 0 {
			track = trackIndexes[i%len(trackIndexes)]
		}
		err = p.SetAudioTrack(track)
		assert.Nil(t, err)

		_, actualTrack, err := p.AudioTracks()
		assert.Nil(t, err)
		assert.Equal(t, track, actualTrack)

		time.Sleep(time.Second * 7)
	}

	err = p.Stop()
	assert.Nil(t, err)

	err = p.Shutdown()
	assert.Nil(t, err)
}

func TestPlayer_Subtitles(t *testing.T) {
	p := NewPlayer(nil).SetDebugMode(debugMode)
	err := p.Start()
	assert.Nil(t, err)

	err = p.Play(sampleVideoPath)
	assert.Nil(t, err)

	length, err := p.Length()
	assert.Nil(t, err)
	assert.True(t, length > 0)

	time.Sleep(time.Second * 2)

	tracks, track, err := p.SubtitleTrack()
	assert.Nil(t, err)
	assert.Equal(t, -1, track)

	trackIndexes := make([]int, 0, len(tracks))
	for trackIndex, trackTitle := range tracks {
		if !strings.Contains(strings.ToLower(trackTitle), "forced") {
			trackIndexes = append(trackIndexes, trackIndex)
		}
	}
	sort.Ints(trackIndexes)

	for i := 0; i < 3; i++ {
		seekTo := Duration(rand.Intn(int(length) - 10))
		err := p.Seek(seekTo)
		assert.Nil(t, err)

		track := -1
		if len(trackIndexes) > 0 {
			track = trackIndexes[i%len(trackIndexes)]
		}
		err = p.SetSubtitleTrack(track)
		assert.Nil(t, err)

		_, actualTrack, err := p.SubtitleTrack()
		assert.Nil(t, err)
		assert.Equal(t, track, actualTrack)

		time.Sleep(time.Second * 7)
	}

	err = p.Stop()
	assert.Nil(t, err)

	err = p.Shutdown()
	assert.Nil(t, err)
}
