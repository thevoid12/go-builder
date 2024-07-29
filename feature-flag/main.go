package main

import (
	"flag"
	"fmt"
)

// command line flag parsing: https://docs.google.com/document/d/1gjXiQMmRFNJqMlIY8yUIUMiq0GnLfPuOIRhfbLNcqAc/edit?usp=sharing
func main() {
	//format: actual flag,default value if proper value is not entered for this flag, flag definition.
	// The reason why we give default value because we can use it in help mode(./abc.exe -h)
	testflag := flag.String("test", "default test msg", "Enter a string test message")
	testintflag := flag.Int("intflag", 5, "Enter the number of times the test msg flag should run")
	flag.Parse() // this should always be in the last line and before starting to use the value

	//we can start using the flags. if no value entered for testinflag it takes default 5 as we are using it here in range query
	for i := range *testintflag {
		fmt.Println(i, *testflag)
	}
}
