package main

import (
  "fmt"
)

func Fibonacci(n int) int {
  if(n == 0 || n == 1){
    return 1
  }
  return (Fibonacci(n - 1) + Fibonacci(n - 2))
}

func MasGrande(args ...int) int {
  temp := args[0]
  
  for _, v := range args {
    if(temp < v){
      temp = v
      }
  }

  return temp
}

func generadorImpares() func() int {
	i := int(1)
	return func() int {
		var impar = i
		i += 2
		return impar
	}
}

func intercambia(a, b *int) {
  temp := *a
  *a = *b
  *b = temp
}

func main(){
  var temp, p int
  var slice []int
  nextImpar := generadorImpares()
  
  fmt.Println("--Secuencia Fibonacci--")
  fmt.Print("Ingrese posición para fiboniacci: ")
  fmt.Scan(&temp)
  fmt.Println("Numero Fibonacci", Fibonacci(temp))

  fmt.Println("--Implementacion Variadic--")
  fmt.Print("Ingrese longitud del arreglo: ")
  fmt.Scan(&p)
  for i:=0; i < p; i++{
    fmt.Print("Ingrese numero: ")
    fmt.Scan(&temp)
    slice = append(slice, temp)
  }
  
  fmt.Println("El número más grande fue", MasGrande(slice...))

  fmt.Println("--Generador Impares--")
  fmt.Print("Ingrese cantidad de numeros impares: ")
  fmt.Scan(&temp)
  
  for i:=1; i <= temp; i++{
    fmt.Println(nextImpar())
  }

  fmt.Println("--Intercambio de numeros--")
  fmt.Print("Ingresa numero a: ")
  fmt.Scan(&temp)
  fmt.Print("Ingresa numero b: ")
  fmt.Scan(&p)
  intercambia(&temp, &p)
  fmt.Println("Valor a:", temp)
  fmt.Print("Valor b:", p)
}