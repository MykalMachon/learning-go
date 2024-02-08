package maps

import "testing"

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertErrors(t testing.TB, err error, want string) {
	t.Helper()

	if err == nil {
		t.Fatal("error was expected but none received")
	}

	if err.Error() != want {
		t.Errorf("got error %s, wanted error %s", err.Error(), want)
	}
}

func assertNoErrors(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Error("no error was expected but one was received")
	}
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("Find string in dictionary", func(t *testing.T) {
		query := "test"

		got, err := dictionary.Search(query)
		want := "this is just a test"

		assertStrings(t, got, want)
		assertNoErrors(t, err)
	})

	t.Run("Don't find string in dictionary", func(t *testing.T) {
		query := "not_found"

		got, err := dictionary.Search(query)
		want := ""

		assertErrors(t, err, ErrKeyNotFound.Error())
		assertStrings(t, got, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("Add a word to the dictionary", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "this is just a test")

		want := "this is just a test"
		got, err := dictionary.Search("test")

		assertNoErrors(t, err)
		assertStrings(t, got, want)
	})

	t.Run("Add a word to the dictionary", func(t *testing.T) {
		key := "test"
		value := "this is just a test"

		dictionary := Dictionary{key: value}
		err := dictionary.Add(key, value)

		got, searchErr := dictionary.Search(key)

		assertErrors(t, err, ErrKeyAlreadyExists.Error())
		assertNoErrors(t, searchErr)
		assertStrings(t, got, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update an existing word", func(t *testing.T) {
		dict := Dictionary{"test": "this is just a test"}
		dict.Update("test", "this is also just a test")

		got, err := dict.Search("test")

		assertNoErrors(t, err)
		assertStrings(t, got, "this is also just a test")
	})

	t.Run("Update a word that DNE", func(t *testing.T) {
		dict := Dictionary{"test": "this is just a test"}
		err := dict.Update("other_test", "this is also just a test")

		// make sure key wasn't added with update
		got, searchErr := dict.Search("other_test")

		assertErrors(t, err, ErrWordDoesNotExist.Error())
		assertErrors(t, searchErr, ErrKeyNotFound.Error())
		assertStrings(t, got, "")
	})
}

func TestDelete(t *testing.T) {

	t.Run("Delete a word", func(t *testing.T) {
		dict := Dictionary{"test": "test entry"}
		err := dict.Delete("test")

		got, searchErr := dict.Search("test")

		// assert delete doesn't throw an error
		assertNoErrors(t, err)

		// assert that the search throws an error
		assertErrors(t, searchErr, ErrKeyNotFound.Error())
		assertStrings(t, got, "")
	})

	t.Run("Delete a word that DNE", func(t *testing.T) {
		dict := Dictionary{"test": "test entry"}
		err := dict.Delete("other_test")

		assertErrors(t, err, ErrCannotDeleteNonExistantEntry.Error())
	})
}
