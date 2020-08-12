package main

import (
	"crypto/md5"
	"log"
)

func main() {
	log.Println(md5.New())
}
