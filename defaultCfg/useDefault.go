package defaultCfg

var (
	useDefault = false
)

func SetUseDefaultConfig(flag bool) {
	useDefault = flag
}

func IsUseDefaultConfig() bool {
	return useDefault
}
