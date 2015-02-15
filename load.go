package transportdata

import (
	"encoding/csv"
	"log"
	"os"
)

func LoadStops(filePath string) Database {
	file := loadFileOrExit(filePath)
	reader := csv.NewReader(file)
	database := Database{}

	// Skip over headers
	reader.Read()

	for record, _ := reader.Read(); record != nil; record, _ = reader.Read() {
		stopName := record[2]
		hopEdges := make([]HopEdge, 0)
		stopNode := StopNode{stopName, hopEdges}
		database[stopNode.Name] = stopNode
	}

	return database
}

func loadFileOrExit(filePath string) *os.File {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	return file
}
