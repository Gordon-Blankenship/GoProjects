package main

import (
	"fmt"
)

func main() {
	var dealerships [3]Dealership

	carList1 := []Car{}
	carList1 = append(carList1, Car{"Tesla", "Model3", 2018, 55000}, Car{"Ford", "F-150", 2020, 60000}, Car{"Chevrolet", "Silverado", 2019, 47000}, Car{"Honda", "Civic", 2022, 45000}, Car{"Toyata", "Corolla", 2014, 18000})

	carList2 := []Car{}
	carList2 = append(carList2, Car{"Honda", "CR-V", 2022, 36000}, Car{"Toyota", "Camry", 2019, 21000}, Car{"Ford", "Escape", 2020, 26000}, Car{"Chevrolet", "Equinox", 2022, 32000}, Car{"Toyota", "Prius", 2021, 28000})

	carList3 := []Car{}
	carList3 = append(carList3, Car{"Honda", "Accord", 2021, 30000}, Car{"Mazda", "CX-5", 2022, 32000}, Car{"Ford", "Explorer", 2019, 31000}, Car{"Lincoln", "Navigator", 2021, 60000}, Car{"Cadillac", "Escalade", 2022, 70000})

	dealerships[0] = Dealership{"Mike's Auto Park", carList1}
	dealerships[1] = Dealership{"O'Neil Family Vehicles", carList2}
	dealerships[2] = Dealership{"Easy Cars Auto Emporeum", carList3}

	choice := -1
	curDealership := 0
	fmt.Println("Welcome to the Car Dealership Simulation in Golang")
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)
	profile := Customer{name, []Car{}, 100000}

	for choice != 0 {
		fmt.Println("\nYou are at the following dealership: ", dealerships[curDealership].name)
		choice = getChoice()
		switch choice {
		case 1:
			curDealership = goToNewDealership(curDealership, dealerships[:])
		case 2:
			newCarIndex := buyACar(dealerships[curDealership])
			dealerCars := dealerships[curDealership].car_list
			if int(dealerCars[newCarIndex].price) > profile.balance {
				fmt.Println("You didn't have enough money to purchase the car")
				fmt.Println("Current Balance: ", profile.balance)
			} else {
				purchasedCar := dealerCars[newCarIndex]
				profile.car_list = append(profile.car_list, purchasedCar)
				profile.balance -= int(purchasedCar.price)
				newCarList := []Car{}
				for i := 0; i < len(dealerCars); i++ {
					if i != newCarIndex {
						newCarList = append(newCarList, dealerCars[i])
					}
				}
				dealerships[curDealership].car_list = newCarList
				fmt.Println("You purchased a ", purchasedCar.year, " ", purchasedCar.make, " ", purchasedCar.model)

			}

		case 3:
			sellCarIndex := sellACar(profile)
			updatedGarage := []Car{}
			for i := 0; i < len(profile.car_list); i++ {
				if i != sellCarIndex {
					updatedGarage = append(updatedGarage, profile.car_list[i])
				}
			}
			profile.car_list = updatedGarage

		case 4:
			fmt.Print("Your current balance: ", profile.balance)

		case 5:
			fmt.Println("\nHere are your cars: ")
			custCarList := profile.car_list
			for i := 0; i < len(custCarList); i++ {
				curCar := custCarList[i]
				fmt.Println(curCar.make, " ", curCar.model, " ", curCar.year, " ", curCar.price)
			}
		}
	}
	fmt.Println("\nThank you for playing!")

}

func getChoice() int {
	fmt.Println("\nHere are your choices: ")
	fmt.Println("1. Go To a New Dealership")
	fmt.Println("2. View Car Lineup/Buy a Car")
	fmt.Println("3. Sell a Car")
	fmt.Println("4. Check Your Balance")
	fmt.Println("5. View Your Cars")
	fmt.Println("0. Quit")
	fmt.Print("\nEnter Choice: ")
	var choice int
	fmt.Scan(&choice)
	return choice
}

func goToNewDealership(curDealership int, dealerships []Dealership) int {
	for i := 0; i < len(dealerships); i++ {
		if i != curDealership {
			fmt.Println(i, " ", dealerships[i].name)
		}
	}
	fmt.Print("\nEnter the number next to the dealership of your choice: ")
	var newDealershipIndex int
	fmt.Scan(&newDealershipIndex)
	return newDealershipIndex
}

func buyACar(dealership Dealership) int {
	car_list := dealership.car_list
	fmt.Println("\nHere are all the cars available at ", dealership.name)
	for i := 0; i < len(car_list); i++ {
		fmt.Println(i, " ", car_list[i])
	}
	fmt.Print("\nEnter the number next to the Car of your choice: ")
	var newCarIndex int
	fmt.Scan(&newCarIndex)
	return newCarIndex
}

func sellACar(profile Customer) int {
	car_list := profile.car_list
	fmt.Println("\nHere are all your cars that you can sell: ")
	for i := 0; i < len(car_list); i++ {
		fmt.Println(i, " ", car_list[i])
	}
	fmt.Print("\nEnter the number next to the Car you'd like to sell: ")
	var sellCarIndex int
	fmt.Scan(&sellCarIndex)
	return sellCarIndex
}

type Dealership struct {
	name     string
	car_list []Car
}

type Car struct {
	make  string
	model string
	year  uint
	price uint
}

type Customer struct {
	name     string
	car_list []Car
	balance  int
}
