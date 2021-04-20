package filters

import (
	"github.com/pasdam/go-scaffold/pkg/core"
)

type noOpFilter struct{}

func NewNoOpFilter() core.Filter {
	return &noOpFilter{}
}

func (f *noOpFilter) Accept(value string) bool {
	return true
}
