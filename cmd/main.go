package main

import (
	"fmt"
	"log"

	psutil "github.com/codescalersinternships/psutil-golang-Fatma-Ebrahim/pkg"
)

func main() {
	filesinfo := &psutil.FilesInfo{}
	ps, err := psutil.RunningProcesses(filesinfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ps)

}
