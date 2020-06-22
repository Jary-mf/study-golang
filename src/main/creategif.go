package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

type lconfig struct {
	cycles float64
	res    float64
	freq   float64
	size   int
	frames int
	delay  int
}

func gif_main() {
	rand.Seed(time.Now().UnixNano())
	lconf := lconfig{
		cycles: 5,
		res:    0.001,
		freq:   rand.Float64() * 3.0,
		size:   100,
		frames: 64,
		delay:  8,
	}
	lissajous(os.Stdout, lconf)
}

func lissajous(out io.Writer,set lconfig) {
	anim := gif.GIF{LoopCount: set.frames}
	phase := 0.0 // phase difference
	for i := 0; i < set.frames; i++ {
		rect := image.Rect(0, 0, 2*set.size+1, 2*set.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < set.cycles*2*math.Pi; t += set.res {
			x := math.Sin(t)
			y := math.Sin(t*set.freq + phase)
			img.SetColorIndex(set.size+int(x*float64(set.size)+0.5), set.size+int(y*float64(set.size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, set.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}