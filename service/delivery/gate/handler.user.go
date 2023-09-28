package gate

import (
	"fmt"
	"net/http"

	"github.com/DeniesKresna/bkn/models"
	"github.com/DeniesKresna/gobridge/serror"
)

func (c *Gate) UserCreate(w http.ResponseWriter, r *http.Request) {
	functionName := "[Gate.UserCreate]"
	ctx := r.Context()
	var errx serror.SError

	request := models.CreateUserRequest{}

	errx = c.DecodeRequestAndValidate(r, functionName, &request)
	if errx != nil {
		c.ResponseJSON(w, errx.GetStatusCode(), errx, errx.GetMessage())
		return
	}

	user, errx := c.UserUsecase.UserCreate(ctx, request)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While UserCreate", functionName))
		c.ResponseJSON(w, errx.GetStatusCode(), errx, errx.GetMessage())
		return
	}

	c.ResponseJSON(w, http.StatusCreated, user, "Berhasil tambah pengguna")
	return
}

func (c *Gate) UserRegularCreate(w http.ResponseWriter, r *http.Request) {
	functionName := "[Gate.UserRegularCreate]"
	ctx := r.Context()
	var errx serror.SError

	request := models.CreateUserRequest{}

	errx = c.DecodeRequestAndValidate(r, functionName, &request)
	if errx != nil {
		c.ResponseJSON(w, errx.GetStatusCode(), errx, errx.GetMessage())
		return
	}

	userCreateResponse, errx := c.UserUsecase.UserRegularCreate(ctx, request)

	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While UserRegularCreate Error", functionName))
		c.ResponseJSON(w, errx.GetStatusCode(), errx, errx.GetMessage())
		return
	}

	c.ResponseJSON(w, http.StatusCreated, userCreateResponse, "Berhasil tambah pengguna")
	return
}

func (c *Gate) UserLogin(w http.ResponseWriter, r *http.Request) {
	functionName := "[Gate.UserLogin]"
	ctx := r.Context()
	var errx serror.SError

	request := models.AuthRequest{}

	errx = c.DecodeRequestAndValidate(r, functionName, &request)
	if errx != nil {
		c.ResponseJSON(w, errx.GetStatusCode(), errx, errx.GetMessage())
		return
	}

	authResp, errx := c.AuthUsecase.AuthLogin(ctx, request.Email, request.Password)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While UserLogin Error", functionName))
		c.ResponseJSON(w, errx.GetStatusCode(), errx, errx.GetMessage())
		return
	}

	c.ResponseJSON(w, http.StatusOK, authResp, "Berhasil Login")
	return
}

func (c *Gate) UserIndex(w http.ResponseWriter, r *http.Request) {
	functionName := "[Gate.UserIndex]"
	ctx := r.Context()
	var errx serror.SError

	paginationData := c.GetRequestPaginationData(r)

	request := models.UserSearch{}
	errx = c.DecodeRequestAndValidate(r, functionName, &request)
	if errx != nil {
		c.ResponseJSON(w, errx.GetStatusCode(), errx, errx.GetMessage())
		return
	}

	users, errx := c.UserUsecase.UserIndexWithPagination(ctx, request, paginationData)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While Get User Pagination Error", functionName))
		c.ResponseJSON(w, errx.GetStatusCode(), errx, errx.GetMessage())
		return
	}

	c.ResponseJSON(w, http.StatusOK, users, "Berhasil")
	return
}

func (c *Gate) UserGetDetailById(w http.ResponseWriter, r *http.Request) {
	functionName := "[Gate.UserGetDetailById]"
	ctx := r.Context()
	var errx serror.SError

	userID := c.GetInt64Var(r, "id")

	userAndProfile, errx := c.UserUsecase.UserProfileGetByID(ctx, userID)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While UserProfileGetByID Error", functionName))
		c.ResponseJSON(w, errx.GetStatusCode(), errx, errx.GetMessage())
		return
	}

	c.ResponseJSON(w, http.StatusOK, userAndProfile, "Berhasil")
	return
}

func (c *Gate) UserUpdatePassword(w http.ResponseWriter, r *http.Request) {
	functionName := "[Gate.UserUpdatePassword]"
	ctx := r.Context()
	var errx serror.SError

	userID := c.GetInt64Var(r, "id")

	request := models.CreatePasswordRequest{}

	errx = c.DecodeRequestAndValidate(r, functionName, &request)
	if errx != nil {
		c.ResponseJSON(w, errx.GetStatusCode(), errx, errx.GetMessage())
		return
	}

	user, errx := c.UserUsecase.UserUpdatePassword(ctx, userID, request)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While UserUpdatePassword Error", functionName))
		c.ResponseJSON(w, errx.GetStatusCode(), errx, errx.GetMessage())
		return
	}

	c.ResponseJSON(w, http.StatusAccepted, user, "Berhasil Sunting Password")
	return
}

func (c *Gate) UserSelfUpdatePassword(w http.ResponseWriter, r *http.Request) {
	functionName := "[Gate.UserSelfUpdatePassword]"
	ctx := r.Context()
	var errx serror.SError

	request := models.CreatePasswordRequest{}

	errx = c.DecodeRequestAndValidate(r, functionName, &request)
	if errx != nil {
		c.ResponseJSON(w, errx.GetStatusCode(), errx, errx.GetMessage())
		return
	}

	user, errx := c.UserUsecase.UserSelfUpdatePassword(ctx, request)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While UserSelfUpdatePassword Wrong", functionName))
		c.ResponseJSON(w, errx.GetStatusCode(), errx, errx.GetMessage())
		return
	}

	c.ResponseJSON(w, http.StatusAccepted, user, "Berhasil Sunting Password")
	return
}

func (c *Gate) UserUpdateImg(w http.ResponseWriter, r *http.Request) {
	functionName := "[Gate.UserUpdateImg]"
	ctx := r.Context()
	var errx serror.SError

	userID := c.GetInt64Var(r, "id")

	request := models.CreateImgRequest{}

	errx = c.DecodeRequestAndValidate(r, functionName, &request)
	if errx != nil {
		c.ResponseJSON(w, errx.GetStatusCode(), errx, errx.GetMessage())
		return
	}

	user, errx := c.UserUsecase.UserUpdateImageURL(ctx, userID, &request)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While User Upload image error", functionName))
		c.ResponseJSON(w, errx.GetStatusCode(), errx, errx.GetMessage())
		return
	}

	c.ResponseJSON(w, http.StatusAccepted, user, "Berhasil Sunting Gambar")
	return
}

func (c *Gate) UserUpdate(w http.ResponseWriter, r *http.Request) {
	functionName := "[Gate.UserUpdate]"
	ctx := r.Context()
	var errx serror.SError

	userID := c.GetInt64Var(r, "id")

	request := models.UpdateUserRequest{}

	errx = c.DecodeRequestAndValidate(r, functionName, &request)
	if errx != nil {
		c.ResponseJSON(w, errx.GetStatusCode(), errx, errx.GetMessage())
		return
	}

	user, errx := c.UserUsecase.UserUpdate(ctx, userID, request)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While UserUpdate", functionName))
		c.ResponseJSON(w, errx.GetStatusCode(), errx, errx.GetMessage())
		return
	}

	c.ResponseJSON(w, http.StatusAccepted, user, "Berhasil Sunting pengguna")
	return
}

func (c *Gate) UserGetSession(w http.ResponseWriter, r *http.Request) {
	functionName := "[Gate.UserGetSession]"
	ctx := r.Context()
	var errx serror.SError

	resp, errx := c.AuthUsecase.AuthGetSession(ctx)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While get user Session", functionName))
		c.ResponseJSON(w, http.StatusUnauthorized, errx, errx.GetMessage())
		return
	}

	c.ResponseJSON(w, http.StatusOK, resp, "Berhasil")
	return
}

func (c *Gate) UserRegister(w http.ResponseWriter, r *http.Request) {
	functionName := "[Gate.UserRegister]"
	ctx := r.Context()
	var errx serror.SError

	request := models.UserRegister{}

	errx = c.DecodeRequestAndValidate(r, functionName, &request)
	if errx != nil {
		c.ResponseJSON(w, errx.GetStatusCode(), errx, errx.GetMessage())
		return
	}

	errx = c.UserUsecase.UserRegister(ctx, &request)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While Register user", functionName))
		c.ResponseJSON(w, errx.GetStatusCode(), errx, errx.GetMessage())
		return
	}

	c.ResponseJSON(w, http.StatusCreated, nil, "Berhasil registrasi, silakan cek email anda untuk proses selanjutnya")
	return
}

func (c *Gate) UserVerifyToken(w http.ResponseWriter, r *http.Request) {
	functionName := "[Gate.UserVerifyToken]"
	ctx := r.Context()
	var errx serror.SError

	var request struct {
		Token string `json:"token"`
	}

	errx = c.DecodeRequestAndValidate(r, functionName, &request)
	if errx != nil {
		c.ResponseJSON(w, errx.GetStatusCode(), errx, errx.GetMessage())
		return
	}

	errx = c.UserUsecase.UserVerifyToken(ctx, request.Token)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While UserVerifyToken", functionName))
		c.ResponseJSON(w, errx.GetStatusCode(), errx, errx.GetMessage())
		return
	}

	c.ResponseJSON(w, http.StatusAccepted, nil, "Berhasil meregistrasi pengguna, kamu bisa login sekarang")
	return
}

func (c *Gate) UserTransactionHistoryByID(w http.ResponseWriter, r *http.Request) {
	functionName := "[Gate.UserTransactionHistory]"
	ctx := r.Context()
	var errx serror.SError

	paginationData := c.GetRequestPaginationData(r)
	search := c.GetQuery(r, "search")

	userID := c.GetInt64Var(r, "id", 0)

	users, errx := c.UserUsecase.UserTransactionHistoryByIDWithPagination(ctx, search, userID, paginationData)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While Get UserTransactionHistoryWithPagination Error", functionName))
		c.ResponseJSON(w, errx.GetStatusCode(), errx, errx.GetMessage())
		return
	}

	c.ResponseJSON(w, http.StatusOK, users, "Berhasil")
	return
}

func (c *Gate) UserProfile(w http.ResponseWriter, r *http.Request) {
	functionName := "[Gate.UserProfile]"
	ctx := r.Context()
	var errx serror.SError

	resp, errx := c.UserUsecase.UserProfile(ctx)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While get user profile", functionName))
		return
	}

	c.ResponseJSON(w, http.StatusOK, resp, "Berhasil")
	return
}

func (c *Gate) UserSelfUpdate(w http.ResponseWriter, r *http.Request) {
	functionName := "[Gate.UserSelfUpdate]"
	ctx := r.Context()
	var errx serror.SError

	request := models.UpdateUserRequest{}

	errx = c.DecodeRequestAndValidate(r, functionName, &request)
	if errx != nil {
		c.ResponseJSON(w, errx.GetStatusCode(), errx, errx.GetMessage())
		return
	}

	user, errx := c.UserUsecase.UserSelfUpdate(ctx, request)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While UserSelfUpdate", functionName))
		c.ResponseJSON(w, errx.GetStatusCode(), errx, errx.GetMessage())
		return
	}

	c.ResponseJSON(w, http.StatusCreated, user, "Berhasil Sunting pengguna")
	return
}

func (c *Gate) UserDelete(w http.ResponseWriter, r *http.Request) {
	functionName := "[Gate.UserDelete]"
	ctx := r.Context()
	var errx serror.SError

	requestID := c.GetInt64Var(r, "id", 0)

	errx = c.UserUsecase.UserDelete(ctx, requestID)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While UserDelete Error", functionName))
		c.ResponseJSON(w, errx.GetStatusCode(), errx, errx.GetMessage())
		return
	}

	c.ResponseJSON(w, http.StatusAccepted, nil, "Berhasil")
	return
}

func (c *Gate) UserSelfAddOrUpdateProfile(w http.ResponseWriter, r *http.Request) {
	functionName := "[Gate.UserSelfAddOrUpdateProfile]"
	ctx := r.Context()
	var errx serror.SError

	request := models.ProfileRequest{}

	errx = c.DecodeRequestAndValidate(r, functionName, &request)
	if errx != nil {
		c.ResponseJSON(w, errx.GetStatusCode(), errx, errx.GetMessage())
		return
	}

	errx = c.UserUsecase.UserSelfAddOrUpdateProfile(ctx, request)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While UserSelfAddOrUpdateProfile", functionName))
		c.ResponseJSON(w, errx.GetStatusCode(), errx, errx.GetMessage())
		return
	}

	c.ResponseJSON(w, http.StatusNoContent, nil, "Berhasil Sunting pengguna")
	return
}

func (c *Gate) UserAddOrUpdateProfileByID(w http.ResponseWriter, r *http.Request) {
	functionName := "[Gate.UserAddOrUpdateProfileByID]"
	ctx := r.Context()
	var errx serror.SError

	userID := c.GetInt64Var(r, "id", 0)

	request := models.ProfileRequest{}

	errx = c.DecodeRequestAndValidate(r, functionName, &request)
	if errx != nil {
		c.ResponseJSON(w, errx.GetStatusCode(), errx, errx.GetMessage())
		return
	}

	errx = c.UserUsecase.UserAddOrUpdateProfileByID(ctx, request, userID)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While UserAddOrUpdateProfileByID", functionName))
		c.ResponseJSON(w, errx.GetStatusCode(), errx, errx.GetMessage())
		return
	}

	c.ResponseJSON(w, http.StatusNoContent, nil, "Berhasil Sunting pengguna")
	return
}
