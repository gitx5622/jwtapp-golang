package Middlewares

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"jwtapp/Config"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"strings"
	"time"
)

type SqlRes struct {
	Path        string
	AuthorityId string
	ApiId       uint
	Id          uint
}

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "qmPlus"
)

type CustomClaims struct {
	UUID        uuid.UUID
	ID          uint
	NickName    string
	AuthorityId string
	jwt.StandardClaims
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get("Authorization")
		if auth == "" {
			c.String(http.StatusForbidden, "No Authorization header provided")
			c.Abort()
			return
		}
		token := strings.TrimPrefix(auth, "Bearer ")
		if token == "" {
			c.JSON(http.StatusForbidden, gin.H{
				"msg":"Token is Empty",
				"reload": true,
			})
			c.Abort()
			return
		}
		j := NewJWT()
		// parseToken
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				//c.JSON(http.StatusForbidden, gin.H{
				//	"msg":"Token Expired",
				//	"reload": true,
				//})
				_, err = j.RefreshToken(token)
				//if err != nil {
				c.JSON(http.StatusBadRequest, gin.H {
					"code": "400",
					"msg":"Error refreshing token",
				})
				//}
				c.Abort()
				return
			}
			c.JSON(http.StatusForbidden, gin.H{
				"msg":"Token Error",
				"reload": true,
			})
			c.Abort()
			return
		}
		var sqlRes SqlRes
		row := Config.DB.Raw("SELECT apis.path,api_authorities.authority_id,api_authorities.api_id,apis.id FROM apis INNER JOIN api_authorities ON api_authorities.api_id = apis.id 	WHERE apis.path = ? AND	api_authorities.authority_id = ?", c.Request.RequestURI, claims.AuthorityId)
		err = row.Scan(&sqlRes).Error
		if fmt.Sprintf("%v", err) == "record not found" {
			c.JSON(http.StatusBadRequest, gin.H{"msg":"Not Found"})
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}


func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// Get Sign Key
func GetSignKey() string {
	return SignKey
}

// Set SignKey
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

//Create token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

//Parse  token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid

	}

}

// Refresh Token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
