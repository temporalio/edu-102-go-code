package translation

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"go.temporal.io/sdk/activity"
)

func TranslateTerm(ctx context.Context, input TranslationActivityInput) (TranslationActivityOutput, error) {
	logger := activity.GetLogger(ctx)

	logger.Info("Translating term", "LanguageCode", input.LanguageCode, "Term", input.Term)

	lang := url.QueryEscape(input.LanguageCode)
	term := url.QueryEscape(input.Term)
	url := fmt.Sprintf("http://localhost:9998/translate?lang=%s&term=%s", lang, term)

	resp, err := http.Get(url)
	if err != nil {
		logger.Error("Translation request failed", "Error", err)
		return TranslationActivityOutput{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Failed to read response content", "Error", err)
		return TranslationActivityOutput{}, err
	}

	// This string will contain either the translated term, if the service could
	// perform the translation, or the error message, if it was unsuccessful
	content := string(body)

	status := resp.StatusCode
	if status >= 400 {
		// This means that we succcessfully called the service, but it could not
		// perform the translation for some reason
		logger.Error("Translation failed", "Status", status, "Message", content)
		message := fmt.Sprintf("HTTP Error %d: %s", status, content)
		return TranslationActivityOutput{}, errors.New(message)
	}

	logger.Debug("Translation successful", "Translation", content)
	output := TranslationActivityOutput{
		Translation: content,
	}

	return output, nil
}
