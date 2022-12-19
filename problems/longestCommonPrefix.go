package problems

func LongestCommonPrefix(strs []string) string {

	word1 := strs[0]
	common := ""
	for _, word := range strs {
		for _, l := range word {
			for _, k := range word1 {
				if l == k {
					common += string(l)
				}
			}
		}
	}
	return common
}
