package client_two

import (
	"github.com/thebigyovadiaz/design-pattern-go/creacionales/1-Singleton/singleton"
)

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
