package routing

import (
	"strconv"
	"strings"
)

func extractIdFromPath(path string, commonPath string) (uint32, error) {
	idRaw := strings.TrimSuffix(path[len(commonPath):], "/")

	id64, err := strconv.ParseUint(idRaw, 10, 32)

	if err != nil {
		return 0, err
	}

	return uint32(id64), nil
}
