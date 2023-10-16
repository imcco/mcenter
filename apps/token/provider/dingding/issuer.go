package dingding

import (
	"context"
	"fmt"

	"github.com/infraboard/mcenter/apps/domain"
	"github.com/infraboard/mcenter/apps/domain/password"
	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcenter/apps/token/provider"
	"github.com/infraboard/mcenter/apps/user"
	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/ioc"
	"github.com/infraboard/mcube/ioc/config/logger"
	"github.com/rs/zerolog"
)

type issuer struct {
	domain domain.Service
	user   user.Service
	log    *zerolog.Logger
}

func (i *issuer) Init() error {
	i.domain = ioc.GetController(domain.AppName).(domain.Service)
	i.user = ioc.GetController(user.AppName).(user.Service)
	i.log = logger.Sub("issuer.dingding")
	return nil
}

func (i *issuer) GrantType() token.GRANT_TYPE {
	return token.GRANT_TYPE_DINGDING
}

func (i *issuer) validate(ctx context.Context, username, code string) (*user.User, error) {
	// 从用户名中 获取到DN, 比如oldfish@default, 比如username: oldfish domain: default
	_, domainName := user.SpliteUserAndDomain(username)

	// 查询域下 对应的飞书设置
	dom, err := i.domain.DescribeDomain(ctx, domain.NewDescribeDomainRequestByName(domainName))
	if err != nil {
		return nil, err
	}

	if dom.Spec.DingdingSetting == nil {
		return nil, fmt.Errorf("domain dingding not settting")
	}

	// 获取Token
	client := NewDingDingClient(dom.Spec.DingdingSetting)
	dt, err := client.Login(ctx, code)
	if err != nil {
		return nil, err
	}

	// 获取用户信息
	du, err := client.GetUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	// 同步钉钉用户
	// 判断用户是否在数据库存在, 如果不存在需要同步到本地数据库
	lu, err := i.user.DescribeUser(ctx, user.NewDescriptUserRequestByName(du.Username()))
	if err != nil {
		if exception.IsNotFoundError(err) {
			i.log.Debug().Msgf("sync user: %s(%s) to db", du.Username(), dom.Spec.Name)
			gen := password.New(dom.Spec.PasswordConfig)
			randomPass, err := gen.Generate()
			if err != nil {
				return nil, err
			}
			// 创建本地用户
			newReq := du.ToCreateUserRequest(dom.Spec.Name, *randomPass, "系统自动生成")
			lu, err = i.user.CreateUser(ctx, newReq)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	// 更新用户Profile
	updateReq := user.NewPatchUserRequest(lu.Meta.Id)
	updateReq.Profile = du.ToProfile()
	updateReq.DingdingToken = dt
	u, err := i.user.UpdateUser(ctx, updateReq)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (i *issuer) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	u, err := i.validate(ctx, req.Username, req.AuthCode)
	if err != nil {
		return nil, err
	}

	// 颁发Token
	tk := token.NewToken(req)
	tk.Domain = u.Spec.Domain
	tk.Username = u.Spec.Username
	tk.UserType = u.Spec.Type
	tk.UserId = u.Meta.Id

	return tk, nil
}

func (i *issuer) IssueCode(ctx context.Context, req *token.IssueCodeRequest) (*token.Code, error) {
	_, err := i.validate(ctx, req.Username, req.AuthCode)
	if err != nil {
		return nil, err
	}

	// 颁发Token
	c, err := token.NewCode(req)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func init() {
	provider.Registe(&issuer{})
}
