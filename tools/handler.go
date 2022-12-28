package tools

import (
	"fmt"
	"log"
	"time"
)

type handler interface {
	Err(err error, callback func(), sleep time.Duration) bool
	Retry(error,
		int,
		time.Duration,
		func() error,
		func(),
		func()) bool
}
type handle struct {
}

func NewHandle() *handle {
	return &handle{}
}

// Usage:
//
//	if Pkgtools.NewHandle().Err(err, func(){}, time.Second*0){
//		dosomething...
//		continue
//	}
func (h *handle) Err(err error, callback func(), sleep time.Duration) bool {
	if err != nil {
		callback()
		log.Println(err)
		time.Sleep(sleep)
		return true
	}
	return false
}

// Usage:
//
//	if Pkgtools.NewHandle().Retry(err, 5, time.Second*0, DB.Conn,
//		func() {
//			log.Println("somthing alive")
//		}, func() {
//			log.Println("somthing dead already")
//		}) {
//		return
//	}
func (h *handle) Retry(entryErr error, retrys int, sleep time.Duration, tryFunc func() error, success func(), failure func()) (stillErr bool) {

	if entryErr != nil {
		log.Println(entryErr)
		fmt.Println("Start retry procedure")
		for i := 1; i <= retrys; i++ {
			err := tryFunc()
			if err != nil {
				log.Println("still error:", err, "retry:", i)
				time.Sleep(sleep)
				continue
			}
			success()
			return false
		}
		failure()
		return true
	}
	return false
}
