package doc

import (
	"fmt"
	"reflect"
	"sort"

	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/types"
	"github.com/naturalselectionlabs/pregod/pkg/jschema"
	"github.com/naturalselectionlabs/pregod/service/hub/internal/server/model"
	"github.com/shopspring/decimal"
)

type ExchangeResult struct{}

func (d *Doc) Define() {
	// Add some common handlers for primitive types
	d.schemas.AddTimeHandler()
	d.schemas.AddBigIntHandler()
	d.schemas.AddJSONRawMessageHandler()
	d.AddDecimalHandler()

	// Define the types that specific for this project
	{
		d.schemas.Define(model.Response{})

		d.DefineTransfer()

		d.schemas.SetSchema(ExchangeResult{}, d.schemas.AnyOf(model.CexResult{}, model.DexResult{}))

		d.DefinePlatformList()

		d.DefineNetworkList()
	}
}

func (d *Doc) AddDecimalHandler() {
	d.schemas.AddHandler(decimal.Decimal{}, func() *jschema.Schema {
		return &jschema.Schema{
			Description: "github.com/shopspring/decimal.Decimal",
			Title:       "Decimal",
			Type:        jschema.TypeString,
		}
	})
}

func (d *Doc) DefineTransfer() {
	d.schemas.Define(dbModel.Transfer{})

	type TransferTypes struct{}
	ts := d.schemas.Const(types.TransferTypes)
	d.schemas.SetSchema(TransferTypes{}, ts)
	d.schemas.GetSchema(TransferTypes{}).Description = `The name field on items of actions are the possible values of action field on metadata field.`

	s := d.schemas.GetSchema(dbModel.Transfer{})

	anyOf := []*jschema.Schema{}

	for _, tt := range types.TransferTypes {
		meta := d.schemas.DefineT(reflect.TypeOf(tt.Metadata).Elem())

		n := s.Clone()
		n.Description = fmt.Sprintf("Transfer object for tag: '%s', type: '%s', meta: '%s'", tt.Tag, tt.Type, meta.Ref.Name)

		n.Properties["tag"] = d.schemas.Const(tt.Tag)
		n.Properties["type"] = d.schemas.Const(tt.Type)
		n.Properties["metadata"] = meta

		anyOf = append(anyOf, n)
	}

	s.Properties = nil
	s.Description = "For all the possible types of transfer, see the TransferTypes in this doc."
	s.AnyOf = anyOf
}

func (d *Doc) DefinePlatformList() {
	list := []string{}

	for _, ps := range protocol.PlatformList {
		list = append(list, ps...)
	}

	sort.Strings(list)

	type PlatformName string

	d.schemas.SetSchema(PlatformName(""), &jschema.Schema{
		Type: jschema.TypeString,
		Enum: jschema.ToJValList(list...),
	})
}

func (d *Doc) DefineNetworkList() {
	list := []string{}

	list = append(list, protocol.EthclientNetworks...)
	list = append(list, protocol.SupportNetworks...)

	sort.Strings(list)

	list = removeDuplicate(list)

	type NetworkName string

	d.schemas.SetSchema(NetworkName(""), &jschema.Schema{
		Type: jschema.TypeString,
		Enum: jschema.ToJValList(list...),
	})
}

func removeDuplicate[T string | int](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
