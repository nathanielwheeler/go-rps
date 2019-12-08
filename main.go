package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	// initial prompt
	fmt.Print(`
RRRR           k        PPPP                        SSS                              
R   R          k k      P   P                      S          ii                     
RRRR  ooo  ccc kk       PPPP   aa ppp  eee rrr      SSS   ccc     ss  ss ooo rrr  ss 
R R   o o c    k k      P     a a p  p e e r           S c    ii  s   s  o o r    s  
R  RR ooo  ccc k  k     P     aaa ppp  ee  r       SSSS   ccc ii ss  ss  ooo r   ss  
                                  p                                                  
								  p                                                  

	`)

	// get input
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		err = runCommand(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func runCommand(commandStr string) error {
	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr)
	switch arrCommandStr[0] {
	case "exit":
	case "quit":
		os.Exit(0)
	case "plus":
	}
	cmd := exec.Command(arrCommandStr[0], arrCommandStr[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// generate random computer option
func botChoice() (botsChoice int) {

	// NOTE Always seed your RNG in Go before you use it
	rand.Seed(time.Now().UTC().UnixNano())

	return rand.Intn(3)
}
