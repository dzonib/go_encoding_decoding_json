package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// with uppercase you can access it without being member of this package

type User struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
}

type UserDB struct {
	Users []User `json:"users,omitempty"`
	Type  string `json:"type,omitempty"`
}

func main() {

	f, err := os.Open("user.db.json")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	dec := json.NewDecoder(f)

	db := UserDB{}

	dec.Decode(&db)

	fmt.Println(db)

}

func createJSONFile(db UserDB) {
	users := []User{
		{Username: "John Doe", Password: "change me", Email: "johndoe@email.com"},
		{Username: "John Seconds", Password: "me change", Email: "johndoe2@email.com"},
	}

	db = UserDB{Users: users, Type: "Simple db"}

	var buf = new(bytes.Buffer)

	enc := json.NewEncoder(buf)

	enc.Encode(db)

	f, err := os.Create("user.db.json")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	io.Copy(f, buf)
}
