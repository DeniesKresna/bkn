package queries

const (
	GetExpert = `
		{
			"select":[
				{"col":"e.*"},
				{"col":"u.*"}
			],
			"from":{
				"value": "experts", "as": "e"
			},
			"join": [
				{"value": "users", "as": "u", "type": "inner", "conn": "u.id = e.user_id"}
			],
			"where":{
				"and": [
					{"col":"id", "value":"e.id"},
					{"col":"userid", "value":"e.user_id"},
					{"col":"active", "value":"e.active"},
					{"col":"e.lowername", "value":"LOWER(JSON_VALUE(e.datas, '$.name' RETURNING CHAR(255)))"},
					{"col":"-", "value":"JSON_VALUE(e.datas, '$.deleted_at' RETURNING CHAR(30)) is null"}
				]
			}
		}
	`

	UpdateExpert = `
		{
			"set": [
				{"col": "e.*"}
			],
			"from": {
				"value": "experts", "as": "e"
			},
			"where": {
				"and": [
					{"col":"id", "value":"e.id"},
					{"col":"active", "value":"e.active"},
					{"col":"-", "value":"JSON_VALUE(e.datas, '$.deleted_at' RETURNING CHAR(30)) is null"}
				]
			}
		}
	`

	CreateExpertRow = `
		insert into experts (user_id) values (?)
	`

	UpdateExpertByID = "update experts set datas = ? where id = ?"

	GetExpertTempProfileByID = "select datas from experts where `datas`->>'$.id' = ? and `datas`->>'$.deleted_at' = 'null' and `datas`->>'$.active' = 0"

	UpdateActiveExpertById = "update experts set datas = JSON_SET(datas, '$.updated_at', DATE_FORMAT(NOW(), '%Y-%m-%dT%TZ'), '$.active', ?, '$.updated_by', ?) WHERE id = ?"

	DeleteExpert = "update experts set datas = JSON_SET(datas, '$.deleted_at', DATE_FORMAT(NOW(), '%Y-%m-%dT%TZ'), '$.deleted_by', ?) WHERE id = ?"

	GetUserIDByExpertID = "select user_id from experts where `datas`->>'$.id' = ? and `datas`->>'$.deleted_at' = 'null'"

	GetExpertAcceptedProfileByID = "select datas from experts where `datas`->>'$.id' = ? and `datas`->>'$.deleted_at' = 'null' and `datas`->>'$.active' = 1"
)
