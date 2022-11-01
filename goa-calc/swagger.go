package calcapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
)

type SwaggerHTTPServer struct {
}

func New(
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	externalServerBaseUrl string,
	swaggerDocFilePath string,
	uiEndpoint string,
) (*SwaggerHTTPServer, error) {
	// externalURL, err := url.Parse(externalServerBaseUrl)
	// if err != nil {
	// 	return nil, fmt.Errorf("url.Parse() error: %s", err.Error())
	// }
	// externalURL.Scheme
	if !fileExists(swaggerDocFilePath) {
		return nil, fmt.Errorf("swagger file %s does not exist", swaggerDocFilePath)
	}
	swaggerFile, err := os.Open(swaggerDocFilePath)
	if err != nil {
		return nil, fmt.Errorf("os.Open() error: %s", err.Error())
	}
	var swaggerFileContentMap map[string]interface{}
	err = json.NewDecoder(swaggerFile).Decode(&swaggerFileContentMap)
	if err != nil {
		return nil, fmt.Errorf("json.Decode() error: %s", err.Error())
	}
	swaggerFileContentMap["servers"] = []map[string]string{
		{
			"url": externalServerBaseUrl,
		},
	}
	return nil, nil
}

// fileExists returns a bool indicating whether a file exists or not. Also
// in the case the input filename is a directory, even if it "exists", the function
// will also return false as it is not a file
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if errors.Is(err, fs.ErrNotExist) {
		return false
	}
	return !info.IsDir()
}
