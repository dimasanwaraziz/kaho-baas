package services

import (
	"golang.org/x/crypto/bcrypt"
)

type UserData struct {
	ID          int
	Username    string
	Email       string
	Password    string
	CompanyId   string
	CompanyName string
	Roles       []string
	Name        string
}

func (a *authService) Login(email, password string) (UserData, error) {

	var userData UserData
	err := a.db.QueryRow(`SELECT
			u.id,
			username,
			email,
			password,
			u.company,
			c.name,
			full_name
		from
			auth.users u
		left join vistalenta.company c on
			u.company = c.company_id 
		WHERE email = $1 or username = $1`,
		email).Scan(&userData.ID, &userData.Username, &userData.Email, &userData.Password, &userData.CompanyId, &userData.CompanyName, &userData.Name)
	if err != nil {
		return UserData{}, err
	}

	var roleData []string
	rows, err := a.db.Query(`
		SELECT r.name
		FROM auth.roles_users ru
		LEFT JOIN auth.roles r ON r.id = ru.role_id
		LEFT JOIN auth.users u ON u.id = ru.user_id
		WHERE u.email = $1
	`, email)

	if err != nil {
		return UserData{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var role string
		if err := rows.Scan(&role); err != nil {
			return UserData{}, err
		}
		roleData = append(roleData, role)
	}

	if err := rows.Err(); err != nil {
		return UserData{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(password))
	if err != nil {
		return UserData{}, err
	}

	userData.Roles = roleData
	return userData, nil
}
