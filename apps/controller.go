package apps

import (
	// 注册所有内部服务模块, 无须对外暴露的服务, 用于内部依赖
	_ "github.com/infraboard/mcenter/apps/counter/impl"
	_ "github.com/infraboard/mcenter/apps/ip2region/impl"

	// 引入第三方存储模块(Mongo)
	_ "github.com/infraboard/mcube/ioc/apps/oss/mongo"

	// 注册所有GRPC服务模块, 暴露给框架GRPC服务器加载, 注意 导入有先后顺序
	_ "github.com/infraboard/mcenter/apps/domain/impl"
	_ "github.com/infraboard/mcenter/apps/endpoint/impl"
	_ "github.com/infraboard/mcenter/apps/instance/impl"
	_ "github.com/infraboard/mcenter/apps/label/impl"
	_ "github.com/infraboard/mcenter/apps/namespace/impl"
	_ "github.com/infraboard/mcenter/apps/notify/impl"
	_ "github.com/infraboard/mcenter/apps/policy/impl"
	_ "github.com/infraboard/mcenter/apps/resource/impl"
	_ "github.com/infraboard/mcenter/apps/role/impl"
	_ "github.com/infraboard/mcenter/apps/service/impl"
	_ "github.com/infraboard/mcenter/apps/token/impl"
	_ "github.com/infraboard/mcenter/apps/user/impl"
)
