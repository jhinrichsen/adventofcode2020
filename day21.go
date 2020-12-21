package aoc2020

import (
	"strings"
)

func NewDay21(lines []string) (Day21, error) {
	const separator = ":"
	var d Day21

	prepare := func(s string) string {
		s = strings.Replace(s, "(contains", separator, 1)
		s = strings.Replace(s, ")", "", 1)
		s = strings.ReplaceAll(s, ",", "")
		return s
	}
	// tried a regexp but gave up
	for i, line := range lines {
		var f food
		f.ID = i
		f.allergens = make(map[allergen]bool)
		f.ingredients = make(map[ingredient]bool)

		line = prepare(line)

		fs := strings.Fields(line)
		parseIngredients := true
		for j := range fs {
			if fs[j] == separator {
				parseIngredients = false
			} else if parseIngredients {
				f.ingredients[ingredient(fs[j])] = true
			} else {
				f.allergens[allergen(fs[j])] = true
			}
		}
		d = append(d, f)
	}
	return d, nil
}

// If Go has strong typing, why not use it to make sure there's no
// confusion between allergens and ingredients?

type allergen string
type ingredient string

type food struct {
	ID          int
	allergens   map[allergen]bool
	ingredients map[ingredient]bool
}

type Day21 []food

// Delete remove this allergen/ ingredient combination from all foods, not just
// from food that have this allergen.
func (a *Day21) Delete(al allergen, in ingredient) {
	for i := range *a {
		delete((*a)[i].allergens, al)
		delete((*a)[i].ingredients, in)
	}
}

func (a *Day21) Part1() uint {
backtrack:
	// use a combination of reductions until stable
	if a.reduce1() {
		goto backtrack
	}
	if a.reduceN() {
		goto backtrack
	}

	// count all remaining ingredients
	var n uint
	for i := range *a {
		n += uint(len((*a)[i].ingredients))
	}
	return n
}

// reduce1 removes an allergen/ ingredient combination from a food if it is the
// only combination.
// returns true if reduced, false for no change.
func (a *Day21) reduce1() bool {
	for i := range *a {
		if len((*a)[i].allergens) == 1 &&
			len((*a)[i].ingredients) == 1 {

			al := anyAllergen((*a)[i].allergens)
			in := anyIngredient((*a)[i].ingredients)
			a.Delete(al, in)
			return true
		}
	}
	return false
}

// reduceN searches for matching allergen/ ingredient combinations in all
// foods.
// If two foods f₁ and f₂ have an allergen aₙ and an ingredient iₙ, and the
// intersection of ingredients of f1 and f2 has exactly one element, it is
// removed.
// returns true if reduced, false for no change.
func (a *Day21) reduceN() bool {
	as := a.Allergens()
	sas := SortByOccurenceDesc(as)

	var fs []food // foods that contain a certain allergen
	for i := range sas {
		// find all foods that contain allergen
		for j := range *a {
			if _, ok := (*a)[j].allergens[sas[i]]; ok {
				fs = append(fs, (*a)[j])
			}
		}
		// see if we have exactly one ingredient in all foods, in this
		// case it is the right ingredient/allergen match
		// multi-intersection:
		in0 := fs[0].ingredients
		for j := 1; j < len(fs); j++ {
			in0 = intersect(in0, fs[j].ingredients)
		}
		if len(in0) != 1 {
			// found more than 1 allergen <-> ingredient
			// combination, does not help, try next
			continue
		}

		// found a unique combination of allergen/ ingredient
		al := sas[i]
		in := anyIngredient(in0) // there's only one
		a.Delete(al, in)

		// only one change per run, done for now
		return true
	}
	return false
}

func SortByOccurenceDesc(m map[allergen]uint) []allergen {
	var as []allergen
	for len(m) > 0 {
		var max uint
		var maxA allergen
		for k, v := range m {
			if v > max {
				max = v
				maxA = k
			}
		}
		// move from map to array
		as = append(as, maxA)
		delete(m, maxA)
	}
	return as
}

// Allergens returns a list of allergens, and their occurence in all food.
func (a Day21) Allergens() map[allergen]uint {
	m := make(map[allergen]uint)
	for i := range a {
		for k := range a[i].allergens {
			m[k]++
		}
	}
	return m
}

func intersect(m1, m2 map[ingredient]bool) map[ingredient]bool {
	m := make(map[ingredient]bool)
	for k := range m1 {
		if m2[k] {
			m[k] = true
		}
	}
	return m
}

func anyAllergen(m map[allergen]bool) allergen {
	var mu allergen
	for k := range m {
		mu = k
		break
	}
	return mu
}

func anyIngredient(m map[ingredient]bool) ingredient {
	var mu ingredient
	for k := range m {
		mu = k
		break
	}
	return mu
}
