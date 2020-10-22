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

func main(){
  var temp, p int
  var slice []int
  nextImpar := generadorImpares()
  fmt.Print("Ingrese posición para fiboniacci: ")
  fmt.Scan(&temp)
  fmt.Println("Numero Fibonacci", Fibonacci(temp))
  fmt.Print("Ingrese longitud del arreglo: ")
  fmt.Scan(&p)
  for i:=0; i < p; i++{
    fmt.Print("Ingrese numero: ")
    fmt.Scan(&temp)
    slice = append(slice, temp)
  }
  
  fmt.Println("El número más grande fue", MasGrande(slice...))
  fmt.Print("Ingrese cantidad de numeros impares: ")
  fmt.Scan(&temp)
  
  for i:=1; i <= temp; i++{
    fmt.Println(nextImpar())
  }
}