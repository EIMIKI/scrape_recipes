package scrape

import "github.com/PuerkitoBio/goquery"

func scrapeAjinomoto(url string) ScrapedRecipe {
	doc, _ := goquery.NewDocument(url)

	scrapedRecipe := ScrapedRecipe{}
	scrapedRecipe.Title = doc.Find("h1").Text()

	ingredientSelection := doc.Find(".recipeMaterialList>dl")
	ingredient := Ingredient{}
	ingredientSelection.Each(func(index int, s *goquery.Selection) {
		ingredient.Name = s.Find("dt").Text()
		ingredient.Amount = s.Find("dd").Text()
		scrapedRecipe.Ingredients = append(scrapedRecipe.Ingredients, ingredient)
	})

	directionSelection := doc.Find(".inGallery")
	direction := Direction{}
	directionSelection.Each(func(index int, s *goquery.Selection) {
		if s.Find("h3").Text() != "" {
			direction.Position = s.Find("h3").Text()
			direction.Text = s.Find(".txt").Text()
			scrapedRecipe.Directions = append(scrapedRecipe.Directions, direction)
		}
	})

	scrapedRecipe.Err = ""

	return scrapedRecipe
}
