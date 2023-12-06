package engine

type Neighbor struct {
	Symbol rune
	Row    int
	Column int
}

func findNeighbors(part EnginePart, schema *Schema) []Neighbor {
	minRow := part.Row - 1
	maxRow := part.Row + 1
	minColumn := part.Column - 1
	maxColumn := part.Column + part.Length

	neighbors := []Neighbor{}
	for row := minRow; row <= maxRow; row++ {
		for column := minColumn; column <= maxColumn; {
			// skip digit indexes
			rowMatch := row == part.Row
			columnMatch := part.Column <= column && column < part.Column+part.Length
			if rowMatch && columnMatch {
				column += part.Length
				continue
			}

			// check bounds
			rowInBounds := 0 <= row && row < schema.Rows
			columnInBounds := 0 <= column && column < schema.Columns
			if rowInBounds && columnInBounds {
				neighbors = append(neighbors, Neighbor{schema.Content[row][column], column, row})
			}

			column++
		}
	}

	return neighbors
}
