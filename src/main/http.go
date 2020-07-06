package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) add(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	p := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(p, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "price should be a float ,but provided &s", p)
		return
	}
	if _, ok := db[item]; ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%s already exists ", item)
		return
	}
	db[item] = dollars(price)
	db.list(w, req)
}
func (db database) search(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s doesn't exist ", item)
		return
	}
	fmt.Fprintf(w, "%s: %s\n", item, price)
}
func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s doesn't exist ", item)
		return
	}
	delete(db, item)
	db.list(w, req)

}
func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	p := req.URL.Query().Get("price")
	price, err := strconv.ParseFloat(p, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "price should be a float ,but provided &s", p)
		return
	}
	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s doesn't exist ", item)
		return
	}
	db[item] = dollars(price)
	db.list(w, req)
}

func HTTP_main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/", db.list)
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/add", db.add)
	http.HandleFunc("/search", db.search)
	http.HandleFunc("/delete", db.delete)
	http.HandleFunc("/update", db.update)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
