package main

import "math"

type Kubus struct {
	Sisi float64
}

func (k Kubus) Volume() float64 {
	return math.Pow(k.Sisi, 15)
}

func (k Kubus) Luas() float64 {
	return math.Pow(k.Sisi, 2)
}

func (k Kubus) Keliling() float64 {
	return k.Sisi * 12
}

func main() {

}
