package logger

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

var (
	Info            *log.Logger
	ReqestInfo      *log.Logger
	SubmissionInfo  *log.Logger
	Error           *log.Logger
	ReqestError     *log.Logger
	SubmissionError *log.Logger
)

func init() {
	file, err := os.OpenFile("sys.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Faild to open error logger file:", err)
	}

	Info = log.New(io.MultiWriter(file, os.Stderr), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ReqestInfo = log.New(io.MultiWriter(file, os.Stderr), "REQEST_INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	SubmissionInfo = log.New(io.MultiWriter(file, os.Stderr), "SUBMISSION_INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	ReqestError = log.New(io.MultiWriter(file, os.Stderr), "REQEST_ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	SubmissionError = log.New(io.MultiWriter(file, os.Stderr), "SUBMISSION_ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func ShowNormalAuxSubmission() {
	file, err := os.Open("sys.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lineText string
	var index int
	scanner := bufio.NewScanner(file)
	fmt.Println("---------[start print]-----------")

	for scanner.Scan() {
		lineText = scanner.Text()
		match, _ := regexp.MatchString(`^SUBMISSION_INFO`, lineText)
		if match != false {
			index++
			fmt.Println(lineText)
		}
	}
	fmt.Println("[total number}: ", index)
	fmt.Println("---------[end print]-----------")
}

func ShowNormalAuxRequest() {
	file, err := os.Open("sys.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lineText string
	var index int
	scanner := bufio.NewScanner(file)
	fmt.Println("---------[start print]-----------")

	for scanner.Scan() {
		lineText = scanner.Text()
		match, _ := regexp.MatchString(`^REQEST_INFO`, lineText)
		if match != false {
			index++
			fmt.Println(lineText)
		}
	}
	fmt.Println("[total number}: ", index)
	fmt.Println("---------[end print]-----------")
}

func ShowErrorAuxSubmission() {
	file, err := os.Open("sys.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lineText string
	var index int
	scanner := bufio.NewScanner(file)
	fmt.Println("---------[start print]-----------")
	for scanner.Scan() {
		lineText = scanner.Text()
		match, _ := regexp.MatchString(`^SUBMISSION_ERROR`, lineText)
		if match != false {
			index++
			fmt.Println(lineText)
		}
	}
	fmt.Println("[total number}: ", index)
	fmt.Println("---------[end print]-----------")
}

func ShowErrorAuxRequest() {
	file, err := os.Open("sys.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lineText string
	var index int
	scanner := bufio.NewScanner(file)
	fmt.Println("---------[start print]-----------")
	for scanner.Scan() {
		lineText = scanner.Text()
		match, _ := regexp.MatchString(`^REQEST_ERROR`, lineText)
		if match != false {
			index++
			fmt.Println(lineText)
		}
	}
	fmt.Println("[total number}: ", index)
	fmt.Println("---------[end print]-----------")
}
