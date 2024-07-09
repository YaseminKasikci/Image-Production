package main

import (
	"flag"
	"fmt"
	"imgproc/filter"
	"imgproc/task"
	"time"
)

func main() {
	//CLI
	var srcDir = flag.String("src", "", "Input directory")
	var dstDir = flag.String("dst", "", "Output directory")
	var filterType = flag.String("filter", "grayscale", "grayscale/blur")
	var taskType = flag.String("task", "waitgrp","waitgrp/channel")
	var PoolSize = flag.Int("poolsize", 4, "Worker pool size for the channel task")
	flag.Parse()

	var f filter.Filter
	switch *filterType {
		case "grayscale":
			f = filter.Grayscale{}
		case "blur":
			f =filter.Blur{}
	}

	var t task.Tasker
	switch *taskType {
	case "waitgrp":
		t= task.NewWaitGrpTask(*srcDir, *dstDir, f)
	case "channel":
		t = task.NewChanTask(*srcDir, *dstDir, f, *PoolSize)
	}


	start := time.Now()
	t.Process()
	elapsed := time.Since(start)
	fmt.Printf("Image processing took %s\n", elapsed)

}