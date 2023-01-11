package vismanet_test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestEmployeesGet(t *testing.T) {
	req := client.NewEmployeesGet()
	req.PathParams().JobID = "879689bd-b219-4f05-8660-2d3f06d9808b"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	log.Println(resp.EmployeeFileUris[0])
	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
