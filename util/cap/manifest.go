package cap

import "strings"

func ManifestToMap(manifest string) map[string]string {
	elements := make(map[string]string)

	var k, v string
	for _, line := range strings.Split(manifest, "\n") {
		if strings.HasPrefix(line, " ") {
			v += strings.TrimLeft(line, " ")
		} else if pos := strings.Index(line, ":"); pos != -1 {
			if k != "" {
				elements[k] = v
			}
			k = strings.TrimSpace(line[:pos])
			v = strings.TrimSpace(line[pos+1:])
		}
	}

	if k != "" {
		elements[k] = v
	}

	return elements
}
