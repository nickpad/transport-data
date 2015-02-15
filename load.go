package transportdata

import (
	"encoding/csv"
	"log"
	"os"
)

func LoadStops(filePath string, database Database) {
	file := loadFileOrExit(filePath)
	reader := csv.NewReader(file)

	// Skip over headers
	reader.Read()

	for record, _ := reader.Read(); record != nil; record, _ = reader.Read() {
		stopId := record[0]
		stopName := record[2]
		hopEdges := make([]HopEdge, 0)
		stopNode := StopNode{stopId, stopName, hopEdges}
		database[stopNode.Id] = stopNode
	}
}

func loadFileOrExit(filePath string) *os.File {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	return file
}
