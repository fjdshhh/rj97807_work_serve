package middleware

import (
	"net/http"
	"rj97807_work_serve/funcs"
	"strconv"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("用于认证失败"))
			return
		}
		uc, err := funcs.AnalyzeToken(auth)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}
		r.Header.Set("UserRole", strconv.Itoa(uc.Role))
		r.Header.Set("Uid", uc.Uid)
		r.Header.Set("UserName", uc.Name)
		next(w, r)
	}
}
