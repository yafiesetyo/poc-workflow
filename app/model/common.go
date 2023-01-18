package model

var (
	TypeFirstEndpoint  = "first_endpoint"
	TypeSecondEndpoint = "second_endpoint"
	TypeThirdEndpoint  = "third_endpoint"
)

type (
	// for pub/sub and cloud function
	CommonPublishReq struct {
		ID       uint64 `json:"id"`
		RuleName string `json:"rule_name"`
		Type     string `json:"type"`
	}

	// create ktp request
	CommonCreateReq struct {
		Name string `json:"name"`
	}
)
