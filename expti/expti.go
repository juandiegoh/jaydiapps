package expti

import (
	"expvar"
	"log"
	"sync"
	"time"

	"github.com/paulbellamy/ratecounter"
)

const (
	diffTime  = -2 * time.Minute
	sleepTime = 10 * time.Second
)

var (
	counters  = make(map[string]*ratecounter.RateCounter)
	requests  = make(map[string]*expvar.Int)
	lastTimes = make(map[string]time.Time)
	mut       = &sync.RWMutex{}
)

func init() {
	go fixLastTimes()
}

func fixLastTimes() {
	for {
		for n, t := range lastTimes {
			del := expvarTimeout(t)
			if del {
				resetExpvar(n, t)
			}
		}
		time.Sleep(sleepTime)
	}
}

func expvarTimeout(t time.Time) bool {
	mut.RLock()
	del := t.Before(time.Now().Add(diffTime))
	mut.RUnlock()
	return del
}

func resetExpvar(n string, lastTime time.Time) {
	mut.Lock()
	del := lastTime.Before(time.Now().Add(diffTime))
	if del {
		log.Printf("Deleting %s from lastTimes", n)
		delete(lastTimes, n)
		request := requests[n]
		request.Set(0)
	}
	mut.Unlock()
}

// AddCount metrics
func AddCount(name string) {
	updLastTime(name, time.Now())

	var counter *ratecounter.RateCounter
	counter, ok := counters[name]
	if !ok {
		counter = ratecounter.NewRateCounter(1 * time.Minute)
		counters[name] = counter
	}

	var request *expvar.Int
	request, ok = requests[name]
	if !ok {
		request = expvar.NewInt(name)
		requests[name] = request
	}
	counter.Incr(1)
	request.Set(counter.Rate())
}

func updLastTime(name string, t time.Time) {
	mut.Lock()
	lastTimes[name] = t
	mut.Unlock()
}
