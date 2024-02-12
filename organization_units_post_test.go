package vismanet_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestOrganizationUnitsPost(t *testing.T) {
	req := client.NewOrganizationUnitsPost()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}

