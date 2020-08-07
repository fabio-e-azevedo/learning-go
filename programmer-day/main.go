package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.Now()
    if t.YearDay() == 256 {
        fmt.Println("Programmer Congratulations!!!")
    } else {
        fmt.Println("...")
    }
}
