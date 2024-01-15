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
	if err != nil {
		return fmt.Errorf("PDF xato" + err.Error())
	}
	return nil
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



// func(r *Post) getPostData(postID string) (Post, error) {
// 	var p models.Post

// 	err := r.db.QueryRow(`
// 		SELECT
// 			p.id AS post_id,
// 			p.user_id,
// 			p.title AS post_title,
// 			p.body AS post_body,
// 			p.created_at AS post_created_at,
// 			p.updated_at AS post_updated_at,
// 			c.id AS comment_id,
// 			c.user_id AS comment_user_id,
// 			c.body AS comment_body,
// 			c.created_at AS comment_created_at,
// 			c.updated_at AS comment_updated_at,
// 			l.id AS like_id,
// 			l.user_id AS like_user_id,
// 			l.created_at AS like_created_at
// 		FROM "post" p
// 		LEFT JOIN "comment" c ON p.id = c.post_id
// 		LEFT JOIN "like" l ON p.id = l.post_id
// 		WHERE p.id = $1
// 	`, postID).Scan(
// 		p.Id,
// 		p.UserId,
// 		p.Title,
// 		p.Body,
// 		p.CreatedAt,
// 		p.UpdatedAt,
// 		p.Comment[0].Id
// 		p.Like[0].Id
// 		&post.ID,
// 		&post.UserID,
// 		&post.Title,
// 		&post.Body,
// 		&post.CreatedAt,
// 		&post.UpdatedAt,
// 		&post.Comments[0].ID,
// 		&post.Comments[0].UserID,
// 		&post.Comments[0].Body,
// 		&post.Comments[0].CreatedAt,
// 		&post.Comments[0].UpdatedAt,
// 		&post.like[0].ID,
// 		&post.Likes[0].UserID,
// 		&post.Likes[0].CreatedAt,
// 	)
// 	if err != nil {
// 		return post, err
// 	}

// 	return post, nil
// }