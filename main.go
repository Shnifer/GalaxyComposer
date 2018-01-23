package main

import (
	"image"
	"image/png"
	"image/color"
	"os"
	"log"
	"math/rand"
	"time"
)

const GalaxyRadius = 50000 //in light years
const HalfImageSize = 500 //in pixels, better even



func main(){

	rand.Seed(time.Now().Unix())

	stars := make([]image.Point,1000)

	for i,_:=range stars{
		stars[i].X = rand.Intn(HalfImageSize*2)-HalfImageSize
		stars[i].Y = rand.Intn(HalfImageSize*2)-HalfImageSize
	}

	r:=image.Rect(-HalfImageSize,-HalfImageSize,HalfImageSize,HalfImageSize)
	img:=image.NewRGBA(r)

	for _,star:=range stars {
		x:=star.X
		y:=star.Y
		img.Set(x, y, color.NRGBA{
			R: uint8((x + y) & 255),
			G: uint8((x + y) << 1 & 255),
			B: uint8((x + y) << 2 & 255),
			A: 255,
		})
	}

	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
