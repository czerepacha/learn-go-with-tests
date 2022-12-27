package main

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound         = DictionaryErr("could not find the term you were looking for.")
	ErrTermExists       = DictionaryErr("cannot add a term because it already exists.")
	ErrTermDoesNotExist = DictionaryErr("cannot update/delete term because it does not exist.")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(term string) (string, error) {
	definition, ok := d[term]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(term, definition string) error {
	_, err := d.Search(term)

	switch err {
	case ErrNotFound:
		d[term] = definition
	case nil:
		return ErrTermExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(term, definition string) error {
	_, err := d.Search(term)

	switch err {
	case ErrNotFound:
		return ErrTermDoesNotExist
	case nil:
		d[term] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(term string) error {
	_, err := d.Search(term)

	switch err {
	case ErrNotFound:
		return ErrTermDoesNotExist
	case nil:
		delete(d, term)
	default:
		return err
	}

	return nil
}
