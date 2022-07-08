// The memory package is the one in charge of managing the writing and reading
// in the embedded memory of the program, in this case bbolt library.
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

// Get, gets an element with the element key.
// example: memory.Get("key")
func Get(key string) {
	err := db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		check.CompareErr(err)

		todoJSON := bucket.Get([]byte(key))

		var tmpItem todo.Todo
		err = json.Unmarshal([]byte(todoJSON), &tmpItem)
		check.Err(err)
		printer.TodoTable([]todo.Todo{tmpItem})

		return nil
	})

	check.CompareErr(err)
}

// SetDone marks a task as ready or not ready.
// example: memory.SetDone("key")
func SetDone(key string) {
	err := db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		check.CompareErr(err)

		todoJSON := bucket.Get([]byte(key))

		var tmpItem todo.Todo
		err = json.Unmarshal([]byte(todoJSON), &tmpItem)
		check.Err(err)

		tmpItem.Done = !tmpItem.Done
		value := tmpItem.ToJSON()
		err = bucket.Put([]byte(key), []byte(value))
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

// GetList fetches all items from memory and displays them in a list
// example: memory.GetList()
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
