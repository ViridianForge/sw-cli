/*
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/
package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	RootCmd.AddCommand(randoSpellCmd)
}

var randoSpellMap = [][]string{
	{"Awesome", "Busted", "Dope", "Epic", "Gnarly", "Hyped", "Janky", "Killer", "Psyched", "Rad", "Sketchy", "Stoked"},
	{"Animating", "Attracting", "Bewildering", "Concealing", "Consuming", "Crushing", "Duplicating", "Expanding", "Revealing", "Sealing", "Shielding", "Summoning"},
	{"Acid", "Air", "Dust", "Earth", "Fire", "Light", "Reflecting", "Shadow", "Smoke", "Sound", "Spirit", "Water"},
	{"Armor", "Boot", "Bread", "Bucket", "Chain", "Door", "Hammer", "Lute", "Mattress", "Tower", "Tree", "Well"},
}

func randoSpellProperty(row int) string {
	return randoSpellMap[row][rand.Intn(11)+1]
}

func RandoSpell() string {
	return fmt.Sprintf("%s %s %s %s", randoSpellProperty(0), randoSpellProperty(1), randoSpellProperty(2), randoSpellProperty(3))
}

var randoSpellCmd = &cobra.Command{
	Use:   "rando-spell",
	Short: "Generates a random rando spell",
	Long:  "Generates a rando spell based on the table available in the standard rule book",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Your rando-spell is %s\n", RandoSpell())
	},
}
