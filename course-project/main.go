package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Job struct {
	Name   string
	Delay  time.Duration
	Number int
}

type Worker struct {
	Id         int
	JobQueue   chan Job
	WorkerPool chan chan Job
	QuitChan   chan bool
}

type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorkers int
	JobQueue   chan Job
}

func NewWorker(id int, workerPool chan chan Job) *Worker {
	return &Worker{
		Id:         id,
		JobQueue:   make(chan Job),
		WorkerPool: workerPool,
		QuitChan:   make(chan bool)}
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobQueue
			select {
			case job := <-w.JobQueue:
				fmt.Printf("worker %d started job %s\n", w.Id, job.Name)
				fib := Fibonacci(job.Number)
				time.Sleep(job.Delay)
				fmt.Printf("worker %d finished job %s with result %d\n", w.Id, job.Name, fib)
			case <-w.QuitChan:
				fmt.Printf("worker %d stopped\n", w.Id)
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

func Fibonacci(n int) int {
	if n <= 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {
	workerPool := make(chan chan Job, maxWorkers)
	return &Dispatcher{WorkerPool: workerPool, MaxWorkers: maxWorkers, JobQueue: jobQueue}
}

func (d *Dispatcher) Dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			fmt.Printf("dispatcher started job %s\n", job.Name)
			go func(job Job) {
				workerJobChannel := <-d.WorkerPool
				workerJobChannel <- job
			}(job)
		}
	}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(i, d.WorkerPool)
		worker.Start()
	}
	go d.Dispatch()
}

func RequestHandler(w http.ResponseWriter, r *http.Request, jobQueue chan Job) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "invalid value for delay: "+err.Error(), http.StatusBadRequest)
		return
	}

	value, err := strconv.Atoi(r.FormValue("number"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "invalid value for number: "+err.Error(), http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "missing name", http.StatusBadRequest)
		return
	}

	job := Job{Name: name, Delay: delay, Number: value}
	jobQueue <- job
	w.WriteHeader(http.StatusCreated)
}

func main() {
	const (
		maxWorkers   = 10
		maxQueueSize = 100
		port         = ":8081"
	)
	jobQueue := make(chan Job, maxQueueSize)
	dispatcher := NewDispatcher(jobQueue, maxWorkers)
	dispatcher.Run()
	http.HandleFunc("/fibo", func(w http.ResponseWriter, r *http.Request) {
		RequestHandler(w, r, jobQueue)
	})
	log.Fatal(http.ListenAndServe(port, nil))
}
