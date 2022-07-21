// The memory package is the one in charge of managing the writing and reading
// in the embedded memory of the program, in this case bbolt library.
package memory

import (
	"fmt"
	"todo/internal/check"
	"todo/internal/printer"

	"go.etcd.io/bbolt"
)

const (
	// path and file name of the embedded database file
	memoryPath = ".database"
	// name of the bucket in the database
	bucketName = "TODOS"
)

var (
	// db is the database connection
	db *bbolt.DB
)

// Open, opens the communication with the database.
func Open() {
	database, err := bbolt.Open(memoryPath, 0666, nil)
	check.Err(err)
	db = database
	createBucket()
}

// Close communication
func Close() {
	db.Close()
}

// Create Bucket, Action to create bucket, a process necessary for the bbolt
// database.
func createBucket() {
	err := db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		return err
	})

	check.CompareErr(err)
}

// Set saves an element, this function receives two arguments key and value,
// both must be string type
// example: memory.Set("key", "value")
func Set(key, value string) error {
	err := db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		err := bucket.Put([]byte(key), []byte(value))
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

// Get, gets an element with the element key.
// example: memory.Get("key")
func Get(key string) (string, error) {
	result := ""

	err := db.View(func(tx *bbolt.Tx) error {
		value := tx.Bucket([]byte(bucketName)).Get([]byte(key))
		result = string(value)
		return nil
	})

	return result, err
}

// GetAll fetches all items from memory and displays them in a list
// example: memory.GetList()
func GetAll() ([]string, error) {
	todoList := []string{}

	err := db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		cursor := bucket.Cursor()

		for k, v := cursor.Last(); k != nil; k, v = cursor.Prev() {
			todoList = append(todoList, string(v))
		}

		return nil
	})

	return todoList, err
}

// Remove, removes an item from memory using its key.
// example: memory.Remove("key")
func Remove(key string) {
	err := db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		check.CompareErr(err)

		cursor := bucket.Cursor()
		exist := false

		for k, _ := cursor.Last(); k != nil; k, _ = cursor.Prev() {
			if string(k) == key {
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

// DeleteAllTodos, Eleminates all the elements stored in memory
// example: memory.DeleteAllTodos()
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
