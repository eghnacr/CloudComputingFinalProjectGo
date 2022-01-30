// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    result, err := UnmarshalResult(bytes)
//    bytes, err = result.Marshal()

package main

import "encoding/json"

func UnmarshalResult(data []byte) (Result, error) {
	var r Result
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Result) Marshall() ([]byte, error) {
	return json.Marshal(r)
}

type Result struct {
	Status              string        `json:"status"`
	CreatedDateTime     string        `json:"createdDateTime"`
	LastUpdatedDateTime string        `json:"lastUpdatedDateTime"`
	AnalyzeResult       AnalyzeResult `json:"analyzeResult"`
}

type AnalyzeResult struct {
	Version     string       `json:"version"`
	ReadResults []ReadResult `json:"readResults"`
	PageResults []PageResult `json:"pageResults"`
}

type PageResult struct {
	Page   int64         `json:"page"`
	Tables []interface{} `json:"tables"`
}

type ReadResult struct {
	Page           int64           `json:"page"`
	Angle          int64           `json:"angle"`
	Width          int64           `json:"width"`
	Height         int64           `json:"height"`
	Unit           string          `json:"unit"`
	Lines          []Line          `json:"lines"`
	SelectionMarks []SelectionMark `json:"selectionMarks"`
}

type Line struct {
	BoundingBox []int64    `json:"boundingBox"`
	Text        string     `json:"text"`
	Appearance  Appearance `json:"appearance"`
	Words       []Word     `json:"words"`
}

type Appearance struct {
	Style Style `json:"style"`
}

type Style struct {
	Name       Name    `json:"name"`
	Confidence float64 `json:"confidence"`
}

type Word struct {
	BoundingBox []int64 `json:"boundingBox"`
	Text        string  `json:"text"`
	Confidence  float64 `json:"confidence"`
}

type SelectionMark struct {
	BoundingBox []int64 `json:"boundingBox"`
	Confidence  float64 `json:"confidence"`
	State       string  `json:"state"`
}

type Name string

const (
	Handwriting Name = "handwriting"
	Other       Name = "other"
)
