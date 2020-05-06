package apitest

var (
	Host = "localhost"
	Port = 9001

	C = &CommonParams{
		TeamUUID: "K4Cthi6p",
		OrgUUID:  "AytnNdkN",
		UserUUID: "JhWrCvGN",
		Token:    "Wqmz4hoypb54fQhmscNaxcZN1ACpYqzC7qbirP64Y4Ihw2wgRrPTF8YgtCddPbhd",
	}
)

type CommonParams struct {
	TeamUUID    string
	UserUUID    string
	Token       string
	OrgUUID     string
	ProjectUUID string
}
