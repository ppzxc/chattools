package inbound

type DeviceInfo struct {
	DeviceId        string `json:"device_id,omitempty" validate:"required"`
	BrowserId       string `json:"browser_id,omitempty" validate:"required"`
	UserAgent       string `json:"user_agent,omitempty" validate:"required"`
	OperationSystem string `json:"operation_system,omitempty" validate:"required,oneof=ios mac windows unix linux android blackberry etc"`
	Platform        string `json:"platform,omitempty" validate:"required,oneof=mobile desktop"`
}
