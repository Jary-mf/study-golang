package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func GetUrl_main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		fmt.Println(url)
		res, err := http.Get(url)
		fmt.Printf("%s\n", res.Status)

		if err != nil {
			continue
		}
		// _, e := io.Copy(os.Stdout, res.Body)
		// if e != nil {
		// 	fmt.Fprintf(os.Stderr, "%s", e)
		// }
		defer res.Body.Close()
	}
}
