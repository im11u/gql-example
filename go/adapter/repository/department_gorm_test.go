package repository_test

import (
	"testing"

	"github.com/im11u/gql-example/go/adapter/repository"
)

func TestDepartmentGormRepository_FindAll(t *testing.T) {
	repo := repository.NewDepartmentGormRepository(db)

	result, err := repo.FindAll()
	if err != nil {
		t.Fatal(err)
	}

	wants := []struct {
		ID   uint
		Name string
	}{
		{ID: 1, Name: "営業部"},
		{ID: 2, Name: "システム開発部"},
	}

	if length := len(result); length != len(wants) {
		t.Errorf("[Length] Got: %d | Want: %d", length, len(wants))
	}
	for i, want := range wants {
		got := result[i]

		// ID
		if got.ID != want.ID {
			t.Errorf("[ID(%d)] Got: %d | Want: %d", i, got.ID, want.ID)
		}
		// Name
		if got.Name != want.Name {
			t.Errorf("[Name(%d)] Got: '%s' | Want: '%s'", i, got.Name, want.Name)
		}
	}
}
