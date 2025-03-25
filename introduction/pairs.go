package introduction

func pairs(arr []string) [][]string {

	pairs := make([][]string, 0)

	for i := 0; i < len(arr); i++ {

		for j := i + 1; j < len(arr); j++ {

			pair := []string{arr[i], arr[j]}

			pairs = append(pairs, pair)

		}

	}

	return pairs
}
