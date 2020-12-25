package lc0028

func strStr(haystack string, needle string) int {
	var kmp_table = build_kmp_table(needle)

	var idx_h, idx_n int = 0, 0

	for idx_h < len(haystack) {
		if idx_n == len(needle) {
			break
		}

		switch {
		case haystack[idx_h] == needle[idx_n]:
			idx_h++
			idx_n++

		case idx_n == 0:
			idx_h++

		default:
			idx_n = kmp_table[idx_n-1]
		}
	}

	if idx_n == len(needle) {
		return idx_h - idx_n
	}

	return -1
}

func build_kmp_table(pattern string) []int {
	var kmp_table = make([]int, len(pattern))

	var lo, hi int = 0, 1

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
