package repository

import (
	"fmt"
	"github.com/bersen66/wb_task0/pkg/entities"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type PGStorageConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type PGStorage struct {
	db *sqlx.DB
}

const (
	connStr      = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s"
	qInsertOrder = "INSERT INTO orders VALUES (%s)"
)

func openDB(c PGStorageConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres",
		fmt.Sprintf(connStr, c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode))
	return db, err
}

func NewPGStorage(config PGStorageConfig) *PGStorage {
	result := new(PGStorage)
	db, err := openDB(config)
	if err != nil {
		log.Fatalf("Error while creating DB: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Can't ping DB: %v", err)
	}
	result.db = db
	return result
}

func (st *PGStorage) Insert(order *entities.Order) (bool, error) {
	tx := st.db.MustBegin()
	query := fmt.Sprintf(qInsertOrder, order.DBString())
	fmt.Println(query)
	_, err := tx.Exec(query)
	fmt.Println()
	if err != nil {
		return false, err
	}

	err = tx.Commit()
	if err != nil {
		return false, err
	}
	return true, nil
}

func (st *PGStorage) Contains(order *entities.Order) (bool, error) {
	return false, nil
}

func (st *PGStorage) GetOrder(uuid uuid.UUID) (entities.Order, error) {
	return entities.Order{}, nil
}
