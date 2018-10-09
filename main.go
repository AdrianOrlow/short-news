/* 

Source code of the Short News - a simple website
which provides news shortcuts. Made with Golang and vanilla CSS.

Avaiable online at https://short-news.herokuapp.com

(C) Made by Adrian Orłów

*/

package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	strip "github.com/grokify/html-strip-tags-go"
	rss "github.com/ungerik/go-rss"
)

//Article type
type Article struct {
	Title  string
	Desc   string
	Source string
	Link   string
}

//Category type
type Category struct {
	Name     string
	Articles []Article
}

//PageData type
type PageData struct {
	Categories []Category
}

//Returns user-friendly media name 
func getHumanMediaName(media string) (name string) {
	switch media {
	case "www.polsatnews.pl":
		name = "Polsat News"
		break
	case "www.rmf24.pl":
		name = "RMF 24"
		break
	case "www.tvn24.pl":
		name = "TVN24"
		break
	case "www.wykop.pl":
		name = "Wykop"
		break
	}

	return name
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

//Start template execution
func start(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/main.html")
	t.Execute(w, nil)
}

//Gets news from RSS channels and returns results
func results(w http.ResponseWriter, r *http.Request) {

	//Map with media rss channels link for every category 
	media := map[string]map[string]string{
		"polandCat": {
			"tvn24Med":      "https://www.tvn24.pl/wiadomosci-z-kraju,3.xml",
			"rmf24Med":      "https://www.rmf24.pl/fakty/polska/feed",
			"polsatnewsMed": "http://www.polsatnews.pl/rss/polska.xml",
			"wykopMed":       "https://www.wykop.pl/tag/znaleziska/polska/znaleziska/rss",
		},
		"worldCat": {
			"tvn24Med":      "https://www.tvn24.pl/wiadomosci-ze-swiata,2.xml",
			"rmf24Med":      "https://www.rmf24.pl/fakty/swiat/feed",
			"polsatnewsMed": "http://www.polsatnews.pl/rss/swiat.xml",
			"wykopMed":      "https://www.wykop.pl/tag/znaleziska/swiat/znaleziska/rss",
		},
		"economicsCat": {
			"tvn24Med":      "https://www.tvn24.pl/biznes-gospodarka,6.xml",
			"rmf24Med":      "https://www.rmf24.pl/ekonomia/feed",
			"polsatnewsMed": "http://www.polsatnews.pl/rss/biznes.xml",
			"wykopMed":      "https://www.wykop.pl/tag/znaleziska/ekonomia/znaleziska/rss",
		},
		"cultureCat": {
			"tvn24Med":      "https://www.tvn24.pl/kultura-styl,8.xml",
			"rmf24Med":      "https://www.rmf24.pl/kultura/feed",
			"polsatnewsMed": "http://www.polsatnews.pl/rss/kultura.xml",
			"wykopMed":      "https://www.wykop.pl/tag/znaleziska/kultura/znaleziska/rss",
		},
		"sportCat": {
			"tvn24Med":      "https://sport.tvn24.pl/sport,81,m.xml",
			"rmf24Med":      "https://www.rmf24.pl/sport/feed",
			"polsatnewsMed": "http://www.polsatnews.pl/rss/sport.xml",
			"wykopMed":      "https://www.wykop.pl/tag/znaleziska/sport/znaleziska/rss",
		},
	}

	pageData := PageData{
		Categories: []Category{},
	}

	if r.Method == "GET" {
		http.Redirect(w, r, "/", 301)
	} else {
		catKeys := [6]string{"polandCat", "worldCat", "politicsCat", "economicsCat", "cultureCat", "sportCat"}
		catNames := [6]string{"Polska", "Świat", "Polityka", "Gospodarka", "Kultura", "Sport"}
		medKeys := [4]string{"tvn24Med", "rmf24Med", "polsatnewsMed", "wykopMed"}

		r.ParseForm()
		form := r.Form
		for i := 0; i < len(catKeys); i++ {
			if _, ok := form[catKeys[i]]; ok {
				key := catKeys[i]
				var rssLinks []string

				item := Category{Name: catNames[i]}
				pageData.Categories = append(pageData.Categories, item)
				var index int
				for l := 0; l < len(pageData.Categories); l++ {
					if pageData.Categories[l].Name == catNames[i] {
						index = l
						break
					}
				}

				//Creates array with choosen media rss links
				for j := 0; j < len(medKeys); j++ {
					if _, ok := form[medKeys[j]]; ok {
						if val, ok := media[key][medKeys[j]]; ok {
							rssLinks = append(rssLinks, val)
						}
					}
				}

				//Gets articles
				q, _ := strconv.ParseInt(r.FormValue("quantity")[0:], 10, 64)
				var k int64
				for ; k < q; k++ {
					rLink := rssLinks[random(0, len(rssLinks))]
					channel, _ := rss.Read(rLink)

					rL, _ := url.Parse(rLink)
					rLink = getHumanMediaName(rL.Host)

					rand.Seed(time.Now().UnixNano())
					items := len(channel.Item) - 1
					rArticle := random(0, items)

					c := channel.Item[rArticle]
					article := Article{
						Title: c.Title,
						Desc: strip.StripTags(c.Description),
						Source: rLink,
						Link: c.Link
					}

					pageData.Categories[index].Articles = append(pageData.Categories[index].Articles, article)
				}
			}
		}
	}

	//Executes template with data
	t, _ := template.ParseFiles("templates/results.html")
	t.Execute(w, pageData)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", start)
	http.HandleFunc("/results", results)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":"+port, nil)
}
