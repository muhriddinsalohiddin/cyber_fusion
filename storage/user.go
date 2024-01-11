package storage

import (
	"app/models"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	db *sql.DB
}

type UserStorageImpl struct {
	DB *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorageImpl {
	return &UserStorageImpl{DB: db}
}

func NewUser(db *sql.DB) *User {
	return &User{db: db}
}

func (r *User) Create(u *models.User) error {
	_, err := r.db.Exec(`insert into "user" (id, name, gender, birthday, email, login, password, bio) values ($1, $2,$3,$4,$5,$6,$7,$8)`,
		uuid.NewString(),
		u.Name,
		u.Gender,
		u.Birthday,
		u.Email,
		u.Login,
		u.Password,
		u.Bio,
	)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Update(u1 models.User, id string) error {
	_, err := u.db.Exec(`UPDATE "user" SET name=$1, 
gender=$2, 
birthday=$3, 
email=$4, 
login=$5, 
password=$6, 
bio=$7,
updated_at=now() where id=$8`, u1.Name, u1.Gender, u1.Birthday, u1.Email, u1.Login, u1.Password, u1.Bio, id)
	if err != nil {
		return err
	}
	return nil

}

func (d *User) Delete(id string) error {
	_, err := d.db.Exec(`DELETE FROM "user" WHERE id=$1`, id)
	if err != nil {
		return fmt.Errorf("User delete funksiyada xatolik kettiyou brat" + err.Error())
	}
	return nil
}

func (n *User) GetById(id string) (*models.User, error) {
	var (
		updatedAt     sql.NullString
		u             models.User
		messages      string
		notifications string
		posts         string
		query         = `
		SELECT 
    u.id, 
    u.name, 
    u.gender, 
    u.birthday, 
    u.email, 
    u.login, 
    u.password, 
    u.bio, 
    to_char(u.created_at, 'YYYY-MM-DD') as created_at,
    to_char(u.updated_at, 'YYYY-MM-DD') as updated_at,
    COALESCE(
      (
        SELECT 
        json_agg(
          json_build_object(
            'id', m.id,
            'sender_id', m.sender_id,
            'receiver_id', m.receiver_id,
            'body', m.body,
            'created_at', to_char(m.created_at, 'YYYY-MM-DD'),
            'updated_at', to_char(m.updated_at, 'YYYY-MM-DD')
          )
        )
        FROM "message" m
        WHERE u.id = m.sender_id 
      ),
      '[]'
    ) as messages, 
    COALESCE(
      (
        SELECT 
        json_agg(
        	json_build_object(
            'id', n.id,
            'user_id', n.user_id,
            'type', n.type,
            'body', n.body,
            'created_at', to_char(n.created_at, 'YYYY-MM-DD')
          )
        )
        FROM "notification" n
        WHERE u.id = n.user_id 
      ),
      '[]'
    ) as notifications, 
    COALESCE(
      (
      	SELECT 
        json_agg(
          json_build_object(
        	  'id', p.id, 
        	  'user_id', p.user_id, 
        	  'title', p.title, 
        	  'body', p.body, 
        	  'created_at', to_char(p.created_at, 'YYYY-MM-DD'),
        	  'updated_at', to_char(p.updated_at, 'YYYY-MM-DD'),

        	  'likes', COALESCE(
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
            ),

            'comments', COALESCE(
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
            )
						
          )
        )
        FROM "post" p
        WHERE u.id = p.user_id 
      ),
      '[]'
    ) as posts 
		FROM 
    	"user" u
		WHERE 
    	u.id = $1
		GROUP BY  
    	u.id;
	`
	)

	err := n.db.QueryRow(query, id).Scan(
		&u.Id,
		&u.Name,
		&u.Gender,
		&u.Birthday,
		&u.Email,
		&u.Login,
		&u.Password,
		&u.Bio,
		&u.CreatedAt,
		&updatedAt,
		&messages,
		&notifications,
		&posts,
	)

	if err != nil {
		return nil, err
	}

	if updatedAt.Valid {
		u.UpdetadAt = updatedAt.String
	}

	// Unmarshal the messages string into []*models.Message
	err = json.Unmarshal([]byte(messages), &u.Messages)
	if err != nil {
		return nil, err
	}

	// Unmarshal the notifications string into []*models.Message
	err = json.Unmarshal([]byte(notifications), &u.Notifications)
	if err != nil {
		return nil, err
	}

	// Unmarshal the posts string into []*models.Post
	err = json.Unmarshal([]byte(posts), &u.Posts)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (n *User) Get(req *models.UserReq) (*models.Users, error) {
	var (
		updatedAt sql.NullString
		m         models.Users
		filter    = " WHERE true "
		end       = " GROUP BY  u.id "
		arr       []any
	)

	if req.FromDate != "" && req.ToDate != "" {
		arr = append(arr, req.FromDate, req.ToDate)
		filter += fmt.Sprintf(" AND created_at BETWEEN $%d AND $%d", len(arr)-1, len(arr))
	}

	if req.Limit != 0 {
		arr = append(arr, req.Limit)
		filter += " LIMIT $" + fmt.Sprint(len(arr))
	}

	if req.Offset != 0 {
		arr = append(arr, req.Offset)
		filter += " OFFSET $" + fmt.Sprint(len(arr))
	}

	query := `
	SELECT 
    u.id, 
    u.name, 
    u.gender, 
    u.birthday, 
    u.email, 
    u.login, 
    u.password, 
    u.bio, 
    to_char(u.created_at, 'YYYY-MM-DD') as created_at,
    to_char(u.updated_at, 'YYYY-MM-DD') as updated_at,
    COALESCE(
      (
        SELECT 
        json_agg(
          json_build_object(
            'id', m.id,
            'sender_id', m.sender_id,
            'receiver_id', m.receiver_id,
            'body', m.body,
            'created_at', to_char(m.created_at, 'YYYY-MM-DD'),
            'updated_at', to_char(m.updated_at, 'YYYY-MM-DD')
          )
        )
        FROM "message" m
        WHERE u.id = m.sender_id 
      ),
      '[]'
    ) as messages, 
    COALESCE(
      (
        SELECT 
        json_agg(
        	json_build_object(
            'id', n.id,
            'user_id', n.user_id,
            'type', n.type,
            'body', n.body,
            'created_at', to_char(n.created_at, 'YYYY-MM-DD')
          )
        )
        FROM "notification" n
        WHERE u.id = n.user_id 
      ),
      '[]'
    ) as notifications, 
    COALESCE(
      (
      	SELECT 
        json_agg(
          json_build_object(
        	  'id', p.id, 
        	  'user_id', p.user_id, 
        	  'title', p.title, 
        	  'body', p.body, 
        	  'created_at', to_char(p.created_at, 'YYYY-MM-DD'),
        	  'updated_at', to_char(p.updated_at, 'YYYY-MM-DD'),

        	  'likes', COALESCE(
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
            ),

            'comments', COALESCE(
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
            )
          )
        )
        FROM "post" p
        WHERE u.id = p.user_id 
      ),
      '[]'
    ) as posts 
		FROM 
    	"user" u
	` + filter + end
	rows, err := n.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var (
			u             models.User
			messages      string
			notifications string
			posts         string
		)
		err = rows.Scan(
			&u.Id,
			&u.Name,
			&u.Gender,
			&u.Birthday,
			&u.Email,
			&u.Login,
			&u.Password,
			&u.Bio,
			&u.CreatedAt,
			&updatedAt,
			&messages,
			&notifications,
			&posts,
		)
		if err != nil {
			return nil, fmt.Errorf("Get dagi rowsni scan qilishda xatolik bor ekan xatolik: " + err.Error())
		}

		if updatedAt.Valid {
			u.UpdetadAt = updatedAt.String
		} else {
			u.UpdetadAt = "bu profilda yangilanish kiritilmagan"
		}
		// Unmarshal the messages string into []*models.Message
		err = json.Unmarshal([]byte(messages), &u.Messages)
		if err != nil {
			return nil, err
		}

		// Unmarshal the notification string into []*models.Notification
		err = json.Unmarshal([]byte(notifications), &u.Notifications)
		if err != nil {
			return nil, err
		}

		// Unmarshal the posts string into []*models.Post
		err = json.Unmarshal([]byte(posts), &u.Posts)
		if err != nil {
			return nil, err
		}

		m.Users = append(m.Users, &u)
	}

	return &m, nil
}
