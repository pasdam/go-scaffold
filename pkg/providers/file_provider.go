package providers

import (
	"github.com/pasdam/go-scaffold/pkg/filters"
	"github.com/pasdam/go-scaffold/pkg/processors"
)

// FileProvider identifies a type used to process files
type FileProvider interface {

	// ProvideFiles provides the files that match filesFilter to the given processor
	ProvideFiles(filesFilter filters.Filter, processor processors.Processor) error
}
