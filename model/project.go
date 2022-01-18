package model

type Project struct {
	RecordMeta

	// Id Generate by Snowflake, is unique in all the project
	// ContractId contract with the supplier
	ContractId  string
	Description string
	LocationId  string
}
