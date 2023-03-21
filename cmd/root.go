/*
This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/
package cmd

import (
	"crypto/rand"
	crand "crypto/rand"
	"fmt"
	"math/big"
	mrand "math/rand"
	"regexp"
	"sort"
	"time"

	"os"

	"github.com/spf13/cobra"
)

// Holders for flag data
var swName string
var rollType string

var nonAlphaNumericRegex = regexp.MustCompile(`[^\p{L}\p{N}\.\/\'\- ]+`)

// Static Game Data

var randoAbilityMap = [][]uint8{
	{2, 1, 0},
	{2, 0, 1},
	{1, 2, 0},
	{0, 2, 1},
	{1, 0, 2},
	{0, 1, 2},
}

var itemMap = []string{
	"Staff",
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
	"Trailblaze",
	"Sweet Jamz",
	"High Times",
	"Gleam the Cube",
}

var skateTrickPropertyMap = [][]string{
	{"Steady", "Floating", "Rotated", "Inverted", "Greasy", "Upside", "Greyside", "Purring", "Barky", "Soggy", "Reversed", "Fuzzy", "Hanging", "Salty", "Thirsty", "Curvy", "Parallax", "Highlow", "Side", "Crunchy"},
	{"Roll", "Puff", "Skizzle", "Nutsome", "Chafer", "Yanker", "Burner", "Monger", "Dilly", "Jortle", "Crunchy", "Thwappy", "Snappy", "Clapper", "Stoney", "Punted", "Wiggle", "Shooketh", "Pendulus", "Doppler"},
	{"Frazz", "Burner", "Dump", "Chunt", "Breff", "Plop", "Nay", "Whizz", "Knob", "Goat", "Bruh", "Flex", "Cap", "Tope", "Smash", "Plug", "Mofo", "Seam", "Blaze", "Streak"},
	{"Grumble", "Stunted", "Lazy", "Glitchy", "Wrinkle", "Rising", "Saucy", "Flabby", "Wumpus", "Fuddy", "Lickety", "Muffin", "Whirly", "Squeegee", "Wishy", "Diddle", "Fripper", "Wishy", "Saluting", "Blubber"},
	{"Twist", "Pop", "Mast", "Drip", "Hole", "Melter", "Smack", "Schmooze", "Dollop", "Boo", "Daddle", "Pronk", "Hog", "Wink", "Spleen", "Fop", "Roll", "Bend", "Choke", "Dingle"},
}

var randoSpellMap = [][]string{
	{"Awesome", "Busted", "Dope", "Epic", "Gnarly", "Hyped", "Janky", "Killer", "Psyched", "Rad", "Sketchy", "Stoked"},
	{"Animating", "Attracting", "Bewildering", "Concealing", "Consuming", "Crushing", "Duplicating", "Expanding", "Revealing", "Sealing", "Shielding", "Summoning"},
	{"Acid", "Air", "Dust", "Earth", "Fire", "Light", "Reflecting", "Shadow", "Smoke", "Sound", "Spirit", "Water"},
	{"Armor", "Boot", "Bread", "Bucket", "Chain", "Door", "Hammer", "Lute", "Mattress", "Tower", "Tree", "Well"},
}

// Randomization Methods

func shuffleStringSlice(slice []string) {
	mrand.Seed(time.Now().UnixMicro())
	mrand.Shuffle(len(slice), func(i int, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

func shuffleAbilities(abilities [][]uint8) {
	mrand.Seed(time.Now().UnixMicro())
	mrand.Shuffle(len(abilities), func(i int, j int) {
		abilities[i], abilities[j] = abilities[j], abilities[i]
	})
}

func skateTrickProperty(row int) string {
	return skateTrickPropertyMap[row][0]
}

func randoSkateTrick() string {
	return fmt.Sprintf("%s %s %s with a %s %s", skateTrickProperty(0), skateTrickProperty(1), skateTrickProperty(2), skateTrickProperty(3), skateTrickProperty(4))
}

func randoSpellProperty(row int) string {
	return randoSpellMap[row][0]
}

func randoSpell() string {
	return fmt.Sprintf("%s %s %s %s", randoSpellProperty(0), randoSpellProperty(1), randoSpellProperty(2), randoSpellProperty(3))
}

func randoBootleg() string {
	return bootlegMap[0]
}

func randoAbilities() []uint8 {
	return randoAbilityMap[0]
}

func randoItems() []string {
	return itemMap[:4]
}

func gameRoll(dice int, modifier string) int {
	max := big.NewInt(6)
	rolledDice := make([]int, 3)
	for i := range rolledDice {
		randInt, err := crand.Int(rand.Reader, max)
		if err != nil {
			fmt.Println("sw-cli encountered a critical error generating random numbers: ", err)
			os.Exit(2)
		}
		rolledDice[i] = int(randInt.Uint64()) + 1
	}
	// Sort the list
	switch modifier {
	case "disadvantage":
		sort.Ints(rolledDice)
	case "advantage":
		sort.Sort(sort.Reverse(sort.IntSlice(rolledDice)))
	default:
		break
	}
	// Sum the number of dice needed.
	rollTotal := 0
	for i := 0; i < dice; i++ {
		rollTotal += rolledDice[i]
	}
	return rollTotal
}

// COMMANDS

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sw-cli",
	Short: "Doesn't do anything yet!",
	Long: `Doesn't do anything yet, but may in the future. If it does
	that'll probably be launching a game loop for interacting with a
	character.
	
	So until then, just run one of the following commands:
	rando-trick - get a random skate trick.
	rando-spell - get a random spell.
	rando-character - generate a random character`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var randoSkateTrickCmd = &cobra.Command{
	Use:   "rando-trick",
	Short: "Generates a random skate trick name",
	Long:  "Generates a skate trick name based on the table available in the standard rule book",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Your skate trick name is: %s\n", randoSkateTrick())
	},
}

var randoSpellCmd = &cobra.Command{
	Use:   "rando-spell",
	Short: "Generates a random rando spell",
	Long:  "Generates a rando spell based on the table available in the standard rule book",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Your rando-spell is %s\n", randoSpell())
	},
}

var randoDiceCmd = &cobra.Command{
	Use:   "rando-roll",
	Short: "Rolls some dice",
	Long:  "Rolls a number of dice",
	Run: func(cmd *cobra.Command, args []string) {
		switch rollType {
		case "initiative":
			fmt.Printf("Initiatve roll: %d\n", gameRoll(1, "none"))
		case "advantage":
			fmt.Printf("Advantaged roll: %d\n", gameRoll(2, "advantage"))
		case "disadvantage":
			fmt.Printf("Disadvantaged roll: %d\n", gameRoll(2, "disadvantage"))
		default:
			fmt.Printf("Just a roll %d\n", gameRoll(2, "none"))
		}

	},
}

var randoCharacterCmd = &cobra.Command{
	Use:   "rando-character",
	Short: "Generates a random character",
	Long:  "Generates a rando character in the standard rule book",
	Run: func(cmd *cobra.Command, args []string) {
		cleanedName := nonAlphaNumericRegex.ReplaceAllString(swName, "")
		randoStats := randoAbilities()
		randoItems := randoItems()
		fmt.Printf("    ▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄      \n")
		fmt.Printf("   ▐      ╔══════════SKATE WIZARD═══════╗  ╔════════════════BOOTLEG SPELL══════════════╗  ▌   \n")
		fmt.Printf("  ▐       ║ %-27s ║  ║ %-41s ║   ▌  \n", cleanedName, randoBootleg())
		fmt.Printf(" ▐        ╚═════════════════════════════╝  ╚═══════════════════════════════════════════╝    ▌ \n")
		fmt.Printf(" ▐        ╔═LVL═╗         ╔═EXP═╗          ╔════════════════════ITEMS══════════════════╗    ▌ \n")
		fmt.Printf("▐         ║  1  ║         ║  0  ║          ║ 1. %-38s ║     ▌\n", randoItems[0])
		fmt.Printf("▐         ╚═════╝         ╚═════╝          ║ 2. %-38s ║     ▌\n", randoItems[1])
		fmt.Printf("▐         ╔═HIT═╗         ╔═DEF═╗          ║ 3. %-38s ║     ▌\n", randoItems[2])
		fmt.Printf("▐         ║  4  ║         ║  6  ║          ║ 4. %-38s ║     ▌\n", randoItems[3])
		fmt.Printf("▐         ╚═════╝         ╚═════╝          ╚═══════════════════════════════════════════╝     ▌\n")
		fmt.Printf("▐         ╔═STR═╗ ╔═DEX═╗ ╔WILL═╗          ╔════════════════RANDO SPELL════════════════╗     ▌\n")
		fmt.Printf("▐         ║  %d  ║ ║  %d  ║ ║  %d  ║          ║ %-41s ║     ▌\n", randoStats[0], randoStats[1], randoStats[2], randoSpell())
		fmt.Printf(" ▐        ╚═════╝ ╚═════╝ ╚═════╝          ╚═══════════════════════════════════════════╝    ▌ \n")
		fmt.Printf(" ▐                                     ╔════════════════SIGNATURE TRICK════════════════╗    ▌ \n")
		fmt.Printf("  ▐                                    ║ %-45s ║   ▌  \n", randoSkateTrick())
		fmt.Printf("   ▐                                   ╚═══════════════════════════════════════════════╝  ▌   \n")
		fmt.Printf("    ▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀     \n")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Printf("sw-cli has suffered a critical error. \n %s", err.Error())
		os.Exit(1)
	}
}

func init() {

	// Initialize data sources
	shuffleStringSlice(itemMap)
	shuffleStringSlice(bootlegMap)
	shuffleAbilities(randoAbilityMap)
	for i := range skateTrickPropertyMap {
		shuffleStringSlice(skateTrickPropertyMap[i])
	}
	for i := range randoSpellMap {
		shuffleStringSlice(randoSpellMap[i])
	}

	// Add all commands
	randoCharacterCmd.Flags().StringVarP(&swName, "name", "n", "Steve", "a non-random name for a random Skate Wizard")
	rootCmd.AddCommand(randoCharacterCmd)

	rootCmd.AddCommand(randoSkateTrickCmd)
	rootCmd.AddCommand(randoSpellCmd)

	randoDiceCmd.Flags().StringVarP(&rollType, "roll", "r", "danger", "the type of random roll to make")
	rootCmd.AddCommand(randoDiceCmd)
}
