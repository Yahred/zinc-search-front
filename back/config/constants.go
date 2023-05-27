package config

import "encoding/base64"

const RootDirectoryData = "data"
const ZincSearchURL = "http://localhost:4080/api"
const ZincSearchUsername = "admin"
const ZincSearchPassword = "admin"

func GetZincSearchCredentials() string {
	return base64.StdEncoding.EncodeToString([]byte(ZincSearchUsername + ":" + ZincSearchPassword))
}
