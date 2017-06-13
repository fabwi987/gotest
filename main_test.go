package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/fabwi987/gotest/models"
)

type mockDB struct{}

func (mdb *mockDB) GetMeets() ([]*models.Meet, error) {
	met := make([]*models.Meet, 0)
	met = append(met, &models.Meet{1, "Stockholm", time.Now(), "Meeting no 1", time.Now(), time.Now(), "/meet/single/1", 11})
	met = append(met, &models.Meet{2, "Karlstad", time.Now(), "Meeting no 2", time.Now(), time.Now(), "/meet/single/2", 12})
	return met, nil
}

func TestMeets(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/meets", nil)

	env := Env{db: &mockDB{}}

	http.HandlerFunc(env.meetsIndex).ServeHTTP(rec, req)

	expected := "978-1503261969, Emma, Jayne Austen, £9.44\n978-1505255607, The Time Machine, H. G. Wells, £5.99\n"
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}
