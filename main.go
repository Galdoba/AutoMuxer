package main

import (
	"fmt"

	"github.com/Galdoba/AutoMuxer/amux"
	"github.com/malashin/ffinfo"
)

func main() {
	fmt.Println(amux.FfToPrem("20:20:20.200"))
	str, err := ffinfo.Probe("f:\\Work\\petr_proj\\___IN\\IN_2020-07-13--test\\Tracks\\Tracks_HD.mp4")
	if err != nil {
		panic(0)
	}
	fmt.Println(str)
	fl, err := str.StreamDuration(0)
	if err != nil {
		panic(1)
	}
	fmt.Println(fl)
}
