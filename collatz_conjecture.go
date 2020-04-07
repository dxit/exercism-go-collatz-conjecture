package collatzconjecture

import "errors"

// CollatzConjecture returns the numbers of steps required to reach 1
func CollatzConjecture(number int) (int, error) {
	if number < 1 {
		return 0, errors.New("unexpected")
	}

	i := 0
	for {
		if number == 1 {
			break
		}
		if number%2 == 0 {
			number /= 2
		} else {
			number = (number * 3) + 1
		}
		i++
	}

	return i, nil
}
