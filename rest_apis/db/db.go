package db

type User struct {
	ID          string `json:"ID"`
	Account     string `json:"Account"`
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	Birthday    string `json:"BirthDay"`
	Description string `json:"Description"`
	Gender      int    `json:"Gender"`
}

type Users []User

var UserList Users

const (
	Male = iota
	Female
	Other
)

func Initialize() {
	UserList = Users{
		User{ID: "1", Account: "bob1", FirstName: "Bob", LastName: "Dylan", Birthday: "1941-05-24", Description: "アメリカ・ミネソタ州出身のミュージシャン。", Gender: Male},
		User{ID: "2", Account: "aretha2", FirstName: "Aretha", LastName: "Franklin", Birthday: "1942-03-25", Description: "アメリカ合衆国出身の女性ソウル歌手。", Gender: Female},
		User{ID: "3", Account: "david3", FirstName: "David", LastName: "Bowie", Birthday: "1947-01-08", Description: "イングランド出身のミュージシャン、シンガーソングライター、音楽プロデューサー、俳優。", Gender: Other},
	}
}

/*
User 4 json
{
"ID": "4",
"Account": "karen4",
"FirstName": "Karen",
"LastName": "Carpenter",
"Birthday": "1950-03-02",
"Description": "カーペンターズのヴォーカリスト、ドラマー。 ",
"Gender":1
}
*/
