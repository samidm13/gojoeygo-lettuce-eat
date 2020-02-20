package main
  
import (
	_ "github.com/lib/pq"

	"golang.org/x/crypto/bcrypt"
)

var (
	userCount int
	user_id   int
	email     string
	password  string
)

func userSignUp(firstname string, lastname string, mail string, pass string) int {

	hashedpass, _ := HashPassword(pass)

	sqlStatement := `

  INSERT INTO users (first_name, last_name, email, password)
  VALUES ($1, $2, $3, $4) RETURNING user_id`
	err := DB.QueryRow(sqlStatement, firstname, lastname, mail, hashedpass).Scan(&user_id)
	if err != nil {
		panic(err)
	}
	return user_id
}
func userLogIn(remail string, rpassword string) (int, string) {
	DB.QueryRow("SELECT COUNT(user_id) AS userCount, user_id, email, password FROM users WHERE email=$1 GROUP BY user_id", remail).Scan(&userCount, &user_id, &email, &password)
	return user_id, password
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
