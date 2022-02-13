package main

func FindUniqueNumber(arr []float32) float32 {

	occurrence := make(map[float32]int)

	for n := range arr {
		value := occurrence[arr[n]]
		occurrence[arr[n]] = value + 1
	}

	for k, v := range occurrence {
		if v == 1 {
			return k
		}
	}

	return 0
}
