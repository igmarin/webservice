package models

// User is (this line is for the linter)
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

func getUsers() []*User {
	return users
}

func addUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}
