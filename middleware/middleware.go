package middleware

import (
	"net/http"
)

type contextKey string

const userContextKey = contextKey("authenticatedUser")

type Middleware func(http.HandlerFunc) http.HandlerFunc
