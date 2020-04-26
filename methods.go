package main

import "fmt"

type employee struct {
	employee_id uint16 //min 0 max 65536
	name        string
	role        string
}

func (e employee) getData() string {
	return fmt.Sprintf("Employee id of %s is %d", e.name, e.employee_id)
}

func main() {
	emp := employee{employee_id: 1, name: "Homi", role: "developer"}

	fmt.Println(emp)
	fmt.Println(emp.getData())
}
