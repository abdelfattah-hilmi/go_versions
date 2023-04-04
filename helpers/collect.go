package main

import (
	"fmt"
	"strings"

	// importing Colly
	"github.com/gocolly/colly"
)

func main() {

	c := colly.NewCollector()

	var url = "https://repology.org/project/mongodb/information"

	c.OnHTML("span.version-newest", func(e *colly.HTMLElement) {
		// printing all URLs associated with the a links in the page
		fmt.Println("lates version:", e.Text)
	})

	c.Visit(url)

	url = "https://repology.org/project/mongodb/cves"

	c.OnHTML("tbody tr", func(e *colly.HTMLElement) {
		cve_id := e.ChildText("td.minimal-column a")
		cve_link := e.ChildAttr("td.minimal-column a", "href")

		affected_versions := e.ChildText("td.text-left")

		v := "1.2.0"
		if strings.Contains(affected_versions, v) {
			fmt.Println(cve_id, "->", cve_link, "affected:", affected_versions)
		}
	})

	c.Visit(url)
}
