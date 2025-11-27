package aoc2020

import (
	"sort"
	"strings"
)

// NewDay21 parses text lines into a Day21 struct.
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
		f.allergens = make(map[allergen]struct{})
		f.ingredients = make(map[ingredient]struct{})

		line = prepare(line)

		fs := strings.Fields(line)
		parseIngredients := true
		for j := range fs {
			if fs[j] == separator {
				parseIngredients = false
			} else if parseIngredients {
				f.ingredients[ingredient(fs[j])] = struct{}{}
			} else {
				f.allergens[allergen(fs[j])] = struct{}{}
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
	allergens   map[allergen]struct{}
	ingredients map[ingredient]struct{}
}

// Day21 represents a food list.
type Day21 []food

// Delete remove this allergen/ ingredient combination from all foods, not just
// from food that have this allergen.
func (a *Day21) Delete(al allergen, in ingredient) {
	for i := range *a {
		delete((*a)[i].allergens, al)
		delete((*a)[i].ingredients, in)
	}
}

// Part1 solves Day 21, part #1.
func (a *Day21) Part1() uint {
backtrack:
	// use a combination of reductions until stable
	_, _, reduced := a.reduce1()
	if reduced {
		goto backtrack
	}
	_, _, reduced = a.reduceN()
	if reduced {
		goto backtrack
	}

	// count all remaining ingredients
	var n uint
	for i := range *a {
		n += uint(len((*a)[i].ingredients))
	}
	return n
}

// Part2 returns a comma separated list of ingredients, sorted by their
// corresponding allergen.
func (a *Day21) Part2() string {
	m := make(map[allergen]ingredient)
backtrack:
	// use a combination of reductions until stable
	al, in, reduced := a.reduce1()
	if reduced {
		m[al] = in
		goto backtrack
	}
	al, in, reduced = a.reduceN()
	if reduced {
		m[al] = in
		goto backtrack
	}

	// sort by allergen
	var ss []string
	for k := range m {
		ss = append(ss, string(k))
	}
	sort.Strings(ss)

	var result []string
	for i := range ss {
		result = append(result, string(m[allergen(ss[i])]))
	}
	return strings.Join(result, ",")
}

// reduce1 removes an allergen/ ingredient combination from a food if it is the
// only combination.
// returns true if reduced, false for no change.
func (a *Day21) reduce1() (allergen, ingredient, bool) {
	for i := range *a {
		if len((*a)[i].allergens) == 1 &&
			len((*a)[i].ingredients) == 1 {

			al := anyAllergen((*a)[i].allergens)
			in := anyIngredient((*a)[i].ingredients)
			a.Delete(al, in)
			return al, in, true
		}
	}
	return "", "", false
}

// reduceN searches for matching allergen/ ingredient combinations in all
// foods.
// If two foods f₁ and f₂ have an allergen aₙ and an ingredient iₙ, and the
// intersection of ingredients of f1 and f2 has exactly one element, it is
// removed.
// returns true if reduced, false for no change.
func (a *Day21) reduceN() (allergen, ingredient, bool) {
	as := a.allergens()
	sas := sortByOccurenceDesc(as)

	for i := range sas {
		// foods that contain the current allergen
		var fs []food
		// find all foods that contain allergen
		for j := range *a {
			if _, ok := (*a)[j].allergens[sas[i]]; ok {
				fs = append(fs, (*a)[j])
			}
		}
		if len(fs) == 0 {
			// no foods carry this allergen anymore
			continue
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
		return al, in, true
	}
	return "", "", false
}

// sortByOccurenceDesc converts a map into a sorted list.
func sortByOccurenceDesc(m map[allergen]uint) []allergen {
	// Collect keys and sort deterministically by count desc, then name asc.
	as := make([]allergen, 0, len(m))
	for k := range m {
		as = append(as, k)
	}
	sort.Slice(as, func(i, j int) bool {
		ci, cj := m[as[i]], m[as[j]]
		if ci != cj {
			return ci > cj
		}
		return string(as[i]) < string(as[j])
	})
	return as
}

// allergens returns a list of allergens, and their occurence in all food.
func (a Day21) allergens() map[allergen]uint {
	m := make(map[allergen]uint)
	for i := range a {
		for k := range a[i].allergens {
			m[k]++
		}
	}
	return m
}

func intersect(m1, m2 map[ingredient]struct{}) map[ingredient]struct{} {
	m := make(map[ingredient]struct{})
	for k := range m1 {
		if _, ok := m2[k]; ok {
			m[k] = struct{}{}
		}
	}
	return m
}

func anyAllergen(m map[allergen]struct{}) allergen {
	var mu allergen
	for k := range m {
		mu = k
		break
	}
	return mu
}

func anyIngredient(m map[ingredient]struct{}) ingredient {
	var mu ingredient
	for k := range m {
		mu = k
		break
	}
	return mu
}
