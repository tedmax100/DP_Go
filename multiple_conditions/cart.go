package main

import "errors"

type double = float64

type Cart struct{}

type ShippingFeeFunc func(shipper string, length, width, height, weight double) (double, error)

func ShippingFactory(shipper string) (IShipping, error) {
	if shipper == "BlackCat" {
		return &BlackCat{}, nil
	}
	if shipper == "HCT" {
		return &HCT{}, nil
	}
	if shipper == "PostOffice" {
		return &PostOffice{}, nil
	}
	return nil, errors.New("shipper not exist")
}

// 根據物流廠商, 貨品尺寸(長寬高), 貨品重量來計算運費
// shipper : 物流廠商
// 貨品尺寸 : length, width, height
// 貨品重量 : weight
// Return : 運費, error
func (c *Cart) ShippingFee(shipper string, length, width, height, weight double) (double, error) {
	shipperObj, err := ShippingFactory(shipper)
	if err != nil {
		return -1, err
	}
	return shipperObj.ShippingFee(length, width, height, weight)
}

type IShipping interface {
	ShippingFee(double, double, double, double) (double, error)
}

type BlackCat struct {
}

func (b *BlackCat) ShippingFee(length, width, height, weight double) (double, error) {
	if weight > 20 {
		return 500, nil
	}
	return 100 + weight*10, nil
}

type HCT struct {
}

func (h *HCT) ShippingFee(length, width, height, weight double) (double, error) {
	var size double = length * width * height
	if length > 100 || width > 100 || height > 100 {
		return size*0.00002*1100 + 500, nil
	}
	return size * 0.00002 * 1200, nil
}

type PostOffice struct {
}

func (p *PostOffice) ShippingFee(length, width, height, weight double) (double, error) {
	var feeByWeight double = 80 + weight*10
	var size double = length * width * height
	var feeBySize = size * 0.00002 * 1100
	if feeByWeight < feeBySize {
		return feeByWeight, nil
	} else {
		return feeBySize, nil
	}
}
