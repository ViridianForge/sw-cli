package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

var skateTrickPropertyMap = [][]string{
	{"Steady", "Floating", "Rotated", "Inverted", "Greasy", "Upside", "Greyside", "Purring", "Barky", "Soggy", "Reversed", "Fuzzy", "Hanging", "Salty", "Thirsty", "Curvy", "Parallax", "Highlow", "Side", "Crunchy"},
	{"Roll", "Puff", "Skizzle", "Nutsome", "Chafer", "Yanker", "Burner", "Monger", "Dilly", "Jortle", "Crunchy", "Thwappy", "Snappy", "Clapper", "Stoney", "Punted", "Wiggle", "Shooketh", "Pendulus", "Doppler"},
	{"Frazz", "Burner", "Dump", "Chunt", "Breff", "Plop", "Nay", "Whizz", "Knob", "Goat", "Bruh", "Flex", "Cap", "Tope", "Smash", "Plug", "Mofo", "Seam", "Blaze", "Streak"},
	{"Grumble", "Stunted", "Lazy", "Glitchy", "Wrinkle", "Rising", "Saucy", "Flabby", "Wumpus", "Fuddy", "Lickety", "Muffin", "Whirly", "Squeegee", "Wishy", "Diddle", "Fripper", "Wishy", "Saluting", "Blubber"},
	{"Twist", "Pop", "Mast", "Drip", "Hole", "Melter", "Smack", "Schmooze", "Dollop", "Boo", "Daddle", "Pronk", "Hog", "Wink", "Spleen", "Fop", "Roll", "Bend", "Choke", "Dingle"},
}

func init() {
	rand.Seed(time.Now().UnixNano())
	rootCmd.AddCommand(randoSkateTrickCmd)
}

func skateTrickProperty(row int) string {
	return skateTrickPropertyMap[row][rand.Intn(19)+1]
}

func randoSkateTrick() string {
	return fmt.Sprintf("%s %s %s with a %s %s", skateTrickProperty(0), skateTrickProperty(1), skateTrickProperty(2), skateTrickProperty(3), skateTrickProperty(4))
}

var randoSkateTrickCmd = &cobra.Command{
	Use:   "rando-trick",
	Short: "Generates a random skate trick name",
	Long:  "Generates a skate trick name based on the table available in the standard rule book",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Your skate trick name is: %s\n", randoSkateTrick())
	},
}
