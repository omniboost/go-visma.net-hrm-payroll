package vismanet_test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestCostCentersGet(t *testing.T) {
	req := client.NewCostCentersGet()
	req.PathParams().JobID = "b2d4a3f2-52d7-4377-843d-f6a91a31138b"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	log.Println(resp.CostCentersFileUris[0])
	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
