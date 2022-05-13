package test

import (
	"log"
	"os"
	"strconv"
	"time"
)

func LogSettest() {
	log.SetFlags(log.Ltime)
	log.Println("hello,logsettest")

}

func Logfile() *log.Logger {
	datestr := time.Now().Format("2006-01-02")
	file_a := "./" + datestr + ".log"

	logg, errinfo := WriteFilelog(file_a)
	if errinfo != nil {
		logg.Output(2, errinfo.Error())
	} else {
		tips := time.Now().Format("2006-01-02 15:04:05") + "没有创建日期目录"
		logg.Output(2, tips)
		log.Println(tips)
	}

	hour := time.Now().Hour()
	path := "./" + datestr + "/"
	file_b := path + strconv.Itoa(hour) + ".log"

	_, err := os.Stat(path)
	if err != nil {
		ok := os.Mkdir(path, os.ModePerm)
		if ok != nil {
			logg.SetPrefix("make_")
			logg.Output(2, ok.Error())
		}
	}

	logg, err_b := WriteFilelog(file_b)
	if err_b != nil {
		panic(err_b)
	}

	return logg
}

func LogFileSet(logg *log.Logger) *log.Logger {
	logg.SetFlags(log.Ldate | log.Ltime | log.Llongfile) //默认输出选项
	// logg.Output(2, "ddd")
	return logg
}

func WriteFilelog(file string) (*log.Logger, error) {
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		return nil, err
	}

	logg := log.New(logFile, "pre_", log.Lshortfile)

	return logg, err

}
