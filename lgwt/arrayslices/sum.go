package arrayslices

// import "fmt"

// arrays have fixed sizes in the tpe
// slices

// takes an array of size 5 of intergers
// and return the sum of all numbers
func Sum(numbers []int) (sum int) {
	sum = 0
	// for i := 0; i < 5; i++ {
	// 	sum += numbers[i]
	// }

	// lets look at the range function
	for _, number := range numbers {
		// fmt.Printf("%d: adding %d to %d \n", _, number, sum)
		sum += number
	}
	return
}

// the ... []int in the params
// denotes size n parameters of type []int
func SumAll(numbersToSum ...[]int) []int {
	// numbersLength := len(numbersToSum)
	// make will create a slice of type int of starting size numbersLength
	// sums := make([]int, numbersLength)

	// for i, numbers := range numbersToSum {
	// 	sums[i] = Sum(numbers)
	// }

	// you can use append to create a NEW slice of size of old_slice+1
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
