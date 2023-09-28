package sql

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/DeniesKresna/gobridge/serror"
	"github.com/DeniesKresna/gohelper/utstruct"
	"github.com/DeniesKresna/myqgen2/qgen"

	"github.com/DeniesKresna/bkn/models"
	"github.com/DeniesKresna/bkn/service/repository/sql/queries"
)

func (r *UserSqlRepository) UserGetByID(ctx context.Context, id int64) (u models.User, errx serror.SError) {
	functionName := "[UserSqlRepository.UserGetByID]"

	err := r.db.Take(&u, r.q.Build(queries.GetActiveUser, qgen.Args{
		Fields: []string{
			"u.*",
		},
		Conditions: map[string]interface{}{
			"id": id,
		},
	}))
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusNotFound, fmt.Sprintf("%s While Query UserGetByID (id: %d)", functionName, id))
	}
	return
}

func (r *UserSqlRepository) UserGetByIDIgnoreActive(ctx context.Context, id int64) (u models.User, errx serror.SError) {
	functionName := "[UserSqlRepository.UserGetByIDIgnoreActive]"

	err := r.db.Take(&u, r.q.Build(queries.GetUser, qgen.Args{
		Fields: []string{
			"u.*",
		},
		Conditions: map[string]interface{}{
			"id": id,
		},
	}))
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusNotFound, fmt.Sprintf("%s While Query GetUserByIDIgnoreActive (id: %d)", functionName, id))
	}
	return
}

func (r *UserSqlRepository) UserRoleGetByID(ctx context.Context, id int64) (u models.UserRole, errx serror.SError) {
	functionName := "[UserSqlRepository.UserRoleGetByID]"

	err := r.db.Take(&u, r.q.Build(queries.GetActiveUser, qgen.Args{
		Fields: []string{
			"u.id",
			"u.created_at",
			"u.updated_at",
			"u.deleted_at",
			"u.first_name",
			"u.last_name",
			"u.email",
			"u.phone",
			"u.image_url",
			"u.role_id",
			"r.name",
		},
		Conditions: map[string]interface{}{
			"id": id,
		},
	}))
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusNotFound, fmt.Sprintf("%s While Query UserRoleGetByID (id: %d)", functionName, id))
	}
	return
}

func (r *UserSqlRepository) UserGetByEmail(ctx context.Context, email string) (u models.User, errx serror.SError) {
	functionName := "[UserSqlRepository.UserGetByEmail]"

	err := r.db.Take(&u, r.q.Build(queries.GetActiveUser, qgen.Args{
		Fields: []string{
			"u.*",
		},
		Conditions: map[string]interface{}{
			"email": email,
		},
	}))
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusNotFound, fmt.Sprintf("%s While Query UserGetByEmail (email: %s)", functionName, email))
	}
	return
}

func (r *UserSqlRepository) UserGetByEmailIgnoreActive(ctx context.Context, email string) (u models.User, errx serror.SError) {
	functionName := "[UserSqlRepository.UserGetByEmail]"

	err := r.db.Take(&u, r.q.Build(queries.GetUser, qgen.Args{
		Fields: []string{
			"u.*",
		},
		Conditions: map[string]interface{}{
			"email": email,
		},
	}))
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusNotFound, fmt.Sprintf("%s While Query GetUserByEmailIgnoreActive (email: %s)", functionName, email))
	}
	return
}

func (r *UserSqlRepository) UserCreate(ctx context.Context, req models.CreateUserRequest, operator string) (uID int64, errx serror.SError) {
	functionName := "[UserSqlRepository.UserCreate]"

	res, err := r.db.Exec(queries.CreateUser,
		req.FirstName, req.LastName, req.Email, req.Phone, req.RoleId, req.Password, req.ImageUrl, req.Active, operator, operator,
	)
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, fmt.Sprintf("%s While Query UserCreate (data: %+v)", functionName, req))
		return
	}
	uID, err = res.LastInsertId()
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, fmt.Sprintf("%s While Get Last Inserted ID", functionName))
		return
	}
	return
}

func (r *UserSqlRepository) UserUpdatePassword(ctx context.Context, u models.User, password string, operator string) (errx serror.SError) {
	functionName := "[UserSqlRepository.UserUpdatePassword]"

	_, err := r.db.Exec(r.q.Build(queries.UpdateActiveUser, qgen.Args{
		Updates: map[string]interface{}{
			"u.password": password,
		},
		Conditions: map[string]interface{}{
			"id": u.ID,
		},
	}), operator)
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, fmt.Sprintf("%s While Query Update Password", functionName))
		return
	}
	return
}

func (r *UserSqlRepository) UserUpdateImageURL(ctx context.Context, u models.User, imageUrl string, operator string) (errx serror.SError) {
	functionName := "[UserSqlRepository.UserUpdateImageURL]"

	_, err := r.db.Exec(r.q.Build(queries.UpdateActiveUser, qgen.Args{
		Updates: map[string]interface{}{
			"u.image_url": imageUrl,
		},
		Conditions: map[string]interface{}{
			"id": u.ID,
		},
	}), operator)

	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, fmt.Sprintf("%s While Query Update Image URL", functionName))
		return
	}
	return
}

func (r *UserSqlRepository) UserUpdate(ctx context.Context, id int64, req models.UpdateUserRequest, operator string) (errx serror.SError) {
	functionName := "[UserSqlRepository.UserUpdate]"

	_, err := r.db.Exec(r.q.Build(queries.UpdateActiveUser, qgen.Args{
		Updates: map[string]interface{}{
			"u.first_name": req.FirstName,
			"u.last_name":  req.LastName,
			"u.phone":      req.Phone,
			"u.email":      req.Email,
		},
		Conditions: map[string]interface{}{
			"id": id,
		},
	}), operator)
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, fmt.Sprintf("%s While Query Update User (id:%d)", functionName, id))
		return
	}
	return
}

func (r *UserSqlRepository) UserUpdateRole(ctx context.Context, u models.User, roleName string, operator string) (errx serror.SError) {
	functionName := "[UserSqlRepository.UserUpdateRole]"

	_, err := r.db.Exec(queries.UpdateUserRoleByID, roleName, operator, u.ID)

	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, fmt.Sprintf("%s While Query UpdateUserRoleByID (userID:%d, roleName: %s)", functionName, u.ID, roleName))
		return
	}
	return
}

func (r *UserSqlRepository) UserIndexWithPagination(ctx context.Context, filter models.UserSearch, paging models.PaginationData) (resp models.PaginationResponse, errx serror.SError) {
	functionName := "[UserSqlRepository.UserIndexWithPagination]"

	var (
		users     []models.UserTableData
		totalRows int64 = 0
		condition       = make(map[string]interface{})
	)

	if paging.Sort == "" {
		paging.Sort = "-u.created_at"
	}

	if filter.Search != "" {
		search := strings.ToLower(filter.Search)
		searchArg := "%" + search + "%"
		condition["name:LIKE"] = searchArg
	}

	if filter.Email != "" {
		email := "%" + filter.Email + "%"
		condition["email:LIKE"] = email
	}

	if filter.Phone != "" {
		phone := "%" + filter.Phone + "%"
		condition["phone:Like"] = phone
	}

	condition["active"] = 1

	err := r.db.Take(&totalRows, r.q.Build(queries.GetUser, qgen.Args{
		Fields: []string{
			"u.count",
		},
		Conditions: condition,
	}))
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusNotFound, fmt.Sprintf("%s While Build Query User Pagination", functionName))
		return
	}

	err = r.db.Select(&users, r.q.Build(queries.GetUser, qgen.Args{
		Fields: []string{
			"u.id",
			"u.full_name",
			"u.image_url",
			"u.email",
			"u.role_id",
			"u.phone",
		},
		Conditions: condition,
		Sorting:    []string{paging.Sort},
		Limit:      paging.Limit,
		Offset:     int64(paging.Offset),
	}))
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusNotFound, fmt.Sprintf("%s While Query GetUsersTableData", functionName))
		return
	}

	err = utstruct.InjectStructValue(paging, &resp.Pagination)
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, fmt.Sprintf("%s While inject value to struct. value: %+v", functionName, paging))
		return
	}
	resp.Pagination.Total = totalRows

	usersData, err := utstruct.ConvertToSliceMapInterface(users)
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, fmt.Sprintf("%s While json.Marshal", functionName))
		return
	}
	resp.Data = usersData

	return
}

func (r *UserSqlRepository) UserDelete(ctx context.Context, id int64, operator string) (errx serror.SError) {
	functionName := "[UserSqlRepository.UserDelete]"

	_, err := r.db.Exec(r.q.Build(queries.UpdateUser, qgen.Args{
		Updates: map[string]interface{}{
			"u.deleted_by": operator,
			"u.deleted_at": "__NOW()__",
		},
		Conditions: map[string]interface{}{
			"id": id,
		},
	}), operator)
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, fmt.Sprintf("%s While Query Delete User (id:%d)", functionName, id))
	}
	return
}

func (r *UserSqlRepository) UserUpdateActive(ctx context.Context, u models.User, operator string) (errx serror.SError) {
	functionName := "[UserSqlRepository.UserUpdateActive]"

	_, err := r.db.Exec(r.q.Build(queries.UpdateUser, qgen.Args{
		Updates: map[string]interface{}{
			"u.active": 1,
		},
		Conditions: map[string]interface{}{
			"id": u.ID,
		},
	}), operator)
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, fmt.Sprintf("%s While Query Update activation status (userId:%d)", functionName, u.ID))
		return
	}
	return
}

func (r *UserSqlRepository) UserTransactionHistoryWithPaginationByUserID(ctx context.Context, search string, userID int64, paging models.PaginationData) (resp models.PaginationResponse, errx serror.SError) {
	functionName := "[UserSqlRepository.UserTransactionHistoryWithPaginationByUserID]"

	var (
		orders    []models.OrderTableData
		totalRows int64 = 0
	)

	if paging.Sort == "" {
		paging.Sort = "-o.created_at"
	}

	err := r.db.Take(&totalRows, r.q.Build(queries.GetOrder, qgen.Args{
		Fields: []string{
			"o.count",
		},
		Conditions: map[string]interface{}{
			"userID": userID,
		},
	}))
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusNotFound, fmt.Sprintf("%s While Build Query User Pagination", functionName))
		return
	}

	err = r.db.Select(&orders, r.q.Build(queries.GetOrder, qgen.Args{
		Fields: []string{
			"o.id",
			"o.finishStatus",
			"o.serviceName",
			"o.updated_at",
		},
		Conditions: map[string]interface{}{
			"userID": userID,
		},
		Sorting: []string{paging.Sort},
		Limit:   paging.Limit,
		Offset:  int64(paging.Offset),
	}))
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusNotFound, fmt.Sprintf("%s While Query GetOrdersTableData", functionName))
		return
	}

	err = utstruct.InjectStructValue(paging, &resp.Pagination)
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, fmt.Sprintf("%s While inject value to struct. value: %+v", functionName, paging))
		return
	}
	resp.Pagination.Total = totalRows

	ordersData, err := utstruct.ConvertToSliceMapInterface(orders)
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, fmt.Sprintf("%s While json.Marshal", functionName))
		return
	}
	resp.Data = ordersData

	return
}

func (r *UserSqlRepository) UserTransactionFinishedWithPagination(ctx context.Context, paging models.PaginationData, userRoleID int64) (resp models.PaginationResponse, errx serror.SError) {
	functionName := "[UserSqlRepository.UserTransactionFinishedWithPagination]"

	var (
		orders         []models.OrderTableData
		ordersResponse []models.OrderTableResponse
		totalRows      int64 = 0
	)

	queryCondition := fmt.Sprintf(" where o.deleted_at is null and JSON_VALUE(d.deal, '$.service.deleted_at' RETURNING CHAR(30)) is null and o.is_finished = 1 and o.user_id = %d", userRoleID)

	queryCountGetOrderTable := queries.GetCountOrderTableRow + queryCondition
	queryPaginationGetOrderTable := queries.GetOrderTableRow + queryCondition + " order by o.created_at desc limit ? offset ? "

	err := r.db.Take(&totalRows, queryCountGetOrderTable)
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusNotFound, fmt.Sprintf("%s While Build Query Order Pagination", functionName))
		return
	}

	err = r.db.Select(&orders, queryPaginationGetOrderTable, paging.Limit, paging.Offset)
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusNotFound, fmt.Sprintf("%s While Query GetOrderTableData", functionName))
		return
	}

	err = utstruct.InjectStructValue(paging, &resp.Pagination)
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, fmt.Sprintf("%s While inject value to struct. value: %+v", functionName, paging))
		return
	}
	resp.Pagination.Total = totalRows

	for _, order := range orders {
		split := strings.Split(order.Status, ", ")
		o := models.OrderTableResponse{}
		o.ID = order.ID
		o.Status = split[0]
		o.StatusPaid = split[1]
		o.CreatedAt = order.CreatedAt
		o.User.Id = order.UserId
		o.User.Name = order.UserName
		o.Expert.Id = order.ExpertId
		o.Expert.Name = order.ExpertName
		o.Service = order.Service
		ordersResponse = append(ordersResponse, o)
	}

	ordersData, err := utstruct.ConvertToSliceMapInterface(ordersResponse)
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, fmt.Sprintf("%s While json.Marshal", functionName))
		return
	}
	resp.Data = ordersData

	return
}

func (r *UserSqlRepository) UserProfileGetByID(ctx context.Context, id int64) (u models.UserProfileDetail, errx serror.SError) {
	functionName := "[UserSqlRepository.UserGetByID]"

	err := r.db.Take(&u, r.q.Build(queries.GetActiveUser, qgen.Args{
		Fields: []string{
			"u.*",
			"p.profession",
			"p.company",
			"p.domicile",
		},
		Conditions: map[string]interface{}{
			"id": id,
		},
	}))
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusNotFound, fmt.Sprintf("%s While Query UserGetByID (id: %d)", functionName, id))
	}
	return
}

func (r *UserSqlRepository) UserAddProfile(ctx context.Context, userID int64, req models.ProfileRequest, operator string) (errx serror.SError) {
	functionName := "[UserSqlRepository.UserUpdate]"

	res, err := r.db.Exec(queries.CreateProfile,
		userID, req.Profession, req.Company, req.Domicile, operator, operator,
	)
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, fmt.Sprintf("%s While Query Add Profile (data: %+v)", functionName, req))
		return
	}
	_, err = res.LastInsertId()
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, fmt.Sprintf("%s While Get Last Inserted ID", functionName))
		return
	}
	return
}

func (r *UserSqlRepository) UserUpdateProfile(ctx context.Context, userID int64, req models.ProfileRequest, operator string) (errx serror.SError) {
	functionName := "[UserSqlRepository.UserUpdateProfile]"

	_, err := r.db.Exec(r.q.Build(queries.UpdateProfile, qgen.Args{
		Updates: map[string]interface{}{
			"p.profession": req.Profession,
			"p.company":    req.Company,
			"p.domicile":   req.Domicile,
		},
		Conditions: map[string]interface{}{
			"user_id": userID,
		},
	}), operator)
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, fmt.Sprintf("%s While Query Update Profile (user_id:%d)", functionName, userID))
		return
	}
	return
}

func (r *UserSqlRepository) ProfileGetByID(ctx context.Context, userID int64) (p models.ProfileDetail, errx serror.SError) {
	functionName := "[UserSqlRepository.ProfileGetByID]"

	err := r.db.Take(&p, r.q.Build(queries.GetProfileByID, qgen.Args{
		Fields: []string{
			"p.profession",
			"p.company",
			"p.domicile",
		},
		Conditions: map[string]interface{}{
			"user_id": userID,
		},
	}))
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusNotFound, fmt.Sprintf("%s While Query ProfileGetByID (user_id: %d)", functionName, userID))
	}
	return
}
