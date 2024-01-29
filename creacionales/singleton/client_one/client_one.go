package client_one

import (
	"github.com/thebigyovadiaz/design-pattern-go/creacionales/singleton/singleton"
)

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
