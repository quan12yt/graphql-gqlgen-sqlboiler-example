package graph

import "github.com/quan12yt/graphql-sqlboiler-example/internal/users"

type Resolver struct {
	sv users.UserService
}
