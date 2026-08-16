package engram

import (
	"github.com/LimiteRoche/kto-ai/v2/internal/installcmd"
	"github.com/LimiteRoche/kto-ai/v2/internal/model"
	"github.com/LimiteRoche/kto-ai/v2/internal/system"
)

func InstallCommand(profile system.PlatformProfile) ([][]string, error) {
	return installcmd.NewResolver().ResolveComponentInstall(profile, model.ComponentEngram)
}
