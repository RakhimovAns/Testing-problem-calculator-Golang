package main

import (
	"bufio"
	"github.com/RakhimovAns/Testing-problem-calculator-Golang/functions"
	"github.com/RakhimovAns/Testing-problem-calculator-Golang/types"
	"log"
	"os"
	"strings"
)

func main() {
	Reader := bufio.NewReader(os.Stdin)
	text, err := Reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	parts := strings.Split(text, " ")
	if len(parts) < 3 {
		log.Fatal(types.ErrNotMathEq)
	}
	if len(parts) > 3 {
		log.Fatal(types.ErrWrongFormat)
	}
	First, KindOfFirst, err := functions.Parse(parts[0])
	if err != nil {
		log.Fatal(err)
	}
	secondary := strings.Split(parts[2], "\r\n")
	Second, KindOfSecond, err := functions.Parse(secondary[0])
	if err != nil {
		log.Fatal(err)
	}
	if KindOfFirst != KindOfSecond {
		log.Fatal(types.ErrWrongKind)
	}
	var answer int
	if parts[1] == "+" {
		answer = First + Second
		if KindOfFirst == "roman" {
			log.Println(functions.ConvertToRoman(answer))
			return
		}
	} else if parts[1] == "-" {
		answer = First - Second
		if KindOfFirst == "roman" {
			if answer < 0 {
				log.Fatal(types.ErrNegativeRoman)
			}
			if answer == 0 {
				log.Fatal(types.ErrZero)
			}
			log.Println(functions.ConvertToRoman(answer))
			return
		}
	} else if parts[1] == "*" {
		answer = First * Second
		if KindOfFirst == "roman" {
			log.Println(functions.ConvertToRoman(answer))
			return
		}
	} else if parts[1] == "/" {
		answer = First / Second
		if KindOfFirst == "roman" {
			log.Println(functions.ConvertToRoman(answer))
			return
		}
	}
	log.Println(answer)
}
