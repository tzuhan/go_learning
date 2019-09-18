package restapi

import (
	"fmt"
	"net/http"
)

type User struct {
	ID      int
	Account string
}
type UserInfo struct {
	ID          int
	FirstName   string
	LastName    string
	Birthday    string
	Description string
	Gender      int
}
type Users []User
type UserInfos []UserInfo

const (
	Male = iota
	Female
	Other
)

var UserList Users
var UserInfoList UserInfos

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome\n")
}

func preset() {
	UserList = Users{User{ID: 1, Account: "bob"}, User{ID: 2, Account: "aretha"}, User{ID: 3, Account: "david"}}
	UserInfoList = UserInfos{
		UserInfo{ID: 1, FirstName: "Bob", LastName: "Dylan", Birthday: "1941-05-24", Description: "アメリカ・ミネソタ州出身のミュージシャン。", Gender: Male},
		UserInfo{ID: 2, FirstName: "Aretha", LastName: "Franklin", Birthday: "1942-03-25", Description: "アメリカ合衆国出身の女性ソウル歌手。", Gender: Female},
		UserInfo{ID: 3, FirstName: "David", LastName: "Bowie", Birthday: "1947-01-08", Description: "イングランド出身のミュージシャン、シンガーソングライター、音楽プロデューサー、俳優。", Gender: Other},
	}
}
