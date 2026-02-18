package main

import (
	"fmt"
	"os"

	"github.com/cocodrilette/researchdiary/infra"
	"github.com/cocodrilette/researchdiary/models"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	db, err := infra.Database.Connect2()
	if err != nil {
		fmt.Fprintf(os.Stderr, "database connection failed: %v\n", err)
		os.Exit(1)
	}

	db.AutoMigrate(&models.Article{}, &models.Author{})

	articleManager := models.ArticleManager{DB: db}
	result, _ := articleManager.Find("")

	fmt.Println(result)

	// article, err := articleManager.NewFromTerminal(os.Stdin)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "cannot create article from terminal: %v\n", err)
	// }

	// err = articleManager.Create(&article)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "cannot create in db: %v\n", err)
	// }

	// fmt.Println(article)

}
