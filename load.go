package transportdata

import (
	"encoding/csv"
	"log"
	"os"
)

func LoadStops(filePath string, database StopIndex) {
	reader := makeCsvReader(filePath)

	for record, _ := reader.Read(); record != nil; record, _ = reader.Read() {
		stopId := record[0]
		stopName := record[2]
		hopEdges := make([]HopEdge, 0)
		stopNode := StopNode{stopId, stopName, hopEdges}
		database[stopNode.Id] = stopNode
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
