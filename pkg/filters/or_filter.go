package filters

type orFilter struct {
	filters []Filter
}

// Or returns a new Filter that merges the input ones, it performs a logical OR
// (if there is one that matches the value it will return true)
func Or(filters ...Filter) Filter {
	return &orFilter{
		filters: filters,
	}
}

func (f *orFilter) Accept(value string) bool {
	for _, filter := range f.filters {
		if filter.Accept(value) {
			return true
		}
	}
	return false
}
