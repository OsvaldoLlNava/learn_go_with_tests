package dictionary

const (
	ErrNotFound        = DictionaryErr("could not find the word you were looking for")
	ErrKeyExists       = DictionaryErr("cannot add word because it already exists")
	ErrKeyDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {

	value, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) error {

	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrKeyExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, value string) error {

	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrKeyDoesNotExist
	case nil:
		d[key] = value
	default:
		return err
	}

	return nil
}
