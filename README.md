metis [![GoDoc](https://godoc.org/github.com/elos/metis?status.svg)](https://godoc.org/github.com/elos/metis)
-----
Package metis provides logical structures for modeling persisted data structures and their relationships.

### Architecture
The overall architecture of metis is quite simple. All model definitions are valid JSON. Models can have traits (attributes) and links (relationships). A group of models collectively and implicitly defines a schema. Metis has basic io parsing, which reads in model definition files and then you can use `BuildSchema(models)` to get a schema from these definitions. Metis first serializes the JSON to go structs (ModelDef, TraitDef, LinkDef, and SchemaDef), at this stage in warns of any errors (invalid link references, bad codomains etc).

### Components
#### Trait
A trait has a name, which you use to refer to the attribute and a metis primitive type (e.g. string, int, boolean)
#### Link
A link has a name, a multiplicity (one or mul), a singular form (iff multiplicity = mul), a codomain (the space of models which can be assigned to this link), and an inverse (the name of the corresponding model's link to this model).
#### Model
A model has a kind (it's name), a space (it's plural - but also its physical domain), domains (the physcial and abstract spaces this model implements), and a list of traits and a list of links.
#### Schema
A schema has a list of models, a list of spaces and a list of domains. Plus a version. A virtual domain can be thought of as an interface and is any domain referenced in a model that has no physical space counterpart.
