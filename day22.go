package aoc2020

import (
	"fmt"
	"strconv"
	"strings"
)

// Day22Part1 returns the score for a one game.
func Day22Part1(p1, p2 []uint) (uint, error) {
	for len(p1) > 0 && len(p2) > 0 {
		n1 := pop(&p1)
		n2 := pop(&p2)
		if n1 > n2 {
			push(&p1, n1)
			push(&p1, n2)
		} else {
			push(&p2, n2)
			push(&p2, n1)
		}
	}
	if len(p1) > 0 {
		return score(p1), nil
	}
	return score(p2), nil
}

// Day22Part2 returns the score for a card game, or the winner (1 or 2) for
// recursive games (`game > 1` ).
func Day22Part2(p1, p2 []uint, game uint) uint {
	recurse := func(draw uint, deckSize uint) bool {
		if draw < deckSize {
			return true
		}
		return false
	}

	// cannot use []uint as map key, using string rep instead
	rep := func(deck []uint) string {
		return fmt.Sprintf("%v", deck)
	}
	seen1, seen2 := make(map[string]bool), make(map[string]bool)

	var winner uint // 1 -> player 1, 2 -> player 2
	for len(p1) > 0 && len(p2) > 0 {
		// avoid infinite recursion
		if seen1[rep(p1)] || seen2[rep(p2)] {
			winner = 1
			break
		}
		seen1[rep(p1)], seen2[rep(p2)] = true, true

		c1, c2 := pop(&p1), pop(&p2)

		if recurse(c1, uint(len(p1)+1)) &&
			recurse(c2, uint(len(p2)+1)) {

			cp1 := make([]uint, c1)
			cp2 := make([]uint, c2)
			copy(cp1, p1)
			copy(cp2, p2)
			winner = Day22Part2(cp1, cp2, game+1)
		} else {
			// regular draw
			if c1 > c2 {
				winner = 1
			} else {
				winner = 2
			}
		}
		if winner == 1 {
			push(&p1, c1)
			push(&p1, c2)
		} else {
			push(&p2, c2)
			push(&p2, c1)
		}
	}
	if game > 1 {
		return winner
	}

	if len(p1) > 0 {
		return score(p1)
	}
	return score(p2)
}

// NewDay22 parses puzzle input into two sets of decks for player #1 and player
// #2.
func NewDay22(lines []string) ([]uint, []uint, error) {
	var p1, p2 []uint
	var player1 bool
	for i := range lines {
		if lines[i] == "" {
			continue
		}
		if strings.HasPrefix(lines[i], "Player 1") {
			player1 = true
			continue
		}
		if strings.HasPrefix(lines[i], "Player 2") {
			player1 = false
			continue
		}
		n, err := strconv.Atoi(lines[i])
		if err != nil {
			msg := "line %d: error parsing number %q"
			return nil, nil, fmt.Errorf(msg, i, lines[i])
		}
		if player1 {
			p1 = append(p1, uint(n))
		} else {
			p2 = append(p2, uint(n))
		}
	}
	return p1, p2, nil
}

func push(a *[]uint, n uint) {
	*a = append(*a, n)
}

func pop(a *[]uint) uint {
	m := (*a)[0]
	*a = (*a)[1:]
	return m
}

func score(deck []uint) uint {
	var n, times uint
	for i := len(deck) - 1; i >= 0; i-- {
		times++
		n += deck[i] * times
	}
	return n
}
