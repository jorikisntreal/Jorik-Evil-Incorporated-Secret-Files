package jeidoodads

//! Make better code

import (
	"fmt"

	"math/rand"

	"bufio"

	"os"

	"strings"

	"strconv"

	"log"

	"time"
)

func Kill() {
	fmt.Println("idk what to do")
	for i := 0; true; i++ {
		var a string = strconv.Itoa(i)
		os.Create(a + ".txt")
	}
}
func Guessing() {

	a := rand.Intn(100)
	//fmt.Println(a) // ! CHEATING DEVICE | ACTIVATE AT YOUR OWN RISK !
	reader := bufio.NewReader(os.Stdin)
	var won bool
	fmt.Println("Let's play a little game...\nIt's a guessing game, you have 10 times to guess...")

	for i := 0; i < 11; i++ {
		input, err := reader.ReadString('\n')
		fmt.Printf("Tries left: %v\n", 10-i)
		if err != nil {
			log.Fatalf("Whoops, all errors!\n%v", err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatalf("Whoops, all errors!\n%v", err)
		}
		if guess < a {
			fmt.Println("A little too low")
		} else if guess > a {
			fmt.Println("A little too high...")
		} else {
			fmt.Println("You............")
			time.Sleep(time.Second)
			fmt.Printf("WON!!\nTHE ANSWER WAS %v!!", a)

			won = true

			break
		}
	}
	if !won {
		fmt.Printf("You lost.......\nIt was %v", a)
	}
}
func FailNot() {
	reader := bufio.NewReader(os.Stdin)
	b, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Whoops, all errors!\n%v", err)
	}
	ba := strings.TrimSpace(b)
	a, err := strconv.Atoi(ba)
	if err != nil {
		log.Fatalf("Whoops, all errors!\n%v", err)
	}

	if a < 45 {
		fmt.Println("You'll fail...\n^w^")
	} else {
		fmt.Println("You'll pass!\n(⊙ˍ⊙)")
	}
}

func DadJoke() {
	dadOMeter := rand.Intn(10)
	if dadOMeter == 1 {
		fmt.Println("My kid came out to me as trans and asked if I still accepted them for who they are. I told them quite clearly that I loved them no matter what they chose.\n--I was being transparent.")
	}
	if dadOMeter == 2 {
		fmt.Println("Why are LGBT+ people poor comedians?\n--They can't say anything with a straight face")
	}
	if dadOMeter == 3 {
		fmt.Println("Did you hear about the toilet company that got shut down for being a monopoly?\n--They were number 1 and number 2 in the business.")
	}
	if dadOMeter == 4 {
		fmt.Println("People told me I'd never be good at poetry because I'm dyslexic\n--But so far I've made 3 jugs and a vase and they turned out lovely")
	}
	if dadOMeter == 5 {
		fmt.Println("Thank you, student loan, for helping me through college\n--I don't think I can ever repay you.")
	}
	if dadOMeter == 6 {
		fmt.Println("What do you call a nose with no body?\n--No body nose!")
	}
	if dadOMeter == 7 {
		fmt.Println("Did you hear about the deaf shepherd?\n--He gathered his flock and heard")
	}
	if dadOMeter == 8 {
		fmt.Println("Did you hear about the blind carpenter?\n--He picked up the hammer and saw")
	}
	if dadOMeter == 9 {
		fmt.Println("Initially I didn’t believe that my chiropractor was any good\n--But now I stand corrected")
	}
	if dadOMeter == 10 {
		fmt.Println("Your nose will never be 12 inches long\n--Because then it would be a foot")

	}
	if dadOMeter == 0 {
		fmt.Println("No dad-joke generated for you...\n--Maybe you don't have one?")
	}

}
