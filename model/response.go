package model

import "solidgate-test/util"

type Response struct {
	Validness bool        `json:"valid"`
	Error     *util.Error `json:"error,omitempty"`
}
