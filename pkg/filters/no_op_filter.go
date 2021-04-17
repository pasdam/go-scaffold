package filters

type noOpFilter struct{}

func NewNoOpFilter() Filter {
	return &noOpFilter{}
}

func (f *noOpFilter) Accept(value string) bool {
	return true
}
