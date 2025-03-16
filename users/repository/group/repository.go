package group

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	groupentity "github.com/turao/topics/users/entity/group"
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

func (r *repository) Save(ctx context.Context, group groupentity.Group) error {
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

func (r *repository) FindByID(ctx context.Context, groupID groupentity.ID) (groupentity.Group, error) {
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

	var groupMemberModels []GroupMemberModel
	err = r.database.SelectContext(
		ctx,
		&groupMemberModels,
		"SELECT * FROM group_member WHERE group_id = $1 AND group_version = $2",
		groupModel.ID,
		groupModel.Version,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return ToEntity(groupModel, groupMemberModels)
}

func (r *repository) FindByMemberID(ctx context.Context, memberID groupentity.MemberID) (map[groupentity.ID]struct{}, error) {
	var groupMemberModels []GroupMemberModel
	err := r.database.SelectContext(
		ctx,
		&groupMemberModels,
		`
		WITH latest_group_version AS (
			SELECT id, max(version) as version 
			FROM groups GROUP BY id
		)
		
		SELECT gm.*
		FROM group_member gm
		JOIN latest_group_version g ON g.id = gm.group_id AND g.version = gm.group_version
		WHERE gm.member_id = $1
		`,
		memberID,
	)
	if err != nil {
		return nil, err
	}

	groupIDs := make(map[groupentity.ID]struct{}, len(groupMemberModels))
	for _, groupMemberModel := range groupMemberModels {
		groupID := groupentity.ID(groupMemberModel.GroupID)
		groupIDs[groupID] = struct{}{}
	}

	return groupIDs, nil
}
