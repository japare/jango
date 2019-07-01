package jan

import (
	"bytes"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

// Search searches the website myfigurecollection.net for JANs matching the passed keywords
func Search(keywords []string) ([]string, error) {

	// construct the URI
	joined := strings.Join(keywords[:], "+")
	uri := "https://myfigurecollection.net/browse.v4.php?keywords=" + joined

	// get the document
	doc, err := goquery.NewDocument(uri)
	if err != nil {
		return []string{}, err
	}

	// use selectors to get mfc item ids
	matches := []string{}
	doc.Find(".tbx-tooltip").
		Each(func(i int, g *goquery.Selection) {
			link, exists := g.Attr("href")
			if !exists {
				//skip
			}
			matches = append(matches, link)
		})

	matches = filter(matches, func(s string) bool {
		return strings.Contains(s, "item")
	})
	log.Println("matches:", matches)

	// look up the JAN for each item ID
	uri = "https://myfigurecollection.net"
	// results := make(chan string, len(matches))
	resultchan := make(chan string, 1000)

	re := regexp.MustCompile(`meta itemprop="productID" content="jan:(\d+)`)
	var wg sync.WaitGroup

	for i, itemID := range matches {
		wg.Add(1)
		go func(itemID string, baseuri string, re *regexp.Regexp, results chan string, waitgroup *sync.WaitGroup, id int) {
			defer waitgroup.Done()

			response, err := http.Get(baseuri + itemID)
			if err != nil {
				return
			}
			buf := new(bytes.Buffer)
			buf.ReadFrom(response.Body)
			body := buf.String()
			match := re.FindStringSubmatch(body)
			// We want the element inside the 1st capturing group
			// If it has a different length, something broke and this is probably useless to us
			if len(match) != 2 {
				return
			}
			jan := match[1]
			results <- jan
			return
		}(itemID, uri, re, resultchan, &wg, i)
	}

	jans := make([]string, len(matches))
	wg.Wait()
	close(resultchan)

	i := 0
	for jan := range resultchan {
		jans[i] = jan
		i++
	}
	return jans, nil
}

// returns all elements of `strings` that satisfy function `f`
func filter(strings []string, f func(string) bool) []string {
	satisfied := make([]string, 0)
	for _, v := range strings {
		if f(v) {
			satisfied = append(satisfied, v)
		}
	}
	return satisfied
}
