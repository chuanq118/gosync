package main

import (
	"flag"
	cmn "gosync/common"
	"gosync/logger"
)

var mode = cmn.RunMode(*flag.String("mode", "watch", "watch or batch."))

func main() {
	flag.Parse()
	if mode == cmn.Batch {
		logger.Debugln("Running gosync in Batch mode.")
	} else if mode == cmn.Watch {
		logger.Debugln("Running gosync in Watch mode.")
	}
}
