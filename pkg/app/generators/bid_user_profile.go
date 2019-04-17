package generators

import (
	"github.com/chenhg5/go-admin/template/types"
	"github.com/chenhg5/go-admin/plugins/admin/models"
)

func GetBid_user_profileTable() (bid_user_profileTable models.Table) {

	bid_user_profileTable.Info.FieldList = []types.FieldStruct{{
			Head:     "Id",
			Field:    "id",
			TypeName: "int",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     "Created_at",
			Field:    "created_at",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     "Updated_at",
			Field:    "updated_at",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     "Email",
			Field:    "email",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     "Telephone",
			Field:    "telephone",
			TypeName: "varchar",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     "Last_login",
			Field:    "last_login",
			TypeName: "timestamp",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     "User_id",
			Field:    "user_id",
			TypeName: "int",
			Sortable: false,
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},}

	bid_user_profileTable.Info.Table = "bid_user_profile"
	bid_user_profileTable.Info.Title = "Bid_user_profile"
	bid_user_profileTable.Info.Description = "Bid_user_profile"

	bid_user_profileTable.Form.FormList = []types.FormStruct{{
			Head:     "Id",
			Field:    "id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "default",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     "Created_at",
			Field:    "created_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     "Updated_at",
			Field:    "updated_at",
			TypeName: "timestamp",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     "Email",
			Field:    "email",
			TypeName: "varchar",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     "Telephone",
			Field:    "telephone",
			TypeName: "varchar",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     "Last_login",
			Field:    "last_login",
			TypeName: "timestamp",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},{
			Head:     "User_id",
			Field:    "user_id",
			TypeName: "int",
			Default:  "",
			Editable: false,
			FormType: "text",
			ExcuFun: func(model types.RowModel) interface{} {
				return model.Value
			},
		},	}

	bid_user_profileTable.Form.Table = "bid_user_profile"
	bid_user_profileTable.Form.Title = "Bid_user_profile"
	bid_user_profileTable.Form.Description = "Bid_user_profile"

	bid_user_profileTable.ConnectionDriver = "mysql"

	return
}