package main

import (
    "Go-BootCamp/Examples/Interface/employees"
    "fmt"
)

func main() {
    worker := &employees.Person{
                          FirstName: "Martha",
                          LastName:  "Steward",
                          Age:       79,
                          Title:     "Doctor",
                          Salary:    200000,
                          IsActive:  true,
                          }
    employees.Details(worker)
    fmt.Println("\n")


    pino := &employees.Ticket{
                       Name: "M6-Paywall",
                       ID: []int{2019,2018,2201,2392,1212,1123},
                       Squad: "PINO",
    }
    employees.TicketDetails(pino)






}
