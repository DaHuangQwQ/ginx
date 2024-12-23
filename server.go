package ginx

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"path/filepath"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/DaHuangQwQ/ginx/openapi"
	"github.com/gin-gonic/gin"
)

type Server struct {
	*gin.Engine
	OpenAPI *openapi.OpenAPI
	addr    string
}

func NewServer(addr string, opts ...gin.OptionFunc) *Server {
	return &Server{
		addr:   addr,
		Engine: gin.Default(opts...),
	}
}

func (s *Server) Handle(method, path string, handler gin.HandlerFunc) {
	s.Engine.Handle(method, path, handler)
}

func (s *Server) Start() error {
	return s.Engine.Run(s.addr)
}

func (s *Server) MarshalSpec() ([]byte, error) {
	s.OpenAPI = Oai
	return json.MarshalIndent(s.OpenAPI.Description(), "", "	")
}

func (s *Server) SaveOpenAPIToFile(path string) error {
	jsonFolder := filepath.Dir(path)

	err := os.MkdirAll(jsonFolder, 0o750)
	if err != nil {
		return errors.New("error creating docs directory")
	}

	f, err := os.Create(path)
	if err != nil {
		return errors.New("error creating file")
	}
	defer f.Close()

	marshal, err := s.MarshalSpec()
	if err != nil {
		return err
	}

	_, err = f.Write(marshal)
	if err != nil {
		return errors.New("error writing file ")
	}

	return nil
}

// RegisterOpenAPIRoutes Registers the routes to serve the OpenAPI spec and Swagger UI.
func (s *Server) RegisterOpenAPIRoutes(path string) {
	s.GET(path, func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		spec, err := s.MarshalSpec()
		if err != nil {
			return
		}
		ctx.String(http.StatusOK, string(spec))
	})
	s.GET(path+"/*any", httpToGinHandler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost"+s.addr+path),
	)))
}

// 转换 http.HandlerFunc 为 Gin HandlerFunc
func httpToGinHandler(httpHandler http.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 创建一个标准的 http.ResponseWriter 和 *http.Request
		rw := c.Writer
		req := c.Request

		// 调用 http.HandlerFunc 处理请求
		httpHandler(rw, req)
	}
}
