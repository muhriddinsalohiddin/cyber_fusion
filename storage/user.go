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

func (n *User) GetByIdWithAllItems(id string) (*models.User, error) {
	var (
		updatedAt    sql.NullString
		u            models.User
		notification string
		message      string
		comment      string
		like         string
		posts        string
	)
	query := `
	select
		u.id,
		u.name,
		u.gender,
		u.birthday,
		u.email,
		u.login,
		u.password,
		u.bio,
		u.created_at,
		u.updated_at,
		json_agg(n.*)::TEXT as notification,
		json_agg(c.*)::TEXT as comment,
		json_agg(m.*)::TEXT as message,
		json_agg(l.*)::TEXT as likes,
		json_agg(post.*)::TEXT as posts
	from 
	"user" u
	join notification n on n.user_id = u.id
	join "comment" c on c.user_id = u.id
	join message m on m.sender_id = u.id
	join "like" l on l.user_id = u.id
	join lateral (
		select
			p.id,
			p.user_id,
			p.title ,
			p.body ,
			p.created_at ,
			p.updated_at ,
			json_agg(l.*) as likes,
			json_agg(c.*) as "comments"
		from post p 
		join "like" l on l.post_id  = p.id
		join "comment" c on c.post_id = p.id 
		where p.user_id = u.id
		group by p.id
	) as post on true
	where u.id = $1
	group by u.id`

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
		&notification,
		&comment,
		&message,
		&like,
		&posts,
	)
	if err != nil {
		return nil, fmt.Errorf("GetByIdWithAllItems dagi queryda xatolik bor ekan xatolik: " + err.Error())
	}

	err = json.Unmarshal([]byte(notification), &u.Notifications)
	if err != nil {
		return nil, fmt.Errorf("GetByIdWithAllItems dagi UnmarshalNotificationda xatolik bor ekan xatolik: " + err.Error())
	}

	err = json.Unmarshal([]byte(comment), &u.Comments)
	if err != nil {
		return nil, fmt.Errorf("GetByIdWithAllItems dagi UnmarshalCommentda xatolik bor ekan xatolik: " + err.Error())
	}

	err = json.Unmarshal([]byte(message), &u.Message)
	if err != nil {
		return nil, fmt.Errorf("GetByIdWithAllItems dagi UnmarshalMessageda xatolik bor ekan xatolik: " + err.Error())
	}

	err = json.Unmarshal([]byte(like), &u.Likes)
	if err != nil {
		return nil, fmt.Errorf("GetByIdWithAllItems dagi UnmarshalLikedda xatolik bor ekan xatolik: " + err.Error())
	}

	err = json.Unmarshal([]byte(posts), &u.Posts)
	if err != nil {
		return nil, fmt.Errorf("GetByIdWithAllItems dagi UnmarshalPostda xatolik bor ekan xatolik: " + err.Error())
	}

	u.UpdetadAt = updatedAt.String
	return &u, err
}
func (n *User) GetById(id string) (*models.User, error) {
	var updatedAt sql.NullString
	var u models.User
	err := n.db.QueryRow(`SELECT id, 
	name, 
	gender, 
	birthday, 
	email, 
	login, 
	password, 
	bio, 
	to_char(created_at, 'YYYY-MM-DD') as created_at,
	to_char(updated_at, 'YYYY-MM-DD') as updated_at 
	FROM 
	"user" 
	where id=$1`, id).Scan(
		&u.Id,
		&u.Name,
		&u.Gender,
		&u.Birthday,
		&u.Email,
		&u.Login,
		&u.Password,
		&u.Bio,
		&u.CreatedAt,
		&updatedAt)

	if updatedAt.Valid {
		u.UpdetadAt = updatedAt.String
	}

	return &u, err

}
func (n *User) Get(req *models.UserReq) (*models.Users, error) {
	var (
		updatedAt sql.NullString
		m         models.Users
		filter    = "WHERE true"
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
	SELECT id, 
		name, 
		gender, 
		birthday, 
		email, 
		login, 
		password, 
		bio, 
		to_char(created_at, 'YYYY-MM-DD') as created_at,
		to_char(updated_at, 'YYYY-MM-DD') as updated_at 
	FROM "user" ` + filter
	fmt.Println("query", query, arr)
	rows, err := n.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var u models.User

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
		)
		if err != nil {
			return nil, fmt.Errorf("Get dagi rowsni scan qilishda xatolik bor ekan xatolik: " + err.Error())
		}

		if updatedAt.Valid {
			u.UpdetadAt = updatedAt.String
		} else {
			u.UpdetadAt = "bu profilda yangilanish kiritilmagan"
		}

		m.Users = append(m.Users, &u)
	}

	return &m, nil
}
