package transportdata

import (
	"encoding/csv"
	"log"
	"os"
)

func LoadStops(filePath string, graph graph) {
	reader := makeCsvReader(filePath)

	for record, _ := reader.Read(); record != nil; record, _ = reader.Read() {
		stopId := record[0]
		stopName := record[2]
		graph.AddNode(stopId, stopName)
	}
}

func makeCsvReader(filePath string) *csv.Reader {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)

	// Skip over headers
	reader.Read()

	return reader
}
