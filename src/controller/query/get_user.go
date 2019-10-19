package query
import(
	mo "project_reservasi/src/model"
	"database/sql"

)

var db *sql.DB
var err error

func QueryUser(username string) mo.User {
	var users = mo.User{}
	err = db.QueryRow(`
		SELECT id, 
		username, 
		password 
		FROM user WHERE username=?
		`, username).
		Scan(
			&users.Id,
			&users.Username,
			&users.Password,
		)
	return users
}