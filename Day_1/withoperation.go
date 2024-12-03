package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

type Record struct {
	Field1, Field2, Field3, Field4, Field5 string
}

func readCSV(filePath string) ([]Record, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	var records []Record

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		records = append(records, Record{
			Field1: record[0], Field2: record[1], Field3: record[2],
			Field4: record[3], Field5: record[4],
		})
	}
	return records, nil
}

func writeCSV(filePath string, records []Record) error {
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	for _, record := range records {
		err := w.Write([]string{record.Field1, record.Field2, record.Field3, record.Field4, record.Field5})
		if err != nil {
			return fmt.Errorf("failed to write record: %v", err)
		}
	}
	w.Flush()
	return nil
}

func listRecords(records []Record) {
	fmt.Println("Listing all records:")
	for _, record := range records {
		fmt.Printf("%s, %s, %s, %s, %s\n", record.Field1, record.Field2, record.Field3, record.Field4, record.Field5)
	}
}

func queryRecords(records []Record, query string) {
	fmt.Printf("Querying records for: %s\n", query)
	for _, record := range records {
		if strings.Contains(record.Field1, query) || strings.Contains(record.Field2, query) ||
			strings.Contains(record.Field3, query) || strings.Contains(record.Field4, query) ||
			strings.Contains(record.Field5, query) {
			fmt.Printf("%s, %s, %s, %s, %s\n", record.Field1, record.Field2, record.Field3, record.Field4, record.Field5)
		}
	}
}

func sortRecords(records []Record) {
	sort.Slice(records, func(i, j int) bool {
		return records[i].Field1 < records[j].Field1
	})
	fmt.Println("Sorted records by first field:")
	listRecords(records)
}

func addRecord(records []Record, newRecord Record) []Record {
	records = append(records, newRecord)
	fmt.Println("Record added:")
	fmt.Printf("%s, %s, %s, %s, %s\n", newRecord.Field1, newRecord.Field2, newRecord.Field3, newRecord.Field4, newRecord.Field5)
	return records
}

func deleteRecord(records []Record, field string) []Record {
	var updatedRecords []Record
	for _, record := range records {
		if record.Field1 != field {
			updatedRecords = append(updatedRecords, record)
		}
	}
	return updatedRecords
}

func main() {
	filePath := "fixlets.csv"

	records, err := readCSV(filePath)
	if err != nil {
		log.Fatal(err)
	}

	listRecords(records)

	queryRecords(records, "Low")

	sortRecords(records)

	newRecord := Record{"1","501217007","MS22-AUG: Security Update for Windows Server 2022 - Windows Server 2022 - KB5012170001 (x64)","Low","100"}
	records = addRecord(records, newRecord)

	records = deleteRecord(records, "")

	err = writeCSV(filePath, records)
	if err != nil {
		log.Fatal(err)
	}

	listRecords(records)
}
