package token

type Sign string

const (
	SignToken       = "x"
	SignXToken      = "X"
	SignAccessToken = "X-Access-Token"
	SignLoginToken  = "X-Auth-Token"
)
