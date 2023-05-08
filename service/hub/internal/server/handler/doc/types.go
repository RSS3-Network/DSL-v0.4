package doc

import (
	"fmt"
	"reflect"
	"sort"

	dbModel "github.com/naturalselectionlabs/pregod/common/database/model"
	"github.com/naturalselectionlabs/pregod/common/protocol"
	"github.com/naturalselectionlabs/pregod/common/protocol/filter"
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

	s := d.schemas.GetSchema(dbModel.Transfer{})

	desc := s.Description + "\n\nAll the possible mappings for the transfer object:\n\n"

	anyOf := []*jschema.Schema{}

	for _, m := range filter.MetadataMapping {
		meta := d.schemas.DefineT(reflect.TypeOf(m.Metadata).Elem())

		desc += fmt.Sprintf("- %s:\n", meta.Ref.ID)

		for _, c := range m.Criteria {
			n := s.Clone()
			n.Description = fmt.Sprintf("Transfer object for tag: '%s', type: '%s', meta: '%s'", c.Tag, c.Type, meta.Ref.Name)

			desc += fmt.Sprintf("  - tag: %s, type: %s\n", c.Tag, c.Type)

			n.Properties["tag"] = d.schemas.Const(c.Tag)
			n.Properties["type"] = d.schemas.Const(c.Type)
			n.Properties["metadata"] = meta

			anyOf = append(anyOf, n)
		}
	}

	s.Properties = nil
	s.Description = desc
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
