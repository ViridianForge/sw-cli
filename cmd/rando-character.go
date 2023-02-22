/*
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/
package cmd

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(randoCharacterCmd)
}

/*var randoSpellMap = [][]string{
	{"Awesome", "Busted", "Dope", "Epic", "Gnarly", "Hyped", "Janky", "Killer", "Psyched", "Rad", "Sketchy", "Stoked"},
	{"Animating", "Attracting", "Bewildering", "Concealing", "Consuming", "Crushing", "Duplicating", "Expanding", "Revealing", "Sealing", "Shielding", "Summoning"},
	{"Acid", "Air", "Dust", "Earth", "Fire", "Light", "Reflecting", "Shadow", "Smoke", "Sound", "Spirit", "Water"},
	{"Armor", "Boot", "Bread", "Bucket", "Chain", "Door", "Hammer", "Lute", "Mattress", "Tower", "Tree", "Well"},
}*/

var randoAbilityMap = [][]uint8{
	{2, 1, 0},
	{2, 0, 1},
	{1, 2, 0},
	{0, 2, 1},
	{1, 0, 2},
	{0, 1, 2},
}

var itemMap = []string{
	"Staff or wand",
	"Skate Key",
	"Smoking Pipe",
	"Shield",
	"Melee Weapon",
	"Ranged Weapon",
	"Lore from Yore",
	"Lockpicking Tools",
	"Oil Flask",
	"Rope",
	"Torch and Tinder",
	"Ball Bearing",
}

var bootlegMap = []string{
	"Mattress",
	"Wannabe",
	"Traiblaze",
	"Sweet Jamz",
	"High Times",
	"Gleam the Cube",
}

func randoAbilities() []uint8 {
	col, err := rand.Int(rand.Reader, big.NewInt(5))
	if err != nil {
		panic(fmt.Sprintf("randoAbilities critically failed during random number generation with error %v", err))
	}
	return randoAbilityMap[col.Int64()+1]
}

func randoItems() []string {
	items := make([]string, 4)
	for i := range items {
		item, err := rand.Int(rand.Reader, big.NewInt(5))
		if err != nil {
			panic(fmt.Sprintf("randoItems critically failed during random number generation with error %v", err))
		}
		items[i] = (itemMap[item.Int64()+1])
	}
	return items
}

func randoBootleg() string {
	col, err := rand.Int(rand.Reader, big.NewInt(5))
	if err != nil {
		panic(fmt.Sprintf("randoBootleg critically failed during random number generation with error %v", err))
	}
	return bootlegMap[col.Int64()+1]
}

var randoCharacterCmd = &cobra.Command{
	Use:   "rando-character",
	Short: "Generates a random character",
	Long:  "Generates a rando character in the standard rule book",
	Run: func(cmd *cobra.Command, args []string) {
		randoStats := randoAbilities()
		randoItems := randoItems()
		fmt.Printf("Your ability modifiers are is Strength: +%d, Dexterity: +%d, Will: +%d\n", randoStats[0], randoStats[1], randoStats[2])
		fmt.Print("You have 4 Health\n")
		fmt.Print("You have 6 Defense\n")
		fmt.Print("Your permanent spells are:\n")
		fmt.Print("\t1. Ramp\n")
		fmt.Print("\t1. Sidewalk\n")
		fmt.Print("\t1. Rail\n")
		fmt.Printf("Your rando spell is: %s\n", RandoSpell())
		fmt.Printf("Your bootleg spell is: %s\n", randoBootleg())
		fmt.Printf("Your starting items are:\n")
		fmt.Printf("\t1. %s\n", randoItems[0])
		fmt.Printf("\t2. %s\n", randoItems[1])
		fmt.Printf("\t3. %s\n", randoItems[2])
		fmt.Printf("\t4. %s\n", randoItems[3])
		fmt.Printf("Your signature skate trick is: %s\n", RandoSkateTrick())
	},
}
