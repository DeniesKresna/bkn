package models

// role name
const (
	RoleNameExpert = "expert"
	RoleNameAdmin  = "administrator"
	RoleNameUser   = "user"
)

// expert status
const (
	ExpertStatusPending   = 0
	ExpertStatusAccepted  = 1
	ExpertStatusPublished = 2
	ExpertStatusRejected  = 3

	ExpertStatusStringPending   = "pending"
	ExpertStatusStringAccepted  = "accept"
	ExpertStatusStringRejected  = "reject"
	ExpertStatusStringPublished = "publish"
)

// expert services
const (
	ConsultationObject  = `consultation`
	InviteExpertObject  = `invitation`
	RecruitExpertObject = `recruitment`
	TrainingObject      = `training`
)

// course types
const (
	CourseHybrid    = "hybrid"
	CourseOnline    = "online"
	CourseOffline   = "offline"
	CourseRecording = "recording"
)

// course types code
const (
	HYBRID_CODE      = "HY"
	ONLINE_CODE      = "ON"
	OFFLINE_CODE     = "OF"
	RECORDING_CODING = "RC"
)

// course program
const (
	CourseJobhunAcademy                = "jobhun academy"
	CourseJobhunClass                  = "jobhun class"
	CourseStudiIndependenBersertifikat = "studi independen bersertifikat"
	CourseAirlanggaCareerPreparation   = "airlangga career preparation"
	CourseYesTech                      = "yes tech"
	CourseReadyToWork                  = "ready to work"
	CourseJobhunAcademyForCorporate    = "jobhun academy for corporate"
	CourseKartuPrakerja                = "kartu prakerja"
)

// course program code
const (
	JOBHUN_ACADEMY_CODE                 = "JA"
	JOBHUN_CLASS_CODE                   = "JC"
	STUDI_INDEPENDEN_BERSERTIFIKAT_CODE = "SIB"
	AIRLANGGA_CAREER_PREPARATION_CODE   = "ACP"
	YES_TECH_CODE                       = "YES"
	READY_TO_WORK_CODE                  = "RTW"
	JOBHUN_ACADEMY_FOR_CORPORATE_CODE   = "JAC"
	KARTU_PRAKERJA_CODE                 = "KP"
)

// product code
const (
	EXPERT_CODE = "EXP"
	COURSE_CODE = "CRS"
)

// product type
const (
	EXPERT_TYPE = "expert"
	COURSE_TYPE = "course"
)

type HttpDump struct {
	ReqDump string
	ResDump string
}
