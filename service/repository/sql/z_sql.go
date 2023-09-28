package sql

import (
	"context"

	"github.com/DeniesKresna/gobridge/sdb"
	"github.com/DeniesKresna/gobridge/serror"
	"github.com/DeniesKresna/myqgen2/qgen"

	"github.com/DeniesKresna/bkn/models"
)

// FORMER OF USER SERVICE

type UserSqlRepository struct {
	db *sdb.DBInstance
	q  *qgen.Obj
}

func InitUserSqlRepository(db *sdb.DBInstance, svc *models.Service, q *qgen.Obj) IUserSqlRepository {
	return &UserSqlRepository{
		db: db,
		q:  q,
	}
}

type IUserSqlRepository interface {
	//user
	UserGetByID(ctx context.Context, id int64) (u models.User, errx serror.SError)
	UserGetByIDIgnoreActive(ctx context.Context, id int64) (u models.User, errx serror.SError)
	UserRoleGetByID(ctx context.Context, id int64) (u models.UserRole, errx serror.SError)
	UserGetByEmail(ctx context.Context, email string) (u models.User, errx serror.SError)
	UserGetByEmailIgnoreActive(ctx context.Context, email string) (u models.User, errx serror.SError)
	UserCreate(ctx context.Context, req models.CreateUserRequest, operator string) (uID int64, errx serror.SError)
	UserUpdatePassword(ctx context.Context, u models.User, password string, operator string) (errx serror.SError)
	UserUpdateImageURL(ctx context.Context, u models.User, imageUrl string, operator string) (errx serror.SError)
	UserUpdate(ctx context.Context, id int64, req models.UpdateUserRequest, operator string) (errx serror.SError)
	UserUpdateRole(ctx context.Context, u models.User, roleName string, operator string) (errx serror.SError)
	UserIndexWithPagination(ctx context.Context, filter models.UserSearch, paging models.PaginationData) (resp models.PaginationResponse, errx serror.SError)
	UserDelete(ctx context.Context, id int64, operator string) (errx serror.SError)
	UserUpdateActive(ctx context.Context, u models.User, operator string) (errx serror.SError)
	UserProfileGetByID(ctx context.Context, id int64) (u models.UserProfileDetail, errx serror.SError)

	UserTransactionHistoryWithPaginationByUserID(ctx context.Context, search string, userID int64, paginationData models.PaginationData) (resp models.PaginationResponse, errx serror.SError)
	UserTransactionFinishedWithPagination(ctx context.Context, paging models.PaginationData, userRoleID int64) (resp models.PaginationResponse, errx serror.SError)

	//role
	RoleGetByID(ctx context.Context, id int64) (role models.Role, errx serror.SError)
	RoleCreate(ctx context.Context, req models.CreateRoleRequest, operator string) (rID int64, errx serror.SError)
	RoleGetByName(ctx context.Context, name string) (role models.Role, errx serror.SError)

	//profiles
	UserAddProfile(ctx context.Context, userID int64, req models.ProfileRequest, operator string) (errx serror.SError)
	UserUpdateProfile(ctx context.Context, userID int64, req models.ProfileRequest, operator string) (errx serror.SError)
	ProfileGetByID(ctx context.Context, id int64) (p models.ProfileDetail, errx serror.SError)

	//verified
	VerifiedCreateVerified(ctx context.Context, u models.VerifiedUser) (id int64, errx serror.SError)
	VerifiedGetVerifiedByUserID(ctx context.Context, id int64) (u models.VerifiedUser, errx serror.SError)
	VerifiedDeleteVerified(ctx context.Context, u models.VerifiedUser) (u2 models.VerifiedUser, errx serror.SError)
	VerifiedGetVerifiedByCode(ctx context.Context, code string) (u models.VerifiedUser, errx serror.SError)
}

type ExpertSqlRepository struct {
	db *sdb.DBInstance
	q  *qgen.Obj
}

func InitExpertSqlRepository(db *sdb.DBInstance, svc *models.Service, q *qgen.Obj) IExpertSqlRepository {
	return &ExpertSqlRepository{
		db: db,
		q:  q,
	}
}

type IExpertSqlRepository interface {
	ExpertPublishedProfileGetByID(ctx context.Context, id int64) (expert models.Expert, errx serror.SError)
	ExpertPermanentProfileGetByID(ctx context.Context, id int64) (expert models.Expert, errx serror.SError)
	ExpertFeatures(ctx context.Context) (featureExperts []models.ExpertFeature, errx serror.SError)
	ExpertPublishedSearchWithPagination(ctx context.Context, paginationData models.PaginationData, filter models.ExpertSearch) (resp models.PaginationResponse, errx serror.SError)
	ExpertListPublishedSearch(ctx context.Context, search string) (experts []models.ExpertOptionData, errx serror.SError)
	ExpertPermanentSearchWithPagination(ctx context.Context, paginationData models.PaginationData, filter models.ExpertSearch) (resp models.PaginationResponse, errx serror.SError)
	ExpertCreate(ctx context.Context, request models.ExpertCreateRequest, userID int64, isPermanent bool, operator string) (errx serror.SError)
	ExpertGetByUserID(ctx context.Context, userID int64) (expert models.Expert, errx serror.SError)
	ExpertDataIgnoreActiveGetByID(ctx context.Context, id int64) (expert models.Expert, errx serror.SError)
	ExpertUpdateStatus(ctx context.Context, active int, id int64, operator string) (errx serror.SError)
	ExpertEdit(ctx context.Context, id int64, request models.ExpertUpdateRequest, operator string) (errx serror.SError)
	ExpertGetUserIDByExpertID(ctx context.Context, id int64) (uID int64, errx serror.SError)
	ExpertDelete(ctx context.Context, id int64, operator string) (errx serror.SError)
	ExpertAcceptedProfileGetByID(ctx context.Context, id int64) (expert models.Expert, errx serror.SError)

	// expert temp
	ExpertTempProfileGetByID(ctx context.Context, id int64) (expert models.Expert, errx serror.SError)
	ExpertTempSearchWithPagination(ctx context.Context, paginationData models.PaginationData, filter models.ExpertSearch) (resp models.PaginationResponse, errx serror.SError)
	ExpertCreateHistory(ctx context.Context, paginationData models.PaginationData, filter models.ExpertSearch, userID int64) (resp models.PaginationResponse, errx serror.SError)

	//transaction
	DetailOrderByServiceAndId(ctx context.Context, id int64, service string) (proposal models.DetailDashboardJSON, errx serror.SError)
	TableOrderWithPaginationByObject(ctx context.Context, filter models.OrderFilterAndSearch, paging models.PaginationData, userRoleID int64, object string) (resp models.PaginationResponse, errx serror.SError)
	DetailRequestExpertById(ctx context.Context, id int64) (res models.DetailServiceRequestExpert, errx serror.SError)
	TableServiceRequestExpertWithPagination(ctx context.Context, filter models.ServiceRequestFilterAndSearch, paging models.PaginationData) (resp models.PaginationResponse, errx serror.SError)
	OrderGetByServiceAndId(ctx context.Context, id int64, object string) (order models.Order, errx serror.SError)
	DealUpdate(ctx context.Context, id int64, req models.Deal) (errx serror.SError)
	OrderUpdate(ctx context.Context, id int64, orderArg models.Order, operator string) (errx serror.SError)

	CreateExpertRequestment(ctx context.Context, orderArg models.ServiceRequestExpert, operator string) (errx serror.SError)

	ProposeExpertService(ctx context.Context, req models.Proposal) (detailID int64, errx serror.SError)
	CreateOrder(ctx context.Context, orderArg models.Order, operator string) (errx serror.SError)
}

type CourseSqlRepository struct {
	db *sdb.DBInstance
	q  *qgen.Obj
}

func InitCourseSqlRepository(db *sdb.DBInstance, svc *models.Service, q *qgen.Obj) ICourseSqlRepository {
	return &CourseSqlRepository{
		db: db,
		q:  q,
	}
}

type ICourseSqlRepository interface {
	CoursePublishedSearchWithPagination(ctx context.Context, paging models.PaginationData, filter models.CourseSearch) (resp models.PaginationResponse, errx serror.SError)
	CourseInterestPublishedSearchWithPagination(ctx context.Context, paging models.PaginationData, filter models.CourseSearch, userID int64) (resp models.PaginationResponse, errx serror.SError)
	CourseDataWithExpertIgnoreActiveGetByID(ctx context.Context, id int64) (course models.CourseWithExpert, errx serror.SError)
	CourseDataIgnoreActiveGetByCode(ctx context.Context, code string) (course models.Course, errx serror.SError)
	CourseDataIgnoreActiveGetByID(ctx context.Context, id int64) (course models.Course, errx serror.SError)
	CoursePublishedGetByID(ctx context.Context, id int64) (course models.Course, errx serror.SError)
	CourseCreate(ctx context.Context, req models.CreateCourseRequest, operator string) (uID int64, errx serror.SError)
	CourseUpdateByID(ctx context.Context, u models.Course, operator string) (errx serror.SError)
	CourseUpdateImageByID(ctx context.Context, id int64, image string, operator string) (errx serror.SError)
	CourseGetCodePrev(ctx context.Context, courseType string, courseProgram string) (codePrev string, errx serror.SError)
	CourseDashboardSearchWithPagination(ctx context.Context, paging models.PaginationData, filter models.CourseDashboardSearch) (resp models.PaginationResponse, errx serror.SError)
	CourseListUserRegisterSearchWithPagination(ctx context.Context, paging models.PaginationData, courseID int64, search string) (resp models.PaginationResponse, errx serror.SError)
	CourseListUserInterestSearchWithPagination(ctx context.Context, paging models.PaginationData, courseID int64, search string) (resp models.PaginationResponse, errx serror.SError)
	CourseInterestDataWithExpertGetByID(ctx context.Context, userID int64, id int64) (course models.CourseWithExpert, errx serror.SError)

	// Course User
	CourseCreateUser(ctx context.Context, req models.CreateCourseUserRequest, operator string) (uID int64, errx serror.SError)
	CourseGetCourseUserByID(ctx context.Context, id int64) (courseUser models.CourseUser, errx serror.SError)

	// course user interest
	CourseInterestCreate(ctx context.Context, courseID int64, userID int64) (courseInterestID int64, errx serror.SError)
	CourseGetCourseInterestByID(ctx context.Context, ID int64) (courseInterest models.CourseUserInterest, errx serror.SError)
	CourseGetCourseInterestByCourseIDAndUserID(ctx context.Context, courseID int64, userID int64) (courseInterest models.CourseUserInterest, errx serror.SError)
	CourseInterestUpdate(ctx context.Context, courseInterest models.CourseUserInterest) (errx serror.SError)
}

type PaymentSqlRepository struct {
	db *sdb.DBInstance
	q  *qgen.Obj
}

func InitPaymentSqlRepository(db *sdb.DBInstance, svc *models.Service, q *qgen.Obj) IPaymentSqlRepository {
	return &PaymentSqlRepository{
		db: db,
		q:  q,
	}
}

type IPaymentSqlRepository interface {
	PaymentGetByCode(ctx context.Context, code string) (resp models.Payment, errx serror.SError)
	PaymentGetBaseCodeNotExpireNotPaid(ctx context.Context, base_code string) (resp []models.Payment, errx serror.SError)
	PaymentGetByID(ctx context.Context, id int64) (resp models.Payment, errx serror.SError)
	PaymentCreateInvoice(ctx context.Context, req models.CreatePaymentRequest, operator string) (errx serror.SError)
	PaymentUpdateInvoiceAfterCreated(ctx context.Context, req models.CreatePaymentRequest, operator string) (errx serror.SError)
	PaymentUpdateCallback(ctx context.Context, req models.UpdatePaymentCallback, operator string) (errx serror.SError)
}
