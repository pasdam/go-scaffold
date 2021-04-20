package filters

import (
	"regexp"

	"github.com/pasdam/go-scaffold/pkg/core"
)

type patternFilter struct {
	patterns  []*regexp.Regexp
	inclusive bool
}

// NewPatternFilter returns a new instance of Filter configured with the
// specified regExp or an error if one of them is invalid.
// The filter accept a string if it is inclusive and the value matches one of
// the regexp, of if it is exclusive (inclusive=false) and the value doesn't
// match any of the patterns.
func NewPatternFilter(inclusive bool, patterns ...string) (core.Filter, error) {
	regExps := make([]*regexp.Regexp, len(patterns))

	for i := 0; i < len(patterns); i++ {
		regex, err := regexp.Compile(patterns[i])
		if err != nil {
			return nil, err
		}
		regExps[i] = regex
	}

	return &patternFilter{
		patterns:  regExps,
		inclusive: inclusive,
	}, nil
}

// NewPatternFilterFromInstance duplicates the filter, using the same pattern(s)
// and the specified inclusive flag
func NewPatternFilterFromInstance(f core.Filter, inclusive bool) core.Filter {
	return &patternFilter{
		patterns:  f.(*patternFilter).patterns,
		inclusive: inclusive,
	}
}

// Accept returns true if it is inclusive and the value matches one of the
// regexp, of if it is exclusive (inclusive=false) and the value doesn't match
// any of the patterns.
func (f *patternFilter) Accept(value string) bool {
	var valueWhenFound bool
	if f.inclusive {
		valueWhenFound = true

	} else {
		valueWhenFound = false
	}

	for _, pattern := range f.patterns {
		if pattern.MatchString(value) {
			// the filter is exclusive but the name matches
			return valueWhenFound
		}
	}
	return !valueWhenFound
}
