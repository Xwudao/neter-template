package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/knadh/koanf/v2"

	"go.uber.org/zap"

	"go-kitboxpro/internal/biz"
	"go-kitboxpro/internal/core"
	"go-kitboxpro/internal/data/ent/user"
	"go-kitboxpro/internal/domain/params"
	"go-kitboxpro/internal/routes/mdw"
	"go-kitboxpro/internal/routes/valid"
)

type DataListRoute struct {
	conf *koanf.Koanf
	g    *gin.Engine
	log  *zap.SugaredLogger

	dlb *biz.DataListBiz
}

func NewDataListRoute(g *gin.Engine, dlb *biz.DataListBiz, log *zap.SugaredLogger, conf *koanf.Koanf) *DataListRoute {
	r := &DataListRoute{
		conf: conf,
		g:    g, dlb: dlb,
		log: log.Named("data-list-route"),
	}

	return r
}

func (r *DataListRoute) Reg() {
	// r.g.GET("/v1/data_list", core.WrapData(r.dataList()))

	group := r.g.Group("/v1/data_list")
	{
		group.GET("", core.WrapData(r.dataList()))
	}
	authGroup := r.g.Group("/auth/v1/data_list").Use(mdw.MustLoginMiddleware())
	{
		// authGroup.GET("/auth", core.WrapData(r.dataList()))
		_ = authGroup
	}
	adminGroup := r.g.Group("/admin/v1/data_list").Use(mdw.MustWithRoleMiddleware(user.RoleAdmin))
	{
		adminGroup.GET("/list", core.WrapData(r.list()))
		adminGroup.GET("/sort_data", core.WrapData(r.sortData()))
		adminGroup.POST("/create", core.WrapData(r.create()))
		adminGroup.POST("/update", core.WrapData(r.update()))
		adminGroup.POST("/update_order", core.WrapData(r.updateOrder()))
		adminGroup.POST("/delete", core.WrapData(r.delete()))
	}
}

func (r *DataListRoute) dataList() core.WrappedHandlerFunc {
	return func(c *gin.Context) (any, *core.RtnStatus) {
		//var (
		//    ctx = c.Request.Context()
		//)
		return "hello", nil
	}
}

func (r *DataListRoute) list() core.WrappedHandlerFunc {
	return func(c *gin.Context) (any, *core.RtnStatus) {
		var (
			ctx = c.Request.Context()
			pm  params.ListDataByKindParams
		)

		if err := c.ShouldBind(&pm); err != nil {
			return nil, core.NewRtnWithErr(valid.GetErrorMsg(&pm, err))
		}

		data, total, err := r.dlb.ListByKind(ctx, &pm)
		if err != nil {
			return nil, core.NewRtnWithErr(err)
		}

		return core.NewListRtn(data, total)
	}
}

func (r *DataListRoute) create() core.WrappedHandlerFunc {
	return func(c *gin.Context) (any, *core.RtnStatus) {
		var (
			ctx = c.Request.Context()
			pm  params.CreateDataListParams
		)

		if err := c.ShouldBindJSON(&pm); err != nil {
			return nil, core.NewRtnWithErr(valid.GetErrorMsg(&pm, err))
		}

		data, err := r.dlb.Create(ctx, &pm)
		if err != nil {
			return nil, core.NewRtnWithErr(err)
		}

		return data, nil
	}
}

func (r *DataListRoute) delete() core.WrappedHandlerFunc {
	return func(c *gin.Context) (any, *core.RtnStatus) {
		var (
			ctx = c.Request.Context()
			pm  params.DeleteIDParams
		)

		if err := c.ShouldBindJSON(&pm); err != nil {
			return nil, core.NewRtnWithErr(valid.GetErrorMsg(&pm, err))
		}

		if err := r.dlb.Delete(ctx, pm.ID); err != nil {
			return nil, core.NewRtnWithErr(err)
		}

		return nil, nil
	}
}

func (r *DataListRoute) update() core.WrappedHandlerFunc {
	return func(c *gin.Context) (any, *core.RtnStatus) {
		var (
			ctx = c.Request.Context()
			pm  params.UpdateDataListParams
		)

		if err := c.ShouldBindJSON(&pm); err != nil {
			return nil, core.NewRtnWithErr(valid.GetErrorMsg(&pm, err))
		}

		data, err := r.dlb.Update(ctx, &pm)
		if err != nil {
			return nil, core.NewRtnWithErr(err)
		}

		return data, nil
	}
}

func (r *DataListRoute) sortData() core.WrappedHandlerFunc {
	return func(c *gin.Context) (any, *core.RtnStatus) {
		var (
			ctx = c.Request.Context()
			pm  params.GetDataListSortDataParams
		)

		if err := c.ShouldBind(&pm); err != nil {
			return nil, core.NewRtnWithErr(valid.GetErrorMsg(&pm, err))
		}

		data, err := r.dlb.GetSortData(ctx, &pm)
		if err != nil {
			return nil, core.NewRtnWithErr(err)
		}

		return data, nil
	}
}

func (r *DataListRoute) updateOrder() core.WrappedHandlerFunc {
	return func(c *gin.Context) (any, *core.RtnStatus) {
		var (
			ctx = c.Request.Context()
			pm  params.ItemOrderParams
		)

		if err := c.ShouldBindJSON(&pm); err != nil {
			return nil, core.NewRtnWithErr(valid.GetErrorMsg(&pm, err))
		}

		if err := r.dlb.UpdateOrder(ctx, &pm); err != nil {
			return nil, core.NewRtnWithErr(err)
		}

		return nil, nil
	}
}
