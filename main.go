package main

import (
	"fmt"
	"math/rand"
	"time"
)

// generate random computer option
func botChoice() (botsChoice int) {

	// NOTE Always seed your RNG in Go before you use it
	rand.Seed(time.Now().UTC().UnixNano())

	return rand.Intn(3)
}

func main() {
	// prompt
	fmt.Print(`
RRRR           k        PPPP                        SSS                              
R   R          k k      P   P                      S          ii                     
RRRR  ooo  ccc kk       PPPP   aa ppp  eee rrr      SSS   ccc     ss  ss ooo rrr  ss 
R R   o o c    k k      P     a a p  p e e r           S c    ii  s   s  o o r    s  
R  RR ooo  ccc k  k     P     aaa ppp  ee  r       SSSS   ccc ii ss  ss  ooo r   ss  
                                  p                                                  
                                  p                                                  
	`)
	fmt.Print("Enter 'rock', 'paper', or 'scissors'")

	// get input
	var input string
	fmt.Scanln(input)
}
