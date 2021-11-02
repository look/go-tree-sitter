package svelte_test

import (
	"context"
	"testing"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/svelte"
	"github.com/stretchr/testify/assert"
)

func TestGrammar(t *testing.T) {
	assert := assert.New(t)

	n, err := sitter.Parse(context.Background(), []byte(`<script context="module">
    let name = 'world';
	</script>
	<h1>Hello {name'<>{}"\''""{}}!</h1>
	`), svelte.GetLanguage())

	assert.NoError(err)
	assert.Equal(
		"(document (script_element (start_tag (tag_name) (attribute (attribute_name) (quoted_attribute_value (attribute_value)))) (raw_text) (end_tag (tag_name))) (text) (element (start_tag (tag_name)) (text) (expression (raw_text_expr)) (text) (end_tag (tag_name))) (text))",
		n.String(),
	)
}
