package efp

import (
	"fmt"
	"testing"
)

func TestEFP(t *testing.T) {
	testStrings := [][]string{
		{"+", "[]"},
		{"", "[]"},
		{"                 ", "[]"},
		{"=SUM())", "[{SUM Function Start} { Function Stop} { Function Stop}]"},
		{"=SUM(\"\")", "[{SUM Function Start} { Operand Text} { Function Stop}]"},
		{"=;", "[{; Unknown }]"},
		{"=1;2", "[{1 Operand Number} {; Unknown } {2 Operand Number}]"},
		{"=SUM(1;2)", "[{SUM Function Start} {1 Operand Number} {; Unknown } {2 Operand Number} { Function Stop}]"},
		{"0(((;)))", "[{0 Function Start} { Subexpression Start} { Subexpression Start} {; Unknown } { Subexpression Stop} { Subexpression Stop} { Function Stop}]"},
		{"={1,2;3,4}", "[{ARRAY Function Start} {ARRAYROW Function Start} {1 Operand Number} {, Argument } {2 Operand Number} { Function Stop} {, Argument } {ARRAYROW Function Start} {3 Operand Number} {, Argument } {4 Operand Number} { Function Stop} { Function Stop}]"},
		{"={(1;2);3}", "[{ARRAY Function Start} {ARRAYROW Function Start} { Subexpression Start} {1 Operand Number} {; Unknown } {2 Operand Number} { Subexpression Stop} { Function Stop} {, Argument } {ARRAYROW Function Start} {3 Operand Number} { Function Stop} { Function Stop}]"},
		{"=\"あいうえお\"&H3&\"b\"", "[{あいうえお Operand Text} {& OperatorInfix Concatenation} {H3 Operand Range} {& OperatorInfix Concatenation} {b Operand Text}]"},
		{"=1+3+5", "[{1 Operand Number} {+ OperatorInfix Math} {3 Operand Number} {+ OperatorInfix Math} {5 Operand Number}]"},
		{"=3 * 4 + 5", "[{3 Operand Number} {* OperatorInfix Math} {4 Operand Number} {+ OperatorInfix Math} {5 Operand Number}]"},
		{"=50", "[{50 Operand Number}]"},
		{"=1+1", "[{1 Operand Number} {+ OperatorInfix Math} {1 Operand Number}]"},
		{"=$A1", "[{$A1 Operand Range}]"},
		{"=$B$2", "[{$B$2 Operand Range}]"},
		{"=SUM(B5:B15)", "[{SUM Function Start} {B5:B15 Operand Range} { Function Stop}]"},
		{"=SUM(B5:B15,D5:D15)", "[{SUM Function Start} {B5:B15 Operand Range} {, Argument } {D5:D15 Operand Range} { Function Stop}]"},
		{"=SUM(B5:B15 A7:D7)", "[{SUM Function Start} {B5:B15 Operand Range} { OperatorInfix Intersection} {A7:D7 Operand Range} { Function Stop}]"},
		{"=SUM(sheet1!$A$1:$B$2)", "[{SUM Function Start} {sheet1!$A$1:$B$2 Operand Range} { Function Stop}]"},
		{"=[data.xls]sheet1!$A$1", "[{[data.xls]sheet1!$A$1 Operand Range}]"},
		{"=[#data.xls]", "[{[#data.xls] Operand Range}]"},
		{"=[{data.xls]", "[{[{data.xls] Operand Range}]"},
		{"=SUM((A:A 1:1))", "[{SUM Function Start} { Subexpression Start} {A:A Operand Range} { OperatorInfix Intersection} {1:1 Operand Range} { Subexpression Stop} { Function Stop}]"},
		{"=SUM((A:A,1:1))", "[{SUM Function Start} { Subexpression Start} {A:A Operand Range} {, OperatorInfix Union} {1:1 Operand Range} { Subexpression Stop} { Function Stop}]"},
		{"=SUM((A:A A1:B1))", "[{SUM Function Start} { Subexpression Start} {A:A Operand Range} { OperatorInfix Intersection} {A1:B1 Operand Range} { Subexpression Stop} { Function Stop}]"},
		{"=SUM(D9:D11,E9:E11,F9:F11)", "[{SUM Function Start} {D9:D11 Operand Range} {, Argument } {E9:E11 Operand Range} {, Argument } {F9:F11 Operand Range} { Function Stop}]"},
		{"=SUM((D9:D11,(E9:E11,F9:F11)))", "[{SUM Function Start} { Subexpression Start} {D9:D11 Operand Range} {, OperatorInfix Union} { Subexpression Start} {E9:E11 Operand Range} {, OperatorInfix Union} {F9:F11 Operand Range} { Subexpression Stop} { Subexpression Stop} { Function Stop}]"},
		{"=((D2 * D3) + D4) & \" should be 10\"", "[{ Subexpression Start} { Subexpression Start} {D2 Operand Range} {* OperatorInfix Math} {D3 Operand Range} { Subexpression Stop} {+ OperatorInfix Math} {D4 Operand Range} { Subexpression Stop} {& OperatorInfix Concatenation} { should be 10 Operand Text}]"},
		{"=AND(1=1),1=1", "[{AND Function Start} {1 Operand Number} {= OperatorInfix Logical} {1 Operand Number} { Function Stop} {, OperatorInfix Union} {1 Operand Number} {= OperatorInfix Logical} {1 Operand Number}]"},
		{"='x'", "[{x Operand Range}]"},
		{"=a\"b\"\"", "[{a Unknown } {b\" Operand Range}]"},
		{"=#]#NUM!", "[{#]#NUM! Operand Range}]"},
		{"=3.1E-24-2.1E-24", "[{3.1E-24 Operand Number} {- OperatorInfix Math} {2.1E-24 Operand Number}]"},
		{"''", "[]"},
		{"=IF(R#", "[{IF Function Start} {R Unknown } {# Operand Range}]"},
		{"=IF(R{", "[{IF Function Start} {R Unknown } {ARRAY Function Start} {ARRAYROW Function Start}]"},
		{"=\"\"+'''", "[{ Operand Text} {+ OperatorInfix Math} {' Operand Range}]"},
		{"=1%2", "[{1 Operand Number} {% OperatorPostfix } {2 Operand Number}]"},
		{"={1,2}", "[{ARRAY Function Start} {ARRAYROW Function Start} {1 Operand Number} {, Argument } {2 Operand Number} { Function Stop} { Function Stop}]"},
		{"=TRUE", "[{TRUE Operand Logical}]"},
		{"=--1-1", "[{- OperatorPrefix } {- OperatorPrefix } {1 Operand Number} {- OperatorInfix Math} {1 Operand Number}]"},
		{"=1 .  +\" \"", "[{1 Operand Number} { OperatorInfix Intersection} {. Operand Range} {+ OperatorInfix Math} {  Operand Text}]"},
		{"=1 ''", "[{1 Operand Number}]"},
		{"=10*2^(2*(1+1))% (=10.28114; % has greater precedence than ^)", "[{10 Operand Number} {* OperatorInfix Math} {2 Operand Number} {^ OperatorInfix Math} { Subexpression Start} {2 Operand Number} {* OperatorInfix Math} { Subexpression Start} {1 Operand Number} {+ OperatorInfix Math} {1 Operand Number} { Subexpression Stop} { Subexpression Stop} {% OperatorPostfix } { Subexpression Start} {= OperatorInfix Logical} {10.28114 Operand Number} {; Unknown } {% OperatorPostfix } {has Operand Range} { OperatorInfix Intersection} {greater Operand Range} { OperatorInfix Intersection} {precedence Operand Range} { OperatorInfix Intersection} {than Operand Range} {^ OperatorInfix Math} { Subexpression Stop}]"},
		{"=2+(10*2^(2*(1+1)+SUM(A2)))*3 (who knows, but you'll push and pop here multiple times)", "[{2 Operand Number} {+ OperatorInfix Math} { Subexpression Start} {10 Operand Number} {* OperatorInfix Math} {2 Operand Number} {^ OperatorInfix Math} { Subexpression Start} {2 Operand Number} {* OperatorInfix Math} { Subexpression Start} {1 Operand Number} {+ OperatorInfix Math} {1 Operand Number} { Subexpression Stop} {+ OperatorInfix Math} {SUM Function Start} {A2 Operand Range} { Function Stop} { Subexpression Stop} { Subexpression Stop} {* OperatorInfix Math} {3 Operand Number} { OperatorInfix Intersection} { Subexpression Start} {who Operand Range} { OperatorInfix Intersection} {knows Operand Range} {, OperatorInfix Union} {but Operand Range} {you Unknown } {ll push and pop here multiple times) Operand Range}]"},
		{"'C:\\Docs\\Book[A-B.XLS]C-D'!$BC14", "[{C:\\Docs\\Book[A-B.XLS]C-D!$BC14 Operand Range}]"},
		{"=\"a\"\"b\"", "[{a\"b Operand Text}]", "\"a\"\"b\""},
		{"=IF(P5=1.0,\"NA\",IF(P5=2.0,\"A\",IF(P5=3.0,\"B\",IF(P5=4.0,\"C\",IF(P5=5.0,\"D\",IF(P5=6.0,\"E\",IF(P5=7.0,\"F\",IF(P5=8.0,\"G\"))))))))", "[{IF Function Start} {P5 Operand Range} {= OperatorInfix Logical} {1.0 Operand Number} {, Argument } {NA Operand Text} {, Argument } {IF Function Start} {P5 Operand Range} {= OperatorInfix Logical} {2.0 Operand Number} {, Argument } {A Operand Text} {, Argument } {IF Function Start} {P5 Operand Range} {= OperatorInfix Logical} {3.0 Operand Number} {, Argument } {B Operand Text} {, Argument } {IF Function Start} {P5 Operand Range} {= OperatorInfix Logical} {4.0 Operand Number} {, Argument } {C Operand Text} {, Argument } {IF Function Start} {P5 Operand Range} {= OperatorInfix Logical} {5.0 Operand Number} {, Argument } {D Operand Text} {, Argument } {IF Function Start} {P5 Operand Range} {= OperatorInfix Logical} {6.0 Operand Number} {, Argument } {E Operand Text} {, Argument } {IF Function Start} {P5 Operand Range} {= OperatorInfix Logical} {7.0 Operand Number} {, Argument } {F Operand Text} {, Argument } {IF Function Start} {P5 Operand Range} {= OperatorInfix Logical} {8.0 Operand Number} {, Argument } {G Operand Text} { Function Stop} { Function Stop} { Function Stop} { Function Stop} { Function Stop} { Function Stop} { Function Stop} { Function Stop}]"},
		{"={SUM(B2:D2*B3:D3)}", "[{ARRAY Function Start} {ARRAYROW Function Start} {SUM Function Start} {B2:D2 Operand Range} {* OperatorInfix Math} {B3:D3 Operand Range} { Function Stop} { Function Stop} { Function Stop}]"},
		{"=SUM(123 + SUM(456) + (45<6))+456+789", "[{SUM Function Start} {123 Operand Number} {+ OperatorInfix Math} {SUM Function Start} {456 Operand Number} { Function Stop} {+ OperatorInfix Math} { Subexpression Start} {45 Operand Number} {< OperatorInfix Logical} {6 Operand Number} { Subexpression Stop} { Function Stop} {+ OperatorInfix Math} {456 Operand Number} {+ OperatorInfix Math} {789 Operand Number}]"},
		{"=AVG(((((123 + 4 + AVG(A1:A2))))))", "[{AVG Function Start} { Subexpression Start} { Subexpression Start} { Subexpression Start} { Subexpression Start} {123 Operand Number} {+ OperatorInfix Math} {4 Operand Number} {+ OperatorInfix Math} {AVG Function Start} {A1:A2 Operand Range} { Function Stop} { Subexpression Stop} { Subexpression Stop} { Subexpression Stop} { Subexpression Stop} { Function Stop}]"},
		{"=IF(\"a\"={\"a\",\"b\";\"c\",#N/A;-1,TRUE}, \"yes\", \"no\") &   \"  more \"\"test\"\" text\"", "[{IF Function Start} {a Operand Text} {= OperatorInfix Logical} {ARRAY Function Start} {ARRAYROW Function Start} {a Operand Text} {, Argument } {b Operand Text} { Function Stop} {, Argument } {ARRAYROW Function Start} {c Operand Text} {, Argument } {#N/A Operand Error} { Function Stop} {, Argument } {ARRAYROW Function Start} {- OperatorPrefix } {1 Operand Number} {, Argument } {TRUE Operand Logical} { Function Stop} { Function Stop} {, Argument } {yes Operand Text} {, Argument } {no Operand Text} { Function Stop} {& OperatorInfix Concatenation} {  more \"test\" text Operand Text}]"},
		{"=+ AName- (-+-+-2^6) = {\"A\",\"B\"} + @SUM(R1C1) + (@ERROR.TYPE(#VALUE!) = 2)", "[{+ OperatorInfix Math} {AName Operand Range} {- OperatorInfix Math} { Subexpression Start} {- OperatorPrefix } {- OperatorPrefix } {- OperatorPrefix } {2 Operand Number} {^ OperatorInfix Math} {6 Operand Number} { Subexpression Stop} {= OperatorInfix Logical} {ARRAY Function Start} {ARRAYROW Function Start} {A Operand Text} {, Argument } {B Operand Text} { Function Stop} { Function Stop} {+ OperatorInfix Math} {SUM Function Start} {R1C1 Operand Range} { Function Stop} {+ OperatorInfix Math} { Subexpression Start} {ERROR.TYPE Function Start} {#VALUE! Operand Error} { Function Stop} {= OperatorInfix Logical} {2 Operand Number} { Subexpression Stop}]"},
		{"=IF(R13C3>DATE(2002,1,6),0,IF(ISERROR(R[41]C[2]),0,IF(R13C3>=R[41]C[2],0, IF(AND(R[23]C[11]>=55,R[24]C[11]>=20),R53C3,0))))", "[{IF Function Start} {R13C3 Operand Range} {> OperatorInfix Logical} {DATE Function Start} {2002 Operand Number} {, Argument } {1 Operand Number} {, Argument } {6 Operand Number} { Function Stop} {, Argument } {0 Operand Number} {, Argument } {IF Function Start} {ISERROR Function Start} {R[41]C[2] Operand Range} { Function Stop} {, Argument } {0 Operand Number} {, Argument } {IF Function Start} {R13C3 Operand Range} {>= OperatorInfix Logical} {R[41]C[2] Operand Range} {, Argument } {0 Operand Number} {, Argument } {IF Function Start} {AND Function Start} {R[23]C[11] Operand Range} {>= OperatorInfix Logical} {55 Operand Number} {, Argument } {R[24]C[11] Operand Range} {>= OperatorInfix Logical} {20 Operand Number} { Function Stop} {, Argument } {R53C3 Operand Range} {, Argument } {0 Operand Number} { Function Stop} { Function Stop} { Function Stop} { Function Stop}]"},
	}
	for _, item := range testStrings {
		p := ExcelParser()
		result := fmt.Sprint(p.Parse(item[0]))
		if result != item[1] {
			t.Errorf("not equal:\nexpected:%+v\nactual  :%+v\nmessages:%s", item[1], result, item[0])
		}
		if len(item) == 3 { // Test Render
			r := p.Render()
			if r != item[2] {
				t.Errorf("Render not equal:\nexpected:%+v\nactual  :%+v\nmessages:%s", item[2], r, item[0])
			}
		}
		t.Log("Pretty printed:\n", p.PrettyPrint())
		t.Log("----------------------------------------")
		t.Log("Render printed:\n", p.Render())
		p.Tokens.tp()
		p.Tokens.value()
		p.Tokens.subtype()
	}
	tk := Tokens{Index: -1}
	tk.current()
	tk.next()
	tk.previous()
}
