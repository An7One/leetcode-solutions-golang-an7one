package kmp_algorithm

func build_kmp_table(pattern string) []int {
	var kmp_table = make([]int, len(pattern))

	var lo, hi = 0, 1

	for hi < len(pattern) {
		switch {
		case pattern[lo] == pattern[hi]:
			lo++
			kmp_table[hi] = lo
			hi++
		case lo == 0:
			hi++
		default:
			lo = kmp_table[lo-1]
		}
	}

	return kmp_table
}
