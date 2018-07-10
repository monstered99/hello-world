// You can edit this code!
package main

import (
	"fmt"
	"net/http"
)

/*
func add(x, y float64) float64 {
	return x + y
}

func multiple(a, b string) (string, string) {
	return a, b
}
*/
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, Go is neat!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Expert web design by Sentdex")
}

func main() {
	//var num1, num2 float64 = 5.6, 9.5
	//fmt.Println(add(num1, num2))
	/*w1, w2 := "Hey", "there"
	fmt.Println(multiple(w1, w2))
	x := 15
	a := &x
	fmt.Println(x, a, *a)
	*/
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about/", aboutHandler)
	http.ListenAndServe(":8000", nil)
}

//Pierwiastek drugiego stopnia z n
/*
func main() {
	fmt.Println("Square root of 4 = ", math.Sqrt(4))
}
*/

//RNG od 1 do 100
/*
func main() {
	fmt.Println("A number from 1-100", rand.Intn(100))
}
*/
