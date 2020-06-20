package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"os"
)

func main() {

	var x, y, w, h int
	var r *rover

	r = &rover{
		loc:             &point{},
		min:             &point{},
		max:             &point{},
		dir:             &direction{},
		locationHistory: make([]point, 400000),
		beaconLocations: make([]point, 1000),
	}

	for _, cmd := range roverData() {
		r.execute(string(cmd))
	}

	fmt.Println(len(r.locationHistory), len(r.beaconLocations), len(roverData()), r.min, r.max)

	w = r.max.x - r.min.x
	h = r.max.y - r.min.y

	img := image.NewRGBA(image.Rect(0, 0, w, h))

	for _, loc := range r.locationHistory {
		x = loc.x + r.min.x
		y = loc.y + r.min.y
		img.Set(x, y, color.Black)
	}
	for _, loc := range r.beaconLocations {
		x = loc.x + r.min.x
		y = loc.y + r.min.y
		img.Set(x, y, color.White)
	}

	f, _ := os.Create("image.png")
	_ = png.Encode(f, img)

	// the generated image prints "gcpsrzf5evr1"

}

func roverData() []byte {
	data, err := ioutil.ReadFile("rover-retrieved-program.txt")
	if err != nil {
		panic(err)
	}

	return data
}