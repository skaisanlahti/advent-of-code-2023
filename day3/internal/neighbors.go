package internal

type neighbor struct {
	symbol rune
	row    int
	column int
}

func findNeighbors(p enginePart, s *schema) []neighbor {
	var neighbors []neighbor

	startRow := p.row - 1
	endRow := p.row + 1
	startColumn := p.column - 1
	endColumn := p.column + p.length

	for row := startRow; row <= endRow; row++ {
		for column := startColumn; column <= endColumn; {
			if isDigitPosition(row, column, p) {
				column += p.length
				continue
			}

			if isWithinBounds(row, column, s) {
				neighbors = append(neighbors, neighbor{s.content[row][column], column, row})
			}

			column++
		}
	}

	return neighbors
}

const (
	digitRowOk = 1 << iota
	digitMinColumnOk
	digitMaxColumnOk
	digitOk = digitRowOk | digitMinColumnOk | digitMaxColumnOk
)

func isDigitPosition(row, column int, part enginePart) bool {
	var checks int
	if row == part.row {
		checks |= digitRowOk
	}

	if column >= part.column {
		checks |= digitMinColumnOk
	}

	if column < part.column+part.length {
		checks |= digitMaxColumnOk
	}

	return checks == digitOk
}

const (
	boundsMinRowOk = 1 << iota
	boundsMaxRowMaxOk
	boundsMinColumnOk
	boundsMaxColumnOk
	boundsOk = boundsMinRowOk | boundsMaxRowMaxOk | boundsMinColumnOk | boundsMaxColumnOk
)

func isWithinBounds(row, column int, schema *schema) bool {
	var checks int
	if row >= 0 {
		checks |= boundsMinRowOk
	}

	if row < schema.rows {
		checks |= boundsMaxRowMaxOk
	}

	if column >= 0 {
		checks |= boundsMinColumnOk
	}

	if column < schema.columns {
		checks |= boundsMaxColumnOk
	}

	return checks == boundsOk
}
