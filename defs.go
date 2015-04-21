package metis

// Definition for a trait
// { ...
//		"trait_name": {
//			"type": "datetime"
//		}
// ... }
type TraitDef struct {
	Type string `json:"type"`
}

// Definition for a relationship
// { ...
//		"relationship_name": {
//			"other": "event",
//			"type": "mul",
//		}
// ... }
type RelationshipDef struct {
	Kind  string `json:"kind"`
	Other string `json:"other"`
}

// Definition for a model
// {
//		"kind": "user", "plural": "users", "version": "1.0.0",
//		"relationships": { ... },
//		"traits": { ... }
// }
type ModelDef struct {
	Kind    string `json:"kind"`
	Plural  string `json:"plural"`
	Version string `json:"version"`

	Traits        map[string]TraitDef
	Relationships map[string]RelationshipDef
}
