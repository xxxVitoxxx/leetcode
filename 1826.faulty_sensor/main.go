package main

func badSensor(sensor1, sensor2 []int) int {
	i, j, l := 0, 0, len(sensor1)
	for idx := 0; idx < l; idx++ {
		if sensor1[i] == sensor2[idx] {
			i++
		}
		if sensor2[j] == sensor1[idx] {
			j++
		}
	}

	if i == l-1 && j == l-1 {
		return -1
	} else if i == l-1 {
		return 1
	} else if j == l-1 {
		return 2
	}
	return -1
}

func badSensor2(sensor1, sensor2 []int) int {
	i, l := 0, len(sensor1)
	for ; i < l-1 && sensor1[i] == sensor2[i]; i++ {
	}

	for ; i < l-1; i++ {
		if sensor1[i] != sensor2[i+1] {
			return 2
		} else if sensor1[i+1] != sensor2[i] {
			return 1
		}
	}
	return -1
}
