package transportdata

import (
	"encoding/csv"
	"log"
	"os"
)

func loadStops(filePath string, graph graph) {
	reader := makeCsvReader(filePath)

	for record, _ := reader.Read(); record != nil; record, _ = reader.Read() {
		stopName := record[2]
		graph.addVertex(stopName)
	}
}

func makeCsvReader(filePath string) *csv.Reader {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)

	// Skip over headers
	_, err = reader.Read()

	if err != nil {
		log.Fatal(err)
	}

	return reader
}
