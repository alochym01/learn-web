# Gin hello world - Day 01

## Install Auto Reload Air

1. `curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s`
2. air -v

## Create Folder structure

1. mkdir web-w-gin
2. cd web-w-gin
3. air init
4. go mod init github.com/alochym01/web-w-gin

## Create Entry Point file

1. touch [main.go](main.go)
2. Define Album struct

   ```go
   // Album represents data about a record Album.
   type Album struct {
       ID     int
       Title  string
       Artist string
       Price  float64
   }
   ```

## Using Auto Reload air

1. air

## Using curl to test

1. curl localhost:8080/albums