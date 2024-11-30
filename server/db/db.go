package db

import (
	"go-react-blog/model"
	"database/sql"
	"log"
)

type DB interface {
	GetTechnologies() ([]*model.Technology, error)
	GetBlogs() ([]*model.Blog, error)
    CreateBlog(blog *model.Blog) error
    UpdateBlog(id int, blog *model.Blog) error
    DeleteBlog(id int)  error
    GetBlog(id int )(*model.Blog, error) 
}

type PostgresDB struct {
	db *sql.DB
}

func NewDB(db *sql.DB) DB {
	return PostgresDB{db: db}
}

func (d PostgresDB) GetTechnologies() ([]*model.Technology, error) {
	rows, err := d.db.Query("SELECT name, details FROM public.technologies")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var technologies []*model.Technology
	for rows.Next() {
		t := new(model.Technology)
		if err := rows.Scan(&t.Name, &t.Details); err != nil {
			return nil, err
		}
		technologies = append(technologies, t)
	}
	return technologies, nil
}

func (d PostgresDB) CreateBlog(blog *model.Blog) error {
    query := `INSERT INTO blogs (title, body, coverURL) VALUES ($1, $2, $3) RETURNING id`
    return d.db.QueryRow(query, blog.Title, blog.Body, blog.CoverURL).Scan(&blog.ID)
}

func (d PostgresDB) GetBlogs() ([]*model.Blog, error) {
    log.Println("Fetching blogs...")
    rows, err := d.db.Query("SELECT id, title, coverURL, body FROM public.blogs")
    if err != nil {
        log.Printf("Error fetching blogs: %v", err)
        return nil, err
    }
    defer rows.Close()

    var blogs []*model.Blog
    for rows.Next() {
        b := new(model.Blog)
        if err := rows.Scan(&b.ID, &b.Title, &b.CoverURL, &b.Body); err != nil {
            log.Printf("Error scanning blog row: %v", err)
            return nil, err
        }
        blogs = append(blogs, b)
    }
    log.Printf("Found %d blogs", len(blogs))
    return blogs, nil
}

func (d PostgresDB) GetBlog(id int) (*model.Blog, error) {
    println(id)
    t := new(model.Blog)
    query := `SELECT id, title, body, coverURL FROM blogs WHERE id = $1`
    err := d.db.QueryRow(query, id).Scan(&t.ID, &t.Title, &t.Body, &t.CoverURL)
    if err != nil {
        return nil, err
    }
    return t, nil
}

func (d PostgresDB) UpdateBlog(id int, blog *model.Blog) error {
    query := `UPDATE blogs SET title = $1, body = $2, coverURL = $3 WHERE id = $4`
    _, err := d.db.Exec(query, blog.Title, blog.Body, blog.CoverURL, id)
    return err
}

func (d PostgresDB) DeleteBlog(id int) error {
    query := `DELETE FROM blogs WHERE id = $1`
    _, err := d.db.Exec(query, id)
    return err
}