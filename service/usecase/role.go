package usecase

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DeniesKresna/bkn/models"
	"github.com/DeniesKresna/gobridge/serror"
)

func (h *UserUsecase) RoleCreate(ctx context.Context, req models.CreateRoleRequest) (role models.Role, errx serror.SError) {
	functionName := "[Usecase.RoleCreate]"

	//handle tx
	err := h.db.StartTx()
	if err != nil {
		errx = serror.NewWithErrorCommentMessage(err, http.StatusInternalServerError, fmt.Sprintf("%s While StartTx", functionName), "Kesalahan Server")
		return
	}
	defer h.db.SubmitTx(errx.GetError())

	operator, errx := h.authCase.AuthGetFromContext(ctx)
	if errx != nil {
		errx.AddComment(fmt.Sprintf("%s While Get Session", functionName))
		return
	}

	rID, errx := h.userRepo.RoleCreate(ctx, req, operator.Email)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While RoleCreate", functionName), "Tambah role gagal")
		return
	}

	role, errx = h.RoleGetByID(ctx, rID)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While RoleGetByID", functionName), "Role tidak ditemukan")
		return
	}

	return
}

func (h *UserUsecase) RoleGetByID(ctx context.Context, id int64) (role models.Role, errx serror.SError) {
	functionName := "[Usecase.RoleGetByID]"

	role, errx = h.userRepo.RoleGetByID(ctx, id)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While RoleGetByID", functionName), "Role tidak ditemukan")
		return
	}
	return
}

func (h *UserUsecase) RoleGetByName(ctx context.Context, name string) (role models.Role, errx serror.SError) {
	functionName := "[Usecase.RoleGetByName]"

	role, errx = h.userRepo.RoleGetByName(ctx, name)
	if errx != nil {
		errx.AddCommentMessage(fmt.Sprintf("%s While RoleGetByName", functionName), "Role tidak ditemukan")
		return
	}
	return
}
