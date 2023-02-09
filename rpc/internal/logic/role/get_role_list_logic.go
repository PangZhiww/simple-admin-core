package role

import (
	"context"

	"github.com/suyuan32/simple-admin-core/pkg/ent/predicate"
	"github.com/suyuan32/simple-admin-core/pkg/ent/role"
	"github.com/suyuan32/simple-admin-core/rpc/internal/svc"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
)

type GetRoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleListLogic {
	return &GetRoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRoleListLogic) GetRoleList(in *core.RoleListReq) (*core.RoleListResp, error) {
	var predicates []predicate.Role
	if in.Name != "" {
		predicates = append(predicates, role.NameContains(in.Name))
	}
	if in.Value != "" {
		predicates = append(predicates, role.ValueContains(in.Value))
	}
	if in.DefaultRouter != "" {
		predicates = append(predicates, role.DefaultRouterContains(in.DefaultRouter))
	}
	result, err := l.svcCtx.DB.Role.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)
	if err != nil {
		logx.Error(err.Error())
		return nil, statuserr.NewInternalError(i18n.DatabaseError)
	}

	resp := &core.RoleListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &core.RoleInfo{
			Id:            v.ID,
			CreatedAt:     v.CreatedAt.UnixMilli(),
			UpdatedAt:     v.UpdatedAt.UnixMilli(),
			Status:        uint32(v.Status),
			Name:          v.Name,
			Value:         v.Value,
			DefaultRouter: v.DefaultRouter,
			Remark:        v.Remark,
			Sort:          v.Sort,
		})
	}

	return resp, nil
}
