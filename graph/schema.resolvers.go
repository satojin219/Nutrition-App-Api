package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"Nutrition-App-Api/graph/generated"
	"Nutrition-App-Api/graph/model"
	"context"
	"fmt"
)

// CreateMenu is the resolver for the createMenu field.
func (r *mutationResolver) CreateMenu(ctx context.Context, input model.CrateMenuInput, foodstuffs []*model.CreateFoodstuffInput, recipes []*model.CreateRecipeInput, totalNutrition *model.CreateNutritionInput) (*model.Menu, error) {
	panic(fmt.Errorf("not implemented: CreateMenu - createMenu"))
}

// UpdateMenu is the resolver for the UpdateMenu field.
func (r *mutationResolver) UpdateMenu(ctx context.Context, input model.UpdateMenuInput, foodstuffs []*model.UpdateFoodstuffInput, recipes []*model.UpdateRecipeInput, totalNutrition model.UpdateNutritionInput) (*model.Menu, error) {
	panic(fmt.Errorf("not implemented: UpdateMenu - UpdateMenu"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
