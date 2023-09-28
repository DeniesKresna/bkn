package usecase

import (
	"context"

	"github.com/DeniesKresna/bkn/models"
	"github.com/DeniesKresna/bkn/service/repository/mail"
	"github.com/DeniesKresna/bkn/service/repository/sql"
	"github.com/DeniesKresna/bkn/service/repository/storage"
	"github.com/DeniesKresna/gobridge/sdb"
	"github.com/DeniesKresna/gobridge/serror"
)

type UserUsecase struct {
	serviceConfig *models.Service
	db            *sdb.DBInstance
	cloudStorage  storage.IStorage
	mailRepo      mail.IMail
	userRepo      sql.IUserSqlRepository
	authCase      IAuthUsecase
}

func InitUserUsecase(db *sdb.DBInstance, serviceConfig *models.Service, authCase IAuthUsecase, userRepo sql.IUserSqlRepository, cloudStorage storage.IStorage, mailRepo mail.IMail) IUserUsecase {
	return &UserUsecase{
		serviceConfig: serviceConfig,
		db:            db,
		cloudStorage:  cloudStorage,
		mailRepo:      mailRepo,
		authCase:      authCase,
		userRepo:      userRepo,
	}
}

type IUserUsecase interface {
	//user
	UserCreate(ctx context.Context, req models.CreateUserRequest) (u models.User, errx serror.SError)
	UserUpdate(ctx context.Context, id int64, req models.UpdateUserRequest) (u models.User, errx serror.SError)
	UserRoleGetByID(ctx context.Context, id int64) (u models.UserRole, errx serror.SError)
	UserGetByID(ctx context.Context, id int64) (u models.User, errx serror.SError)
	UserGetByEmail(ctx context.Context, email string) (u models.User, errx serror.SError)
	UserSelfUpdatePassword(ctx context.Context, req models.CreatePasswordRequest) (u models.User, errx serror.SError)
	UserUpdatePassword(ctx context.Context, id int64, req models.CreatePasswordRequest) (u models.User, errx serror.SError)
	UserUpdateImageURL(ctx context.Context, id int64, req *models.CreateImgRequest) (u models.User, errx serror.SError)
	UserRegularCreate(ctx context.Context, req models.CreateUserRequest) (u models.User, errx serror.SError)
	UserUpdateRole(ctx context.Context, id int64, roleName string) (u models.User, errx serror.SError)
	UserIndexWithPagination(ctx context.Context, filter models.UserSearch, paginationData models.PaginationData) (resp models.PaginationResponse, errx serror.SError)
	UserDelete(ctx context.Context, id int64) (errx serror.SError)
	UserRegister(ctx context.Context, req *models.UserRegister) (errx serror.SError)
	UserVerifyToken(ctx context.Context, token string) (errx serror.SError)
	UploadUserImage(ctx context.Context, userID int64, image string) (url string, errx serror.SError)
	UserProfile(ctx context.Context) (user models.UserProfileDetail, errx serror.SError)
	UserProfileGetByID(ctx context.Context, id int64) (user models.UserProfileDetail, errx serror.SError)
	UserSelfUpdate(ctx context.Context, req models.UpdateUserRequest) (u models.User, errx serror.SError)

	// role
	RoleCreate(ctx context.Context, req models.CreateRoleRequest) (role models.Role, errx serror.SError)
	RoleGetByID(ctx context.Context, id int64) (role models.Role, errx serror.SError)
	RoleGetByName(ctx context.Context, name string) (role models.Role, errx serror.SError)

	//profile
	UserSelfAddOrUpdateProfile(ctx context.Context, req models.ProfileRequest) (errx serror.SError)
	UserAddOrUpdateProfileByID(ctx context.Context, req models.ProfileRequest, userID int64) (errx serror.SError)
}

type AuthUsecase struct {
	db       *sdb.DBInstance
	userRepo sql.IUserSqlRepository
}

func InitAuthUsecase(db *sdb.DBInstance, userRepo sql.IUserSqlRepository) IAuthUsecase {
	return &AuthUsecase{
		db:       db,
		userRepo: userRepo,
	}
}

type IAuthUsecase interface {
	//auth
	AuthGetFromContext(ctx context.Context) (res models.UserRole, errx serror.SError)
	AuthGetSession(ctx context.Context) (a models.AuthResponse, errx serror.SError)
	AuthLogin(ctx context.Context, email string, password string) (authResp models.AuthResponse, errx serror.SError)
}
