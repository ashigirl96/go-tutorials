package main

import "fmt"

func helloEnd() {
    defer fmt.Println("End")
    fmt.Println("Hello")
}

func panicEnd() {
    defer fmt.Println("End")
    fmt.Println("Hello")
    panic("Panic!!")
}

func recoverEnd() {
    defer func() {
        fmt.Println("End")
        err := recover()
        if err != nil {
            fmt.Println("Recover!:", err)
        }
    }()

    fmt.Println("Start")
    panic("Panic!")
}

func main() {
    helloEnd()
    fmt.Println()
    // panicEnd()
    // fmt.Println()
    recoverEnd()

}
