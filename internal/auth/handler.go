package auth

import (
	"net/http"
	"temporary/config"
	"temporary/pkg/jwt"
	"temporary/pkg/res"
	"temporary/req"
)


type AuthHandler struct {
	Config config.AuthConfig
	SessionsStore SessionStore
	JWT *jwt.JWT
}

type AuthHandlerDeps struct {
	Config config.AuthConfig
	SessionsStore SessionStore
	JWT *jwt.JWT
}

func NewAuthHandler(router *http.ServeMux, deps *AuthHandlerDeps) *AuthHandler {
	handler := &AuthHandler{
		Config:       deps.Config,
		SessionsStore: deps.SessionsStore,
		JWT:          deps.JWT,
	}

	router.HandleFunc("POST /phone-verification", handler.GetPhoneCode())
	router.HandleFunc("POST /phone-verification/verify", handler.VerifyPhoneCode())

	return handler
}

func (h *AuthHandler) GetPhoneCode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, err := req.HandleBody[GetPhoneCodeRequest](&w, r)
		if err != nil {
			return
		}

		sessionId, err := h.SessionsStore.Generate(body.Phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = h.SessionsStore.Save(sessionId, body.Phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.WriteJSON(w, GetPhoneCodeResponse{
			SessionId: sessionId,
		}, http.StatusOK)
	}
}

func (h *AuthHandler) VerifyPhoneCode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, err := req.HandleBody[VerifyPhoneCodeRequest](&w, r)
		if err != nil {
			return
		}
		
		session, err := h.SessionsStore.Get(body.SessionId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if body.Code != session.Code {
			http.Error(w, ErrInvalidCode, http.StatusBadRequest)
			return
		}

		err = h.SessionsStore.Delete(body.SessionId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		token, err := h.JWT.Create(session.Phone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.WriteJSON(w, VerifyPhoneCodeResponse{
			Token: token,
		}, http.StatusOK)
	}
}