package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func ParseXML_main() {
	res := fetchurl(os.Args[1])
	r := strings.NewReader(res)
	dec := xml.NewDecoder(r)
	var stack []string
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect : %v\n", err)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if contrainsAll(stack, os.Args[2:]) {
				fmt.Printf("%s : %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

func fetchurl(url string) string {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch:reading %s : %v\n", url, err)
		os.Exit(1)
	}
	return fmt.Sprintf("%s", b)
}

func contrainsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
