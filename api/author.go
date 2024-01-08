package api

import (
	"app/models"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (a *Api) CreateAuthor(c *fiber.Ctx) error {
	var b models.Author
	err := c.BodyParser(&b)
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "body parser da xatolik: "+err.Error())
	}
	fmt.Println(b)
	a.stg.Author.CreateAuthor(&b)
	return handlerResponse(c, http.StatusCreated, b)
}

// func (a *Api) CreateAuthor(c *fiber.Ctx) error {
// 	var b models.Author
// 	err := c.BodyParser(&b)
// 	if err != nil {
// 		return handlerResponse(c, http.StatusBadRequest, "body parser da xatolik: "+err.Error())
// 	}

// 	// Inserting the ID generation and creation logic here
// 	b.Id = uuid.NewString()

// 	// Database insertion logic
// 	_, err = a.db.Exec(`
// 		INSERT INTO "author" (
// 			id,
// 			name
// 		) VALUES (
// 			$1,$2
// 		)`, b.Id, b.Name)

// 	if err != nil {
// 		return handlerResponse(c, http.StatusInternalServerError, "Create qilishda xatolik: "+err.Error())
// 	}

// 	return handlerResponse(c, http.StatusCreated, b)
// }

func (a *Api) UpdateAuthor(c *fiber.Ctx) error {
	var m models.Author

	err := c.BodyParser(&m)
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "body parser da xatolik: "+err.Error())
	}

	id := c.Params("id")

	// if id == "" {
	// 	return handlerResponse(c, http.StatusBadRequest, "ID berilmadi"+err.Error())
	// }

	err = a.stg.Author.Update(&m, &id)
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, "Update storage da xatolik: "+err.Error())
	}

	return handlerResponse(c, http.StatusOK, "author updated successfully")
}

// func (a *Api) GetAuthorList(c *fiber.Ctx) error {
// 	var m models.AuthorList

// 	// Retrieve the list of authors
// 	query := `
// 		SELECT
// 			id,
// 			name,
// 			created_at,
// 			updated_at
// 		FROM "author"`

// 	rows, err := a.db.Query(query)
// 	if err != nil {
// 		return handlerResponse(c, http.StatusInternalServerError, "Get qilishda xatolik: "+err.Error())
// 	}

// 	defer func() {
// 		err = rows.Close()
// 		if err != nil {
// 			fmt.Println("aka kanal yopilmadi", err)
// 		}
// 	}()

// 	// Process each row and append to the list
// 	for rows.Next() {
// 		var (
// 			updated sql.NullString
// 			b       models.Author
// 		)
// 		err = rows.Scan(
// 			&b.Id,
// 			&b.Name,
// 			&b.CreatedAt,
// 			&updated,
// 		)
// 		if err != nil {
// 			return handlerResponse(c, http.StatusInternalServerError, "Scan qilishda xatolik: "+err.Error())
// 		}
// 		if updated.Valid {
// 			b.UpdatedAt = updated.String
// 		}

// 		m.Authors = append(m.Authors, &b)
// 	}

// 	// Retrieve the total count of authors
// 	err = a.db.QueryRow(`SELECT COUNT(1) FROM "author"`).Scan(&m.Count)
// 	if err != nil {
// 		return handlerResponse(c, http.StatusInternalServerError, "Count qilishda xatolik: "+err.Error())
// 	}

// 	return handlerResponse(c, http.StatusOK, m)
// }

func (a *Api) GetAuthorList(c *fiber.Ctx) error {
	// var l models.Author
	m, err := a.stg.Author.GetList()
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, "Get qilishda xatolik: "+err.Error())
	}

	return handlerResponse(c, http.StatusOK, m)
}

// func (a *Api) DeleteAuthor(c *fiber.Ctx) error {
// 	var k models.Author
// 	err := c.BodyParser(&k)
// 	if err != nil {
// 		return handlerResponse(c, http.StatusBadRequest, "Body parcer da xatolik"+err.Error())
// 	}

// 	// Delete the author directly within the Api type
// 	_, err = a.db.Exec(`
// 		DELETE FROM 
// 			"author" 
// 		WHERE 
// 			id = $1
// 		`, k.Id)

// 	if err != nil {
// 		return handlerResponse(c, http.StatusInternalServerError, "Delete qilishda xatolik: "+err.Error())
// 	}

// 	return handlerResponse(c, http.StatusAccepted, "Author deleted successfully")
// }

func (a *Api) DeleteAuthor(c *fiber.Ctx) error {
	var k models.Author
	err := c.BodyParser(&k)
	if err != nil {
		return handlerResponse(c, http.StatusBadRequest, "Body parcer da xatolik"+err.Error())
	}
	err = a.stg.Author.Delete(&k)
	if err != nil {
		return handlerResponse(c, http.StatusInternalServerError, "Delete qilishda xatolik: "+err.Error())
	}

	return handlerResponse(c, http.StatusAccepted, "Author deleted successfully")
}
