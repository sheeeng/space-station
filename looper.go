// Ctrl-C (in older Unixes, DEL) sends an INT signal (SIGINT); by default, this causes the process to terminate.
// Capture a Ctrl+C signal and run a cleanup function in Golang.
// http://stackoverflow.com/a/18158859/4763512

// It seems Ctrl-C doesn't work on PTY-based terminal emulators (such as mintty) on Windows unless the program is compiled by Cygwin.
// http://grokbase.com/p/gg/golang-nuts/148tgdtw7k/go-nuts-my-app-dont-end-on-ctrl-c-aka-sigint

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime/pprof"
	"time"
)

func main() {
	// Capture ctrl+c and stop CPU profiler.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			log.Printf("Captured %v, stopping profiler & exiting...\n",
				sig)
			pprof.StopCPUProfile()
			os.Exit(1)
		}
	}()

	for {
		fmt.Printf("%v\n", "#")
		time.Sleep(time.Second)
	}
}
