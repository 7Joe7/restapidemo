package resources

import "fmt"

type Pizza struct {
	Name        string
	Ingredients []string `json:",omitempty"`
	Id          string   `json:",omitempty"`
}

type Ingredient struct {
	Name string
	Id   string `json:",omitempty"`
}

type IngredientOnPizza struct {
	Id           string `json:",omitempty"`
	IngredientId string
	PizzaId      string
}

func NewPizza() *Pizza {
	return &Pizza{Name: "", Ingredients: []string{}}
}

func NewIngredient() *Ingredient {
	return &Ingredient{Name: ""}
}

func NewIngredientOnPizza() *IngredientOnPizza {
	return &IngredientOnPizza{}
}

func (p *Pizza) IsValid() error {
	if p.Name == "" {
		return fmt.Errorf("Missing pizza name.")
	}
	if len(p.Ingredients) == 0 {
		return fmt.Errorf("Missing ingredients for pizza.")
	}
	for i := 0; i < len(p.Ingredients); i++ {
		if p.Ingredients[i] == "" {
			return fmt.Errorf("Missing id of assigned ingredient.")
		}
	}
	return nil
}

func (in *Ingredient) IsValid() error {
	if in.Name == "" {
		return fmt.Errorf("Missing ingredient name.")
	}
	return nil
}

func (iop *IngredientOnPizza) IsValid() error {
	if iop.IngredientId == "" {
		return fmt.Errorf("Missing ingredient reference.")
	}
	if iop.PizzaId == "" {
		return fmt.Errorf("Missing pizza reference.")
	}
	return nil
}

func (p *Pizza) ToMap() map[string]string {
	return map[string]string{"Name": p.Name, "Id": p.Id}
}

func (i *Ingredient) ToMap() map[string]string {
	return map[string]string{"Name": i.Name, "Id": i.Id}
}

func (iop *IngredientOnPizza) ToMap() map[string]string {
	return map[string]string{"PizzaId": iop.PizzaId, "Id": iop.Id, "IngredientId": iop.IngredientId}
}
