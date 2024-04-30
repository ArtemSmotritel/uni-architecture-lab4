package types

type Author struct {
	ID        int64
	FullName  string
	ShortName string
}

func NewAuthor(id int64, fullName, shortName string) Author {
	return Author{
		ID:        id,
		FullName:  fullName,
		ShortName: shortName,
	}
}
