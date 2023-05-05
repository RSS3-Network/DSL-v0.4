package doc

import (
	"fmt"
	"reflect"

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

	s := d.schemas.GetRawSchema(dbModel.Transfer{})

	desc := "All the possible mappings for the transfer object:\n\n"

	t := &jschema.Schema{
		Type:  s.Type,
		Title: s.Title,
		AnyOf: []*jschema.Schema{},
	}

	for meta, conditions := range filter.MetadataTypeMap {
		metaS := d.schemas.DefineT(reflect.TypeOf(meta).Elem())

		desc += fmt.Sprintf("- %s:\n", metaS.Ref.ID)

		for _, cond := range conditions {
			n := s.Clone()

			desc += fmt.Sprintf("  - tag: %s, type: %s\n", cond.Tag, cond.Type)

			n.Properties["tag"] = d.schemas.Const(cond.Tag)
			n.Properties["type"] = d.schemas.Const(cond.Type)
			n.Properties["metadata"] = metaS

			t.AnyOf = append(t.AnyOf, n)
		}
	}

	// TODO: Remove it once the MetadataTypeMap is complete
	t.AnyOf = append(t.AnyOf, s.Clone())

	*s = *t

	s.Description = desc
}

func (d *Doc) DefinePlatformList() {
	list := []jschema.JVal{}

	for _, ps := range protocol.PlatformList {
		for _, p := range ps {
			list = append(list, p)
		}
	}

	type PlatformName string

	d.schemas.SetSchema(PlatformName(""), &jschema.Schema{
		Type: jschema.TypeString,
		Enum: list,
	})
}
