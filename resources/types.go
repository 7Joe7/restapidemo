package resources

import "fmt"

type pizza struct {
	Name string
	Ingredients []string
}

func NewPizza() *pizza {
	return &pizza{Name:"", Ingredients:[]string{}}
}

func (p *pizza) IsValid() error {
	if p.Name == "" {
		return fmt.Errorf("Missing pizza name.")
	}
	if len(p.Ingredients) == 0 {
		return fmt.Errorf("Missing ingredients for pizza.")
	}
	for i := 0; i < len(p.Ingredients); i++ {
		if p.Ingredients[i] == "" {

		}
	}
	return nil
}