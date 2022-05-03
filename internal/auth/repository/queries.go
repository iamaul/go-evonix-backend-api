package repository

const (
	createUserQuery = `INSERT INTO users (name, password, email, registered_date, register_ip) VALUES ($1, $2, $3, $4, $5) RETURNING *`

	getUserByEmailOrUsername = `SELECT id, name, email FROM users WHERE name = $1 OR email = $1`
)
