package config

import (
	"io"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/afero"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Reader interface {
	Common
	WatchConfig()
	OnConfigChange(run func(in fsnotify.Event))
}

type Common interface {
	AllKeys() []string
	AllSettings() map[string]interface{}
	SetFs(fs afero.Fs)
	ConfigFileUsed() string
	SetTypeByDefaultValue(enable bool)
	Get(key string) interface{}
	Sub(key string) *viper.Viper
	GetString(key string) string
	GetBool(key string) bool
	GetInt(key string) int
	GetInt32(key string) int32
	GetInt64(key string) int64
	GetUint(key string) uint
	GetUint32(key string) uint32
	GetUint64(key string) uint64
	GetFloat64(key string) float64
	GetTime(key string) time.Time
	GetDuration(key string) time.Duration
	GetIntSlice(key string) []int
	GetStringSlice(key string) []string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetSizeInBytes(key string) uint
	Debug()
	IsSet(key string) bool
	RegisterAlias(alias string, key string)
	InConfig(key string) bool
	SetDefault(key string, value interface{})
	Set(key string, value interface{})
}

type ReaderFat interface {
	Reader
	MergeInConfig() error
	ReadConfig(in io.Reader) error
	MergeConfig(in io.Reader) error
	MergeConfigMap(cfg map[string]interface{}) error
	UnmarshalKey(key string, rawVal interface{}, opts ...viper.DecoderConfigOption) error
	Unmarshal(rawVal interface{}, opts ...viper.DecoderConfigOption) error
	UnmarshalExact(rawVal interface{}, opts ...viper.DecoderConfigOption) error
	BindPFlags(flags *pflag.FlagSet) error
	BindPFlag(key string, flag *pflag.Flag) error
	BindFlagValues(flags viper.FlagValueSet) (err error)
	BindFlagValue(key string, flag viper.FlagValue) error
	Common
	Writer
	ReadRemote
}

type Writer interface {
	WriteConfig() error
	SafeWriteConfig() error
	WriteConfigAs(filename string) error
	SafeWriteConfigAs(filename string) error
}

type ReadRemote interface {
	ReadRemoteConfig() error
	WatchRemoteConfig() error
	WatchRemoteConfigOnChannel() error
	AddRemoteProvider(provider, endpoint, path string) error
	AddSecureRemoteProvider(provider, endpoint, path, secretkeyring string) error
}

func New(file string) (Reader, error) {
	vipr := viper.New()
	vipr.SetConfigFile(file)
	vipr.AutomaticEnv()
	vipr.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := vipr.ReadInConfig()
	return vipr, err
}

// SetEnvPrefix(in string)
// 	AllowEmptyEnv(allowEmptyEnv bool)
// 	BindEnv(input ...string) error
// SetConfigFile(in string)
// 	AddConfigPath(in string)
// 	SetConfigName(in string)
// 	SetConfigType(in string)
// 	AutomaticEnv()
// 	SetEnvKeyReplacer(r *strings.Replacer)
// 	ReadInConfig() error
// 	SetConfigPermissions(perm os.FileMode)
