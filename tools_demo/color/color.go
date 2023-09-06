package color

import (
	"fmt"
	"github.com/fatih/color"
)

// Cyan 可以用来输出关键信息
func Cyan() {
	color.Cyan("测试输出")
	color.Blue("Prints %s in blue.", "text")
	color.Red("We have red")
	color.Magenta("And many others ..")

	c := color.New(color.FgCyan).Add(color.Underline)
	c.Println("Prints cyan text with an underline.")

	// Create a custom print function for convenience
	red := color.New(color.FgRed).PrintfFunc()
	red("Warning")

	// Mix up multiple attributes
	notice := color.New(color.Bold, color.FgGreen).PrintlnFunc()
	notice("Don't forget this...")

	// Create SprintXxx functions to mix strings with other non-colorized strings:
	yellow := color.New(color.FgYellow).SprintFunc()
	red2 := color.New(color.FgRed).SprintFunc()
	fmt.Printf("This is a %s and this is %s.\n", yellow("warning"), red2("error"))

	info := color.New(color.FgWhite, color.BgGreen).SprintFunc()
	fmt.Printf("This %s rocks!\n", info("package"))

	// Use helper functions
	fmt.Println("This", color.RedString("warning"), "should be not neglected.")
	fmt.Printf("%v %v\n", color.GreenString("Info:"), "an important message.")

	// Windows supported too! Just don't forget to change the output to color.Output
	fmt.Fprintf(color.Output, "Windows support: %s", color.GreenString("PASS"))

	c = color.New(color.FgCyan)
	c.Println("Prints cyan text")

	c.DisableColor()
	c.Println("This is printed without any color")

	c.EnableColor()
	c.Println("This prints again cyan...")
}
