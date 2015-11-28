/*
* Copyright (c) 2015, zheng-ji.info
* */

package gophone

import (
	"fmt"
	"testing"
	"time"
)

func since(t time.Time) int {
	return int(time.Since(t).Nanoseconds() / 1e6)
}

func TestFindPhone(t *testing.T) {

	timeStart := time.Now()
	pr, err := Find("13580198235123123213213")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(since(timeStart))

	timeStart = time.Now()
	pr, err = Find("15813581745")
	if err == nil {
		fmt.Println(pr)
	}
	fmt.Println(since(timeStart))
}
