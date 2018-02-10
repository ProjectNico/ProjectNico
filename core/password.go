/* Copyright TuziMoe 2018 All Rights Reserved
   This is the password.go for ProjectNico */
package core

import "time"

const PasswordLength = 256
type Password [PasswordLength]byte

func init() {
	rand.Seed(time.Now().Unix())
}
func RandPassword() *Password {
	intArr := rand.Perm(PasswordLength)
	password := &Password{}
	for i, v := range intArr {
		password[i] = byte(v)
		if i == v {
			return RandPassword()
		}
	}
	return password
}