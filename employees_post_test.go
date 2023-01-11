package vismanet_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestEmployeesPost(t *testing.T) {
	req := client.NewEmployeesPost()
	resp, err := req.DoCSV()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
