package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"repath.io/data"
	"repath.io/graph/generated"
	"repath.io/graph/model"
)

func (r *mutationResolver) CreatePerson(ctx context.Context, input model.NewPerson) (*model.Person, error) {
	record := data.Person{FirstName: input.FirstName, LastName: input.LastName, Title: input.Title}
	result, err := data.Store().Create(record)
	if err != nil {
		return nil, err
	}
	response := &model.Person{ID: result.Id, FirstName: result.FirstName, LastName: result.LastName, Title: *result.Title}
	return response, nil
}

func (r *mutationResolver) UpdatePerson(ctx context.Context, input model.UpdatePerson) (*model.PersonUpdate, error) {
	record := data.Person{Id: input.ID, FirstName: input.FirstName, LastName: input.LastName, Title: input.Title}
	out, err := data.Store().Update(record)
	if err != nil {
		return nil, err
	}
	result := &model.PersonUpdate{
		Current: &model.Person{
			ID:        out.Current["id"].(string),
			FirstName: out.Current["FirstName"].(string),
			LastName:  out.Current["LastName"].(string),
			Title:     out.Current["Title"].(string)},
		Previous: &model.Person{
			ID:        out.Previous["id"].(string),
			FirstName: out.Previous["FirstName"].(string),
			LastName:  out.Previous["LastName"].(string),
			Title:     out.Previous["Title"].(string)}}
	return result, nil
}

func (r *mutationResolver) DeletePerson(ctx context.Context, input model.DeletePerson) (*model.PersonDelete, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) People(ctx context.Context) ([]*model.Person, error) {
	people, err := data.Store().FindAll()
	if err != nil {
		return nil, err
	}
	out := make([]*model.Person, 0)
	for _, v := range people {
		out = append(out, &model.Person{
			ID:        v.Id,
			FirstName: v.FirstName,
			LastName:  v.LastName,
			Title: *v.Title,
		})
	}
	return out, nil
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
