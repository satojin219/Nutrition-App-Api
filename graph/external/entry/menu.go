package entry

type Menu struct {
	Id             uint
	RecipeName     string
	ImgUrl         string
	Foodstuffs     []Foodstuff
	TotalNutrition Nutrition
	Recipes        []Recipe
	Tips           string
	Cost           int
	Time           int
}
