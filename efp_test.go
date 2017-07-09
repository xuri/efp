package efp

import "testing"

func TestEFP(t *testing.T) {
	formulae := []string{
		// Simple test formulae
		`=1+3+5`,
		`=3 * 4 + 5`,
		`=50`,
		`=1+1`,
		`=$A1`,
		`=$B$2`,
		`=SUM(B5:B15)`,
		`=SUM(B5:B15,D5:D15)`,
		`=SUM(B5:B15 A7:D7)`,
		`=SUM(sheet1!$A$1:$B$2)`,
		`=[data.xls]sheet1!$A$1`,
		`=SUM((A:A 1:1))`,
		`=SUM((A:A,1:1))`,
		`=SUM((A:A A1:B1))`,
		`=SUM(D9:D11,E9:E11,F9:F11)`,
		`=SUM((D9:D11,(E9:E11,F9:F11)))`,
		`=((D2 * D3) + D4) & " should be 10"`,
		// E. W. Bachtal's test formulae
		`=IF(P5=1.0,"NA",IF(P5=2.0,"A",IF(P5=3.0,"B",IF(P5=4.0,"C",IF(P5=5.0,"D",IF(P5=6.0,"E",IF(P5=7.0,"F",IF(P5=8.0,"G"))))))))`,
		`={SUM(B2:D2*B3:D3)}`,
		`=SUM(123 + SUM(456) + (45<6))+456+789`,
		`=AVG(((((123 + 4 + AVG(A1:A2))))))`,
		`=IF("a"={"a","b";"c",#N/A;-1,TRUE}, "yes", "no") &   "  more ""test"" text"`,
		`=+ AName- (-+-+-2^6) = {"A","B"} + @SUM(R1C1) + (@ERROR.TYPE(#VALUE!) = 2)`,
		`=IF(R13C3>DATE(2002,1,6),0,IF(ISERROR(R[41]C[2]),0,IF(R13C3>=R[41]C[2],0, IF(AND(R[23]C[11]>=55,R[24]C[11]>=20),R53C3,0))))`,
	}
	for _, f := range formulae {
		p := ExcelParser()
		t.Log("========================================")
		t.Log("Formula:     ", f)
		p.Parse(f)
		t.Log("Pretty printed:\n", p.PrettyPrint())
		t.Log("----------------------------------------")
		t.Log("Render printed:\n", p.Render())
		p.Tokens.tp()
		p.Tokens.value()
		p.Tokens.subtype()
	}
}
