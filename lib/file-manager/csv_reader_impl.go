package file_manager

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type csvReader struct {
	Path string
}

func NewCSVReader(path string) CSVReader {
	return &csvReader{
		Path: path,
	}
}

type Row struct {
	Column []string
}

func (c *csvReader) ReadCSV(file string) (rows []Row, err error) {
	rows = make([]Row, 0)
	csvFile, _ := os.Open(fmt.Sprintf("%v/%v", c.Path, file))
	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		var line []string
		line, err = reader.Read()
		if err == io.EOF {
			err = nil
			break
		} else if err != nil {
			continue
		}
		row := Row{
			Column: line,
		}
		rows = append(rows, row)
	}
	return
}
