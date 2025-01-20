package main

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	o := NewOrdinal()
	m := NewSafeMap()
	o.Init(1123)
	fmt.Println("initial ordinal is:", o.GetOrdinal())
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Set(fmt.Sprint(i), "success")
			o.Increment()

		}(i)
	}

	wg.Wait()
	for i := 0; i < 10; i++ {
		v, err := m.Get(fmt.Sprint(i))
		if err != nil || v != "success" {
			panic(err)
		}
	}
	fmt.Println("final ordinal is:", o.GetOrdinal())
	fmt.Println("all keys found and marked as: 'success'")
}

//----------------------------------------------------------

// SafeMap uses a mutex to allow
// getting and setting in a thread-safe way
type SafeMap struct {
	m  map[string]string
	mu *sync.RWMutex
}

// NewSafeMap creates a SafeMap
func NewSafeMap() SafeMap {
	return SafeMap{m: make(map[string]string), mu: &sync.RWMutex{}}
}

// Set uses a write lock and sets the value given
// a key
func (t *SafeMap) Set(key, value string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.m[key] = value
}

// Get uses a RW lock and gets the value if it exists,
// otherwise an error is returned
func (t *SafeMap) Get(key string) (string, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	if v, ok := t.m[key]; ok {
		return v, nil
	}

	return "", errors.New("key not found")
}

// -----------------------------------------------------
// Ordinal holds a global a value
// and can only be initialized once
type Ordinal struct {
	ordinal uint64
	once    *sync.Once
}

// NewOrdinal returns ordinal with once
// setup
func NewOrdinal() *Ordinal {
	return &Ordinal{once: &sync.Once{}}
}

// Init sets the ordinal value
// can only be done once
func (o *Ordinal) Init(val uint64) {
	o.once.Do(func() {
		atomic.StoreUint64(&o.ordinal, val)
	})
}

// GetOrdinal will return the current
// ordinal
func (o *Ordinal) GetOrdinal() uint64 {
	return atomic.LoadUint64(&o.ordinal)
}

// Increment will increment the current
// ordinal
func (o *Ordinal) Increment() {
	atomic.AddUint64(&o.ordinal, 1)
}

// initial ordinal is: 1123
// final ordinal is: 1133
// all keys found and marked as: 'success'
