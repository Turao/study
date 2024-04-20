package group

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/turao/topics/users/entity/group"
)

var (
	ErrNotFound = errors.New("not found")
)

type repository struct {
	database *sqlx.DB
}

func NewRepository(database *sqlx.DB) (*repository, error) {
	if database == nil {
		return nil, errors.New("database connection is nil")
	}

	return &repository{
		database: database,
	}, nil
}

func (r *repository) Save(ctx context.Context, group group.Group) error {
	groupModel, err := ToGroupModel(group)
	if err != nil {
		return err
	}

	groupMemberModels, err := ToGroupMemberModels(group)
	if err != nil {
		return err
	}

	transaction, err := r.database.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	_, err = transaction.NamedExecContext(
		ctx,
		`INSERT INTO groups (id, version, name, tenancy, created_at, deleted_at)
		VALUES (:id, :version, :name, :tenancy, :created_at, :deleted_at)`,
		groupModel,
	)
	if err != nil {
		transaction.Rollback()
		return err
	}

	for _, groupMemberModel := range groupMemberModels {
		_, err = transaction.NamedExecContext(
			ctx,
			`INSERT INTO group_member (group_id, group_version, member_id)
			VALUES (:group_id, :group_version, :member_id)`,
			groupMemberModel,
		)
		if err != nil {
			transaction.Rollback()
			return err
		}
	}

	err = transaction.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindByID(ctx context.Context, groupID group.ID) (group.Group, error) {
	var groupModel GroupModel
	err := r.database.GetContext(
		ctx,
		&groupModel,
		"SELECT * FROM groups WHERE id = $1 ORDER BY version DESC LIMIT 1",
		groupID,
	)
	if err != nil {
		return nil, err
	}

	return ToEntity(groupModel, []GroupMemberModel{})
}
