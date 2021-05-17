package graph

import "github.com/quan12yt/graphql-sqlboiler-example/internal/service/users"

type Resolver struct {
	sv users.UserService
}
