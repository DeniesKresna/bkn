package usecase

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/DeniesKresna/bkn/models"
	"github.com/DeniesKresna/gobridge/serror"
	"github.com/DeniesKresna/gohelper/utstring"
	"github.com/DeniesKresna/gohelper/utstruct"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func (h *AuthUsecase) AuthGetFromContext(ctx context.Context) (res models.UserRole, errx serror.SError) {
	functionName := "[AuthUsecase.AuthGetFromContext]"

	session := ctx.Value("session")
	sessionType, ok := session.(models.Session)
	if !ok {
		errx = serror.NewWithErrorCommentMessage(nil, http.StatusUnauthorized, fmt.Sprintf("%s Session Not Found", functionName), "Sesi tidak ditemukan")
		return
	}
	userID := sessionType.UserID
	if userID <= 0 {
		errx = serror.NewWithErrorCommentMessage(nil, http.StatusUnauthorized, fmt.Sprintf("%s Session Not Found", functionName), "Sesi tidak ditemukan")
		return
	}

	userRes, errx := h.userRepo.UserGetByID(ctx, userID)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While User UserGetByID", functionName), "Gagal dapat sesi user")
		return
	}

	r, errx := h.userRepo.RoleGetByID(ctx, userRes.RoleId)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While RoleGetByID", functionName))
		return
	}

	utstruct.InjectStructValue(userRes, &res)
	res.RoleName = r.Name

	return
}

func (h *AuthUsecase) AuthLogin(ctx context.Context, email string, password string) (authResp models.AuthResponse, errx serror.SError) {
	functionName := "[AuthUsecase.AuthLogin]"

	user, errx := h.userRepo.UserGetByEmail(ctx, email)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserGetByEmail", functionName), "Pengguna tidak ditemukan")
		return
	}

	if user.Active == 0 {
		errx = serror.NewWithErrorCommentMessage(nil, http.StatusNotFound, "[Usecase.UserLogin] While Check activation", "Pengguna tidak aktif")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		errx = serror.NewWithErrorCommentMessage(err, http.StatusBadRequest, "[Usecase.UserLogin] While CompareHashAndPassword", "Password salah")
		return
	}

	var (
		tokenString string
		expires     time.Time
	)
	// token generation
	{
		expires = time.Now().Add(time.Hour * 3)

		// Create the JWT token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			ExpiresAt: expires.Unix(),
			Issuer:    "my-app",
			Subject:   fmt.Sprintf("%d", user.ID),
		})

		// Sign the token with a secret key
		tokenString, err = token.SignedString([]byte(utstring.GetEnv(models.AppApiSecret)))
		if err != nil {
			errx = serror.NewWithErrorCommentMessage(err, http.StatusInternalServerError, fmt.Sprintf("%s While SignedString token", functionName), "Tidak dapat membuat sesi")
			return
		}
	}

	r, errx := h.userRepo.RoleGetByID(ctx, user.RoleId)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While RoleGetByID", functionName))
		return
	}

	authResp = models.AuthResponse{
		User:      user,
		Token:     tokenString,
		Role:      r,
		ExpiredAt: expires.Format(time.RFC3339),
	}

	return
}

func (h *AuthUsecase) AuthGetSession(ctx context.Context) (a models.AuthResponse, errx serror.SError) {
	functionName := "[UserUsecase.UserGetSession]"

	userRole, errx := h.AuthGetFromContext(ctx)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While AuthGetFromContext", functionName), "Sesi tidak ditemukan")
		return
	}

	u, errx := h.userRepo.UserGetByID(ctx, userRole.ID)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserGetByID", functionName), "Pengguna tidak ditemukan")
		return
	}

	r, errx := h.userRepo.RoleGetByID(ctx, u.RoleId)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While RoleGetByID", functionName))
		return
	}

	a.User = u
	a.Role = r

	return
}
