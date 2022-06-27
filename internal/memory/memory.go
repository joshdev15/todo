package memory

import (
	"encoding/json"
	"fmt"
	"todo/internal/check"
	"todo/internal/printer"
	"todo/internal/structs/todo"

	"go.etcd.io/bbolt"
)

const (
	memoryPath    = "memory/user"
	bucketName    = "TODOS"
	subBucketName = "LIST"
)

var (
	db *bbolt.DB
)

func Open() {
	database, err := bbolt.Open(memoryPath, 0666, nil)
	check.Err(err)
	db = database
	createBucket()
}

func Close() {
	db.Close()
}

func createBucket() {
	err := db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		return err
	})

	check.CompareErr(err)
}

func Set(key, value string) {
	err := db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		check.CompareErr(err)

		err = bucket.Put([]byte(key), []byte(value))
		check.Err(err)

		printer.Success("Add new todo")

		return nil
	})

	check.CompareErr(err)
}

func Get(id string) {
	err := db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		check.CompareErr(err)

		key := bucket.Get([]byte(id))

		var tmpItem todo.Todo
		err = json.Unmarshal([]byte(key), &tmpItem)
		check.Err(err)
		printer.TodoTable([]todo.Todo{tmpItem})

		return nil
	})

	check.CompareErr(err)
}

func SetDone(id string) {
	err := db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		check.CompareErr(err)

		key := bucket.Get([]byte(id))

		var tmpItem todo.Todo
		err = json.Unmarshal([]byte(key), &tmpItem)
		check.Err(err)

		tmpItem.Done = !tmpItem.Done
		value := tmpItem.ToJSON()
		err = bucket.Put([]byte(id), []byte(value))
		check.Err(err)

		isDone := "done"
		if tmpItem.Done {
			isDone = "done"
		} else {
			isDone = "not done"
		}

		isDone = fmt.Sprintf("Set as %v", isDone)
		printer.Success(isDone)

		return nil
	})

	check.CompareErr(err)
}

func Remove(id string) {
	err := db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		check.CompareErr(err)

		cursor := bucket.Cursor()
		exist := false

		for k, _ := cursor.Last(); k != nil; k, _ = cursor.Prev() {
			if string(k) == id {
				exist = true
				cursor.Delete()
			}
		}

		if !exist {
			fmt.Println("El id del todo no existe")
		} else {
			printer.Success("Remove todo")
		}

		return nil
	})

	check.CompareErr(err)
}

func GetList() {
	err := db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		check.CompareErr(err)

		cursor := bucket.Cursor()
		todoList := []todo.Todo{}

		for k, v := cursor.Last(); k != nil; k, v = cursor.Prev() {
			var tmpItem todo.Todo
			err = json.Unmarshal([]byte(v), &tmpItem)
			check.Err(err)
			todoList = append(todoList, tmpItem)
		}

		printer.TodoTable(todoList)

		return nil
	})

	check.CompareErr(err)
}

func DeleteAllTodos() {
	err := db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		check.CompareErr(err)

		cursor := bucket.Cursor()

		for k, _ := cursor.Last(); k != nil; k, _ = cursor.Prev() {
			cursor.Delete()
		}

		printer.Success("Delete all to dos")

		return nil
	})

	check.CompareErr(err)
}
