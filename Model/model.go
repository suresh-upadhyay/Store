package Model

import (
	"database/sql"
	"strconv"
)

func (p *Product) GetProduct(db *sql.DB) error {
	return db.QueryRow("SELECT name, price FROM Product WHERE id=$1",
		p.ID).Scan(&p.Name, &p.Price)
}

func (p *Product) UpdateProduct(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE Product SET name=$1, price=$2 WHERE id=$3",
			p.Name, p.Price, p.ID)

	return err
}

func (p *Product) DeleteProduct(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM Product WHERE id=$1", p.ID)

	return err
}

func (p *Product) CreateProduct(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO Product(name, price) VALUES($1, $2) RETURNING id",
		p.Name, p.Price).Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}

func GetProducts(db *sql.DB, start, count int) ([]Product, error) {
	rows, err := db.Query(
		"SELECT id, name,  price FROM Product LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []Product{}

	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}


/////////////////........ DIY........///////////////////

func (s *Store) GetStoreProduct(db *sql.DB) error {

	return db.QueryRow("SELECT p_id FROM Store WHERE s_id=$1 and is_available = $2",
		s.StoreId , true).Scan(&s.ProductId)
}


func (s *Store) SetStoreProduct(db *sql.DB , Productlist []int) error {
	DbQuery := `INSERT INTO Store (s_id, p_id, is_available) VALUES `

	for  v := range Productlist {
		DbQuery += `(` + strconv.Itoa(s.StoreId) + `,` + strconv.Itoa(Productlist[v])
		DbQuery += "," + "true"

		DbQuery+= `),`
	}
	DbQuery = DbQuery[:len(DbQuery)-1]
	DbQuery += `;`
	_, err := db.Exec(DbQuery)
	return err
}
