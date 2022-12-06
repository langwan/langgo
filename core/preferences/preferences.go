package preferences

type Preferences interface {
	GetModule(name string) []byte
}
