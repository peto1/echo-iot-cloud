package model

type Project struct {
	RecordMeta

	// ProjectId Generate by Snowflake, is unique in all the project
	ProjectId int64
	// ContractId contract with the supplier
	ContractId  string
	Description string
	LocationId  string
}
