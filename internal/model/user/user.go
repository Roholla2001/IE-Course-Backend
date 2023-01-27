package user

import "golang.org/x/crypto/bcrypt"

type UserModel struct {
	ID       int64 `gorm:"primaryKey"`
	Username string
	Password string
}

func (u *UserModel) BeforeSave() error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)

	return nil

}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
