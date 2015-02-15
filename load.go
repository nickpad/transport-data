package transportdata

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func LoadStops(filePath string) Database {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(file)
	database := Database{}

	record, err := reader.Read()

	if err == io.EOF {
		return database
	} else if err != nil {
		log.Fatal(err)
	}

	for record != nil {
		record, err = reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		stopName := record[2]
		hopEdges := make([]HopEdge, 0)
		stopNode := StopNode{stopName, hopEdges}
		database[stopNode.Name] = stopNode
	}

	return database
}
