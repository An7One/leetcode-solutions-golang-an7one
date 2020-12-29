package lc0028

func strStr(haystack string, needle string) int {
	var kmpTable = buildKMPTable(needle)

	var idxH, idxN int = 0, 0

	for idxH < len(haystack) {
		if idxN == len(needle) {
			break
		}

		switch {
		case haystack[idxH] == needle[idxN]:
			idxH++
			idxN++

		case idxN == 0:
			idxH++

		default:
			idxN = kmpTable[idxN-1]
		}
	}

	if idxN == len(needle) {
		return idxH - idxN
	}

	return -1
}

func buildKMPTable(pattern string) []int {
	var kmpTable = make([]int, len(pattern))

	lo, hi := 0, 1

	for hi < len(pattern) {
		switch {
		case pattern[lo] == pattern[hi]:
			lo++
			kmpTable[hi] = lo
			hi++

		case lo == 0:
			hi++

		default:
			lo = kmpTable[lo-1]
		}
	}

	return kmpTable
}
