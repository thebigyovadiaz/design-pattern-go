package client_two

import (
	"github.com/thebigyovadiaz/design-pattern-go/creacionales/singleton/singleton"
)

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
