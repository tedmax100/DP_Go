package main

import (
	"testing"
)

const blackCat = "BlackCat"
const hct = "HCT"
const postOffice = "PostOffice"

type args struct {
	shipper string
	length  double
	width   double
	height  double
	weight  double
}

func TestCarShippingFee(t *testing.T) {
	testCases := []struct {
		name    string
		args    args
		want    double
		wantErr bool
	}{
		{
			name: "black cat with light weight",
			args: args{
				shipper: blackCat,
				length:  30,
				width:   20,
				height:  10,
				weight:  5,
			},
			want:    150,
			wantErr: false,
		},
		{
			name: "black cat with heavy weight",
			args: args{
				shipper: blackCat,
				length:  30,
				width:   20,
				height:  10,
				weight:  50,
			},
			want:    500,
			wantErr: false,
		},
		{
			name: "fedex with small size",
			args: args{
				shipper: hct,
				length:  30,
				width:   20,
				height:  10,
				weight:  50,
			},
			want:    144,
			wantErr: false,
		},
		{
			name: "fedex with huge size",
			args: args{
				shipper: hct,
				length:  110,
				width:   20,
				height:  10,
				weight:  50,
			},
			want:    984,
			wantErr: false,
		},
		{
			name: "post office by weight",
			args: args{
				shipper: postOffice,
				length:  100,
				width:   20,
				height:  10,
				weight:  3,
			},
			want:    110,
			wantErr: false,
		},
		{
			name: "unknown shipper",
			args: args{
				shipper: "unknown",
				length:  0,
				width:   0,
				height:  0,
				weight:  0,
			},
			want:    -1,
			wantErr: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			c := Cart{}
			got, err := c.ShippingFee(testCase.args.shipper, testCase.args.length, testCase.args.width, testCase.args.height, testCase.args.weight)
			if (err != nil) != testCase.wantErr {
				t.Errorf("ShippingFee() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}
			if got != testCase.want {
				t.Errorf("ShippingFee() got = %v, want %v", got, testCase.want)
			}
		})
	}
}
