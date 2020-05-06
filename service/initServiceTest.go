package service

import (
	"flag"
	"log"
	"os/exec"
	"path/filepath"
	"runtime"
)

var logQuery = flag.Bool("log", false, "Print queries")
var binary = flag.Bool("binary", false, "Execute binary reseting")
var debug = flag.Bool("debug", false, "Change the level of log to the debug")
var runCounter uint64

func initServiceTest() (bool, bool) {
	if runCounter == 0 {
		_, dir, _, _ := runtime.Caller(0)
		if *binary {
			exeBinery(dir)
		} else {
			runMain(dir)
		}

	}
	runCounter++
	return *logQuery, *debug
}

func exeBinery(dir string) {
	exeFile := filepath.Join(filepath.Dir(dir), "..", "cmd", "testinsertion", "testinsertion")
	command := exec.Command(exeFile)
	if err := command.Run(); err != nil {
		log.Fatal("Error in reseting the database\n", err)
	}
}

func runMain(dir string) {
	exeFile := filepath.Join(filepath.Dir(dir), "..", "cmd", "testinsertion", "main.go")
	command := exec.Command("sh", "-c", "go run "+exeFile)
	if err := command.Run(); err != nil {
		log.Fatal("Error in reseting the database\n", err)
	}
}
