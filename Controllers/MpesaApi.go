package Controllers

import (
	"github.com/AndroidStudyOpenSource/mpesa-api-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	appKey    = "q5vLG7QGE3tLomGmNppDHIMWPHr5aYHD"
	appSecret = "CeMO8cYR11dC9gwm"
)

func MpesaExpress(c *gin.Context) {

	svc, err := mpesa.New(appKey, appSecret, mpesa.SANDBOX)
	if err != nil {
		panic(err)
	}

	res, err := svc.Simulation(mpesa.Express{
		BusinessShortCode: "174379",
		Password:          "MTc0Mzc5YmZiMjc5ZjlhYTliZGJjZjE1OGU5N2RkNzFhNDY3Y2QyZTBjODkzMDU5YjEwZjc4ZTZiNzJhZGExZWQyYzkxOTIwMTgwNDA5MDkzMDAy",
		Timestamp:         "20180409093002",
		TransactionType:   "CustomerPayBillOnline",
		Amount:            "1",
		PartyA:            "254740790736",
		PartyB:            "174379",
		PhoneNumber:       "254741790736",
		CallBackURL:       "https://wamu.co.ke",
		AccountReference:  "account",
		TransactionDesc:   "test",
	})

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":404,
			"message": "Failed to pay",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message": "Success",
		"res": res,

	})

}
