//go:build windows

package i18n

import "embed"

//go:generate xcopy "..\\..\\frontend\\src\\assets\\i18n\\" .\\langs\\
//go:embed langs
var langs embed.FS
