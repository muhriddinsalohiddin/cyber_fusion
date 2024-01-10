package storage

import (
	"app/models"
	"database/sql"
	"encoding/json"
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
				id,
				user_id,
				title,
				body
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
	var (
		m     models.PostListResp
		query = `
		SELECT 
			p.id, 
			p.user_id, 
			p.title, 
			p.body, 
			to_char(p.created_at, 'YYYY-MM-DD') as created_at,
			to_char(p.updated_at, 'YYYY-MM-DD') as updated_at,
			COALESCE(
				(
					SELECT 
					json_agg(
						json_build_object(
							'id', l.id,
							'user_id', l.user_id,
							'post_id', l.post_id,
							'created_at', to_char(l.created_at, 'YYYY-MM-DD')
						)
					)
				FROM "like" l
				WHERE p.id = l.post_id
				),
				'[]'
			) as likes , 
			COALESCE(
				(
					SELECT 
					json_agg(
						json_build_object(
							'id', c.id,
							'user_id', c.user_id,
							'post_id', c.post_id,
							'body', c.body,
							'created_at', to_char(c.created_at, 'YYYY-MM-DD'),
							'updated_at', to_char(c.created_at, 'YYYY-MM-DD')
						)
					)
					FROM "comment" c
					WHERE p.id = c.post_id
				),
				'[]'
			) as comments 
		FROM 
			"post" p
		GROUP BY  
			p.id;
	`
	)
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
			updateAt sql.NullString
			likes    string
			comments string
			b        models.Post
		)
		err = rows.Scan(
			&b.Id,
			&b.UserId,
			&b.Title,
			&b.Body,
			&b.CreatedAt,
			&updateAt,
			&likes,
			&comments,
		)
		if err != nil {
			return nil, err
		}

		if updateAt.Valid {
			b.UpdatedAt = updateAt.String
		}

		err = json.Unmarshal([]byte(likes), &b.Likes)
		if err != nil {
			return nil, err
		}
		
		err = json.Unmarshal([]byte(comments), &b.Comments)
		if err != nil {
			return nil, err
		}

		m.Post = append(m.Post, &b)
	}

	err = r.db.QueryRow(`SELECT COUNT(1) FROM "post"`).Scan(&m.Count)

	return &m, err
}

func (r *Post) GetByIdPost(id string) (*models.Post, error) {
	var (
		likes    string
		comments string
		updateAt sql.NullString
		post     models.Post
		query    = `
			SELECT 
  			p.id, 
  			p.user_id, 
  			p.title, 
  			p.body, 
  			to_char(p.created_at, 'YYYY-MM-DD') as created_at,
  			to_char(p.updated_at, 'YYYY-MM-DD') as updated_at,
				COALESCE(
  				(
    				SELECT 
    				json_agg(
      				json_build_object(
        				'id', l.id,
        				'user_id', l.user_id,
        				'post_id', l.post_id,
        				'created_at', to_char(l.created_at, 'YYYY-MM-DD')
      				)
    				)
    			FROM "like" l
    			WHERE p.id = l.post_id
  				),
  				'[]'
				) as likes , 
				COALESCE(
					(
						SELECT 
						json_agg(
							json_build_object(
								'id', c.id,
								'user_id', c.user_id,
								'post_id', c.post_id,
								'body', c.body,
								'created_at', to_char(c.created_at, 'YYYY-MM-DD'),
								'updated_at', to_char(c.created_at, 'YYYY-MM-DD')
							)
						)
						FROM "comment" c
						WHERE p.id = c.post_id
					),
					'[]'
				) as comments 
			FROM 
  			"post" p
			WHERE 
  			p.id = $1
			GROUP BY  
  			p.id;
		`
	)
	err := r.db.QueryRow(query, id).Scan(
		&post.Id,
		&post.UserId,
		&post.Title,
		&post.Body,
		&post.CreatedAt,
		&updateAt,
		&likes,
		&comments,
	)

	if err != nil {
		return nil, err
	}
	if updateAt.Valid {
		post.UpdatedAt = updateAt.String
	}

	// Unmarshal the likes string into []*models.Message
	err = json.Unmarshal([]byte(likes), &post.Likes)
	err = json.Unmarshal([]byte(comments), &post.Comments)

	return &post, err
}
