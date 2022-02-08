package main

import (
	"strings"
)

func HighScoreWord(s string) string {

	alpha := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	//string to slice
	arr := strings.Fields(s)

	hScore := 0    //high score
	wrdScore := "" //word high score

	//iterate slice
	for _, w := range arr {

		//calc score by word
		score := 0
		letters := []rune(strings.ToUpper(w)) //word to slice

		for i := 0; i < len(letters); i++ { //iterate slice word

			l := string(letters[i])                     //convert rune to string
			score = score + strings.Index(alpha, l) + 1 //add to score letter by letter
		}

		if score > hScore { //verify if highscore
			wrdScore = w
			hScore = score
		}

	}

	return wrdScore

}
