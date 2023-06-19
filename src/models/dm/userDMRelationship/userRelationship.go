package userDMRelationship

import (
	"gochat/src/connections/database"
	"gochat/src/models/user"
)

type RelationshipInterface interface {
	Create() error
}

type UserDMRelationship struct {
	Id     int
	UserId int
	DMId   int
}

func (relationship UserDMRelationship) Create() error {
	conn := database.GetConnection()
	defer conn.Close()

	_, err := (*conn).Exec("INSERT INTO user_dms (user_id, dm_id) VALUES (?, ?)", relationship.UserId, relationship.DMId)
	if err != nil {
		return err
	}

	return nil
}

func (relationship UserDMRelationship) Delete() error {
	conn := database.GetConnection()
	defer conn.Close()

	_, err := (*conn).Exec("DELETE FROM user_dms WHERE user_id = ? AND dm_id = ?", relationship.UserId, relationship.DMId)
	if err != nil {
		return err
	}
	return nil
}

func (relationship UserDMRelationship) Exists() bool {
	conn := database.GetConnection()
	defer conn.Close()

	var count int
	(*conn).QueryRow("SELECT COUNT(*) FROM user_dms WHERE user_id = ? AND dm_id = ?", relationship.UserId, relationship.DMId).Scan(&count)

	return count > 0
}

func (relationship UserDMRelationship) GetMembers() ([]user.UserClient, error) {
	conn := database.GetConnection()
	defer conn.Close()

	query := `
	SELECT users.id, users.name, users.email 
	FROM users
	INNER JOIN user_dms ON users.id = user_dms.user_id
	WHERE user_dms.dm_id = ?`

	rows, err := (*conn).Query(query, relationship.DMId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []user.UserClient
	for rows.Next() {
		var member user.UserClient
		err := rows.Scan(&member.Id, &member.Name, &member.Email)
		if err != nil {
			return nil, err
		}
		members = append(members, member)
	}

	return members, nil
}
