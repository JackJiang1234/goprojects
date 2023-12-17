package log

import (
	"log"
	"os"
)

func init() {
	log.SetPrefix("trace:")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

	file, err := os.OpenFile("errors.txt", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}
	Warning = log.New(os.Stdout, "WARNING: ", log.Ldate | log.Ltime | log.Lshortfile)
	Error = log.New(file, "Error: ", log.Ldate | log.Ltime | log.Lshortfile)
}

func usage() {
	log.Println("message")
	//log.Fatalln("fatal message")
	//log.Panicln("panic message")
	Warning.Println("There is something you need to know about")
	Error.Println("Something is wrong")
}

var (
	Error *log.Logger
	Warning *log.Logger
)

