package tickets

import (
    "fmt"
)

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
    fmt.Printf("Squad Name: %s\n", s.Name)
    fmt.Printf("Squad Amount: %d\n", s.People)
    fmt.Printf("Project Name:%s\n", s.Project.Name)
    fmt.Println("Technology used:", s.Project.Technology[0],s.Project.Technology[3])
    fmt.Printf("Priority: %d\n", s.Project.Priority)
    fmt.Printf("Points: %d\n", s.Project.Points)
    fmt.Printf("An active Sprint: %t\n", s.isSprint)
}

func (s *Squad) UpdatePoints() int {
     s.Project.Points += 2
     return s.Project.Points
}

func (s *Squad) GetProjectPriority() int {
    return s.Project.Priority
}

func PinoProject() {
pino := Squad{
       Name: "Pino",
       People: 20,
       isSprint: true,
       Project: Project{
                Name: "M6-Paywall",
                Priority: 1,
                Technology: []string{"Go","Php","JavaScript","Docker","React"},
                Points: 13,
       },
    }
    pino.Print()
    fmt.Println("New priority point:",pino.UpdatePoints())
    fmt.Println("Current Project Priority:",pino.GetProjectPriority())
}




