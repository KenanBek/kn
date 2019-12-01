package se

// Storage
type Storage interface{}

// Events
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

func NewSE() *SE {
	return &SE{}
}
