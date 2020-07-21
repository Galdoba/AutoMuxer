package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := readArgs()
	if len(args) < 2 {
		fmt.Println("No args... Exit")
		return
	}
	cmd := newComand(f3, "Name", args...)
	fmt.Println(cmd)
	cmd.run()

}

type argSet struct {
	of map[string][]string
}

type ArgFeeder interface {
	Feed(funcName string) []string
}

func readArgs() []string {
	args := os.Args
	var nArgs []string
	for i := range args {
		if i == 0 {
			continue
		}
		nArgs = append(nArgs, args[i])
	}
	return nArgs
}

type comand struct {
	comandName         string
	comandArguments    []string
	bindedFunctionName func(interface{}) (interface{}, error)
}

func newComand(f func(interface{}) (interface{}, error), name string, args ...string) comand {
	cmd := comand{}
	cmd.comandName = name
	cmd.comandArguments = args
	cmd.bindedFunctionName = f
	return cmd
}

func (cmd *comand) run() {
	fn := cmd.bindedFunctionName
	fn(cmd.comandArguments)
}

func plus(argsFeeder interface{}) (interface{}, error) {
	var input []int
	switch argsFeeder.(type) {
	default:
		return nil, errors.New(fmt.Sprintf("Wrong Argument Type: func f3f, args: %d", argsFeeder))
	case []string:
		for i := range argsFeeder.([]string) {
			d, err := strconv.Atoi(argsFeeder.([]string)[i])
			if err != nil {
				panic(errors.New(fmt.Sprintf("Wrong Argument Type: func f3f, args: %d", argsFeeder)))
			}
			input = append(input, d)
		}
	}
	fmt.Println("start f3f")
	var res []int
	for i := range input {
		fmt.Println("f3f add arg", i)
		res = append(res, input[i]+1)
	}
	fmt.Println("f3f executed")
	fmt.Println(res, nil)
	fmt.Println(" ")
	return res, nil
}

func toInts(sl []string) []int {
	var res []int
	for i := range sl {
		n, err := strconv.Atoi(sl[i])
		if err != nil {
			n = 0
		}
		res = append(res, n)
	}
	return res
}

func sum(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

/*
возможные типы команд:
- без аргументов
- с одним аргументом
- с множеством аргументов (!)















Команды:
-ss	---	Стартовая точка
-t	---	Желаемая длинна
-i 	---	input file
-setInFolder	---	Установить папку с исходниками (хранится в конфиге)
-setProcessFolder	---	Установить промежуточную папку для ускорения мукса (хранится в конфиге)  - а надо ли?
-setDestinationFolder	---	Установить папку куда мы будем сливать конечный результат
-resetAll	---	Сбросить весь конфиг на исходные позиции
-setOutput	--- (set Audio Map)	задать карту звука для целевого файла   -taskAudio?
-setbase	---	задать базовое имя результата

Example:
amux -i file.mp4 -ss 00:00:02:00 -t 00:20:00:00 -setoutput hd_ar6e2

-i Perri_meyson_s01e05_SER_09493_AUDIOENG20.m4a
Perri_meyson_s01e05_SER_09493_AUDIORUS51.m4a
Perri_meyson_s01e05_SER_09493_HD.mp4



*/
