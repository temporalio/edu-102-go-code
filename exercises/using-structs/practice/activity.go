package translation

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// TODO Replace the last two input parameters with the struct you defined as input
// TODO Replace the first output type (string) with the name of the struct you defined as output
func TranslateTerm(ctx context.Context, inputTerm string, languageCode string) (string, error) {
	// TODO Change the parameters used in these two calls to QueryEscape with
	//      the appopriate fields from your struct
	lang := url.QueryEscape(languageCode)
	term := url.QueryEscape(inputTerm)
	url := fmt.Sprintf("http://localhost:9998/translate?lang=%s&term=%s", lang, term)

	resp, err := http.Get(url)
	if err != nil {
		// TODO Return an empty output struct instead of an empty string
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// TODO Return an empty output struct instead of an empty string
		return "", err
	}

	// This string will contain either the translated term, if the service could
	// perform the translation, or the error message, if it was unsuccessful
	content := string(body)

	status := resp.StatusCode
	if status >= 400 {
		// This means that we succcessfully called the service, but it could not
		// perform the translation for some reason
		// TODO Return an empty output struct instead of an empty string
		return "", fmt.Errorf("HTTP Error %d: %s", status, content)
	}

	// TODO Replace 'content' below with the struct your using as output,
	//      populated with the translation
	return content, nil
}
