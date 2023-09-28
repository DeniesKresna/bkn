package queries

const (
	GetUser = `
		{
			"select": [
				{"col":"u.count", "as": "user_count", "value": "count(*)"},
				{"col": "u.*"},
				{"col": "r.name", "as": "role_name", "value": "r.name"},
				{"col": "u.full_name", "as":"name", "value": "concat(u.first_name,' ',u.last_name) "}
			],
			"from": {
				"value": "users", "as": "u"
			},
			"join": [
				{"value": "roles", "as": "r", "type": "inner", "conn": "r.id = u.role_id"}
			],
			"where": {
				"and": [
					{"col":"id", "value":"u.id"},
					{"col":"email", "value":"u.email"},
					{"col":"phone", "value":"u.phone"},
					{"col":"name", "value":"LOWER(CONCAT(u.first_name, ' ', u.last_name))"},
					{"col":"active", "value":"u.active"},
					{"col":"-", "value":"u.deleted_at is null"}
				]
			}
		}
	`

	GetActiveUser = `
		{
			"select": [
				{"col": "u.*"},
				{"col": "r.name", "as": "role_name", "value": "r.name"},
				{"col": "p.profession", "as": "profession", "value": "p.profession"},
				{"col": "p.company", "as": "company", "value": "p.company"},
				{"col": "p.domicile", "as": "domicile", "value": "p.domicile"}
			],
			"from": {
				"value": "users", "as": "u"
			},
			"join": [
				{"value": "roles", "as": "r", "type": "inner", "conn": "r.id = u.role_id"},
				{"value": "profiles", "as": "p", "type": "left", "conn": "p.user_id = u.id"}
			],
			"where": {
				"and": [
					{"col":"id", "value":"u.id"},
					{"col":"email", "value":"u.email"},
					{"col":"-", "value":"u.active = 1"},
					{"col":"-", "value":"u.deleted_at is null"}
				]
			}
		}
	`

	UpdateUser = `
		{
			"set": [
				{"col": "u.*"},
				{"col": "-", "value":"updated_at = NOW()"},
				{"col": "-", "value":"updated_by = ?"}
			],
			"from": {
				"value": "users", "as": "u"
			},
			"where": {
				"and": [
					{"col":"id", "value":"u.id"},
					{"col":"email", "value":"u.email"},
					{"col":"active", "value":"u.active"},
					{"col":"-", "value":"u.deleted_at is null"}
				]
			}
		}
	`

	UpdateActiveUser = `
		{
			"set": [
				{"col": "u.*"},
				{"col": "-", "value":"updated_at = NOW()"},
				{"col": "-", "value":"updated_by = ?"}
			],
			"from": {
				"value": "users", "as": "u"
			},
			"where": {
				"and": [
					{"col":"id", "value":"u.id"},
					{"col":"email", "value":"u.email"},
					{"col":"-", "value":"u.active = 1"},
					{"col":"-", "value":"u.deleted_at is null"}
				]
			}
		}
	`

	CreateUser = `
		insert into users (first_name, last_name, email, phone, role_id, password, image_url, active, created_by, updated_by, created_at, updated_at)
		values(?,?,?,?,?,?,?,?,?,?, NOW(),NOW())
	`

	UpdateUserRoleByID = `
		update users set
			role_id = (select r.id from roles r where r.name = ? and r.deleted_at is null limit 1),
			updated_at = NOW(),
			updated_by = ?
		where id = ? and active = 1 and deleted_at is null
	`

	CreateProfile = `
		insert into profiles (user_id, profession, company, domicile, created_by, updated_by, created_at, updated_at)
		values(?,?,?,?,?,?, NOW(),NOW())
	`

	UpdateProfile = `
		{
			"set": [
				{"col": "p.*"},
				{"col": "-", "value":"updated_at = NOW()"},
				{"col": "-", "value":"updated_by = ?"}
			],
			"from": {
				"value": "profiles", "as": "p"
			},
			"where": {
				"and": [
					{"col":"user_id", "value":"p.user_id"},
					{"col":"-", "value":"p.deleted_at is null"}
				]
			}
		}
	`
	GetProfileByID = `
		{
			"select": [
				{"col": "p.*"}
			],
			"from": {
				"value": "profiles", "as": "p"
			},
			"where": {
				"and": [
					{"col":"user_id", "value":"p.user_id"},
					{"col":"-", "value":"p.deleted_at is null"}
				]
			}
		}
	`
)
