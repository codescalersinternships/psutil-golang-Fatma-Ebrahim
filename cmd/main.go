package main

import (
	"fmt"
	"log"

	psutil "github.com/codescalersinternships/psutil-golang-Fatma-Ebrahim/pkg"
)

func main() {
	ps, err := psutil.RunningProcesses()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ps)
}
