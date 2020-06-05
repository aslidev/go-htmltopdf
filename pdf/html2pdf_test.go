package pdf

import (
	"io/ioutil"
	"testing"
)

const HTML = `<html>
<style type="text/css">
/* Background */ .chroma { color: #f8f8f2; background-color: #282a36 }
/* LineTableTD */ .chroma .lntd { vertical-align: top; padding: 0; margin: 0; border: 0; }
/* LineTable */ .chroma .lntable { border-spacing: 0; padding: 0; margin: 0; border: 0; width: auto; overflow: auto; display: block; }
/* LineHighlight */ .chroma .hl { display: block; width: 100%;background-color: #3d3f4a }
/* LineNumbersTable */ .chroma .lnt { margin-right: 0.4em; padding: 0 0.4em 0 0.4em;color: #7f7f7f }
/* LineNumbers */ .chroma .ln { margin-right: 0.4em; padding: 0 0.4em 0 0.4em;color: #7f7f7f }
/* Keyword */ .chroma .k { color: #ff79c6 }
/* KeywordConstant */ .chroma .kc { color: #ff79c6 }
/* KeywordDeclaration */ .chroma .kd { color: #8be9fd; font-style: italic }
/* KeywordNamespace */ .chroma .kn { color: #ff79c6 }
/* KeywordPseudo */ .chroma .kp { color: #ff79c6 }
/* KeywordReserved */ .chroma .kr { color: #ff79c6 }
/* KeywordType */ .chroma .kt { color: #8be9fd }
/* NameAttribute */ .chroma .na { color: #50fa7b }
/* NameBuiltin */ .chroma .nb { color: #8be9fd; font-style: italic }
/* NameClass */ .chroma .nc { color: #50fa7b }
/* NameFunction */ .chroma .nf { color: #50fa7b }
/* NameLabel */ .chroma .nl { color: #8be9fd; font-style: italic }
/* NameTag */ .chroma .nt { color: #ff79c6 }
/* NameVariable */ .chroma .nv { color: #8be9fd; font-style: italic }
/* NameVariableClass */ .chroma .vc { color: #8be9fd; font-style: italic }
/* NameVariableGlobal */ .chroma .vg { color: #8be9fd; font-style: italic }
/* NameVariableInstance */ .chroma .vi { color: #8be9fd; font-style: italic }
/* LiteralString */ .chroma .s { color: #f1fa8c }
/* LiteralStringAffix */ .chroma .sa { color: #f1fa8c }
/* LiteralStringBacktick */ .chroma .sb { color: #f1fa8c }
/* LiteralStringChar */ .chroma .sc { color: #f1fa8c }
/* LiteralStringDelimiter */ .chroma .dl { color: #f1fa8c }
/* LiteralStringDoc */ .chroma .sd { color: #f1fa8c }
/* LiteralStringDouble */ .chroma .s2 { color: #f1fa8c }
/* LiteralStringEscape */ .chroma .se { color: #f1fa8c }
/* LiteralStringHeredoc */ .chroma .sh { color: #f1fa8c }
/* LiteralStringInterpol */ .chroma .si { color: #f1fa8c }
/* LiteralStringOther */ .chroma .sx { color: #f1fa8c }
/* LiteralStringRegex */ .chroma .sr { color: #f1fa8c }
/* LiteralStringSingle */ .chroma .s1 { color: #f1fa8c }
/* LiteralStringSymbol */ .chroma .ss { color: #f1fa8c }
/* LiteralNumber */ .chroma .m { color: #bd93f9 }
/* LiteralNumberBin */ .chroma .mb { color: #bd93f9 }
/* LiteralNumberFloat */ .chroma .mf { color: #bd93f9 }
/* LiteralNumberHex */ .chroma .mh { color: #bd93f9 }
/* LiteralNumberInteger */ .chroma .mi { color: #bd93f9 }
/* LiteralNumberIntegerLong */ .chroma .il { color: #bd93f9 }
/* LiteralNumberOct */ .chroma .mo { color: #bd93f9 }
/* Operator */ .chroma .o { color: #ff79c6 }
/* OperatorWord */ .chroma .ow { color: #ff79c6 }
/* Comment */ .chroma .c { color: #6272a4 }
/* CommentHashbang */ .chroma .ch { color: #6272a4 }
/* CommentMultiline */ .chroma .cm { color: #6272a4 }
/* CommentSingle */ .chroma .c1 { color: #6272a4 }
/* CommentSpecial */ .chroma .cs { color: #6272a4 }
/* CommentPreproc */ .chroma .cp { color: #ff79c6 }
/* CommentPreprocFile */ .chroma .cpf { color: #ff79c6 }
/* GenericDeleted */ .chroma .gd { color: #8b080b }
/* GenericEmph */ .chroma .ge { text-decoration: underline }
/* GenericHeading */ .chroma .gh { font-weight: bold }
/* GenericInserted */ .chroma .gi { font-weight: bold }
/* GenericOutput */ .chroma .go { color: #44475a }
/* GenericSubheading */ .chroma .gu { font-weight: bold }
/* GenericUnderline */ .chroma .gl { text-decoration: underline }
body { color: #f8f8f2; background-color: #282a36; }
</style><body class="chroma">
<pre class="chroma">
<span class="kd">var</span> <span class="nx">buf</span> <span class="nx">bytes</span><span class="p">.</span><span class="nx">Buffer</span>
<span class="nx">fmt</span><span class="p">.</span><span class="nf">Fprintf</span><span class="p">(</span><span class="o">&amp;</span><span class="nx">buf</span><span class="p">,</span> <span class="s">&#34;Size: %d MB.&#34;</span><span class="p">,</span> <span class="mi">85</span><span class="p">)</span>
<span class="nx">s</span> <span class="o">:=</span> <span class="nx">buf</span><span class="p">.</span><span class="nf">String</span><span class="p">())</span> <span class="c1">// s == &#34;Size: 85 MB.&#34;
</span></pre>
</body>
</html>
`

func TestHtml2pdf(t *testing.T) {
	newObj := New()
	defer newObj.Destroy()

	/* newObj.SetData("Testing html to pdf go bindings")
	newObj.SetOutputFileName("test.pdf")

	newObj.OnProgressChanged(func(a int) { fmt.Println("Progress::", a, "%") })

	newObj.OnPhaseChanged(func(msg string) { fmt.Println("Phase::", msg) })

	newObj.OnError(func(msg string) { fmt.Println("Error::", msg) })

	newObj.OnWarning(func(msg string) { fmt.Println("Warning::", msg) })

	err, _ := newObj.CreatePDF()
	if err != nil {
		t.Fatal(err)
	} */

	// newObj.SetURL("https://www.baidu.com")
	newObj.SetData(HTML)
	newObj.SetBufferedOutput()
	err, data := newObj.CreatePDF()
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile("baidu.pdf", data, 0x777)
	if err != nil {
		t.Fatal(err)
	}
}
