package client_one

import (
	"github.com/thebigyovadiaz/design-pattern-go/creacionales/1-Singleton/singleton"
)

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
