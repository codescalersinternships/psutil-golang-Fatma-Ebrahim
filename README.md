# psutil-golang-Fatma-Ebrahim

This repository contains an implementation of psutil-go, a lightweight library in Go that provides essential system information similar to [psutil](https://github.com/giampaolo/psutil) in Python.

## Functions

- **`CPUCoresNum(info Info) (string, error)`**
  Returns the number of CPU cores.

- **`CPUVendor(info Info) (string, error)`**
  Returns the CPU vendor name.

- **`CPUModel(info Info) (string, error)`**
  Returns the CPU model name.

- **`CPUcacheSize(info Info) (string, error)`**
  Returns the CPU cache size in KB.

- **`CPUMHZ(info Info) (string, error)`**
  Returns the CPU frequency in MHz.

- **`AvailMem(info Info) (string, error)`**
  Returns the available memory in KB.

- **`TotalMem(info Info) (string, error)`**
  Returns the total memory in KB.

- **`UsedMem(info Info) (string, error)`**
  Returns the used memory in KB.

- **`RunningProcesses(info Info) ([]ProcessDetails, error)`**
  Returns running processes, including their PIDs and names.

## How to Use

### Step 1: Install the Package Using `go get`

```shell
go get github.com/codescalersinternships/psutil-golang-Fatma-Ebrahim
```

This command fetches the package and adds it to your project's `go.mod` file.

### Step 2: Install the Needed Dependencies

```shell
go mod download
```

### Step 3: Import and Use the Package in Your Code

Now, you can import the package into your project and use the functions as described:

```go
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
```
