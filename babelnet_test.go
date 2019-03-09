package babelnet

import (
	"fmt"
	"testing"
)

func TestClient_GetVersion(t *testing.T) {
	client := NewClient("https://babelnet.io", "PutYourKeyHere")
	version := client.GetVersion()
	fmt.Println(version)
	ids := client.GetSynSetIds("apple", "EN")
	fmt.Println(ids[0].ID)
	synSet := client.GetSynSet("bn:14792761n")
	fmt.Println(synSet.Categories)
}
