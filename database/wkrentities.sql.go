// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.21.0
// source: wkrentities.sql

package database

import (
	"context"
	"database/sql"
)

const kvstoreGetLastStatus = `-- name: KvstoreGetLastStatus :one

WITH d AS (VALUES(1))
SELECT
	kvc.v AS next_page,
	kvt.v AS entity_ts
FROM d
LEFT JOIN kvstore kvc ON kvc.k = 'next-page'
LEFT JOIN kvstore kvt ON kvt.k = 'entity-ts'
`

type KvstoreGetLastStatusRow struct {
	NextPage sql.NullString
	EntityTs sql.NullString
}

func (q *Queries) KvstoreGetLastStatus(ctx context.Context) (KvstoreGetLastStatusRow, error) {
	row := q.db.QueryRowContext(ctx, kvstoreGetLastStatus)
	var i KvstoreGetLastStatusRow
	err := row.Scan(&i.NextPage, &i.EntityTs)
	return i, err
}

const kvstoreInsertEntityTs = `-- name: KvstoreInsertEntityTs :exec

INSERT INTO kvstore (k, v)
VALUES ('entity-ts', ?)
ON CONFLICT (k)
DO UPDATE
	SET v = EXCLUDED.v
`

func (q *Queries) KvstoreInsertEntityTs(ctx context.Context, v string) error {
	_, err := q.db.ExecContext(ctx, kvstoreInsertEntityTs, v)
	return err
}

const kvstoreInsertNextPage = `-- name: KvstoreInsertNextPage :exec

INSERT INTO kvstore (k, v)
VALUES ('next-page', ?)
ON CONFLICT (k)
DO UPDATE
	SET v = EXCLUDED.v
`

func (q *Queries) KvstoreInsertNextPage(ctx context.Context, v string) error {
	_, err := q.db.ExecContext(ctx, kvstoreInsertNextPage, v)
	return err
}

const reposInsert = `-- name: ReposInsert :exec

INSERT INTO repos (owner_name, repo_name, last_ts)
VALUES (?, ?, UNIXEPOCH())
ON CONFLICT (owner_name, repo_name)
DO NOTHING
`

type ReposInsertParams struct {
	OwnerName string
	RepoName  string
}

func (q *Queries) ReposInsert(ctx context.Context, arg ReposInsertParams) error {
	_, err := q.db.ExecContext(ctx, reposInsert, arg.OwnerName, arg.RepoName)
	return err
}
