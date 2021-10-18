package main

import (
	"log"
)

func removeEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}

	return r
}

func main() {
	log.Println("nothing much here")
}
