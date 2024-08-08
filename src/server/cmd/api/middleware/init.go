package middleware

import "github.com/cloudwego/hertz/pkg/app/server"

func Init(r *server.Hertz) {
	r.Use(configRecovery())
}
