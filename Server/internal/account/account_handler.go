package account

import (
	"context"
	"database/sql"
	"dungeons_helper/util"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/mail"
	"os"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/api/idtoken"
)

type Handler struct {
	Service
}

func NewHandler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func isValidEmail(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}

func (h *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account CreateAccountReq
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}(r.Body)

	if err := isValidEmail(account.Email); err != nil {
		http.Error(w, fmt.Sprintf("Invalid email address: %s", err.Error()), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	err = h.Service.CreateAccount(ctx, &account)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var account LoginAccountReq
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}(r.Body)

	ctx := r.Context()
	res, err := h.Service.Login(ctx, &account)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	setCookieAndRespond(w, "jwt", res.accessToken, res)
}

func (h *Handler) Logout(w http.ResponseWriter, _ *http.Request) {
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		Domain:   "localhost",
		Secure:   false,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"message": "logout successful"}`))
}

func (h *Handler) LoginGoogle(w http.ResponseWriter, r *http.Request) {
	var googleToken Token
	err := json.NewDecoder(r.Body).Decode(&googleToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}(r.Body)

	idToken := googleToken.Token
	audience := os.Getenv("GOOGLE_OAUTH_CLIENT_ID")

	payload, err := idtoken.Validate(context.Background(), idToken, audience)
	if err != nil {
		http.Error(w, "Invalid ID token", http.StatusUnauthorized)
		return
	}
	email, found := payload.Claims["email"].(string)
	if !found {
		http.Error(w, "Email not found in ID token", http.StatusUnauthorized)
		return
	}
	imageURL, found := payload.Claims["picture"].(string)
	if !found {
		http.Error(w, "Picture URL not found in ID token", http.StatusUnauthorized)
		return
	}

	var googleData GoogleAcc
	googleData.Email = email

	ctx := r.Context()
	res, err := h.Service.GoogleAuth(ctx, &googleData)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			var account CreateAccountReq
			account.Email = googleData.Email
			response, err := http.Get(imageURL)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer response.Body.Close()

			imageData, err := io.ReadAll(response.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			base64Image := base64.StdEncoding.EncodeToString(imageData)

			account.Avatar = base64Image
			account.Nickname = strings.Split(account.Email, "@")[0]
			account.Password = util.GeneratePassword()
			err = h.Service.CreateAccount(ctx, &account)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			res, err := h.Service.GoogleAuth(ctx, &googleData)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			setCookieAndRespond(w, "jwt", res.accessToken, res)
		} else {
			handleError(w, err, http.StatusInternalServerError)
		}
		return
	}

	setCookieAndRespond(w, "jwt", res.accessToken, res)
}

func (h *Handler) RestorePassword(w http.ResponseWriter, r *http.Request) {
	var restoreReq RestoreReq
	if err := json.NewDecoder(r.Body).Decode(&restoreReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}(r.Body)

	if restoreReq.Email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	err := h.Service.RestorePassword(ctx, &restoreReq)
	if err != nil {
		handleError(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdateNickname(w http.ResponseWriter, r *http.Request) {
	var req UpdateNicknameReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}(r.Body)

	id, err := util.GetIdFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	req = UpdateNicknameReq{Id: id, Nickname: req.Nickname}
	ctx := r.Context()
	err = h.Service.UpdateNickname(ctx, &req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			http.Error(w, "Wrong password or email", http.StatusUnauthorized)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdatePassword(w http.ResponseWriter, r *http.Request) {
	var req UpdatePasswordReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}(r.Body)

	id, err := util.GetIdFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	req = UpdatePasswordReq{Id: id, OldPassword: req.OldPassword, NewPassword: req.NewPassword}
	ctx := r.Context()
	err = h.Service.UpdatePassword(ctx, &req)
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			http.Error(w, "Wrong old password", http.StatusUnauthorized)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdateAvatar(w http.ResponseWriter, r *http.Request) {
	var req UpdateAvatarReq

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}(r.Body)

	id, err := util.GetIdFromToken(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	req = UpdateAvatarReq{Id: id, Avatar: req.Avatar}
	ctx := r.Context()
	err = h.Service.UpdateAvatar(ctx, &req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			http.Error(w, "Wrong password or email", http.StatusUnauthorized)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}

func handleError(w http.ResponseWriter, err error, status int) {
	switch {
	case errors.Is(err, context.DeadlineExceeded):
		http.Error(w, "Wrong password or email", http.StatusRequestTimeout)
	case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) || errors.Is(err, sql.ErrNoRows):
		http.Error(w, "Wrong password or email", http.StatusUnauthorized)
	case errors.As(err, new(*mysql.MySQLError)):
		mysqlErr := err.(*mysql.MySQLError)
		if mysqlErr != nil && mysqlErr.Number == 1062 {
			http.Error(w, err.Error(), http.StatusConflict)
		} else {
			http.Error(w, err.Error(), status)
		}
	default:
		http.Error(w, err.Error(), status)
	}
}

func setCookieAndRespond(w http.ResponseWriter, cookieName, cookieValue string, res *LoginAccountRes) {
	cookie := http.Cookie{
		Name:     cookieName,
		Value:    cookieValue,
		Expires:  time.Now().Add(time.Hour * 24 * 7),
		Path:     "/",
		Domain:   "localhost",
		Secure:   false,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	jsonResponse, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(jsonResponse)
}
