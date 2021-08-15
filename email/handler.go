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

// SendVerifyEmail
// @Summary Send Verify Email
// @Description send token to verify email
// @Tags Email
// @Accept json
// @Produce json
// @Param SendVerifyEmail body email.SendVerifyEmailRequest true "request body to send verify email"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.ErrResponse "Bad Request"
// @Failure 500 {object} response.ErrResponse "Internal Server Error"
// @Router /verification [post]
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

// SendOTP
// @Summary Send OTP
// @Description send otp to email
// @Tags Email
// @Accept json
// @Produce json
// @Param SendOTP body email.SendOtpRequest true "request body to send otp"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.ErrResponse "Bad Request"
// @Failure 500 {object} response.ErrResponse "Internal Server Error"
// @Router /otp [post]
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

// SendMarginCall
// @Summary Send Warning Margin Call
// @Description send warning margin call
// @Tags Email
// @Accept json
// @Produce json
// @Param SendMarginCall body email.SendMarginCallRequest true "request body to send warning margin call"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.ErrResponse "Bad Request"
// @Failure 500 {object} response.ErrResponse "Internal Server Error"
// @Router /margin-call [post]
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

// SendLiquidateFund
// @Summary Send Liquidate Fund
// @Description send liquidate fund
// @Tags Email
// @Accept json
// @Produce json
// @Param SendLiquidateFund body email.SendLiquidateFundRequest true "request body to send liquidate fund"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.ErrResponse "Bad Request"
// @Failure 500 {object} response.ErrResponse "Internal Server Error"
// @Router /liquidation [post]
func (s *mailhandler) SendLiquidateFund(c *handler.Ctx) error {
	var req SendLiquidateFundRequest
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

	c.Log().Info(fmt.Sprintf("Send liquidate fund to %s success.", req.To))
	return c.Status(http.StatusOK).JSON(response.NewResponse(response.ResponseContextLocale(c.Context()).SendMailSuccess, nil))
}

// SendSubscription
// @Summary Send Subscription
// @Description send subscription
// @Tags Email
// @Accept json
// @Produce json
// @Param SendSubscription body email.SendSubscriptionRequest true "request body to send subscription"
// @Success 200 {object} response.Response "Success"
// @Failure 400 {object} response.ErrResponse "Bad Request"
// @Failure 500 {object} response.ErrResponse "Internal Server Error"
// @Router /subscription [post]
func (s *mailhandler) SendSubscription(c *handler.Ctx) error {
	var req SendSubscriptionRequest
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

	c.Log().Info(fmt.Sprintf("Send Subscription to %s success.", req.To))
	return c.Status(http.StatusOK).JSON(response.NewResponse(response.ResponseContextLocale(c.Context()).SendMailSuccess, nil))
}
