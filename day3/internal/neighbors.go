package internal

type neighbor struct {
	symbol rune
	row    int
	column int
}

func findNeighbors(part enginePart, schema *schema) []neighbor {
	minRow := part.row - 1
	maxRow := part.row + 1
	minColumn := part.column - 1
	maxColumn := part.column + part.length

	var neighbors []neighbor
	for row := minRow; row <= maxRow; row++ {
		for column := minColumn; column <= maxColumn; {
			// skip digit indexes
			rowMatch := row == part.row
			columnMatch := part.column <= column && column < part.column+part.length
			if rowMatch && columnMatch {
				column += part.length
				continue
			}

			// check bounds
			rowInBounds := 0 <= row && row < schema.rows
			columnInBounds := 0 <= column && column < schema.columns
			if rowInBounds && columnInBounds {
				neighbors = append(neighbors, neighbor{schema.content[row][column], column, row})
			}

			column++
		}
	}

	return neighbors
}
