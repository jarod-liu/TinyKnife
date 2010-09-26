package main

import (
  "os"
  "fmt"
  "flag"
  "rand"
)

const (
  Numbers = "0123456789"
  UpperAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
  LowerAlpha = "abcdefghijklmnopqrstuvwxyz"
  Special = "~!@#$%^&*()_-+=\\|}]{[:;\"'<,>.?/"
)

var withNumbers = flag.Bool("n", true, "with numbers characters")
var withUpperAlpha = flag.Bool("A", true, "with uppercase alphabet characters")
var withLowerAlpha = flag.Bool("a", true, "with lowercase alphabet characters")
var withSpecial = flag.Bool("s", true, "with non alpha-numeric characters")
var passLen = flag.Int("l", 12, "password length")
var passCount = flag.Int("c", 8, "password count")
var version = flag.Bool("v", false, "print version info")

func main() {
  flag.Parse()

	if *version {
		fmt.Printf("rpass v1.0925\n")
		os.Exit(0)
	}

  chars := ""
  if *withNumbers {
    chars += Numbers
  }
  if *withUpperAlpha {
    chars += UpperAlpha
  }
  if *withLowerAlpha {
    chars += LowerAlpha
  }
  if *withSpecial {
    chars += Special
  }
  
  charsLen := len(chars)
  sec, nsec, err := os.Time()
  if err != nil {
    fmt.Printf("", sec, err)
  }
  rand.Seed(nsec)
  for i:=0; i<*passCount; i++ {
    for j:=0; j<*passLen; j++ {
	  r := rand.Intn(charsLen)
	  fmt.Printf("%c", chars[r])
	}
	fmt.Println()
  }  
}
