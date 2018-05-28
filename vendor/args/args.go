package args

var (
	// Port ..
	Port = ":8080"
)

const (
	// SecretKey ..
	SecretKey = "My Secret"
	// this const arg should in args/args.go

	// AuthCodeSize ..
	AuthCodeSize = 6

	// AuthCodeLifeTime is the life time of auth code (minutes)
	// default to 30.
	AuthCodeLifeTime = 30
)
