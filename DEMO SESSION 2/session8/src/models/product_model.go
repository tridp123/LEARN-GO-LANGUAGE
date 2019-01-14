package models

import (
	"database/sql"
	"entities"
)

type ProductModel struct {
	Db *sql.DB
}

func (ProductModel ProductModel) FindAll() ([]entities.Product, error) {
	rows, err := ProductModel.Db.Query("select * from product")
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int

			rows.Scan(&id, &name, &price, &quantity)
			products = append(products, entities.Product{
				Id:       id,
				Name:     name,
				Price:    price,
				Quantity: quantity,
			})
		}
		return products, nil
	}
}

func (ProductModel ProductModel) Find(id int) (entities.Product, error) {
	rows, err := ProductModel.Db.Query("select * from product where id=?", id)
	//rows, err := ProductModel.Db.Query("select * from product where id=? and price=?", id, price) => giong preparestatement
	if err != nil {
		//tra ve doi tuong null
		return entities.Product{}, err
	} else {
		var products entities.Product
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int

			rows.Scan(&id, &name, &price, &quantity)
			products.Id = id
			products.Name = name
			products.Price = price
			products.Quantity = quantity
		}
		return products, nil
	}
}

func (productModel ProductModel) Create(product *entities.Product) error {
	result, err := productModel.Db.Exec("insert into product(name, price, quantity) values(?,?,?)", product.Name, product.Price, product.Quantity)
	if err != nil {
		return err
	} else {
		product.Id, _ = result.LastInsertId()
		return nil
	}
}

func (productModel ProductModel) Delete(id int64) (int64, error) {
	result, err := productModel.Db.Exec("delete from product where id = ?", id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}

func (productModel ProductModel) Update(product entities.Product) (int64, error) {
	result, err := productModel.Db.Exec("update product set name = ?, price=?, quantity=?", product.Name, product.Price, product.Quantity)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
