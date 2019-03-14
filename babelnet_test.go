package babelnet

import (
	"gotest.tools/assert"
	"testing"
)

var client = NewClient("https://babelnet.io", "Your API Key Here")

func TestClient_GetVersion(t *testing.T) {
	version := client.GetVersion()
	assert.Equal(t, version, "V4_0")
}

func TestClient_GetSynSetIds(t *testing.T) {
	ids := client.GetSynSetIds("apple", "EN")
	assert.Equal(t, len(ids), 17)
	assert.Equal(t, ids[0].ID, "bn:00289737n")
}

func TestClient_GetSynSet(t *testing.T) {
	synSet := client.GetSynSet("bn:14792761n")
	assert.Equal(t, len(synSet.Categories), 4)
}

func TestClient_GetSenses(t *testing.T) {
	req := new(BabelSenseRequest)
	req.SetLemma("apple").SetSearchLang("EN")
	result := client.GetSenses(req)
	assert.Equal(t, len(result), 261)
	assert.Equal(t, result[0].Type, "WordNetSense")
}

func TestClient_GetBabelNetIds(t *testing.T) {
	req := new(BabelIdRequest)
	req.SetId("trousers").SetSearchLang("EN").SetPos("NOUN").SetSource("WIKI")
	response := client.GetBabelNetId(req)
	assert.Equal(t, len(response), 2)
	assert.Equal(t, response[0].ID, "bn:00060423n")
}

func TestClient_GetOutgoingEdges(t *testing.T) {
	response := client.GetOutgoingEdges("bn:14792761n")
	assert.Equal(t, len(response), 561)
	assert.Equal(t,response[0].Language, "MUL")
}
