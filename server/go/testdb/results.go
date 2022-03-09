package testdb

import (
	"context"
	"encoding/json"
	"fmt"

	openapi "github.com/kubeshop/tracetest/server/go"
)

func (td *TestDB) CreateResult(ctx context.Context, testid string, run *openapi.Result) error {
	stmt, err := td.db.Prepare("INSERT INTO results(id, testid, result) VALUES( $1, $2, $3 )")
	if err != nil {
		return fmt.Errorf("sql prepare: %w", err)
	}
	defer stmt.Close()
	b, err := json.Marshal(run)
	if err != nil {
		return fmt.Errorf("json Marshal: %w", err)
	}
	_, err = stmt.ExecContext(ctx, run.Id, testid, b)
	if err != nil {
		return fmt.Errorf("sql exec: %w", err)
	}

	return nil
}

func (td *TestDB) GetResult(ctx context.Context, id string) (*openapi.Result, error) {
	stmt, err := td.db.Prepare("SELECT result FROM results WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var b []byte
	err = stmt.QueryRowContext(ctx, id).Scan(&b)
	if err != nil {
		return nil, err
	}
	var run openapi.Result

	err = json.Unmarshal(b, &run)
	if err != nil {
		return nil, err
	}
	return &run, nil
}

func (td *TestDB) GetResultsByTestID(ctx context.Context, testID string) ([]openapi.Result, error) {
	stmt, err := td.db.Prepare("SELECT result FROM results WHERE testid = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, testID) //.Scan(&b)
	if err != nil {
		return nil, err
	}
	var run []openapi.Result

	for rows.Next() {
		var b []byte
		if err := rows.Scan(&b); err != nil {
			return nil, err
		}
		var result openapi.Result
		err = json.Unmarshal(b, &result)
		if err != nil {
			return nil, err
		}

		run = append(run, result)
	}

	return run, nil
}

func (td *TestDB) UpdateResult(ctx context.Context, run *openapi.Result) error {
	stmt, err := td.db.Prepare("UPDATE results SET result = $2 WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	b, err := json.Marshal(run)
	if err != nil {
		return fmt.Errorf("json Marshal: %w", err)
	}
	_, err = stmt.ExecContext(ctx, run.Id, b)
	if err != nil {
		return fmt.Errorf("sql exec: %w", err)
	}

	return nil
}