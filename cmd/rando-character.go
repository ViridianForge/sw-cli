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
	ShuffleStringSlice(itemMap)
	ShuffleStringSlice(bootlegMap)
	ShuffleAbilities(randoAbilityMap)
	RootCmd.AddCommand(randoCharacterCmd)
}

func ShuffleStringSlice(slice []string) {
	rand.Seed(time.Now().UnixMicro())
	rand.Shuffle(len(slice), func(i int, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

func ShuffleAbilities(abilities [][]uint8) {
	rand.Seed(time.Now().UnixMicro())
	rand.Shuffle(len(abilities), func(i int, j int) {
		abilities[i], abilities[j] = abilities[j], abilities[i]
	})
}

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
	"Traiblaze",
	"Sweet Jamz",
	"High Times",
	"Gleam the Cube",
}

func randoAbilities() []uint8 {
	return randoAbilityMap[0]
}

func randoItems() []string {
	return itemMap[:4]
}

func randoBootleg() string {
	return bootlegMap[0]
}

var randoCharacterCmd = &cobra.Command{
	Use:   "rando-character",
	Short: "Generates a random character",
	Long:  "Generates a rando character in the standard rule book",
	Run: func(cmd *cobra.Command, args []string) {
		randoStats := randoAbilities()
		randoItems := randoItems()
		fmt.Printf("    ▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄      \n")
		fmt.Printf("   ▐      ╔══════════SKATE WIZARD═══════╗  ╔════════════════BOOTLEG SPELL══════════════╗  ▌   \n")
		fmt.Printf("  ▐       ║ PLACEHOLDER                 ║  ║ %-41s ║   ▌  \n", randoBootleg())
		fmt.Printf(" ▐        ╚═════════════════════════════╝  ╚═══════════════════════════════════════════╝    ▌ \n")
		fmt.Printf(" ▐        ╔═LVL═╗         ╔═EXP═╗          ╔════════════════════ITEMS══════════════════╗    ▌ \n")
		fmt.Printf("▐         ║  1  ║         ║  0  ║          ║ 1. %-38s ║     ▌\n", randoItems[0])
		fmt.Printf("▐         ╚═════╝         ╚═════╝          ║ 2. %-38s ║     ▌\n", randoItems[1])
		fmt.Printf("▐         ╔═HIT═╗         ╔═DEF═╗          ║ 3. %-38s ║     ▌\n", randoItems[2])
		fmt.Printf("▐         ║  4  ║         ║  6  ║          ║ 4. %-38s ║     ▌\n", randoItems[3])
		fmt.Printf("▐         ╚═════╝         ╚═════╝          ╚═══════════════════════════════════════════╝     ▌\n")
		fmt.Printf("▐         ╔═STR═╗ ╔═DEX═╗ ╔WILL═╗          ╔════════════════RANDO SPELL════════════════╗     ▌\n")
		fmt.Printf("▐         ║  %d  ║ ║  %d  ║ ║  %d  ║          ║ %-41s ║     ▌\n", randoStats[0], randoStats[1], randoStats[2], RandoSpell())
		fmt.Printf(" ▐        ╚═════╝ ╚═════╝ ╚═════╝          ╚═══════════════════════════════════════════╝    ▌ \n")
		fmt.Printf(" ▐                                     ╔════════════════SIGNATURE TRICK════════════════╗    ▌ \n")
		fmt.Printf("  ▐                                    ║ %-45s ║   ▌  \n", RandoSkateTrick())
		fmt.Printf("   ▐                                   ╚═══════════════════════════════════════════════╝  ▌   \n")
		fmt.Printf("    ▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀     \n")
	},
}
