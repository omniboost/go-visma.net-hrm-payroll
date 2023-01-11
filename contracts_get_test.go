package vismanet_test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestContractsGet(t *testing.T) {
	req := client.NewContractsGet()
	req.PathParams().JobID = "96f53e50-9d2b-4358-ae1e-10ae64811289"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	log.Println(resp.ContractFileUris[0])
	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
