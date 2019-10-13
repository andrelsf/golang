package main

import (
    "fmt"
    "os"
    "strconv"
    "math/rand"
    "time"
)

func password_generate(amount int64, size int64) {
    const values = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ#%$@*&!?+{}[]"
    min, max := 0, len(values)
    var x int64 = 0
    for x < amount {
        password := ""
        var i int64 = 0
        for i < size {
            rand.Seed(time.Now().UnixNano())
            var num int = rand.Intn(max - min)
            password = password + values[num:num+1]
            i++
        }
        fmt.Println(password)
        x++
    }
}

func help() {
    fmt.Println("GO Password Generate\n" +
                "Uso: gopg -a <amount> -s <size>\n" +
                "   -a : passwords amount.\n" +
                "   -s : passwords size.\n" +
                "golang \\o/")
    os.Exit(1)
}

func main() {
    if len(os.Args) <  5 {
        help()
    }
    p := make([]int64, 2)
    params := os.Args[:]
    for i, param := range params {
        if i == 1 {
            if param == "-a" {
                continue
            } else {
                help()
            }
        } else if i == 2 {
            amount, err := strconv.ParseInt(param, 0, 8)
            if err != nil {
                help()
            }
            p[0] = amount
        } else if i == 3 {
            if param == "-s" {
                continue
            } else {
                help()
            }
        } else if i == 4 {
            size, err := strconv.ParseInt(param, 0, 8)
            if err != nil {
                help()
            }
            p[1] = size
        }
    }
    password_generate(p[0], p[1])
}
