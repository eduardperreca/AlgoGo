package main

import (
	"fmt"
	"log"
y"
)

func main() {
	doc, err := goquery.NewDocument("https://sell.wethenew.com/consignment/product/1106")
	if err != nil {
		log.Fatal(err)
	}

	// Estrarre il titolo del prodotto
	title := doc.Find("h1.title").Text()
	fmt.Println("Titolo:", title)

	// Estrarre il prezzo del prodotto
	price := doc.Find("span.price").Text()
	fmt.Println("Prezzo:", price)

	// Estrarre l'immagine del prodotto
	image, _ := doc.Find("div.product-img-zoom > img").Attr("src")
	fmt.Println("Immagine:", image)
}
