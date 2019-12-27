package envd

import "syscall"

const (
	// ENForBESdPort is exported.
	ENForBESdPort = "KN_BE_SD_PORT"
	// ENForBESd is exported.
	ENForBESd = "KN_BE_SD"
	// ENForBEApi is exported.
	ENForBEApi = "KN_BE_API"
	// ENForFEWeb is exported.
	ENForFEWeb = "KN_FE_WEB"
)

// SD struct used to keep information about internal services.
type SD struct {
	BESdPort string `json:"be_sd_port"`
	BESd     string `json:"be_sd"`
	BEApi    string `json:"be_api"`
	FEWeb    string `json:"fe_web"`
}

// NewSD harvests specific env variables and returns a new instance of SD.
func NewSD() (*SD, error) {
	sdPort, _ := syscall.Getenv(ENForBESdPort)
	sdURL, _ := syscall.Getenv(ENForBESd)
	apiURL, _ := syscall.Getenv(ENForBEApi)
	webURL, _ := syscall.Getenv(ENForFEWeb)

	sd := SD{
		BESdPort: sdPort,
		BESd:     sdURL,
		BEApi:    apiURL,
		FEWeb:    webURL,
	}

	return &sd, nil
}
