package main

import (
  "os/exec"
  "github.com/deckarep/gosx-notifier"
)

func PlaySound(name string) {
  args := "/System/Library/Sounds/" + name + ".aiff"
  exec.Command("/usr/bin/afplay", args).Output()
}

func Notify(url string) {
  note := gosxnotifier.NewNotification("Url copied to clipboard")
  note.Title = "Uploaded"
  note.Link = url
  note.Group = "pl.bunsch.snappie"
  note.Sender = "pl.bunsch.snappie"
  note.Push()
}

