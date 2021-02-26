package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"repath.io/graph/generated"
	"repath.io/graph/model"
)

func (r *mutationResolver) CreatePerson(ctx context.Context, input model.NewPerson) (*model.Person, error) {
	person := &model.Person{ID: input.ID, FirstName: input.FirstName, LastName: input.LastName}
	people = append(people, person)
	return person, nil
}

func (r *queryResolver) People(ctx context.Context) ([]*model.Person, error) {
	return people, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var people = make([]*model.Person, 0)

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Person, error) {
	panic(fmt.Errorf("not implemented"))
}
