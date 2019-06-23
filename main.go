package main

import (
	"log"
	"net/http"
	"scrape_recipes/scrape"

	"github.com/ant0ine/go-json-rest/rest"
)

// Recipes レシピのURL
type Recipes struct {
	Urls []string `json:"urls"`
}

func getRecipes(w rest.ResponseWriter, req *rest.Request) {
	recipes := Recipes{}
	err := req.DecodeJsonPayload(&recipes)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	scrapedRecipes := []scrape.ScrapedRecipe{}
	for _, url := range recipes.Urls {
		scrapedRecipe := scrape.ScrapeRecipe(url)
		scrapedRecipes = append(scrapedRecipes, scrapedRecipe)
	}
	w.WriteJson(&scrapedRecipes)
}

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	r, err := rest.MakeRouter(
		rest.Post("/recipes", getRecipes),
	)

	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(r)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
