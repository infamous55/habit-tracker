package graphql

import (
	"github.com/infamous55/habit-tracker/internal/mongodb"
)

type Resolver struct {
	Database mongodb.DatabaseWrapper
}
