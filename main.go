package main

import (
	"fmt"

	"github.com/Galdoba/AutoMuxer/amux"
)

func main() {

	// report, err := ffinfo.Probe("f:\\Work\\petr_proj\\___IN\\IN_2020-07-13--test\\Tracks\\Tracks_HD.mp4")
	// if err != nil {
	// 	panic(0)
	// }
	// fmt.Println(report)
	// fmt.Println("-----------------")
	// fl, err := report.StreamDuration(0)

	// if err != nil {
	// 	panic(1)
	// }
	// fmt.Println(fl)
	//	amux.SetFPS(amux.Fps29_97)

	tcStr := "02:19:21.950"
	tc, err := amux.NewTimecode(tcStr)
	fmt.Println("Error:", err)

	fmt.Println("tc:", tc, "| tcStr:", tcStr)

	fmt.Println(tc.FrameNumber())

	fmt.Println(tc.PremireString())

	//fmt.Println(amux.DisassembleData("20:20:20:05"))
	//fmt.Println(amux.DisassembleData("10:20:30.42"))
}
