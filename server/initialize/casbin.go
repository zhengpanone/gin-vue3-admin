package initialize

import (
	"errors"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"

	"go.uber.org/zap"

	"github.com/zhengpanone/gin-vue3-admin/global"
)

var (
	enforcerInstance *casbin.SyncedCachedEnforcer
	once             sync.Once
	initErr          error
)

// InitCasbin 持久化到数据库  引入自定义规则
//
//	用 sync.Once 确保只初始化一次 Enforcer
func InitCasbin() {
	once.Do(func() {
		adapter, err := gormadapter.NewAdapterByDB(global.GVA_DB)
		if err != nil {
			initErr = errors.New("适配数据库失败，请检查 casbin 表是否为 InnoDB 引擎")
			global.GVA_LOG.Error(initErr.Error(), zap.Error(err))
			return
		}
		text := `
		[request_definition]
		r = sub, obj, act # sub=主体（用户/角色），obj=资源，act=动作
		
		[policy_definition]
		p = sub, obj, act # sub=角色，obj=资源路径，act=动作
		
		[role_definition]
		g = _, _ # g = 用户 -> 角色 绑定关系
		
		[policy_effect]
		e = some(where (p.eft == allow)) # allow优先
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			global.GVA_LOG.Error("加载 Casbin 模型失败", zap.Error(err))
			return
		}
		enforcerInstance, err = casbin.NewSyncedCachedEnforcer(m, adapter)
		if err != nil {
			global.GVA_LOG.Error("创建 Casbin Enforcer 失败", zap.Error(err))
			return
		}
		enforcerInstance.SetExpireTime(60 * 60)
		if err = enforcerInstance.LoadPolicy(); err != nil {
			global.GVA_LOG.Error("加载 Casbin 策略失败", zap.Error(err))
			enforcerInstance = nil
		}
	})

}
