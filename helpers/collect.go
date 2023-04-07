package helpers

import (
	"fmt"

	"github.com/gocolly/colly"
)

func getReleaseNotes(pkg string) string {
	var base_url = "https://www.wikidata.org"
	var url = base_url + "/w/index.php?search=" + pkg

	c := colly.NewCollector()

	ids := []string{}

	c.OnHTML("ul.mw-search-results", func(h *colly.HTMLElement) {
		ids = append(ids, h.ChildAttr("li div a", "href"))
	})
	c.Visit(url)

	id := ids[0]
	fmt.Println(id)

	url = base_url + id

	link := []string{}

	c.OnHTML("#P348 .wb-preferred .external", func(h *colly.HTMLElement) {
		link = append(link, h.Text)
	})
	c.Visit(url)
	return link[0]
}

func getRepologyData(pkg string) {

	c := colly.NewCollector()

	var url = "https://repology.org/project/" + pkg + "/information"

	c.OnHTML("span.version-newest", func(e *colly.HTMLElement) {

		fmt.Println("lates version:", e.Text)
	})

	c.Visit(url)

	url = "https://repology.org/project/" + pkg + "/cves"

	c.OnHTML("tbody tr", func(e *colly.HTMLElement) {
		cve_id := e.ChildText("td.minimal-column a")
		cve_link := e.ChildAttr("td.minimal-column a", "href")

		affected_versions := e.ChildText("td.text-left")

		// v := "1.2.0"
		// if strings.Contains(affected_versions, v) {
		// 	fmt.Println(cve_id, "->", cve_link, "affected:", affected_versions)
		// }
		fmt.Println(cve_id, "->", cve_link, "affected:", affected_versions)
	})

	c.Visit(url)
}

func pkgIsinUbnutuData(distro_name string, pkg string) bool {
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
