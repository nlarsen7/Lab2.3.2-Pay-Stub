//Nicholas Larsen
//February 27, 2020
//Calculates a paystub depending on user input of their hourly wage and hours
package main

import "fmt"

func main() {
  var wage float64
  var workHours float64
  var name string
//creates the variables 
  fmt.Println("enter your name")
  fmt.Scanln(&name)

  fmt.Println("enter your hourly wage")
  fmt.Scanln(&wage)

  fmt.Println("enter how many hours you worked this week")
  fmt.Scanln(&workHours)
  //scans each variable
  
  overtime:=(workHours-40)
  
  if workHours>40 {
    workHours=workHours-overtime

  }else{
    overtime=0
    //makes overtime 0 so it doesn't mess with the paystub and makes sure workHours don't mess with it either
  }
  paystub:= ((workHours*wage)+(overtime*wage*1.5))- ((workHours*wage)+(overtime*wage*1.5))*.12
  //calculates paystub

  fmt.Println(name,"you made",paystub,"this week accounting for taxes and overtime")
//prints out the paystub

}