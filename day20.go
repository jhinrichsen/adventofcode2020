package aoc2020

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// NewDay20 parses text lines into a Day 20 struct.
func NewDay20(lines []string) (Day20, error) {
	var d Day20
	d.grid = make(map[uint][]string)

	isTileLine := func(s string) bool {
		return strings.HasPrefix(s, "Tile")
	}

	// make sure lines is an empty line
	if lines[len(lines)-1] > "" {
		lines = append(lines, "")
	}

	var ID uint
	var grid []string
	for i, line := range lines {
		if isTileLine(line) {
			s := line[5 : len(line)-1] // ignore trailing colon
			n, err := strconv.Atoi(s)
			if err != nil {
				msg := "line %d: error parsing tile ID in %q: %w"
				return d, fmt.Errorf(msg, i, line, err)
			}
			ID = uint(n)
		} else if line == "" {
			d.grid[ID] = grid
			grid = nil
		} else {
			// append regular grid line
			grid = append(grid, line)
		}
	}
	return d, nil
}

// Day20 represents a grid.
type Day20 struct {
	grid map[uint][]string
}

// CornerProduct returns the product of all corner tile IDs.
func (a Day20) CornerProduct() (uint, error) {
	bis := a.borders()

	asSet := func(ns []uint) map[uint]struct{} {
		m := make(map[uint]struct{}, len(ns))
		for i := range ns {
			m[ns[i]] = struct{}{}
		}
		return m
	}

	corresponding := func(i int) int {
		// borders come in pairs, return index of partner
		// 0 -- north
		// 1 -- north, flipped
		// 2 -- east
		// 3 -- east, flipped
		// ...
		if i%2 == 0 {
			return i + 1
		}
		return i - 1
	}

	var centers = make(map[uint]struct{})
	var middles = make(map[uint]struct{})
	var corners = make(map[uint]struct{})

	// search for unique borders
	for i := range bis {
		my := asSet(bis[i].borders[:])
		for j := range bis {
			if i == j { // do not remove my borders
				continue
			}
			for k := range bis[j].borders {
				if _, ok := my[bis[j].borders[k]]; ok {
					delete(my, bis[j].borders[k])

					// need to delete flipped border as well
					delete(my, bis[j].borders[corresponding(k)])
				}
			}
		}
		uniqueBorders := len(my) >> 1 // ignore flipped duplicates
		switch uniqueBorders {
		case 0:
			centers[bis[i].tileID] = struct{}{}
		case 1:
			middles[bis[i].tileID] = struct{}{}
		case 2:
			corners[bis[i].tileID] = struct{}{}
		}
	}

	product := uint(1)
	for k := range corners {
		product *= k
	}
	return product, nil
}

func northBorder(lines []string) string {
	return lines[0]
}

func southBorder(lines []string) string {
	return lines[len(lines)-1]
}

func verticalBorder(lines []string, l int) string {
	var sb strings.Builder
	for i := range lines {
		sb.WriteByte(lines[i][l])
	}
	return sb.String()
}

func eastBorder(lines []string) string {
	return verticalBorder(lines, len(lines[0])-1)
}

func westBorder(lines []string) string {
	return verticalBorder(lines, 0)
}

// BorderInfo keeps a numerical representation of a tile's borders.
type BorderInfo struct {
	tileID  uint
	borders [8]uint // north, east, south, west, both regular and flipped
}

// Borders sets up border information for each tile.
func (a *Day20) borders() []BorderInfo {
	var bis []BorderInfo
	for k, v := range a.grid {
		bi := BorderInfo{
			k,
			[8]uint{
				BorderID(northBorder(v)),
				BorderID(Reverse(northBorder(v))),
				BorderID(eastBorder(v)),
				BorderID(Reverse(eastBorder(v))),
				BorderID(southBorder(v)),
				BorderID(Reverse(southBorder(v))),
				BorderID(westBorder(v)),
				BorderID(Reverse(westBorder(v))),
			},
		}
		bis = append(bis, bi)
	}
	return bis
}

// BorderID generates a unique ID for a textual border.
func BorderID(s string) uint {
	var n uint
	for i := 0; i < len(s); i++ {
		n = n << 1
		if s[i] == '#' {
			n++
		}
	}
	return n
}

// Reverse implements the missing strings.Reverse (string -> string)
func Reverse(s string) string {
	var sb strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}
	return sb.String()
}

// tileGrid represents a concrete orientation of a tile (no borders removed yet).
type tileGrid struct {
	id   uint
	data []string
}

// rotate90 rotates a grid 90 degrees clockwise.
func rotate90(g []string) []string {
	h := len(g)
	w := len(g[0])
	out := make([]string, w)
	for x := 0; x < w; x++ {
		var sb strings.Builder
		for y := h - 1; y >= 0; y-- {
			sb.WriteByte(g[y][x])
		}
		out[x] = sb.String()
	}
	return out
}

// flipH flips a grid horizontally.
func flipH(g []string) []string {
	out := make([]string, len(g))
	for i := range g {
		out[i] = Reverse(g[i])
	}
	return out
}

// orientations returns the 8 orientations (4 rotations x optional horizontal flip).
func orientations(g []string) [][]string {
	res := make([][]string, 0, 8)
	r := g
	for i := 0; i < 4; i++ {
		res = append(res, r)
		res = append(res, flipH(r))
		r = rotate90(r)
	}
	return res
}

func topBorder(g []string) string    { return northBorder(g) }
func leftBorder(g []string) string   { return westBorder(g) }
func rightBorder(g []string) string  { return eastBorder(g) }
func bottomBorder(g []string) string { return southBorder(g) }

// assemble places all tiles on a puzzle grid using backtracking so that all borders match.
func (a Day20) assemble() ([][]tileGrid, error) {
	// Prepare all orientation variants for each tile ID
	type variants struct{ grids [][]string }
	all := make(map[uint]variants)
	ids := make([]uint, 0, len(a.grid))
	for id, g := range a.grid {
		ids = append(ids, id)
		all[id] = variants{grids: orientations(g)}
	}

	// grid size
	n := len(ids)
	size := int(math.Sqrt(float64(n)))
	if size*size != n {
		return nil, fmt.Errorf("tile count %d not a perfect square", n)
	}

	placed := make([][]tileGrid, size)
	for i := range placed {
		placed[i] = make([]tileGrid, size)
	}
	used := make(map[uint]struct{})

	var dfs func(pos int) bool
	dfs = func(pos int) bool {
		if pos == size*size {
			return true
		}
		r := pos / size
		c := pos % size
		for _, id := range ids {
			if _, ok := used[id]; ok {
				continue
			}
			for _, g := range all[id].grids {
				// Check top neighbor
				if r > 0 {
					up := placed[r-1][c].data
					if topBorder(g) != bottomBorder(up) {
						continue
					}
				}
				// Check left neighbor
				if c > 0 {
					left := placed[r][c-1].data
					if leftBorder(g) != rightBorder(left) {
						continue
					}
				}
				placed[r][c] = tileGrid{id: id, data: g}
				used[id] = struct{}{}
				if dfs(pos + 1) {
					return true
				}
				delete(used, id)
			}
		}
		return false
	}

	if !dfs(0) {
		return nil, fmt.Errorf("failed to assemble image")
	}
	return placed, nil
}

// removeBorders trims the outermost border from each tile and stitches to a single image.
func removeBorders(placed [][]tileGrid) []string {
	size := len(placed)
	tileH := len(placed[0][0].data)
	tileW := len(placed[0][0].data[0])
	innerH := tileH - 2
	innerW := tileW - 2

	out := make([]string, 0, size*innerH)
	for r := 0; r < size; r++ {
		for inr := 1; inr < tileH-1; inr++ { // skip top/bottom
			var sb strings.Builder
			for c := 0; c < size; c++ {
				line := placed[r][c].data[inr]
				sb.WriteString(line[1 : tileW-1])
			}
			_ = innerW // keep for readability
			out = append(out, sb.String())
		}
	}
	return out
}

// sea monster pattern as relative coordinates (dx, dy) where '#'
var seaMonster = []struct{ x, y int }{
	{18, 0},
	{0, 1}, {5, 1}, {6, 1}, {11, 1}, {12, 1}, {17, 1}, {18, 1}, {19, 1},
	{1, 2}, {4, 2}, {7, 2}, {10, 2}, {13, 2}, {16, 2},
}

func findSeaMonsters(img []string) (int, map[[2]int]bool) {
	H := len(img)
	W := len(img[0])
	marked := make(map[[2]int]bool)
	count := 0
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			ok := true
			for _, p := range seaMonster {
				yy := y + p.y
				xx := x + p.x
				if yy >= H || xx >= W || img[yy][xx] != '#' {
					ok = false
					break
				}
			}
			if ok {
				count++
				for _, p := range seaMonster {
					marked[[2]int{x + p.x, y + p.y}] = true
				}
			}
		}
	}
	return count, marked
}

// orientationsImg returns 8 orientations for the full image.
func orientationsImg(g []string) [][]string {
	res := make([][]string, 0, 8)
	r := g
	for i := 0; i < 4; i++ {
		res = append(res, r)
		res = append(res, flipH(r))
		r = rotate90(r)
	}
	return res
}

// WaterRoughness computes the number of '#' not part of any sea monster.
func (a Day20) WaterRoughness() (uint, error) {
	placed, err := a.assemble()
	if err != nil {
		return 0, err
	}
	img := removeBorders(placed)
	var bestMarked map[[2]int]bool
	maxFound := 0
	for _, g := range orientationsImg(img) {
		n, marked := findSeaMonsters(g)
		if n > maxFound {
			maxFound = n
			bestMarked = marked
		}
	}
	// Count total '#'
	var total uint
	for y := range img {
		for x := 0; x < len(img[y]); x++ {
			if img[y][x] == '#' {
				total++
			}
		}
	}
	// Subtract those covered by sea monsters
	if maxFound == 0 {
		return total, nil
	}
	return total - uint(len(bestMarked)), nil
}
