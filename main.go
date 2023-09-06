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
		logger.Debugf("Running gosync in %s mode.", mode)
	} else if mode == cmn.Watch {
		logger.Debugf("Running gosync in %s mode.", mode)
	}
}
