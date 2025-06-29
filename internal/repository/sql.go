package repository

type SQLRepository interface {
	CreateUser(userEmail, pwd string) error
	RetrieveHashedPwd(userEmail string) ([]byte, error)
}
