// this program demonstrates how fetches through gopherjs seem to be memory leaking.
package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Set("gocode", map[string]interface{}{
		"fetch": fetch,
	})
}

const (
	maxFetch = 10000
	fetchURL = "http://localhost:8080/gocode.js"
)

var fetchInterval = 1 * time.Millisecond

func fetch() {
	go func() {
		for i := 0; i < maxFetch; i++ {
			time.Sleep(fetchInterval)
			req, err := http.NewRequest("GET", fetchURL, nil)
			if err != nil {
				log.Fatal(err)
			}
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Fatal(err)
			}
			func() {
				defer res.Body.Close()
				if _, err := io.Copy(ioutil.Discard, res.Body); err != nil {
					log.Fatal(err)
				}
			}()
		}
	}()
}
