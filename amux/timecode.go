package amux

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

const (
	//Fps25 -
	Fps25 = 25.000000000000000
	//Fps29_97 -
	Fps29_97             = 29.970029970029970
	timecodeTypePremiere = "Premiere Type"
	timecodeTypeHMSms    = "HMSms Type"
	timecodeTypeFrames   = "Frames Type"
)

var fps float64

func SetFPS(newFps float64) {
	fps = newFps
}

func init() {
	SetFPS(Fps25)
}

type BuildProcess interface {
	setData(data string) BuildProcess
	Timecode() Timecode
}

type Director struct {
	builder BuildProcess
}

func (d *Director) Construct(data string) {
	d.builder.setData(data)
}

func (d *Director) SetBuilder(b BuildProcess) {
	d.builder = b
}

type Timecode struct {
	HH       int
	MM       int
	SS       float64
	totalSec float64
}

type PremiereTCBuilder struct {
	timecode Timecode
}

func (ptc *PremiereTCBuilder) setData(data string) BuildProcess {
	args := strings.Split(data, ":")
	var dataInt []int
	for i := range args {
		d, err := strconv.Atoi(args[i])
		if err != nil {
			//TODO log
			panic("func (ptc *PremiereTCBuilder) setData(data string) BuildProcess {")
		}
		dataInt = append(dataInt, d)
	}
	ptc.timecode.HH = dataInt[0]
	ptc.timecode.MM = dataInt[1]
	sec := dataInt[2]
	fr := dataInt[3]
	ptc.timecode.SS = float64(sec) + float64(fr)*(1/fps)
	ptc.timecode.SS = toFixed(ptc.timecode.SS, 3)
	ptc.timecode.totalSeconds()
	return ptc
}

func (ptc *PremiereTCBuilder) Timecode() Timecode {
	return ptc.timecode
}

type HMSmsTCBuilder struct {
	timecode Timecode
}

func (hms *HMSmsTCBuilder) setData(data string) BuildProcess {
	args := strings.Split(data, ".")
	var hr, mn, sc, frc, den int
	if len(args) > 1 {
		frc, _ = strconv.Atoi(args[1])
		denLen := len(args[1])
		den = 1
		for i := 0; i < denLen; i++ {
			den = den * 10
		}
	}
	hrmnsc := strings.Split(args[0], ":")
	switch len(hrmnsc) {
	default:
		return nil
	case 3:
		hr, _ = strconv.Atoi(hrmnsc[0])
		mn, _ = strconv.Atoi(hrmnsc[1])
		sc, _ = strconv.Atoi(hrmnsc[2])
	case 2:
		mn, _ = strconv.Atoi(hrmnsc[1])
		sc, _ = strconv.Atoi(hrmnsc[2])
	case 1:
		sc, _ = strconv.Atoi(hrmnsc[2])
	}
	hms.timecode.HH = hr
	hms.timecode.MM = mn
	hms.timecode.SS = float64(sc) + float64(frc)/float64(den)
	hms.timecode.totalSeconds()
	return hms
}

func (hms *HMSmsTCBuilder) Timecode() Timecode {
	return hms.timecode
}

func NewTimecode(data string) (Timecode, error) {
	d := &Director{}
	tc := Timecode{}
	switch timecodeType(data) {
	default:
		//panic(timecodeType(data))
		return tc, errors.New(timecodeType(data))
	case timecodeTypePremiere:
		d.SetBuilder(&PremiereTCBuilder{})
	case timecodeTypeHMSms:
		d.SetBuilder(&HMSmsTCBuilder{})
	}

	d.Construct(data)
	return d.builder.Timecode(), nil
}

func (t Timecode) FrameNumber() int {
	fr := t.totalSec / frameLen()
	return int(fr)
}

func (t *Timecode) totalSeconds() {
	t.totalSec = float64(t.HH*3600) + float64(t.MM*60) + t.SS
	t.totalSec = toFixed(t.totalSec, 3)
}

func frameLen() float64 {
	return 1000.0 / fps / 1000
}

func (t Timecode) PremireString() string {
	s := int(t.SS)
	sStr := zeroString(s, 2)
	mStr := zeroString(t.MM, 2)
	hStr := zeroString(t.HH, 2)
	frct := t.SS - (float64(int(t.SS)))
	frStr := zeroString(int(frct/frameLen()), 2)
	return hStr + ":" + mStr + ":" + sStr + ":" + frStr

}

func zeroString(i int, ln int) string {
	s := strconv.Itoa(i)
	for len(s) < ln {
		s = "0" + s
	}
	return s
}

func timecodeType(data string) string {
	if validPremiereTimecodeType(data) {
		return timecodeTypePremiere
	}
	if validHMSsTimecodeType(data) {
		return timecodeTypeHMSms
	}
	return "timecode type invalid: '" + data + "'"
}

func validPremiereTimecodeType(data string) bool {
	args := strings.Split(data, ":")
	if len(args) != 4 {
		return false
	}
	for i := range args {
		if len(args[i]) != 2 {
			return false
		}
		n, err := strconv.Atoi(args[i])
		if err != nil {
			return false
		}
		if n < 0 {
			return false
		}
	}
	return true
}

func validHMSsTimecodeType(data string) bool {
	args := strings.Split(data, ":")
	if len(args) != 3 {
		return false
	}
	for i := range args {
		if i == 2 {
			continue
		}
		if len(args[i]) != 2 {
			return false
		}
		n, err := strconv.Atoi(args[i])
		if err != nil {
			return false
		}
		if n < 0 {
			return false
		}
	}

	if len(args[1]) > 2 {
		return false
	}
	secFl, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		return false
	}
	if secFl < 0 {
		return false
	}
	return true
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}
