package main

import (
	"Week06/pkg"
	"fmt"
	"time"
)

func DebugRc(rc *pkg.RollingCounter) {
	fmt.Printf("Avg:%v Sum:%v Max:%v Min:%v TimeSpan:%v\n", rc.Avg(), rc.Sum(), rc.Max(), rc.Min(), rc.Timespan())
}

func main() {
	opts := pkg.RollingCounterOpts{
		Size:           3,
		BucketDuration: time.Duration(1 * time.Second),
	}
	rc := pkg.NewRollingCounter(opts)
	fmt.Printf("Add %v\n", 1.15)
	rc.Add(1.15)
	DebugRc(rc)
	time.Sleep(time.Duration(1 * time.Second))

	fmt.Printf("Add %v\n", 2.0)
	rc.Add(2.0)
	DebugRc(rc)

	time.Sleep(time.Duration(2 * time.Second))
	DebugRc(rc)

	fmt.Printf("Add %v\n", 1.33)
	rc.Add(1.33)
	DebugRc(rc)
	time.Sleep(time.Duration(1 * time.Second))
	DebugRc(rc)
}
