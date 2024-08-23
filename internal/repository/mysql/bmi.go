package mysql

import (
	"context"
	"database/sql"
	"fmt"
)

type BMIRepository struct {
	DB *sql.DB
}

// NewMysqlAuthorRepository will create an implementation of author.Repository
func NewBMIRepository(db *sql.DB) *BMIRepository {
	return &BMIRepository{
		DB: db,
	}
}

func (b *BMIRepository) SaveBMI(ctx context.Context, name string, bmi float32) error {
	query := `insert into bmi_table (name, bmi) values (?, ?) `
	stmt, err := b.DB.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("prepare: %s", err.Error())
	}

	fmt.Println("name", name)
	fmt.Println("bmi", bmi)

	_, err = stmt.ExecContext(ctx, name, bmi)
	if err != nil {
		return fmt.Errorf("exec: %s", err.Error())
	}

	return nil
}