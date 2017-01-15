package main

import (
    "github.com/go-vgo/robotgo"
)

func main() {
  robotgo.MoveMouseSmooth(250, 10, 1.0, 100.0)
  robotgo.MouseClick("left",false)
  robotgo.KeyTap("t", "control")
  robotgo.MoveMouseSmooth(130, 67, 1.0, 100.0)
  robotgo.MouseClick("left",false)
  robotgo.TypeString("https://www.strava.com/dashboard/following/100") //importing "Hello World"
  robotgo.KeyTap("enter")           //Press "enter"
} 