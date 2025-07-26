package scraper

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rs/zerolog/log"
)

func getPageDoc(path string) *goquery.Document {
	url := fmt.Sprintf(path)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible)")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal().Err(err).Msg("request failed")
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatal().Int("code", res.StatusCode).Str("status", res.Status).Msg("status code error")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse document")
	}

	return doc
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Error().Err(err).Str("value", s).Msg("failed to convert")
		return -1
	}

	return i
}

func parseMadeAttempt(s string) (int, int) {
	parts := strings.Split(s, "-")
	if len(parts) != 2 {
		log.Error().Str("value", s).Msg("unexpected made/attempt value")
		return 0, 0
	}

	made := stringToInt(strings.TrimSpace(parts[0]))
	att := stringToInt(strings.TrimSpace(parts[1]))
	return made, att
}
