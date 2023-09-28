package sql

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DeniesKresna/bkn/models"
	"github.com/DeniesKresna/bkn/service/repository/sql/queries"
	"github.com/DeniesKresna/gobridge/serror"
	"github.com/DeniesKresna/myqgen2/qgen"
)

func (r *UserSqlRepository) RoleGetByID(ctx context.Context, id int64) (role models.Role, errx serror.SError) {
	functionName := "[UserSqlRepository.RoleGetByID]"

	err := r.db.Take(&role, r.q.Build(queries.GetRole, qgen.Args{
		Fields: []string{
			"r.*",
		},
		Conditions: map[string]interface{}{
			"id": id,
		},
	}))
	if err != nil {
		return role, serror.NewWithErrorComment(err, http.StatusNotFound, fmt.Sprintf("%s While Query RoleGetByID (id: %d)", functionName, id))
	}
	return role, nil
}

func (r *UserSqlRepository) RoleGetByName(ctx context.Context, name string) (role models.Role, errx serror.SError) {
	functionName := "[UserSqlRepository.RoleGetByName]"

	err := r.db.Take(&role, r.q.Build(queries.GetRole, qgen.Args{
		Fields: []string{
			"r.*",
		},
		Conditions: map[string]interface{}{
			"name": name,
		},
	}))
	if err != nil {
		return role, serror.NewWithErrorComment(err, http.StatusNotFound, fmt.Sprintf("%s While Query RoleGetByName (name: %s)", functionName, name))
	}
	return role, nil
}

func (r *UserSqlRepository) RoleCreate(ctx context.Context, req models.CreateRoleRequest, operator string) (rID int64, errx serror.SError) {
	functionName := "[UserSqlRepository.RoleCreate]"

	res, err := r.db.Exec(queries.CreateRole,
		req.Name,
		operator,
	)
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, fmt.Sprintf("%s While Query RoleCreate (role name: %s)", functionName, req.Name))
		return
	}
	rID, err = res.LastInsertId()
	if err != nil {
		errx = serror.NewWithErrorComment(err, http.StatusInternalServerError, fmt.Sprintf("%s While Get Last Inserted ID", functionName))
		return
	}
	return
}
