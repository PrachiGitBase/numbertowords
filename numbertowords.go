// Package numbertowords allows  numbers to be converted into English words
package numbertowords

import (
	"errors"
)

var words = [20]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "tweleve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}

var tenwords = [10]string{"", "", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}

//MaxNumber is the largest number that this version of numberstowords can convert into words
const MaxNumber = 99999

//Convert function convert the number 0 to maxnumber integer into words
func Convert(number int) (string, error) {

	if number < 0 || number > MaxNumber {

		return "", errors.New("number not in valid range")
	}

	tenThousandRes := ""

	tenthousand := number / 1000

	if tenthousand > 0 {

		if tenthousand < 20 {
			tenThousandRes = words[tenthousand] + " " + "thousand"
		} else {
			resultVal := tenthousand / 10
			tenThousandRes = tenwords[resultVal]

			secondNum := tenthousand % 10
			tenThousandRes = tenThousandRes + " " + words[secondNum] + " thousand"
		}
	}

	numberHund := number - (number/1000)*1000
	if numberHund > 99 {
		numberHunIndex := numberHund / 100
		if numberHunIndex > 0 {
			tenThousandRes = tenThousandRes + " " + words[numberHunIndex] + " hundread"
		}
	}

	decimalNumb := number - (number/100)*100
	if decimalNumb < 100 && decimalNumb > 0 {
		decimalNumbInd := decimalNumb / 10
		if decimalNumbInd > 1 {
			tenThousandRes = tenThousandRes + " " + tenwords[decimalNumbInd]

			unitIndex := decimalNumb % 10
			//fmt.Println(words[unitIndex])
			tenThousandRes = tenThousandRes + " " + words[unitIndex]
		} else {
			tenThousandRes = tenThousandRes + " " + words[decimalNumb]
		}
	}

	return tenThousandRes, nil
}
