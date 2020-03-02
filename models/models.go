package models

import "github.com/rs/xid"

type (
	//Client
	Client struct {
		ID        xid.ID `json:"id" form:"id" binding:""`
		Name      string `json:"name" form:"name" binding:"required"`
		CreatedAt int64  `json:"create_at"`
	}

	//Order
	Order struct {
		ID        xid.ID `json:"id" form:"id" binding:""`
		Serial    string `json:"serial"`
		Invoice   string `json:"invoice"`
		Files     []File `json:"files"`
		CreatedAt int64  `json:"create_at"`
	}

	//File
	File struct {
		Path      string `json:"path"`
		Name      string `json:"name"`
		CreatedAt int64  `json:"create_at"`
	}
)
