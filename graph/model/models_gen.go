// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CrateMenuInput struct {
	RecipeName *string `json:"recipeName"`
	ImgURL     *string `json:"imgUrl"`
	Tips       *string `json:"tips"`
	Cost       *int    `json:"cost"`
	Time       *int    `json:"time"`
}

type CreateFoodstuffInput struct {
	Name      *string               `json:"name"`
	Weight    *int                  `json:"weight"`
	Nutrition *CreateNutritionInput `json:"nutrition"`
}

type CreateNutritionInput struct {
	Calorie         *int `json:"calorie"`
	Carbohydrates   *int `json:"carbohydrates"`
	Protein         *int `json:"protein"`
	Lipids          *int `json:"lipids"`
	Suger           *int `json:"suger"`
	DietaryFiber    *int `json:"dietaryFiber"`
	Salt            *int `json:"salt"`
	Na              *int `json:"na"`
	K               *int `json:"k"`
	Ca              *int `json:"ca"`
	Mg              *int `json:"mg"`
	P               *int `json:"p"`
	Fe              *int `json:"fe"`
	Zn              *int `json:"zn"`
	Cu              *int `json:"cu"`
	Mn              *int `json:"mn"`
	I               *int `json:"i"`
	Se              *int `json:"se"`
	Cr              *int `json:"cr"`
	Mo              *int `json:"mo"`
	VitA            *int `json:"vitA"`
	VitD            *int `json:"vitD"`
	VitE            *int `json:"vitE"`
	VitK            *int `json:"vitK"`
	VitB1           *int `json:"vitB1"`
	VitB2           *int `json:"vitB2"`
	VitB6           *int `json:"vitB6"`
	VitB12          *int `json:"vitB12"`
	VitC            *int `json:"vitC"`
	Niacin          *int `json:"niacin"`
	PantothenicAcid *int `json:"pantothenicAcid"`
	Folate          *int `json:"folate"`
	Biotin          *int `json:"biotin"`
}

type CreateRecipeInput struct {
	Content *string `json:"content"`
}

type DishData struct {
	Breakfast []*Menu `json:"breakfast"`
	Lunch     []*Menu `json:"lunch"`
	Dinner    []*Menu `json:"dinner"`
	Snack     []*Menu `json:"snack"`
}

type Foodstuff struct {
	ID        string     `json:"id"`
	Name      *string    `json:"name"`
	Weight    *int       `json:"weight"`
	Nutrition *Nutrition `json:"nutrition"`
}

type Menu struct {
	ID             string       `json:"id"`
	RecipeName     *string      `json:"recipeName"`
	ImgURL         *string      `json:"imgUrl"`
	Foodstuffs     []*Foodstuff `json:"foodstuffs"`
	TotalNutrition *Nutrition   `json:"totalNutrition"`
	Recipes        []*Recipe    `json:"recipes"`
	Tips           *string      `json:"tips"`
	Cost           *int         `json:"cost"`
	Time           *int         `json:"time"`
}

type Nutrition struct {
	Calorie         *int `json:"calorie"`
	Carbohydrates   *int `json:"carbohydrates"`
	Protein         *int `json:"protein"`
	Lipids          *int `json:"lipids"`
	Suger           *int `json:"suger"`
	DietaryFiber    *int `json:"dietaryFiber"`
	Salt            *int `json:"salt"`
	Na              *int `json:"na"`
	K               *int `json:"k"`
	Ca              *int `json:"ca"`
	Mg              *int `json:"mg"`
	P               *int `json:"p"`
	Fe              *int `json:"fe"`
	Zn              *int `json:"zn"`
	Cu              *int `json:"cu"`
	Mn              *int `json:"mn"`
	I               *int `json:"i"`
	Se              *int `json:"se"`
	Cr              *int `json:"cr"`
	Mo              *int `json:"mo"`
	VitA            *int `json:"vitA"`
	VitD            *int `json:"vitD"`
	VitE            *int `json:"vitE"`
	VitK            *int `json:"vitK"`
	VitB1           *int `json:"vitB1"`
	VitB2           *int `json:"vitB2"`
	VitB6           *int `json:"vitB6"`
	VitB12          *int `json:"vitB12"`
	VitC            *int `json:"vitC"`
	Niacin          *int `json:"niacin"`
	PantothenicAcid *int `json:"pantothenicAcid"`
	Folate          *int `json:"folate"`
	Biotin          *int `json:"biotin"`
}

type Recipe struct {
	ID      string  `json:"id"`
	Content *string `json:"content"`
}

type UpdateFoodstuffInput struct {
	ID        string                `json:"id"`
	Name      *string               `json:"name"`
	Weight    *int                  `json:"weight"`
	Nutrition *UpdateNutritionInput `json:"nutrition"`
}

type UpdateMenuInput struct {
	ID         string  `json:"id"`
	RecipeName *string `json:"recipeName"`
	ImgURL     *string `json:"imgUrl"`
	Tips       *string `json:"tips"`
	Cost       *int    `json:"cost"`
	Time       *int    `json:"time"`
}

type UpdateNutritionInput struct {
	Calorie         *int `json:"calorie"`
	Carbohydrates   *int `json:"carbohydrates"`
	Protein         *int `json:"protein"`
	Lipids          *int `json:"lipids"`
	Suger           *int `json:"suger"`
	DietaryFiber    *int `json:"dietaryFiber"`
	Salt            *int `json:"salt"`
	Na              *int `json:"na"`
	K               *int `json:"k"`
	Ca              *int `json:"ca"`
	Mg              *int `json:"mg"`
	P               *int `json:"p"`
	Fe              *int `json:"fe"`
	Zn              *int `json:"zn"`
	Cu              *int `json:"cu"`
	Mn              *int `json:"mn"`
	I               *int `json:"i"`
	Se              *int `json:"se"`
	Cr              *int `json:"cr"`
	Mo              *int `json:"mo"`
	VitA            *int `json:"vitA"`
	VitD            *int `json:"vitD"`
	VitE            *int `json:"vitE"`
	VitK            *int `json:"vitK"`
	VitB1           *int `json:"vitB1"`
	VitB2           *int `json:"vitB2"`
	VitB6           *int `json:"vitB6"`
	VitB12          *int `json:"vitB12"`
	VitC            *int `json:"vitC"`
	Niacin          *int `json:"niacin"`
	PantothenicAcid *int `json:"pantothenicAcid"`
	Folate          *int `json:"folate"`
	Biotin          *int `json:"biotin"`
}

type UpdateRecipeInput struct {
	ID      string  `json:"id"`
	Content *string `json:"content"`
}
