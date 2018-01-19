package main

import (
    "fmt"
    "os"
    "os/exec"
	 "runtime"
	 "strconv"
    "time"
)

import gof "./GameOfLife"

/*
 * From: https://stackoverflow.com/questions/22891644/how-can-i-clear-the-terminal-screen-in-go
 */
var clear map[string]func() //create a map for storing clear funcs

func init() {
    clear = make(map[string]func()) //Initialize it
    clear["linux"] = func() { 
        cmd := exec.Command("clear") //Linux example, its tested
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
    clear["windows"] = func() {
      //   cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
        cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested 
        cmd.Stdout = os.Stdout
        cmd.Run()
	 }
}


func CallClear() {
    value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
    if ok { //if we defined a clear func for that platform:
        value()  //we execute it
    } else { //unsupported platform
        panic("Your platform is unsupported! I can't clear terminal screen :(")
    }
}
// **************************************



func main() {
	cols, rows, nGen := 10, 10, 100
	var seed int
	var err error

	n := len(os.Args)
	if n == 2 { // Gen number
		nGen, err = strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
	}

	if n == 3 { // cols and rows number
		cols, err = strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		rows, err = strconv.Atoi(os.Args[2])
		if err != nil {
			panic(err)
		}
	}

	if n >= 4 { // generations, cols and rows
		nGen, err = strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		cols, err = strconv.Atoi(os.Args[2])
		if err != nil {
			panic(err)
		}
		rows, err = strconv.Atoi(os.Args[3])
		if err != nil {
			panic(err)
		}

		if n == 5 {
			seed, err = strconv.Atoi(os.Args[4])
			if err != nil {
				panic(err)
			}
		}
	}

	fmt.Println("** Initializing with:")
	fmt.Printf("%d generations of %d x %d matrix\n", nGen, cols, rows)
	fmt.Print("Press ENTER to start")
	var s string
	fmt.Scanf("%s", &s)

	g := gof.GameOfLife{}
	g.Init(cols, rows, nGen, int64(seed))
	g.GenerateAll()

	for _, gen := range g.Generations {
		for i := range gen {
			for j := range gen[i] {
				fmt.Printf("%d ", gen[i][j])
			}
			fmt.Printf("\n")
		}

		time.Sleep(50 * time.Millisecond)
		CallClear()
	}
}