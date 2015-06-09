package spruce

type User struct {
	Key   string
	Email string
}

type UserService interface {
	Get(id int) (*User, error)
	FindByEmail(email string) (*User, error)
	Save(user *User) (created bool, err error)
	Delete(user *User) (deleted bool, err error)
}

type UserDatastore struct {
}

func (store *UserDatastore) Save(user *User) (bool, error) {

	// url := "https://heedowereariedimentouthe:BwClLuPcRu8qtBkr4ACMKf23@adlongwell.cloudant.com/texere/"

	return true, nil
}
