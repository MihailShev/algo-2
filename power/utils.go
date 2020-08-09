package power

import "strconv"

func parseData(data []string) (float64, int, error) {
	a, err := strconv.ParseFloat(data[0], 64)

	if err != nil {
		return 0, 0, err
	}

	n, err := strconv.Atoi(data[1])

	if err != nil {
		return 0, 0, err
	}

	return a, n, nil
}
