package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/howeyc/fsnotify"
)

func watchFolder(path string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fatal(err)
	}

	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				if isWatchedFile(ev.Name) {
					watcherLog("sending event %s", ev)
					startChannel <- ev.String()
				}
			case watcherErr := <-watcher.Error:
				watcherLog("error: %s", watcherErr)
			}
		}
	}()

	err = watcher.Watch(path)

	if err != nil {
		fatal(err)
	}
}

func watch() {
	paths := watchPaths()
	for i := range paths {
		wp := strings.TrimSpace(paths[i])

		_, err := os.Stat(wp)
		if err != nil {
			watcherLog("Path not exists: %s", wp)
			continue
		}

		watcherLog("Watching %s", wp)
		filepath.Walk(wp, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() && !isTmpDir(path) {
				if len(path) > 1 && strings.HasPrefix(filepath.Base(path), ".") {
					return filepath.SkipDir
				}

				if strings.Contains(path, "vendor/") {
					return filepath.SkipDir
				}

				watchFolder(path)
			}

			return err
		})
	}
}
