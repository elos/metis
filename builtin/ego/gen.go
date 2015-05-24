package ego

import (
	"fmt"
	"path/filepath"

	"github.com/elos/metis"
)

func Generate(s *metis.Schema, root string) error {
	if err := s.Valid(); err != nil {
		return fmt.Errorf("schema invalid: %s", err)
	}

	for kind, model := range s.Models {
		MakeGo(model).WriteFile(filepath.Join(root, kind+".go"))
	}

	WriteKindsFile(s, filepath.Join(root, "kinds.go"))
	WriteDynamicFile(s, filepath.Join(root, "dynamic.go"))
	WriteDBsFile(s, filepath.Join(root, "dbs.go"))

	return nil
}
