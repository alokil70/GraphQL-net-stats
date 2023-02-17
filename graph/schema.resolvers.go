package graph


import (
	"context"
	"test/graph/model"
)

func (r *mutationResolver) CreateIface(ctx context.Context, input model.NewIface) (*model.Iface, error) {
	iface := &model.Iface{
		Name:  "",
		TxSec: "",
		RxSec: "",
		Tx:    "",
		Rx:    "",
	}

	r.IfaceList = append(r.IfaceList, iface)

	return iface, nil
}

func (r *queryResolver) Ifaces(ctx context.Context) ([]*model.Iface, error) {

	var ifaces []*model.Iface = getInterfaces()
	
	return ifaces, nil
}

func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
