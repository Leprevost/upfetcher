package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	const VERSION = "1.0.0"

	file := flag.String("in", "", "File with a list of UniProt IDs")
	out := flag.String("out", "", "Output name")
	flag.Parse()

	if len(*file) > 0 {

		f, err := ioutil.ReadFile(*file)
		if err != nil {
			log.Println("[ERROR] Cannot open file")
		}

		list := strings.Split(string(f), "\n")

		o, err := os.Create(*out)
		if err != nil {
			log.Println("[ERROR] Could not create outpur file")
			return
		}
		defer o.Close()

		for i := range list {

			var query string
			if len(list[i]) > 0 {
				query = fmt.Sprintf("http://www.uniprot.org/uniprot/%s.fasta", list[i])

				fmt.Println(query)

				response, err := http.Get(query)
				if err != nil {
					log.Println("[ERROR] Could not download annotation file", query, "-", err)
					return
				}
				defer response.Body.Close()

				n, err := io.Copy(o, response.Body)
				if err != nil {
					log.Println("[ERROR] Could not download uniprot record", list[i])
					return
				}
				_ = n
			}
		}

	}

	return
}
