package graph

import (
	"database/sql"

	"github.com/im11u/gql-example/go/adapter/action"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	actions action.Actions
}

func NewResolver(db *sql.DB) *Resolver {
	return &Resolver{
		actions: action.NewActions(db),
	}
}
