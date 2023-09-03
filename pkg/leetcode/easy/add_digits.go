package easy

func addDigits(num int) int {
	var sum, s int

	for num != 0 {
		sum = sum + num%10
		num = num / 10
	}
	for sum != 0 {
		s = s + sum%10
		sum = sum / 10
	}
	return s
}

func addDigitsOne(num int) int {
	if num < 10 {
		return num
	}
	var dSum int
	for num != 0 {
		dSum = dSum + num%10
		num = num / 10
	}
	return addDigitsOne(dSum)
}

func addDigitsTwo(num int) int {
	if num < 10 {
		return num
	}
	if res := num % 9; res == 0 {
		return 9
	} else {
		return res
	}

}
