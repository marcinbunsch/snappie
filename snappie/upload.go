package main

import (
  "time"
  "log"
  "path"
  "bytes"
  "io/ioutil"
  "github.com/jlaffaye/ftp"
  "github.com/jmcvetta/randutil"
  "github.com/mitchellh/goamz/aws"
  "github.com/mitchellh/goamz/s3"
)

func UploadFile(filename string) string {
  config := LoadConfig()
  use, _ := config.Get("use")
  var result string
  if use == "ftp" {
    result = UploadFileToFtp(filename)
  } else if use == "s3" {
    result = UploadFileToS3(filename)
  }
  return result
}

func UploadFileToFtp(filename string) string {
  config := LoadConfig()

  host, _ := config.Get("ftp.host")
  user, _ := config.Get("ftp.user")
  pass, _ := config.Get("ftp.pass")
  dir, _ := config.Get("ftp.dir")
  http, _ := config.Get("ftp.http")

  connection, err := ftp.Connect(host + ":21")

  connection.Login(user, pass)

  contents, err := ioutil.ReadFile(filename)

  if err != nil {
      log.Fatal(err)
  }

  hash, err := randutil.AlphaString(32)
  newName := hash + ".png"
  var targetPath string = path.Join(dir, newName)

  reader := bytes.NewReader(contents)
  connection.Stor(targetPath, reader)

  connection.Quit()

  url := http + newName

  return url

}

func UploadFileToS3(filename string) string {
  config := LoadConfig()

  key, _ := config.Get("s3.access_key_id")
  secret, _ := config.Get("s3.secret_access_key")
  bucket_name, _ := config.Get("s3.bucket")
  http, _ := config.Get("s3.http")

  auth, _ := aws.GetAuth(key, secret)

  client := s3.New(auth, aws.EUWest)

  bucket := client.Bucket(bucket_name)

  contents, _ := ioutil.ReadFile(filename)

  prefix := time.Now().Format("2006/01/02/")
  hash, _ := randutil.AlphaString(32)
  newName := prefix + hash + ".png"

  bucket.Put(newName, contents, "image/png", s3.PublicRead)

  url := http + newName

  return url

}

