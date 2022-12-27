package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known term", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown term", func(t *testing.T) {
		_, got := dictionary.Search("nonexistent")

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new term", func(t *testing.T) {
		dictionary := Dictionary{}
		term := "test"
		definition := "this is just a test"

		err := dictionary.Add(term, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, term, definition)
	})

	t.Run("existing term", func(t *testing.T) {
		term := "test"
		definition := "this is just a test"
		dictionary := Dictionary{term: definition}

		err := dictionary.Add(term, "something")

		assertError(t, err, ErrTermExists)
		assertDefinition(t, dictionary, term, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing term", func(t *testing.T) {
		term := "test"
		definition := "this is just a test"
		dictionary := Dictionary{term: definition}
		newDefinition := "new definition"

		dictionary.Update(term, newDefinition)

		assertDefinition(t, dictionary, term, newDefinition)
	})

	t.Run("new term", func(t *testing.T) {
		term := "test"
		dictionary := Dictionary{}
		newDefinition := "new definition"

		err := dictionary.Update(term, newDefinition)

		assertError(t, err, ErrTermDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing term", func(t *testing.T) {
		term := "test"
		definition := "this is just a test"
		dictionary := Dictionary{term: definition}

		dictionary.Delete(term)
		_, err := dictionary.Search(term)

		assertError(t, err, ErrNotFound)
	})

	t.Run("missing term", func(t *testing.T) {
		term := "test"
		dictionary := Dictionary{}

		err := dictionary.Delete(term)

		assertError(t, err, ErrTermDoesNotExist)
	})
}

func assertDefinition(t testing.TB, dictionary Dictionary, term, definition string) {
	t.Helper()

	got, err := dictionary.Search(term)
	if err != nil {
		t.Fatal("should find added term:", err)
	}

	if definition != got {
		t.Errorf("got %q, want %q", got, definition)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q, want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
