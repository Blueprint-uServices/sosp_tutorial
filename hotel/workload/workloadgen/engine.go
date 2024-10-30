package workloadgen

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"golang.org/x/exp/rand"
	rand2 "golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat/distuv"
)

type Stat struct {
	Start    int64
	Duration int64
	IsError  bool
}

type RequestFunction func(context.Context) Stat

type Engine struct {
	Stats    []Stat
	OutFile  string
	Tput     int64
	Duration time.Duration
	Wrk      *Workload
}

type ApiInfo struct {
	Name       string
	Fn         RequestFunction
	Proportion int
}

type Workload struct {
	ApiInfos []ApiInfo
}

func NewWorkload() *Workload {
	w := &Workload{}
	return w
}

func (w *Workload) AddAPI(name string, fn RequestFunction, proportion int) error {
	info := ApiInfo{Name: name, Fn: fn, Proportion: proportion}
	w.ApiInfos = append(w.ApiInfos, info)
	return nil
}

func NewEngine(outfile string, tput int64, duration string, w *Workload) (*Engine, error) {
	dur, err := time.ParseDuration(duration)
	if err != nil {
		return nil, err
	}
	return &Engine{OutFile: outfile, Tput: tput, Duration: dur, Wrk: w}, nil
}

func (e *Engine) init_prop_map() map[int]RequestFunction {
	sort.Slice(e.Wrk.ApiInfos, func(i, j int) bool { return e.Wrk.ApiInfos[i].Proportion > e.Wrk.ApiInfos[j].Proportion })
	proportion_map := make(map[int]RequestFunction)
	var last_proportion_val int
	for _, api := range e.Wrk.ApiInfos {
		var i int
		for i = 0; i < api.Proportion; i += 1 {
			proportion_map[last_proportion_val+i] = api.Fn
		}
		last_proportion_val += i
	}
	return proportion_map
}

func (e *Engine) RunOpenLoop(ctx context.Context) {
	prop_map := e.init_prop_map()
	log.Println("Target throughput", e.Tput)
	// Launch stat collector channel
	stat_channel := make(chan Stat, e.Tput)
	done := make(chan bool)
	go func() {
		count := 0
		for stat := range stat_channel {
			count += 1
			if count%1000 == 0 {
				log.Println("Processed", count, "requests")
			}
			e.Stats = append(e.Stats, stat)
		}
		close(done)
	}()

	// Launch the request maker goroutine that launches a request every tick_val
	tick_every := float64(1e9) / float64(e.Tput)
	tick_val := time.Duration(int64(1e9 / float64(e.Tput)))
	log.Println("Ticking after every", tick_val)
	stop := make(chan bool)
	var wg sync.WaitGroup
	var i int64
	go func() {
		src := rand2.NewSource(0)
		g := distuv.Poisson{100, src}
		timer := time.NewTimer(0 * time.Second)
		next := time.Now()
		for {
			select {
			case <-stop:
				return
			case <-timer.C:
				n := rand.Intn(100)
				fn := prop_map[n]
				wg.Add(1)
				go func() {
					defer wg.Done()
					stat := fn(ctx)
					stat_channel <- stat
				}()
				next = next.Add(time.Duration(g.Rand()*tick_every/100) * time.Nanosecond)
				waitt := next.Sub(time.Now())
				timer.Reset(waitt)
			}
		}
	}()
	// Let the requests happen for the desired duration
	time.Sleep(e.Duration)
	stop <- true
	// Wait for all the launched routines to finish
	wg.Wait()
	log.Println("Total launched routines:", i)
	// Finish gathering stats!
	close(stat_channel)
	<-done
	log.Println("Finished all requests")
}

func (e *Engine) PrintStats() error {
	var num_errors int64
	var num_reqs int64
	var sum_durations int64
	stat_strings := []string{}
	for _, stat := range e.Stats {
		num_reqs += 1
		if stat.IsError {
			num_errors += 1
		}
		sum_durations += stat.Duration
		stat_strings = append(stat_strings, fmt.Sprintf("%d,%d,%t", stat.Start, stat.Duration, stat.IsError))
	}

	fmt.Println("Total Number of Requests:", num_reqs)
	fmt.Println("Successful Requests:", num_reqs-num_errors)
	fmt.Println("Error Responses:", num_errors)
	fmt.Println("Average Latency:", float64(sum_durations)/float64(num_reqs))
	// Write to file
	header := "Start,Duration,IsError\n"
	data := header + strings.Join(stat_strings, "\n")
	f, err := os.OpenFile(e.OutFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}
