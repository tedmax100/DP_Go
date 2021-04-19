package main

import "fmt"

func main() {
	smsOTP := &sms{}
	o := otp{
		iOtp:    smsOTP,
		otpName: "sms",
	}
	args := make([]interface{}, 0)
	args = append(args, "sms")

	o.genAndSendOTPWithHooks(4, func([]interface{}) { fmt.Println(args, "hook1 executed") })
	fmt.Printf("\n")
	emailOTP := &email{}
	o = otp{
		iOtp:    emailOTP,
		otpName: "email",
	}
	o.genAndSendOTPWithHooks(4, func([]interface{}) { fmt.Println(args, "hook1 executed") }, func([]interface{}) { fmt.Println(args, "hook2 executed") })
}
