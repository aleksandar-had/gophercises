package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/aleksandar-had/gophercises/linksparser"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:"loc"`
}
type urlset struct {
	Xmlns string `xml:"xmlns,attr"`
	Urls  []loc  `xml:"url"`
}

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")
	maxDepth := flag.Int("depth", 3, "the maximum number of links' depth to traverse")

	flag.Parse()

	// fmt.Println(*urlFlag) //  GET the webpage
	// fmt.Println(*maxDepth)

	// 1. GET the webpage
	// 2. Parse all the links on the page
	// 3. Build proper urls with the links
	// 4. Filter out any links with a diff domain
	// 5. Find all the pages (BFS)

	links := linksBFS(*urlFlag, *maxDepth)

	// 6. Print out XML
	encodeXML(links)

	//for _, link := range links {
	//	fmt.Println(link)
	//}

}

func linksBFS(href string, maxDepth int) []string {
	discovered := make(map[string]struct{})
	var q map[string]struct{}
	nq := map[string]struct{}{
		href: {},
	}
	for i := 0; i <= maxDepth; i++ {
		q, nq = nq, make(map[string]struct{})
		if len(q) == 0 {
			break
		}
		for url := range q {
			if _, ok := discovered[url]; ok {
				continue
			}
			discovered[url] = struct{}{} // mark url as seen
			for _, link := range get(url) {
				nq[link] = struct{}{} // get all the links from the links of the current depth
			}
		}
	}
	res := make([]string, 0, len(discovered))
	for url := range discovered {
		res = append(res, url)
	}
	return res
}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}

	base := baseUrl.String()

	return filter(hrefs(resp.Body, base), withPrefix(base))
}

func filter(links []string, keepFn func(string) bool) []string {
	var res []string

	for _, link := range links {
		if keepFn(link) {
			res = append(res, link)
		}
	}

	return res
}

func withPrefix(pref string) func(string) bool {
	return func(link string) bool {
		return strings.Contains(link, pref)
	}
}

func hrefs(r io.Reader, base string) []string {
	links, err := linksparser.Parse(r)
	if err != nil {
		panic(err)
	}
	var res []string
	for _, link := range links {
		switch {
		case strings.HasPrefix(link.Href, "/"):
			res = append(res, base+link.Href)
		case strings.HasPrefix(link.Href, "http"):
			res = append(res, link.Href)
		}
	}

	var filt_hrefs []string
	for _, href := range res {
		if strings.Contains(href, base) {
			filt_hrefs = append(filt_hrefs, href)
		}
	}
	return res
}

func encodeXML(links []string) {
	var toXml urlset
	for _, link := range links {
		toXml.Urls = append(toXml.Urls, loc{link})
	}
	toXml.Xmlns = xmlns

	fmt.Print(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")

	if err := enc.Encode(toXml); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
