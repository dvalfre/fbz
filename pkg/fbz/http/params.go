package http

import (
	"net/url"

	"github.com/ess/fbz/pkg/fbz"
)

func paramsToValues(params fbz.Params) url.Values {
	return url.Values(params)
}
