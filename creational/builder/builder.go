/*
Builder Pattern Example

This file demonstrates the Builder Pattern, a creational design pattern that separates the construction of a complex object from its representation. The Builder Pattern allows the same construction process to create different representations of an object.

Prerequisite Components for the Builder Pattern:
1. Product:
   - The complex object to be created (here: Meal).
2. Builder (Interface):
   - Specifies the steps required to build the product (e.g., SetDrink, SetMainCourse, SetDessert, Build).
3. ConcreteBuilder:
   - Implements the Builder interface to provide specific implementations for each step (e.g., VegetarianMeal, KidsMeal).
4. Director:
   - Orchestrates the building process using a Builder. It defines the order in which to call the building steps to construct the product.

How this file implements the Builder Pattern:
- Product: The 'Meal' struct represents the complex object to be built.
- Builder Interface: 'MealBuilder' defines the methods for setting each part of the meal and building the final product.
- Concrete Builders: 'VegetarianMeal' and 'KidsMeal' structs implement the MealBuilder interface to build specific types of meals.
- Director: The 'Director' struct manages the construction process by using a MealBuilder to assemble a Meal with the desired components.

Usage:
- Create a builder (e.g., NewVegetarianMeal or NewKidsMeal).
- Set the builder in the Director.
- Call Director.Construct() with the desired meal components to get a fully built Meal.

This approach encapsulates the construction logic, allows for different meal variations, and improves code readability and maintainability.
*/

package builder

type Meal struct {
	Drink      string
	MainCourse string
	Dessert    string
}

type MealBuilder interface {
	SetDrink(drink string)
	SetMainCourse(mainCourse string)
	SetDessert(dessert string)
	Build() Meal
}

type VegetarianMeal struct {
	Meal
}

func NewVegetarianMeal() MealBuilder {
	return &VegetarianMeal{}
}

func (v *VegetarianMeal) SetDrink(drink string) {
	v.Drink = drink
}

func (v *VegetarianMeal) SetMainCourse(main string) {
	v.MainCourse = main
}

func (v *VegetarianMeal) SetDessert(dessert string) {
	v.Dessert = dessert
}

func (v *VegetarianMeal) Build() Meal {
	return v.Meal
}

type KidsMeal struct {
	Meal
}

func NewKidsMeal() MealBuilder {
	return &KidsMeal{}
}

func (k *KidsMeal) SetDrink(drink string) {
	k.Drink = drink
}

func (k *KidsMeal) SetMainCourse(main string) {
	k.MainCourse = main
}

func (k *KidsMeal) SetDessert(dessert string) {
	k.Dessert = dessert
}

func (k *KidsMeal) Build() Meal {
	return k.Meal
}

type Director struct {
	builder MealBuilder
}

func NewDirector(builder MealBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) SetBuilder(builder MealBuilder) {
	d.builder = builder
}

func (d *Director) Build() Meal {
	return d.builder.Build()
}

func (d *Director) Construct(drink, mainCourse, dessert string) Meal {
	d.builder.SetDrink(drink)
	d.builder.SetMainCourse(mainCourse)
	d.builder.SetDessert(dessert)
	return d.builder.Build()
}

/*
Usage Example:

	builder := NewVegetarianMeal()
	director := NewDirector(builder)
	meal := director.Construct("Juice", "Salad", "Fruit")

	fmt.Println(meal.Drink)      // Output: Juice
	fmt.Println(meal.MainCourse) // Output: Salad
	fmt.Println(meal.Dessert)    // Output: Fruit
*/
