package main

import (
	"fmt"
	"log"

	psutil "github.com/codescalersinternships/psutil-golang-Fatma-Ebrahim/pkg"
)

func main() {
	ps := &psutil.PSUtil{}
	err := ps.NewPSUtil()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ps.RunningProcesses())

}
