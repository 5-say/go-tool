package change

import (
	"fmt"
)

func ToString[FromType Simple | SimplePointor](data FromType) string {
	return fmt.Sprint(data)
}
