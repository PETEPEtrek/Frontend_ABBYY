package helpers

import "log"

func Trace(s string) string {
	log.Println("entering:", s)
	return s
}

func Untrace(s string) {
	log.Println("leaving:", s)
}
