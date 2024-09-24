package main

import (
	"github.com/codescalersinternships/psutil-golang-Fatma-Ebrahim/pkg"
)

func main() {	
ps:=&psutil.PSUtil{}
ps.GetCPUInfo()
}