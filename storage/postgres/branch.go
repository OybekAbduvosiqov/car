package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"app/models"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/google/uuid"
)

type BranchRepo struct {
	db *pgxpool.Pool
}

func NewBranchRepo(db *pgxpool.Pool) *BranchRepo {
	return &BranchRepo{
		db: db,
	}
}

func (r *BranchRepo) Insert(ctx context.Context, Branch *models.CreateBranch) (string, error) {

	var (
		id = uuid.New().String()
	)

	query := `
		INSERT INTO branch (
			id,
			name,
			updated_at
		) VALUES ($1, $2, now())
	`

	_, err := r.db.Exec(ctx, query,
		id,
		Branch.Name,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *BranchRepo) GetByID(ctx context.Context, req *models.BranchPrimeryKey) (*models.Branch, error) {

	var (
		id        sql.NullString
		name      sql.NullString
		createdAt sql.NullString
		updatedAt sql.NullString
	)

	query := `
		SELECT
			id,
			name,
			created_at,
			updated_at
		FROM branch
		WHERE id = $1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.Branch{
		Id:        id.String,
		Name:      name.String,
		CreatedAt: createdAt.String,
		UpdatedAt: updatedAt.String,
	}, err
}

func (r *BranchRepo) GetList(ctx context.Context, req *models.GetListBranchRequest) (*models.GetListBranchResponse, error) {
	var (
		offset = "OFFSET 0"
		limit  = "LIMIT 10"
		resp   = &models.GetListBranchResponse{}
	)

	if req.Offset > 0 {
		offset = fmt.Sprintf("OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf("LIMIT %d", req.Limit)
	}

	query := `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			created_at,
			updated_at
		FROM branch
	`

	query += offset + limit

	rows, err := r.db.Query(ctx, query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {

		var (
			id        sql.NullString
			name      sql.NullString
			createdAt sql.NullString
			updatedAt sql.NullString
		)

		err = rows.Scan(
			&resp.Count,
			&id,
			&name,
			&createdAt,
			&updatedAt,
		)

		resp.Branchs = append(resp.Branchs, &models.Branch{
			Id:        id.String,
			Name:      name.String,
			CreatedAt: createdAt.String,
			UpdatedAt: updatedAt.String,
		})
	}

	return resp, err
}

func (r *BranchRepo) Update(ctx context.Context, Branch *models.UpdateBranch) error {
	query := `
		UPDATE
			branch
		SET
			name = $2,
			updated_at = now()
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query,
		Branch.Id,
		Branch.Name,
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *BranchRepo) Delete(ctx context.Context, req *models.BranchPrimeryKey) error {

	_, err := r.db.Exec(ctx, "delete from investor where id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
