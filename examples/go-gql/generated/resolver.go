package generated

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"

	"nuuka.com/go-gql/models"
)

type Resolver struct{}

func (r *mutationResolver) Reverse(ctx context.Context, str string) (*Message, error) {
	panic("not implemented")
}

func (r *queryResolver) Test(ctx context.Context) (*string, error) {
	panic("not implemented")
}

func (r *queryResolver) User(ctx context.Context, id int) (*models.User, error) {
	panic("not implemented")
}

func (r *userResolver) LinkedUsers(ctx context.Context, obj *models.User) ([]*models.User, error) {
	panic("not implemented")
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
