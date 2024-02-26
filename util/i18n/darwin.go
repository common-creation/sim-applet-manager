//go:build darwin

package i18n

import "embed"

//go:generate cp -r "../../frontend/src/assets/i18n/" ./langs/
//go:embed langs
var langs embed.FS
