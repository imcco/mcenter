package protocol

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/emicklei/go-restful/v3"
	"github.com/infraboard/mcenter/apps/endpoint"
	"github.com/infraboard/mcenter/swagger"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/apps/apidoc"
	"github.com/infraboard/mcube/v2/ioc/apps/health"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/contrib/instrumentation/github.com/emicklei/go-restful/otelrestful"

	"github.com/infraboard/mcenter/protocol/auth"
	"github.com/infraboard/mcube/v2/ioc/config/application"
	ioc_http "github.com/infraboard/mcube/v2/ioc/config/http"
	"github.com/infraboard/mcube/v2/ioc/config/log"
)

// NewHTTPService 构建函数
func NewHTTPService() *HTTPService {
	r := restful.DefaultContainer
	restful.DefaultResponseContentType(restful.MIME_JSON)
	restful.DefaultRequestContentType(restful.MIME_JSON)

	// CORS中间件
	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"*"},
		AllowedDomains: []string{"*"},
		AllowedMethods: []string{"HEAD", "OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE"},
		CookiesAllowed: false,
		Container:      r,
	}
	r.Filter(cors.Filter)
	// trace中间件
	filter := otelrestful.OTelFilter(application.Get().AppName)
	restful.DefaultContainer.Filter(filter)
	// 添加鉴权中间件
	r.Filter(auth.NewHttpAuther().GoRestfulAuthFunc)

	server := &http.Server{
		ReadHeaderTimeout: 60 * time.Second,
		ReadTimeout:       60 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1M
		Addr:              ioc_http.Get().Addr(),
		Handler:           r,
	}

	return &HTTPService{
		r:      r,
		server: server,
		l:      log.Sub("http"),

		apiDocPath: "/apidocs.json",
	}
}

// HTTPService http服务
type HTTPService struct {
	r      *restful.Container
	l      *zerolog.Logger
	server *http.Server

	apiDocPath string
}

func (s *HTTPService) PathPrefix() string {
	return fmt.Sprintf("/%s/api", application.Get().AppName)
}

// Start 启动服务
func (s *HTTPService) Start() {
	// 装置子服务路由
	ioc.LoadGoRestfulApi(s.PathPrefix(), s.r)

	// API Doc
	// Optionally, you can install the Swagger Service which provides a nice Web UI on your REST API
	// You need to download the Swagger HTML5 assets and change the FilePath location in the config below.
	// Open http://localhost:8080/apidocs/?url=http://localhost:8080/apidocs.json
	// http.Handle("/apidocs/", http.StripPrefix("/apidocs/", http.FileServer(http.Dir("/Users/emicklei/Projects/swagger-ui/dist"))))
	s.r.Add(apidoc.APIDocs(s.apiDocPath, swagger.Docs))
	s.l.Info().Msgf("Swagger API Doc访问地址: http://%s%s", ioc_http.Get().Addr(), s.apiDocPath)

	// HealthCheck
	hc := health.NewDefaultHealthChecker()
	s.r.Add(hc.WebService())
	s.l.Info().Msgf("健康检查地址: http://%s%s", ioc_http.Get().Addr(), hc.HealthCheckPath)

	// 注册路由条目
	s.RegistryEndpoint()

	// 启动 HTTP服务
	s.l.Info().Msgf("HTTP服务启动成功, 监听地址: %s", s.server.Addr)
	if err := s.server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			s.l.Info().Msgf("service is stopped")
			return
		}
		s.l.Error().Msgf("start service error, %s", err.Error())
	}
}

// 注册服务权限条目
func (s *HTTPService) RegistryEndpoint() {
	s.l.Info().Msgf("start registry endpoints ...")

	entries := []*endpoint.Entry{}
	wss := s.r.RegisteredWebServices()
	for i := range wss {
		es := endpoint.TransferRoutesToEntry(wss[i].Routes())
		entries = append(entries, endpoint.GetPRBACEntry(es)...)
	}

	req := endpoint.NewRegistryRequest(application.Get().AppName, entries)
	req.ServiceId = application.Get().AppName
	controller := ioc.Controller().Get(endpoint.AppName).(endpoint.Service)
	_, err := controller.RegistryEndpoint(context.Background(), req)
	if err != nil {
		s.l.Warn().Msgf("registry endpoints error, %s", err)
	} else {
		s.l.Debug().Msgf("service endpoints registry success")
	}
}

// Stop 停止server
func (s *HTTPService) Stop() error {
	s.l.Info().Msgf("start graceful shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	// 优雅关闭HTTP服务
	if err := s.server.Shutdown(ctx); err != nil {
		s.l.Error().Msgf("graceful shutdown timeout, force exit")
	}
	return nil
}
