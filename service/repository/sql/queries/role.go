package queries

const (
	GetRole = `{
		"select": [
			{"col": "r.*"}
		],
		"from": {
			"value": "roles", "as": "r"
		},
		"where": {
			"and": [
				{"col":"id", "value":"r.id"},
				{"col":"name", "value":"r.name"},
				{"col":"-", "value":"r.deleted_at is null"}
			]
		}
	}`

	CreateRole = `
		insert into roles (name, created_by, created_at, updated_by, updated_at)
		values (?, ?, NOW(), ?, NOW())
	`
)
