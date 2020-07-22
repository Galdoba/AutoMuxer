package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := readArgs()

	kaSl := KeyArgs()
	fmt.Println(kaSl)
	fmt.Println("---------")
	test := "123456789"
	fmt.Println(string(test[0]))

	if len(args) < 2 {
		fmt.Println("No args... Exit")
		return
	}
	var cmd comand
	switch args[0] {
	default:
		fmt.Println("Key is not in White list")
		os.Exit(1)
	case "-Sum":
		cmd = newComand(plus, args[0], args[1:]...)
	case "-Minus":
		cmd = newComand(minus, args[0], args[1:]...)
	}
	fmt.Println(cmd)
	res, err := cmd.run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
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

func (cmd *comand) run() (interface{}, error) {
	fn := cmd.bindedFunctionName
	output, err := fn(cmd.comandArguments)
	return output, err
}

func plus(argsFeeder interface{}) (interface{}, error) {
	var input []int
	switch argsFeeder.(type) {
	default:
		return nil, errors.New(fmt.Sprintf("Wrong Argument Type: func f3f, args: %d", argsFeeder))
	case []string:
		input = toInts(argsFeeder.([]string))
	}
	fmt.Println("Args =", input)
	var res int
	for i := range input {
		fmt.Println("add", input[i])
		res = res + input[i]
	}
	fmt.Println("Result:", res)
	return res, nil
}

func minus(argsFeeder interface{}) (interface{}, error) {
	var input []int
	switch argsFeeder.(type) {
	default:
		return nil, errors.New(fmt.Sprintf("Wrong Argument Type: func f3f, args: %d", argsFeeder))
	case []string:
		input = toInts(argsFeeder.([]string))
	}
	fmt.Println("Args =", input)
	var res int
	for i := range input {
		if i == 0 {
			res = input[i]
			continue
		}
		fmt.Println("sub", input[i])
		res = res - input[i]
	}
	fmt.Println("Result:", res)
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





amux -i file.mp4 -ss 00:00:02:00 -t 00:20:00:00 -setoutput hd_ar6e2


*/
