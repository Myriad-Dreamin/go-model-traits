package modeltraits

type DatabaseTraitsInterface interface {
	Migrate() error
}
