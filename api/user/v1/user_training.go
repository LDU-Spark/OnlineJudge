package v1

import "github.com/gogf/gf/v2/frame/g"

type TrainingReq struct {
	g.Meta `path:"/training/{username}" method:"get"`
}
type TrainingRes struct {
	Submitted int `json:"submitted"`
	Rating    int `json:"rating"`
	Solved    int `json:"solved"`
}
