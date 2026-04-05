package models

import "sync"

type AutoIncID struct {
	sync.Mutex
	current int
}

func (a *AutoIncID) Next() (id int) {
	// set lock to ensure atomic access
	a.Lock()
	defer a.Unlock()
	// return current ID and increment for next call
	id = a.current
	a.current++
	return
}

func (a *AutoIncID) Set(value int) {
	a.Lock()
	defer a.Unlock()
	a.current = value
}

func (a *AutoIncID) Reset() {
	a.Lock()
	defer a.Unlock()
	a.current = 0
}
