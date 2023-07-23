package metrics

import (
	"expvar"
	"math/rand"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/mem"
)

type MetricMapUint32 map[string]uint32

type MetricMapUint64 map[string]uint64

type MetricMapFloat64 map[string]float64

type MetricStorage struct {
	Timestamp        int64            `json:"timestamp"`
	MetricMapUint32  MetricMapUint32  `json:"uint32"`
	MetricMapUint64  MetricMapUint64  `json:"uint64"`
	MetricMapFloat64 MetricMapFloat64 `json:"float64"`
}

func Get() (mMap MetricStorage, err error) {
	rand.Seed(time.Now().UnixNano())
	memstatsFunc := expvar.Get("memstats").(expvar.Func)
	memstats := memstatsFunc().(runtime.MemStats)
	v, err := mem.VirtualMemory()
	if err != nil {
		return
	}

	mMap = MetricStorage{
		Timestamp: time.Now().Unix(),
		MetricMapUint32: MetricMapUint32{
			"NumForcedGC": memstats.NumForcedGC,
			"NumGC":       memstats.NumGC,
		},
		MetricMapUint64: MetricMapUint64{
			"BuckHashSys": memstats.BuckHashSys,
			"Frees":       memstats.Frees,
			"GCSys":       memstats.GCSys,

			"HeapAlloc":    memstats.HeapAlloc,
			"HeapIdle":     memstats.HeapIdle,
			"HeapInuse":    memstats.HeapInuse,
			"HeapObjects":  memstats.HeapObjects,
			"HeapReleased": memstats.HeapReleased,
			"HeapSys":      memstats.HeapSys,

			"LastGC":  memstats.LastGC,
			"Lookups": memstats.Lookups,

			"MCacheInuse": memstats.MCacheInuse,
			"MCacheSys":   memstats.MCacheSys,
			"MSpanInuse":  memstats.MSpanInuse,
			"MSpanSys":    memstats.MSpanSys,

			"Mallocs":      memstats.Mallocs,
			"NextGC":       memstats.NextGC,
			"OtherSys":     memstats.OtherSys,
			"PauseTotalNs": memstats.PauseTotalNs,

			"StackInuse": memstats.StackInuse,
			"StackSys":   memstats.StackSys,

			"Alloc":       memstats.Alloc,
			"Sys":         memstats.Sys,
			"TotalAlloc":  memstats.TotalAlloc,
			"RandomValue": rand.Uint64(),

			// Memory
			"TotalMemory": v.Total,
			"FreeMemory":  v.Free,
		},
		MetricMapFloat64: MetricMapFloat64{
			"GCCPUFraction": memstats.GCCPUFraction,
		},
	}
	return
}
