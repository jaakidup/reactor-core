package store

// CreateStore ...
// @storeType string, for createing a store type of SQLITE etc.
func CreateStore() *Store {

	userStore := MakeUserStore()

	return &Store{
		User: userStore,
	}
}

// Store ...
type Store struct {
	User *UserStore
}
