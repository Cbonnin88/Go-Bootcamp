package employees

import "fmt"

type Employee interface {
    GetName() string
    GetAge() (string, int)
    GetStatus() (string, bool)
}

type IDnumber interface {
    GetIDNumber() []int
}

type Person struct {
    FirstName  string
    LastName   string
    Age        int
    Title      string
    Salary     int
    IsActive   bool
}

type Ticket struct {
    Name   string
    ID     []int
    Squad  string
}

func (p *Person) GetName() string {
    return "Employee Name: " + p.FirstName + " " + p.LastName
}

func (p *Person) GetAge() (string, int) {
    return "Employee age:", p.Age
}

func (p *Person) GetStatus() (string, bool) {
    return "Employee currently employed:", p.IsActive
}


func Details(e Employee) {
    fmt.Println(e.GetName())
    fmt.Println(e.GetAge())
    fmt.Println(e.GetStatus())
}


func (t *Ticket) GetIDNumber() []int {
    return t.ID
}

func TicketDetails(i IDnumber) {
    fmt.Println(i.GetIDNumber())
}


