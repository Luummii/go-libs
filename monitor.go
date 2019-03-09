package libs

import (
	"encoding/json"
	"runtime"
	"time"
)

type monitor struct {
	Alloc,
	TotalAlloc,
	Sys,
	Mallocs,
	Frees,
	LiveObj,
	PauseTNs uint64

	NumGC  uint32
	NumGor int
}

func RunMonitor(duration int) {
	var m monitor
	var rtm runtime.MemStats
	var interval = time.Duration(duration) * time.Second

	for {
		<-time.After(interval)

		// read full mem stats
		runtime.ReadMemStats(&rtm)

		// number of goroutines
		m.NumGor = runtime.NumGoroutine()

		// misc memory stats
		m.Alloc = rtm.Alloc             // currently allocated number of bytes on the heap,
		m.TotalAlloc = rtm.TotalAlloc   // cumulative max bytes allocated on the heap (will not decrease),
		m.Sys = rtm.Sys                 // total memory obtained from the OS,
		m.Mallocs = rtm.Mallocs         // --
		m.Frees = rtm.Frees             // number of allocations, deallocations, and live objects (mallocs - frees),
		m.LiveObj = m.Mallocs - m.Frees // live objects = Mallocs - Frees

		// GC stats
		m.PauseTNs = rtm.PauseTotalNs //  total GC pauses since the app has started,
		m.NumGC = rtm.NumGC           // number of completed GC cycles

		// just encode to json and print
		b, _ := json.Marshal(m)
		Mnt.Println(string(b))
	}
}
