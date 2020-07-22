/*
Author: Jasmin Smith
Date: 06/17/2020
*/
package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// What makes up our Hero
type Hero struct {
	Health       int
	MaxHealth    int
	AttackPower  int
	DefensePower int
	MaxDefense   int
	Level        int
	GoblinKills  int
	Potions      [5]int
	Gold         int
}

// Keeping up with initial Health and Defense for Hero

// What makes up our goblins
type Goblin struct {
	MaxHealth    int
	AttackPower  int
	DefensePower int
}

// Builds our Hero
func (hero *Hero) assembleHero() {
	maxHealth := rand.Intn(11) + 20
	maxDefense := rand.Intn(5) + 1
	hero.GoblinKills = 0
	hero.MaxHealth = maxHealth
	hero.Health = maxHealth
	hero.AttackPower = rand.Intn(3) + 1
	hero.MaxDefense = maxDefense
	hero.DefensePower = maxDefense
	hero.Potions = [5]int{2, 2, 2, 2, 2}
	hero.Gold = 0
	hero.Level = 1

}

// builds our goblins
func assembleGoblin() *Goblin {

	var goblin = Goblin{
		MaxHealth:    rand.Intn(6) + 5,
		AttackPower:  rand.Intn(2) + 2,
		DefensePower: rand.Intn(2) + 1,
	}

	return &goblin
}

// fights Goblin
func combat(hero *Hero) bool {
	goblin := assembleGoblin()
	alive := true
	// Keeps looping until there is only one alive
	for alive == true {
		// Makes sure the goblin isn't getting damage after the hero is defeated

		if hero.Health != 0 {
			// hero attacks goblin
			for i := 0; i < hero.AttackPower; i++ {
				// hits defense power first
				if goblin.DefensePower == 0 {
					goblin.MaxHealth -= 1
					if goblin.MaxHealth == 0 {
						alive = false
						break
					}
				} else {
					goblin.DefensePower -= 1
				}
			}
		}
		// makes sure hero is not getting damage after the goblin is defeated
		if goblin.MaxHealth != 0 {
			// goblin attacks
			for i := 0; i < goblin.AttackPower; i++ {
				if hero.DefensePower == 0 {
					hero.Health -= 1
					if hero.Health == 0 {
						alive = false
						break
					}
				} else {
					hero.DefensePower -= 1
				}
			}
		}
	}
	// returns if hero lost
	if hero.Health == 0 {
		return false
	}
	// returns if hero won
	return true
}

// Hero levels up
func levelUp(hero *Hero) {
	var potionResponse, response string
	var potions int = 0
	var isValid bool = false
	hero.Level += 1
	fmt.Println("Let me restore your defense power for you")
	// Restoring Defense Power
	for hero.DefensePower < hero.MaxDefense {
		hero.DefensePower += 1
	}
	potionCount := hero.potionCount()
	// Shops for potion
Prompt:
	fmt.Println("You have", hero.Gold, "pieces of gold")
	fmt.Println("You have", potionCount, "many potions")
	fmt.Print("Would you like to buy more potions for four pieces of gold per potion? ")
	fmt.Scan(&response)
	response = strings.ToLower(response)
	if response == "yes" {

		for isValid == false {
			fmt.Print("How many would you like to buy? ")
			fmt.Scan(&potionResponse)
			potions, err := strconv.Atoi(potionResponse)
			if err != nil || (potions > potionCount && potions < 0) {
				fmt.Println("You do not have room in your bag :/")
			} else if hero.Gold < potions*4 {
				fmt.Println("You are too broke to pay for this")
			} else {
				hero.Gold -= potions * 4
				isValid = true
			}
		}
		j := 0
		// gives potions that hero purchaes
		for i := 0; i < potions; i++ {
			if j < potions && hero.Potions[i] == 0 {
				hero.Potions[i] = 2
				j++
			}
		}

	} else if response != "no" {
		fmt.Println("That is an invalid choice sir")
		goto Prompt
	}

}

// Determines if hero runs into a goblin
func encounter(hero *Hero) bool {
	rand.Seed(time.Now().UnixNano())
	fight := rand.Intn(3)
	if fight == 0 {
		return true
	}
	return false
}

// See if Hero can take potions
func takePotion(hero *Hero) bool {
	potionCount := hero.potionCount()
	var response, answer string

	if potionCount == 0 {
		return false
	}
Prompt:
	fmt.Println("You have", potionCount, "potions!")
	fmt.Print("Would you like to take a potion? ")
	fmt.Scan(&answer)
	answer = strings.ToLower(answer)
	if answer == "yes" {
		fmt.Print("How many potions would you like to take? ")
		fmt.Scan(&response)
		amount, _ := strconv.Atoi(response)

		if amount < 0 || amount > potionCount {
			fmt.Println("That is not a valid amount")
			goto Prompt
		} else {
			j := 0
			// Takes up to five potions
			for i := 0; i < len(hero.Potions); i++ {
				if hero.Health != hero.MaxHealth-1 && hero.Health < hero.MaxHealth {
					hero.Health += 2
					hero.Potions[i] = 0
					j++
					if j == amount {
						fmt.Println("You have taken your potions. Your current health is " + strconv.Itoa(hero.Health))
					}

					//Takes only half of potion
				} else if hero.Health == hero.MaxHealth-1 {
					hero.Health += 1
					hero.Potions[i] = 0
					j++
				} else {
					fmt.Println("You are at max health.")

					break
				}
				if j <= amount && hero.Health > hero.MaxHealth {
					fmt.Println("You are at max health. Here is your potion back.")
					hero.Potions[i] = 2
					break
				}

			}
			fmt.Println("We shall move foreward!")
		}
		return true
	} else if answer != "no" {
		fmt.Println("That is not a valid choice")
		goto Prompt
	}
	return false

}

// Prints out hero stats
func heroStats(hero *Hero) {
	fmt.Println()
	fmt.Println("Your health is", hero.Health)
	fmt.Println("Your attack power is:", hero.AttackPower)
	fmt.Println("Your defense power is:", hero.DefensePower)
	fmt.Println("The amount of gold you have is:", hero.Gold)
	fmt.Println()

}

// Counts potions
func (hero *Hero) potionCount() int {
	potionCount := 0
	for i := 0; i < 5; i++ {
		if hero.Potions[i] == 2 {
			potionCount++
		}
	}
	return potionCount
}

// Shows how wounded the Hero is after battle
func postBattleStats(hero *Hero) {

	potions := strconv.Itoa(hero.potionCount())
	fmt.Println()
	fmt.Println("Your health is", hero.Health)
	fmt.Println("Your attack power is:", hero.AttackPower)
	fmt.Println("Your defense power is:", hero.DefensePower)
	fmt.Println("The amount of gold you have is:", hero.Gold)
	fmt.Println("You have", potions, "potions left")
	fmt.Println()

}

// Plays game
func gamePlay(hero *Hero) (bool, *Hero) {

	isOver := true
	var response string
	var steps int
	for isOver == true {
		isValid := false
		for isValid != true {

			fmt.Print("Would you like to take a step or take a potion? ")
			fmt.Scan(&response)
			response = strings.ToLower(response)
			if response == "potion" {
				isValid = takePotion(hero)
				steps += 1
			} else if response == "step" {
				steps += 1
				isValid = true
			} else {
				fmt.Println("That is not a choice bro")
				isValid = false
			}
		}
		encounter := encounter(hero)
		if encounter == true {

			fmt.Println("You have encountered a goblin! ")
			win := combat(hero)
			if win == true {
				fmt.Println("Congrats on your victory Goblin slayer!")
				hero.Gold += 2
				hero.GoblinKills += 1
				postBattleStats(hero)
			} else if win == false {
				fmt.Println("Despite your valiant effort you have been defeated :(")
				isOver = false
			}
		} else {
			fmt.Println("You have avoid the enemy! ")
		}

		if steps%10 == 0 && steps != 0 {
			levelUp(hero)
		}
	}
	if isOver == false {
		return false, hero
	}
	return true, hero
}

// Randomizes things
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Puts everything together
func main() {

	playAgain := true
	rounds := 0
	totalGold := 0
	hero := new(Hero)
	for playAgain != false {
		if rounds == 0 {
			hero.assembleHero()
			heroStats(hero)
		} else {
			hero.assembleHero()
			hero.Gold = totalGold
			heroStats(hero)
		}

		play, hero := gamePlay(hero)
		if play == false {
			var response string
			fmt.Print("Would you like to play again? ")
			fmt.Scan(&response)
			response = strings.ToLower(response)
			if response == "no" {
				level := strconv.Itoa(hero.Level)
				kills := strconv.Itoa(hero.GoblinKills)
				fmt.Println("You got up to level " + level + " you killed a total of " + kills + " Goblins")
				playAgain = false
			} else {
				totalGold = hero.Gold
				rounds++
			}
		}

	}

}
