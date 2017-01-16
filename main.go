package main

import (
    "github.com/go-vgo/robotgo"
    "time"
)

func main() {

  robotgo.MoveMouseSmooth(800, 67, 1.0, 100.0)
  robotgo.MoveMouseSmooth(250, 10, 1.0, 100.0)
  robotgo.MouseClick("left",false)
  robotgo.KeyTap("t", "control")
  robotgo.MoveMouseSmooth(130, 67, 1.0, 100.0)
  robotgo.MouseClick("left",false)
  robotgo.TypeString("https://www.strava.com/dashboard/following/100")
  robotgo.KeyTap("enter")
  time.Sleep(10* time.Second)
  robotgo.MoveMouseSmooth(200, 67, 1.0, 100.0)
  robotgo.MouseClick("left",false)
  robotgo.TypeString("javascript:jQuery(document).ready( function($){ $('.js-add-kudo').each(function(){ $(this).click() }); });")
  robotgo.KeyTap("enter")

  } 