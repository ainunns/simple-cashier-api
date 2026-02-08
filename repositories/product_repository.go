package repositories

import (
	"database/sql"
	"errors"

	"simple-cashier-api/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) GetAll(nameFilter string) ([]models.ProductDetail, error) {
	query := `SELECT p.id, p.name, p.price, p.stock,
	                 p.category_id,
	                 c.id, c.name, c.description
	          FROM products p
	          LEFT JOIN categories c ON c.id = p.category_id`

	args := []any{}
	if nameFilter != "" {
		query += " WHERE p.name ILIKE $1"
		args = append(args, "%"+nameFilter+"%")
	}

	rows, err := repo.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]models.ProductDetail, 0)
	for rows.Next() {
		var p models.ProductDetail
		var categoryID sql.NullInt64
		var catID sql.NullInt64
		var catName sql.NullString
		var catDesc sql.NullString

		err := rows.Scan(
			&p.ID, &p.Name, &p.Price, &p.Stock,
			&categoryID,
			&catID, &catName, &catDesc,
		)
		if err != nil {
			return nil, err
		}

		if categoryID.Valid {
			val := int(categoryID.Int64)
			p.CategoryID = &val
		}

		if catID.Valid && catName.Valid {
			p.Category = &models.Category{
				ID:          int(catID.Int64),
				Name:        catName.String,
				Description: catDesc.String,
			}
		}

		products = append(products, p)
	}

	return products, nil
}

func (repo *ProductRepository) Create(product *models.Product) error {
	query := "INSERT INTO products (name, price, stock, category_id) VALUES ($1, $2, $3, $4) RETURNING id"
	err := repo.db.QueryRow(query, product.Name, product.Price, product.Stock, product.CategoryID).Scan(&product.ID)
	return err
}

func (repo *ProductRepository) GetByID(id int) (*models.ProductDetail, error) {
	query := `SELECT p.id, p.name, p.price, p.stock,
									 p.category_id,
									 c.id AS category_id,
									 c.name AS category_name,
									 c.description AS category_description
    FROM products p
    LEFT JOIN categories c ON c.id = p.category_id
    WHERE p.id = $1`

	var p models.ProductDetail
	var categoryID sql.NullInt64
	var catID sql.NullInt64
	var catName sql.NullString
	var catDesc sql.NullString

	err := repo.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &categoryID, &catID, &catName, &catDesc)

	if err == sql.ErrNoRows {
		return nil, errors.New("product not found")
	}
	if err != nil {
		return nil, err
	}

	if categoryID.Valid {
		val := int(categoryID.Int64)
		p.CategoryID = &val
	}

	if catID.Valid && catName.Valid {
		p.Category = &models.Category{
			ID:          int(catID.Int64),
			Name:        catName.String,
			Description: catDesc.String,
		}
	}

	return &p, nil
}

func (repo *ProductRepository) Update(product *models.Product) error {
	query := "UPDATE products SET name = $1, price = $2, stock = $3, category_id = $4 WHERE id = $5"
	result, err := repo.db.Exec(query, product.Name, product.Price, product.Stock, product.CategoryID, product.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("product not found")
	}

	return nil
}

func (repo *ProductRepository) Delete(id int) error {
	query := "DELETE FROM products WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("product not found")
	}

	return err
}
