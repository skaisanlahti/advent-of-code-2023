package internal

type neighbor struct {
	char   rune
	row    int
	column int
}

func findNeighbors(p enginePart, s *schema) []neighbor {
	row := p.row
	column := p.column
	length := p.length

	rowCount := s.rows
	columnCount := s.columns
	schema := s.content

	var neighbors []neighbor

	// above
	if row-1 >= 0 {
		// left
		if column-1 >= 0 {
			neighbors = append(neighbors, neighbor{schema[row-1][column-1], row - 1, column - 1})
		}

		// middle
		for i, char := range schema[row-1][column : column+length] {
			neighbors = append(neighbors, neighbor{char, row - 1, column + i})
		}

		// right
		if column+length < columnCount {
			neighbors = append(neighbors, neighbor{schema[row-1][column+length], row - 1, column + length})
		}
	}

	// left
	if column-1 >= 0 {
		neighbors = append(neighbors, neighbor{schema[row][column-1], row, column - 1})
	}

	// right
	if column+length < columnCount {
		neighbors = append(neighbors, neighbor{schema[row][column+length], row, column + length})
	}

	// below
	if row+1 < rowCount {
		// left
		if column-1 >= 0 {
			neighbors = append(neighbors, neighbor{schema[row+1][column-1], row + 1, column - 1})
		}

		// middle
		for i, char := range schema[row+1][column : column+length] {
			neighbors = append(neighbors, neighbor{char, row + 1, column + i})
		}

		// right
		if column+length < columnCount {
			neighbors = append(neighbors, neighbor{schema[row+1][column+length], row + 1, column + length})
		}
	}

	return neighbors
}
