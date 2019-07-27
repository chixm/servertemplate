package main

import (
	Encoder "encoding/base64"
	_ "log"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
)

//
// Create UniqueID with hashed hostname
// so even this code worked in multiple servers,
// ID does not duplicate between servers.
// https://github.com/chilts/sid/blob/master/sid.go

var lastTime int64
var lastRand int64
var chars = make([]string, 11, 11)
var mu = &sync.Mutex{}

var hostName string

func initializeUniqueIDMaker() {
	if h, err := os.Hostname(); err == nil {
		hostName = h
	} else {
		panic(`Failed to get HostName `)
	}
}

func createUniqID() string {
	return createID()
}

func createID() string {
	// lock for lastTime, lastRand, and chars
	mu.Lock()
	defer mu.Unlock()

	now := time.Now().UTC().UnixNano()
	var r int64

	// if we have the same time, just inc lastRand, else create a new one
	if now == lastTime {
		lastRand++
	} else {
		lastRand = rand.Int63()
	}
	r = lastRand

	// remember this for next time
	lastTime = now

	id := toStr(now) + "-" + toStr(r) + `-` + hostHash()
	//log.Println(`created string::` + id)
	return id
}

const base64 string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz~"

func toStr(now int64) string {
	// now do the generation (backwards, so we just %64 then /64 along the way)
	for i := 10; i >= 0; i-- {
		index := now % 64
		chars[i] = string(base64[index])
		now = now / 64
	}

	return strings.Join(chars, "")
}

func hostHash() string {
	//log.Println(`host name is ` + hostname)
	s := Encoder.StdEncoding.EncodeToString([]byte(hostName))
	return s
}
