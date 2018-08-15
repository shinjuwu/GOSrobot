package tool

import (
	"log"
	"strings"

	"github.com/shopspring/decimal"
)

func Mul(v1 string, v2 string, place int32) string {
	dec1, err := decimal.NewFromString(v1)
	if err != nil {
		log.Fatal(err)
	}
	dec2, err := decimal.NewFromString(v2)
	if err != nil {
		log.Fatal(err)
	}

	dec3 := dec1.Mul(dec2)
	vTruncate := dec3.Truncate(place).String()
	result := addZeroByPlace(vTruncate, place)
	return result
}

func Sub(v1 string, v2 string, place int32) string {
	dec1, err := decimal.NewFromString(v1)
	if err != nil {
		log.Fatal(err)
	}
	dec2, err := decimal.NewFromString(v2)
	if err != nil {
		log.Fatal(err)
	}
	dec3 := dec1.Sub(dec2)
	vTruncate := dec3.Truncate(place).String()
	result := addZeroByPlace(vTruncate, place)
	return result
}

func Add(v1, v2 string, place int32) string {
	dec1, err := decimal.NewFromString(v1)
	if err != nil {
		log.Fatal(err)
	}
	dec2, err := decimal.NewFromString(v2)
	if err != nil {
		log.Fatal(err)
	}
	dec3 := dec1.Add(dec2)
	vTruncate := dec3.Truncate(place).String()
	result := addZeroByPlace(vTruncate, place)
	return result
}

func Div(v1 string, v2 string, place int32) string {
	dec1, err := decimal.NewFromString(v1)
	if err != nil {
		log.Fatal(err)
	}
	dec2, err := decimal.NewFromString(v2)
	if err != nil {
		log.Fatal(err)
	}

	dec3 := dec1.Div(dec2)
	vTruncate := dec3.Truncate(place).String()
	result := addZeroByPlace(vTruncate, place)
	return result
}

func addZeroByPlace(v string, place int32) string {
	strSplit := strings.Split(v, ".")
	if len(strSplit) > 2 {
		log.Fatal("Wrong number to add zero")
	}
	var suffix string
	if len(strSplit) == 1 {
		suffix = ""
	} else {
		suffix = strSplit[1]
	}

	startIndex := len(suffix)
	if len(suffix) < int(place) {
		for i := startIndex; i < int(place); i++ {
			suffix = suffix + "0"
		}
	}
	if place == 0 {
		return strSplit[0]
	}
	return strSplit[0] + "." + suffix
}

func GetWinAmount(v1, v2, v3 string, place int32) string {
	dec1, err := decimal.NewFromString(v1)
	if err != nil {
		log.Fatal(err)
	}
	dec2, err := decimal.NewFromString(v2)
	if err != nil {
		log.Fatal(err)
	}
	dec3, err := decimal.NewFromString(v3)
	if err != nil {
		log.Fatal(err)
	}
	decAmount := dec1.Mul(dec2).Div(dec3)
	vTruncate := decAmount.Truncate(place).String()
	result := addZeroByPlace(vTruncate, place)
	return result
}

func GetMultiplier(v1, v2, v3 string, place int32) string {
	dec1, err := decimal.NewFromString(v1)
	if err != nil {
		log.Fatal(err)
	}
	dec2, err := decimal.NewFromString(v2)
	if err != nil {
		log.Fatal(err)
	}
	dec3, err := decimal.NewFromString(v3)
	if err != nil {
		log.Fatal(err)
	}
	decAmount := dec1.Mul(dec2).Div(dec3)
	vTruncate := decAmount.Truncate(place).String()
	result := addZeroByPlace(vTruncate, place)
	return result
}

func GetNewBalance(v1, v2, v3 string, place int32) string {
	dec1, err := decimal.NewFromString(v1)
	if err != nil {
		log.Fatal(err)
	}
	dec2, err := decimal.NewFromString(v2)
	if err != nil {
		log.Fatal(err)
	}
	dec3, err := decimal.NewFromString(v3)
	if err != nil {
		log.Fatal(err)
	}
	decAmount := dec1.Sub(dec2).Add(dec3)
	vTruncate := decAmount.Truncate(place).String()
	result := addZeroByPlace(vTruncate, place)
	return result
}

func GetSum(data []float64, place int32) string {
	decSum := decimal.New(0, 0)
	for _, v := range data {
		dec := decimal.NewFromFloat(v)
		decSum = decSum.Add(dec)
	}
	vTruncate := decSum.Truncate(place).String()
	result := addZeroByPlace(vTruncate, place)
	return result
}
