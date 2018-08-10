package beep

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func getDefaultSoundTrack() *os.File {
	f, err := os.Open("sound/beep_1.mp3")
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func Beep(id ...int) {
	// Get the beep sound
	var f *os.File
	var err error
	if len(id) > 0 {
		inputFile := "sound/beep_" + strconv.Itoa(id[0]) + ".mp3"
		f, err = os.Open(inputFile)
		if err != nil {
			f = getDefaultSoundTrack()
		}
	} else {
		f = getDefaultSoundTrack()
	}

	// Decode the .mp3 File
	s, format, _ := mp3.Decode(f)

	// Init the Speaker with the SampleRate of the format and a buffer size of 1/10s
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	// Channel, which will signal the end of the playback.
	playing := make(chan struct{})

	// Now we Play our Streamer on the Speaker
	speaker.Play(beep.Seq(s, beep.Callback(func() {
		// Callback after the stream Ends
		close(playing)
	})))
	<-playing
}
