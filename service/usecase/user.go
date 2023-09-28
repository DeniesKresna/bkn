package usecase

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/DeniesKresna/bkn/models"
	"github.com/DeniesKresna/bkn/service/helpers"
	"github.com/DeniesKresna/gobridge/serror"
	"github.com/DeniesKresna/gohelper/utlog"
	"github.com/DeniesKresna/gohelper/utstorage"
	"github.com/DeniesKresna/gohelper/utstring"
	"github.com/DeniesKresna/gohelper/utstruct"
	"golang.org/x/crypto/bcrypt"
)

func (h *UserUsecase) UserCreate(ctx context.Context, req models.CreateUserRequest) (u models.User, errx serror.SError) {
	functionName := "[UserUsecase.UserCreate]"

	//handle tx
	err := h.db.StartTx()
	if err != nil {
		errx = serror.NewWithErrorCommentMessage(err, http.StatusInternalServerError, fmt.Sprintf("%s While StartTx", functionName), "Kesalahan Server")
		return
	}
	defer func() {
		if errx != nil {
			err = errx.GetError()
		}
		h.db.SubmitTx(err)
	}()

	if u, errx = h.userRepo.UserGetByEmail(ctx, req.Email); errx == nil {
		return u, serror.NewWithCommentMessage(http.StatusConflict, fmt.Sprintf("%s While UserGetByEmail", functionName), "Pengguna tersebut sudah ada")
	}

	if req.Password == "" {
		req.Password = req.FirstName + "." + req.LastName + time.Now().Format("02.01.06")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		errx = serror.NewWithErrorCommentMessage(err, http.StatusInternalServerError, fmt.Sprintf("%s While GenerateFromPassword", functionName), "Gagal Generate Password")
		return
	}

	req.Password = string(hashedPassword)
	req.Active = 1

	var tempImageCode string
	if req.ImageUrl != nil {
		tempImageCode = *req.ImageUrl
		req.ImageUrl = nil
	}

	operator, errx := h.authCase.AuthGetFromContext(ctx)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While Get Session", functionName))
		return
	}

	uID, errx := h.userRepo.UserCreate(ctx, req, operator.Email)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserCreate", functionName), "Gagal Tambah User")
		return
	}

	if tempImageCode != "" {
		imageReq := models.CreateImgRequest{
			Image: tempImageCode,
		}

		_, errx = h.UserUpdateImageURL(ctx, uID, &imageReq)
		if errx != nil {
			errx.AddCommentMessage(fmt.Sprintf("%s While UserUpdateImageURL", functionName), "Gagal Update Foto Pengguna")
			return
		}
	}

	u, errx = h.UserGetByID(ctx, uID)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserGetByID", functionName), "Gagal Ambil data User Baru")
		return
	}

	return
}

func (h *UserUsecase) UserUpdate(ctx context.Context, id int64, req models.UpdateUserRequest) (u models.User, errx serror.SError) {
	functionName := "[UserUsecase.UserUpdate]"

	//handle tx
	err := h.db.StartTx()
	if err != nil {
		errx = serror.NewWithErrorCommentMessage(err, http.StatusInternalServerError, fmt.Sprintf("%s While StartTx", functionName), "Kesalahan Server")
		return
	}
	defer func() {
		if errx != nil {
			err = errx.GetError()
		}
		h.db.SubmitTx(err)
	}()

	u, errx = h.UserGetByID(ctx, id)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While UserGetByID. id: %d", functionName, id))
		return
	}

	operator, errx := h.authCase.AuthGetFromContext(ctx)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While Get Session", functionName))
		return
	}

	if operator.RoleName != models.RoleNameAdmin {
		if operator.ID != id {
			errx = serror.NewWithCommentMessage(http.StatusForbidden, fmt.Sprintf("%s While Check user id", functionName), "Operasi ini tidak diijinkan")
			return
		}
	}

	var tempImageCode string
	if req.ImageUrl != nil {
		tempImageCode = *req.ImageUrl
		req.ImageUrl = nil
	}

	errx = h.userRepo.UserUpdate(ctx, u.ID, req, operator.Email)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserUpdate", functionName), "Gagal Sunting Pengguna")
		return
	}

	if tempImageCode != "" {
		if tempImageCode != *u.ImageUrl {
			imageReq := models.CreateImgRequest{
				Image: tempImageCode,
			}

			_, errx = h.UserUpdateImageURL(ctx, u.ID, &imageReq)
			if errx != nil {
				errx.AddCommentMessage(fmt.Sprintf("%s While UserUpdateImageURL", functionName), "Gagal Update Foto Pengguna")
				return
			}
		}
	}

	u, errx = h.UserGetByID(ctx, u.ID)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserGetByID", functionName), "Gagal Ambil data User Baru")
		return
	}

	return
}

func (h *UserUsecase) UserGetByEmail(ctx context.Context, email string) (u models.User, errx serror.SError) {
	functionName := "[UserUsecase.UserGetByEmail]"

	u, errx = h.userRepo.UserGetByEmail(ctx, email)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserGetByEmail", functionName), "Pengguna tidak ditemukan")
		return
	}

	return
}

func (h *UserUsecase) UserRoleGetByID(ctx context.Context, id int64) (u models.UserRole, errx serror.SError) {
	functionName := "[UserUsecase.UserRoleGetByID]"

	u, errx = h.userRepo.UserRoleGetByID(ctx, id)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserRoleGetByID", functionName), "Gagal Lihat Pengguna")
		return
	}

	return
}

func (h *UserUsecase) UserGetByID(ctx context.Context, id int64) (u models.User, errx serror.SError) {
	functionName := "[UserUsecase.UserGetByID]"

	u, errx = h.userRepo.UserGetByID(ctx, id)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserGetByID", functionName), "Gagal Lihat User by ID")
		return
	}
	return
}

func (h *UserUsecase) UserSelfUpdatePassword(ctx context.Context, req models.CreatePasswordRequest) (u models.User, errx serror.SError) {
	functionName := "[UserUsecase.UserSelfUpdatePassword]"

	//handle tx
	err := h.db.StartTx()
	if err != nil {
		errx = serror.NewWithErrorCommentMessage(err, http.StatusInternalServerError, fmt.Sprintf("%s While StartTx", functionName), "Kesalahan Server")
		return
	}
	defer func() {
		if errx != nil {
			err = errx.GetError()
		}
		h.db.SubmitTx(err)
	}()

	operator, errx := h.authCase.AuthGetFromContext(ctx)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While Get Session", functionName))
		return
	}

	u, errx = h.UserUpdatePassword(ctx, operator.ID, req)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserUpdatePassword", functionName), "Gagal update password user")
		return
	}

	return
}

func (h *UserUsecase) UserUpdatePassword(ctx context.Context, id int64, req models.CreatePasswordRequest) (u models.User, errx serror.SError) {
	functionName := "[UserUsecase.UserUpdatePassword]"

	//handle tx
	err := h.db.StartTx()
	if err != nil {
		errx = serror.NewWithErrorCommentMessage(err, http.StatusInternalServerError, fmt.Sprintf("%s While StartTx", functionName), "Kesalahan Server")
		return
	}
	defer func() {
		if errx != nil {
			err = errx.GetError()
		}
		h.db.SubmitTx(err)
	}()

	if req.NewPassword2 != req.NewPassword {
		errx = serror.NewWithCommentMessage(http.StatusBadRequest, fmt.Sprintf("%s While Compare Password 1 and 2", functionName), "Password baru tidak sama")
		return
	}

	u, errx = h.userRepo.UserGetByID(ctx, id)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserGetByID", functionName), "Pengguna tidak terdaftar")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.OldPassword))
	if err != nil {
		errx = serror.NewWithErrorCommentMessage(err, http.StatusMethodNotAllowed, fmt.Sprintf("%s While CompareHashAndPassword", functionName), "Password Lama salah")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		errx = serror.NewWithErrorCommentMessage(err, http.StatusInternalServerError, fmt.Sprintf("%s While GenerateFromPassword", functionName), "Tidak bisa generate password baru")
		return
	}

	operator, errx := h.authCase.AuthGetFromContext(ctx)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While Get Session", functionName))
		return
	}

	errx = h.userRepo.UserUpdatePassword(ctx, u, string(hashedPassword), operator.Email)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserUpdatePassword", functionName), "Tidak bisa update password baru")
		return u, errx
	}
	return u, nil
}

func (h *UserUsecase) UserUpdateRole(ctx context.Context, id int64, roleName string) (u models.User, errx serror.SError) {
	functionName := "[UserUsecase.UserUpdateRole]"

	u, errx = h.UserGetByID(ctx, id)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserGetByID", functionName), "Pengguna tidak terdaftar")
		return
	}

	role, errx := h.RoleGetByName(ctx, roleName)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While RoleGetByName", functionName))
		return
	}

	operator, errx := h.authCase.AuthGetFromContext(ctx)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While Get Session", functionName))
		return
	}

	errx = h.userRepo.UserUpdateRole(ctx, u, roleName, operator.Email)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While UserUpdateRole", functionName))
		return
	}

	u.RoleId = role.ID

	return
}

func (h *UserUsecase) UserUpdateImageURL(ctx context.Context, id int64, req *models.CreateImgRequest) (u models.User, errx serror.SError) {
	functionName := "[UserUsecase.UserUpdateImageURL]"

	//handle tx
	err := h.db.StartTx()
	if err != nil {
		errx = serror.NewWithErrorCommentMessage(err, http.StatusInternalServerError, fmt.Sprintf("%s While StartTx", functionName), "Kesalahan Server")
		return
	}
	defer func() {
		if errx != nil {
			err = errx.GetError()
		}
		h.db.SubmitTx(err)
	}()

	if u, errx = h.UserGetByID(ctx, id); errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserGetByID", functionName), "Pengguna tidak terdaftar")
		return
	}

	imageUrl, errx := h.UploadUserImage(ctx, u.ID, req.Image)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While uploadUserImage", functionName), "Gagal upload foto pengguna")
		return
	}

	operator, errx := h.authCase.AuthGetFromContext(ctx)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While Get Session", functionName))
		return
	}

	errx = h.userRepo.UserUpdateImageURL(ctx, u, imageUrl, operator.Email)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserUpdateImgUrl", functionName), "Tidak bisa menyimpan gambar ke database")
		return
	}

	return
}

func (h *UserUsecase) UserRegularCreate(ctx context.Context, req models.CreateUserRequest) (u models.User, errx serror.SError) {
	functionName := "[UserUsecase.UserRegularCreate]"

	//handle tx
	err := h.db.StartTx()
	if err != nil {
		errx = serror.NewWithErrorCommentMessage(err, http.StatusInternalServerError, fmt.Sprintf("%s While StartTx", functionName), "Kesalahan Server")
		return
	}
	defer h.db.SubmitTx(errx.GetError())

	u, errx = h.userRepo.UserGetByEmail(ctx, req.Email)
	if errx == nil {
		errx = serror.NewWithCommentMessage(http.StatusConflict, fmt.Sprintf("%s While UserGetByEmail", functionName), "Pengguna tersebut sudah ada")
		return
	}

	randomPassword := utstring.RandomString(12)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(randomPassword), bcrypt.DefaultCost)
	if err != nil {
		errx = serror.NewWithErrorCommentMessage(err, http.StatusInternalServerError, fmt.Sprintf("%s While GenerateFromPassword", functionName), "Tidak bisa generate password baru")
		return
	}

	req.Password = string(hashedPassword)

	// Hard Coded Role id, with 2 means "Regular User"
	req.RoleId = 2
	// activate user
	req.Active = 1

	operator, errx := h.authCase.AuthGetFromContext(ctx)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While Get Session", functionName))
		return
	}

	uID, errx := h.userRepo.UserCreate(ctx, req, operator.Email)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserCreate", functionName), "Gagal Tambah User")
		return
	}

	u, errx = h.UserGetByID(ctx, uID)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserGetByID", functionName), "Gagal Ambil Data User")
		return
	}

	return
}

func (h *UserUsecase) UserIndexWithPagination(ctx context.Context, filter models.UserSearch, paginationData models.PaginationData) (resp models.PaginationResponse, errx serror.SError) {
	functionName := "[UserUsecase.UserIndexWithPagination]"

	if resp, errx = h.userRepo.UserIndexWithPagination(ctx, filter, paginationData); errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserIndexWithPagination", functionName), "Gagal Ambil data pengguna")
		return resp, errx
	}
	return resp, nil
}

func (h *UserUsecase) UserDelete(ctx context.Context, id int64) (errx serror.SError) {
	functionName := "[UserUsecase.UserDelete]"

	_, errx = h.UserGetByID(ctx, id)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserGetByID", functionName), "Pengguna tidak terdaftar")
		return
	}

	operator, errx := h.authCase.AuthGetFromContext(ctx)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While Get Session", functionName))
		return
	}

	errx = h.userRepo.UserDelete(ctx, id, operator.Email)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserDelete", functionName), "Gagal hapus pengguna")
		return
	}
	return
}

func (h *UserUsecase) UserRegister(ctx context.Context, req *models.UserRegister) (errx serror.SError) {
	functionName := "[UserUsecase.UserRegister]"

	//handle tx
	err := h.db.StartTx()
	if err != nil {
		errx = serror.NewWithErrorCommentMessage(err, http.StatusInternalServerError, fmt.Sprintf("%s While StartTx", functionName), "Kesalahan Server")
		return
	}
	defer func() {
		if errx != nil {
			err = errx.GetError()
		}
		h.db.SubmitTx(err)
	}()

	// Check if user with that email exist
	u, errx := h.userRepo.UserGetByEmailIgnoreActive(ctx, req.Email)
	if errx == nil {
		_, errx = h.UserGetByID(ctx, u.ID)
		if errx == nil {
			errx = serror.NewWithErrorCommentMessage(nil, http.StatusConflict, fmt.Sprintf("%s While UserGetAuthByID", functionName), "Pengguna tersebut sudah aktif")
			return
		}
		vu, errx := h.userRepo.VerifiedGetVerifiedByUserID(ctx, u.ID)
		if errx == nil {
			now := time.Now()
			then := now.AddDate(0, 0, -1)
			if vu.CreatedAt.Before(then) {
				if vu, errx = h.userRepo.VerifiedDeleteVerified(ctx, vu); errx != nil {
					utlog.Errorf("%s While VerifiedDeleteVerified", functionName)
				}
			} else {
				errx = serror.NewWithCommentMessage(http.StatusConflict, fmt.Sprintf("%s While VerifiedGetVerifiedByUserID", functionName), "Kode aktivasi sudah dikirim sebelumnya, silakan cek email anda")
				return errx
			}
		}

		// Create sha256 encrypt code from userid_currenttimestamp_jobhunregis
		idstr := strconv.FormatUint(uint64(u.ID), 10)
		now := time.Now().Format("02.01.06")
		// code := h.utils.Encrypt256(idstr, now, "jobhunregis")  //TODO: need to be fixed.
		code := helpers.EncryptMD5(idstr, now, "jobhunregis")

		// Create new verified user
		verifiedUser := models.VerifiedUser{}
		verifiedUser.UserId = u.ID
		verifiedUser.Code = code

		if _, errx := h.userRepo.VerifiedCreateVerified(ctx, verifiedUser); errx != nil {
			errx.AddCommentMessage(fmt.Sprintf("%s While VerifiedCreateVerified", functionName), "Tambah user verifikasi gagal")
			return errx
		}

		// Mail user
		content := helpers.CreateRegistrationMailContent(u.FirstName, u.LastName, h.serviceConfig.WebURL, code)

		if errx := h.mailRepo.SystemSendMail(req.Email, "Verifikasi Akun Kamu", content); errx != nil {
			errx.AddCommentMessage(fmt.Sprintf("%s While SendEmail", functionName), "Kirim email registrasi gagal")
			return errx
		}

		return nil
	}

	//hash password
	if req.Password == "" {
		req.Password = req.FirstName + "." + req.LastName + time.Now().Format("02.01.06")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		errx = serror.NewWithErrorCommentMessage(err, http.StatusInternalServerError, "[UserUsecase.UserRegister] While GenerateFromPassword", "Tidak bisa generate password baru")
		return
	}

	req.Password = string(hashedPassword)

	// Create user with role user and active = 0
	userCreateReq := models.CreateUserRequest{}
	utstruct.InjectStructValue(req, &userCreateReq)
	// Hard Coded Role id, with 2 means "Regular User"
	userCreateReq.RoleId = 2
	userCreateReq.Active = 0

	operatorMail := h.serviceConfig.Name
	operator, erry := h.authCase.AuthGetFromContext(ctx)
	if erry == nil {
		operatorMail = operator.Email
	}
	uID, errx := h.userRepo.UserCreate(ctx, userCreateReq, operatorMail)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserCreate", functionName), "Tidak bisa tambah user baru")
		return errx
	}

	// Create sha256 encrypt code from userid_currenttimestamp_jobhunregis
	now := time.Now().Format("02.01.06")
	// code := h.utils.Encrypt256(idstr, now, "jobhunregis")  //TODO: need to be fixed.
	code := helpers.EncryptMD5(fmt.Sprintf("%d", uID), now, "jobhunregis")

	// Create new verified user
	verifiedUser := models.VerifiedUser{}
	verifiedUser.UserId = uID
	verifiedUser.Code = code

	if _, errx := h.userRepo.VerifiedCreateVerified(ctx, verifiedUser); errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While VerifiedCreateVerified", functionName), "tidak bisa tambah user verifikasi")
		return errx
	}

	// Mail user
	content := helpers.CreateRegistrationMailContent(userCreateReq.FirstName, userCreateReq.LastName, h.serviceConfig.WebURL, code)

	if errx := h.mailRepo.SystemSendMail(req.Email, "Verifikasi Akun Kamu", content); errx != nil {
		errx = serror.NewWithErrorCommentMessage(err, http.StatusInternalServerError, fmt.Sprintf("%s While SendEmail", functionName), "tidak bisa kirim email registrasi")
		return errx
	}

	return
}

func (h *UserUsecase) UserVerifyToken(ctx context.Context, token string) (errx serror.SError) {
	functionName := "[UserUsecase.UserVerifyToken]"

	//handle tx
	err := h.db.StartTx()
	if err != nil {
		errx = serror.NewWithErrorCommentMessage(err, http.StatusInternalServerError, fmt.Sprintf("%s While StartTx", functionName), "Kesalahan Server")
		return
	}
	defer func() {
		if errx != nil {
			err = errx.GetError()
		}
		h.db.SubmitTx(err)
	}()

	vu, errx := h.userRepo.VerifiedGetVerifiedByCode(ctx, token)
	if errx != nil {
		errx = serror.NewWithCommentMessage(http.StatusInternalServerError, fmt.Sprintf("%s While VerifiedGetVerifiedByCode", functionName), "tidak bisa menemukan kode aktivasi")
		return
	}

	now := time.Now()
	then := now.AddDate(0, 0, -1)
	if vu.CreatedAt.Before(then) {
		if vu, errx = h.userRepo.VerifiedDeleteVerified(ctx, vu); errx != nil {
			errx.AddCommentMessage(fmt.Sprintf("%s While VerifiedDeleteVerified", functionName), "tidak bisa tambah user verifikasi")
			return errx
		}
		errx = serror.NewWithCommentMessage(http.StatusInternalServerError, fmt.Sprintf("%s While VerifiedDeleteVerified", functionName), "kode aktivasi sudah kadaluarsa")
		return
	}

	u, errx := h.userRepo.UserGetByIDIgnoreActive(ctx, vu.UserId)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserGetByIDIgnoreActive", functionName), "pengguna tidak ditemukan")
		return
	}

	operatorMail := h.serviceConfig.Name
	operator, erry := h.authCase.AuthGetFromContext(ctx)
	if erry == nil {
		operatorMail = operator.Email
	}

	if errx = h.userRepo.UserUpdateActive(ctx, u, operatorMail); errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserUpdateActive", functionName), "tidak bisa mengaktifkan pengguna")
		return errx
	}

	if _, errx = h.userRepo.VerifiedDeleteVerified(ctx, vu); errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While VerifiedDeleteVerified", functionName), "kesalahan server")
		return errx
	}

	return
}

func (h *UserUsecase) UploadUserImage(ctx context.Context, userID int64, image string) (url string, errx serror.SError) {
	functionName := "[UserUsecase.UploadUserImage]"

	// Upload image
	currentTime := time.Now().Format("20060102150405")
	imagePath := fmt.Sprintf("assets/images/user/user_%d_%s.jpg", userID, currentTime)
	if err := utstorage.UploadFileToLocal(image, imagePath); err != nil {
		errx = serror.NewWithErrorCommentMessage(err, http.StatusInternalServerError, fmt.Sprintf("%s While UploadFileToLocal", functionName), "Gagal Upload gambar ke server")
		return
	}
	defer func() {
		if err := os.Remove(imagePath); err != nil {
			// not doing anything
			errx.AddCommentMessage(fmt.Sprintf("%s While os.Remove", functionName), "Tidak bisa hapus file di local server")
		}
	}()

	var err error
	if url, err = h.cloudStorage.AddCloudFileAndGetURL(imagePath); err != nil {
		errx = serror.NewWithErrorCommentMessage(err, http.StatusInternalServerError, fmt.Sprintf("%s While AddFileToPublicS3", functionName), "Upload file ke cloud gagal")
		return
	}
	return
}

func (h *UserUsecase) UserTransactionHistoryByIDWithPagination(ctx context.Context, search string, userID int64, paginationData models.PaginationData) (resp models.PaginationResponse, errx serror.SError) {
	functionName := "[UserUsecase.UserTransactionHistoryByIDWithPagination]"

	resp, errx = h.userRepo.UserTransactionHistoryWithPaginationByUserID(ctx, search, userID, paginationData)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserTransactionHistoryWithPaginationByUserID", functionName), "Tidak bisa mendapatkan data history")
		return
	}

	return
}

func (h *UserUsecase) UserTransactionFinishedWithPagination(ctx context.Context, paginationData models.PaginationData) (resp models.PaginationResponse, errx serror.SError) {
	functionName := "[UserUsecase.UserTransactionFinishedWithPagination]"

	var (
		userRoleID int64 = 0
	)
	auth, errx := h.authCase.AuthGetFromContext(ctx)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While AuthGetFromContext", functionName))
		return
	}

	userRoleID = auth.ID

	if resp, errx = h.userRepo.UserTransactionFinishedWithPagination(ctx, paginationData, userRoleID); errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While TableSelfOrderWithPagination", functionName), "Gagal Ambil data order")
		return resp, errx
	}
	return resp, nil
}

func (h *UserUsecase) UserProfileGetByID(ctx context.Context, id int64) (user models.UserProfileDetail, errx serror.SError) {
	functionName := "[UserUsecase.UserProfileGetByID]"

	user, errx = h.userRepo.UserProfileGetByID(ctx, id)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserGetByID", functionName), "Pengguna tidak ditemukan")
		return
	}

	return
}

func (h *UserUsecase) UserProfile(ctx context.Context) (user models.UserProfileDetail, errx serror.SError) {
	functionName := "[UserUsecase.UserProfile]"

	operator, errx := h.authCase.AuthGetFromContext(ctx)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While Get Session", functionName))
		return
	}

	user, errx = h.userRepo.UserProfileGetByID(ctx, operator.ID)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserGetByID", functionName), "Pengguna tidak ditemukan")
		return
	}

	return
}

func (h *UserUsecase) UserSelfUpdate(ctx context.Context, req models.UpdateUserRequest) (u models.User, errx serror.SError) {
	functionName := "[UserUsecase.UserSelfUpdate]"

	//handle tx
	err := h.db.StartTx()
	if err != nil {
		errx = serror.NewWithErrorCommentMessage(err, http.StatusInternalServerError, fmt.Sprintf("%s While StartTx", functionName), "Kesalahan Server")
		return
	}
	defer func() {
		if errx != nil {
			err = errx.GetError()
		}
		h.db.SubmitTx(err)
	}()

	operator, errx := h.authCase.AuthGetFromContext(ctx)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While Get Session", functionName))
		return
	}

	var tempImageCode string
	if req.ImageUrl != nil {
		tempImageCode = *req.ImageUrl
		req.ImageUrl = nil
	}
	errx = h.userRepo.UserUpdate(ctx, operator.ID, req, operator.Email)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserUpdate", functionName), "Gagal Sunting Pengguna")
		return
	}

	if tempImageCode != "" {
		imageReq := models.CreateImgRequest{
			Image: tempImageCode,
		}

		_, errx = h.UserUpdateImageURL(ctx, operator.ID, &imageReq)
		if errx != nil {
			errx.AddCommentMessage(fmt.Sprintf("%s While UserUpdateImageURL", functionName), "Gagal Update Foto Pengguna")
			return
		}
	}

	u, errx = h.UserGetByID(ctx, operator.ID)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While UserGetByID", functionName), "Gagal Ambil data User Baru")
		return
	}

	return
}

func (h *UserUsecase) UserSelfAddOrUpdateProfile(ctx context.Context, req models.ProfileRequest) (errx serror.SError) {
	functionName := "[UserUsecase.UserSelfAddOrUpdateProfile]"

	//handle tx
	err := h.db.StartTx()
	if err != nil {
		errx = serror.NewWithErrorCommentMessage(err, http.StatusInternalServerError, fmt.Sprintf("%s While StartTx", functionName), "Kesalahan Server")
		return
	}
	defer func() {
		if errx != nil {
			err = errx.GetError()
		}
		h.db.SubmitTx(err)
	}()

	operator, errx := h.authCase.AuthGetFromContext(ctx)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While Get Session", functionName))
		return
	}

	_, errx = h.userRepo.ProfileGetByID(ctx, operator.ID)
	if errx != nil {
		errx = h.userRepo.UserAddProfile(ctx, operator.ID, req, operator.Email)
		if errx != nil {
			errx.AddCommentMessage(fmt.Sprintf("%s While CreateUserProfile", functionName), "Gagal Sunting Pengguna")
		}
		return
	}

	errx = h.userRepo.UserUpdateProfile(ctx, operator.ID, req, operator.Email)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While CreateUserProfile", functionName), "Gagal Sunting Pengguna")
		return
	}
	return
}

func (h *UserUsecase) UserAddOrUpdateProfileByID(ctx context.Context, req models.ProfileRequest, userID int64) (errx serror.SError) {
	functionName := "[UserUsecase.UserAddOrUpdateProfileByID]"

	//handle tx
	err := h.db.StartTx()
	if err != nil {
		errx = serror.NewWithErrorCommentMessage(err, http.StatusInternalServerError, fmt.Sprintf("%s While StartTx", functionName), "Kesalahan Server")
		return
	}
	defer func() {
		if errx != nil {
			err = errx.GetError()
		}
		h.db.SubmitTx(err)
	}()

	operator, errx := h.authCase.AuthGetFromContext(ctx)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While Get Session", functionName))
		return
	}

	_, errx = h.userRepo.ProfileGetByID(ctx, userID)
	if errx != nil {
		errx = h.userRepo.UserAddProfile(ctx, userID, req, operator.Email)
		if errx != nil {
			errx.AddCommentMessage(fmt.Sprintf("%s While CreateUserProfile", functionName), "Gagal Sunting Pengguna")
		}
		return
	}

	errx = h.userRepo.UserUpdateProfile(ctx, userID, req, operator.Email)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While CreateUserProfile", functionName), "Gagal Sunting Pengguna")
		return
	}
	return
}
