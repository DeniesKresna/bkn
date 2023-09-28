package queries

const (
	CreateUserVerify = `
		insert into verified_user (user_id, code, created_at) values (?, ?, NOW())
	`

	GetUserVerifyByUserID = `
		select * from verified_user where user_id = ? order by created_at desc limit 1
	`

	HardDeleteUserVerifyByID = `
		delete from verified_user where id = ?
	`

	GetUserVerifyByCode = `
	select * from verified_user where code = ?
	`
)
