package watch

import (
	"github.com/fsnotify/fsnotify"
	"gosync/logger"
	"os"
	"path/filepath"
)

var printer = logger.Log.Sugar()

func recursiveAdd(name string, watcher *fsnotify.Watcher) error {
	files, err := os.ReadDir(name)
	if err != nil {
		return err
	}
	err = watcher.Add(name)
	if err != nil {
		return err
	}
	for _, f := range files {
		if f.IsDir() {
			dirName := filepath.Join(name, f.Name())
			err = watcher.Add(dirName)
			if err != nil {
				return err
			}
			err = recursiveAdd(dirName, watcher)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func watch(base string) error {
	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				printer.Infoln("event:", event)
				// 处理新文件被创建
				if event.Has(fsnotify.Create) {
					// 判断是否为文件夹
					stat, err := os.Stat(event.Name)
					if err != nil {
						printer.Fatal(err)
						return
					}
					if stat.IsDir() {
						err = watcher.Add(event.Name)
						if err != nil {
							printer.Fatal(err)
							return
						}
					}
				} else if event.Has(fsnotify.Write) {

				} else if event.Has(fsnotify.Rename) {
					
				} else if event.Has(fsnotify.Remove) {

				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				printer.Error("error:", err)
			}
		}
	}()

	// Add a path.
	err = recursiveAdd("C:\\Users\\legen\\Documents\\temp\\sync", watcher)
	if err != nil {
		return err
	}

	// Block main goroutine forever.
	<-make(chan struct{})

	return nil
}
