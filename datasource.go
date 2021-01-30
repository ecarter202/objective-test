package main

type DataSource interface {
	Load(string) error
}

var (
	supportedFileTypes = map[string]DataSource{
		"json": JSON{},
	}
)

var selectedDataSource DataSource
