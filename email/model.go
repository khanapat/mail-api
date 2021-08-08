package email

import (
	"fmt"
	"mail-api/response"
	"unicode/utf8"

	"github.com/pkg/errors"
)

type SendOtpRequest struct {
	From     string             `json:"from" example:"icfin999@gmail.com"`
	To       []string           `json:"to" example:"yoisak09446@gmail.com"`
	Subject  string             `json:"subject" example:"otp request"`
	Template string             `json:"template" example:"otp.html"`
	Body     BodySendOtpRequest `json:"body"`
	Auth     bool               `json:"auth" example:"true"`
}

type BodySendOtpRequest struct {
	Name  string `json:"name" example:"trust"`
	RefNo string `json:"refNo" example:"tog2C7"`
	Otp   string `json:"otp" example:"999999"`
}

func (req *SendOtpRequest) validate() error {
	if utf8.RuneCountInString(req.From) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'from' must be REQUIRED field but the input is '%v'", req.From)), response.ValidateFieldError)
	}
	if len(req.To) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'to' must be REQUIRED field but the input is '%v'", req.To)), response.ValidateFieldError)
	}
	if utf8.RuneCountInString(req.Subject) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'subject' must be REQUIRED field but the input is '%v'", req.Subject)), response.ValidateFieldError)
	}
	if utf8.RuneCountInString(req.Template) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'template' must be REQUIRED field but the input is '%v'", req.Template)), response.ValidateFieldError)
	}
	if utf8.RuneCountInString(req.Body.Name) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'name' must be REQUIRED field but the input is '%v'", req.Body.Name)), response.ValidateFieldError)
	}
	if utf8.RuneCountInString(req.Body.RefNo) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'refNo' must be REQUIRED field but the input is '%v'", req.Body.RefNo)), response.ValidateFieldError)
	}
	if utf8.RuneCountInString(req.Body.Otp) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'otp' must be REQUIRED field but the input is '%v'", req.Body.Otp)), response.ValidateFieldError)
	}
	return nil
}

type SendVerifyEmailRequest struct {
	From     string                     `json:"from" example:"icfin999@gmail.com"`
	To       []string                   `json:"to" example:"yoisak09446@gmail.com"`
	Subject  string                     `json:"subject" example:"otp request"`
	Template string                     `json:"template" example:"otp.html"`
	Body     BodySendVerifyEmailRequest `json:"body"`
	Auth     bool                       `json:"auth" example:"true"`
}

type BodySendVerifyEmailRequest struct {
	Name string `json:"name" example:"trust momo"`
	Link string `json:"link" example:"www.lending.com/WERaOJOsfX"`
}

func (req *SendVerifyEmailRequest) validate() error {
	if utf8.RuneCountInString(req.From) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'from' must be REQUIRED field but the input is '%v'", req.From)), response.ValidateFieldError)
	}
	if len(req.To) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'to' must be REQUIRED field but the input is '%v'", req.To)), response.ValidateFieldError)
	}
	if utf8.RuneCountInString(req.Subject) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'subject' must be REQUIRED field but the input is '%v'", req.Subject)), response.ValidateFieldError)
	}
	if utf8.RuneCountInString(req.Template) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'template' must be REQUIRED field but the input is '%v'", req.Template)), response.ValidateFieldError)
	}
	if utf8.RuneCountInString(req.Body.Name) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'name' must be REQUIRED field but the input is '%v'", req.Body.Name)), response.ValidateFieldError)
	}
	if utf8.RuneCountInString(req.Body.Link) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'link' must be REQUIRED field but the input is '%v'", req.Body.Link)), response.ValidateFieldError)
	}
	return nil
}

type SendMarginCallRequest struct {
	From     string                `json:"from" example:"icfin999@gmail.com"`
	To       []string              `json:"to" example:"yoisak09446@gmail.com"`
	Subject  string                `json:"subject" example:"margin call"`
	Template string                `json:"template" example:"margin-call.html"`
	Body     BodyMarginCallRequest `json:"body"`
	Auth     bool                  `json:"auth" example:"true"`
}

type BodyMarginCallRequest struct {
	Name            string `json:"name" example:"trust momo"`
	MarginCallCount int    `json:"marginCallCount" example:"1"`
}

func (req *SendMarginCallRequest) validate() error {
	if utf8.RuneCountInString(req.From) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'from' must be REQUIRED field but the input is '%v'", req.From)), response.ValidateFieldError)
	}
	if len(req.To) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'to' must be REQUIRED field but the input is '%v'", req.To)), response.ValidateFieldError)
	}
	if utf8.RuneCountInString(req.Subject) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'subject' must be REQUIRED field but the input is '%v'", req.Subject)), response.ValidateFieldError)
	}
	if utf8.RuneCountInString(req.Template) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'template' must be REQUIRED field but the input is '%v'", req.Template)), response.ValidateFieldError)
	}
	if utf8.RuneCountInString(req.Body.Name) == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'name' must be REQUIRED field but the input is '%v'", req.Body.Name)), response.ValidateFieldError)
	}
	if req.Body.MarginCallCount == 0 {
		return errors.Wrapf(errors.New(fmt.Sprintf("'marginCallCount' must be REQUIRED field but the input is '%v'", req.Body.MarginCallCount)), response.ValidateFieldError)
	}
	return nil
}
