package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"https://github.com/AllensYgl/truss-patch/linux"
	"https://github.com/AllensYgl/truss-patch/windows"
)

var (
	help    bool
	version bool
	path    string
)

func init() {

	pt := os.Getenv("GOPATH")
	pt += "/src/github.com/tuneinc/truss"

	flag.BoolVar(&help, "h", false, "this help")
	flag.BoolVar(&version, "v", false, "show version")
	flag.StringVar(&path, "p", pt, "set truss path manually")
}

func main() {
	flag.Parse()
	if help {
		showUsage()
	} else if version {
		showVersion()
	} else {
		run()
	}
}

func showUsage() {
	fmt.Print(`
truss_tools version: truss-patch/1.0
Usage: truss_tools [-h help] [-v version][-p truss_filePath]
Options:
`)
	flag.PrintDefaults()
	fmt.Print(`
	`)
}

func showVersion() {
	fmt.Print("version: truss-patch/1.0")
}

func showPath() {
	fmt.Println("path is: ", path)
}

func run() {
	fmt.Println("start patching")
	env := runtime.GOOS
	if env == "windows" {
		fmt.Println("windows environment")
		showPath()
		windows.Option(path)
		windows.Windows(path)
	} else if env == "linux" {
		fmt.Println("linux environment")
		showPath()
		linux.Option(path)
		linux.Linux(path)
	} else {
		fmt.Println("unknown environment...quit")
	}
}
