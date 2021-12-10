// Code generated by lark suite oapi sdk gen
package v3

import (
	"github.com/larksuite/oapi-sdk-go/api/core/tools"
)

type Bot struct {
}

type BotInfo struct {
	ActivateStatus  int      `json:"activate_status,omitempty"`
	AppName         string   `json:"app_name,omitempty"`
	AvatarUrl       string   `json:"avatar_url,omitempty"`
	IpWhiteList     []string `json:"ip_white_list,omitempty"`
	OpenId          string   `json:"open_id,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *BotInfo) MarshalJSON() ([]byte, error) {
	type cp BotInfo
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type BotGetResult struct {
	Bot *BotInfo `json:"bot,omitempty"`
}
