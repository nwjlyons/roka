package main

import (
	"flag"
	"fmt"
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

	// Vowels
	romaji["a"] = Kana{Hiragana: "あ", Katakana: "ア"}
	romaji["i"] = Kana{Hiragana: "い"}
	romaji["u"] = Kana{Hiragana: "う"}
	romaji["e"] = Kana{Hiragana: "え"}
	romaji["o"] = Kana{Hiragana: "お"}

	// K row
	romaji["ka"] = Kana{Hiragana: "か"}
	romaji["ki"] = Kana{Hiragana: "き"}
	romaji["ku"] = Kana{Hiragana: "く"}
	romaji["ke"] = Kana{Hiragana: "け"}
	romaji["ko"] = Kana{Hiragana: "こ"}

	// Dakuten
	// G row
	romaji["ga"] = Kana{Hiragana: "が"}
	romaji["gi"] = Kana{Hiragana: "ぎ"}
	romaji["gu"] = Kana{Hiragana: "ぐ"}
	romaji["ge"] = Kana{Hiragana: "げ"}
	romaji["go"] = Kana{Hiragana: "ご"}

	// Yoon
	// K row
	romaji["kya"] = Kana{Hiragana: "かゃ"}
	romaji["kyu"] = Kana{Hiragana: "くゃ"}
	romaji["kyo"] = Kana{Hiragana: "こゃ"}

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
