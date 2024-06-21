package main

func intersectionBruteforce(setA, setB []int) []int {
	var result []int

	if len(setB) < len(setA) {
		setA, setB = setB, setA
	}

	for i := range setA {
		for j := range setB {
			if setA[i] == setB[j] {
				result = append(result, setA[i])

				break
			}
		}
	}

	return result
}

func intersectionTwoPointer(setA, setB []int) []int {
	var result []int
	var i, j int //nolint:varnamelen // short scope

	if len(setB) < len(setA) {
		setA, setB = setB, setA
	}

	for i < len(setA) && j < len(setB) {
		switch {
		case setA[i] == setB[j]:
			result = append(result, setA[i])
			i++
			j++
		case setA[i] < setB[j]:
			i++
		default:
			j++
		}
	}

	return result
}

func main() {}
