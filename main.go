package main

import (
  "os"
  "strings"
  "path"
  "github.com/howeyc/fsnotify"
)

func HandleEvent(ev *fsnotify.FileEvent) {

  var name string = path.Base(ev.Name)

  if ev.IsCreate() && strings.Contains(name, "Screen Shot") && !strings.HasPrefix(name, ".") {
    url := UploadFile(ev.Name)
    Notify(url)
    CopyToClipboard(url)
    PlaySound("Glass")
    os.Remove(ev.Name)
  }

}

func main() {

    config := LoadConfig()

    if config != nil {
      WatchChanges(GetDesktopLocation(), HandleEvent)
    }

}
