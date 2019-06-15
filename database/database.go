package database

import (
	"sync"
)

type Database struct {
	DBH  DBHandler
	Chan chan []string
}

type DBHandler interface {
	setup() error
	insert(chan []string)
}

func New(name string) *Database {
	var register = map[string]DBHandler{
		name: new(DBObject),
	}

	return &Database{
		DBH: register[name],
	}
}

func (d *Database) Run() error {
	var (
		wg  sync.WaitGroup
		err error
	)

	err = d.DBH.setup()
	if err != nil {
		return err
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		d.DBH.insert(d.Chan)
	}()
	wg.Wait()

	return nil
}

func (d *Database) End() {
	close(d.Chan)
}
