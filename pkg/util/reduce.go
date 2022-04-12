package util

func Reduce(container map[rune]int, data map[rune]int) map[rune]int {
	for k, v := range data {
		if _, ok := container[k]; !ok {
			container[k] = 0
		}
		container[k] += v
	}
	return container
}
