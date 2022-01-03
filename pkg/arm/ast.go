package arm

import "github.com/alecthomas/participle/v2/lexer"

var ArmLexer = lexer.MustSimple([]lexer.Rule{
	// {Name: "Comment", Pattern: `(?i)rem[^\n]*`},
	{Name: "String", Pattern: `'(\\'|[^'])*'`},
	{Name: "Number", Pattern: `[-+]?(\d*\.)?\d+`},
	{Name: "Ident", Pattern: `[a-zA-Z_]\w*`},
	{Name: "Punct", Pattern: `[-[!@#$%^&*()+_={}\|:;"'<,>.?/]|]`},
	// {Name: "EOL", Pattern: `[\n\r]+`},
	{Name: "whitespace", Pattern: `[ \t]+`},
})

type ArmProgram struct {
	Pos lexer.Position

	Expression *ArmExpression `"[" @@ "]"`
}

type ArmExpression struct {
	Pos lexer.Position

	Call          *ArmCall       `( @@`
	Value         *ArmValue      `| @@`
	Subexpression *ArmExpression `| "(" @@ ")" )`
}

type ArmCall struct {
	Pos lexer.Position

	Name string           `@Ident`
	Args []*ArmExpression `"(" ( @@ ( ","  @@ )* )? ")"`
}

type ArmValue struct {
	Pos lexer.Position

	Integer  *float64 `  @Number`
	Variable *string  `| @Ident`
	String   *string  `| @String`
}
