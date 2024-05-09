package main

import (
	"fmt"
	"log"
	"time"

	"github.com/werbot/lime/pkg/fsutil"
	"github.com/werbot/lime/pkg/license"
)

var fixturePath = "../../../fixtures/licenses/"

func main() {
	// create new license key
	licenseData := license.License{
		IssuedBy:     "Werbot, Inc.",
		CustomerID:   "8ED96811-1804-4A13-9CE7-05874869A1CF",
		SubscriberID: "EED1CA19-4DC5-4376-83F5-61077B501961",
		Type:         "Enterprise",
		Limit: license.Limits{
			Limits: []license.Limit{
				{
					Key:   "servers",
					Value: 99,
				},
				{
					Key:   "companies",
					Value: 99,
				},
				{
					Key:   "users",
					Value: 99,
				},
				{
					Key:  "modules",
					List: []string{"module1", "module2", "module3"},
				},
			},
		},
		IssuedAt:  time.Date(2023, time.January, 10, 16, 55, 43, 277033166, time.UTC),
		ExpiresAt: time.Date(2050, time.May, 27, 16, 55, 43, 277033272, time.UTC),
	}

	licenseKey, err := createLicenseKey(licenseData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", licenseKey)

	// read license key
	dataFromLicense, err := readLicenseKey(licenseKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", dataFromLicense.Info())
}

func createLicenseKey(data license.License) ([]byte, error) {
	forDecode, err := license.DecodePrivateKey(fsutil.MustReadFile(fixturePath + "privateKey_ok.key"))
	if err != nil {
		return nil, err
	}

	forDecode.License = data
	return forDecode.Encode()
}

func readLicenseKey(licenseKey []byte) (*license.Public, error) {
	forEncode, err := license.DecodePublicKey(fsutil.MustReadFile(fixturePath + "publicKey_ok.key"))
	if err != nil {
		return nil, err
	}

	return forEncode.Decode(licenseKey)
}
