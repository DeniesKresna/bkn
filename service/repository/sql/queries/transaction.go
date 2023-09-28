package queries

const (
	GetDealByServiceAndId = `select deal from order_detail where id = ? and JSON_CONTAINS(deal->'$.service', ?)`

	CreateProposal = `
		insert into order_detail (proposal, deal) values (?, ?)
	`

	CreateOrder = `
		insert into orders (
			user_id, 
			expert_id, 
			order_detail_id, 
			price, rest_price, 
			is_paid, is_paid_off, 
			no_invoice,
			invoice_url,
			is_started, 
			is_finished, 
			expert_rate, 
			expert_suggestion, 
			user_rate,
			user_suggestion,
			updated_at, updated_by, created_at, created_by)
		values
			(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,NOW(), ?,NOW(), ?)
	`

	GetDeal = `
		{
			"select": [
				{"col": "d.*"}
			],
			"from": {
				"value": "order_detail", "as": "d"
			},
			"where": {
				"and": [
					{"col":"id", "value":"d.id"},
					{"col":"dealConsultationStartDate", "value":"d.consultation_start_date"},
					{"col":"dealInviteStartDate", "value":"d.invite_start_date"},
					{"col":"dealTrainingStartDate", "value":"d.training_start_date"},
					{"col":"dealRecruitEmail", "value":"d.recruit_email"},
					{"col":"-", "value":"JSON_VALUE(d.deal, '$.service.deleted_at' RETURNING CHAR(30)) is null"}
				]
			}
		}
	`

	GetOrder = `
		{
			"select": [
				{"col":"o.count", "as": "order_count", "value": "count(*)"},
				{"col": "d.*"},
				{"col": "o.*"},
				{
					"col": "o.finishStatus", "as":"status",
					"value": "
						CASE
							WHEN o.is_paid = 1 AND o.is_paid_off = 1 AND o.is_finished = 1 THEN 'Selesai, Lunas'
							WHEN o.is_paid = 1 AND o.is_paid_off = 0 AND o.is_finished = 1 THEN 'Selesai, Belum lunas'
							WHEN o.is_paid = 0 AND o.is_paid_off = 0 AND o.is_finished = 1 THEN 'Selesai, Belum dibayar'
							WHEN o.is_paid = 1 AND o.is_paid_off = 1 AND o.is_started = 1 THEN 'Sedang berjalan, Lunas'
							WHEN o.is_paid = 1 AND o.is_paid_off = 0 AND o.is_started = 1 THEN 'Sedang berjalan, Belum lunas'
							WHEN o.is_paid = 0 AND o.is_paid_off = 0 AND o.is_started = 1 THEN 'Sedang berjalan, Belum dibayar'
							WHEN o.is_paid = 1 AND o.is_paid_off = 1 AND o.is_started = 0 THEN 'Belum dimulai, Lunas'
							WHEN o.is_paid = 1 AND o.is_paid_off = 0 AND o.is_started = 0 THEN 'Belum dimulai, Belum lunas'
							WHEN o.is_paid = 0 AND o.is_paid_off = 0 AND o.is_started = 0 THEN 'Belum dimulai, Belum dibayar'
							ELSE 'Tidak terdefinisi'
						END
					"
				},
				{
					"col": "o.serviceName", "as":"service",
					"value": "
						CASE
							WHEN d.deal->>'$.service.recruit_expert.email' != '' THEN 'recruitment'
							WHEN d.deal->>'$.service.training.start_time' != 'null' THEN 'training'
							WHEN d.deal->>'$.service.invite_expert.start_date' != 'null' THEN 'invitation'
							WHEN d.deal->>'$.service.consultation.start_date' != 'null' THEN 'consultation'
							ELSE 'Tidak terdefinisi'
						END
					"
				}
			],
			"from": {
				"value": "orders", "as": "o"
			},
			"join": [
				{"value": "order_detail", "as": "d", "type": "inner", "conn": "d.id = o.order_detail_id"}
			],
			"where": {
				"and": [
					{"col":"id", "value":"o.id"},
					{"col":"dealConsultationStartDate", "value":"d.consultation_start_date"},
					{"col":"dealInviteStartDate", "value":"d.invite_start_date"},
					{"col":"dealTrainingStartTime", "value":"d.training_start_time"},
					{"col":"dealRecruitEmail", "value":"d.recruit_email"},
					{"col":"userID", "value":"o.user_id"},
					{"col":"-", "value":"JSON_VALUE(d.deal, '$.service.deleted_at' RETURNING CHAR(30)) is null"},
					{"col":"-", "value":"o.deleted_at is null"}
				]
			}
		}
	`

	GetOrderTable = `
		{
			"select": [
				{"col":"o.count", "as": "order_count", "value": "count(*)"},
				{"col": "d.*"},
				{"col": "u.full_name", "as":"user_name", "value": "concat(u.first_name,' ',u.last_name)"},
				{"col": "e.*"},
				{
					"col": "o.finishStatus", "as":"status",
					"value": "
						CASE
							WHEN o.is_paid = 1 AND o.is_paid_off = 1 AND o.is_finished = 1 THEN 'Selesai, Lunas'
							WHEN o.is_paid = 1 AND o.is_paid_off = 0 AND o.is_finished = 1 THEN 'Selesai, Belum lunas'
							WHEN o.is_paid = 0 AND o.is_paid_off = 0 AND o.is_finished = 1 THEN 'Selesai, Belum dibayar'
							WHEN o.is_paid = 1 AND o.is_paid_off = 1 AND o.is_started = 1 THEN 'Sedang berjalan, Lunas'
							WHEN o.is_paid = 1 AND o.is_paid_off = 0 AND o.is_started = 1 THEN 'Sedang berjalan, Belum lunas'
							WHEN o.is_paid = 0 AND o.is_paid_off = 0 AND o.is_started = 1 THEN 'Sedang berjalan, Belum dibayar'
							WHEN o.is_paid = 1 AND o.is_paid_off = 1 AND o.is_started = 0 THEN 'Belum dimulai, Lunas'
							WHEN o.is_paid = 1 AND o.is_paid_off = 0 AND o.is_started = 0 THEN 'Belum dimulai, Belum lunas'
							WHEN o.is_paid = 0 AND o.is_paid_off = 0 AND o.is_started = 0 THEN 'Belum dimulai, Belum dibayar'
							ELSE 'Tidak terdefinisi'
						END
					"
				},
				{"col": "o.*"}
			],
			"from": {
				"value": "orders", "as": "o"
			},
			"join": [
				{"value": "order_detail", "as": "d", "type": "inner", "conn": "d.id = o.order_detail_id"},
				{"value": "users", "as": "u", "type": "inner", "conn": "u.id = o.user_id"},
				{"value": "experts", "as": "e", "type": "inner", "conn": "e.id = o.expert_id"}
			],
			"where": {
				"and": [
					{"col":"dealConsultationStartDate", "value":"d.consultation_start_date"},
					{"col":"dealInviteStartDate", "value":"d.invite_start_date"},
					{"col":"dealTrainingStartTime", "value":"d.training_start_time"},
					{"col":"dealRecruitEmail", "value":"d.recruit_email"},
					{"col":"first_name", "value":"u.first_name"},
					{"col":"last_name", "value":"u.last_name"},
					{"col":"e.lowername", "value":"LOWER(JSON_VALUE(e.datas, '$.name' RETURNING CHAR(255)))"},
					{"col":"userID", "value":"o.user_id"},
					{"col":"expertID", "value":"o.expert_id"},
					{"col":"-", "value":"o.deleted_at is null"},
					{"col":"-", "value":"JSON_VALUE(d.deal, '$.service.deleted_at' RETURNING CHAR(30)) is null"}
				]
			}
		}
	`
	CreateExpertRequestment = `
		insert into service_request (
			user_id, 
			service, 
			description, 
			updated_at, updated_by, created_at, created_by)
		values
			(?,?,?,NOW(), ?,NOW(), ?)
	`

	GetServiceRequestDetail = `
		{
			"select": [
				{"col":"s.count", "as": "service_request_count", "value": "count(*)"},
				{"col": "s.*"}
			],
			"from": {
				"value": "service_request", "as": "s"
			},
			"where": {
				"and": [
					{"col":"id", "value":"s.id"},
					{"col":"-", "value":"s.deleted_at is null"}
				]
			}
		}
	`

	GetServiceRequestTable = `
		{
			"select": [
				{"col": "s.count", "as": "service_request_count", "value": "count(*)"},				
				{"col": "u.full_name", "as":"user_name", "value": "concat(u.first_name,' ',u.last_name)"},
				{"col": "s.*"}
			],
			"from": {
				"value": "service_request", "as": "s"
			},
			"join": [
				{"value": "users", "as": "u", "type": "inner", "conn": "u.id = s.user_id"}
			],
			"where": {
				"and": [
					{"col":"id", "value":"s.id"},
					{"col":"-", "value":"s.deleted_at is null"},
					{"col":"first_name", "value":"u.first_name"},
					{"col":"last_name", "value":"u.last_name"}
				]
			}
		}
	`
	UpdateOrder = `
		{
			"set": [
				{"col": "o.*"},
				{"col": "-", "value":"updated_at = NOW()"},
				{"col": "-", "value":"updated_by = ?"}
			],
			"from": {
				"value": "orders", "as": "o"
			},
			"where": {
				"and": [
					{"col":"id", "value":"o.id"},
					{"col":"-", "value":"o.deleted_at is null"}
				]
			}
		}
	`
	UpdateDeal = `
		{
			"set": [
				{"col": "d.*"}
			],
			"from": {
				"value": "order_detail", "as": "d"
			},
			"where": {
				"and": [
					{"col":"id", "value":"d.id"},
					{"col":"-", "value":"JSON_VALUE(d.deal, '$.service.deleted_at' RETURNING CHAR(30)) is null"}
				]
			}
		}
	`

	GetOrderTableRow = `
		select concat(u.first_name,' ',u.last_name) as user_name, e.datas->>"$.name" as name, 
		CASE
		WHEN o.is_paid = 1 AND o.is_paid_off = 1 AND o.is_finished = 1 THEN 'Selesai, Lunas'
		WHEN o.is_paid = 1 AND o.is_paid_off = 0 AND o.is_finished = 1 THEN 'Selesai, Belum lunas'
		WHEN o.is_paid = 0 AND o.is_paid_off = 0 AND o.is_finished = 1 THEN 'Selesai, Belum dibayar'
		WHEN o.is_paid = 1 AND o.is_paid_off = 1 AND o.is_started = 1 THEN 'Sedang berjalan, Lunas'
		WHEN o.is_paid = 1 AND o.is_paid_off = 0 AND o.is_started = 1 THEN 'Sedang berjalan, Belum lunas'
		WHEN o.is_paid = 0 AND o.is_paid_off = 0 AND o.is_started = 1 THEN 'Sedang berjalan, Belum dibayar'
		WHEN o.is_paid = 1 AND o.is_paid_off = 1 AND o.is_started = 0 THEN 'Belum dimulai, Lunas'
		WHEN o.is_paid = 1 AND o.is_paid_off = 0 AND o.is_started = 0 THEN 'Belum dimulai, Belum lunas'
		WHEN o.is_paid = 0 AND o.is_paid_off = 0 AND o.is_started = 0 THEN 'Belum dimulai, Belum dibayar'
		ELSE 'Tidak terdefinisi'
		END  as status, 
		CASE
		WHEN JSON_VALUE(d.deal, '$.service.consultation.start_date' RETURNING CHAR(225)) not in ( "null", "" ) THEN 'consultation'
		WHEN JSON_VALUE(d.deal, '$.service.invite_expert.start_date' RETURNING CHAR(225)) not in ( "null", "" ) THEN 'invitation'
		WHEN JSON_VALUE(d.deal, '$.service.recruit_expert.email' RETURNING CHAR(225)) not in ( "null", "" ) THEN 'recruitment'
		WHEN JSON_VALUE(d.deal, '$.service.training.start_time' RETURNING CHAR(225)) not in ( "null", "" ) THEN 'training'
		ELSE 'Tidak terdefinisi'
		END  as service, 
		o.id as id, o.expert_id as expert_id, o.created_at as created_at, o.user_id as user_id from orders o inner join order_detail d on d.id = o.order_detail_id inner join users u on u.id = o.user_id inner join experts e on e.id = o.expert_id
	`

	GetCountOrderTableRow = "select count(*) from orders o inner join order_detail d on d.id = o.order_detail_id inner join users u on u.id = o.user_id inner join experts e on e.id = o.expert_id"

	GetServiceRequestTableRow = `
	select concat(u.first_name,' ',u.last_name) as user_name, s.user_id as user_id, s.service as service, s.created_at as 
	created_at, s.id as id from service_request s inner join users u on u.id = s.user_id
	`
	GetCountServiceRequestTableRow = "select count(*) as service_request_count from service_request s inner join users u on u.id = s.user_id"

	// ######################
	// Payments
	// ######################

	GetPayment = `
		{
			"select": [
				{"col": "p.count", "as": "payment_count", "value": "count(*)"},				
				{"col": "p.*"}
			],
			"from": {
				"value": "payments", "as": "p"
			},
			"where": {
				"and": [
					{"col":"id", "value":"p.id"},
					{"col":"code", "value":"p.code"},
					{"col":"base_code", "value":"p.base_code"},
					{"col":"success", "value":"p.success"},
					{"col":"invoice_at", "value":"p.invoice_request_at"},
					{"col":"expired_at", "value":"p.invoice_expired_at"},
					{"col":"-", "value":"p.deleted_at is null"}
				]
			}
		}
	`

	UpdatePayment = `
		{
			"set": [
				{"col": "p.*"},
				{"col": "-", "value":"updated_at = NOW()"},
				{"col": "-", "value":"updated_by = ?"}
			],
			"from": {
				"value": "payments", "as": "p"
			},
			"where": {
				"and": [
					{"col":"id", "value":"p.id"},
					{"col":"code", "value":"p.code"},
					{"col":"-", "value":"p.deleted_at is null"}
				]
			}
		}
	`

	CreatePayment = `insert into payments(base_code, code, product_type, amount, gateway_vendor, invoice_request_url, invoice_request_data, invoice_request_at, invoice_expired_at,invoice_response, callback_url,callback_data, callback_at, success, created_by, updated_by)
		values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)
	`

	CreateInvoice = `insert into payments(base_code, code, product_type,amount, gateway_vendor, invoice_request_url, created_by,updated_by)
		values(?,?,?,?,?,?,?,?)
	`
)
