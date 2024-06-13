package postgres

import (
	model "User_Product/models"
	"database/sql"
)

type UserProductRepo struct {
	DB *sql.DB
}

func NewUserProductRepo(db *sql.DB) *UserProductRepo {
	return &UserProductRepo{DB: db}
}

func (u *UserProductRepo) GetUserProducts(userID int) ([]model.Product, error) {
	tr, err := u.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()

	query := `
	SELECT 	p.id, p.name, p.description, p.price,
            u.quantity
    FROM users u
    JOIN user_products u ON u.id = u.user_id
    JOIN products p ON u.product_id = p.id
	WHERE u.id = $1`
	rows, err := u.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var pr model.Product
		err := rows.Scan(&pr.ID, &pr.Name, &pr.Description, &pr.Price, &pr.Quantity)
		if err != nil {
			return nil, err
		}
		products = append(products, pr)
	}
	return products, nil
}

func (u *UserProductRepo) AddProduct(userID, productID, quantity int) error {
	tr, err := u.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()

	query := `insert into user_products(user_id, product_id, quantity)
	values($1, $2, $3)
	on conflict (user_id, product_id)
	do update set quantity = user_products.quantity + $3`
	_, err = u.DB.Exec(query, userID, productID, quantity)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserProductRepo) UpdateProductQuantityForUser(userID, productID, quantity int) error {
	tr, err := u.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()

	query := "update user_products set quantity = $3 where user_id = $1 and product_id = $2"
	_, err = u.DB.Exec(query, userID, productID, quantity)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserProductRepo) RemoveProductFromUser(userID, productID int) error {
	tr, err := u.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()

	_, err = u.DB.Exec("delete from user_products where user_id = $1 and product_id = $2",
		userID, productID)
	if err != nil {
		return err
	}
	return nil
}
