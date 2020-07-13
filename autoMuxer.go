package main

// import (
// 	"errors"
// 	"fmt"
// 	"math"
// 	"os"
// 	"os/exec"
// 	"regexp"
// 	"strconv"
// 	"strings"
// 	"sync"
// 	"time"

// 	"github.com/Galdoba/utils"
// )

// const (
// 	inRoot       = "f:\\Work\\petr_proj\\___IN\\"
// 	outRoot      = "e:\\_OUT\\"
// 	senderRoot   = "d:\\SENDER\\"
// 	inPrefix     = "IN_"
// 	outPrefix    = "OUT_"
// 	senderPrefix = "SENDER_"
// 	taskFile     = "_TaskFile.txt"
// 	taskCon0     = "TaskStatus-0"
// 	taskCon1     = "TaskStatus-1"
// 	taskCon2     = "TaskStatus-2"
// 	taskCon3     = "TaskStatus-3"
// 	taskCon4     = "TaskStatus-4"
// 	taskCon5     = "TaskStatus-5"
// 	taskCon6     = "TaskStatus-6"
// 	taskCon7     = "TaskStatus-7"
// 	taskDone     = "TaskStatus-DONE"
// )

// var startTime time.Time
// var inFolder string
// var outFolder string
// var senderFolder string
// var taskFilePath string
// var activeTask int

// func check(e error) {
// 	if e != nil {
// 		panic(e)
// 	}
// }

// func checkFolder(adress string) bool {
// 	//permitionMode uint32 - 0755 дефолтное значение
// 	//TODO: разобраться с системными доступами к папкам
// 	err := os.Mkdir(adress, 0755)
// 	if err != nil {
// 		if err.Error() == "mkdir "+adress+": Cannot create a file when that file already exists." {
// 			fmt.Println(adress, "-- folder VALID")
// 			return true
// 		}
// 		panic(err.Error())
// 	}
// 	fmt.Println("Folder created: '" + adress + "'")
// 	return true
// }

// func todaysFolderName() string {
// 	yyyy := strconv.Itoa(startTime.Year())
// 	mm := fmtCurMonth()
// 	dd := strconv.Itoa(startTime.Day())
// 	if len(dd) == 1 {
// 		dd = "0" + dd
// 	}
// 	dateKey := yyyy + "-" + mm + "-" + dd + "--test"
// 	return dateKey
// }

// func fmtCurMonth() string {
// 	switch startTime.Month() {
// 	default:
// 		panic("Date 1 - unknown Month: " + startTime.Month().String())
// 	case time.January:
// 		return "01"
// 	case time.February:
// 		return "02"
// 	case time.March:
// 		return "03"
// 	case time.April:
// 		return "04"
// 	case time.May:
// 		return "05"
// 	case time.June:
// 		return "06"
// 	case time.July:
// 		return "07"
// 	case time.August:
// 		return "08"
// 	case time.September:
// 		return "09"
// 	case time.October:
// 		return "10"
// 	case time.November:
// 		return "11"
// 	case time.December:
// 		return "12"
// 	}
// }

// func preCheck() {
// 	startTime = time.Now()
// 	//dateKey := todaysFolderName()
// 	inFolder = inRoot + "IN_" + todaysFolderName()
// 	outFolder = outRoot + "OUT_" + todaysFolderName()
// 	senderFolder = senderRoot + "SENDER_" + todaysFolderName()
// 	checkFolder(inFolder)
// 	checkFolder(outFolder)
// 	checkFolder(senderFolder)
// 	createTaskFile()
// 	taskFilePath = inFolder + "\\" + taskFile
// 	// if !fileAvailableM(inFolder, taskFile) {
// 	// 	fmt.Println("Error: No", taskFile, "in location\n", inFolder)
// 	// 	os.Exit(2)
// 	// }
// 	fmt.Println(taskFile, "-- VALID")
// }

// func createTaskFile() {
// 	if !fileExists(inFolder, taskFile) {
// 		f, err := os.Create(inFolder + "\\" + taskFile)
// 		check(err)
// 		f.Close()

// 		fmt.Println(taskFile, "-- Created")
// 	}

// }

// type Task struct {
// 	dataLine    string
// 	inVideo     string
// 	inDurat     string
// 	offset      string
// 	prVideo     string
// 	prAudio     []string
// 	prSubs      string
// 	outBaseName string
// 	outFullName string
// 	outDurat    string
// 	status      int
// 	inFilePos   int
// 	outputTags  map[string]string
// }

// func mapTags(arg string) map[string]string {
// 	tagMap := make(map[string]string)
// 	argParts := strings.Split(arg, "__")
// 	part1 := strings.Split(argParts[0], "_")
// 	nameLen := len(part1) - 1
// 	name := strings.Join(part1[:nameLen], "_")
// 	tagMap["name"] = name
// 	year := part1[nameLen:][0]
// 	tagMap["year"] = year
// 	part2 := strings.Split(argParts[1], "_")
// 	for i := range part2 {
// 		switch i {
// 		default:
// 			fmt.Println("Err??", part2[i])
// 		case 0:
// 			if validateTag(part2[i], resolutionTagWL()) {
// 				tagMap["resolutionTag"] = part2[i]
// 			}
// 		case 1:
// 			if validateTag(part2[i], audioTagWL()) {
// 				tagMap["audioTag"] = part2[i]
// 			}
// 		case 2:
// 			if validateTag(part2[i], subsTagWL()) {
// 				tagMap["subsTag"] = part2[i]
// 			}
// 		}
// 	}

// 	return tagMap
// }

// func (t *Task) predictAudioFiles() (audio1, audio2 string) {
// 	switch t.outputTags["audioTag"] {
// 	case "ar2":
// 		audio1 = t.outBaseName + "_rus20"
// 	case "ae2":
// 		audio1 = t.outBaseName + "_eng20"
// 	case "ar6":
// 		audio1 = t.outBaseName + "_rus51"
// 	case "ae6":
// 		audio1 = t.outBaseName + "_eng51"
// 	case "ar2e2":
// 		audio1 = t.outBaseName + "_rus20"
// 		audio2 = t.outBaseName + "_eng20"
// 	case "ar2e6":
// 		audio1 = t.outBaseName + "_rus20"
// 		audio2 = t.outBaseName + "_eng51"
// 	case "ar6e2":
// 		audio1 = t.outBaseName + "_rus51"
// 		audio2 = t.outBaseName + "_eng20"
// 	case "ar6e6":
// 		audio1 = t.outBaseName + "_rus51"
// 		audio2 = t.outBaseName + "_eng51"
// 	}
// 	t.prAudio = append(t.prAudio, audio1+".ac3")
// 	if audio2 != "" {
// 		t.prAudio = append(t.prAudio, audio2+".ac3")
// 	}
// 	return audio1, audio2
// }

// func (t *Task) predictSubsFile() (subs string) {
// 	if t.outputTags["subsTag"] == "[NO_SUBS]" {
// 		return ""
// 	}
// 	t.prSubs = "sync_" + t.outFullName + "_.srt"
// 	return "sync_" + t.outFullName + "_.srt"
// }

// func initStatus() {
// 	args := dataLineArgs(activeTask)
// 	newDl := "TaskStatus-0 " + strings.Join(args, " ")
// 	editActiveTaskLine(newDl)
// }

// func editActiveTaskLine(newLine string) {
// 	utils.EditLineInFile(taskFilePath, activeTask, newLine)
// }

// func newTask(dataLine string) *Task {
// 	task := &Task{}
// 	task.dataLine = dataLine
// 	task.outputTags = make(map[string]string)
// 	args := dataLineArgs(activeTask)
// 	err := checkArgs(args)
// 	if err != nil {
// 		if err.Error() == "Status Unknown" {
// 			initStatus()
// 		}
// 		if err.Error() != "Task Done" {
// 			fmt.Println(args)
// 			fmt.Println(err)
// 		}
// 		task.status = -1
// 		return task
// 	}
// 	fmt.Println("activeTask", activeTask)
// 	task.outDurat = args[2]
// 	task.offset = args[3]
// 	if fileAvailableM(inFolder, args[1]) {
// 		task.inVideo = args[1]
// 		task.inDurat = videoDuration(inFolder, args[1])
// 	}

// 	task.outputTags = mapTags(args[4])

// 	task.outBaseName = task.outputTags["name"] + "_" +
// 		task.outputTags["year"] + "__" +
// 		task.outputTags["resolutionTag"]
// 	task.prVideo = "cut_" + task.outBaseName + ".mp4"
// 	task.outFullName = args[4]
// 	task.predictAudioFiles()
// 	task.predictProcessFiles()
// 	task.predictSubsFile()
// 	//fmt.Println("Test status", readStatus(dataLine))
// 	//task.Info()

// 	return task
// }

// func (t *Task) CheckStatus() {
// 	fmt.Println(t.outFullName, "Status:", t.status)
// 	if t.status < 0 {
// 		return
// 	}
// 	time.Sleep(time.Second)
// 	if t.status == 0 {
// 		t.Info()
// 		if !fileAvailableM(inFolder, t.inVideo) {
// 			return
// 		}
// 		if fileAvailableM(outFolder, t.prVideo) {
// 			if videoDuration(outFolder, t.prVideo) == t.outDurat {
// 				fmt.Println(outFolder+"\\"+t.prVideo, "VALID")
// 				fmt.Println(t.outFullName, "set status 1")
// 				changeStatusInFile(taskCon1, t.outFullName)
// 				t.status = 1
// 			}
// 		}
// 	}
// 	pr := printer{defaultColor: "red"}
// 	if t.status == 1 {
// 		for i, audio := range t.prAudio {
// 			if !fileAvailableM(outFolder, t.prAudio[i]) {
// 				fmt.Println(audio, "is not available!")
// 			}
// 			fmt.Println(outFolder+"\\"+t.prAudio[i], utils.ASCIIColor("green", "VALID"))
// 			pr.PrintLine("Testststst")
// 		}
// 	}

// 	panic(t.status)
// }

// type cPrinter interface {
// 	PrintLine(string)
// 	SetColor(string)
// }

// type printer struct {
// 	defaultColor string
// }

// func (pr *printer) PrintLine(line string) {
// 	fmt.Println(utils.ASCIIColor(pr.defaultColor, line))
// }

// func (task *Task) predictProcessFiles() {
// 	if task.outBaseName == "" {
// 		panic("BASE NAME")
// 	}
// }

// func (task *Task) args() []string {
// 	args := strings.Split(task.dataLine, " ")
// 	return args
// }

// func readStatus(dataline string) int {
// 	newStatus := 0
// 	args := strings.Split(dataline, " ")
// 	if len(args) < 1 {
// 		return -1
// 	}
// 	if !strings.Contains(args[0], "TaskCon") {
// 		return -2
// 	}
// 	return newStatus
// }

// func (task *Task) Info() {
// 	fmt.Println("dataLine    string", task.dataLine)
// 	fmt.Println("inVideo     string", task.inVideo)
// 	fmt.Println("inDurat     string", task.inDurat)
// 	fmt.Println("offset      string", task.offset)
// 	fmt.Println("prVideo     string", task.prVideo)
// 	fmt.Println("prAudio     []string", task.prAudio)
// 	fmt.Println("prSubs      string", task.prSubs)
// 	fmt.Println("outBaseName string", task.outBaseName)
// 	fmt.Println("outFullName string", task.outFullName)
// 	fmt.Println("outDurat    string", task.outDurat)
// 	fmt.Println("status      int", task.status)
// }

// func taskFileReadable() bool {
// 	if !fileAvailableM(inFolder, taskFile) {
// 		return false
// 	}
// 	return true
// }

// func dataLineArgs(i int) []string {
// 	lines := utils.LinesFromTXT(taskFilePath)

// 	args := strings.Split(lines[i], " ")
// 	return args
// }

// func statusValid(status string) bool {
// 	switch status {
// 	default:
// 		return false
// 	case taskCon0, taskCon1, taskCon2, taskCon3, taskCon4, taskCon5, taskCon6, taskDone:
// 		return true
// 	}
// }

// func resolutionTagWL() []string {
// 	return []string{"sd", "hd", "3d"}
// }

// func audioTagWL() []string {
// 	return []string{"ar2e2", "ar6e2", "ar2e6", "ar6e6",
// 		"ar6", "ar2",
// 		"ae6", "ae2",
// 	}
// }

// func subsTagWL() []string {
// 	return []string{"sr", "[NO_SUBS]"}

// }

// func validateTag(tag string, whiteList []string) bool {
// 	for i := range whiteList {
// 		if tag == whiteList[i] {
// 			return true
// 		}
// 	}
// 	return false
// }

// func checkArgs(args []string) error {
// 	if len(args) == 4 {
// 		return errors.New("Status Unknown")
// 	}

// 	if len(args) != 5 {
// 		return errors.New("Error: Task " + strconv.Itoa(activeTask) + " have " + strconv.Itoa(len(args)) + " arguments (expecting 5)")
// 	}
// 	if !statusValid(args[0]) {
// 		return errors.New("Error: Task " + strconv.Itoa(activeTask) + " have INVALID status")
// 	}
// 	if args[0] == taskDone {
// 		return errors.New("Task Done")
// 	}
// 	fmt.Println("StatusArg:", args[0], "ok")

// 	fmt.Println("ResultArg:", args[0])

// 	// tn, err := tagname.NewFromFilename(args[4], tagname.CheckNormal)
// 	// if err != nil {
// 	// 	if err.Error() == "\x22sdhd\x22 type does not exist" {
// 	// 		return errors.New("Error: Task " + strconv.Itoa(activeTask) + " sdhd tag INVALID")
// 	// 	}
// 	// 	return errors.New("Error: Task " + strconv.Itoa(activeTask) + " " + err.Error())
// 	// }

// 	// atag, err := tn.GetTag("atag")
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	// fmt.Println(atag)
// 	//tag, errsdhd := tn.GetTag("sdhd")
// 	//fmt.Println("-------------tag", tag)
// 	// if errsdhd != nil {
// 	// 	panic("New Error")
// 	// 	return errors.New("Error: Task " + strconv.Itoa(activeTask) + " tag 'sdhd' not detected")
// 	// }
// 	// fmt.Println(tag)
// 	//fmt.Println(atag)

// 	if !isTimeStamp(args[2]) {
// 		return errors.New("Error: Task " + strconv.Itoa(activeTask) + " - args[2] [" + args[2] + "] is INVALID")
// 	}

// 	if !isTimeStamp(args[3]) {
// 		return errors.New("Error: Task " + strconv.Itoa(activeTask) + " - args[3] [" + args[3] + "] is INVALID")
// 	}

// 	if !fileAvailableM(inFolder, args[1]) {
// 		return errors.New("Error: Task " + strconv.Itoa(activeTask) + " - args[1] [" + args[1] + "] inFile is not available")
// 	}

// 	return nil
// }

// func (t *Task) Status() int {
// 	return t.status
// }

// func (t *Task) SetStatus(st int) {
// 	t.status = st
// }

// func main2() {
// 	// file := "word1_word2_partNum_0000__hd_ar2e6_sr"
// 	// tn, err := tagname.NewFromFilename(file, tagname.CheckNormal)
// 	// tag, err2 :=tn.GetTag("atag")
// 	// fmt.Println(tag)
// 	// fmt.Println(err)
// 	// fmt.Println(err2)
// 	// return
// 	preCheck()
// 	if !taskFileReadable() {
// 		fmt.Println("TaskFile is not readable...")
// 		fmt.Println("Resolve and restart")
// 		return
// 	}
// 	for i := 0; i < 1000; i++ {
// 		time.Sleep(time.Second)
// 		utils.ClearScreen()
// 		curTime := time.Now()
// 		activeTask = 0
// 		fmt.Println("cycle", i)
// 		fmt.Println("Start time :", startTime)
// 		fmt.Println("Curent time:", curTime)
// 		fmt.Println("")
// 		dataLines := dataLines()
// 		for dl := range dataLines {
// 			fmt.Println("	Task", dl, dataLines[dl])
// 			activeTask = dl
// 			task := newTask(dataLines[dl])
// 			task.CheckStatus()

// 		}

// 	}

// 	return
// 	// err = os.Mkdir("subdir", 0755)
// 	// check(err)

// 	// err = os.Mkdir("f:\\Work\\petr_proj\\___IN\\Test", 0755)

// 	// defer os.RemoveAll("subdir")

// 	// createEmptyFile := func(name string) {
// 	// 	d := []byte("")
// 	// 	check(ioutil.WriteFile(name, d, 0644))
// 	// }

// 	// createEmptyFile("subdir/file1")

// 	// err = os.MkdirAll("subdir/parent/child", 0755)
// 	// check(err)

// 	// createEmptyFile("subdir/parent/file2")
// 	// createEmptyFile("subdir/parent/file3")
// 	// createEmptyFile("subdir/parent/child/file4")

// 	// c, err := ioutil.ReadDir("subdir/parent")
// 	// check(err)

// 	// fmt.Println("Listing subdir/parent")
// 	// for _, entry := range c {
// 	// 	fmt.Println(" ", entry.Name(), entry.IsDir())
// 	// }

// 	// err = os.Chdir("subdir/parent/child")
// 	// check(err)

// 	// c, err = ioutil.ReadDir(".")
// 	// check(err)

// 	// fmt.Println("Listing subdir/parent/child")
// 	// for _, entry := range c {
// 	// 	fmt.Println(" ", entry.Name(), entry.IsDir())
// 	// }

// 	// err = os.Chdir("../../..")
// 	// check(err)

// 	// fmt.Println("Visiting subdir")
// 	// err = filepath.Walk("subdir", visit)
// 	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// fmt.Println(dir)
// 	// _, filename, _, ok := runtime.Caller(0)
// 	// if !ok {
// 	// 	panic("No caller information")
// 	// }
// 	// fmt.Printf("Filename : %q, Dir : %q\n", filename, path.Dir(filename))
// }

// func visit(p string, info os.FileInfo, err error) error {
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println(" ", p, info.IsDir())
// 	return nil
// }

// func fileAvailableM(folder string, file string) bool {
// 	fmt.Println("checking:", folder+"\\"+file)
// 	err := os.Rename(folder+"\\"+file, folder+"\\"+"RENAMED_"+file)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return false
// 	}
// 	os.Rename(folder+"\\"+"RENAMED_"+file, folder+"\\"+file)
// 	return true
// }

// func fileAvailable(file string) bool {
// 	fmt.Println("checking:", file)
// 	err := os.Rename(file, "RENAMED_"+file)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return false
// 	}
// 	os.Rename("RENAMED_"+file, file)
// 	return true
// }

// // exists returns whether the given file or directory exists
// func fileExists(folder, file string) bool {
// 	_, err := os.Stat(folder + "\\" + file)
// 	if err == nil {
// 		return true
// 	}
// 	return false
// }

// func dataLines() (result []string) {
// 	lines := utils.LinesFromTXT(taskFilePath)
// 	for _, dataLine := range lines {
// 		result = append(result, dataLine)
// 	}
// 	return result
// }

// /*
// Status:
// TaskCon

// TaskCon			| 1 | 2 | 3 | 4 | 5 |
// конечная задача	| + | + | + | + | + | - readTask() - есть ли ошибки в чтении таска
// ИСХвидео		| - | + | + | + | + | - checkInput() - свободен ли видеофайл InVideo в папке INfolder
// ОБРвидео		| - | - | + | + | + | - checkPrVideo() - свободен ли видеофайл shortname.mp4 в папке PRfolder
// ОБРзвук 		| - | - | - | + | + | - checkAudio() - свободен ли аудиофайл/ы shortname.aac в папке PRfolder
// ОБРсабы			| - | - | - | - | + | - checkSubs() - свободен ли файл сабов в папке PRfolder с тегом Sync

// TaskCon-0
// readTask():
// получает ShortName; tags; имена предмуксовых видео и аудио; проверка (длинну файла)
// TaskCon-1
// TaskCon-2
// TaskCon-3
// TaskCon-4
// TaskCon-5

// */

// func videoDuration(folder, file string) string {
// 	if !fileAvailableM(folder, file) {
// 		return "00:00:00.00"
// 	}
// 	cmd := exec.Command("ffmpeg", "-i", folder+"\\"+file)
// 	output, _ := cmd.CombinedOutput()
// 	stringOUT := string(output)
// 	str1 := strings.Split(stringOUT, "Duration: ")
// 	time.Sleep(time.Second)
// 	if len(str1) > 0 {
// 		durationSTR := strings.Split(str1[1], ", ")
// 		return ffToPrem(durationSTR[0])
// 	}
// 	return "00:00:00.00"
// }

// func ffToPrem(duration string) string {
// 	premiereDur := ""
// 	timeArgs := strings.Split(duration, ":")
// 	hours, errHour := strconv.Atoi(timeArgs[0])
// 	if errHour != nil {
// 		panic(errHour)
// 	}
// 	minutes, errMin := strconv.Atoi(timeArgs[1])
// 	if errMin != nil {
// 		panic(errMin)
// 	}
// 	secs, err := strconv.ParseFloat(timeArgs[2], 64)
// 	if err != nil {
// 		panic(err)
// 	}
// 	secsInt := int(secs)
// 	strHour := strconv.Itoa(hours)
// 	if hours < 10 {
// 		strHour = "0" + strHour
// 	}
// 	strMin := strconv.Itoa(minutes)
// 	if minutes < 10 {
// 		strMin = "0" + strMin
// 	}
// 	strSec := strconv.Itoa(secsInt)
// 	if secsInt < 10 {
// 		strSec = "0" + strSec
// 	}
// 	part := secs
// 	for part > 1 {
// 		part = part - 1
// 	}
// 	part = toFixed(part, 3)
// 	part = part / 0.04
// 	frame := int(part)
// 	strFrame := strconv.Itoa(frame)
// 	if frame < 10 {
// 		strFrame = "0" + strFrame
// 	}
// 	if frame == 25 {
// 		strFrame = "00"
// 	}
// 	premiereDur = strHour + ":" + strMin + ":" + strSec + ":" + strFrame
// 	return premiereDur
// }

// func floatToString(input_num float64) string {
// 	// to convert a float number to a string
// 	return strconv.FormatFloat(input_num, 'f', 2, 64)
// }

// func round(num float64) int {
// 	return int(num + math.Copysign(0.5, num))
// }

// func toFixed(num float64, precision int) float64 {
// 	output := math.Pow(10, float64(precision))
// 	return float64(round(num*output)) / output
// }

// func premToFF(duration string) string {
// 	parts := strings.Split(duration, ":")
// 	secsInt, _ := strconv.Atoi(parts[2])
// 	partsInt, _ := strconv.Atoi(parts[3])
// 	partsFl := float64(partsInt)*40/1000 + float64(secsInt)
// 	sec := floatToString(partsFl)
// 	return parts[0] + ":" + parts[1] + ":" + sec
// }

// func premToFrames(duration string) int {
// 	parts := strings.Split(duration, ":")
// 	hour, _ := strconv.Atoi(parts[0])
// 	min, _ := strconv.Atoi(parts[1])
// 	sec, _ := strconv.Atoi(parts[2])
// 	frm, _ := strconv.Atoi(parts[3])
// 	return frm + 25*sec + 1500*min + 90000*hour
// }

// func exe_cmd(cmd string, wg *sync.WaitGroup) {
// 	fmt.Println("command is ", cmd)
// 	// splitting head => g++ parts => rest of the command
// 	parts := strings.Fields(cmd)
// 	head := parts[0]
// 	parts = parts[1:len(parts)]

// 	out, err := exec.Command(head, parts...).Output()
// 	if err != nil {
// 		fmt.Printf("%s", err)
// 	}
// 	fmt.Printf("%s", out)
// 	wg.Done() // Need to signal to waitgroup that this goroutine is done
// }

// func isTimeStamp(arg string) bool {
// 	match, _ := regexp.MatchString("[0-9][0-9]:[0-9][0-9]:[0-9][0-9]:[0-9][0-9]", arg)
// 	//     if err != nil {
// 	// fmt.Println(err.Error())
// 	// return false
// 	// 	}
// 	return match
// }

// func loadTaskFileBuffer() []string {
// 	return dataLines()
// }

// func saveTaskFile(buffer []string) {
// 	for i, val := range buffer {
// 		utils.EditLineInFile(taskFilePath, i, val)
// 	}
// }

// func currentTaskArgs(outputName string) ([]string, int) {
// 	buffer := loadTaskFileBuffer()
// 	args := []string{}
// 	iPos := 0
// 	for i, val := range buffer {
// 		if strings.Contains(val, outputName) {
// 			args = dataLineArgs(i)
// 			iPos = i
// 		}
// 	}
// 	return args, iPos
// }

// func changeStatusTF(argsOld []string, newStatus string) string {
// 	argsOld[0] = newStatus
// 	return strings.Join(argsOld, " ")
// }

// func currentDataLine(outputName string) (data string, num int) {
// 	buffer := loadTaskFileBuffer()
// 	for i, val := range buffer {
// 		if strings.Contains(val, outputName) {
// 			return val, i
// 		}
// 	}
// 	return "[LINE NOT FOUND]", -1
// }

// func changeStatusInFile(newStatus string, outName string) {
// 	args, i := currentTaskArgs(outName)
// 	newData := changeStatusTF(args, newStatus)
// 	buffer := loadTaskFileBuffer()
// 	buffer[i] = newData
// 	saveTaskFile(buffer)
// }
