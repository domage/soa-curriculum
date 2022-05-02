package graph

import (
	"context"
	"time"

	"nuuka.com/go-gql/generated"
	"nuuka.com/go-gql/models"
)

func New() generated.Config {
	c := generated.Config{
		Resolvers: &Resolver{},
	}

	return c
}

type Resolver struct{}

func (r *mutationResolver) Reverse(ctx context.Context, str string) (*generated.Message, error) {
	now := time.Now().String()

	return &generated.Message{
		String: &str,
		MadeAt: &now,
	}, nil
}

func (r *queryResolver) Test(ctx context.Context) (*string, error) {
	res := "Test from ggo!"

	return &res, nil
}

func (r *queryResolver) User(ctx context.Context, id int) (*models.User, error) {
	return &models.User{
		ID:        id,
		FirstName: "FirstName",
		LastName:  "LastName",
	}, nil
}

func (r *userResolver) LinkedUsers(ctx context.Context, obj *models.User) ([]*models.User, error) {
	return nil, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
