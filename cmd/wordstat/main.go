package main

import (
	"fmt"
	"log"

	"github.com/kotmyshev/wordstat/internal/app/wsapp"
	"github.com/kotmyshev/wordstat/pkg/filestat"
	"github.com/kotmyshev/wordstat/pkg/rndstr"
)

const (
	// separators for split text to words (for filestat pkg)
	separators = " .,:[]()!?\"\n\t\r"
	// length of names for new files
	fnamelen = 10
	// randomize seed additional number
	rndseed = 42
)

func main() {

	fmt.Print("Enter URL: ")
	inurl, err := wsapp.ReadLineFromStdIn()

	if err != nil {
		log.Fatal(err)
	}

	rndwrd := rndstr.GenerateRandomString(rndseed, fnamelen)

	err = wsapp.MakeFilesFromURL(inurl, rndwrd)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rndwrd)

	stat, err := filestat.WordStatFromFile(rndwrd+".txt", separators)

	wsapp.PrintStatInStdOut(stat)

	if err != nil {
		log.Fatal(err)
	}
}
