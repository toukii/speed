package speed

import (
    // "time"
    "fmt"
)

func Speed(a float32, intrvl int) float32  {
    fmt.Println(a,intrvl)
    return a*float32(intrvl)
}