package queries

const (
	GetCourse = `
		{
			"select":[
				{"col":"c.*"},
				{"col":"c.count", "as": "course_count", "value": "count(*)"},
				{
					"col": "c.status", "as":"status",
					"value": "
						CASE
							WHEN CURRENT_DATE < c.start_time AND CURRENT_DATE < c.end_time THEN 'Belum dimulai'
							WHEN CURRENT_DATE >= c.start_time AND CURRENT_DATE <= c.end_time THEN 'Sedang berjalan'
							WHEN CURRENT_DATE > c.start_time AND CURRENT_DATE > c.end_time THEN 'Selesai'
							ELSE 'Tidak terdefinisi'
						END
					"
				}
			],
			"from":{
				"value": "courses", "as": "c"
			},
			"where":{
				"and": [
					{"col":"id", "value":"c.id"},
					{"col":"c.code", "value":"c.code"},
					{"col":"c.name", "value":"c.name"},
					{"col":"c.active", "value":"c.active"},
					{"col":"c.sector", "value":"JSON_VALUE(c.sectors, '$.sectors')"},
					{"col":"-", "value":"c.deleted_at is null"}
				]
			}
		}
	`

	CreateCourse = `
		insert into courses (name, program, code, type, facility, sectors, start_time, end_time, price, detail, expert_id, active, created_by, updated_by)
		values (?,?,?,?,?,?,?,?,?,?,?,2,?,?)
	`

	UpdateCourse = `
		update courses set name = ?, facility = ?, sectors = ?, start_time = ?, end_time = ?, detail = ?, expert_id = ?, active = ? ,updated_at = NOW(), updated_by = ?
		where id = ?
	`

	UpdateCourseImageURL = `
		update courses set image_url = ?, updated_at = NOW(), updated_by = ?
		where id = ?
	`

	GetCourseWithExpertByID = `
		select
			c.*,
			CASE
				WHEN CURRENT_DATE < c.start_time AND CURRENT_DATE < c.end_time THEN 'Belum dimulai'
				WHEN CURRENT_DATE >= c.start_time AND CURRENT_DATE <= c.end_time THEN 'Sedang berjalan'
				WHEN CURRENT_DATE > c.start_time AND CURRENT_DATE > c.end_time THEN 'Selesai'
				ELSE 'Tidak terdefinisi'
			END as status,
			e.expert_name,
			e.expert_image,
			e.expert_experiences
			from courses c
		left join (
			select 
				id,
				datas->>"$.name" as expert_name,
				datas->>"$.image" as expert_image,
				datas->>"$.experiences" as expert_experiences
			from experts
			where datas->>"$.deleted_at" = "null" and datas->>"$.active" in (1,2)
		) e on e.id = c.expert_id
		where c.deleted_at is null and c.id = ?
	`

	CreateCourseUser = `
		insert into course_user (
			course_id, user_id, course_rate, course_suggestion, certificate_link, payment_id, created_by, updated_by
		) values (?,?,?,?,?,?,?,?)
	`

	GetCourseUser = `
		{
			"select": [
				{"col": "cu.count", "as": "course_user_count", "value": "count(*)"},				
				{"col": "cu.*"}
			],
			"from": {
				"value": "course_user", "as": "cu"
			},
			"where": {
				"and": [
					{"col":"id", "value":"cu.id"},
					{"col":"course_id", "value":"cu.code"},
					{"col":"user_id", "value":"cu.base_code"},
					{"col":"payment_id", "value":"cu_payment_id"},
					{"col":"-", "value":"cu.deleted_at is null"}
				]
			}
		}
	`

	GetCourseDataIgnoreActiveGetByID = `
		select c.* from courses c where c.deleted_at is null and c.id = ?
	`

	GetCourseCodePrevByProgramAndType = `
		select code from courses where program = ? and type = ? order by created_at desc limit 1
	`

	GetCourseDashboardRaw = `
		select c.id, c.name, c.image_url,
		CASE
		WHEN CURRENT_DATE < c.start_time AND CURRENT_DATE < c.end_time THEN 'Belum dimulai'
		WHEN CURRENT_DATE >= c.start_time AND CURRENT_DATE <= c.end_time THEN 'Sedang berjalan'
		WHEN CURRENT_DATE > c.start_time AND CURRENT_DATE > c.end_time THEN 'Selesai'
		ELSE 'Tidak terdefinisi'
		END as status,
		count(cu.course_id) as total_register,
		count(cui.course_id) as total_interest
		from courses c
		left join course_user cu on c.id = cu.course_id
		left join course_user_interest cui on c.id = cui.course_id and cui.interest = 1
	`

	GetCountCourseRaw = `
		select count(*)
		from courses c
	`

	CreateCourseInterest = `
		insert into course_user_interest (course_id, user_id, interest, created_at)
		values (?,?,1,NOW())
	`

	GetCourseInterestByID = `
		select * from course_user_interest where id = ?
	`

	GetCourseInterestByCourseIDAndUserID = `
		select * from course_user_interest where course_id = ? and user_id = ?
	`

	UpdateCourseInterest = `
		update course_user_interest set interest = ?
		where id = ?
	`

	GetCourseTablePublishedRaw = `
		select 
			c.image_url as image_url, c.type as type, c.start_time as start_time, c.price as price, c.end_time as end_time, c.id as id, c.name as name, c.sectors as sectors
		from courses c
	`

	GetCourseTablePublishedWithInterestRaw = `
		select 
			c.image_url as image_url, c.type as type, c.start_time as start_time, c.price as price, c.end_time as end_time, c.id as id, c.name as name, c.sectors as sectors, (select interest from course_user_interest cui where cui.course_id = c.id and cui.user_id = ?) as interest
		from courses c
	`

	GetCourseUserRegister = `
		{
			"select": [
				{"col":"cu.count", "as": "count", "value": "count(*)"},
				{"col": "cu.*"},
				{"col": "u.user_name", "as":"user_name", "value": "concat(u.first_name,' ',u.last_name)"},
				{"col": "u.*"}
			],
			"from": {
				"value": "course_user", "as": "cu"
			},
			"join": [
				{"value": "users", "as": "u", "type": "inner", "conn": "u.id = cu.user_id"}
			],
			"where": {
				"and": [
					{"col":"user_name", "value":"LOWER(CONCAT(u.first_name, ' ', u.last_name))"},
					{"col":"course_id", "value":"cu.course_id"},
					{"col":"-", "value":"cu.deleted_at is null"}
				]
			}
		}
	`

	GetCourseUserInterest = `
		{
			"select": [
				{"col":"cui.count", "as": "count", "value": "count(*)"},
				{"col": "cui.*"},
				{"col": "u.user_name", "as":"user_name", "value": "concat(u.first_name,' ',u.last_name)"},
				{"col": "u.*"}
			],
			"from": {
				"value": "course_user_interest", "as": "cui"
			},
			"join": [
				{"value": "users", "as": "u", "type": "inner", "conn": "u.id = cui.user_id"}
			],
			"where": {
				"and": [
					{"col":"user_name", "value":"LOWER(CONCAT(u.first_name, ' ', u.last_name))"},
					{"col":"course_id", "value":"cui.course_id"},
					{"col":"-", "value":"cui.interest = 1"}
				]
			}
		}
	`
	GetCourseInterestWithExpertByID = `
		select
			c.*,
			CASE
				WHEN CURRENT_DATE < c.start_time AND CURRENT_DATE < c.end_time THEN 'Belum dimulai'
				WHEN CURRENT_DATE >= c.start_time AND CURRENT_DATE <= c.end_time THEN 'Sedang berjalan'
				WHEN CURRENT_DATE > c.start_time AND CURRENT_DATE > c.end_time THEN 'Selesai'
				ELSE 'Tidak terdefinisi'
			END as status,
			e.expert_name,
			e.expert_image,
			e.expert_experiences,
			cui.interest
			from courses c
		left join (
			select 
				id,
				datas->>"$.name" as expert_name,
				datas->>"$.image" as expert_image,
				datas->>"$.experiences" as expert_experiences
			from experts
			where datas->>"$.deleted_at" = "null" and datas->>"$.active" in (1,2)
		) e on e.id = c.expert_id
		left join course_user_interest cui on c.id = cui.course_id and cui.user_id = ?
		where c.deleted_at is null and c.active = 2 and c.id = ?
	`
)
