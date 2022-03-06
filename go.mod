module github.com/hiroki-kondo-git/dayMemo_api_go

// +heroku goVersion go1.17
go 1.17

require (
	cloud.google.com/go/storage v1.18.2
	firebase.google.com/go v3.13.0+incompatible
	github.com/labstack/echo v3.3.10+incompatible
	google.golang.org/api v0.63.0
)

require gopkg.in/go-playground/assert.v1 v1.2.1 // indirect

require (
	cloud.google.com/go/firestore v1.6.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8
	gopkg.in/go-playground/validator.v9 v9.31.0
	gorm.io/driver/postgres v1.2.3
	gorm.io/gorm v1.22.5
)
