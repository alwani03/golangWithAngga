package management

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvalaible bool
}

func (user User) Display() string {

	return fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.Email)

}

func (group Group) DisplayGroup() {
	fmt.Printf("Name Group: %s", group.Name)
	fmt.Println("")
	fmt.Printf("Jumlah Anggota: %d", len(group.Users))
	fmt.Println("")

	for _, value := range group.Users {
		fmt.Println(value.FirstName)
	}
}
