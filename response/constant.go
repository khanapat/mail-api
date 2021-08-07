package response

const _externalErrorCode = 4000
const _internalErrorCode = 5000

const (
	SuccessCode uint64 = 2000

	ErrInvalidRequestCode uint64 = _externalErrorCode + 1

	ErrDatabaseCode     uint64 = _internalErrorCode + 1
	ErrOperationCode    uint64 = _internalErrorCode + 2
	ErrSmtpSendMailCode uint64 = _internalErrorCode + 3
	ErrThirdPartyCode   uint64 = _internalErrorCode + 4
)

const (
	SuccessMessageEN           string = "Success."
	ErrInternalServerMessageEN string = "Internal server error."
	// BasicAuthen
	ErrBasicAuthenticationMessageEN string = "Authentication failed."
	// Mail
	SuccessSendMailMessageEN string = "Send e-mail Success."
	ErrTemplateNoMatchEN     string = "Template e-mail not found."
	ErrSmtpSendMailMessageEN string = "Cannot send e-mail."
	// Desc
	ErrRequestDataDescEN    string = "Please check request data again."
	ErrContactAdminDescEN   string = "Please contact administrator for more information."
	ErrThirdPartyDescEN     string = "Service is unavailable. Please try again later."
	ErrAuthenticationDescEN string = "Unable to access data. Please check user & password."
)

const (
	SuccessMessageTH           string = "สำเร็ว."
	ErrInternalServerMessageTH string = "มีข้อผิดพลาดภายในเซิร์ฟเวอร์."
	// BasicAuthen
	ErrBasicAuthenticationMessageTH string = "ยืนยันตัวตนล้มเหลว"
	// Mail
	SuccessSendMailMessageTH string = "ส่งอีเมล์สำเร็จ."
	ErrTemplateNoMatchTH     string = "ไม่พบแบบฟอร์มอีเมล์."
	ErrSmtpSendMailMessageTH string = "ไม่สามารถส่งอีเมล์ได้."
	// Desc
	ErrRequestDataDescTH    string = "กรุณาตรวจสอบข้อมูลอีกครั้ง."
	ErrContactAdminDescTH   string = "กรุณาติดต่อเจ้าหน้าที่ดูแลระบบเพื่อรับข้อมูลเพิ่มเติม."
	ErrThirdPartyDescTH     string = "ไม่สามารถใช้บริการได้. กรุณาทำรายการใหม่อีกครั้งภายหลัง."
	ErrAuthenticationDescTH string = "ไม่สามารถเข้าถึงข้อมูลได้. กรุณาตรวจสอบรหัสผู้ใช้งานใหม่อีกครั้ง."
)

const (
	ValidateFieldError string = "Invalid Parameters"
)
