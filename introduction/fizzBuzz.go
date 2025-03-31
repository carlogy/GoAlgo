package introduction

func fizzBuzz(num int) []interface{} {

	result := make([]interface{}, 0)

	for i := 1; i <= num; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, "fizzbuzz")
		} else if i%5 == 0 {
			result = append(result, "buzz")
		} else if i%3 == 0 {
			result = append(result, "fizz")
		} else {
			result = append(result, i)
		}
	}

	return result
}
