package main

import "math"

func FtoB(f float64) byte {
	switch {
	case f > 1:
		return 255
	case f < 0:
		return 0
	default:
		return byte(f * 255)
	}
}

//integer angle degree [0,360)
func getDeg(x,y int) (deg int){
	const Pi  = math.Pi
	if x==0 {
		deg = 90
		if y<0 {
			deg*=-1
		}
	}else{
		ang :=math.Atan(float64(y)/float64(x)) //TODO:получение по таблице целочисленных тангенсов
		ang = ang/Pi*2*90
		deg = int(ang)
	}
	if x<0 {
		deg+=180
	}
	if deg<0{
		deg+=360
	}
	return deg
}

func Whirl(){}


//takes x,y in GalaxyUnits, return Density of star population
//0 -can't, 255 - max
func Dens(x, y int) byte {
	//radial blur
	//	f := math.Sqrt(float64(x*x+y*y)) / GalaxyRadius
	//	f=1-f
	//angle beams
	deg:=getDeg(x,y)

	r1:=Whirl

	return FtoB(float64(deg)/360)
}
