// Code generated by goctl. DO NOT EDIT.
package types

type StudentGpaItem struct {
	StudentId string `json:"student_id"`
	Gpa       string `json:"gpa"`
}

type StudentGpaResponse struct {
	Results []StudentGpaItem `json:"results"`
}
