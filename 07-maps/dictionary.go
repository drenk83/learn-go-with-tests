package maps

import "errors"

var (
	ErrNotFound         = errors.New("could not find the word you were looking for")
	ErrWordExists       = errors.New("cannot add word because it already exists")
	ErrWordDoesNotExist = errors.New("cannot perform operation on word because it does not exist")
)

type Dictionary map[string]string

func (d Dictionary) Search(elem string) (string, error) {
	if res, ok := d[elem]; ok {
		return res, nil
	}
	return "", ErrNotFound
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
		return nil
	case nil:
		return ErrWordExists
	default:
		return err
	}
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case nil:
		d[key] = value
		return nil
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)

	switch err {
	case nil:
		delete(d, key)
		return nil
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}
}
