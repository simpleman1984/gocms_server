package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Review struct {
	item.Item

	Title  string `json:"title"`
	Body   string `json:"body"`
	Rating int    `json:"rating"`
}

// MarshalEditor writes a buffer of html to edit a Review within the CMS
// and implements editor.Editable
func (r *Review) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(r,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Review field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Title", r, map[string]string{
				"label":       "Title",
				"type":        "text",
				"placeholder": "Enter the Title here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Body", r, map[string]string{
				"label":       "Body",
				"placeholder": "Enter the Body here",
			}),
		},
		editor.Field{
			View: editor.Input("Rating", r, map[string]string{
				"label":       "Rating",
				"type":        "text",
				"placeholder": "Enter the Rating here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Review editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Review"] = func() interface{} { return new(Review) }
}

// String defines how a Review is printed. Update it using more descriptive
// fields from the Review struct type
func (r *Review) String() string {
	return fmt.Sprintf("Review: %s", r.UUID)
}
