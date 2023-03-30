package main

import "fmt"

type workedHoursError struct {
}

func (e workedHoursError) Error() string {
	return "Error: el trabajador no puede haber trabajado menos de 0 hs mensuales"
}

func calculateSalary(hours, hoursCost float32) (float32, error) {

	if hours < 80 {
		return 0, workedHoursError{}
	}

	total := hours * hoursCost

	if total <= 150000 {
		total *= 0.90
	}

	return total, nil
}

func main() {
	salary, err := calculateSalary(70, 20.0)

	if err != nil {
		panic(err)
	}
	fmt.Println(salary)

}
