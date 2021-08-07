package response

import (
	"context"
	"mail-api/common"
)

var (
	EN = Global{
		SendMailSuccess:         Response{Code: SuccessCode, Title: SuccessMessageEN},
		SendMailValidateReq:     ErrResponse{Code: ErrInvalidRequestCode, Title: ErrSmtpSendMailMessageEN, Description: ErrRequestDataDescEN},
		SendMailTemplateNoMatch: ErrResponse{Code: ErrSmtpSendMailCode, Title: ErrTemplateNoMatchEN, Description: ErrContactAdminDescEN},
		SendMailSmtpFail:        ErrResponse{Code: ErrSmtpSendMailCode, Title: ErrSmtpSendMailMessageEN, Description: ErrThirdPartyDescEN},
		InternalOperation:       ErrResponse{Code: ErrOperationCode, Title: ErrInternalServerMessageEN, Description: ErrContactAdminDescEN},
	}

	TH = Global{
		SendMailSuccess:         Response{Code: SuccessCode, Title: SuccessMessageTH},
		SendMailValidateReq:     ErrResponse{Code: ErrInvalidRequestCode, Title: ErrSmtpSendMailMessageTH, Description: ErrRequestDataDescTH},
		SendMailTemplateNoMatch: ErrResponse{Code: ErrSmtpSendMailCode, Title: ErrTemplateNoMatchTH, Description: ErrContactAdminDescTH},
		SendMailSmtpFail:        ErrResponse{Code: ErrSmtpSendMailCode, Title: ErrSmtpSendMailMessageTH, Description: ErrThirdPartyDescTH},
		InternalOperation:       ErrResponse{Code: ErrOperationCode, Title: ErrInternalServerMessageTH, Description: ErrContactAdminDescTH},
	}

	Language = map[interface{}]Global{
		"en": EN,
		"th": TH,
	}
)

type Global struct {
	SendMailSuccess         Response
	SendMailValidateReq     ErrResponse
	SendMailTemplateNoMatch ErrResponse
	SendMailSmtpFail        ErrResponse
	InternalOperation       ErrResponse
}

func ResponseContextLocale(ctx context.Context) *Global {
	v := ctx.Value(common.LocaleKey)
	if v == nil {
		return nil
	}
	l, ok := Language[v]
	if ok {
		return &l
	}
	return &EN
}
