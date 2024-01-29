package main

import (
	"fmt"
	"github.com/thebigyovadiaz/design-pattern-go/creacionales/singleton/client_one"
	"github.com/thebigyovadiaz/design-pattern-go/creacionales/singleton/client_two"
	"github.com/thebigyovadiaz/design-pattern-go/creacionales/singleton/singleton"
	"sync"
)

/**
* Singleton design pattern with concurrent instance
 */
func main() {
	wg := sync.WaitGroup{}
	wg.Add(200)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			client_one.IncrementAge()
		}()

		go func() {
			defer wg.Done()
			client_two.IncrementAge()
		}()
	}

	wg.Wait()
	p := singleton.GetInstance()
	age := p.GetAge()
	fmt.Println(age)
}
