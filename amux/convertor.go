package amux

import (
	"math"
	"strconv"
	"strings"
)

const (
	fps25 = 25
)

type VidDuration interface {
	Frames() int
	PremDuration() string
	FFmpegDuration() string
}

func validPrem(data string, frames int) bool {
	args := strings.Split(data, ":")
	if len(args) != 4 {
		return false
	}
	for i := range args {
		num, err := strconv.Atoi(args[i])
		if err != nil {
			return false
		}
		switch i {
		case 1, 2:
			if num > 59 || num < 0 {
				return false
			}
		case 3:
			if num > frames || num < 0 {
				return false
			}
		}
	}
	return true
}

func validFFMPEG(data string, frames ...int) bool {
	args := strings.Split(data, ":")
	if len(args) != 3 {
		return false
	}
	for i := range args {
		if i == 2 {
			continue
		}
		num, err := strconv.Atoi(args[i])
		if err != nil {
			return false
		}
		switch i {
		case 1:
			if num > 59 || num < 0 {
				return false
			}
		}
	}
	secs, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		return false
	}
	part := secs
	for part > 1 {
		part = part - 1
	}
	part = toFixed(part, 3)
	part = part / (1 / float64(frames))
	frame := int(part)
	if frame > frames {
		return false
	}
	return true
}

func FfToPrem(duration string) string {
	premiereDur := ""
	timeArgs := strings.Split(duration, ":")
	hours, errHour := strconv.Atoi(timeArgs[0])
	if errHour != nil {
		panic(errHour)
	}
	minutes, errMin := strconv.Atoi(timeArgs[1])
	if errMin != nil {
		panic(errMin)
	}
	secs, err := strconv.ParseFloat(timeArgs[2], 64)
	if err != nil {
		panic(err)
	}
	secsInt := int(secs)
	strHour := strconv.Itoa(hours)
	if hours < 10 {
		strHour = "0" + strHour
	}
	strMin := strconv.Itoa(minutes)
	if minutes < 10 {
		strMin = "0" + strMin
	}
	strSec := strconv.Itoa(secsInt)
	if secsInt < 10 {
		strSec = "0" + strSec
	}
	part := secs
	for part > 1 {
		part = part - 1
	}
	part = toFixed(part, 3)
	part = part / 0.04
	frame := int(part)
	strFrame := strconv.Itoa(frame)
	if frame < 10 {
		strFrame = "0" + strFrame
	}
	if frame == 25 {
		strFrame = "00"
	}
	premiereDur = strHour + ":" + strMin + ":" + strSec + ":" + strFrame
	return premiereDur
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
