package helpers

import (
	"example/go_versions/models"
	"fmt"

	"github.com/gocolly/colly"
)

func GetReleaseNotes(pkg string) []string {
	var base_url = "https://www.wikidata.org"
	var url = base_url + "/w/index.php?search=" + pkg

	c := colly.NewCollector()

	ids := []string{}

	c.OnHTML("ul.mw-search-results", func(h *colly.HTMLElement) {
		ids = append(ids, h.ChildAttr("li div a", "href"))
	})
	c.Visit(url)

	id := ids[0]

	url = base_url + id

	Link := []string{}

	c.OnHTML("#P348 .wb-preferred .external", func(h *colly.HTMLElement) {
		Link = append(Link, h.Text)
	})
	c.Visit(url)

	return Link
}

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

	return repo
}

func PkgIsinUbnutuData(distro_name string, pkg string) bool {
	url := "https://packages.ubuntu.com/" + distro_name + "/" + pkg

	c := colly.NewCollector()

	value := []string{}
	c.OnHTML("#content h1", func(h *colly.HTMLElement) {
		value = append(value, h.Text)
	})

	c.Visit(url)
	if value[0] == "Error" {
		fmt.Println(false)
		return false
	} else {
		fmt.Println(true)
		return true
	}
}
