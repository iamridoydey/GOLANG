package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func guessGame() {
    // Pick a random number
    rand.Seed(time.Now().UnixNano())
    targetNumber := rand.Intn(10) + 1
    var guess, attempts int

    fmt.Println("Guess Game!")

    for {
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Guess a number 1 to 10: ")

        strNum, _ := reader.ReadString('\n')
        strNum = strings.TrimSpace(strNum)
        num, err := strconv.Atoi(strNum)
        guess = num
        attempts++

        if err != nil {
            fmt.Println("Not a valid input")
            continue
        }

        if guess < targetNumber {
            fmt.Println("Too low!")
        } else if guess > targetNumber {
            fmt.Println("Too high!")
        } else {
            fmt.Printf("Congratulations! You guessed it in %d attempts\n", attempts)
            break
        }
    }
}
