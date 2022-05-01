package settingutil

import (
	"encoding/json"
	"errors"

	"github.com/warrior21st/go-utils/commonutil"
	"github.com/warrior21st/go-utils/jsonutil"
)

var (
	_settingsCache   interface{}
	_settingFilePath string
)

//获取指定设置，嵌套key使用":"分隔，如"AppSettings:DBConnectionString"
func GetAppSetting(keys string) string {
	if _settingsCache == nil {
		json.Unmarshal(commonutil.ReadFileBytes(GetSettingFilePath()), &_settingsCache)
	}

	return jsonutil.ReadJsonValFromDecodedBytes(_settingsCache, keys)
}

//获取appsetting文件位置
func GetSettingFilePath() string {
	if commonutil.IsNilOrWhiteSpace(_settingFilePath) {
		_settingFilePath = commonutil.CombinePath(commonutil.GetProgramRootPath(), "appsettings.development.json")
		if !commonutil.IsExistPath(_settingFilePath) {
			_settingFilePath = commonutil.CombinePath(commonutil.GetProgramRootPath(), "appsettings.test.json")
		}
		if !commonutil.IsExistPath(_settingFilePath) {
			_settingFilePath = commonutil.CombinePath(commonutil.GetProgramRootPath(), "appsettings.production.json")
		}
		if !commonutil.IsExistPath(_settingFilePath) {
			_settingFilePath = commonutil.CombinePath(commonutil.GetProgramRootPath(), "appsettings.json")
		}
		if !commonutil.IsExistPath(_settingFilePath) {
			panic(errors.New("can not find setting file in " + commonutil.GetProgramRootPath() + "."))
		}
	}

	//
	return _settingFilePath
}
