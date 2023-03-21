package converter

type AppConverters struct {
	UserConverter
}

func NewAppConverters() *AppConverters {
	return &AppConverters{
		UserConverter: NewUserConverter(),
	}
}
