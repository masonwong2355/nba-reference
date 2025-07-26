package scraper

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getPageDoc(path string) *goquery.Document {
	url := fmt.Sprintf(path)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible)")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("failed to convert '%s': %v", s, err)
		return -1
	}

	return i
}

func parseMadeAttempt(s string) (int, int) {
	parts := strings.Split(s, "-")
	if len(parts) != 2 {
		log.Printf("unexpected made/attempt value: %q", s)
		return 0, 0
	}

	made := stringToInt(strings.TrimSpace(parts[0]))
	att := stringToInt(strings.TrimSpace(parts[1]))
	return made, att
}
