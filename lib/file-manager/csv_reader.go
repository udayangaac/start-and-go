package file_manager

type CSVReader interface {
	ReadCSV(file string) (rows []Row, err error)
}
