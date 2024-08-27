package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	deeplapi "github.com/leschuster/deepl-cli/pkg/deepl-api"
)

func main() {
	fmt.Println("Welcome to the inofficial DeepL-CLI tool!")

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf(".env file not found")
		os.Exit(2)
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Printf("You need to provide an API_KEY env variable.")
		os.Exit(2)
	}

	api := deeplapi.New(apiKey)

	opt := deeplapi.TranslateParams{
		Text:       []string{"Sets whether the translation engine should respect the original formatting, even if it would usually correct some aspects."},
		SourceLang: "EN",
		TargetLang: "FR",
	}

	res, err := api.Translate(opt)
	// res, err := api.GetLanguages()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	}

	fmt.Printf("%v", res)
}
