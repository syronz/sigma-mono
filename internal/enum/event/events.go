package event

import "fmt"

// Event is used for type of events
type Event string

// event enums
const (
	UserCreate Event = "user-create"
	UserUpdate Event = "user-update"
	UserDelete Event = "user-delete"
	UserList   Event = "user-list"
	UserView   Event = "user-view"
	UserExcel  Event = "user-excel"

	LicenseCreate Event = "license-create"

	CompanyCreate   Event = "company-create"
	CompanyUpdate   Event = "company-update"
	CompanyDelete   Event = "company-delete"
	CompanyList     Event = "company-list"
	CompanyView     Event = "company-view"
	CompanyExcel    Event = "company-excel"
	CompanyRegister Event = "company-register"

	StationCreate Event = "station-create"
	StationUpdate Event = "station-update"
	StationDelete Event = "station-delete"
	StationList   Event = "station-list"
	StationView   Event = "station-view"
	StationExcel  Event = "station-excel"

	NodeCreate   Event = "node-create"
	NodeUpdate   Event = "node-update"
	NodeDelete   Event = "node-delete"
	NodeList     Event = "node-list"
	NodeView     Event = "node-view"
	NodeExcel    Event = "node-excel"
	NodeActivate Event = "node-activate"

	SyncSessionCreate Event = "syncSession-create"
	SyncSessionUpdate Event = "syncSession-update"
	SyncSessionDelete Event = "syncSession-delete"
	SyncSessionList   Event = "syncSession-list"
	SyncSessionView   Event = "syncSession-view"
	SyncSessionExcel  Event = "syncSession-excel"

	AccountCreate Event = "account-create"
	AccountUpdate Event = "account-update"
	AccountDelete Event = "account-delete"
	AccountList   Event = "account-list"
	AccountView   Event = "account-view"
	AccountExcel  Event = "account-excel"

	SettingCreate Event = "setting-create"
	SettingUpdate Event = "setting-update"
	SettingDelete Event = "setting-delete"
	SettingList   Event = "setting-list"
	SettingView   Event = "setting-view"
	SettingExcel  Event = "setting-excel"

	PhoneCreate Event = "phone-create"
	PhoneUpdate Event = "phone-update"
	PhoneDelete Event = "phone-delete"
	PhoneList   Event = "phone-list"
	PhoneView   Event = "phone-view"
	PhoneExcel  Event = "phone-excel"

	VersionCreate Event = "version-create"
	VersionUpdate Event = "version-update"
	VersionDelete Event = "version-delete"
	VersionList   Event = "version-list"
	VersionView   Event = "version-view"
	VersionExcel  Event = "version-excel"

	RoleCreate Event = "role-create"
	RoleUpdate Event = "role-update"
	RoleDelete Event = "role-delete"
	RoleList   Event = "role-list"
	RoleView   Event = "role-view"
	RoleExcel  Event = "role-excel"

	Login Event = "login"
)

func (e *Event) String() string {
	return fmt.Sprint(*e)
}
