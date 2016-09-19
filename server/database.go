package main

import (
    "path"
    "os"
    "log"

    leveldb "github.com/syndtr/goleveldb/leveldb"
)

type DbCallback func(string, string)

type Database struct {
    connector   *leveldb.DB
    isOpen      bool
}


//
// Opens a database folder
// If the folder doesn't exist, creates it
//
func (this *Database) Open( name string ) {
    var cwd, _		= os.Getwd();
    var db, dbErr   = leveldb.OpenFile(path.Join(cwd, "db", name), nil)

    log.Println("[Database] Opening: " + path.Join(cwd, "db", name))

    this.isOpen = (dbErr == nil)
    if (this.isOpen) {
        this.connector = db
    }
}


//
// GET
//
func (this *Database) Get(key string) string {
    if !this.isOpen {
        log.Println("Error: database is closed")
        return ""
    }

    var data, err = this.connector.Get([]byte(key), nil)
    var value = string(data); 
    if (err == nil && len(value) > 0) {
        return value
    }

    return ""
}


//
// SET
//
func (this *Database) Set(key string, value string) bool {
    if !this.isOpen {
        return false;
    }
    err := this.connector.Put([]byte(key), []byte(value), nil)
    return (err == nil)
}


//
// UNSET
//
func (this *Database) Unset(key string) bool {
    if !this.isOpen {
        log.Println("Error: database is closed")
        return false;
    }

    err :=  this.connector.Delete([]byte(key), nil)
    return err == nil
}


//
// FOREACH
//
func (this *Database) ForEach(fn DbCallback) {
    iter := this.connector.NewIterator(nil, nil)
    for iter.Next() {
        fn( string(iter.Key()), string(iter.Value()) )
    }
    iter.Release()
}


//
// CLOSE
//
func (this *Database) Close() {
    if this.isOpen {
        log.Println("[Database] Closing")
        this.connector.Close()
        this.isOpen = false;
    }
}
