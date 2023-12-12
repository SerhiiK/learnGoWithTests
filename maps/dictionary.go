package maps

const (
	ErrNotFound         = DictionaryErr("could not find word")
	ErrWordExist        = DictionaryErr("word exist")
	ErrWordDoesNotExist = DictionaryErr("word doesn't exist")
)

type DictionaryErr string
type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExist
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
	if _, ok := d[word]; !ok {
		return ErrWordDoesNotExist
	}
	return nil
}

func (e DictionaryErr) Error() string {
	return string(e)
}
