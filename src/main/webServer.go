package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"

	//"strconv"
	"sync"
	"time"
)

var mu sync.Mutex
var count int

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

func handler(w http.ResponseWriter, r *http.Request) {

	mu.Lock()
	count++
	fmt.Printf("count++ time is %s\n", time.Now())
	fmt.Fprintf(w, "%s  %s  %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
	mu.Unlock()
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Printf("counter time is %s\n", time.Now())

	fmt.Fprintf(w, "count=%d\n", count)
	mu.Unlock()
}

func lissajous(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	lconf := lconfig{
		cycles: 5,
		res:    0.001,
		freq:   rand.Float64() * 3.0,
		size:   100,
		frames: 64,
		delay:  8,
	}
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		if k == "cycles" {
			if cycles, err := strconv.Atoi(v[0]); err != nil {
				fmt.Printf("Error")
			} else {
				lconf.cycles = float64(cycles)
				fmt.Printf("cycles=%d\n", cycles)
			}
		}
	}
	lissajous1(w, lconf)
}

func lissajous1(out io.Writer, set lconfig) {
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

func web_main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/lissajous", lissajous)

	log.Fatal(http.ListenAndServe("localhost:8081", nil))

}
