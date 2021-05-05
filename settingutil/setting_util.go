package settingutil

import (
	"encoding/json"

	"github.com/warrior21st/go-utils/commonutil"
	"github.com/warrior21st/go-utils/jsonutil"
)

var settingCache interface{}

//获取指定设置，嵌套key使用":"分隔，如"AppSettings:DBConnectionString"
func GetAppSetting(keys string) string {
	if settingCache == nil {
		settingFilePath := commonutil.CombinePath(commonutil.GetProgramRootPath(), "appsettings.json")
		json.Unmarshal(commonutil.ReadFileBytes(settingFilePath), &settingCache)
	}

	return jsonutil.ReadJsonValFromDecodedBytes(settingCache, keys)
}
