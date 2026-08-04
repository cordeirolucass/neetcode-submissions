type Solution struct{}


func (s *Solution) Encode(strs []string) string {
	var sb strings.Builder
	for _, str := range strs {
		sizeStr := len(str)
		sb.WriteString(strconv.Itoa(sizeStr))
		sb.WriteString("#")
		sb.WriteString(str)
	}

	return sb.String();
}

func (s *Solution) Decode(encoded string) []string {
	var decodedStrs []string

	i := 0
	for i < len(encoded) {
			j := i
			for encoded[j] != '#' {
				j++
			}

		sizeStr := encoded[i:j]
		originalSize, _ := strconv.Atoi(sizeStr)

		strStart := j+1
		strEnd := j+1+originalSize
		originalStr := encoded[strStart:strEnd]
		decodedStrs = append(decodedStrs, originalStr)

		i = strEnd
	}

	return decodedStrs;
}
