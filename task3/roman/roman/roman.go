package roman

type RomanMap[T int | int64] map[rune]int

func New[T int | int64]() RomanMap[T] {
	return RomanMap[T]{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
}

func (m RomanMap[T]) Convert(roman string) T {
	res := 0
	length := len(roman)

	for i := 0; i < length; i++ {
		currentVal := m[rune(roman[i])]

		if i+1 < length && currentVal < m[rune(roman[i+1])] {
			res -= currentVal
		} else {
			res += currentVal
		}
	}

	return T(res)
}