package test

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func LogSettest() {
	log.SetFlags(log.Ltime)
	log.Println("hello,logsettest")

}

func Debug(logName string, tips string) {
	logFile, err := os.OpenFile(logName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("create %s err: %v \n ", logName, err)
	}
	if logFile != nil {
		defer func(file *os.File) {
			file.Close()
		}(logFile)
	}

	debugLog := log.New(logFile, "[Debug]", log.Ldate)

	debugLog.SetPrefix("[Debug]")
	debugLog.SetFlags(log.Lshortfile)
	debugLog.Println(tips)

}

func Logfile() (*log.Logger, error) {
	datestr := time.Now().Format("2006-01-02")
	file_a := "./" + datestr + ".log"

	logFile, err := os.OpenFile(file_a, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		return nil, err
	}
	// if logFile != nil {
	// 	defer func() {
	// 		if err := logFile.Close(); err != nil {
	// 			return
	// 		}
	// 	}()
	// }

	logg := log.New(logFile, "pre_", log.Lshortfile)

	return logg, err
}

func LogFileD() (*log.Logger, error) {
	datestr := time.Now().Format("2006-01-02")
	// datetimestr := time.Now().Format("2006-01-02 15:04:05")
	hour := time.Now().Hour()
	path := "./log/" + datestr + "/"
	// path := "./" + datestr + "/"
	file_b := path + strconv.Itoa(hour) + ".log"

	_, err_a := os.Stat(path)
	if err_a != nil {
		// ok := os.Mkdir(path, os.ModePerm)
		ok := os.MkdirAll(path, os.ModePerm)
		if ok != nil {
			log.Println(ok.Error())
			panic(ok)
		}
	}

	logFile, err := os.OpenFile(file_b, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		return nil, err
	}
	// if logFile != nil {
	// 	defer func() {
	// 		if err := logFile.Close(); err != nil {
	// 			return
	// 		}
	// 	}()
	// }

	logg := log.New(logFile, "pre_", log.Lshortfile)

	return logg, err

}

func LogFileSet(logg *log.Logger) *log.Logger {
	logg.SetFlags(log.Ldate | log.Ltime | log.Llongfile) //默认输出选项
	// logg.Output(2, "ddd")
	return logg
}

func OpenFilelog(file string) (*log.Logger, error) {
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		return nil, err
	}

	logg := log.New(logFile, "pre_", log.Lshortfile)

	return logg, err

}
