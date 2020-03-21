package delivery

import (
	"github.com/gofiber/fiber"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/ryssapp/backend/src/go/common/pb"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"go.uber.org/zap"
)

type httpServer struct {
	address      string
	app          *fiber.App
	StoreService pb.StoreServiceClient
}

func NewHTTPServer(address string) *httpServer {
	app := fiber.New()
	s := &httpServer{
		address: address,
		app:     app,
	}
	s.initRoutes()
	return s
}

func (s *httpServer) initRoutes() {
	schemaConfig := graphql.SchemaConfig{Query: s.createQuery()}
	schema, _ := graphql.NewSchema(schemaConfig)
	h := fasthttpadaptor.NewFastHTTPHandler(handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	}))
	s.app.Post("/graphql", func(ctx *fiber.Ctx) {
		ctx.Fasthttp.SetUserValue("token", ctx.Get("Authorization"))
		h(ctx.Fasthttp)
	})
}

func (s *httpServer) Start() {
	zap.L().Fatal("Error while serving.", zap.Error(s.app.Listen(s.address)))
}
