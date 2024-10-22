package logger

import (
	"log"
	"os"
)

type Logger interface {
	Debug(args ...any)
	Info(args ...any)
	Error(args ...any)
	Fatal(args ...any)
}

func Debug(args ...any) {
	d := log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	d.Println(args)
}

func Info(args ...any) {
	i := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	i.Println(args)
}

func Error(args ...any) {
	e := log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	e.Println(args)
}

func Fatal(args ...any) {
	f := log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	f.Fatal(args)
}
