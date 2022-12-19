package problems

func IsPalindrome(x int) bool {
	var zeroVal int = x
	act := 0
	for x > 0 {
		rem := x % 10
		act = (act * 10) + rem
		x = x / 10
	}

	return zeroVal == act
}
