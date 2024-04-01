package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"math/rand"
	"time"
	pb "user_service/genproto/user-service"
)

type userRepo struct {
	db *sqlx.DB
}

func (u *userRepo) Create(reg *pb.User) (*pb.User, error) {
	query := `INSERT INTO users (first_name, 
                   last_name, 
                   username,
                   role,
                   password, 
                   email,
                   id, 
                   refreshtoken,
                   created_at,
                   updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING first_name,
                	last_name, 
                   username, 
                   role,
                   password, 
                   email,
                   id, 
                   refreshtoken,
                   created_at,
                   updated_at`
	var user pb.User
	err := u.db.QueryRow(query,
		reg.FirstName,
		reg.LastName,
		reg.Username,
		reg.Role,
		reg.Password,
		reg.Email,
		reg.Id,
		reg.RefreshToken,
		time.Now(),
		time.Now(),
	).Scan(
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Role,
		&user.Password,
		&user.Email,
		&user.Id,
		&user.RefreshToken,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) GetUser(id string) (*pb.User, error) {
	var user pb.User
	err := u.db.QueryRow(`SELECT first_name, last_name, username, role, password, email, id, refreshtoken, created_at, updated_at FROM users WHERE id = $1 `, id).Scan(
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Role,
		&user.Password,
		&user.Email,
		&user.Id,
		&user.RefreshToken,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepo) GetAll(page, limit int64) (users []*pb.User, err error) {
	offset := limit * (page - 1)
	rows, err := u.db.Query(`SELECT * FROM users LIMIT $1 OFFSET $2`, limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user pb.User
		err := rows.Scan(
			&user.FirstName, &user.LastName, &user.Username, &user.Role, &user.Password, &user.Email, &user.Id, &user.RefreshToken, &user.CreatedAt, &user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (u *userRepo) Update(user *pb.User) (*pb.User, error) {
	err := u.db.QueryRow(`UPDATE users SET first_name = $1, last_name = $2, refreshtoken = $3, updated_at = now() WHERE email = $4 RETURNING first_name, last_name, username, created_at, updated_at`,
		user.FirstName,
		user.LastName,
		user.RefreshToken,
		user.Email,
	).Scan(
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepo) Delete(user_id string) error {
	_, err := u.db.Exec(`DELETE FROM users WHERE id = $1`, user_id)
	if err != nil {
		return err
	}
	return nil
}

func (u *userRepo) CheckUniquess(reg *pb.CheckUniqReq) (int32, error) {
	var count int
	err := u.db.QueryRow(fmt.Sprintf(`SELECT count(1) FROM users WHERE %s = $1`, reg.Field), reg.Value).Scan(&count)
	if err != nil {
		return 0, err
	}
	if count != 0 {
		return 0, err
	}
	n := rand.Int31() % 1000000
	return n, nil
}

func (u *userRepo) Exists(email string) (*pb.User, error) {
	var user pb.User
	err := u.db.QueryRow(`SELECT first_name, last_name, username, role, password, email, id, refreshtoken FROM users WHERE email = $1`, email).Scan(
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Role,
		&user.Password,
		&user.Email,
		&user.Id,
		&user.RefreshToken,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}
