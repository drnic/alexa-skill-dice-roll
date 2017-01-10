package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"time"

	"github.com/apex/go-apex"
	"github.com/b00giZm/golexa"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		app := golexa.Default()
		app.OnLaunch(func(a *golexa.Alexa, req *golexa.Request, session *golexa.Session) *golexa.Response {
			return a.Response().AddPlainTextSpeech("To ask me to roll a six-sided dice, say d6. To roll four six sided dice, say 4d6")
		})

		app.OnIntent(func(a *golexa.Alexa, intent *golexa.Intent, req *golexa.Request, session *golexa.Session) *golexa.Response {
			var err error
			if intent.Name == "RollDiceIntent" {
				log.Printf("RollDiceIntent: Slots: %v", intent.Slots)

				howManyDice := 1
				diceSides := 6

				if val, ok := intent.Slots["HowMany"]; ok {
					howManyDice, err = strconv.Atoi(val.Value)
					if err != nil {
						howManyDice = 1
						log.Println("RollDiceIntent: parsing HowMany: ", err)
					}
				}

				if val, ok := intent.Slots["DiceSides"]; ok {
					diceSides, err = strconv.Atoi(val.Value)
					if err != nil {
						log.Println("RollDiceIntent: parsing DiceSides: ", err)
						// perhaps 4d6 was heard as "46" (forty-six)
						re := regexp.MustCompile(`(\d*)(\d)`)
						parts := re.FindStringSubmatch(fmt.Sprintf("%d", howManyDice))
						howManyDice, _ = strconv.Atoi(parts[1])
						diceSides, _ = strconv.Atoi(parts[2])
					}
				}

				result, rolls := rollDice(howManyDice, diceSides)
				log.Printf("RollDiceIntent: rolling %d D %d result %d; dice %v\n", howManyDice, diceSides, result, rolls)

				return a.Response().AddPlainTextSpeech(fmt.Sprintf("rolling %d D %d ... %d", howManyDice, diceSides, result))
			}
			return a.Response().AddPlainTextSpeech("I don't know that intent yet.")
		})

		return app.Process(event)
	})
}

func rollDice(howManyDice, diceSides int) (result int, rolls []int) {
	result = 0
	rolls = make([]int, howManyDice)
	for i := 0; i < howManyDice; i++ {
		roll := rand.Intn(diceSides) + 1
		result += roll
		rolls[i] = roll
	}
	return
}
