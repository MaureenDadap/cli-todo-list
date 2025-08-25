package main

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
)

func ReadAndParseCSV(filename string) CSV {
	csvData := readCSVFile(filename)

	data := parseCSV(csvData)

	// Transform matrix to CSV object
	csv := CSV{}
	csv.SetHeader(ToRow(data[0]))

	if len(data) > 0 {
		for i := 1; i < len(data); i++ {
			csv.AppendRow(ToRow(data[i]))
		}
	}

	return csv
}

// Read CSV file
func readCSVFile(filename string) []byte {
	f := openFile(filename)
	defer f.Close()

	return Must(io.ReadAll(f))
}

// Open a CSV File.
// If the file does not exist, it will be created and the default headers will be added
// This does not close the file returned, so it must be cleaned up by the function caller
func openFile(filename string) *os.File {
	file, err := os.Open(filename)

	if err != nil {
		println("No CSV file, creating..")

		file = Must(os.Create(filename))
		addDefaultHeaders(file)
	}

	return file
}

// Writes default headers to the file
func addDefaultHeaders(file io.Writer) {
	writer := csv.NewWriter(file)
	defer writer.Flush()

	CheckErr(writer.Write(getHeaders()))
}

// Get the expected header of the CSV file
func getHeaders() []string {
	return []string{"ID", "Name", "Created At", "Is Done"}
}

// Parses CSV and returns it as a matrix of strings
func parseCSV(data []byte) [][]string {
	reader := csv.NewReader(bytes.NewReader(data))

	records := Must(reader.ReadAll())

	return records
}

// ---- CSV Struct stuff ----

type CSV struct {
	header Row
	rows   []Row
}

type Row struct {
	columns []string
}

func (c *CSV) SetHeader(row Row) {
	c.header = row
}

func (c *CSV) AppendRow(row Row) {
	c.rows = append(c.rows, row)
}

// Transform raw array to Row object
func ToRow(data []string) Row {
	row := Row{}

	row.columns = append(row.columns, data...)

	return row
}
