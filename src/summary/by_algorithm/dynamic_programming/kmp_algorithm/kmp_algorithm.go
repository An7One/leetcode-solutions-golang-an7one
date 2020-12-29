package kmp_algorithm

func build_kmp_table(pattern string) []int {
	var kmpTable = make([]int, len(pattern))

	var lo, hi = 0, 1

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
