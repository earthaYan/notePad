package main

import (
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/earthaYan/notePad/internal/apiserver"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	apiserver.NewApp("notepad-apiserver").Run()
}
