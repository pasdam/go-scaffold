package core

// FileProvider identifies a type used to process files
type FileProvider interface {

	// ProvideFiles provides the files that match filesFilter to the given processor
	ProvideFiles(filesFilter Filter, processor Processor) error
}
