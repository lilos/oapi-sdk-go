// Code generated by lark suite oapi sdk gen
package v1

import (
	"github.com/larksuite/oapi-sdk-go/api/core/tools"
)

type Term struct {
	From            string   `json:"from,omitempty"`
	To              string   `json:"to,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *Term) MarshalJSON() ([]byte, error) {
	type cp Term
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type TextTranslateReqBody struct {
	SourceLanguage  string   `json:"source_language,omitempty"`
	Text            string   `json:"text,omitempty"`
	TargetLanguage  string   `json:"target_language,omitempty"`
	Glossary        []*Term  `json:"glossary,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *TextTranslateReqBody) MarshalJSON() ([]byte, error) {
	type cp TextTranslateReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type TextTranslateResult struct {
	Text string `json:"text,omitempty"`
}

type TextDetectReqBody struct {
	Text            string   `json:"text,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *TextDetectReqBody) MarshalJSON() ([]byte, error) {
	type cp TextDetectReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type TextDetectResult struct {
	Language string `json:"language,omitempty"`
}
