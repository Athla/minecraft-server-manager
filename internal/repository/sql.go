package repository

type SQLRepository interface {
	RetrieveHashedPwd(userEmail string) error
}
