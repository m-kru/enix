package line

func Empty() *Line {
	return &Line{Buf: "", Prev: nil, Next: nil}
}

func FromString(str string) *Line {
	if len(str) == 0 {
		return &Line{Buf: ""}
	}

	startIdx := 0
	var firstLine *Line = nil
	var line *Line
	var nextLine *Line

	for i, r := range str {
		if r == '\n' || i == len(str)-1 {
			if firstLine == nil {
				firstLine = &Line{Buf: str[startIdx:i]}
				line = firstLine
				startIdx = i + 1
			} else {
				if r == '\n' {
					nextLine = &Line{Buf: str[startIdx:i], Prev: line}
				} else {
					nextLine = &Line{Buf: str[startIdx : i+1], Prev: line}
				}
				line.Next = nextLine
				line = nextLine
				startIdx = i + 1
			}
		}
	}

	return firstLine
}
