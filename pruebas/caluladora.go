/*
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("ingrese la operacion suma ejemplo 2+2")
	scanner.Scan()
	operacion := scanner.Text()
	fmt.Println("la operacion ingresada es: " + operacion)
	valores := strings.Split(operacion, "+")
	fmt.Println(valores)
	fmt.Println(valores[0] + valores[1])
	operador1, err1 := strconv.Atoi(valores[0])
	if err1 != nil {
		fmt.Println(err1)
	}
	operador2, _ := strconv.Atoi(valores[1])

	resultado := operador1 + operador2
	fmt.Println(resultado)

}

*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	ns := "hacker"
	arr := strings.Split(ns, "")
	fmt.Println(arr)
	var arreven []string
	var arrodd []string
	for i, s := range arr {
		if i%2 == 0 {
			arreven = append(arreven, s)
		}
	}
	for i, s := range arr {
		if i%2 != 0 {
			arrodd = append(arrodd, s)
		}
	}
	for i := 0; i < len(arreven); i++ {
		fmt.Print(arreven[i])
	}
	fmt.Print(" ")
	for i := 0; i < len(arrodd); i++ {
		fmt.Print(arrodd[i])
	}
	fmt.Println("")
	h := strings.Join(arreven, "")
	fmt.Println(h)
}
