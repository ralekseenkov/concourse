package atc

type Space string

type Spaces struct {
	DefaultSpace Space
	AllSpaces    []Space
	Versions     []SpaceVersion
}

type DefaultSpace struct {
	DefaultSpace Space `json:"default_space"`
}

type SpaceVersion struct {
	Space    Space           `json:"space"`
	Version  Version         `json:"version"`
	Metadata []MetadataField `json:"metadata,omitempty"`
}
