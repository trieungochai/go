// we will look at what it used to be like when working with projects that had multiple Go modules that needed their dependencies to be replaced so that they could use local changes.
// We will then update the example code so that it uses Go workspaces to show the improvements
package printer

import (
	"fmt"

	"github.com/google/uuid"
)

func PrintNewUUID() string {
	id := uuid.New()
	return fmt.Sprintf("Generated UUID: %s\n", id)
}
