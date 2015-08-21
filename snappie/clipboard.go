package main

import (
  "github.com/atotto/clipboard"
)

func CopyToClipboard(text string) {
  clipboard.WriteAll(text)
}
