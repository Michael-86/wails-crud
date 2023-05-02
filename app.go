package main

import (
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

//json.Unmarshal(brownBytes, &people)

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) shutdown(ctx context.Context) {

}

type Kabel struct {
	Fname string `json:"first"`
	Sname string `json:"last"`
}

// Print1 Prints all people passed and lists first person
func (a *App) Print1(people []Kabel) string {

	//--------------------------------------------
	//u, err := json.Marshal(Kabel(people))
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(u))
	//--------------------------------------------
	fmt.Println("all people:", people)
	fmt.Println("First person is", people[0].Fname, people[0].Sname)
	return fmt.Sprintf("Hello %s", people[0].Fname)
}

// GreetPerson Greet a singular person
func (a *App) GreetPerson(person Kabel) string {
	fmt.Println("Hello", person.Fname)
	return fmt.Sprintf("Hello %s", person.Fname)
}
