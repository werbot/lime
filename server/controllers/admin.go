package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/werbot/lime/license"
	"github.com/werbot/lime/server/middleware"
	"github.com/werbot/lime/server/models"
)

// MainHandler is a ...
func MainHandler(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(middleware.IdentityKey)

	if user == nil {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "🔑 Autorisation",
		})
	} else {
		customersList := models.CustomersList()
		c.HTML(http.StatusOK, "customers.html", gin.H{
			"title":     "🐱 Customers",
			"customers": customersList,
		})
	}

}

// CustomerSubscrptionList is a ...
func CustomerSubscrptionList(c *gin.Context) {
	customerID := c.Param("id")
	action := c.Param("action")
	subscriptionsList := models.SubscriptionsByCustomerID(customerID)
	licensesList := models.LicensesListBySubscriptionID(customerID)

	switch action {
	case "/":
	case "/new":
		month := time.Hour * 24 * 31
		modelTariff := models.Tariff{}

		_tariff, _ := modelTariff.FindTariffByID((*subscriptionsList)[0].ID)

		limit := license.Limits{
			Servers:   _tariff.Servers,
			Companies: _tariff.Companies,
			Users:     _tariff.Users,
		}
		metadata := []byte(`{"message": "test message"}`)
		_license := &license.License{
			Iss: (*subscriptionsList)[0].CustomerName,
			Cus: (*subscriptionsList)[0].StripeID,
			Sub: (*subscriptionsList)[0].TariffID,
			Typ: _tariff.Name,
			Lim: limit,
			Dat: metadata,
			Exp: time.Now().UTC().Add(month),
			Iat: time.Now().UTC(),
		}
		encoded, _ := _license.Encode(license.GetPrivateKey())

		hash := md5.Sum([]byte(encoded))
		licenseHash := hex.EncodeToString(hash[:])

		models.DeactivateLicenseBySubID((*subscriptionsList)[0].ID)
		key := &models.License{
			SubscriptionID: (*subscriptionsList)[0].ID,
			License:        encoded,
			Hash:           licenseHash,
			Status:         true,
		}
		key.SaveLicense()

		c.Redirect(http.StatusFound, "/admin/subscription/"+customerID+"/")

	default:
		c.Redirect(http.StatusFound, "/admin/subscription/"+customerID+"/")
	}

	c.HTML(http.StatusOK, "subscriptions.html", gin.H{
		"title":         "🧩 Subscription and Licenses by " + (*subscriptionsList)[0].CustomerName,
		"customerID":    customerID,
		"Subscriptions": subscriptionsList,
		"Licensies":     licensesList,
	})

}

// DownloadLicense is a ...
func DownloadLicense(c *gin.Context) {
	licenseID := c.Param("id")
	license := models.License{}

	uid, err := strconv.ParseUint(licenseID, 10, 32)
	if err != nil {
		panic(err)
	}

	license.FindLicenseByID(uint32(uid))

	body := string(license.License)
	reader := strings.NewReader(body)
	contentLength := int64(len(body))
	contentType := "application/octet-stream"
	extraHeaders := map[string]string{"Content-Disposition": `attachment; filename="` + license.Hash + `"`}
	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}
