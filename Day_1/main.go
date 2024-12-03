package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("fixlets.csv")

	if err != nil {
		log.Fatalf("Failed to opne file: %v\n", err)
	}

	r := csv.NewReader(f)

	if _, err := r.Read(); err != nil {
		log.Fatal(err)
	}

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s, %s, %s, %s, %s\n", record[0], record[1], record[2], record[3], record[4])

	}
}
