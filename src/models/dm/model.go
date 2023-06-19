package dm

import (
	"database/sql"
	"gochat/src/connections/database"
	"gochat/src/models/dm/userDMRelationship"
	"gochat/src/models/user"
)

type DMInterface interface {
	Get() (DM, error)
	Create() error
	Update() error
	Delete() error

	Authenticate() bool
	Exists() (bool, error)
	//HasMember() (bool, error)
}

type DM struct {
	Id          int
	WorkspaceId int
}

// DM Model to be returned to client
type ClientDM struct {
	Id          string
	WorkspaceId int
}

func (dm DM) Get() (DM, error) {

	conn := database.GetConnection()
	defer conn.Close()

	query := `
	SELECT id, workspace_id
	FROM dms
	WHERE id = ? AND workspace_id = ?`

	if err := (*conn).QueryRow(query, dm.Id, dm.WorkspaceId).Scan(&dm.Id, &dm.WorkspaceId); err != nil {
		return dm, err
	}

	return dm, nil
}

func (dm DM) Create() error {
	conn := database.GetConnection()
	defer conn.Close()

	query := `
	INSERT INTO dms (id, workspace_id)
	VALUES (?, ?)`

	_, err := (*conn).Exec(query, dm.Id, dm.WorkspaceId)
	if err != nil {
		return err
	}

	return nil
}

func (dm DM) Update() error {
	conn := database.GetConnection()
	defer conn.Close()

	query := `
	UPDATE dms
	SET workspace_id = ?
	WHERE id = ?`

	_, err := (*conn).Exec(query, dm.WorkspaceId, dm.Id)
	if err != nil {
		return err
	}

	return nil
}

func (dm DM) Delete() error {
	conn := database.GetConnection()
	defer conn.Close()

	query := `
	DELETE FROM dms
	WHERE id = ?`

	_, err := (*conn).Exec(query, dm.Id)
	if err != nil {
		return err
	}

	return nil
}

func (dm DM) Authenticate() bool {
	return true
}

func (dm DM) Exists() (bool, error) {
	conn := database.GetConnection()
	defer conn.Close()

	query := `
	SELECT id
	FROM dm
	WHERE id = ?`

	var id int
	if err := (*conn).QueryRow(query, dm.Id).Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func GetDMsByUserId(userId int) ([]DM, error) {
	conn := database.GetConnection()
	defer conn.Close()

	query := `
	SELECT dms.id, dms.workspace_id
	FROM dms
	INNER JOIN user_dms ON dms.id = user_dms.dm_id
	WHERE user_dms.user_id = ?`

	rows, err := (*conn).Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dms []DM
	for rows.Next() {
		var dm DM
		if err := rows.Scan(&dm.Id, &dm.WorkspaceId); err != nil {
			return nil, err
		}
		dms = append(dms, dm)
	}

	return dms, nil
}

func GetDMsByUserAndWorkspace(user user.User, workspaceId int) ([]DM, error) {
	conn := database.GetConnection()
	defer conn.Close()

	query := `
	SELECT dms.id, dms.workspace_id
	FROM dms
	INNER JOIN user_dms ON dms.id = user_dms.dm_id
	WHERE user_dms.user_id = ? AND dms.workspace_id = ?`

	rows, err := (*conn).Query(query, user.Id, workspaceId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dms []DM
	for rows.Next() {
		var dm DM
		if err := rows.Scan(&dm.Id, &dm.WorkspaceId); err != nil {
			return nil, err
		}
		dms = append(dms, dm)
	}

	return dms, nil
}

func (dm DM) Join(userId int) error {
	conn := database.GetConnection()
	defer conn.Close()

	relationship := userDMRelationship.UserDMRelationship{
		UserId: userId,
		DMId:   dm.Id,
	}

	if relationship.Exists() {
		return nil
	}

	return relationship.Create()
}
