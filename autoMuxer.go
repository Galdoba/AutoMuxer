package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Galdoba/utils"
)

const (
	inRoot       = "f:\\Work\\petr_proj\\___IN\\"
	outRoot      = "e:\\_OUT\\"
	senderRoot   = "d:\\SENDER\\"
	inPrefix     = "IN_"
	outPrefix    = "OUT_"
	senderPrefix = "SENDER_"
	taskFile     = "_TaskFile.txt"
)

var startTime time.Time
var inFolder string
var outFolder string
var senderFolder string

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkFolder(adress string) bool {
	//permitionMode uint32 - 0755 дефолтное значение
	//TODO: разобраться с системными доступами к папкам
	err := os.Mkdir(adress, 0755)
	if err != nil {
		if err.Error() == "mkdir "+adress+": Cannot create a file when that file already exists." {
			fmt.Println(adress, "-- folder VALID")
			return true
		}
		panic(err.Error())
	}
	fmt.Println("Folder created: '" + adress + "'")
	return true
}

func todaysFolderName() string {
	yyyy := strconv.Itoa(startTime.Year())
	mm := fmtCurMonth()
	dd := strconv.Itoa(startTime.Day())
	if len(dd) == 1 {
		dd = "0" + dd
	}
	dateKey := yyyy + "-" + mm + "-" + dd + "--test"
	return dateKey
}

func fmtCurMonth() string {
	switch startTime.Month() {
	default:
		panic("Date 1 - unknown Month: " + startTime.Month().String())
	case time.January:
		return "01"
	case time.February:
		return "02"
	case time.March:
		return "03"
	case time.April:
		return "04"
	case time.May:
		return "05"
	case time.June:
		return "06"
	case time.July:
		return "07"
	case time.August:
		return "08"
	case time.September:
		return "09"
	case time.October:
		return "10"
	case time.November:
		return "11"
	case time.December:
		return "12"
	}
}

func preCheck() {
	startTime = time.Now()
	//dateKey := todaysFolderName()
	inFolder = inRoot + "IN_" + todaysFolderName()
	outFolder = outRoot + "OUT_" + todaysFolderName()
	senderFolder = senderRoot + "SENDER_" + todaysFolderName()
	checkFolder(inFolder)
	checkFolder(outFolder)
	checkFolder(senderFolder)
	createTaskFile()

	// if !fileAvailableM(inFolder, taskFile) {
	// 	fmt.Println("Error: No", taskFile, "in location\n", inFolder)
	// 	os.Exit(2)
	// }
	fmt.Println(taskFile, "-- VALID")
}

func createTaskFile() {
	if !fileExists(inFolder, taskFile) {
		f, err := os.Create(inFolder + "\\" + taskFile)
		check(err)
		f.Close()
		fmt.Println(taskFile, "-- Created")
	}
}

type Task struct {
	inVideo     string
	inDurat     string
	offset      string
	prVideo     string
	prAudio     []string
	prSubs      string
	prTags      []string
	outBaseName string
	outDurat    string
	status      string
}

func newTask(dataLine string) *Task {
	task := &Task{}

	task.inVideo = inFolder + "\\" + inVideo
	task.outBaseName = get
	return task
}

func formTask(line string) []string {
	return []string{}
}

func readTaskFile() bool {
	done := false
	if !fileAvailableM(inFolder, taskFile) {

		return done
	}
	path := inFolder + "\\" + taskFile
	fileLines := utils.LinesFromTXT(path)
	for i, v := range fileLines {
		fmt.Println("Line", i, " -- ", v)
	}
	done = true
	return done
}

func main() {
	preCheck()
	fmt.Println(readTaskFile())
	return
	// err = os.Mkdir("subdir", 0755)
	// check(err)

	// err = os.Mkdir("f:\\Work\\petr_proj\\___IN\\Test", 0755)

	// defer os.RemoveAll("subdir")

	// createEmptyFile := func(name string) {
	// 	d := []byte("")
	// 	check(ioutil.WriteFile(name, d, 0644))
	// }

	// createEmptyFile("subdir/file1")

	// err = os.MkdirAll("subdir/parent/child", 0755)
	// check(err)

	// createEmptyFile("subdir/parent/file2")
	// createEmptyFile("subdir/parent/file3")
	// createEmptyFile("subdir/parent/child/file4")

	// c, err := ioutil.ReadDir("subdir/parent")
	// check(err)

	// fmt.Println("Listing subdir/parent")
	// for _, entry := range c {
	// 	fmt.Println(" ", entry.Name(), entry.IsDir())
	// }

	// err = os.Chdir("subdir/parent/child")
	// check(err)

	// c, err = ioutil.ReadDir(".")
	// check(err)

	// fmt.Println("Listing subdir/parent/child")
	// for _, entry := range c {
	// 	fmt.Println(" ", entry.Name(), entry.IsDir())
	// }

	// err = os.Chdir("../../..")
	// check(err)

	// fmt.Println("Visiting subdir")
	// err = filepath.Walk("subdir", visit)
	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(dir)
	// _, filename, _, ok := runtime.Caller(0)
	// if !ok {
	// 	panic("No caller information")
	// }
	// fmt.Printf("Filename : %q, Dir : %q\n", filename, path.Dir(filename))
}

func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}

func fileAvailableM(folder string, file string) bool {
	err := os.Rename(folder+"\\"+file, folder+"\\"+"RENAMED_"+file)
	if err != nil {
		return false
	}
	os.Rename(folder+"\\"+"RENAMED_"+file, folder+"\\"+file)
	return true
}

// exists returns whether the given file or directory exists
func fileExists(folder, file string) bool {
	_, err := os.Stat(folder + "\\" + file)
	if err == nil {
		return true
	}
	return false
}
