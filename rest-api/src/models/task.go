package models

//Create tasks structure
type Tasks struct{
  ID string `json:"id,omitempty"`
  Name  string `json:"taskName,omitempty"`
  Content string `json:"taskContent,omitempty"`
  State string `json:"taskState,omitempty"`
}
