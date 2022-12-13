package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	starting_place, received_error1 := strconv.Atoi(os.Args[1])
	if received_error1 != nil {
		print(received_error1)
	}

	number_customers_in_line, received_error2 := strconv.Atoi(os.Args[2])
	if received_error2 != nil {
		print(received_error2)
	}

	number_of_trips, received_error3 := strconv.Atoi(os.Args[3])
	if received_error3 != nil {
		print(received_error3)
	}

	a := []int{}
	for i := 0; i < number_customers_in_line; i++ {
		if i == starting_place-1 {
			a = append(a, 1)
		} else {
			a = append(a, 0)
		}
	}
	fmt.Println(a)

	var user_place = starting_place - 1
	var iteration_counter = 0
	for trips_left := number_of_trips; trips_left > 0; trips_left-- {
		fmt.Println(trips_left, iteration_counter, user_place)

		for user_place != 0 {
			a[user_place-1] = a[user_place]
			a[user_place] = 0
			user_place--
			iteration_counter++
			if iteration_counter%5 == 0 {
				a = append(a, 0)
			}
		}
		user_place = len(a) - 1
		a[user_place] = 1
		a[0] = 0
		iteration_counter++

	}
	fmt.Println("It took a total of ", iteration_counter, " iterations")

}

type Customer struct {
	line_position uint8
	name          string

	Career struct {
	}
}
