package cmd

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/spf13/cobra"
)

var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Scrape the data from the imdb website",
	RunE: func(cmd *cobra.Command, args []string) error {
		l, err := scrapeActorsIdsFromMovieId("tt0109830")
		fmt.Println("Size: ", len(l))
		return err
	},
}

func init() {
	rootCmd.AddCommand(scrapeCmd)
}

// get all the actors ids from a movie id
// scrape the data from the imdb website
// use the colly library

func scrapeActorsIdsFromMovieId(id string) ([]string, error) {
	// use Colly to scrape the data

	c := colly.NewCollector()

	// create a slice of strings to store the actors ids
	actors := []string{}

	// find the actors ids in the html
	c.OnHTML("table.cast_list td.primary_photo a", func(e *colly.HTMLElement) {
		// get the href attribute
		// add the id to the slice
		actors = append(actors, e.Attr("href")[6:15])
		fmt.Println(e.ChildAttr("img", "title"))
		//fmt.Println(actors)
	})

	// visit the movie page
	// return the actors ids and an error
	c.Visit("https://www.imdb.com/title/" + id + "/fullcredits")

	return actors, nil
}
