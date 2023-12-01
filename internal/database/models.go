package database

type Book struct {
	ID        int32
	Title     string
	Author    string
	Isbn      string
	IssueYear int32
	Available bool
}
}
