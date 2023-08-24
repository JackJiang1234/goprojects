package main

import (
	"time"
)

type WorkerPool struct {
	start time.Time
	end   time.Time

	total int
}

type work func()

func (w *WorkerPool) Report() {

}

func (w *WorkerPool) Run(wk work, total int){

}
