package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func tmpTrash() {

	var wg sync.WaitGroup

	outDir := "/tmp/"

	t, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("error, time parameter missing")
	}

	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	falcoTracer := NewFalcoTracer()

	falcoTracer.setupConnection()

	falcoTracer.loadRulesFromFalco()

	falcoTracer.flushFalcoData()

	wg.Add(1)
	go falcoTracer.loadStatsFromFalco(time.Duration(t), &wg)

	<-sigs

	falcoTracer.exitFlag = true

	wg.Wait()

	jsonStats, err := falcoTracer.MarshalJSON()
	if err != nil {
		log.Fatal("error in object marshaling")
	}

	falcoTracer.statsAggregator.sortAvgSlices()

	writeJSONOnFile(jsonStats, outDir)
}

func writeJSONOnFile(jsonStats []byte, outDir string) {
	f, err := os.Create(outDir + "tracer_data.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := f.Write(jsonStats)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
