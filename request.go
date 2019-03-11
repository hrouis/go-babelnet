package babelnet

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type params map[string]interface{}

type request struct {
	method   string
	endpoint string
	query    url.Values
	header   http.Header
	body     io.Reader
	fullUrl  string
}

/************************************/
/* Babel Sense Request and setters. */
/************************************/
type BabelSenseRequest struct {
	Lemma      string
	SearchLang string
}

func (babelReq *BabelSenseRequest) SetLemma(lemma string) *BabelSenseRequest {
	babelReq.Lemma = lemma
	return babelReq
}

func (babelReq *BabelSenseRequest) SetSearchLang(lang string) *BabelSenseRequest {
	babelReq.SearchLang = lang
	return babelReq
}

/*********************************/
/* Babel Id Request and setters. */
/*********************************/
type BabelIdRequest struct {
	Id         string
	SearchLang string // optional
	TargetLang string // optional
	Pos        string // optional
	Source     string
	WnVersion  string // optional
}

func (babelReq *BabelIdRequest) SetId(id string) *BabelIdRequest {
	babelReq.Id = id
	return babelReq
}

func (babelReq *BabelIdRequest) SetSearchLang(lang string) *BabelIdRequest {
	babelReq.SearchLang = lang
	return babelReq
}

func (babelReq *BabelIdRequest) SetTargetLang(lang string) *BabelIdRequest {
	babelReq.TargetLang = lang
	return babelReq
}

func (babelReq *BabelIdRequest) SetPos(pos string) *BabelIdRequest {
	babelReq.Pos = pos
	return babelReq
}

func (babelReq *BabelIdRequest) SetSource(source string) *BabelIdRequest {
	babelReq.Source = source
	return babelReq
}

func (babelReq *BabelIdRequest) SetWnVersion(wnVersion string) *BabelIdRequest {
	babelReq.WnVersion = wnVersion
	return babelReq
}

/**************************/
/* Request params setter. */
/**************************/

// setParam set param with key/value to query string
func (r *request) setParam(key string, value interface{}) *request {
	if r.query == nil {
		r.query = url.Values{}
	}
	r.query.Set(key, fmt.Sprintf("%v", value))
	return r
}

// setParams set params with key/values to query string
func (r *request) setParams(m params) *request {
	for k, v := range m {
		r.setParam(k, v)
	}
	return r
}

/*****************************/
/* API Response structures.  */
/*****************************/

type VersionResponse struct {
	Version string `json:"version"`
}

type SynSetIdResponse struct {
	ID     string `json:"id"`
	Pos    string `json:"pos"`
	Source string `json:"source"`
}

type SynSetInfoResponse struct {
	Senses []struct {
		Type       string `json:"type"`
		Properties struct {
			FullLemma   string `json:"fullLemma"`
			SimpleLemma string `json:"simpleLemma"`
			Source      string `json:"source"`
			SenseKey    string `json:"senseKey"`
			Frequency   int    `json:"frequency"`
			Language    string `json:"language"`
			Pos         string `json:"pos"`
			SynsetID    struct {
				ID     string `json:"id"`
				Pos    string `json:"pos"`
				Source string `json:"source"`
			} `json:"synsetID"`
			TranslationInfo string `json:"translationInfo"`
			Pronunciations  struct {
				Audios         []interface{} `json:"audios"`
				Transcriptions []interface{} `json:"transcriptions"`
			} `json:"pronunciations"`
			BKeySense bool `json:"bKeySense"`
			IDSense   int  `json:"idSense"`
		} `json:"properties"`
	} `json:"senses"`
	WnOffsets []interface{} `json:"wnOffsets"`
	MainSense string        `json:"mainSense"`
	Glosses   []struct {
		Source      string `json:"source"`
		SourceSense string `json:"sourceSense"`
		Language    string `json:"language"`
		Gloss       string `json:"gloss"`
		Tokens      []struct {
			Start int `json:"start"`
			End   int `json:"end"`
			ID    struct {
				ID     string `json:"id"`
				Pos    string `json:"pos"`
				Source string `json:"source"`
			} `json:"id"`
			Word string `json:"word"`
		} `json:"tokens"`
	} `json:"glosses"`
	Examples []interface{} `json:"examples"`
	Images   []struct {
		Name      string   `json:"name"`
		Languages []string `json:"languages"`
		URLSource string   `json:"urlSource"`
		License   string   `json:"license"`
		ThumbURL  string   `json:"thumbUrl"`
		URL       string   `json:"url"`
		BadImage  bool     `json:"badImage"`
	} `json:"images"`
	SynsetType string `json:"synsetType"`
	Categories []struct {
		Category string `json:"category"`
		Language string `json:"language"`
	} `json:"categories"`
	Translations struct {
	} `json:"translations"`
	Domains struct {
		COMPUTING float64 `json:"COMPUTING"`
	} `json:"domains"`
	LnToCompound struct {
	} `json:"lnToCompound"`
	LnToOtherForm struct {
		EN []string `json:"EN"`
	} `json:"lnToOtherForm"`
	FilterLangs  []string `json:"filterLangs"`
	BkeyConcepts bool     `json:"bkeyConcepts"`
}

type BabelSenseResponse struct {
	Type       string `json:"type"`
	Properties struct {
		FullLemma   string `json:"fullLemma"`
		SimpleLemma string `json:"simpleLemma"`
		Source      string `json:"source"`
		SenseKey    string `json:"senseKey"`
		Frequency   int    `json:"frequency"`
		Language    string `json:"language"`
		Pos         string `json:"pos"`
		SynsetID    struct {
			ID     string `json:"id"`
			Pos    string `json:"pos"`
			Source string `json:"source"`
		} `json:"synsetID"`
		TranslationInfo string `json:"translationInfo"`
		Pronunciations  struct {
			Audios         []interface{} `json:"audios"`
			Transcriptions []interface{} `json:"transcriptions"`
		} `json:"pronunciations"`
		YAGOURL   string `json:"YAGOURL"`
		BKeySense bool   `json:"bKeySense"`
		IDSense   int    `json:"idSense"`
	} `json:"properties"`
}

type BabelEdgeResponse struct {
	Language string `json:"language"`
	Pointer  struct {
		FSymbol       string `json:"fSymbol"`
		Name          string `json:"name"`
		ShortName     string `json:"shortName"`
		RelationGroup string `json:"relationGroup"`
		IsAutomatic   bool   `json:"isAutomatic"`
	} `json:"pointer"`
	Target           string `json:"target"`
	Weight           int    `json:"weight"`
	NormalizedWeight int    `json:"normalizedWeight"`
}
