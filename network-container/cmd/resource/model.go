package resource

/*
This file is autogenerated, do not edit;
changes will be undone by the next 'generate' command.

Updates to this type are made my editing the schema file
and executing the 'generate' command
*/

// Model is autogenerated from the json schema
type Model struct {
	ProjectId      *string           `json:",omitempty"`
	RegionName     *string           `json:",omitempty"`
	Provisioned    *bool             `json:",omitempty"`
	ProviderName   *string           `json:",omitempty"`
	VpcId          *string           `json:",omitempty"`
	AtlasCidrBlock *string           `json:",omitempty"`
	Id             *string           `json:",omitempty"`
	ApiKeys        *ApiKeyDefinition `json:",omitempty"`
}

// ApiKeyDefinition is autogenerated from the json schema
type ApiKeyDefinition struct {
	PublicKey  *string `json:",omitempty"`
	PrivateKey *string `json:",omitempty"`
}
