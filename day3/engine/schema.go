package engine

type Schema struct {
	Rows    int
	Columns int
	Content [][]rune
}

func newSchema(input []string) Schema {
	content := [][]rune{}
	for _, line := range input {
		runes := []rune(line)
		content = append(content, runes)
	}

	return Schema{len(content), len(content[0]), content}
}
