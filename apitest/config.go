package apitest

var (
	Host = "localhost"
	Port = 9001

	C = &CommonParams{
		TeamUUID: "K4Cthi6p",
		OrgUUID:  "AytnNdkN",
		UserUUID: "JhWrCvGN",
		Token:    "j3iHMEvhse6ytPjrR8dxmHP49mFaYZMiPhRAKNLB6bRRpSYbsd8Ah4ixZEaZLVkF",
	}
)

type CommonParams struct {
	TeamUUID    string
	UserUUID    string
	Token       string
	OrgUUID     string
	ProjectUUID string
}
