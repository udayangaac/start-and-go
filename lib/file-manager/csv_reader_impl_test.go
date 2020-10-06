package file_manager

import "testing"

func TestReadCSV(t *testing.T) {
	reader := NewCSVReader("/home/chamith/go/src/github.com/sixensor/y-not-advertising-api/resources")
	rows, _ := reader.ReadCSV("sample.csv")
	t.Log(rows)
}
