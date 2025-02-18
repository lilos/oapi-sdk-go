// Code generated by lark suite oapi sdk gen
package v1

import (
	"github.com/larksuite/oapi-sdk-go/api/core/tools"
)

type Mailgroup struct {
	MailgroupId             string   `json:"mailgroup_id,omitempty"`
	Email                   string   `json:"email,omitempty"`
	Name                    string   `json:"name,omitempty"`
	Description             string   `json:"description,omitempty"`
	DirectMembersCount      string   `json:"direct_members_count,omitempty"`
	IncludeExternalMember   bool     `json:"include_external_member,omitempty"`
	IncludeAllCompanyMember bool     `json:"include_all_company_member,omitempty"`
	WhoCanSendMail          string   `json:"who_can_send_mail,omitempty"`
	ForceSendFields         []string `json:"-"`
}

func (s *Mailgroup) MarshalJSON() ([]byte, error) {
	type cp Mailgroup
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type MailgroupMember struct {
	MemberId        string   `json:"member_id,omitempty"`
	Email           string   `json:"email,omitempty"`
	UserId          string   `json:"user_id,omitempty"`
	DepartmentId    string   `json:"department_id,omitempty"`
	Type            string   `json:"type,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *MailgroupMember) MarshalJSON() ([]byte, error) {
	type cp MailgroupMember
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type MailgroupPermissionMember struct {
	PermissionMemberId string   `json:"permission_member_id,omitempty"`
	UserId             string   `json:"user_id,omitempty"`
	DepartmentId       string   `json:"department_id,omitempty"`
	Type               string   `json:"type,omitempty"`
	ForceSendFields    []string `json:"-"`
}

func (s *MailgroupPermissionMember) MarshalJSON() ([]byte, error) {
	type cp MailgroupPermissionMember
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type PublicMailbox struct {
	PublicMailboxId string   `json:"public_mailbox_id,omitempty"`
	Email           string   `json:"email,omitempty"`
	Name            string   `json:"name,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *PublicMailbox) MarshalJSON() ([]byte, error) {
	type cp PublicMailbox
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type PublicMailboxMember struct {
	MemberId        string   `json:"member_id,omitempty"`
	UserId          string   `json:"user_id,omitempty"`
	Type            string   `json:"type,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *PublicMailboxMember) MarshalJSON() ([]byte, error) {
	type cp PublicMailboxMember
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type MailgroupListResult struct {
	HasMore   bool         `json:"has_more,omitempty"`
	PageToken string       `json:"page_token,omitempty"`
	Items     []*Mailgroup `json:"items,omitempty"`
}

type MailgroupMemberListResult struct {
	HasMore   bool               `json:"has_more,omitempty"`
	PageToken string             `json:"page_token,omitempty"`
	Items     []*MailgroupMember `json:"items,omitempty"`
}

type MailgroupPermissionMemberListResult struct {
	HasMore   bool                         `json:"has_more,omitempty"`
	PageToken string                       `json:"page_token,omitempty"`
	Items     []*MailgroupPermissionMember `json:"items,omitempty"`
}

type PublicMailboxListResult struct {
	HasMore   bool             `json:"has_more,omitempty"`
	PageToken string           `json:"page_token,omitempty"`
	Items     []*PublicMailbox `json:"items,omitempty"`
}

type PublicMailboxMemberListResult struct {
	HasMore   bool                   `json:"has_more,omitempty"`
	PageToken string                 `json:"page_token,omitempty"`
	Items     []*PublicMailboxMember `json:"items,omitempty"`
}
