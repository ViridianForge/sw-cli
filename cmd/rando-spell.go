package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(randoSpellCmd)
}

var propertyMap = [][]string{
	{"Awesome", "Busted", "Dope", "Epic", "Gnarly", "Hyped", "Janky", "Killer", "Psyched", "Rad", "Sketchy", "Stoked"},
	{"Animating", "Attracting", "Bewildering", "Concealing", "Consuming", "Crushing", "Duplicating", "Expanding", "Revealing", "Sealing", "Shielding", "Summoning"},
	{"Acid", "Air", "Dust", "Earth", "Fire", "Light", "Reflecting", "Shadow", "Smoke", "Sound", "Spirit", "Water"},
	{"Armor", "Boot", "Bread", "Bucket", "Chain", "Door", "Hammer", "Lute", "Mattress", "Tower", "Tree", "Well"},
}

func randomProperty(row int) string {
	rand.Seed(time.Now().UnixNano())
	return propertyMap[row][rand.Intn(11)+1]
}

var randoSpellCmd = &cobra.Command{
	Use:   "rando-spell",
	Short: "Generates a random rando spell",
	Long:  "Generates a rando spell based on the table available in the standard rule book",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Your rando-spell is %s %s %s %s", randomProperty(0), randomProperty(1), randomProperty(2), randomProperty(3))
	},
}
