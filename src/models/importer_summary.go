package models

type ImporterSummary struct {
	UUID        string
	FilePath    string
	SuccessRows uint32
	ErrorRows   uint32
	TotalRows   uint32
	ProcessType string
}
