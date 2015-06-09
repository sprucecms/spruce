package main

import (
	"github.com/sprucecms/spruce/api"
	"github.com/sprucecms/spruce/cms"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	api.Start()
	cms.Start()
}
