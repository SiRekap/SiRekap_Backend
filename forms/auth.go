package forms

type UserRegisterData struct {
	Email    string `json:"email" binding:"required"`
	DeviceId string `json:"device_id" binding:"required"`
	Msisdn   string `json:"msisdn" binding:"required"`
	Password string `json:"password" binding:"required"`
}
