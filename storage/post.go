package storage

import (
	"app/models"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Post struct {
	db *sql.DB
}

func NewPost(db *sql.DB) *Post {
	return &Post{db: db}
}

func (r *Post) Create(m *models.Post) error {
	_, err := r.db.Exec(`
		INSERT INTO 
			"post"(
				id,user_id,title,body
			) 
		VALUES(
			$1,$2,$3,$4
		)
	`, uuid.NewString(), m.UserId, m.Title, m.Body)
	return err
}

func (r *Post) Delete(m *models.Post) error {
	_, err := r.db.Exec(`
	DELETE FROM 
		"post"
	WHERE
		id=$1
	`, m.Id)

	return err
}

func (r *Post) Update(n *models.Post) error {
	_, err := r.db.Exec(`
	UPDATE 
		"post"
	SET
		body=$2,
		title=$3,
		updated_at = NOW()
	WHERE
		id=$1
	`, n.Id, n.Body, n.Title)
	if err != nil {
		return fmt.Errorf("PUF xato " + err.Error())
	}
	return nil
}

func (r *Post) GetPostlist() (*models.PostListResp, error) {
	var m models.PostListResp
	query := `
	SELECT
		id,
		user_id,
		title,
		body,
		created_at
	FROM "post"`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("queryda ")
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			fmt.Println("aka kanal yopilmadi", err)
		}
	}()

	for rows.Next() {
		var (
			b models.Post
		)
		err = rows.Scan(
			&b.Id,
			&b.UserId,
			&b.Title,
			&b.Body,
			&b.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		m.Post = append(m.Post, &b)
	}

	err = r.db.QueryRow(`SELECT COUNT(1) FROM "post"`).Scan(&m.Count)

	return &m, err
}
func (r *Post) GetlistWithComments(user_id string) (*models.PostListResp, error) {
	var (
		posts models.PostListResp
	)

	query := `
	select 
		id,
		title,
		user_id,
		body,
		created_at,
		updated_at
	from post p
	where p.user_id = $1`

	rows, err := r.db.Query(query, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			p          models.Post
			updated_at sql.NullString
		)
		err = rows.Scan(
			&p.Id,
			&p.Title,
			&p.UserId,
			&p.Body,
			&p.CreatedAt,
			&updated_at,
		)
		if err != nil {
			return nil, err
		}
		p.UpdatedAt = updated_at.String

		// Comment Started
		queryComment := `
		select
			id,
			user_id,
			post_id,
			body,
			created_at,
			updated_at
		from comment c
		where c.post_id = $1`

		rowsComment, err := r.db.Query(queryComment, p.Id)
		if err != nil {
			return nil, err
		}
		defer rowsComment.Close()

		for rowsComment.Next() {
			var (
				c         models.Comment
				updatedAt sql.NullString
			)
			err = rowsComment.Scan(
				&c.Id,
				&c.UserId,
				&c.PostId,
				&c.Body,
				&c.CreatedAt,
				&updatedAt,
			)
			if err != nil {
				return nil, err
			}
			c.UpdatedAt = updatedAt.String
			p.Comments = append(p.Comments, &c)
		}

		// Like Started
		queryLike := `
		select
			id,
			user_id,
			post_id,
			created_at
		from "like" l
		where l.post_id = $1`

		rowsLike, err := r.db.Query(queryLike, p.Id)
		if err != nil {
			return nil, err
		}
		defer rowsLike.Close()

		for rowsLike.Next() {
			var (
				l models.Like
			)
			err = rowsLike.Scan(
				&l.Id,
				&l.UserId,
				&l.PostId,
				&l.CreatedAt,
			)
			if err != nil {
				return nil, err
			}
			p.Likes = append(p.Likes, &l)
		}

		posts.Post = append(posts.Post, &p)
	}

	return &posts, nil
}
func (r *Post) GetByIdPost(id string) (*models.Post, error) {
	var post models.Post
	query := `
	SELECT
		id,
		user_id,
		title,
		body,
		created_at
	FROM "post"
	WHERE id = $1
	`

	err := r.db.QueryRow(query, id).Scan(
		&post.Id,
		&post.UserId,
		&post.Title,
		&post.Body,
		&post.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &post, nil
}
