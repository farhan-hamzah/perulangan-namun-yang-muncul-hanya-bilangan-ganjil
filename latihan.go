package main
import "fmt"

func main(){
    var masukan, i int
    fmt.Scan(&masukan)
    i = 1
    for i <= masukan{
        if i%2 != 0{
            fmt.Print(i, " ")
        }
        i++
    }
}