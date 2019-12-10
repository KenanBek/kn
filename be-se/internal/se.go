package se

// Storage is exported.
type Storage interface{}

// Events is exported.
type Events interface{}

// SE is a application composer.
type SE struct {
	// tech
	storage Storage
	events  Events

	// modules
	crawler interface{}
	indexer interface{}
}

// NewSE is exported.
func NewSE() *SE {
	return &SE{}
}
