package main

import "github.com/pspiagicw/pelp"

func main() {
	pelp.Print("This is a message")
	pelp.HeaderWithDescription("Example Header", []string{
		"This is description.",
	})
	pelp.Aligned("Aligned Header", []string{"left1", "left2"}, []string{"right1", "right2"})

	pelp.Flags("Flags Header", []string{"flag1", "flag2"}, []string{"description1", "description2"})
	pelp.Examples("Examples Header", []string{"example1", "exampl2 with args"})
}
