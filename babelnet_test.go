package babelnet

import (
	"fmt"
	"testing"
)

func TestClient_GetVersion(t *testing.T) {
	client := NewClient("https://babelnet.io", "PutYourAPIKeyHere")
	version := client.GetVersion()
	fmt.Println(version)
	ids := client.GetSynSetIds("apple", "EN")
	fmt.Println(ids[0].ID)
	synSet := client.GetSynSet("bn:14792761n")
	fmt.Println(synSet.Categories)
	req := new(BabelSenseRequest)
	req.SetLemma("apple").SetSearchLang("EN")
	result := client.GetSenses(req)
	fmt.Println(result[0].Type)

	req2 := new(BabelIdRequest)
	req2.SetId("trousers").SetSearchLang("EN").SetPos("NOUN").SetSource("WIKI")
	response := client.GetBabelNetId(req2)
	fmt.Println(response)

	response2 := client.GetOutgoingEdges("bn:14792761n")
	fmt.Println(response2)
}
