package repository

import (
	"encoding/json"
	"fmt"
	"github.com/bersen66/wb_task0/pkg/entities"
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
	qGetOrder    = "SELECT * FROM orders WHERE order_uid=$1"
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

func (st *PGStorage) InsertOrder(order *entities.Order) error {
	tx := st.db.MustBegin()
	query := fmt.Sprintf(qInsertOrder, order.DBString())
	fmt.Println(query)
	_, err := tx.Exec(query)
	fmt.Println()
	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}

func (st *PGStorage) Contains(order *entities.Order) (bool, error) {
	return false, nil
}

type tempOrder struct {
	Uid             string  `json:"order_uid" db:"order_uid"`
	TrackNumber     string  `json:"track_number" db:"track_number"`
	Entry           string  `json:"entry" db:"entry"`
	Delivery_       []uint8 `json:"delivery" db:"delivery"`
	Payment_        []uint8 `json:"payment" db:"payment"`
	Items           []uint8 `json:"items" db:"items"`
	Locale          string  `json:"locale" db:"locale"`
	InternalSign    string  `json:"internal_signature" db:"internal_signature"`
	CustomerId      string  `json:"customer_id" db:"customer_id"`
	DeliveryService string  `json:"delivery_service" db:"delivery_service"`
	ShardKey        string  `json:"shardkey" db:"shardkey"`
	SmId            uint32  `json:"sm_id" db:"sm_id"`
	DateCreated     string  `json:"date_created" db:"date_created"`
	OofShard        string  `json:"oof_shard" db:"oof_shard"`
}

func toOriginOrder(temp *tempOrder) (*entities.Order, error) {

	result := entities.Order{
		Uid:             temp.Uid,
		TrackNumber:     temp.TrackNumber,
		Entry:           temp.Entry,
		Locale:          temp.Locale,
		InternalSign:    temp.InternalSign,
		CustomerId:      temp.CustomerId,
		DeliveryService: temp.DeliveryService,
		ShardKey:        temp.ShardKey,
		SmId:            temp.SmId,
		DateCreated:     temp.DateCreated,
		OofShard:        temp.OofShard,
	}
	err := json.Unmarshal(temp.Delivery_, &result.Delivery_)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(temp.Payment_, &result.Payment_)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(temp.Items, &result.Items)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (st *PGStorage) GetOrder(uuid string) (*entities.Order, error) {
	var temp tempOrder

	err := st.db.Get(&temp, qGetOrder, uuid)

	if err != nil {
		return nil, err
	}

	return toOriginOrder(&temp)
}
