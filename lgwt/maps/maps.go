package maps

// custom error type
type DictionaryErr string

func (err DictionaryErr) Error() string {
	return string(err)
}

var ErrKeyNotFound = DictionaryErr("key was not found in dictionary")
var ErrKeyAlreadyExists = DictionaryErr("key already exists in the dictionary")
var ErrWordDoesNotExist = DictionaryErr("cannot update word as it does not exist in the dictionary")
var ErrCannotDeleteNonExistantEntry = DictionaryErr("cannot delete this entry as it does not exist in the dictionary")

type Dictionary map[string]string

func (d Dictionary) Search(query string) (string, error) {
	definition, ok := d[query]

	if !ok {
		return "", ErrKeyNotFound
	}

	return definition, nil
}

// notice how we don't have to do d *Dictionary?
// this is because maps can be edited without a pointer. they seem like a ref type, they're not.
// another note, you should never initialize an empty map. always do map[string]string{} with {} being key
func (d Dictionary) Add(key string, value string) error {
	_, ok := d[key]
	if ok {
		return ErrKeyAlreadyExists
	}

	d[key] = value
	return nil
}

func (d Dictionary) Update(key string, newValue string) error {
	_, ok := d[key]
	if !ok {
		return ErrWordDoesNotExist
	}

	d[key] = newValue
	return nil
}

func (d Dictionary) Delete(key string) error {
	_, ok := d[key]
	// not really neccesary but 🤷‍♀️ oh well
	if !ok {
		return ErrCannotDeleteNonExistantEntry
	}

	delete(d, key)
	return nil
}
