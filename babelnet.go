package babelnet

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"net/http"
)

const (
	Key        = "key"
	Lemma      = "lemma"
	SearchLang = "searchLang"
	SynSetId   = "id"
)

// client to connect to the babel REST api server.
type Client struct {
	baseUrl    string
	httpClient *http.Client
	key        string
}

func NewClient(baseUrl string, key string) (client *Client) {
	client = new(Client)
	client.key = key
	client.baseUrl = baseUrl
	client.httpClient = http.DefaultClient
	return
}

func (client *Client) GetSynSet(synSetId string) (resp SynSetInfoResponse) {
	req := &request{
		method:   "GET",
		endpoint: "/v5/getSynset",
	}
	req.setParam(SynSetId, synSetId)
	client.constructRequest(req)
	request, err := http.NewRequest(req.method, req.fullUrl, req.body)
	checkError(err)
	fmt.Println(request)
	response, err := client.httpClient.Do(request)
	checkError(err)
	data := client.parseResponse(response)
	mapstructure.Decode(data, &resp)
	return
}

/**
 * Request template : https://babelnet.io/v5/getSynsetIds?lemma={lemma}&searchLang={searchLang}&key={key}
 */
func (client *Client) GetSynSetIds(lemma string, lang string) (resp SynSetIdResponse) {
	req := &request{
		method:   "GET",
		endpoint: "/v5/getSynsetIds",
	}

	req.setParam(Lemma, lemma)
	req.setParam(SearchLang, lang)
	client.constructRequest(req)
	request, err := http.NewRequest(req.method, req.fullUrl, req.body)
	checkError(err)
	fmt.Println(request)
	response, err := client.httpClient.Do(request)
	checkError(err)
	data := client.parseResponse(response)
	mapstructure.Decode(data, &resp)
	fmt.Println("response is")
	return
}

// Get babelNet api version.
func (client *Client) GetVersion() (version string) {
	r := &request{
		method:   "GET",
		endpoint: "/v5/getVersion",
	}
	client.constructRequest(r)
	request, err := http.NewRequest(r.method, r.fullUrl, r.body)
	checkError(err)
	fmt.Println(request)
	response, err := client.httpClient.Do(request)
	checkError(err)
	var resp VersionResponse
	data := client.parseResponse(response)
	mapstructure.Decode(data, &resp)
	version = resp.Version
	checkError(err)
	fmt.Println(version)
	return
}

func checkError(e error) {
	if e != nil {
		panic(e)
	}
}

func (client *Client) constructRequest(req *request) (err error) {

	fullURL := fmt.Sprintf("%s%s", client.baseUrl, req.endpoint)
	req.setParam(Key, client.key)
	queryString := req.query.Encode()
	fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	req.fullUrl = fullURL
	return nil
}

/**
  Parse Json response to struct>
 */
func (client *Client) parseResponse(response *http.Response) (responseValue interface{}) {
	if response.StatusCode != http.StatusOK {
		errorMessage := fmt.Sprintf("Error while calling API : %d", response.StatusCode)
		fmt.Println(errorMessage)
		panic(errorMessage)
	}
	fmt.Println(response)
	data, err := ioutil.ReadAll(response.Body)
	checkError(err)
	defer response.Body.Close()
	err = json.Unmarshal(data, &responseValue)
	checkError(err)
	return responseValue
}
