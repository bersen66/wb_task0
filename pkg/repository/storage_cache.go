package repository

import (
	"bufio"
	"fmt"
	cont "github.com/bersen66/wb_task0/internal/containers"
	"github.com/bersen66/wb_task0/pkg/entities"
	"os"
	"sync"
)

/*
* Decorator which should be used for caching values while dealing with database
 */
type StorageLRUCache struct {
	storage  OrdersStorage
	resolver map[string]*cont.Node
	values   *cont.LinkedList
	mutex    sync.Mutex
	capacity int
}

func NewStorageCache(capacity int, storage OrdersStorage) *StorageLRUCache {
	result := new(StorageLRUCache)

	result.storage = storage
	result.resolver = make(map[string]*cont.Node)
	result.values = cont.NewLinkedList()
	result.capacity = capacity

	return result
}

func (s *StorageLRUCache) InsertOrder(order *entities.Order) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if _, alreadyCached := s.resolver[order.Uid]; alreadyCached {
		return nil
	}
	err := s.storage.InsertOrder(order)
	if err != nil {
		return err
	}

	if s.values.Size == s.capacity {
		s.values.Erase(s.values.Back())
	}
	s.resolver[order.Uid] = s.values.PushFront(order)

	return nil
}

func (s *StorageLRUCache) GetOrder(uuid string) (*entities.Order, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if node, alreadyCached := s.resolver[uuid]; alreadyCached {
		s.values.MoveToFront(node)
		return node.Value, nil
	}

	order, err := s.storage.GetOrder(uuid)

	if err != nil {
		return nil, err
	}

	if s.values.Size == s.capacity {
		s.values.Erase(s.values.Back())
	}

	s.resolver[uuid] = s.values.PushFront(order)

	return order, nil
}

func (s *StorageLRUCache) Dump(filepath string) error {
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println("File creation error", err)
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	s.mutex.Lock()
	for uuid, _ := range s.resolver {
		_, err = writer.WriteString(uuid + "\n")
		if err != nil {
			s.mutex.Unlock()
			return err
		}
	}
	s.mutex.Unlock()
	err = writer.Flush()
	return err
}

func (s *StorageLRUCache) Load(filepath string) error {
	file, err := os.Open(filepath)
	if os.IsNotExist(err) {
		return nil
	}

	if err != nil {
		return err
	}
	defer file.Close()

	s.mutex.Lock()
	defer s.mutex.Unlock()
	reader := bufio.NewReader(file)

	for {
		uuid, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		order, err := s.storage.GetOrder(string(uuid))
		if err != nil {
			return err
		}

		if s.values.Size == s.capacity {
			s.values.Erase(s.values.Back())
		}

		s.resolver[string(uuid)] = s.values.PushFront(order)

	}

	return nil
}
