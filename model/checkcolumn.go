package model

import (
	"sigmamono/internal/core"
	"sigmamono/internal/term"
	"sigmamono/utils/helper"
	"strings"
)

func checkColumns(cols []string, variate string) (string, error) {
	fieldError := core.NewFieldError(term.Error_in_url)

	if variate == "*" {
		return strings.Join(cols, ","), nil
	}

	variates := strings.Split(variate, ",")
	for _, v := range variates {
		if ok, _ := helper.Includes(cols, v); !ok {
			fieldError.Add(term.V_is_not_valid, v, strings.Join(cols, ", "))
		}
	}
	if fieldError.HasError() {
		return "", fieldError
	}

	return variate, nil

}
