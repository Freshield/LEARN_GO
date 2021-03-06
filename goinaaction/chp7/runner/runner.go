/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: runner.py
@Time: 2021-09-22 16:49
@Last_update: 2021-09-22 16:49
@Desc: None
@==============================================@
@      _____             _   _     _   _       @
@     |   __|___ ___ ___| |_|_|___| |_| |      @
@     |   __|  _| -_|_ -|   | | -_| | . |      @
@     |__|  |_| |___|___|_|_|_|___|_|___|      @
@                                    Freshield @
@==============================================@
*/
package runner

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	interrupt chan os.Signal

	complete chan error

	timeout <-chan time.Time

	tasks []func(int)
}

var ErrTimeout = errors.New("received timeout")

var ErrInterrupt = errors.New("received interrupt")

func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete: make(chan error),
		timeout: time.After(d),
	}
}

func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <-r.complete:
		return err

	case <-r.timeout:
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		log.Println("begin", id)
		if r.gotInterrupt() {
			log.Println("got interrupt", id)
			return ErrInterrupt
		}
		log.Println("after check", id)

		task(id)
		log.Println("done task", id)
	}

	return nil
}

func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true

	default:
		return false
	}
}

