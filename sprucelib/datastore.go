package sprucelib

type DataStore interface {
	// GetUserById(int) (User, error)
	// GetAllUsers() ([]User, error)
	GetUserByUsernameAndPassword(string, string) (User, error)
	GetUserByToken(string) (User, error)
	NewTokenForUser(User) (string, error)
	// StoreUser(*User) error
	// DeleteUser(*User) error
}
