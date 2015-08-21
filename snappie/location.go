package main

import (
  "log"
  "os/user"
  "path"
)


func GetDesktopLocation() string {
  return path.Join(GetHomeLocation(), "Desktop")
}

func GetHomeLocation() string {
  usr, err := user.Current()

  if err != nil {
      log.Fatal(err)
  }

  return usr.HomeDir
}
