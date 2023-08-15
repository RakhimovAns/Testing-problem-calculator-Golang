package functions

import (
	"github.com/RakhimovAns/Testing-problem-calculator-Golang/types"
)

func Parse(part string) (int, types.Kind, error) {
	switch part {
	case "I":
		return 1, "roman", nil
	case "II":
		return 2, "roman", nil
	case "III":
		return 3, "roman", nil
	case "IV":
		return 4, "roman", nil
	case "V":
		return 5, "roman", nil
	case "VI":
		return 6, "roman", nil
	case "VII":
		return 7, "roman", nil
	case "VIII":
		return 8, "roman", nil
	case "IX":
		return 9, "roman", nil
	case "X":
		return 10, "roman", nil
	case "1":
		return 1, "arabic", nil
	case "2":
		return 2, "arabic", nil
	case "3":
		return 3, "arabic", nil
	case "4":
		return 4, "arabic", nil
	case "5":
		return 5, "arabic", nil
	case "6":
		return 6, "arabic", nil
	case "7":
		return 7, "arabic", nil
	case "8":
		return 8, "arabic", nil
	case "9":
		return 9, "arabic", nil
	case "10":
		return 10, "arabic", nil
	default:
		return -1, "", types.ErrWrongNumber
	}
}

func ConvertToRoman(number int) string {
	romanNumerals := []string{
		"I", "II", "III", "IV", "V",
		"VI", "VII", "VIII", "IX", "X",
	}

	if number <= 10 {
		return romanNumerals[number-1]
	} else if number == 100 {
		return "C"
	} else {
		tens := number / 10
		ones := number % 10
		romanNumber := ""

		for i := 0; i < tens; i++ {
			romanNumber += "X"
		}

		if ones > 0 {
			romanNumber += romanNumerals[ones-1]
		}

		return romanNumber
	}
}
