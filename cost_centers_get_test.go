package vismanet_test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestCostCentersGet(t *testing.T) {
	req := client.NewCostCentersGet()
	req.PathParams().JobID = "7882910c-5a7a-4d46-8dde-c6343fee9782"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	log.Println(resp.CostCentersFileUris[0])
	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
