package tickets

import "fmt"

type Squad struct {
    Name      string
    People    int
    isSprint  bool
    Project   Project
}

type Project struct {
    Name       string
    Priority   int
    Technology []string
    Points      int
}


func (s Squad) Print() {
    fmt.Printf("Name: %s\n", s.Name)
    fmt.Printf("People: %s\n", s.People)
    fmt.Printf("An active Sprint: %s\n", s.isSprint)
}

func PinoProject() {

}




