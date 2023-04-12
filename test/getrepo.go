package main

import (
	"example/go_versions/models"
	"fmt"

	"github.com/gocolly/colly"
)

func GetRepologyData(pkg string) models.Repology {

	repo := models.Repology{}
	cve := models.Cve{}

	c := colly.NewCollector()

	var url = "https://repology.org/project/" + pkg + "/information"

	c.OnHTML("span.version-newest", func(e *colly.HTMLElement) {
		repo.LatestVersion = e.Text
		// fmt.Println("lates version:", e.Text)
	})

	c.Visit(url)

	url = "https://repology.org/project/" + pkg + "/cves"

	c.OnHTML("tbody tr", func(e *colly.HTMLElement) {
		// cve_id := e.ChildText("td.minimal-column a")
		// cve_link := e.ChildAttr("td.minimal-column a", "href")

		cve.CveId = e.ChildText("td.minimal-column a")
		cve.CveLink = e.ChildAttr("td.minimal-column a", "href")
		cve.AffectedVersions = e.ChildText("td.text-left")
		repo.Cves = append(repo.Cves, cve)
		// v := "1.2.0"
		// if strings.Contains(affected_versions, v) {
		// 	fmt.Println(cve_id, "->", cve_link, "affected:", affected_versions)
		// }
		// fmt.Println(cve_id, "->", cve_link, "affected:", affected_versions)
	})
	c.Visit(url)

	fmt.Println(repo.Cves)
	return repo
}

func main() {
	GetRepologyData("nginx")
}
