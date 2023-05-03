package main

import (
    "fmt"
)

// fungsi untuk menghitung rata-rata dari beberapa bilangan desimal
func average(nums ...float64) float64 {
    total := 0.0
    for _, n := range nums {
        total += n
    }
    return total / float64(len(nums))
}



func main() {
    // contoh pemanggilan fungsi average
    avg1 := average(2.5, 3.0, 4.5, 1.0)
    fmt.Println(avg1) // output: 2.75
    
    avg2 := average(14.2, 24.1, 27.6, 49.6, 52.3)
    fmt.Println(avg2) // output: 33.56
    
}
