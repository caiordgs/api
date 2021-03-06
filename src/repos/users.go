package repos

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (repository Users) Create(user models.User) (uint64, error) {
	statement, erro := repository.db.Prepare(
		"insert into users (name, nick, email, password) values(?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if erro != nil {
		return 0, erro
	}

	lastInsertedID, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(lastInsertedID), nil
}

func (repository Users) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	lines, erro := repository.db.Query(
		"select id, name, nick, email, createdIn from users where name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)

	if erro != nil {
		return nil, erro
	}

	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedIn,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository Users) IdSearch(ID uint64) (models.User, error) {
	lines, erro := repository.db.Query(
		"select id, name, nick, email, createdIn from users where id = ?",
		ID,
	)
	if erro != nil {
		return models.User{}, erro
	}
	defer lines.Close()

	var user models.User

	if lines.Next() {
		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedIn,
		); erro != nil {
			return models.User{}, erro
		}
	}
	return user, nil
}

func (repository Users) Update(ID uint64, user models.User) error {
	statement, erro := repository.db.Prepare(
		"update users set name = ?, nick = ?, email = ? where id = ?",
	)
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(user.Name, user.Nick, user.Email, ID); erro != nil {
		return erro
	}

	return nil
}

func (repository Users) Delete(ID uint64) error {
	statement, erro := repository.db.Prepare("delete from users where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		return erro
	}

	return nil
}

func (repository Users) SearchByEmail(email string) (models.User, error) {
	line, erro := repository.db.Query("select id, password from users where email = ?", email)
	if erro != nil {
		return models.User{}, erro
	}
	defer line.Close()

	var user models.User
	if line.Next() {
		if erro = line.Scan(&user.ID, &user.Password); erro != nil {
			return models.User{}, erro
		}
	}

	return user, nil
}

func (repository Users) Follow(userID, followerID uint64) error {
	statement, erro := repository.db.Prepare("insert ignore into followers (user_id, follower_id) values (?, ?)")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(userID, followerID); erro != nil {
		return erro
	}

	return nil
}

func (repository Users) Unfollow(userID, followerID uint64) error {
	statement, erro := repository.db.Prepare("delete from followers where user_id = ? and follower_id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(userID, followerID); erro != nil {
		return erro
	}

	return nil
}

func (repository Users) FindFollowers(userID uint64) ([]models.User, error) {
	lines, erro := repository.db.Query(`
		select u.id, u.name, u.nick, u.email, u.createdIn
		from users u inner join followers f on u.id = f.follower_id where f.user_id = ?`,
		userID)
	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var users []models.User
	for lines.Next() {
		var user models.User

		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedIn,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository Users) FindFollowing(userID uint64) ([]models.User, error) {
	lines, erro := repository.db.Query(`
	select u.id, u.name, u.nick, u.email, u.createdIn
	from users u inner join followers f on u.id = f.user_id where f.follower_id = ?`,
		userID,
	)
	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var users []models.User
	for lines.Next() {
		var user models.User

		if erro = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedIn,
		); erro != nil {
			return nil, erro
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository Users) SearchPassword(userID uint64) (string, error) {
	line, erro := repository.db.Query("select password from users where id = ?", userID)
	if erro != nil {
		return "", erro
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if erro = line.Scan(&user.Password); erro != nil {
			return "", erro
		}
	}

	return user.Password, nil
}

func (repository Users) PasswordUpdate(userID uint64, password string) error {
	statement, erro := repository.db.Prepare("update users set password = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(password, userID); erro != nil {
		return erro
	}

	return nil
}
