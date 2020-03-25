package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	// testPage := &WikiPage{Title: "testpage", Body: []byte("I am the stone that the builder refused")}

	// // saveWiki
	// err := testPage.saveWiki()
	// if err != nil {
	// 	fmt.Println("SAVE_ERR: ", err.Error())
	// }

	// // loadWiki
	// wiki, err := loadWiki("testpage")
	// if err != nil {
	// 	fmt.Println("LOAD_ERR: ", err.Error())
	// }

	// fmt.Println(string(wiki.Body), "Accessed at:", wiki.AccessedAt)

	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":5000", nil))
}

// WikiPage shows the fields that make up a wiki page
type WikiPage struct {
	Title      string
	Body       []byte
	AccessedAt time.Time
}

// WikiPage.saveWiki() persists WikiPage's body to disk with WikiPage.title as the filename. It return an err if any
// If file is non-existent, it is created with perm 0600 <an octal integer literal that means read-write for current user only
func (p *WikiPage) saveWiki() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadWiki(title string) (*WikiPage, error) {
	file, err := ioutil.ReadFile(title)
	if err != nil {
		return nil, err
	}
	return &WikiPage{Title: title, Body: file, AccessedAt: time.Now()}, nil
}

func viewHandler(responseWriter http.ResponseWriter, request *http.Request) {
	title := request.URL.Path[len("/view/"):]

	wiki, err := loadWiki(title)
	if err != nil {
		fmt.Fprintf(responseWriter, "Something went wrong: %s", err.Error())
		return
	}

	fmt.Fprintf(responseWriter, "<html><h2>%s</h2><div><pre>%s</pre></div><footer><small>Accessed at: %s</small></footer></html>", wiki.Title, wiki.Body, wiki.AccessedAt)
}
