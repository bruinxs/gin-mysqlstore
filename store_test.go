package ginmysqlstore

import (
	"testing"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/tester"
)

var newStore = func(t *testing.T) sessions.Store {
	store, err := NewMySQLStore("test:test@tcp(localhost:3306)/test?parseTime=true&loc=Local", "test_store", "/", 60, []byte("fake_key"))
	if err != nil {
		t.Fatal(err)
	}
	return store
}

func TestMySQLStore_SessionGetSet(t *testing.T) {
	tester.GetSet(t, newStore)
}

func TestMySQLStore_SessionDeleteKey(t *testing.T) {
	tester.DeleteKey(t, newStore)
}

func TestMySQLStore_SessionFlashes(t *testing.T) {
	tester.Flashes(t, newStore)
}

func TestMySQLStore_SessionClear(t *testing.T) {
	tester.Clear(t, newStore)
}

func TestMySQLStore_SessionOptions(t *testing.T) {
	tester.Options(t, newStore)
}
