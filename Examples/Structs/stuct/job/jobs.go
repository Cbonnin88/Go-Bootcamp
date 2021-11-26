package job

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    FirstName   string
    LastName    string
    Age         int
    Gender      string
    Position    Position
}

type Position struct {
    isWorking   bool
    ID          int
    Title       string
    Rank        int
    Salary      int
}

func Working() {
   emp1:=  Person{
           FirstName: "John",
           LastName: "Smith",
           Age: 30,
           Gender: "M",
           Position: Position{
                     isWorking: true,
                     ID: 2342,
                     Title: "Customer Success Manager",
                     Rank: 3,
                     Salary: 45000,
               },

        }
   byteArray, err := json.MarshalIndent(emp1, "", " ")
   if err != nil {
       fmt.Println(err)
   }
   fmt.Println(string(byteArray))
}

type Engineer struct {
    Name  string
    Age   int
    Title string
    Project  Project
}

type Project struct {
    WorkTitle  string
    Priority   string
    Techno     []string
}

func Projects() {
    pjct := Engineer{
            Name: "Alice Walker",
            Age:  50,
            Title: "Backend Engineer",
            Project: Project{
                     WorkTitle: "M6-Paywall",
                     Priority: "P1",
                     Techno: []string{"Go", "Php","React", "MariaDB"},
            },
    }

    byteArray, err := json.MarshalIndent(pjct, "", " ")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(byteArray))
    fmt.Println(pjct.Project.Techno[0],pjct.Project.Techno[3])


}

