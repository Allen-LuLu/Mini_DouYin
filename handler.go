package main

import (
	publish "Mini_DouYin/kitex_gen/publish"
	"context"
)

// PublishImpl implements the last service interface defined in the IDL.
type PublishImpl struct{}

// Action implements the PublishImpl interface.
func (s *PublishImpl) Action(ctx context.Context, req *publish.ActionReq) (resp *publish.ActionResp, err error) {
	// TODO: Your code here...
	return
}

// List implements the PublishImpl interface.
func (s *PublishImpl) List(ctx context.Context, req *publish.ListReq) (resp *publish.ListResp, err error) {
	// TODO: Your code here...
	return
}
