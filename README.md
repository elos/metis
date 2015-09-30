metis [![GoDoc](https://godoc.org/github.com/elos/metis?status.svg)](https://godoc.org/github.com/elos/metis)
-----
Package metis provides logical structures for modeling persisted data structures and their relationships. It also provides formal verification of the validity of a relational object schema. For an extensive discussion of the theory behind metis, please see the ["Metis Data Model"](https://github.com/elos/documentation/blob/master/guide/2%20-%20Data%20Model.md).

### Architecture

All model definitions are valid JSON. Models can have traits (string primitive tuples) and relations (links to other models). A group of models collectively and implicitly defines a schema. Metis has basic IO parsing, which reads in model definition files and generates metis.Model structures. Then you can `BuildSchema(models)` to get a schema from these definitions. As metis first serializes the JSON to go structs (ModelDef, TraitDef, LinkDef, and SchemaDef) in uses this stage to warn of any errors (invalid relation references, bad codomains). Therefore any metis.Model produced by the package is guaranteed to be logically sound. It may not, however, define a valid schema when combined with other models. You must check the validity of a schema after building it, with `schema.Valid()`.

### Overview of Logical Components

#### Trait
A trait is a string primitive tuple. It has a name (the string component), which is used to refer it's value (the primitive type)

#### Relation
[Theoretically](https://github.com/elos/documentation/blob/master/guide/2%20-%20Data%20Model.md), a relation can be represented as a Trait. But we deal with a relation as a more abstract logical construct in metis, to allow for the generation of code which creates the convenience algorithms to retrieve model relations based on a metis defintion. Therefore a relation has a name, a multiplicity (one or mul), a singular form (iff multiplicity = mul), a codomain (the space of models which can be assigned to this relation), and an inverse (the name of the corresponding model's link to this model).

#### Model
A model has a kind (it's name), a space (it's plural - but also its physical domain), domains (the physcial and abstract spaces this model implements -- think of type system interfaces), and a list of traits and a list of relations.

note to contributors: when updating README, also update go comments
