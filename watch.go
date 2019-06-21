package watch

import (
	"strings"

	"github.com/negbie/logp"
	"github.com/negbie/ngcp-cdr-db/config"
	"github.com/radovskyb/watcher"
)

var watchme *watcher.Watcher

// Start will watch for new files in folder
func Start(f func(string), shut chan struct{}) {
	setupWatcher()
	go func() {
		for {
			select {
			case event := <-watchme.Event:
				switch event.Op {
				case watcher.Create:
					if event.IsDir() {
						// We got a new folder, add it.
						if err := watchDir(event.Path, false); err != nil {
							logp.Err("%v", err)
							continue
						}
					} else {
						// We got a new file, scan it. Ignore .gz
						if strings.HasSuffix(event.Name(), ".cdr") {
							logp.Info("scan file %s", event.Path)
							f(event.Path)
						}
					}
				case watcher.Rename:
					logp.Info("%v", event)
				case watcher.Move:
					logp.Info("%v", event)
				case watcher.Remove:
					logp.Info("%v", event)
				}

			case err := <-watchme.Error:
				logp.Err("%v", err)
				// We got a new folder, add it.
				if err := watchDir(config.Setting.WatchFolder, config.Setting.WatchRecursive); err != nil {
					logp.Err("%v", err)
				}
				continue
			case <-shut:
				logp.Info("close watcher")
				watchme.Close()
				return
			}
		}
	}()

	// Start the watching process to check for changes every second.
	if err := watchme.Start(config.Setting.WatchTime); err != nil {
		logp.Err("%v", err)
	}
}

// watchDir will search for directories to add watchers to
func watchDir(folder string, recursive bool) error {
	if recursive {
		// Watch following folder recursively for changes.
		if err := watchme.AddRecursive(folder); err != nil {
			return err
		}
	} else {
		// Watch following folder for changes.
		if err := watchme.Add(folder); err != nil {
			return err
		}
	}
	return nil
}

func setupWatcher() {
	// Create a new watcher.
	watchme = watcher.New()
	// Ignore hidden files.
	watchme.IgnoreHiddenFiles(true)
	// SetMaxEvents to max to allow at most max event's to be received
	// on the Event channel per watching cycle. If SetMaxEvents is not set,
	// the default is to send all events.
	watchme.SetMaxEvents(config.Setting.WatchMaxEvent)
	// Only notify create file events.
	watchme.FilterOps(watcher.Create, watcher.Rename, watcher.Move, watcher.Remove)

	if err := watchDir(config.Setting.WatchFolder, config.Setting.WatchRecursive); err != nil {
		logp.Critical("%v", err)
	}
	logp.Info("start watcher in folder=%s, events=%d, frequency=%s ,recursive=%v",
		config.Setting.WatchFolder,
		config.Setting.WatchMaxEvent,
		config.Setting.WatchTime,
		config.Setting.WatchRecursive,
	)
}
