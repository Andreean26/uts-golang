package main

import "fmt"

func main() {
    // inisialisasi array pertama
    array1 := []string{"a", "c", "d"}
    
    // inisialisasi array kedua
    array2 := []string{"a", "d", "c"}     
    
    // melakukan perulangan untuk membandingkan isi dari kedua array
    for i, value := range array1 {
        if value != array2[i] {
            fmt.Printf("index ke %d berbeda\n", i)
        }
    }
}
