package autoload

/*
	You can just read the .env file on import just by doing

		import _ "github.com/jl2501/godotenv/autoload"

	And bob's your mother's brother
*/

import (
	"github.com/jl2501/godotenv"
	"github.com/spf13/afero"
)

func init() {
	fs := afero.NewOsFs()
	godotenv.Load(fs)
}
