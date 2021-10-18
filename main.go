package main
import (
				"fmt"
				"bufio"
				"github.com/notandrewdev/morse-decoder/pkg"
				"log"
				"os"
				"strings"
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