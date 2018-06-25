//line pkg/parser/labels.y:2
package parser

import __yyfmt__ "fmt"

//line pkg/parser/labels.y:2
import (
	"github.com/prometheus/prometheus/pkg/labels"
)

//line pkg/parser/labels.y:9
type labelsSymType struct {
	yys          int
	MatchersExpr []*labels.Matcher
	Matchers     []*labels.Matcher
	Matcher      *labels.Matcher
	LabelsExpr   labels.Labels
	Labels       labels.Labels
	Label        labels.Label
	str          string
	int          int64
	Identifier   string
}

const IDENTIFIER = 57346
const STRING = 57347
const MATCHERS = 57348
const LABELS = 57349
const EQ = 57350
const NEQ = 57351
const RE = 57352
const NRE = 57353
const OPEN_BRACE = 57354
const CLOSE_BRACE = 57355
const COMMA = 57356
const DOT = 57357

var labelsToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"IDENTIFIER",
	"STRING",
	"MATCHERS",
	"LABELS",
	"EQ",
	"NEQ",
	"RE",
	"NRE",
	"OPEN_BRACE",
	"CLOSE_BRACE",
	"COMMA",
	"DOT",
}
var labelsStatenames = [...]string{}

const labelsEofCode = 1
const labelsErrCode = 2
const labelsInitialStackSize = 16

//line pkg/parser/labels.y:63

//line yacctab:1
var labelsExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const labelsPrivate = 57344

const labelsLast = 32

var labelsAct = [...]int{

	8, 11, 15, 16, 17, 18, 12, 22, 5, 19,
	20, 21, 7, 4, 19, 13, 14, 2, 3, 30,
	27, 26, 12, 29, 25, 24, 9, 23, 28, 10,
	6, 1,
}
var labelsPact = [...]int{

	11, -1000, 1, -4, 22, 22, 2, -1000, -6, -1000,
	-3, -1000, -1, -1000, 22, 20, 19, 16, 15, 24,
	-1000, 22, 14, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000,
}
var labelsPgo = [...]int{

	0, 31, 30, 12, 29, 1, 0,
}
var labelsR1 = [...]int{

	0, 1, 1, 2, 2, 3, 3, 3, 3, 4,
	4, 5, 6, 6,
}
var labelsR2 = [...]int{

	0, 4, 4, 1, 3, 3, 3, 3, 3, 1,
	3, 3, 1, 3,
}
var labelsChk = [...]int{

	-1000, -1, 6, 7, 12, 12, -2, -3, -6, 4,
	-4, -5, -6, 13, 14, 8, 9, 10, 11, 15,
	13, 14, 8, -3, 5, 5, 5, 5, 4, -5,
	5,
}
var labelsDef = [...]int{

	0, -2, 0, 0, 0, 0, 0, 3, 0, 12,
	0, 9, 0, 1, 0, 0, 0, 0, 0, 0,
	2, 0, 0, 4, 5, 6, 7, 8, 13, 10,
	11,
}
var labelsTok1 = [...]int{

	1,
}
var labelsTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15,
}
var labelsTok3 = [...]int{
	0,
}

var labelsErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	labelsDebug        = 0
	labelsErrorVerbose = false
)

type labelsLexer interface {
	Lex(lval *labelsSymType) int
	Error(s string)
}

type labelsParser interface {
	Parse(labelsLexer) int
	Lookahead() int
}

type labelsParserImpl struct {
	lval  labelsSymType
	stack [labelsInitialStackSize]labelsSymType
	char  int
}

func (p *labelsParserImpl) Lookahead() int {
	return p.char
}

func labelsNewParser() labelsParser {
	return &labelsParserImpl{}
}

const labelsFlag = -1000

func labelsTokname(c int) string {
	if c >= 1 && c-1 < len(labelsToknames) {
		if labelsToknames[c-1] != "" {
			return labelsToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func labelsStatname(s int) string {
	if s >= 0 && s < len(labelsStatenames) {
		if labelsStatenames[s] != "" {
			return labelsStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func labelsErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !labelsErrorVerbose {
		return "syntax error"
	}

	for _, e := range labelsErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + labelsTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := labelsPact[state]
	for tok := TOKSTART; tok-1 < len(labelsToknames); tok++ {
		if n := base + tok; n >= 0 && n < labelsLast && labelsChk[labelsAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if labelsDef[state] == -2 {
		i := 0
		for labelsExca[i] != -1 || labelsExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; labelsExca[i] >= 0; i += 2 {
			tok := labelsExca[i]
			if tok < TOKSTART || labelsExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if labelsExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += labelsTokname(tok)
	}
	return res
}

func labelslex1(lex labelsLexer, lval *labelsSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = labelsTok1[0]
		goto out
	}
	if char < len(labelsTok1) {
		token = labelsTok1[char]
		goto out
	}
	if char >= labelsPrivate {
		if char < labelsPrivate+len(labelsTok2) {
			token = labelsTok2[char-labelsPrivate]
			goto out
		}
	}
	for i := 0; i < len(labelsTok3); i += 2 {
		token = labelsTok3[i+0]
		if token == char {
			token = labelsTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = labelsTok2[1] /* unknown char */
	}
	if labelsDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", labelsTokname(token), uint(char))
	}
	return char, token
}

func labelsParse(labelslex labelsLexer) int {
	return labelsNewParser().Parse(labelslex)
}

func (labelsrcvr *labelsParserImpl) Parse(labelslex labelsLexer) int {
	var labelsn int
	var labelsVAL labelsSymType
	var labelsDollar []labelsSymType
	_ = labelsDollar // silence set and not used
	labelsS := labelsrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	labelsstate := 0
	labelsrcvr.char = -1
	labelstoken := -1 // labelsrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		labelsstate = -1
		labelsrcvr.char = -1
		labelstoken = -1
	}()
	labelsp := -1
	goto labelsstack

ret0:
	return 0

ret1:
	return 1

labelsstack:
	/* put a state and value onto the stack */
	if labelsDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", labelsTokname(labelstoken), labelsStatname(labelsstate))
	}

	labelsp++
	if labelsp >= len(labelsS) {
		nyys := make([]labelsSymType, len(labelsS)*2)
		copy(nyys, labelsS)
		labelsS = nyys
	}
	labelsS[labelsp] = labelsVAL
	labelsS[labelsp].yys = labelsstate

labelsnewstate:
	labelsn = labelsPact[labelsstate]
	if labelsn <= labelsFlag {
		goto labelsdefault /* simple state */
	}
	if labelsrcvr.char < 0 {
		labelsrcvr.char, labelstoken = labelslex1(labelslex, &labelsrcvr.lval)
	}
	labelsn += labelstoken
	if labelsn < 0 || labelsn >= labelsLast {
		goto labelsdefault
	}
	labelsn = labelsAct[labelsn]
	if labelsChk[labelsn] == labelstoken { /* valid shift */
		labelsrcvr.char = -1
		labelstoken = -1
		labelsVAL = labelsrcvr.lval
		labelsstate = labelsn
		if Errflag > 0 {
			Errflag--
		}
		goto labelsstack
	}

labelsdefault:
	/* default state action */
	labelsn = labelsDef[labelsstate]
	if labelsn == -2 {
		if labelsrcvr.char < 0 {
			labelsrcvr.char, labelstoken = labelslex1(labelslex, &labelsrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if labelsExca[xi+0] == -1 && labelsExca[xi+1] == labelsstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			labelsn = labelsExca[xi+0]
			if labelsn < 0 || labelsn == labelstoken {
				break
			}
		}
		labelsn = labelsExca[xi+1]
		if labelsn < 0 {
			goto ret0
		}
	}
	if labelsn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			labelslex.Error(labelsErrorMessage(labelsstate, labelstoken))
			Nerrs++
			if labelsDebug >= 1 {
				__yyfmt__.Printf("%s", labelsStatname(labelsstate))
				__yyfmt__.Printf(" saw %s\n", labelsTokname(labelstoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for labelsp >= 0 {
				labelsn = labelsPact[labelsS[labelsp].yys] + labelsErrCode
				if labelsn >= 0 && labelsn < labelsLast {
					labelsstate = labelsAct[labelsn] /* simulate a shift of "error" */
					if labelsChk[labelsstate] == labelsErrCode {
						goto labelsstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if labelsDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", labelsS[labelsp].yys)
				}
				labelsp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if labelsDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", labelsTokname(labelstoken))
			}
			if labelstoken == labelsEofCode {
				goto ret1
			}
			labelsrcvr.char = -1
			labelstoken = -1
			goto labelsnewstate /* try again in the same state */
		}
	}

	/* reduction by production labelsn */
	if labelsDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", labelsn, labelsStatname(labelsstate))
	}

	labelsnt := labelsn
	labelspt := labelsp
	_ = labelspt // guard against "declared and not used"

	labelsp -= labelsR2[labelsn]
	// labelsp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if labelsp+1 >= len(labelsS) {
		nyys := make([]labelsSymType, len(labelsS)*2)
		copy(nyys, labelsS)
		labelsS = nyys
	}
	labelsVAL = labelsS[labelsp+1]

	/* consult goto table to find next state */
	labelsn = labelsR1[labelsn]
	labelsg := labelsPgo[labelsn]
	labelsj := labelsg + labelsS[labelsp].yys + 1

	if labelsj >= labelsLast {
		labelsstate = labelsAct[labelsg]
	} else {
		labelsstate = labelsAct[labelsj]
		if labelsChk[labelsstate] != -labelsn {
			labelsstate = labelsAct[labelsg]
		}
	}
	// dummy call; replaced with literal code
	switch labelsnt {

	case 1:
		labelsDollar = labelsS[labelspt-4 : labelspt+1]
		//line pkg/parser/labels.y:35
		{
			labelslex.(*lexer).matcher = labelsDollar[3].Matchers
		}
	case 2:
		labelsDollar = labelsS[labelspt-4 : labelspt+1]
		//line pkg/parser/labels.y:36
		{
			labelslex.(*lexer).labels = labelsDollar[3].Labels
		}
	case 3:
		labelsDollar = labelsS[labelspt-1 : labelspt+1]
		//line pkg/parser/labels.y:39
		{
			labelsVAL.Matchers = []*labels.Matcher{labelsDollar[1].Matcher}
		}
	case 4:
		labelsDollar = labelsS[labelspt-3 : labelspt+1]
		//line pkg/parser/labels.y:40
		{
			labelsVAL.Matchers = append(labelsDollar[1].Matchers, labelsDollar[3].Matcher)
		}
	case 5:
		labelsDollar = labelsS[labelspt-3 : labelspt+1]
		//line pkg/parser/labels.y:44
		{
			labelsVAL.Matcher = mustNewMatcher(labels.MatchEqual, labelsDollar[1].Identifier, labelsDollar[3].str)
		}
	case 6:
		labelsDollar = labelsS[labelspt-3 : labelspt+1]
		//line pkg/parser/labels.y:45
		{
			labelsVAL.Matcher = mustNewMatcher(labels.MatchNotEqual, labelsDollar[1].Identifier, labelsDollar[3].str)
		}
	case 7:
		labelsDollar = labelsS[labelspt-3 : labelspt+1]
		//line pkg/parser/labels.y:46
		{
			labelsVAL.Matcher = mustNewMatcher(labels.MatchRegexp, labelsDollar[1].Identifier, labelsDollar[3].str)
		}
	case 8:
		labelsDollar = labelsS[labelspt-3 : labelspt+1]
		//line pkg/parser/labels.y:47
		{
			labelsVAL.Matcher = mustNewMatcher(labels.MatchNotRegexp, labelsDollar[1].Identifier, labelsDollar[3].str)
		}
	case 9:
		labelsDollar = labelsS[labelspt-1 : labelspt+1]
		//line pkg/parser/labels.y:51
		{
			labelsVAL.Labels = labels.Labels{labelsDollar[1].Label}
		}
	case 10:
		labelsDollar = labelsS[labelspt-3 : labelspt+1]
		//line pkg/parser/labels.y:52
		{
			labelsVAL.Labels = append(labelsDollar[1].Labels, labelsDollar[3].Label)
		}
	case 11:
		labelsDollar = labelsS[labelspt-3 : labelspt+1]
		//line pkg/parser/labels.y:56
		{
			labelsVAL.Label = labels.Label{Name: labelsDollar[1].Identifier, Value: labelsDollar[3].str}
		}
	case 12:
		labelsDollar = labelsS[labelspt-1 : labelspt+1]
		//line pkg/parser/labels.y:60
		{
			labelsVAL.Identifier = labelsDollar[1].str
		}
	case 13:
		labelsDollar = labelsS[labelspt-3 : labelspt+1]
		//line pkg/parser/labels.y:61
		{
			labelsVAL.Identifier = labelsDollar[1].Identifier + "." + labelsDollar[3].str
		}
	}
	goto labelsstack /* stack new state and value */
}
