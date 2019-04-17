package generators

import "github.com/chenhg5/go-admin/plugins/admin/models"

// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.

// admincli generate -h=127.0.0.1 -p=3306 -u=zs -P=zs -n=bid -o=./
var Generators = map[string]models.TableGenerator{
	"user":    GetBid_userTable,
	"user-profile":    GetBid_user_profileTable,
}