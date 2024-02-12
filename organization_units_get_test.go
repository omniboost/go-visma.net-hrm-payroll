package vismanet_test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestOrganixationUnitsGet(t *testing.T) {
	req := client.NewOrganixationUnitsGet()
	req.PathParams().JobID = "658d5090-6529-418f-af30-6a8699f56ac8"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	log.Println(resp.OrganizationUnitsFileUris[0])
	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}

