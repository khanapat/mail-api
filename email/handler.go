package email

import (
	"bytes"
	"fmt"
	"mail-api/internal/handler"
	"mail-api/response"
	"net/http"
	"text/template"
)

type mailhandler struct {
}

func NewMailHandler() *mailhandler {
	return &mailhandler{}
}

func (s *mailhandler) SendVerifyEmail(c *handler.Ctx) error {
	var req SendVerifyEmailRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).SendMailValidateReq, err.Error()))
	}
	if err := req.validate(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).SendMailValidateReq, err.Error()))
	}

	templateData := struct {
		Name string
		Link string
	}{
		Name: req.Body.Name,
		Link: req.Body.Link,
	}

	path := fmt.Sprintf("./template/%s", req.Template)
	t, err := template.ParseFiles(path)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).SendMailTemplateNoMatch, err.Error()))
	}
	buf := new(bytes.Buffer)
	if err := t.Execute(buf, templateData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).SendMailTemplateNoMatch, err.Error()))
	}

	if err := SendEmail(req.From, req.To, req.Subject, buf.String(), req.Auth); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).SendMailSmtpFail, err.Error()))
	}

	c.Log().Info(fmt.Sprintf("Send Verify Email to %s success.", req.To))
	return c.Status(http.StatusOK).JSON(response.NewResponse(response.ResponseContextLocale(c.Context()).SendMailSuccess, nil))
}

func (s *mailhandler) SendOtp(c *handler.Ctx) error {
	var req SendOtpRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).SendMailValidateReq, err.Error()))
	}
	if err := req.validate(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).SendMailValidateReq, err.Error()))
	}

	templateData := struct {
		Name  string
		RefNo string
		Otp   string
	}{
		Name:  req.Body.Name,
		RefNo: req.Body.RefNo,
		Otp:   req.Body.Otp,
	}

	path := fmt.Sprintf("./template/%s", req.Template)
	t, err := template.ParseFiles(path)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).SendMailTemplateNoMatch, err.Error()))
	}
	buf := new(bytes.Buffer)
	if err := t.Execute(buf, templateData); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).SendMailTemplateNoMatch, err.Error()))
	}

	if err := SendEmail(req.From, req.To, req.Subject, buf.String(), req.Auth); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).SendMailSmtpFail, err.Error()))
	}

	c.Log().Info(fmt.Sprintf("Send otp to %s success.", req.To))
	return c.Status(http.StatusOK).JSON(response.NewResponse(response.ResponseContextLocale(c.Context()).SendMailSuccess, nil))
}

func (s *mailhandler) SendMarginCall(c *handler.Ctx) error {
	var req SendMarginCallRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).SendMailValidateReq, err.Error()))
	}
	if err := req.validate(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).SendMailValidateReq, err.Error()))
	}

	t, err := template.ParseFiles(fmt.Sprintf("./template/%s", req.Template))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).SendMailTemplateNoMatch, err.Error()))
	}
	buf := new(bytes.Buffer)
	if err := t.Execute(buf, req.Body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).SendMailTemplateNoMatch, err.Error()))
	}

	if err := SendEmail(req.From, req.To, req.Subject, buf.String(), req.Auth); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.NewErrResponse(response.ResponseContextLocale(c.Context()).SendMailSmtpFail, err.Error()))
	}

	c.Log().Info(fmt.Sprintf("Send margin call to %s success.", req.To))
	return c.Status(http.StatusOK).JSON(response.NewResponse(response.ResponseContextLocale(c.Context()).SendMailSuccess, nil))
}
