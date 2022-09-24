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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) CreateFoodstuff(ctx context.Context, input *model.CreateFoodstuffInput) ([]*model.Foodstuff, error) {
	panic(fmt.Errorf("not implemented: CreateFoodstuff - CreateFoodstuff"))
}
func (r *mutationResolver) UpdateFoodstuff(ctx context.Context, input *model.UpdateFoodstuffInput) ([]*model.Foodstuff, error) {
	panic(fmt.Errorf("not implemented: UpdateFoodstuff - UpdateFoodstuff"))
}
func (r *mutationResolver) CreateRecipe(ctx context.Context, input *model.CreateRecipeInput) ([]*model.Recipe, error) {
	panic(fmt.Errorf("not implemented: CreateRecipe - CreateRecipe"))
}
func (r *mutationResolver) UpdateRecipe(ctx context.Context, input *model.UpdateRecipeInput) ([]*model.Recipe, error) {
	panic(fmt.Errorf("not implemented: UpdateRecipe - UpdateRecipe"))
}
func (r *mutationResolver) CreateNutrition(ctx context.Context, input *model.CreateNutritionInput) (*model.Nutrition, error) {
	panic(fmt.Errorf("not implemented: CreateNutrition - CreateNutrition"))
}
func (r *mutationResolver) UpdateNutrition(ctx context.Context, input *model.UpdateNutritionInput) (*model.Nutrition, error) {
	panic(fmt.Errorf("not implemented: UpdateNutrition - UpdateNutrition"))
}
