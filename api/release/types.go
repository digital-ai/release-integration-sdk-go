package release

type Variable struct {
	FolderId           *string                     `json:"folderId,omitempty"`
	Title              *string                     `json:"title,omitempty"`
	Key                *string                     `json:"key,omitempty"`
	RequiresValue      *bool                       `json:"requiresValue,omitempty"`
	ShowOnReleaseStart *bool                       `json:"showOnReleaseStart,omitempty"`
	Label              *string                     `json:"label,omitempty"`
	Description        *string                     `json:"description,omitempty"`
	ValueProvider      *ValueProviderConfiguration `json:"valueProvider,omitempty"`
	Inherited          *bool                       `json:"inherited,omitempty"`
	Value              map[string]interface{}      `json:"value,omitempty"`
	EmptyValue         map[string]interface{}      `json:"emptyValue,omitempty"`
	ValueEmpty         *bool                       `json:"valueEmpty,omitempty"`
	UntypedValue       map[string]interface{}      `json:"untypedValue,omitempty"`
	Password           *bool                       `json:"password,omitempty"`
	ValueAsString      *string                     `json:"valueAsString,omitempty"`
	EmptyValueAsString *string                     `json:"emptyValueAsString,omitempty"`
}

type Variable1 struct {
	Id                    *string                     `json:"id,omitempty"`
	Key                   *string                     `json:"key,omitempty"`
	Type                  *string                     `json:"type,omitempty"`
	RequiresValue         *bool                       `json:"requiresValue,omitempty"`
	ShowOnReleaseStart    *bool                       `json:"showOnReleaseStart,omitempty"`
	Value                 map[string]interface{}      `json:"value,omitempty"`
	Label                 *string                     `json:"label,omitempty"`
	Description           *string                     `json:"description,omitempty"`
	Multiline             *bool                       `json:"multiline,omitempty"`
	Inherited             *bool                       `json:"inherited,omitempty"`
	PreventInterpolation  *bool                       `json:"preventInterpolation,omitempty"`
	ExternalVariableValue *ExternalVariableValue      `json:"externalVariableValue,omitempty"`
	ValueProvider         *ValueProviderConfiguration `json:"valueProvider,omitempty"`
}

type ExternalVariableValue struct {
	Server      *string `json:"server,omitempty"`
	ServerType  *string `json:"serverType,omitempty"`
	Path        *string `json:"path,omitempty"`
	ExternalKey *string `json:"externalKey,omitempty"`
}

type ValueProviderConfiguration struct {
	Variable *Variable `json:"variable,omitempty"`
}

type VariableOrValue struct {
	Variable *string                `json:"variable,omitempty"`
	Value    map[string]interface{} `json:"value,omitempty"`
}
