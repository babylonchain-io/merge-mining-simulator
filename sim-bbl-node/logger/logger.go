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
	BlockInfo       *log.Logger
	Error           *log.Logger
	ReqestError     *log.Logger
	SubmissionError *log.Logger
	OutdatedError   *log.Logger
)

func init() {
	file, err := os.OpenFile("sys.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Faild to open error logger file:", err)
	}

	Info = log.New(io.MultiWriter(file, os.Stderr), "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ReqestInfo = log.New(io.MultiWriter(file, os.Stderr), "REQEST_INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	SubmissionInfo = log.New(io.MultiWriter(file, os.Stderr), "SUBMISSION_INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	BlockInfo = log.New(io.MultiWriter(file, os.Stderr), "BLOCK_INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	ReqestError = log.New(io.MultiWriter(file, os.Stderr), "REQEST_ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	SubmissionError = log.New(io.MultiWriter(file, os.Stderr), "SUBMISSION_ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	OutdatedError = log.New(io.MultiWriter(file, os.Stderr), "OUTDATED_ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
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

func ShowAll() {
	file, err := os.Open("sys.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lineText string
	var REQEST_ERROR_index int
	var REQEST_INFO_index int
	var SUBMISSION_INFO_index int
	var SUBMISSION_ERROR_index int
	var BLOCK_INFO_index int

	var REQEST_ALL_index int
	var SUBMISSION_ALL_index int
	var OUTDATED_ERROR_index int

	scanner := bufio.NewScanner(file)
	//fmt.Println("---------[start print]-----------")
	for scanner.Scan() {
		lineText = scanner.Text()

		match, _ := regexp.MatchString(`^REQEST_ERROR`, lineText)
		if match != false {
			REQEST_ERROR_index++
			REQEST_ALL_index++
		}

		REQEST_INFO_match, _ := regexp.MatchString(`^REQEST_INFO`, lineText)
		if REQEST_INFO_match != false {
			REQEST_INFO_index++
			REQEST_ALL_index++
		}

		SUBMISSION_INFO_match, _ := regexp.MatchString(`^SUBMISSION_INFO`, lineText)
		if SUBMISSION_INFO_match != false {
			SUBMISSION_INFO_index++
			SUBMISSION_ALL_index++
		}

		SUBMISSION_ERROR_match, _ := regexp.MatchString(`^SUBMISSION_ERROR`, lineText)
		if SUBMISSION_ERROR_match != false {
			SUBMISSION_ERROR_index++
			SUBMISSION_ALL_index++
		}

		BLOCK_INFO_match, _ := regexp.MatchString(`^BLOCK_INFO`, lineText)
		if BLOCK_INFO_match != false {
			BLOCK_INFO_index++
		}

		OUTDATED_ERROR_match, _ := regexp.MatchString(`^OUTDATED_ERROR`, lineText)
		if OUTDATED_ERROR_match != false {
			OUTDATED_ERROR_index++
		}
	}

	fmt.Println("BBL block generated: ", BLOCK_INFO_index, " sent: ", REQEST_ALL_index, " auxPoW received: ", SUBMISSION_ALL_index, " valid: ", SUBMISSION_INFO_index, " invalid: ", SUBMISSION_ERROR_index, "outdates: ", OUTDATED_ERROR_index)

	//fmt.Println("---------[end print]-----------")
}
