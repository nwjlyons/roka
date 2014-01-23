package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

type Kana struct {
	Hiragana string
	Katakana string
}

var romaji = make(map[string]Kana)

var (
	flagHiragana = flag.Bool("h", false, "")
	flagKatakana = flag.Bool("k", false, "")
)

var usage = `Usage: roka [options...] <romaji character>

Options:
	-h Display Hiragana.
	-k Display Katakana.
`

func main() {

	file, err := os.Open("kana.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	for {
		fields, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		romaji[fields[0]] = Kana{Hiragana: fields[1], Katakana: fields[2]}
	}

	flag.Parse()

	// Check command line argument is present
	if flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "Missing romaji character.\n\n")
		fmt.Fprintf(os.Stderr, usage+"\n")
		os.Exit(1)
	}

	character := flag.Args()[0]

	// Lookup kana from romaji.
	if kana, ok := romaji[character]; ok {
		if *flagHiragana {
			fmt.Println("Hiragana:", kana.Hiragana)
		}
		if *flagKatakana {
			fmt.Println("Katakana:", kana.Katakana)
		}
	} else {
		fmt.Println("Kana doesn't exist.")
	}

}
