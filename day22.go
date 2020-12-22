package aoc2020

import (
	"fmt"
	"strconv"
	"strings"
)

// Day22Part1 returns the score for a one game.
func Day22Part1(p1, p2 []uint) (uint, error) {
	push := func(a *[]uint, n uint) {
		*a = append(*a, n)
	}
	pop := func(a *[]uint) uint {
		m := (*a)[0]
		*a = (*a)[1:]
		return m
	}

	var round uint
	for len(p1) > 0 && len(p2) > 0 {
		round++
		// fmt.Printf("-- Round %d --\n", round)
		// fmt.Printf("Player 1's deck: %+v\n", p1)
		// fmt.Printf("Player 2's deck: %+v\n", p2)
		n1 := pop(&p1)
		// fmt.Printf("Player 1 plays: %d\n", n1)
		n2 := pop(&p2)
		// fmt.Printf("Player 2 plays: %d\n", n2)
		if n1 > n2 {
			// fmt.Printf("Player 1 wins the round!\n")
			push(&p1, n1)
			push(&p1, n2)
		} else {
			// fmt.Printf("Player 2 wins the round!\n")
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
	push := func(a *[]uint, n uint) {
		*a = append(*a, n)
	}
	pop := func(a *[]uint) uint {
		m := (*a)[0]
		*a = (*a)[1:]
		return m
	}
	recurse := func(draw uint, deckSize uint) bool {
		if draw < deckSize {
			return true
		}
		return false
	}

	// cannot use []uint as map key, using string rep instead
	rep := func(deck []uint) string {
		return fmt.Sprintf("%+v", deck)
	}
	previousDecks1 := make(map[string]bool)
	previousDecks2 := make(map[string]bool)

	// fmt.Printf("\n=== Game %d ===\n", game)
	var round uint
	var winner uint // 1 -> player 1, 2 -> player 2
	for len(p1) > 0 && len(p2) > 0 {
		round++
		// fmt.Printf("\n-- Round %d (Game %d) --\n", round, game)
		// fmt.Printf("Player 1's deck: %+v\n", p1)
		// fmt.Printf("Player 2's deck: %+v\n", p2)

		// avoid infinite recursion
		if previousDecks1[rep(p1)] || previousDecks2[rep(p2)] {
			winner = 1
			break
		}
		previousDecks1[rep(p1)] = true
		previousDecks2[rep(p2)] = true

		c1 := pop(&p1)
		// fmt.Printf("Player 1 plays: %d\n", c1)
		c2 := pop(&p2)
		// fmt.Printf("Player 2 plays: %d\n", c2)

		if recurse(c1, uint(len(p1)+1)) &&
			recurse(c2, uint(len(p2)+1)) {

			// fmt.Printf("Playing a sub-game to determine the " +
			// 	"winner...\n")
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
			// fmt.Printf("Player 1 wins round %d of game %d!\n",
			//		round, game)
			push(&p1, c1)
			push(&p1, c2)
		} else {
			// fmt.Printf("Player 2 wins round %d of game %d!\n",
			//	round, game)
			push(&p2, c2)
			push(&p2, c1)
		}
	}
	if game > 1 {
		// fmt.Printf("The winner of game %d is player %d!\n",
		//	game, winner)
		// fmt.Println("\n...anyway, back to game 1.")
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

func score(deck []uint) uint {
	var n, times uint
	for i := len(deck) - 1; i >= 0; i-- {
		times++
		n += deck[i] * times
	}
	return n
}
