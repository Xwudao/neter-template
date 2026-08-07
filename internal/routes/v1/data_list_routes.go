package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/knadh/koanf/v2"

	"go.uber.org/zap"

	"github.com/Xwudao/neter-template/internal/biz"
	"github.com/Xwudao/neter-template/internal/core"
	"github.com/Xwudao/neter-template/internal/data/ent"
	"github.com/Xwudao/neter-template/internal/data/ent/user"
	"github.com/Xwudao/neter-template/internal/domain/params"
	"github.com/Xwudao/neter-template/internal/routes/mdw"
)

type DataListRoute struct {
	conf *koanf.Koanf
	log  *zap.SugaredLogger

	dlb biz.DataListBizIface
}

func NewDataListRoute(dlb biz.DataListBizIface, log *zap.SugaredLogger, conf *koanf.Koanf) *DataListRoute {
	r := &DataListRoute{
		conf: conf,
		dlb:  dlb,
		log:  log.Named("data-list-route"),
	}

	return r
}

type DataListPageResponse struct {
	List  []*ent.DataList `json:"list"`
	Total int             `json:"total"`
}

func (r *DataListRoute) Register(router gin.IRouter) {
	// router.GET("/v1/data_list", core.NoInputE(r.dataList))

	group := router.Group("/v1/data_list")
	{
		group.GET("", core.NoInputE(r.dataList))
	}
	authGroup := router.Group("/auth/v1/data_list").Use(mdw.MustLoginMiddleware())
	{
		// authGroup.GET("/auth", core.NoInputE(r.dataList))
		_ = authGroup
	}
	adminGroup := router.Group("/admin/v1/data_list").Use(mdw.MustWithRoleMiddleware(user.RoleAdmin))
	{
		adminGroup.GET("/list", core.RequestE(r.list))
		adminGroup.GET("/sort_data", core.RequestE(r.sortData))
		adminGroup.POST("/create", core.JSONE(r.create))
		adminGroup.POST("/update", core.JSONE(r.update))
		adminGroup.POST("/update_order", core.JSONE(r.updateOrder))
		adminGroup.POST("/delete", core.JSONE(r.delete))
	}
}

func (r *DataListRoute) dataList(c *gin.Context) (string, error) {
	return "hello", nil
}

func (r *DataListRoute) list(c *gin.Context, pm *params.ListDataByKindParams) (DataListPageResponse, error) {
	data, total, err := r.dlb.ListByKind(c.Request.Context(), pm)
	if err != nil {
		return DataListPageResponse{}, err
	}
	return DataListPageResponse{List: data, Total: total}, nil
}

func (r *DataListRoute) create(c *gin.Context, pm *params.CreateDataListParams) (*ent.DataList, error) {
	data, err := r.dlb.Create(c.Request.Context(), pm)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *DataListRoute) delete(c *gin.Context, pm *params.DeleteIDParams) (*core.EmptyResponse, error) {
	if err := r.dlb.Delete(c.Request.Context(), pm.ID); err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *DataListRoute) update(c *gin.Context, pm *params.UpdateDataListParams) (*ent.DataList, error) {
	data, err := r.dlb.Update(c.Request.Context(), pm)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *DataListRoute) sortData(c *gin.Context, pm *params.GetDataListSortDataParams) ([]*ent.DataList, error) {
	data, err := r.dlb.GetSortData(c.Request.Context(), pm)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *DataListRoute) updateOrder(c *gin.Context, pm *params.ItemOrderParams) (*core.EmptyResponse, error) {
	if err := r.dlb.UpdateOrder(c.Request.Context(), pm); err != nil {
		return nil, err
	}
	return nil, nil
}
