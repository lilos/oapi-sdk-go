// Code generated by lark suite oapi sdk gen
package v3

import (
	"github.com/larksuite/oapi-sdk-go/api/core/tools"
	"github.com/larksuite/oapi-sdk-go/event/core/model"
)

type AvatarInfo struct {
	Avatar72        string   `json:"avatar_72,omitempty"`
	Avatar240       string   `json:"avatar_240,omitempty"`
	Avatar640       string   `json:"avatar_640,omitempty"`
	AvatarOrigin    string   `json:"avatar_origin,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *AvatarInfo) MarshalJSON() ([]byte, error) {
	type cp AvatarInfo
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type CustomAttrEvent struct {
	ContactFieldKey string   `json:"contact_field_key,omitempty"`
	FieldValueTypes string   `json:"field_value_types,omitempty"`
	I18nFieldNames  string   `json:"i18n_field_names,omitempty"`
	AllowOpenQuery  string   `json:"allow_open_query,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *CustomAttrEvent) MarshalJSON() ([]byte, error) {
	type cp CustomAttrEvent
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type DepartmentI18nName struct {
	ZhCn            string   `json:"zh_cn,omitempty"`
	JaJp            string   `json:"ja_jp,omitempty"`
	EnUs            string   `json:"en_us,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *DepartmentI18nName) MarshalJSON() ([]byte, error) {
	type cp DepartmentI18nName
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type DepartmentStatus struct {
	IsDeleted       bool     `json:"is_deleted,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *DepartmentStatus) MarshalJSON() ([]byte, error) {
	type cp DepartmentStatus
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type DepartmentUnit struct {
	UnitId          string   `json:"unit_id,omitempty"`
	UnitType        string   `json:"unit_type,omitempty"`
	UnitName        string   `json:"unit_name,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *DepartmentUnit) MarshalJSON() ([]byte, error) {
	type cp DepartmentUnit
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type GroupEvent struct {
	UserGroupId     string   `json:"user_group_id,omitempty"`
	Name            string   `json:"name,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *GroupEvent) MarshalJSON() ([]byte, error) {
	type cp GroupEvent
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type Scope struct {
	Departments     []*Department `json:"departments,omitempty"`
	Users           []*User       `json:"users,omitempty"`
	UserGroups      []*UserGroup  `json:"user_groups,omitempty"`
	ForceSendFields []string      `json:"-"`
}

func (s *Scope) MarshalJSON() ([]byte, error) {
	type cp Scope
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type UserCustomAttr struct {
	Type            string               `json:"type,omitempty"`
	Id              string               `json:"id,omitempty"`
	Value           *UserCustomAttrValue `json:"value,omitempty"`
	ForceSendFields []string             `json:"-"`
}

func (s *UserCustomAttr) MarshalJSON() ([]byte, error) {
	type cp UserCustomAttr
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type UserCustomAttrValue struct {
	Text            string   `json:"text,omitempty"`
	Url             string   `json:"url,omitempty"`
	PcUrl           string   `json:"pc_url,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *UserCustomAttrValue) MarshalJSON() ([]byte, error) {
	type cp UserCustomAttrValue
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type UserGroup struct {
	UserGroupId     string   `json:"user_group_id,omitempty"`
	Name            string   `json:"name,omitempty"`
	Type            int      `json:"type,omitempty"`
	MemberCount     int      `json:"member_count,omitempty"`
	Status          int      `json:"status,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *UserGroup) MarshalJSON() ([]byte, error) {
	type cp UserGroup
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type UserGroupMember struct {
}

type UserOrder struct {
	DepartmentId    string   `json:"department_id,omitempty"`
	UserOrder       int      `json:"user_order,omitempty"`
	DepartmentOrder int      `json:"department_order,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *UserOrder) MarshalJSON() ([]byte, error) {
	type cp UserOrder
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type UserPosition struct {
	PositionCode       string   `json:"position_code,omitempty"`
	PositionName       string   `json:"position_name,omitempty"`
	DepartmentId       string   `json:"department_id,omitempty"`
	LeaderUserId       string   `json:"leader_user_id,omitempty"`
	LeaderPositionCode string   `json:"leader_position_code,omitempty"`
	IsMajor            bool     `json:"is_major,omitempty"`
	ForceSendFields    []string `json:"-"`
}

func (s *UserPosition) MarshalJSON() ([]byte, error) {
	type cp UserPosition
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type UserStatus struct {
	IsFrozen        bool     `json:"is_frozen,omitempty"`
	IsResigned      bool     `json:"is_resigned,omitempty"`
	IsActivated     bool     `json:"is_activated,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *UserStatus) MarshalJSON() ([]byte, error) {
	type cp UserStatus
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type Department struct {
	Name               string              `json:"name,omitempty"`
	I18nName           *DepartmentI18nName `json:"i18n_name,omitempty"`
	ParentDepartmentId string              `json:"parent_department_id,omitempty"`
	DepartmentId       string              `json:"department_id,omitempty"`
	OpenDepartmentId   string              `json:"open_department_id,omitempty"`
	LeaderUserId       string              `json:"leader_user_id,omitempty"`
	ChatId             string              `json:"chat_id,omitempty"`
	Order              int64               `json:"order,omitempty,string"`
	UnitIds            []string            `json:"unit_ids,omitempty"`
	MemberCount        int                 `json:"member_count,omitempty"`
	Status             *DepartmentStatus   `json:"status,omitempty"`
	CreateGroupChat    bool                `json:"create_group_chat,omitempty"`
	ForceSendFields    []string            `json:"-"`
}

func (s *Department) MarshalJSON() ([]byte, error) {
	type cp Department
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type DepartmentCount struct {
	DepartmentId          string   `json:"department_id,omitempty"`
	DirectDepartmentCount int      `json:"direct_department_count,omitempty"`
	DirectUserCount       int      `json:"direct_user_count,omitempty"`
	DepartmentCount       int      `json:"department_count,omitempty"`
	UserCount             int      `json:"user_count,omitempty"`
	ForceSendFields       []string `json:"-"`
}

func (s *DepartmentCount) MarshalJSON() ([]byte, error) {
	type cp DepartmentCount
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type DepartmentEvent struct {
	Name               string `json:"name,omitempty"`
	ParentDepartmentId string `json:"parent_department_id,omitempty"`
	DepartmentId       string `json:"department_id,omitempty"`
	OpenDepartmentId   string `json:"open_department_id,omitempty"`
	LeaderUserId       string `json:"leader_user_id,omitempty"`
	ChatId             string `json:"chat_id,omitempty"`
	Order              int    `json:"order,omitempty"`

	Status          *DepartmentStatus `json:"status,omitempty"`
	ForceSendFields []string          `json:"-"`
}

func (s *DepartmentEvent) MarshalJSON() ([]byte, error) {
	type cp DepartmentEvent
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type DepartmentParent struct {
	DepartmentId    string   `json:"department_id,omitempty"`
	ParentIds       []string `json:"parent_ids,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *DepartmentParent) MarshalJSON() ([]byte, error) {
	type cp DepartmentParent
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type EmployeeTypeEnum struct {
	EnumId          string         `json:"enum_id,omitempty"`
	EnumValue       int64          `json:"enum_value,omitempty,string"`
	Content         string         `json:"content,omitempty"`
	EnumType        int            `json:"enum_type,omitempty"`
	EnumStatus      int            `json:"enum_status,omitempty"`
	I18nContent     []*I18nContent `json:"i18n_content,omitempty"`
	ForceSendFields []string       `json:"-"`
}

func (s *EmployeeTypeEnum) MarshalJSON() ([]byte, error) {
	type cp EmployeeTypeEnum
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type I18nContent struct {
	Locale          string   `json:"locale,omitempty"`
	Value           string   `json:"value,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *I18nContent) MarshalJSON() ([]byte, error) {
	type cp I18nContent
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type NotificationOption struct {
	Channels        []string `json:"channels,omitempty"`
	Language        string   `json:"language,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *NotificationOption) MarshalJSON() ([]byte, error) {
	type cp NotificationOption
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type OldDepartmentObject struct {
	Status           *DepartmentStatus `json:"status,omitempty"`
	OpenDepartmentId string            `json:"open_department_id,omitempty"`
	ForceSendFields  []string          `json:"-"`
}

func (s *OldDepartmentObject) MarshalJSON() ([]byte, error) {
	type cp OldDepartmentObject
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type OldUserObject struct {
	DepartmentIds   []string `json:"department_ids,omitempty"`
	OpenId          string   `json:"open_id,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *OldUserObject) MarshalJSON() ([]byte, error) {
	type cp OldUserObject
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type UserEvent struct {
	OpenId string `json:"open_id,omitempty"`
	UserId string `json:"user_id,omitempty"`
	Name   string `json:"name,omitempty"`
	EnName string `json:"en_name,omitempty"`
	Email  string `json:"email,omitempty"`
	Mobile string `json:"mobile,omitempty"`

	Gender        int         `json:"gender,omitempty"`
	Avatar        *AvatarInfo `json:"avatar,omitempty"`
	Status        *UserStatus `json:"status,omitempty"`
	DepartmentIds []string    `json:"department_ids,omitempty"`
	LeaderUserId  string      `json:"leader_user_id,omitempty"`
	City          string      `json:"city,omitempty"`
	Country       string      `json:"country,omitempty"`
	WorkStation   string      `json:"work_station,omitempty"`
	JoinTime      int         `json:"join_time,omitempty"`

	EmployeeNo   string `json:"employee_no,omitempty"`
	EmployeeType int    `json:"employee_type,omitempty"`

	Orders []*UserOrder `json:"orders,omitempty"`

	CustomAttrs     []*UserCustomAttr `json:"custom_attrs,omitempty"`
	ForceSendFields []string          `json:"-"`
}

func (s *UserEvent) MarshalJSON() ([]byte, error) {
	type cp UserEvent
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type User struct {
	UnionId         string      `json:"union_id,omitempty"`
	UserId          string      `json:"user_id,omitempty"`
	OpenId          string      `json:"open_id,omitempty"`
	Name            string      `json:"name,omitempty"`
	EnName          string      `json:"en_name,omitempty"`
	Email           string      `json:"email,omitempty"`
	Mobile          string      `json:"mobile,omitempty"`
	MobileVisible   bool        `json:"mobile_visible,omitempty"`
	Gender          int         `json:"gender,omitempty"`
	AvatarKey       string      `json:"avatar_key,omitempty"`
	Avatar          *AvatarInfo `json:"avatar,omitempty"`
	Status          *UserStatus `json:"status,omitempty"`
	DepartmentIds   []string    `json:"department_ids,omitempty"`
	LeaderUserId    string      `json:"leader_user_id,omitempty"`
	City            string      `json:"city,omitempty"`
	Country         string      `json:"country,omitempty"`
	WorkStation     string      `json:"work_station,omitempty"`
	JoinTime        int         `json:"join_time,omitempty"`
	IsTenantManager bool        `json:"is_tenant_manager,omitempty"`
	EmployeeNo      string      `json:"employee_no,omitempty"`
	EmployeeType    int         `json:"employee_type,omitempty"`

	Orders          []*UserOrder      `json:"orders,omitempty"`
	CustomAttrs     []*UserCustomAttr `json:"custom_attrs,omitempty"`
	EnterpriseEmail string            `json:"enterprise_email,omitempty"`

	JobTitle             string              `json:"job_title,omitempty"`
	NeedSendNotification bool                `json:"need_send_notification,omitempty"`
	NotificationOption   *NotificationOption `json:"notification_option,omitempty"`
	IsFrozen             bool                `json:"is_frozen,omitempty"`
	ForceSendFields      []string            `json:"-"`
}

func (s *User) MarshalJSON() ([]byte, error) {
	type cp User
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type DepartmentGetResult struct {
	Department *Department `json:"department,omitempty"`
}

type UserDeleteReqBody struct {
	DepartmentChatAcceptorUserId string   `json:"department_chat_acceptor_user_id,omitempty"`
	ExternalChatAcceptorUserId   string   `json:"external_chat_acceptor_user_id,omitempty"`
	DocsAcceptorUserId           string   `json:"docs_acceptor_user_id,omitempty"`
	CalendarAcceptorUserId       string   `json:"calendar_acceptor_user_id,omitempty"`
	ApplicationAcceptorUserId    string   `json:"application_acceptor_user_id,omitempty"`
	HelpdeskAcceptorUserId       string   `json:"helpdesk_acceptor_user_id,omitempty"`
	ForceSendFields              []string `json:"-"`
}

func (s *UserDeleteReqBody) MarshalJSON() ([]byte, error) {
	type cp UserDeleteReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type DepartmentListResult struct {
	HasMore   bool          `json:"has_more,omitempty"`
	PageToken string        `json:"page_token,omitempty"`
	Items     []*Department `json:"items,omitempty"`
}

type UserGroupUpdateUserGroupIdReqBody struct {
	NewUserGroupId  string   `json:"new_user_group_id,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *UserGroupUpdateUserGroupIdReqBody) MarshalJSON() ([]byte, error) {
	type cp UserGroupUpdateUserGroupIdReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type DepartmentPatchResult struct {
	Department *Department `json:"department,omitempty"`
}

type UserUpdateResult struct {
	User *User `json:"user,omitempty"`
}

type UserCreateResult struct {
	User *User `json:"user,omitempty"`
}

type UserPatchResult struct {
	User *User `json:"user,omitempty"`
}

type UserUpdateUserIdReqBody struct {
	NewUserId       string   `json:"new_user_id,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *UserUpdateUserIdReqBody) MarshalJSON() ([]byte, error) {
	type cp UserUpdateUserIdReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type UserGroupMemberCreateReqBody struct {
	UserId          string   `json:"user_id,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *UserGroupMemberCreateReqBody) MarshalJSON() ([]byte, error) {
	type cp UserGroupMemberCreateReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type DepartmentCreateResult struct {
	Department *Department `json:"department,omitempty"`
}

type UserGroupPatchResult struct {
	UserGroup *UserGroup `json:"user_group,omitempty"`
}

type DepartmentUnitPatchReqBody struct {
	UnitType        string   `json:"unit_type,omitempty"`
	UnitName        string   `json:"unit_name,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *DepartmentUnitPatchReqBody) MarshalJSON() ([]byte, error) {
	type cp DepartmentUnitPatchReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type DepartmentUnitPatchResult struct {
	DepartmentUnit *DepartmentUnit `json:"department_unit,omitempty"`
}

type UserGroupCreateResult struct {
	UserGroup *UserGroup `json:"user_group,omitempty"`
}

type UserGroupListResult struct {
	HasMore   bool         `json:"has_more,omitempty"`
	PageToken string       `json:"page_token,omitempty"`
	Items     []*UserGroup `json:"items,omitempty"`
}

type UserGetResult struct {
	User *User `json:"user,omitempty"`
}

type UserGroupMemberListResult struct {
	HasMore   bool    `json:"has_more,omitempty"`
	PageToken string  `json:"page_token,omitempty"`
	Items     []*User `json:"items,omitempty"`
}

type DepartmentUpdateResult struct {
	Department *Department `json:"department,omitempty"`
}

type DepartmentUnitCreateResult struct {
	DepartmentUnit *DepartmentUnit `json:"department_unit,omitempty"`
}

type UserListResult struct {
	HasMore   bool    `json:"has_more,omitempty"`
	PageToken string  `json:"page_token,omitempty"`
	Items     []*User `json:"items,omitempty"`
}

type UserGroupGetResult struct {
	UserGroup *UserGroup `json:"user_group,omitempty"`
}

type DepartmentParentResult struct {
	HasMore   bool          `json:"has_more,omitempty"`
	PageToken string        `json:"page_token,omitempty"`
	Items     []*Department `json:"items,omitempty"`
}

type DepartmentUpdateDepartmentIdReqBody struct {
	NewDepartmentId string   `json:"new_department_id,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *DepartmentUpdateDepartmentIdReqBody) MarshalJSON() ([]byte, error) {
	type cp DepartmentUpdateDepartmentIdReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type DepartmentUnbindDepartmentChatReqBody struct {
	DepartmentId    string   `json:"department_id,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *DepartmentUnbindDepartmentChatReqBody) MarshalJSON() ([]byte, error) {
	type cp DepartmentUnbindDepartmentChatReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type DepartmentSearchReqBody struct {
	Query           string   `json:"query,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *DepartmentSearchReqBody) MarshalJSON() ([]byte, error) {
	type cp DepartmentSearchReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type DepartmentSearchResult struct {
	Items     []*Department `json:"items,omitempty"`
	PageToken string        `json:"page_token,omitempty"`
	HasMore   bool          `json:"has_more,omitempty"`
}

type EmployeeTypeEnumListResult struct {
	Items     []*EmployeeTypeEnum `json:"items,omitempty"`
	HasMore   bool                `json:"has_more,omitempty"`
	PageToken string              `json:"page_token,omitempty"`
}

type DepartmentBatchGetReqBody struct {
	DepartmentIdList []string `json:"department_id_list,omitempty"`
	ForceSendFields  []string `json:"-"`
}

func (s *DepartmentBatchGetReqBody) MarshalJSON() ([]byte, error) {
	type cp DepartmentBatchGetReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type DepartmentBatchGetResult struct {
	Items         []*Department `json:"items,omitempty"`
	InvalidIdList []string      `json:"invalid_id_list,omitempty"`
}

type UserBatchGetReqBody struct {
	UserIdList      []string `json:"user_id_list,omitempty"`
	EmailList       []string `json:"email_list,omitempty"`
	UidList         []string `json:"uid_list,omitempty"`
	ForceSendFields []string `json:"-"`
}

func (s *UserBatchGetReqBody) MarshalJSON() ([]byte, error) {
	type cp UserBatchGetReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type UserBatchGetResult struct {
	UserList         []*User  `json:"user_list,omitempty"`
	InvalidIdList    []string `json:"invalid_id_list,omitempty"`
	InvalidEmailList []string `json:"invalid_email_list,omitempty"`
	InvalidUidList   []string `json:"invalid_uid_list,omitempty"`
}

type DepartmentBatchParentReqBody struct {
	DeptIdList       []string `json:"dept_id_list,omitempty"`
	DepartmentIdType string   `json:"department_id_type,omitempty"`
	ForceSendFields  []string `json:"-"`
}

func (s *DepartmentBatchParentReqBody) MarshalJSON() ([]byte, error) {
	type cp DepartmentBatchParentReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type DepartmentBatchParentResult struct {
	ParentList []*DepartmentParent `json:"parent_list,omitempty"`
}

type DepartmentListChildrenResult struct {
	HasMore   bool          `json:"has_more,omitempty"`
	PageToken string        `json:"page_token,omitempty"`
	Items     []*Department `json:"items,omitempty"`
}

type UserScanResult struct {
	Items     []*User `json:"items,omitempty"`
	HasMore   bool    `json:"has_more,omitempty"`
	PageToken string  `json:"page_token,omitempty"`
}

type DepartmentUsersReqBody struct {
	UserIdType       string   `json:"user_id_type,omitempty"`
	DepartmentIdType string   `json:"department_id_type,omitempty"`
	DeptIdList       []string `json:"dept_id_list,omitempty"`
	ForceSendFields  []string `json:"-"`
}

func (s *DepartmentUsersReqBody) MarshalJSON() ([]byte, error) {
	type cp DepartmentUsersReqBody
	raw := cp(*s)
	return tools.MarshalJSON(raw, s.ForceSendFields)
}

type DepartmentUsersResult struct {
	Users     []*User `json:"users,omitempty"`
	HasMore   bool    `json:"has_more,omitempty"`
	PageToken string  `json:"page_token,omitempty"`
}

type EmployeeTypeEnumCreateResult struct {
	EmployeeTypeEnum *EmployeeTypeEnum `json:"employee_type_enum,omitempty"`
}

type EmployeeTypeEnumUpdateResult struct {
	EmployeeTypeEnum *EmployeeTypeEnum `json:"employee_type_enum,omitempty"`
}

type DepartmentCreatedEventData struct {
	Object *DepartmentEvent `json:"object,omitempty"`
}

type DepartmentCreatedEvent struct {
	*model.BaseEventV2
	Event *DepartmentCreatedEventData `json:"event"`
}

type DepartmentDeletedEventData struct {
	Object    *DepartmentEvent     `json:"object,omitempty"`
	OldObject *OldDepartmentObject `json:"old_object,omitempty"`
}

type DepartmentDeletedEvent struct {
	*model.BaseEventV2
	Event *DepartmentDeletedEventData `json:"event"`
}

type DepartmentUpdatedEventData struct {
	Object    *DepartmentEvent `json:"object,omitempty"`
	OldObject *DepartmentEvent `json:"old_object,omitempty"`
}

type DepartmentUpdatedEvent struct {
	*model.BaseEventV2
	Event *DepartmentUpdatedEventData `json:"event"`
}

type UserCreatedEventData struct {
	Object *UserEvent `json:"object,omitempty"`
}

type UserCreatedEvent struct {
	*model.BaseEventV2
	Event *UserCreatedEventData `json:"event"`
}

type UserDeletedEventData struct {
	Object    *UserEvent     `json:"object,omitempty"`
	OldObject *OldUserObject `json:"old_object,omitempty"`
}

type UserDeletedEvent struct {
	*model.BaseEventV2
	Event *UserDeletedEventData `json:"event"`
}

type UserUpdatedEventData struct {
	Object    *UserEvent `json:"object,omitempty"`
	OldObject *UserEvent `json:"old_object,omitempty"`
}

type UserUpdatedEvent struct {
	*model.BaseEventV2
	Event *UserUpdatedEventData `json:"event"`
}

type UserGroupCreatedEventData struct {
	Object *GroupEvent `json:"object,omitempty"`
}

type UserGroupCreatedEvent struct {
	*model.BaseEventV2
	Event *UserGroupCreatedEventData `json:"event"`
}

type UserGroupDeletedEventData struct {
	Object *GroupEvent `json:"object,omitempty"`
}

type UserGroupDeletedEvent struct {
	*model.BaseEventV2
	Event *UserGroupDeletedEventData `json:"event"`
}

type UserGroupUpdatedEventData struct {
	Object    *GroupEvent `json:"object,omitempty"`
	OldObject *GroupEvent `json:"old_object,omitempty"`
}

type UserGroupUpdatedEvent struct {
	*model.BaseEventV2
	Event *UserGroupUpdatedEventData `json:"event"`
}

type ScopeUpdatedEventData struct {
	Added   *Scope `json:"added,omitempty"`
	Removed *Scope `json:"removed,omitempty"`
}

type ScopeUpdatedEvent struct {
	*model.BaseEventV2
	Event *ScopeUpdatedEventData `json:"event"`
}

type UserGroupMemberChangedEventData struct {
	Object *GroupEvent `json:"object,omitempty"`
}

type UserGroupMemberChangedEvent struct {
	*model.BaseEventV2
	Event *UserGroupMemberChangedEventData `json:"event"`
}

type CustomAttrEventUpdatedEventData struct {
	Object    *CustomAttrEvent `json:"object,omitempty"`
	OldObject *CustomAttrEvent `json:"old_object,omitempty"`
}

type CustomAttrEventUpdatedEvent struct {
	*model.BaseEventV2
	Event *CustomAttrEventUpdatedEventData `json:"event"`
}

type EmployeeTypeEnumActivedEventData struct {
	OldEnum *EmployeeTypeEnum `json:"old_enum,omitempty"`
	NewEnum *EmployeeTypeEnum `json:"new_enum,omitempty"`
}

type EmployeeTypeEnumActivedEvent struct {
	*model.BaseEventV2
	Event *EmployeeTypeEnumActivedEventData `json:"event"`
}

type EmployeeTypeEnumCreatedEventData struct {
	NewEnum *EmployeeTypeEnum `json:"new_enum,omitempty"`
}

type EmployeeTypeEnumCreatedEvent struct {
	*model.BaseEventV2
	Event *EmployeeTypeEnumCreatedEventData `json:"event"`
}

type EmployeeTypeEnumDeactivatedEventData struct {
	OldEnum *EmployeeTypeEnum `json:"old_enum,omitempty"`
	NewEnum *EmployeeTypeEnum `json:"new_enum,omitempty"`
}

type EmployeeTypeEnumDeactivatedEvent struct {
	*model.BaseEventV2
	Event *EmployeeTypeEnumDeactivatedEventData `json:"event"`
}

type EmployeeTypeEnumDeletedEventData struct {
	OldEnum *EmployeeTypeEnum `json:"old_enum,omitempty"`
}

type EmployeeTypeEnumDeletedEvent struct {
	*model.BaseEventV2
	Event *EmployeeTypeEnumDeletedEventData `json:"event"`
}

type EmployeeTypeEnumUpdatedEventData struct {
	OldEnum *EmployeeTypeEnum `json:"old_enum,omitempty"`
	NewEnum *EmployeeTypeEnum `json:"new_enum,omitempty"`
}

type EmployeeTypeEnumUpdatedEvent struct {
	*model.BaseEventV2
	Event *EmployeeTypeEnumUpdatedEventData `json:"event"`
}
