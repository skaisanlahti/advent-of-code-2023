package engine

type schema struct {
	rows    int
	columns int
	content [][]rune
}

func newSchema(input []string) schema {
	var content [][]rune
	for _, line := range input {
		runes := []rune(line)
		content = append(content, runes)
	}

	return schema{len(content), len(content[0]), content}
}
