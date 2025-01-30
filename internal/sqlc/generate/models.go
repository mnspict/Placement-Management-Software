// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Application struct {
	ApplicationID int64
	JobID         int64
	StudentID     int64
	DataUrl       pgtype.Text
	CreatedAt     pgtype.Timestamptz
	Status        interface{}
}

type Company struct {
	CompanyID             int64
	CompanyName           string
	RepresentativeEmail   string
	RepresentativeContact string
	RepresentativeName    string
	DataUrl               pgtype.Text
	UserID                int64
}

type Interview struct {
	InterviewID   int64
	ApplicationID int64
	CompanyID     int64
	DateTime      pgtype.Timestamptz
	Type          interface{}
	Status        interface{}
	Notes         pgtype.Text
	Location      string
	CreatedAt     pgtype.Timestamptz
	Extras        []byte
}

type Job struct {
	JobID        int64
	DataUrl      pgtype.Text
	CreatedAt    pgtype.Timestamp
	CompanyID    int64
	Title        string
	Location     string
	Type         string
	Salary       string
	Skills       []string
	Position     string
	Extras       []byte
	ActiveStatus bool
	Description  pgtype.Text
}

type Student struct {
	StudentID    int64
	StudentName  string
	RollNumber   string
	StudentDob   pgtype.Date
	Gender       string
	Course       string
	Department   string
	YearOfStudy  string
	ResumeUrl    pgtype.Text
	ResultUrl    string
	Cgpa         pgtype.Float8
	ContactNo    string
	StudentEmail string
	Address      pgtype.Text
	Skills       pgtype.Text
	UserID       int64
	Extras       []byte
	PictureUrl   pgtype.Text
}

type TempCorrectAnswer struct {
	QuestionID    string
	CorrectAnswer []string
	Points        pgtype.Int4
}

type Test struct {
	TestID       int64
	TestName     string
	Description  pgtype.Text
	Duration     int64
	QCount       int64
	EndTime      pgtype.Timestamptz
	Type         string
	UploadMethod interface{}
	JobID        pgtype.Int8
	CompanyID    int64
	FileID       string
	ResultUrl    pgtype.Text
	Threshold    int32
}

type Testresponse struct {
	ResponseID int64
	ResultID   int64
	QuestionID string
	Response   []string
	TimeTaken  pgtype.Int8
	Points     pgtype.Int4
	CreatedAt  pgtype.Timestamptz
}

type Testresult struct {
	ResultID  int64
	TestID    int64
	UserID    int64
	StartTime pgtype.Timestamptz
	EndTime   pgtype.Timestamptz
	Score     pgtype.Int8
}

type User struct {
	UserID     int64
	Email      string
	Password   string
	Role       int64
	UserUuid   pgtype.UUID
	CreatedAt  pgtype.Timestamp
	Confirmed  bool
	IsVerified bool
}
