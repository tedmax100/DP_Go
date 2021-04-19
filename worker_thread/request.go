package main

type Task struct {
	Handler func(channal int, v ...interface{})
	Params  []interface{}
}
