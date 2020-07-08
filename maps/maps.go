package maps

// Dictionary type
type Dictionary map[string]string

// ErrNotFound - for not found
const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

// DictionaryErr for errors
type DictionaryErr string

// Error function to return string version
func (e DictionaryErr) Error() string {
	return string(e)
}

// Search - Returns a value from the map
func (d Dictionary) Search(find string) (string, error) {
	definition, ok := d[find]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add to add value to dict
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
