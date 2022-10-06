
package main

import "fmt"

func main() {

// Shorthand declaration of array
arr:= [5]string{"Java", "Php", "Javascript", "Go", "Python"}

fmt.Println("Elements of the array:")

for i:= 0; i < 3; i++{
fmt.Println(arr[i])
}

// Double dimention array
arr2 := [3][3]string{{"C #", "C", "Python"}, {"Java", "Scala", "Perl"},
        {"C++", "Go", "HTML"}}
  

    fmt.Println("Elements of Array 1")
    for x := 0; x < 3; x++ {
        for y := 0; y < 3; y++ {
  
            fmt.Println(arr2[x][y])
        }
    }

}
