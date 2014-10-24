package main

import (
  "log"
  "github.com/howeyc/fsnotify"
)

func WatchChanges(directory string, callback func(*fsnotify.FileEvent)) {

    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }

    done := make(chan bool)

    // Process events
    go func() {
        for {
            select {
            case ev := <-watcher.Event:
                callback(ev)
            case err := <-watcher.Error:
                log.Println("error:", err)
            }
        }
    }()

    err = watcher.Watch(directory)
    if err != nil {
        log.Fatal(err)
    }

    <-done

    watcher.Close()

}
