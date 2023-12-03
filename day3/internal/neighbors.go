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
			if isDigitPosition(row, column, part) {
				column += part.length
				continue
			}

			if isWithinBounds(row, column, schema) {
				neighbors = append(neighbors, neighbor{schema.content[row][column], column, row})
			}

			column++
		}
	}

	return neighbors
}

const (
	digitRowOk = 1 << iota
	digitColumnOk
)

func isDigitPosition(row, column int, part enginePart) bool {
	var result int
	if row == part.row {
		result |= digitRowOk
	}

	if part.column <= column && column < part.column+part.length {
		result |= digitColumnOk
	}

	return result == digitRowOk|digitColumnOk
}

const (
	boundsRowsOk = 1 << iota
	boundsColumnsOk
)

func isWithinBounds(row, column int, schema *schema) bool {
	var result int
	if 0 <= row && row < schema.rows {
		result |= boundsRowsOk
	}

	if 0 <= column && column < schema.columns {
		result |= boundsColumnsOk
	}

	return result == boundsRowsOk|boundsColumnsOk
}
