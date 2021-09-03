package model

type JobsMetaData struct {
	Title string
	Logo  string
	Jobs  []*Job
}
type Job struct {
	Position string
	URL      string
}
