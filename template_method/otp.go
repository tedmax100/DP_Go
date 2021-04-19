package main

// OTP interface : define algorithm
type iOtp interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
	publishMetric()
}

type otp struct {
	iOtp         iOtp
	otpName      string
	hookFuncList []func(args []interface{})
}

// Template Method : 將最後的實作推遲至子類別中實現
func (o *otp) genAndSendOTPWithHooks(otpLength int, hooks ...func(args []interface{})) error {
	otp := o.iOtp.genRandomOTP(otpLength)
	o.hookFuncList = hooks
	o.iOtp.saveOTPCache(otp)
	message := o.iOtp.getMessage(otp)
	err := o.iOtp.sendNotification(message)
	if err != nil {
		return err
	}
	o.iOtp.publishMetric()
	for _, hook := range o.hookFuncList {
		hook([]interface{}{o.otpName})
	}
	return nil
}
