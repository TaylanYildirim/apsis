package employeeHandler

import (
	"apsis/util/errorUtil"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreate(t *testing.T) {
	reqBody := map[string]int{
		"score": 5,
	}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/teams/%s/employee", "0"), bytes.NewReader(body))
	var m map[string]int
	err := json.NewDecoder(req.Body).Decode(&m)
	errorUtil.Check(err)
	req.Body.Close()
	w := httptest.NewRecorder()
	res := w.Result()
	defer res.Body.Close()

	if res.Status != "200 OK" {
		t.Errorf("err: response status %v", res.Status)
	}
}

func TestDelete(t *testing.T) {

}
