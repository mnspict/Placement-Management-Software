// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const extraInfoCompany = `-- name: ExtraInfoCompany :one
INSERT INTO companies (company_name, company_email, representative_contact, representative_name, data_url, user_id)
VALUES ($1, $2, $3, $4, $5, (SELECT user_id FROM users WHERE email = $6))
RETURNING company_id, company_name, company_email, representative_contact, representative_name, data_url, user_id, is_verified
`

type ExtraInfoCompanyParams struct {
	CompanyName           string
	CompanyEmail          string
	RepresentativeContact string
	RepresentativeName    string
	DataUrl               pgtype.Text
	Email                 string
}

func (q *Queries) ExtraInfoCompany(ctx context.Context, arg ExtraInfoCompanyParams) (Company, error) {
	row := q.db.QueryRow(ctx, extraInfoCompany,
		arg.CompanyName,
		arg.CompanyEmail,
		arg.RepresentativeContact,
		arg.RepresentativeName,
		arg.DataUrl,
		arg.Email,
	)
	var i Company
	err := row.Scan(
		&i.CompanyID,
		&i.CompanyName,
		&i.CompanyEmail,
		&i.RepresentativeContact,
		&i.RepresentativeName,
		&i.DataUrl,
		&i.UserID,
		&i.IsVerified,
	)
	return i, err
}

const getAll = `-- name: GetAll :many
SELECT user_id, email, password, role, user_uuid, created_at, confirmed FROM users
`

func (q *Queries) GetAll(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, getAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.UserID,
			&i.Email,
			&i.Password,
			&i.Role,
			&i.UserUuid,
			&i.CreatedAt,
			&i.Confirmed,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserData = `-- name: GetUserData :one
SELECT user_id, email, password, role, user_uuid, created_at, confirmed FROM users WHERE email = $1
`

func (q *Queries) GetUserData(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserData, email)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.Password,
		&i.Role,
		&i.UserUuid,
		&i.CreatedAt,
		&i.Confirmed,
	)
	return i, err
}

const getVerificationCompany = `-- name: GetVerificationCompany :one

SELECT is_verified FROM companies WHERE company_email = $1
`

// -- name: GetVerificationStudent :one
// SELECT is_verified FROM students WHERE email = $1;
func (q *Queries) GetVerificationCompany(ctx context.Context, companyEmail string) (pgtype.Bool, error) {
	row := q.db.QueryRow(ctx, getVerificationCompany, companyEmail)
	var is_verified pgtype.Bool
	err := row.Scan(&is_verified)
	return is_verified, err
}

const insertNewJob = `-- name: InsertNewJob :one




INSERT INTO jobs (data_url, company_id, title, location, type, salary, skills, position, extras)
VALUES ($1, (SELECT company_id FROM companies WHERE company_email = $2), $3, $4, $5, $6, $7, $8, $9)
RETURNING job_id, data_url, created_at, company_id, title, location, type, salary, skills, position, extras
`

type InsertNewJobParams struct {
	DataUrl      pgtype.Text
	CompanyEmail string
	Title        string
	Location     string
	Type         string
	Salary       string
	Skills       []string
	Position     string
	Extras       []byte
}

// -- name: GetVerificationAdmin :one
// SELECT is_verified FROM admins WHERE email = $1;
// -- name: GetVerificationSuperUser :one
// SELECT is_verified FROM superusers WHERE email = $1;
// >>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
// Company queries
func (q *Queries) InsertNewJob(ctx context.Context, arg InsertNewJobParams) (Job, error) {
	row := q.db.QueryRow(ctx, insertNewJob,
		arg.DataUrl,
		arg.CompanyEmail,
		arg.Title,
		arg.Location,
		arg.Type,
		arg.Salary,
		arg.Skills,
		arg.Position,
		arg.Extras,
	)
	var i Job
	err := row.Scan(
		&i.JobID,
		&i.DataUrl,
		&i.CreatedAt,
		&i.CompanyID,
		&i.Title,
		&i.Location,
		&i.Type,
		&i.Salary,
		&i.Skills,
		&i.Position,
		&i.Extras,
	)
	return i, err
}

const signupUser = `-- name: SignupUser :one
INSERT INTO users (email, password, role) VALUES ($1, $2, $3)
RETURNING user_id, email, password, role, user_uuid, created_at, confirmed
`

type SignupUserParams struct {
	Email    string
	Password string
	Role     int64
}

func (q *Queries) SignupUser(ctx context.Context, arg SignupUserParams) (User, error) {
	row := q.db.QueryRow(ctx, signupUser, arg.Email, arg.Password, arg.Role)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Email,
		&i.Password,
		&i.Role,
		&i.UserUuid,
		&i.CreatedAt,
		&i.Confirmed,
	)
	return i, err
}

const updateEmailConfirmation = `-- name: UpdateEmailConfirmation :exec
UPDATE users
SET confirmed = true
WHERE email = $1
`

func (q *Queries) UpdateEmailConfirmation(ctx context.Context, email string) error {
	_, err := q.db.Exec(ctx, updateEmailConfirmation, email)
	return err
}

const updatePassword = `-- name: UpdatePassword :exec
UPDATE users
SET password = $2
WHERE email = $1
`

type UpdatePasswordParams struct {
	Email    string
	Password string
}

func (q *Queries) UpdatePassword(ctx context.Context, arg UpdatePasswordParams) error {
	_, err := q.db.Exec(ctx, updatePassword, arg.Email, arg.Password)
	return err
}
