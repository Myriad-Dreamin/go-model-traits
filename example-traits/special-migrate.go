package mytraits

type ObjectModelTraits struct {
	ModelTraits
}

func (traits *ObjectModelTraits) Migrate() error {
	err := traits.ModelTraits.Migrate()
	if err != nil {
		return err
	}
	// do something you want
	return nil
}
