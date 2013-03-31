

/*2:*/


//line goweave.w:66

package main

import(


/*13:*/


//line gocommon.w:93

"io"
"bytes"



/*:13*/



/*16:*/


//line gocommon.w:135

"bufio"



/*:16*/



/*20:*/


//line gocommon.w:175

"unicode"



/*:20*/



/*27:*/


//line gocommon.w:330

"fmt"



/*:27*/



/*34:*/


//line gocommon.w:457

"os"
"strings"



/*:34*/


//line goweave.w:70

)



/*93:*/


//line goweave.w:178

type xref_info struct{
num int32/* section number plus zero or def_flag */
xlink int32/* index of the previous cross-reference */
}



/*:93*/



/*172:*/


//line goweave.w:2151

type scrap struct{
cat int32
mathness int32
trans[]interface{}


/*365:*/


//line goweave.w:6351

head int32



/*:365*/


//line goweave.w:2156

}




/*:172*/



/*174:*/


//line goweave.w:2180

type id_token int

type res_token int

type section_token int32

type list_token[]interface{}

type inner_list_token[]interface{}





/*:174*/



/*176:*/


//line goweave.w:2197

type reducing func()



/*:176*/



/*182:*/


//line goweave.w:2364

type pair struct{
cat int32
mand bool
}



/*:182*/



/*320:*/


//line goweave.w:5360

type mode int



/*:320*/



/*322:*/


//line goweave.w:5369

type output_state struct{
tok_field[]interface{}/* present location of token list */
mode_field mode/* interpretation of control tokens */
}



/*:322*/


//line goweave.w:73




/*1:*/


//line goweave.w:62

const banner= "This is GOWEAVE (Version 0.6)\n"



/*:1*/



/*4:*/


//line goweave.w:105

const(
max_names= 4000/* number of identifiers, strings, section names
		must be less than 10240*/
line_length= 80/* lines of \TEX/ output have at most this many characters
		should be less than 256 */
)



/*:4*/



/*96:*/


//line goweave.w:192

const(
cite_flag= 10240
file_flag= 3*cite_flag
def_flag= 2*cite_flag
)



/*:96*/



/*111:*/


//line goweave.w:436

const(
ignore rune= 00/* control code of no interest to \.{GOWEAVE} */
underline rune= '\n'/* this code will be intercepted without confusion */
noop rune= 0177/* takes the place of ASCII delete */
xref_roman rune= 0213/* control code for `\.{@\^}' */
xref_wildcard rune= 0214/* control code for `\.{@:}' */
xref_typewriter rune= 0215/* control code for `\.{@.}' */
TeX_string rune= 0216/* control code for `\.{@t}' */
ord rune= 0217/* control code for `\.{@'}' */
join rune= 0220/* control code for `\.{@\&}' */
thin_space rune= 0221/* control code for `\.{@,}' */
math_break rune= 0222/* control code for `\.{@\v}' */
line_break rune= 0223/* control code for `\.{@/}' */
big_line_break rune= 0224/* control code for `\.{@\#}' */
no_line_break rune= 0225/* control code for `\.{@+}' */
pseudo_semi rune= 0226/* control code for `\.{@;}' */
verbatim rune= 0227/* control code for `\.{@=}' */
trace rune= 0232/* control code for `\.{@0}', `\.{@1}' and `\.{@2}' */
format_code rune= 0235/* control code for `\.{@f}' and `\.{@s}' */
begin_code rune= 0237/* control code for `\.{@c}' */
section_name rune= 0240/* control code for `\.{@<}' */
new_section rune= 0241/* control code for `\.{@\ }' and `\.{@*}' */
)



/*:111*/



/*119:*/


//line goweave.w:636

const(
constant rune= 0210/* \GO/ constant */
str rune= 0211/* \GO/ string */
identifier rune= 0212/* \GO/ identifier or reserved word */
)



/*:119*/



/*166:*/


//line goweave.w:1777

const(
normal rune= iota/* ordinary identifiers have normal ilk */
roman rune= iota/* normal index entries have roman ilk */
wildcard rune= iota/* user-formatted index entries have wildcard ilk */
typewriter rune= iota/* `typewriter type' entries have typewriter ilk */
custom rune= iota/* identifiers with user-given control sequence */
)

const(
zero rune= iota
ArrayType rune= iota
StructType rune= iota
PointerType rune= iota
InterfaceType rune= iota
SliceType rune= iota
MapType rune= iota
ChannelType rune= iota
FieldDecl rune= iota
AnonymousField rune= iota
Signature rune= iota
Parameters rune= iota
ParameterList rune= iota
ParameterDecl rune= iota
MethodSpec rune= iota
Block rune= iota
Statement rune= iota
ConstDecl rune= iota
TypeDecl rune= iota
VarDecl rune= iota
FunctionDecl rune= iota
MethodDecl rune= iota
ConstSpec rune= iota
IdentifierList rune= iota
ExpressionList rune= iota
TypeSpec rune= iota
VarSpec rune= iota
ShortVarDecl rune= iota
Receiver rune= iota
Operand rune= iota
QualifiedIdent rune= iota
MethodExpr rune= iota
CompositeLit rune= iota
FunctionLit rune= iota
FunctionType rune= iota
LiteralType rune= iota
LiteralValue rune= iota
ElementList rune= iota
Element rune= iota
PrimaryExpr rune= iota
Conversion rune= iota
BuiltinCall rune= iota
Selector rune= iota
Index rune= iota
Slice rune= iota
TypeAssertion rune= iota
Call rune= iota
Expression rune= iota
UnaryExpr rune= iota
ReceiverType rune= iota
LabeledStmt rune= iota
SimpleStmt rune= iota
GoStmt rune= iota
ReturnStmt rune= iota
BreakStmt rune= iota
ContinueStmt rune= iota
GotoStmt rune= iota
fallthrough_token rune= iota
IfStmt rune= iota
SelectStmt rune= iota
ForStmt rune= iota
DeferStmt rune= iota
SendStmt rune= iota
IncDecStmt rune= iota
Assignment rune= iota
ExprSwitchStmt rune= iota
ExprCaseClause rune= iota
TypeSwitchStmt rune= iota
TypeSwitchGuard rune= iota
TypeCaseClause rune= iota
TypeSwitchCase rune= iota
ForClause rune= iota
RangeClause rune= iota
CommClause rune= iota
CommCase rune= iota
RecvStmt rune= iota
BuiltinArgs rune= iota
PackageClause rune= iota
ImportDecl rune= iota
ImportSpec rune= iota
Type rune= iota
package_token rune= iota/* denotes \.{package}*/
import_token rune= iota/* denotes \&{import} */
type_token rune= iota/* \&{type} */
interface_token rune= iota/* \&{interface} */
const_token rune= iota/* \&{const} */
go_token rune= iota/* \&{go} */
return_token rune= iota/* \&{return} */
break_token rune= iota/* \&{break} */
continue_token rune= iota/* \&{continue} */
goto_token rune= iota/* \&{goto} */
if_token rune= iota/* \&{if} */
switch_token rune= iota/* \&{switch} */
select_token rune= iota/* \&{select} */
case_token rune= iota/* \&{case} */
default_token rune= iota/* \&{default} */
for_token rune= iota/* \&{for}*/
else_token rune= iota/* \&{else} */
defer_token rune= iota/* denotes \.{defer} and \.{go} statements*/
func_token rune= iota/* denotes a function declarator */
struct_token rune= iota/* \&{struct} */
var_token rune= iota/* \&{var} */
range_token rune= iota/* \&{range} */
map_token rune= iota/* \&{map} */
chan_token rune= iota/* \&{chan} */
dot rune= iota/* \&{.} */
eq rune= iota/* denotes an assign operator '=' */
binary_op rune= iota
rel_op rune= iota
add_op rune= iota
mul_op rune= iota
unary_op rune= iota
asterisk rune= iota
assign_op rune= iota

lbrace rune= iota/* denotes a left brace */
rbrace rune= iota/* denotes a right brace */
comma rune= iota/* denotes a comma */
lpar rune= iota/* denotes a left parenthesis */
rpar rune= iota/* denotes a right parenthesis */
lbracket rune= iota/* denotes a left bracket */
rbracket rune= iota/* denotes a right bracket */

semi rune= iota/* denotes a semicolon */
colon rune= iota/* denotes a colon */
insert rune= iota/* a scrap that gets combined with its neighbor */
section_scrap rune= iota/* section name */
dead rune= iota/* scrap that won't combine */
)



/*:166*/



/*170:*/


//line goweave.w:2121

const(
math_rel rune= 0244
big_cancel rune= 0245/* like cancel, also overrides spaces */
cancel rune= 0246/* overrides backup, break_space, force, big_force */
indent rune= 0247/* one more tab (\.{\\1}) */
outdent rune= 0250/* one less tab (\.{\\2}) */
opt rune= 0251/* optional break in mid-statement (\.{\\3}) */
backup rune= 0252/* stick out one unit to the left (\.{\\4}) */
break_space rune= 0253/* optional break between statements (\.{\\5}) */
force rune= 0254/* forced break between statements (\.{\\6}) */
big_force rune= 0255/* forced break with additional space (\.{\\7}) */
quoted_char rune= 0256/* introduces a character token in the range 0200--0377 */
end_translation rune= 0257/* special sentinel token at end of list */
inserted rune= 0260/* sentinel to mark translations of inserts */
)



/*:170*/



/*299:*/


//line goweave.w:4755

const(
maybe_math rune= iota/* works in either horizontal or math mode */
yes_math rune= iota/* should be in math mode */
no_math rune= iota/* should be in horizontal mode */
)



/*:299*/



/*321:*/


//line goweave.w:5363

const(
inner mode= 0/* value of mode for \GO/ texts within \TEX/ texts */
outer mode= 1/* value of mode for \GO/ texts in sections */
)



/*:321*/



/*328:*/


//line goweave.w:5421

const(
res_word rune= 0242/* returned by get_output for reserved words */
section_code rune= 0243/* returned by get_output for section names */
)



/*:328*/



/*370:*/


//line goweave.w:6401

const infinity= -1/* $\infty$ (approximately) */



/*:370*/


//line goweave.w:75




/*92:*/


//line goweave.w:154

var change_exists bool/* has any section changed? */



/*:92*/



/*94:*/


//line goweave.w:184

var xmem[]xref_info/* contains cross-reference information */
var xref_switch int32
var section_xref_switch int32/* either zero or def_flag */



/*:94*/



/*113:*/


//line goweave.w:466

var ccode[256]rune/* meaning of a char following \.{@} */



/*:113*/



/*120:*/


//line goweave.w:643

var cur_section int32/* name of section just scanned */
var cur_section_char rune/* the character just before that name */




/*:120*/



/*133:*/


//line goweave.w:1010

var next_control rune/* control code waiting to be acting upon */



/*:133*/



/*140:*/


//line goweave.w:1180

var lhs int32
var rhs int32/* pointers to byte_start for format identifiers */
var res_wd_end int32



/*:140*/



/*145:*/


//line goweave.w:1279

var cur_xref int32/* temporary cross-reference pointer */
var an_output bool/* did file_flag precede cur_xref? */



/*:145*/



/*148:*/


//line goweave.w:1329

var out_buf[line_length+1]rune/* assembled characters */
var out_ptr int32/* just after last character in out_buf */
var out_buf_end int32= line_length/* end of out_buf */
var out_line int/* number of next line to be output */



/*:148*/



/*167:*/


//line goweave.w:1917

var cat_name[256]string



/*:167*/



/*173:*/


//line goweave.w:2160

var scrap_info[]scrap/* memory array for scraps */



/*:173*/



/*177:*/


//line goweave.w:2201

var shift= 0
var empty reducing= func(){}



/*:177*/



/*303:*/


//line goweave.w:4836

var tracing int32/* can be used to show parsing details */



/*:303*/



/*324:*/


//line goweave.w:5381

var cur_state output_state/* cur_state.tok_field, cur_state.mode_field */
var stack[]output_state/* info for non-current levels */



/*:324*/



/*327:*/


//line goweave.w:5418

var cur_name int32= -1



/*:327*/



/*344:*/


//line goweave.w:5883

var save_line int/* former value of out_line */
var save_place int32/* former value of out_ptr */
var sec_depth int32/* the integer, if any, following \.{@*} */
var space_checked bool/* have we done emit_space_if_needed? */
var format_visible bool/* should the next format declaration be output? */
var doing_format bool= false/* are we outputting a format declaration? */
var group_found bool= false/* has a starred section occurred? */



/*:344*/



/*351:*/


//line goweave.w:6058

var this_section int32/* the current section name, or zero */



/*:351*/



/*362:*/


//line goweave.w:6315

var bucket[256]int32
var blink[max_names]int32/* links in the buckets */



/*:362*/



/*366:*/


//line goweave.w:6354

var cur_depth int32/* depth of current buckets */
var cur_byte int32/* index into byte_mem */
var cur_val int32/* current cross-reference number */
var max_sort_ptr int32/* largest value of sort_ptr */
var sort_ptr int32/* ditto */



/*:366*/



/*368:*/


//line goweave.w:6367

/* collation order */
var collate= [102+128]rune{
0,' ',001,002,003,004,005,006,007,010,011,012,013,014,015,016,017,
020,021,022,023,024,025,026,027,030,031,032,033,034,035,036,037,
'!',042,'#','$','%','&','\'','(',')','*','+',',','-','.','/',':',
';','<','=','>','?','@','[','\\',']','^','`','{','|','}','~','_',
'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q',
'r','s','t','u','v','w','x','y','z','0','1','2','3','4','5','6','7','8','9',
0200,0201,0202,0203,0204,0205,0206,0207,0210,0211,0212,0213,0214,0215,0216,0217,
0220,0221,0222,0223,0224,0225,0226,0227,0230,0231,0232,0233,0234,0235,0236,0237,
0240,0241,0242,0243,0244,0245,0246,0247,0250,0251,0252,0253,0254,0255,0256,0257,
0260,0261,0262,0263,0264,0265,0266,0267,0270,0271,0272,0273,0274,0275,0276,0277,
0300,0301,0302,0303,0304,0305,0306,0307,0310,0311,0312,0313,0314,0315,0316,0317,
0320,0321,0322,0323,0324,0325,0326,0327,0330,0331,0332,0333,0334,0335,0336,0337,
0340,0341,0342,0343,0344,0345,0346,0347,0350,0351,0352,0353,0354,0355,0356,0357,
0360,0361,0362,0363,0364,0365,0366,0367,0370,0371,0372,0373,0374,0375,0376,0377}




/*:368*/



/*378:*/


//line goweave.w:6562

var next_xref int32
var this_xref int32
/* pointer variables for rearranging a list */



/*:378*/


//line goweave.w:77




/*:2*/



/*3:*/


//line goweave.w:84

func main(){
flags['c']= true
flags['x']= true
flags['f']= true
flags['e']= true/* controlled by command-line options */
common_init()


/*98:*/


//line goweave.w:202

xmem= append(xmem,xref_info{})
xref_switch= 0
section_xref_switch= 0



/*:98*/



/*114:*/


//line goweave.w:469

{
for c:=0;c<256;c++{
ccode[c]= ignore
}
}
ccode[' ']= new_section
ccode['\t']= new_section
ccode['\n']= new_section
ccode['\v']= new_section
ccode['\r']= new_section
ccode['\f']= new_section
ccode['*']= new_section
ccode['@']= '@'/* `quoted' at sign */
ccode['=']= verbatim
ccode['f']= format_code
ccode['F']= format_code
ccode['s']= format_code
ccode['S']= format_code
ccode['c']= begin_code
ccode['C']= begin_code
ccode['p']= begin_code
ccode['P']= begin_code
ccode['t']= TeX_string
ccode['T']= TeX_string
ccode['q']= noop
ccode['Q']= noop
ccode['&']= join
ccode['<']= section_name
ccode['(']= section_name
ccode['!']= underline
ccode['^']= xref_roman
ccode[':']= xref_wildcard
ccode['.']= xref_typewriter
ccode[',']= thin_space
ccode['|']= math_break
ccode['/']= line_break
ccode['#']= big_line_break
ccode['+']= no_line_break
ccode[';']= pseudo_semi
ccode['\'']= ord


/*115:*/


//line goweave.w:515

ccode['0']= trace// turn the tracing off
ccode['1']= trace// turn on a printing of irreducible scraps
ccode['2']= trace// turn on a printing of a snapshot of the scrap_info
ccode['4']= trace// turn on a printing of a category name is looking for
ccode['8']= trace// turn on a printing of a resulting translation of a scrap
ccode['3']= trace
ccode['5']= trace
ccode['6']= trace
ccode['7']= trace
ccode['9']= trace



/*:115*/


//line goweave.w:510




/*:114*/



/*151:*/


//line goweave.w:1406

out_ptr= 1
out_line= 1
active_file= tex_file
out_buf[out_ptr]= 'c'
fmt.Fprint(active_file,"\\input gowebma")



/*:151*/



/*154:*/


//line goweave.w:1443

out_buf[0]= '\\'



/*:154*/



/*168:*/


//line goweave.w:1920

for cat_index:=0;cat_index<255;cat_index++{
cat_name[cat_index]= "UNKNOWN-"+fmt.Sprintf("%v",cat_index)
}


cat_name[Type]= "Type"
cat_name[ArrayType]= "ArrayType"
cat_name[StructType]= "StructType"
cat_name[PointerType]= "PointerType"
cat_name[InterfaceType]= "InterfaceType"
cat_name[SliceType]= "SliceType"
cat_name[MapType]= "MapType"
cat_name[ChannelType]= "ChannelType"
cat_name[FieldDecl]= "FieldDecl"
cat_name[AnonymousField]= "AnonymousField"
cat_name[Signature]= "Signature"
cat_name[Parameters]= "Parameters"
cat_name[ParameterList]= "ParameterList"
cat_name[ParameterDecl]= "ParameterDecl"
cat_name[MethodSpec]= "MethodSpec"
cat_name[Block]= "Block"
cat_name[Statement]= "Statement"
cat_name[ConstDecl]= "ConstDecl"
cat_name[TypeDecl]= "TypeDecl"
cat_name[VarDecl]= "VarDecl"
cat_name[FunctionDecl]= "FunctionDecl"
cat_name[MethodDecl]= "MethodDecl"
cat_name[ConstSpec]= "ConstSpec"
cat_name[IdentifierList]= "IdentifierList"
cat_name[ExpressionList]= "ExpressionList"
cat_name[TypeSpec]= "TypeSpec"
cat_name[VarSpec]= "VarSpec"
cat_name[ShortVarDecl]= "ShortVarDecl"
cat_name[Receiver]= "Receiver"
cat_name[Operand]= "Operand"
cat_name[QualifiedIdent]= "QualifiedIdent"
cat_name[MethodExpr]= "MethodExpr"
cat_name[CompositeLit]= "CompositeLit"
cat_name[FunctionLit]= "FunctionLit"
cat_name[FunctionType]= "FunctionType"
cat_name[LiteralType]= "LiteralType"
cat_name[LiteralValue]= "LiteralValue"
cat_name[ElementList]= "ElementList"
cat_name[Element]= "Element"
cat_name[PrimaryExpr]= "PrimaryExpr"
cat_name[Conversion]= "Conversion"
cat_name[BuiltinCall]= "BuiltinCall"
cat_name[Selector]= "Selector"
cat_name[Index]= "Index"
cat_name[Slice]= "Slice"
cat_name[TypeAssertion]= "TypeAssertion"
cat_name[Call]= "Call"
cat_name[Expression]= "Expression"
cat_name[UnaryExpr]= "UnaryExpr"
cat_name[ReceiverType]= "ReceiverType"
cat_name[LabeledStmt]= "LabeledStmt"
cat_name[SimpleStmt]= "SimpleStmt"
cat_name[GoStmt]= "GoStmt"
cat_name[ReturnStmt]= "ReturnStmt"
cat_name[BreakStmt]= "BreakStmt"
cat_name[ContinueStmt]= "ContinueStmt"
cat_name[GotoStmt]= "GotoStmt"
cat_name[fallthrough_token]= "fallthrough_token"
cat_name[IfStmt]= "IfStmt"
cat_name[SelectStmt]= "SelectStmt"
cat_name[ForStmt]= "ForStmt"
cat_name[DeferStmt]= "DeferStmt"
cat_name[SendStmt]= "SendStmt"
cat_name[IncDecStmt]= "IncDecStmt"
cat_name[Assignment]= "Assignment"
cat_name[ExprSwitchStmt]= "ExprSwitchStmt"
cat_name[ExprCaseClause]= "ExprCaseClause"
cat_name[TypeSwitchStmt]= "TypeSwitchStmt"
cat_name[TypeSwitchGuard]= "TypeSwitchGuard"
cat_name[TypeCaseClause]= "TypeCaseClause"
cat_name[TypeSwitchCase]= "TypeSwitchCase"
cat_name[ForClause]= "ForClause"
cat_name[RangeClause]= "RangeClause"
cat_name[CommClause]= "CommClause"
cat_name[CommCase]= "CommCase"
cat_name[RecvStmt]= "RecvStmt"
cat_name[BuiltinArgs]= "BuiltinArgs"
cat_name[PackageClause]= "PackageClause"
cat_name[ImportDecl]= "ImportDecl"
cat_name[ImportSpec]= "ImportSpec"

cat_name[package_token]= "package"
cat_name[import_token]= "import"
cat_name[type_token]= "type"
cat_name[interface_token]= "interface"
cat_name[const_token]= "const"
cat_name[go_token]= "go"
cat_name[return_token]= "return"
cat_name[break_token]= "break"
cat_name[continue_token]= "continue"
cat_name[goto_token]= "goto"
cat_name[if_token]= "if"
cat_name[switch_token]= "switch"
cat_name[select_token]= "select"
cat_name[case_token]= "case"
cat_name[default_token]= "default"
cat_name[for_token]= "for"
cat_name[else_token]= "else"
cat_name[defer_token]= "defer"
cat_name[func_token]= "func"
cat_name[struct_token]= "struct"
cat_name[var_token]= "var"
cat_name[range_token]= "range"
cat_name[map_token]= "map"
cat_name[chan_token]= "chan"

cat_name[dot]= "'.'"

cat_name[eq]= "'='"
cat_name[col_eq]= "':='"
cat_name[binary_op]= "binary_op"
cat_name[rel_op]= "rel_op"
cat_name[add_op]= "add_op"
cat_name[mul_op]= "mul_op"
cat_name[unary_op]= "unary_op"
cat_name[asterisk]= "'*'"
cat_name[assign_op]= "assign_op"

cat_name[lbrace]= "'{'"
cat_name[rbrace]= "'}'"
cat_name[comma]= "','"
cat_name[lpar]= "'('"
cat_name[rpar]= "')'"
cat_name[lbracket]= "'['"
cat_name[rbracket]= "']'"
cat_name[semi]= "';'"
cat_name[colon]= "':'"
cat_name[insert]= "insert"
cat_name[section_scrap]= "section_scrap"
cat_name[dead]= "@d"
cat_name[dot_dot_dot]= "'...'"
cat_name[constant]= "constant"
cat_name[str]= "str"
cat_name[identifier]= "identifier"
cat_name[0]= "zero"
cat_name[direct]= "'<-'"
cat_name[plus_plus]= "'++'"
cat_name[minus_minus]= "'--'"
cat_name[verbatim]= "verbatim"



/*:168*/



/*367:*/


//line goweave.w:6361

max_sort_ptr= 0



/*:367*/


//line goweave.w:91

if show_banner(){
fmt.Print(banner)/* print a ``banner line'' */
}


/*108:*/


//line goweave.w:346


id_lookup([]rune("break"),break_token)
id_lookup([]rune("case"),case_token)
id_lookup([]rune("chan"),chan_token)
id_lookup([]rune("const"),const_token)
id_lookup([]rune("continue"),continue_token)
id_lookup([]rune("default"),default_token)
id_lookup([]rune("defer"),defer_token)
id_lookup([]rune("else"),else_token)
id_lookup([]rune("fallthrough"),fallthrough_token)
id_lookup([]rune("for"),for_token)
id_lookup([]rune("func"),func_token)
id_lookup([]rune("go"),go_token)
id_lookup([]rune("goto"),goto_token)
id_lookup([]rune("if"),if_token)
id_lookup([]rune("import"),import_token)
id_lookup([]rune("interface"),interface_token)
id_lookup([]rune("map"),Type)
id_lookup([]rune("package"),package_token)
id_lookup([]rune("range"),range_token)
id_lookup([]rune("return"),return_token)
id_lookup([]rune("select"),select_token)
id_lookup([]rune("struct"),struct_token)
id_lookup([]rune("switch"),switch_token)
id_lookup([]rune("type"),type_token)
id_lookup([]rune("var"),var_token)

id_lookup([]rune("bool"),Type)
id_lookup([]rune("byte"),Type)
id_lookup([]rune("complex64"),Type)
id_lookup([]rune("complex128"),Type)
id_lookup([]rune("error"),Type)
id_lookup([]rune("float32"),Type)
id_lookup([]rune("float64"),Type)
id_lookup([]rune("int"),Type)
id_lookup([]rune("int8"),Type)
id_lookup([]rune("int16"),Type)
id_lookup([]rune("int32"),Type)
id_lookup([]rune("int64"),Type)
id_lookup([]rune("rune"),Type)
id_lookup([]rune("string"),Type)
id_lookup([]rune("uint"),Type)
id_lookup([]rune("uint8"),Type)
id_lookup([]rune("uint16"),Type)
id_lookup([]rune("uint32"),Type)
id_lookup([]rune("uint64"),Type)
id_lookup([]rune("uintptr"),Type)

id_lookup([]rune("true"),constant)
id_lookup([]rune("false"),constant)
id_lookup([]rune("iota"),constant)

id_lookup([]rune("nil"),constant)

id_lookup([]rune("append"),identifier)
id_lookup([]rune("cap"),identifier)
id_lookup([]rune("close"),identifier)
id_lookup([]rune("complex"),identifier)
id_lookup([]rune("copy"),identifier)
id_lookup([]rune("delete"),identifier)
id_lookup([]rune("imag"),identifier)
id_lookup([]rune("len"),identifier)
id_lookup([]rune("make"),identifier)
id_lookup([]rune("new"),identifier)
id_lookup([]rune("panic"),identifier)
id_lookup([]rune("print"),identifier)
id_lookup([]rune("println"),identifier)
id_lookup([]rune("real"),identifier)
id_lookup([]rune("recover"),identifier)
res_wd_end= int32(len(name_dir))
id_lookup([]rune("TeX"),custom)



/*:108*/


//line goweave.w:95

phase_one()/* read all the user's text and store the cross-references */
phase_two()/* read all the text again and translate it to \TEX/ form */
phase_three()/* output the cross-reference index */
os.Exit(wrap_up())/* and exit gracefully */
}



/*:3*/



/*6:*/


//line gocommon.w:25

const(


/*10:*/


//line gocommon.w:63

and_and rune= 04/* `\.{\&\&}'\,; corresponds to MIT's {\tentex\char'4} */
lt_lt rune= 020/* `\.{<<}'\,;  corresponds to MIT's {\tentex\char'20} */
gt_gt rune= 021/* `\.{>>}'\,;  corresponds to MIT's {\tentex\char'21} */
plus_plus rune= 0200/* `\.{++}'\,;  corresponds to MIT's {\tentex\char'13} */
minus_minus rune= 0201/* `\.{--}'\,;  corresponds to MIT's {\tentex\char'1} */
col_eq rune= 0207/* `\.{:=}'\, */
not_eq rune= 032/* `\.{!=}'\,;  corresponds to MIT's {\tentex\char'32} */
lt_eq rune= 034/* `\.{<=}'\,;  corresponds to MIT's {\tentex\char'34} */
gt_eq rune= 035/* `\.{>=}'\,;  corresponds to MIT's {\tentex\char'35} */
eq_eq rune= 036/* `\.{==}'\,;  corresponds to MIT's {\tentex\char'36} */
or_or rune= 037/* `\.{\v\v}'\,;  corresponds to MIT's {\tentex\char'37} */
dot_dot_dot rune= 0202/* `\.{...}' */
begin_comment rune= '\t'/* tab marks will not appear */
and_not rune= 010/*`\.{\&\^}'\,;*/
direct rune= 0203/*`\.{<-}'\,;*/
begin_short_comment rune= 031/* short comment */



/*:10*/



/*31:*/


//line gocommon.w:402

max_sections= 2000/* number of identifiers, strings, section names;
  must be less than 10240 */




/*:31*/



/*42:*/


//line gocommon.w:614

hash_size= 353/* should be prime */



/*:42*/



/*55:*/


//line gocommon.w:757

less= 0/* the first name is lexicographically less than the second */
equal= 1/* the first name is equal to the second */
greater= 2/* the first name is lexicographically greater than the second */
prefix= 3/* the first name is a proper prefix of the second */
extension= 4/* the first name is a proper extension of the second */



/*:55*/



/*63:*/


//line gocommon.w:978

bad_extension= 5



/*:63*/



/*65:*/


//line gocommon.w:1041

spotless= 0/* history value for normal jobs */
harmless_message= 1/* history value when non-serious info was printed */
error_message= 2/* history value when an error was noted */
fatal_message= 3/* history value when we had to stop prematurely */



/*:65*/


//line gocommon.w:27

)



/*12:*/


//line gocommon.w:87

var buffer[]rune/* where each line of input goes */
var loc int= 0/* points to the next character to be read from the buffer */
var section_text[]rune/* name being sought for */
var id[]rune/* slice pointed to the current identifier */



/*:12*/



/*17:*/


//line gocommon.w:138

var include_depth int/* current level of nesting */
var file[]*bufio.Reader/* stack of non-change files */
var change_file*bufio.Reader/* change file */
var file_name[]string
/* stack of non-change file names */
var change_file_name string= "/dev/null"/* name of change file */
var alt_file_name string/* alternate name to try */
var line[]int/* number of current line in the stacked files */
var change_line int/* number of current line in change file */
var change_depth int/* where \.{@y} originated during a change */
var input_has_ended bool/* if there is no more input */
var changing bool/* if the current line is from change_file */



/*:17*/



/*32:*/


//line gocommon.w:407

var section_count int32/* the current section number */
var changed_section[max_sections]bool/* is the section changed? */
var change_pending bool/* if the current change is not yet recorded in
  changed_section[section_count] */
var print_where bool= false/* should \.{GOTANGLE} print line and file info? */



/*:32*/



/*40:*/


//line gocommon.w:589

type name_info struct{
name[]rune


/*41:*/


//line gocommon.w:603

llink int32



/*:41*/



/*50:*/


//line gocommon.w:684

ispref bool/* prefix flag*/
rlink int32/* right link in binary search tree for section names */




/*:50*/



/*91:*/


//line goweave.w:146

ilk int32/* used by identifiers in \.{GOWEAVE} only */



/*:91*/



/*97:*/


//line goweave.w:199

xref int32/* info corresponding to names */



/*:97*/


//line gocommon.w:592

}/* contains information about an identifier or section name */
type name_index int/* index into array of name_infos */
var name_dir[]name_info/* information about names */
var name_root int32



/*:40*/



/*43:*/


//line gocommon.w:617

var hash[hash_size]int32/* heads of hash lists */
var h int32/* index into hash-head array */



/*:43*/



/*68:*/


//line gocommon.w:1061

var history int= spotless/* indicates how bad this run was */



/*:68*/



/*80:*/


//line gocommon.w:1218

var go_file_name string/* name of go_file */
var tex_file_name string/* name of tex_file */
var idx_file_name string/* name of idx_file */
var scn_file_name string/* name of scn_file */
var flags[128]bool/* an option for each 7-bit code */



/*:80*/



/*87:*/


//line gocommon.w:1360

var go_file io.WriteCloser/* where output of \.{GOTANGLE} goes */
var tex_file io.WriteCloser/* where output of \.{GOWEAVE} goes */
var idx_file io.WriteCloser/* where index from \.{GOWEAVE} goes */
var scn_file io.WriteCloser/* where list of sections from \.{GOWEAVE} goes */
var active_file io.WriteCloser/* currently active file for \.{GOWEAVE} output */



/*:87*/


//line gocommon.w:30



/*7:*/


//line gocommon.w:39
var phase int/* which phase are we in? */



/*:7*/



/*18:*/


//line gocommon.w:157

var change_buffer[]rune/* next line of change_file */



/*:18*/


//line gocommon.w:31




/*:6*/



/*8:*/


//line gocommon.w:45

func common_init(){


/*44:*/


//line gocommon.w:621

for i,_:=range hash{
hash[i]= -1
}



/*:44*/



/*51:*/


//line gocommon.w:689

name_root= -1/* the binary search tree starts out with nothing in it */



/*:51*/


//line gocommon.w:47



/*81:*/


//line gocommon.w:1229

flags['b']= true
flags['h']= true
flags['p']= true



/*:81*/


//line gocommon.w:48



/*88:*/


//line gocommon.w:1367

scan_args()


/*384:*/


//line goweave.w:6622

if f,err:=os.OpenFile(tex_file_name,os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0666);
err!=nil{
fatal("! Cannot open output file ",tex_file_name)
}else{
tex_file= f
}




/*:384*/


//line gocommon.w:1369




/*:88*/


//line gocommon.w:49

}




/*:8*/



/*14:*/


//line gocommon.w:98

/* copies a line into buffer or returns error */
func input_ln(fp*bufio.Reader)error{
var prefix bool
var err error
var buf[]byte
var b[]byte
buffer= nil
for buf,prefix,err= fp.ReadLine()
err==nil&&prefix
b,prefix,err= fp.ReadLine(){
buf= append(buf,b...)
}
if len(buf)> 0{
buffer= bytes.Runes(buf)
}
if err==io.EOF&&len(buffer)!=0{
return nil
}
if err==nil&&len(buffer)==0{
buffer= append(buffer,' ')
}
return err
}



/*:14*/



/*19:*/


//line gocommon.w:167

func prime_the_change_buffer(){
change_buffer= nil


/*21:*/


//line gocommon.w:182

for true{
change_line++
if err:=input_ln(change_file);err!=nil{
return
}
if len(buffer)<2{
continue
}
if buffer[0]!='@'{
continue
}
if unicode.IsUpper(buffer[1]){
buffer[1]= unicode.ToLower(buffer[1])
}
if buffer[1]=='x'{
break
}
if buffer[1]=='y'||buffer[1]=='z'||buffer[1]=='i'{
loc= 2
err_print("! Missing @x in change file")

}
}



/*:21*/


//line gocommon.w:170



/*22:*/


//line gocommon.w:209

for true{
change_line++
if err:=input_ln(change_file);err!=nil{
err_print("! Change file ended after @x")

return
}
if len(buffer)!=0{
break
}
}



/*:22*/


//line gocommon.w:171



/*23:*/


//line gocommon.w:222

{
change_buffer= buffer
buffer= nil
}



/*:23*/


//line gocommon.w:172

}



/*:19*/



/*24:*/


//line gocommon.w:243

func if_section_start_make_pending(b bool){
for loc= 0;loc<len(buffer)&&unicode.IsSpace(buffer[loc]);loc++{}
if len(buffer)>=2&&buffer[0]=='@'&&(unicode.IsSpace(buffer[1])||buffer[1]=='*'){
change_pending= b
}
}



/*:24*/



/*25:*/


//line gocommon.w:253

func compare_runes(l[]rune,r[]rune)int{
i:=0
for;i<len(l)&&i<len(r)&&l[i]==r[i];i++{}
if i==len(r){
if i==len(l){
return 0
}else{
return-1
}
}else{
if i==len(l){
return 1
}else if l[i]<r[i]{
return-1
}else{
return 1
}
}
return 0
}



/*:25*/



/*26:*/


//line gocommon.w:276

/* switches to change_file if the buffers match */
func check_change(){
n:=0/* the number of discrepancies found */
if compare_runes(buffer,change_buffer)!=0{
return
}
change_pending= false
if!changed_section[section_count]{
if_section_start_make_pending(true)
if!change_pending{
changed_section[section_count]= true
}
}
for true{
changing= true
print_where= true
change_line++
if err:=input_ln(change_file);err!=nil{
err_print("! Change file ended before @y")

change_buffer= nil
changing= false
return
}
if len(buffer)> 1&&buffer[0]=='@'{
var xyz_code rune
if unicode.IsUpper(buffer[1]){
xyz_code= unicode.ToLower(buffer[1])
}else{
xyz_code= buffer[1]
}


/*28:*/


//line gocommon.w:333

if xyz_code=='x'||xyz_code=='z'{
loc= 2
err_print("! Where is the matching @y?")

}else if xyz_code=='y'{
if n> 0{
loc= 2
fmt.Printf("\n! Hmm... %d ",n)
err_print("of the preceding lines failed to match")

}
change_depth= include_depth
return
}



/*:28*/


//line gocommon.w:309

}


/*23:*/


//line gocommon.w:222

{
change_buffer= buffer
buffer= nil
}



/*:23*/


//line gocommon.w:311

changing= false
line[include_depth]++
for input_ln(file[include_depth])!=nil{/* pop the stack or quit */
if include_depth==0{
err_print("! GOWEB file ended during a change")

input_has_ended= true
return
}
include_depth--
line[include_depth]++
}
if compare_runes(buffer,change_buffer)!=0{
n++
}
}
}



/*:26*/



/*29:*/


//line gocommon.w:353

func reset_input(){
loc= 0
file= file[:0]


/*30:*/


//line gocommon.w:372

if wf,err:=os.Open(file_name[0]);err!=nil{
file_name[0]= alt_file_name
if wf,err= os.Open(file_name[0]);err!=nil{
fatal("! Cannot open input file ",file_name[0])

}else{
file= append(file,bufio.NewReader(wf))
}
}else{
file= append(file,bufio.NewReader(wf))
}
if cf,err:=os.Open(change_file_name);err!=nil{
fatal("! Cannot open change file ",change_file_name)

}else{
change_file= bufio.NewReader(cf)
}



/*:30*/


//line gocommon.w:357

include_depth= 0
line= line[:0]
line= append(line,0)
change_line= 0
change_depth= include_depth
changing= true
prime_the_change_buffer()
changing= !changing
loc= 0
input_has_ended= false
}



/*:29*/



/*33:*/


//line gocommon.w:415

func get_line()bool{/* inputs the next line */
restart:
if changing&&include_depth==change_depth{


/*37:*/


//line gocommon.w:535
{
change_line++
if input_ln(change_file)!=nil{
err_print("! Change file ended without @z")

buffer= append(buffer,[]rune("@z")...)
}
if len(buffer)> 0{/* check if the change has ended */
if change_pending{
if_section_start_make_pending(false)
if change_pending{
changed_section[section_count]= true
change_pending= false
}
}
if len(buffer)>=2&&buffer[0]=='@'{
if unicode.IsUpper(buffer[1]){
buffer[1]= unicode.ToLower(buffer[1])
}
if buffer[1]=='x'||buffer[1]=='y'{
loc= 2
err_print("! Where is the matching @z?")

}else if buffer[1]=='z'{
prime_the_change_buffer()
changing= !changing
print_where= true
}
}
}
}



/*:37*/


//line gocommon.w:419

}
if!changing||include_depth> change_depth{


/*36:*/


//line gocommon.w:505
{
line[include_depth]++
for input_ln(file[include_depth])!=nil{/* pop the stack or quit */
print_where= true
if include_depth==0{
input_has_ended= true
break
}else{
file[include_depth]= nil
file_name= file_name[:include_depth]
file= file[:include_depth]
line= line[:include_depth]
include_depth--
if changing&&include_depth==change_depth{
break
}
line[include_depth]++
}
}
if!changing&&!input_has_ended{
if len(buffer)==len(change_buffer){
if buffer[0]==change_buffer[0]{
if len(change_buffer)> 0{
check_change()
}
}
}
}
}



/*:36*/


//line gocommon.w:422

if changing&&include_depth==change_depth{
goto restart
}
}
if input_has_ended{
return false
}
loc= 0
if len(buffer)>=2&&buffer[0]=='@'&&(buffer[1]=='i'||buffer[1]=='I'){
loc= 2
for loc<len(buffer)&&unicode.IsSpace(buffer[loc]){
loc++
}
if loc>=len(buffer){
err_print("! Include file name not given")

goto restart
}

include_depth++/* push input stack */


/*35:*/


//line gocommon.w:461
{
l:=loc
if buffer[loc]=='"'{
loc++
l++
for loc<len(buffer)&&buffer[loc]!='"'{
loc++
}
}else{
for loc<len(buffer)&&!unicode.IsSpace(buffer[loc]){
loc++
}
}

file_name= append(file_name,string(buffer[l:loc]))


if f,err:=os.Open(file_name[include_depth]);err==nil{
file= append(file,bufio.NewReader(f))
line= append(line,0)
print_where= true
goto restart/* success */
}
temp_file_name:=os.Getenv("GOWEBINPUTS")
if len(temp_file_name)!=0{

for _,fn:=range strings.Split(temp_file_name,":"){
file_name[include_depth]= fn+"/"+file_name[include_depth]
if f,err:=os.Open(file_name[include_depth]);err==nil{
file= append(file,bufio.NewReader(f))
line= append(line,0)
print_where= true
goto restart/* success */
}
}
}
file_name= file_name[:include_depth]
file= file[:include_depth]
line= line[:include_depth]
include_depth--
err_print("! Cannot open include file")
goto restart
}



/*:35*/


//line gocommon.w:443

}
return true
}



/*:33*/



/*38:*/


//line gocommon.w:570

func check_complete(){
if len(change_buffer)> 0{/* changing is false */
buffer= change_buffer
change_buffer= nil
changing= true
change_depth= include_depth
loc= 0
err_print("! Change file entry did not match")

}
}



/*:38*/



/*45:*/


//line gocommon.w:628

/* looks up a string in the identifier table */
func id_lookup(
id[]rune,/* string with id */
t int32/* the ilk; used by \.{GOWEAVE} only */)int32{


/*46:*/


//line gocommon.w:645

h:=id[0]
for i:=1;i<len(id);i++{
h= (h+h+id[i])%hash_size
}



/*:46*/


//line gocommon.w:633



/*47:*/


//line gocommon.w:654

p:=hash[h]
for p!=-1&&!names_match(p,id,t){
p= name_dir[p].llink
}
if p==-1{
p:=int32(len(name_dir))/* the current identifier is new */
name_dir= append(name_dir,name_info{})
name_dir[p].llink= -1
init_node(p)
name_dir[p].llink= hash[h]
hash[h]= p/* insert p at beginning of hash list */
}



/*:47*/


//line gocommon.w:634

if p==-1{


/*49:*/


//line gocommon.w:672

p= int32(len(name_dir)-1)
name_dir[p].name= append(name_dir[p].name,id...)


/*107:*/


//line goweave.w:335

name_dir[p].ilk= t
name_dir[p].xref= 0




/*:107*/


//line gocommon.w:675




/*:49*/


//line gocommon.w:636

}
return p
}



/*:45*/



/*52:*/


//line gocommon.w:709

func print_section_name(p int32){
q:=p+1
for p!=-1{
fmt.Print(string(name_dir[p].name[1:]))
if name_dir[p].ispref{
p= name_dir[q].llink
q= p
}else{
p= -1
q= -2
}
}
if q!=-2{
fmt.Print("...")/* complete name not yet known */
}
}



/*:52*/



/*53:*/


//line gocommon.w:728

func sprint_section_name(p int32)[]rune{
q:=p+1
var dest[]rune
for p!=-1{
dest= append(dest,name_dir[p].name[1:]...)
if name_dir[p].ispref{
p= name_dir[q].llink
q= p
}else{
p= -1
}
}
return dest
}



/*:53*/



/*54:*/


//line gocommon.w:745

func print_prefix_name(p int32){
l:=name_dir[p].name[0]
fmt.Print(string(name_dir[p].name[1:]))
if int(l)<len(name_dir[p].name){
fmt.Print("...")
}
}



/*:54*/



/*56:*/


//line gocommon.w:765

/* fuller comparison than strcmp */
func web_strcmp(
j[]rune,/* first string */
k[]rune/* second string */)int{
i:=0
for;i<len(j)&&i<len(k)&&j[i]==k[i];i++{}
if i==len(k){
if i==len(j){
return equal
}else{
return extension
}
}else{
if i==len(j){
return prefix
}else if j[i]<k[i]{
return less
}else{
return greater
}
}
return equal
}



/*:56*/



/*57:*/


//line gocommon.w:803

/* install a new node in the tree */
func add_section_name(
par int32,/* parent of new node */
c int,/* right or left? */
name[]rune,/* section name */
ispref bool/* are we adding a prefix or a full name? */)int32{
p:=int32(len(name_dir))/* new node */
name_dir= append(name_dir,name_info{})
name_dir[p].llink= -1
init_node(p)
if ispref{
name_dir= append(name_dir,name_info{})
name_dir[p+1].llink= -1
init_node(p+1)
}
name_dir[p].ispref= ispref
name_dir[p].name= append(name_dir[p].name,int32(len(name)))/* length of section name */
name_dir[p].name= append(name_dir[p].name,name...)
name_dir[p].llink= -1
name_dir[p].rlink= -1
init_node(p)
if par==-1{
name_root= p
}else{
if c==less{
name_dir[par].llink= p
}else{
name_dir[par].rlink= p
}
}
return p
}



/*:57*/



/*58:*/


//line gocommon.w:838

func extend_section_name(
p int32,/* index name to be extended */
text[]rune,/* extension text */
ispref bool/* are we adding a prefix or a full name? */){
q:=p+1
for name_dir[q].llink!=-1{
q= name_dir[q].llink
}
np:=int32(len(name_dir))
name_dir[q].llink= np
name_dir= append(name_dir,name_info{})
name_dir[np].llink= -1
init_node(np)
name_dir[np].name= append(name_dir[np].name,int32(len(text)))/* length of section name */
name_dir[np].name= append(name_dir[np].name,text...)
name_dir[np].ispref= ispref

}



/*:58*/



/*59:*/


//line gocommon.w:863

/* find or install section name in tree */
func section_lookup(
name[]rune,/* new name */
ispref bool/* is the new name a prefix or a full name? */)int32{
c:=less/* comparison between two names*/
p:=name_root/* current node of the search tree */
var q int32= -1/* another place to look in the tree */
var r int32= -1/* where a match has been found */
var par int32= -1/* parent of p, if r is NULL; otherwise parent of r */
name_len:=len(name)


/*60:*/


//line gocommon.w:886

for p!=-1{/* compare shortest prefix of p with new name */
c= web_strcmp(name,name_dir[p].name[1:])
if c==less||c==greater{/* new name does not match p */
if r==-1{/* no previous matches have been found */
par= p
}
if c==less{
p= name_dir[p].llink
}else{
p= name_dir[p].rlink
}
}else{/* new name matches p */
if r!=-1{/* and also r: illegal */
fmt.Printf("\n! Ambiguous prefix: matches <")

print_prefix_name(p)
fmt.Printf(">\n and <")
print_prefix_name(r)
err_print(">")
return 0/* the unsection */
}
r= p/* remember match */
p= name_dir[p].llink/* try another */
q= name_dir[r].rlink/* we'll get back here if the new p doesn't match */
}
if p==-1{
p= q
q= -1/* q held the other branch of r */
}
}



/*:60*/


//line gocommon.w:875



/*61:*/


//line gocommon.w:918

if r==-1{/* no matches were found */
return add_section_name(par,c,name,ispref)
}



/*:61*/


//line gocommon.w:876



/*62:*/


//line gocommon.w:927

first,cmp:=section_name_cmp(name,r)
switch cmp{
/* compare all of r with new name */
case prefix:
if!ispref{
fmt.Printf("\n! New name is a prefix of <")

print_section_name(r)
err_print(">")
}else if name_len<int(name_dir[r].name[0]){
name_dir[r].name[0]= int32(len(name)-first)
}
fallthrough
case equal:
return r
case extension:
if!ispref||first<len(name){
extend_section_name(r,name[first:],ispref)
}
return r
case bad_extension:
fmt.Printf("\n! New name extends <")

print_section_name(r)
err_print(">")
return r
default:/* no match: illegal */
fmt.Printf("\n! Section name incompatible with <")

print_prefix_name(r)
fmt.Printf(">,\n which abbreviates <")
print_section_name(r)
err_print(">")
return r
}



/*:62*/


//line gocommon.w:877

return-1
}



/*:59*/



/*64:*/


//line gocommon.w:982

func section_name_cmp(
name[]rune,/* comparison string */
r int32/* section name being compared */)(int,int){
q:=r+1/* access to subsequent chunks */
var ispref bool/* is chunk r a prefix? */
first:=0
for true{
if name_dir[r].ispref{
ispref= true
q= name_dir[q].llink
}else{
ispref= false
q= -1
}
c:=web_strcmp(name,name_dir[r].name[1:])
switch c{
case equal:
if q==-1{
if ispref{
return first+len(name_dir[r].name[1:]),extension/* null extension */
}else{
return first,equal
}
}else{
if compare_runes(name_dir[q].name,name_dir[q+1].name)==0{
return first,equal
}else{
return first,prefix
}
}
case extension:
if!ispref{
return first,bad_extension
}
first+= len(name_dir[r].name[1:])
if q!=-1{
name= name[len(name_dir[r].name[1:]):]
r= q
continue
}
return first,extension
default:
return first,c
}
}
return-2,-1
}



/*:64*/



/*66:*/


//line gocommon.w:1048

func mark_harmless(){
if history==spotless{
history= harmless_message
}
}



/*:66*/



/*67:*/


//line gocommon.w:1056

func mark_error(){
history= error_message
}



/*:67*/



/*69:*/


//line gocommon.w:1071

/* prints `\..' and location of error message */
func err_print(s string){
var l int/* pointers into buffer */
if len(s)> 0&&s[0]=='!'{
fmt.Printf("\n%s",s)
}else{
fmt.Printf("%s",s)
}
if len(file)> 0&&file[0]!=nil{


/*70:*/


//line gocommon.w:1096

{
if changing&&include_depth==change_depth{
fmt.Printf(". (change file %s:%d)\n",file_name[include_depth],change_line)
}else if include_depth==0&&len(line)> 0{
fmt.Printf(". (%s:%d)\n",file_name[include_depth],line[include_depth])
}else if len(line)> include_depth{
fmt.Printf(". (include file %s:%d)\n",file_name[include_depth],line[include_depth])
}
l= len(buffer)
if loc<l{
l= loc
}
if l> 0{
for k:=0;k<l;k++{
if buffer[k]=='\t'{
fmt.Print(" ")
}else{
fmt.Printf("%c",buffer[k])// print the characters already read 
}
}

fmt.Println()
fmt.Printf("%*c",l,' ')
}
fmt.Println(string(buffer[l:]))
if len(buffer)> 0&&buffer[len(buffer)-1]=='|'{
fmt.Print("|")/* end of \GO/ text in section names */
}
fmt.Print(" ")/* to separate the message from future asterisks */
}



/*:70*/


//line gocommon.w:1081

}
os.Stdout.Sync()
mark_error()
}



/*:69*/



/*72:*/


//line gocommon.w:1142

func wrap_up()int{
fmt.Print("\n")
if show_stats(){
print_stats()/* print statistics about memory usage */
}


/*73:*/


//line gocommon.w:1155

switch history{
case spotless:
if show_happiness(){
fmt.Printf("(No errors were found.)\n")
}
case harmless_message:
fmt.Printf("(Did you see the warning message above?)\n")
case error_message:
fmt.Printf("(Pardon me, but I think I spotted something wrong.)\n")
case fatal_message:
fmt.Printf("(That was a fatal error, my friend.)\n")
}/* there are no other cases */



/*:73*/


//line gocommon.w:1148

if history> harmless_message{
return 1
}
return 0
}



/*:72*/



/*74:*/


//line gocommon.w:1175

func fatal(s string,t string){
if len(s)!=0{
fmt.Print(s)
}
err_print(t)
history= fatal_message
os.Exit(wrap_up())
}



/*:74*/



/*76:*/


//line gocommon.w:1195

func show_banner()bool{
return flags['b']/* should the banner line be printed? */
}



/*:76*/



/*77:*/


//line gocommon.w:1201

func show_progress()bool{
return flags['p']/* should progress reports be printed? */
}



/*:77*/



/*78:*/


//line gocommon.w:1207

func show_stats()bool{
return flags['s']/* should statistics be printed at end of run? */
}



/*:78*/



/*79:*/


//line gocommon.w:1213

func show_happiness()bool{
return flags['h']/* should lack of errors be announced? */
}



/*:79*/



/*82:*/


//line gocommon.w:1249

func scan_args(){
dot_pos:=-1/* position of '.' in the argument */
name_pos:=0/* file name beginning, sans directory */
found_web:=false
found_change:=false
found_out:=false
/* have these names been seen? */
flag_change:=false

for i:=1;i<len(os.Args);i++{
arg:=os.Args[i]
if(arg[0]=='-'||arg[0]=='+')&&len(arg)> 1{


/*86:*/


//line gocommon.w:1346

{
if arg[0]=='-'{
flag_change= false
}else{
flag_change= true
}
for i:=1;i<len(arg);i++{
flags[arg[i]]= flag_change
}
}



/*:86*/


//line gocommon.w:1262

}else{
name_pos= 0
dot_pos= -1
for j:=0;j<len(arg);j++{
if arg[j]=='.'{
dot_pos= j
}else if arg[j]=='/'{
dot_pos= -1
name_pos= j+1
}
}
if!found_web{


/*83:*/


//line gocommon.w:1297

{
if dot_pos==-1{
file_name= append(file_name,fmt.Sprintf("%s.w",arg))
}else{
file_name= append(file_name,arg)
arg= arg[:dot_pos]/* string now ends where the dot was */
}
alt_file_name= fmt.Sprintf("%s.web",arg)
tex_file_name= fmt.Sprintf("%s.tex",arg[name_pos:])/* strip off directory name */
idx_file_name= fmt.Sprintf("%s.idx",arg[name_pos:])
scn_file_name= fmt.Sprintf("%s.scn",arg[name_pos:])
go_file_name= fmt.Sprintf("%s.go",arg[name_pos:])
found_web= true
}



/*:83*/


//line gocommon.w:1275

}else if!found_change{


/*84:*/


//line gocommon.w:1313

{
if arg[0]=='-'{
found_change= true
}else{
if dot_pos==-2{
change_file_name= fmt.Sprintf("%s.ch",arg)
}else{
change_file_name= arg
}
found_change= true
}
}



/*:84*/


//line gocommon.w:1277

}else if!found_out{


/*85:*/


//line gocommon.w:1327

{
if dot_pos==-1{
tex_file_name= fmt.Sprintf("%s.tex",arg)
idx_file_name= fmt.Sprintf("%s.idx",arg)
scn_file_name= fmt.Sprintf("%s.scn",arg)
go_file_name= fmt.Sprintf("%s.go",arg)
}else{
tex_file_name= arg
go_file_name= arg
if flags['x']{/* indexes will be generated */
dot_pos= -1
idx_file_name= fmt.Sprintf("%s.idx",arg)
scn_file_name= fmt.Sprintf("%s.scn",arg)
}
}
found_out= true
}



/*:85*/


//line gocommon.w:1279

}else{


/*383:*/


//line goweave.w:6614

{
fatal("! Usage: goweave [options] webfile[.w] [{changefile[.ch]|-} [outfile[.tex]]]\n","")

}



/*:383*/


//line gocommon.w:1281

}
}
}
if!found_web{


/*383:*/


//line goweave.w:6614

{
fatal("! Usage: goweave [options] webfile[.w] [{changefile[.ch]|-} [outfile[.tex]]]\n","")

}



/*:383*/


//line gocommon.w:1286

}
}



/*:82*/



/*89:*/


//line gocommon.w:1374

func xisxdigit(r rune)bool{
if unicode.IsDigit(r){
return true
}
if!unicode.IsLetter(r){
return false
}
r= unicode.ToLower(r)
if r>='a'&&r<='f'{
return true
}
return false
}



/*:89*/



/*99:*/


//line goweave.w:216

func append_xref(c int32){
xmem= append(xmem,xref_info{})
xmem[len(xmem)-1].num= c
xmem[len(xmem)-1].xlink= 0
}




/*:99*/



/*100:*/


//line goweave.w:225

func is_tiny(p int32)bool{
return p<int32(len(name_dir))&&len(name_dir[p].name)==1
}



/*:100*/



/*101:*/


//line goweave.w:231

/* tells if uses of a name are to be indexed */
func unindexed(p int32)bool{
return p<res_wd_end&&name_dir[p].ilk>=custom
}



/*:101*/



/*102:*/


//line goweave.w:238

func new_xref(p int32){
if flags['x']==false{
return
}
if(unindexed(p)||is_tiny(p))&&xref_switch==0{
return
}
m:=section_count+xref_switch
xref_switch= 0
q:=name_dir[p].xref/* pointer to previous cross-reference */
if q>=0{
n:=xmem[q].num/* new and previous cross-reference value */
if n==m||n==m+def_flag{
return
}else if m==n+def_flag{
xmem[q].num= m
return
}
}
append_xref(m)
xmem[len(xmem)-1].xlink= int32(q)
name_dir[p].xref= int32(len(xmem)-1)
}



/*:102*/



/*103:*/


//line goweave.w:274

func new_section_xref(p int32){
var r int32= 0/* pointers to previous cross-references */
q:=name_dir[p].xref

if q>=0{
for q>=0&&q<int32(len(xmem))&&xmem[q].num> section_xref_switch{
r= q
q= xmem[q].xlink
}
}
if r> 0&&r<int32(len(xmem))&&xmem[r].num==section_count+section_xref_switch{
return/* don't duplicate entries */
}
append_xref(section_count+section_xref_switch)
xmem[len(xmem)-1].xlink= q
section_xref_switch= 0
if r==0{
name_dir[p].xref= int32(len(xmem)-1)
}else{
xmem[r].xlink= int32(len(xmem)-1)
}
}



/*:103*/



/*104:*/


//line goweave.w:301

func set_file_flag(p int32){
q:=name_dir[p].xref
if xmem[q].num==file_flag{
return
}
append_xref(file_flag)
xmem[len(xmem)-1].xlink= q
name_dir[p].xref= int32(len(xmem)-1)
}



/*:104*/



/*105:*/


//line goweave.w:313

func names_match(
p int32,/* points to the proposed match */
id[]rune,
t int32/* desired ilk */)bool{
if len(name_dir[p].name)!=len(id){
return false
}
if name_dir[p].ilk!=t&&!(t==normal&&name_dir[p].ilk> zero){
return false
}
return compare_runes(id,name_dir[p].name)==0
}



/*:105*/



/*106:*/


//line goweave.w:328

func init_node(p int32){
name_dir[p].xref= 0
}



/*:106*/



/*116:*/


//line goweave.w:536

func skip_limbo(){
for{
if loc>=len(buffer)&&!get_line(){
return
}
for loc<len(buffer)&&buffer[loc]!='@'{
loc++/* look for '@', then skip two symbols */
}
l:=loc
loc++
if l<len(buffer){
c:=new_section
if loc<len(buffer){
c= ccode[buffer[loc]]
loc++
}
if c==new_section{
return
}
if c==noop{
skip_restricted()
}else if c==format_code{


/*143:*/


//line goweave.w:1233

{
if get_next()!=identifier{
err_print("! Missing left identifier of @s")

}else{
lhs= id_lookup(id,normal)
if get_next()!=identifier{
err_print("! Missing right identifier of @s")

}else{
rhs= id_lookup(id,normal)
name_dir[lhs].ilk= name_dir[rhs].ilk
}
}
}



/*:143*/


//line goweave.w:559

}
}
}
}



/*:116*/



/*117:*/


//line goweave.w:572

/* skip past pure \TEX/ code */
func skip_TeX()rune{
for{
if loc>=len(buffer)&&!get_line(){
return new_section
}
for loc<len(buffer)&&buffer[loc]!='@'&&buffer[loc]!='|'{
loc++
}
l:=loc
loc++
if l<len(buffer)&&buffer[l]=='|'{
return'|'
}
if loc<len(buffer){
l:=loc
loc++
return ccode[buffer[l]]
}
if l<len(buffer)&&buffer[l]=='@'{
return new_section
}
}
return 0
}



/*:117*/



/*121:*/


//line goweave.w:653

/* produces the next input token */
func get_next()rune{
for{
if loc>=len(buffer){
// Looking for last non-insert scrap
i:=len(scrap_info)-1
for;i>=0&&scrap_info[i].cat==insert;i--{}
if i>=0&&
(scrap_info[i].cat==identifier||
scrap_info[i].cat==constant||
scrap_info[i].cat==str||
scrap_info[i].cat==break_token||
scrap_info[i].cat==continue_token||
scrap_info[i].cat==fallthrough_token||
scrap_info[i].cat==return_token||
scrap_info[i].cat==plus_plus||
scrap_info[i].cat==minus_minus||
scrap_info[i].cat==rpar||
scrap_info[i].cat==rbracket||
scrap_info[i].cat==rbrace||
scrap_info[i].cat==Type){
return pseudo_semi
}
if!get_line(){
return new_section
}
}
c:=buffer[loc]/* the current character */
loc++
nc:=' '
if loc<len(buffer){
nc= buffer[loc]
}
if unicode.IsDigit(c)||(c=='.'&&unicode.IsDigit(nc)){


/*124:*/


//line goweave.w:735
{
id= nil
is_dec:=false
if loc<len(buffer)&&buffer[loc-1]=='0'{
if buffer[loc]=='x'||buffer[loc]=='X'{/* hex constant */
id= append(id,'^')
loc++
for loc<len(buffer)&&xisxdigit(buffer[loc]){
id= append(id,buffer[loc])
loc++
}
}else if unicode.IsDigit(buffer[loc]){/* octal constant */
id= append(id,'~')
for loc<len(buffer)&&unicode.IsDigit(buffer[loc]){
id= append(id,buffer[loc])
loc++
}
}else{
is_dec= true/* decimal constant */
}
}else{
is_dec= true
}
if is_dec{/* decimal constant */
if loc<len(buffer)&&buffer[loc-1]=='.'&&!unicode.IsDigit(buffer[loc]){
goto mistake/* not a constant */
}
id= append(id,buffer[loc-1])
for loc<len(buffer)&&(unicode.IsDigit(buffer[loc])||buffer[loc]=='.'){
id= append(id,buffer[loc])
loc++
}
if loc<len(buffer)&&(buffer[loc]=='e'||buffer[loc]=='E'){/* float constant */
id= append(id,'_')
loc++
if loc<len(buffer)&&(buffer[loc]=='+'||buffer[loc]=='-'){
id= append(id,buffer[loc])
loc++
}
for loc<len(buffer)&&unicode.IsDigit(buffer[loc]){
id= append(id,buffer[loc])
loc++
}
}
if loc<len(buffer)&&buffer[loc]=='i'{
id= append(id,'$')
id= append(id,'i')
loc++
}
}
return constant
}



/*:124*/


//line goweave.w:688

}else if c=='\''||c=='"'||c=='`'{


/*125:*/


//line goweave.w:792
{
delim:=c/* what started the string */
section_text= section_text[0:0]

if delim=='\''&&
loc-2<len(buffer)&&loc-2>=0&&buffer[loc-2]=='@'{
section_text= append(section_text,'@')
section_text= append(section_text,'@')
}
section_text= append(section_text,delim)
for{
if loc>=len(buffer){
if!get_line(){
err_print("! Input ended in middle of string")
loc= 0
break

}else{
section_text= append(section_text,'\\','n')
}
}
l:=loc
loc++
if c= buffer[l];c==delim{
section_text= append(section_text,c)
break
}
if c=='\\'{
if loc>=len(buffer){
continue
}
section_text= append(section_text,'\\')
c= buffer[loc]
loc++
}
section_text= append(section_text,c)
}
id= section_text
return str
}



/*:125*/


//line goweave.w:690

}else if unicode.IsLetter(c)||
c=='_'&&(unicode.IsLetter(c)||unicode.IsDigit(c)){


/*123:*/


//line goweave.w:712
{
loc--
id_first:=loc
for loc<len(buffer)&&
(unicode.IsLetter(buffer[loc])||
unicode.IsDigit(buffer[loc])||
buffer[loc]=='_'){
loc++
}
id= buffer[id_first:loc]
return identifier
}



/*:123*/


//line goweave.w:693

}else if c=='@'{


/*126:*/


//line goweave.w:836
{
c= nc
loc++
switch ccode[c]{

case underline:
xref_switch= def_flag
continue
case trace:
tracing= c-'0'
continue
case xref_roman,xref_wildcard,xref_typewriter,noop,TeX_string:
c= ccode[c]
skip_restricted()
return c
case section_name:


/*127:*/


//line goweave.w:865
{
section_text= section_text[0:0]
cur_section_char= nc


/*129:*/


//line goweave.w:887

for{
if loc>=len(buffer){
if!get_line(){
err_print("! Input ended in section name")

loc= 1
break
}
if len(section_text)> 0{
section_text= append(section_text,' ')
}
}
c= buffer[loc]


/*130:*/


//line goweave.w:912

if c=='@'{
if loc+1>=len(buffer){
err_print("! Section name didn't end")
break

}
c= buffer[loc+1]
if(c=='>'){
loc+= 2
break
}
cc:=ignore
if c<int32(len(ccode)){
cc= ccode[c]
}
if cc==new_section{
err_print("! Section name didn't end")
break

}
if cc==section_name{
err_print("! Nesting of section names not allowed")
break

}
section_text= append(section_text,'@')
loc++/* now c==buffer[loc] again */
}



/*:130*/


//line goweave.w:901

loc++
if unicode.IsSpace(c){
c= ' '
if len(section_text)> 0&&section_text[len(section_text)-1]==' '{
section_text= section_text[:len(section_text)-1]
}
}
section_text= append(section_text,c)
}



/*:129*/


//line goweave.w:868

if len(section_text)> 3&&
compare_runes(section_text[len(section_text)-3:],[]rune("..."))==0{
cur_section= section_lookup(section_text[0:len(section_text)-3],
true)/* 1 means is a prefix */
}else{
cur_section= section_lookup(section_text,false)
}
xref_switch= 0
return section_name
}



/*:127*/


//line goweave.w:852

case verbatim:


/*132:*/


//line goweave.w:975
{
id_first:=loc
loc++
for loc<len(buffer){
if buffer[loc]!='@'{
loc++
continue
}
loc++
if loc==len(buffer){
break
}
if buffer[loc]=='>'{
break
}
}
if loc>=len(buffer){
err_print("! Verbatim string didn't end")

}
id= buffer[id_first:loc-1]
loc+= 1
return verbatim
}



/*:132*/


//line goweave.w:854

case ord:


/*125:*/


//line goweave.w:792
{
delim:=c/* what started the string */
section_text= section_text[0:0]

if delim=='\''&&
loc-2<len(buffer)&&loc-2>=0&&buffer[loc-2]=='@'{
section_text= append(section_text,'@')
section_text= append(section_text,'@')
}
section_text= append(section_text,delim)
for{
if loc>=len(buffer){
if!get_line(){
err_print("! Input ended in middle of string")
loc= 0
break

}else{
section_text= append(section_text,'\\','n')
}
}
l:=loc
loc++
if c= buffer[l];c==delim{
section_text= append(section_text,c)
break
}
if c=='\\'{
if loc>=len(buffer){
continue
}
section_text= append(section_text,'\\')
c= buffer[loc]
loc++
}
section_text= append(section_text,c)
}
id= section_text
return str
}



/*:125*/


//line goweave.w:856

default:
return ccode[c]
}
}



/*:126*/


//line goweave.w:695

}else if unicode.IsSpace(c){
continue/* ignore spaces and tabs */
}
mistake:


/*90:*/


//line gocommon.w:1395

switch c{
case'/':
if nc=='*'{
l:=loc
loc++
if l<=len(buffer){
return begin_comment
}
}else if nc=='/'{
l:=loc
loc++
if l<=len(buffer){
return begin_short_comment
}
}
case'+':
if nc=='+'{
l:=loc
loc++
if l<=len(buffer){
return plus_plus
}
}
case'-':
if nc=='-'{
l:=loc
loc++
if l<=len(buffer){
return minus_minus
}
}
case'.':
if nc=='.'&&loc+1<len(buffer)&&buffer[loc+1]=='.'{
loc++
l:=loc
loc++
if l<=len(buffer){
return dot_dot_dot
}
}
case'=':
if nc=='='{
l:=loc
loc++
if l<=len(buffer){
return eq_eq
}
}
case'>':
if nc=='='{
l:=loc
loc++
if l<=len(buffer){
return gt_eq
}
}else if nc=='>'{
l:=loc
loc++
if l<=len(buffer){
return gt_gt
}
}
case'<':
if nc=='<'{
l:=loc
loc++
if l<=len(buffer){
return lt_lt
}
}else if nc=='-'{
l:=loc
loc++
if l<=len(buffer){
return direct
}
}else if nc=='='{
l:=loc
loc++
if l<=len(buffer){
return lt_eq
}
}
case'&':
if nc=='&'{
l:=loc
loc++
if l<=len(buffer){
return and_and
}
}else if nc=='^'{
l:=loc
loc++
if l<=len(buffer){
return and_not
}
}

case'|':
if nc=='|'{
l:=loc
loc++
if l<=len(buffer){
return or_or
}
}
case'!':
if nc=='='{
l:=loc
loc++
if l<=len(buffer){
return not_eq
}
}
case':':
if nc=='='{
l:=loc
loc++
if l<=len(buffer){
return col_eq
}
}
}
//line goweave.w:117



/*:90*/


//line goweave.w:700

return c
}
return 0
}



/*:121*/



/*131:*/


//line goweave.w:944

func skip_restricted(){
id_first:=loc
false_alarm:
for loc<len(buffer)&&buffer[loc]!='@'{
loc++
}
id= buffer[id_first:loc]
loc++
if loc>=len(buffer){
err_print("! Control text didn't end")
loc= len(buffer)

}else{
if buffer[loc]=='@'&&loc<=len(buffer){
loc++
goto false_alarm
}
l:=loc
loc++
if buffer[l]!='>'{
err_print("! Control codes are forbidden in control text")

}
}
}



/*:131*/



/*134:*/


//line goweave.w:1016

func phase_one(){
phase= 1
reset_input()
section_count= 0
skip_limbo()
change_exists= false
for!input_has_ended{


/*135:*/


//line goweave.w:1032

{
section_count++
changed_section[section_count]= changing
/* it will become 1 if any line changes */
if loc-1<len(buffer)&&buffer[loc-1]=='*'&&show_progress(){
fmt.Printf("*%d",section_count)
os.Stdout.Sync()/* print a progress report */
}


/*138:*/


//line goweave.w:1125

for{
next_control= skip_TeX()
switch next_control{
case underline:
xref_switch= def_flag
continue
case trace:
tracing= buffer[loc-1]-'0'
continue
case'|':
Go_xref(section_name)
case xref_roman,xref_wildcard,xref_typewriter,noop,section_name:
loc-= 2
next_control= get_next()/* scan to \.{@>} */
if next_control>=xref_roman&&next_control<=xref_typewriter{


/*139:*/


//line goweave.w:1150

{
i:=0
j:=0
for i<len(id){
if id[i]=='@'{
i++
}
id[j]= id[i]
j++
i++
}
for j<i{
id[j]= ' '/* clean up in case of error message display */
j++
}
}



/*:139*/


//line goweave.w:1141

new_xref(id_lookup(id,next_control-identifier))
}
}
if next_control>=format_code{
break
}
}



/*:138*/


//line goweave.w:1041



/*141:*/


//line goweave.w:1187

for next_control<=format_code{


/*142:*/


//line goweave.w:1198
{
next_control= get_next()
if next_control==identifier{
lhs= id_lookup(id,normal)
name_dir[lhs].ilk= normal
if xref_switch!=0{
new_xref(lhs)
}
next_control= get_next()
if next_control==identifier{
rhs= id_lookup(id,normal)
name_dir[lhs].ilk= name_dir[rhs].ilk
if unindexed(lhs){
/* retain only underlined entries */
var r int32= 0
for q:=name_dir[lhs].xref;q>=0;q= xmem[q].xlink{
if xmem[q].num<def_flag{
if r!=0{
xmem[r].xlink= xmem[q].xlink
}else{
name_dir[lhs].xref= xmem[q].xlink
}
}else{
r= q
}
}
}
next_control= get_next()
}
}
}



/*:142*/


//line goweave.w:1189

outer_xref()
}



/*:141*/


//line goweave.w:1042



/*144:*/


//line goweave.w:1253

if next_control<=section_name{/* begin_code or section_name */
if next_control==begin_code{
section_xref_switch= 0
}else{
section_xref_switch= def_flag
if cur_section_char=='('&&cur_section!=-1{
set_file_flag(cur_section)
}
}
for{
if next_control==section_name&&cur_section!=-1{
new_section_xref(cur_section)
}
next_control= get_next()
outer_xref()
if next_control> section_name{
break
}
}
}



/*:144*/


//line goweave.w:1043

if changed_section[section_count]{
change_exists= true
}
}



/*:135*/


//line goweave.w:1024

}
changed_section[section_count]= change_exists
/* the index changes if anything does */
phase= 2/* prepare for second phase */


/*147:*/


//line goweave.w:1320

section_check(name_root)



/*:147*/


//line goweave.w:1029

}



/*:134*/



/*136:*/


//line goweave.w:1070

/* makes cross-references for \GO/ identifiers */
func Go_xref(spec_ctrl rune){
for next_control<format_code||next_control==spec_ctrl{
if next_control>=identifier&&next_control<=xref_typewriter{
if next_control> identifier{


/*139:*/


//line goweave.w:1150

{
i:=0
j:=0
for i<len(id){
if id[i]=='@'{
i++
}
id[j]= id[i]
j++
i++
}
for j<i{
id[j]= ' '/* clean up in case of error message display */
j++
}
}



/*:139*/


//line goweave.w:1076

}
p:=id_lookup(id,next_control-identifier)
/* a referenced name */
new_xref(p)
}
if next_control==section_name{
section_xref_switch= cite_flag
new_section_xref(cur_section)
}
next_control= get_next()
if next_control=='|'||next_control==begin_comment||
next_control==begin_short_comment{
return
}
}
}



/*:136*/



/*137:*/


//line goweave.w:1098

/* extension of Go_xref */
func outer_xref(){
for next_control<format_code{
if next_control!=begin_comment&&next_control!=begin_short_comment{
Go_xref(ignore)
}else{
is_long_comment:=(next_control==begin_comment)
bal,res:=copy_comment(is_long_comment,1,nil)/* brace level in comment */
next_control= '|'
for bal> 0{
Go_xref(section_name)/* do not reference section names in comments */
if next_control=='|'{
bal,res= copy_comment(is_long_comment,bal,res)
}else{
bal= 0/* an error message will occur in phase two */
}
}
}
}
}



/*:137*/



/*146:*/


//line goweave.w:1287

/* print anomalies in subtree p */
func section_check(p int32){
if p!=-1{
section_check(name_dir[p].llink)
cur_xref= name_dir[p].xref
if xmem[cur_xref].num==file_flag{
an_output= true
cur_xref= xmem[cur_xref].xlink
}else{
an_output= false
}
if xmem[cur_xref].num<def_flag{
fmt.Print("\n! Never defined: <")
print_section_name(p)
fmt.Print(">")
mark_harmless()

}
for cur_xref!=0&&xmem[cur_xref].num>=cite_flag{
cur_xref= xmem[cur_xref].xlink
}
if cur_xref==0&&!an_output{
fmt.Print("\n! Never used: <")
print_section_name(p)
fmt.Print(">")
mark_harmless()

}
section_check(name_dir[p].rlink)
}
}



/*:146*/



/*149:*/


//line goweave.w:1346

/* outputs from out_buf+1 to b,where b<=out_ptr */
func flush_buffer(b int32,per_cent bool,carryover bool){
j:=b/* pointer into out_buf */
if!per_cent{/* remove trailing blanks */
for j> 0&&out_buf[j]==' '{
j--
}
}
fmt.Fprint(active_file,string(out_buf[1:j+1]))
if per_cent{
fmt.Fprint(active_file,"%")
}
fmt.Fprint(active_file,"\n")
out_line++
if carryover{
for j> 0{
jj:=j
j--
if out_buf[jj]=='%'&&(j==0||out_buf[j]!='\\'){
out_buf[b]= '%'
b--
break
}
}
}
if b<out_ptr{
copy(out_buf[1:],out_buf[b+1:])
}
out_ptr-= b
}



/*:149*/



/*150:*/


//line goweave.w:1386

/* do this at the end of a line */
func finish_line(){
if out_ptr> 0{
flush_buffer(out_ptr,false,false)
}else{
for _,v:=range buffer{
if!unicode.IsSpace(v){
return
}
}
flush_buffer(0,false,false)
}
}



/*:150*/



/*152:*/


//line goweave.w:1421

func out(c rune){
if out_ptr>=out_buf_end{
break_out()
}
out_ptr++
out_buf[out_ptr]= c
}



/*:152*/



/*153:*/


//line goweave.w:1431

/* output characters from s to end of string */
func out_str(s string){
for _,v:=range s{
out(v)
}
}



/*:153*/



/*155:*/


//line goweave.w:1450

/* finds a way to break the output line */
func break_out(){
k:=out_ptr/* pointer into out_buf */
for{
if k==0{


/*156:*/


//line goweave.w:1476

{
fmt.Printf("\n! Line had to be broken (output l. %d):\n",out_line)

fmt.Fprint(os.Stdout,string(out_buf[1:out_ptr]))
fmt.Println()
mark_harmless()
flush_buffer(out_ptr-1,true,true)
return
}



/*:156*/


//line goweave.w:1456

}
if out_buf[k]==' '{
flush_buffer(k,false,true)
return
}
kk:=k
k--
if out_buf[kk]=='\\'&&out_buf[k]!='\\'{/* we've decreased k */
flush_buffer(k,true,true)
return
}
}
}



/*:155*/



/*157:*/


//line goweave.w:1492

func section_str(n int32)string{
s:=fmt.Sprintf("%d",n)
if changed_section[n]&&flags['c']{
s+= "\\*"

}
return s
}



/*:157*/



/*158:*/


//line goweave.w:1505

func out_name(p int32,quote_xalpha bool){
out('{')
for _,v:=range name_dir[p].name{
if v=='_'&&quote_xalpha{
out('\\')
}


out(v)
}
out('}')
}



/*:158*/



/*159:*/


//line goweave.w:1532

func copy_limbo(){
for{
if loc>=len(buffer){
finish_line()
if!get_line(){
return
}
}
for;loc<len(buffer)&&buffer[loc]!='@';loc++{
out(buffer[loc])
}
l:=loc
loc++
if l<len(buffer){
c:=' '
if loc<len(buffer){
c= buffer[loc]
loc++
}
if ccode[c]==new_section{
break
}
switch ccode[c]{
case'@':
out('@')
case noop:
skip_restricted()
case format_code:
if get_next()==identifier{
get_next()
}
if loc>=len(buffer){
get_line()/* avoid blank lines in output */
}
/* the operands of \.{@s} are ignored on this pass */
default:
err_print("! Double @ should be used in limbo")

out('@')
}
}
}
}



/*:159*/



/*160:*/


//line goweave.w:1585

func copy_TeX()rune{
for{
if loc>=len(buffer){
finish_line()
if!get_line(){
return new_section
}
}
c:=buffer[loc]
loc++
for c!='|'&&c!='@'{
out(c)
if out_ptr==1&&unicode.IsSpace(c){
out_ptr--
}
if loc==len(buffer){
break
}
c= buffer[loc]
loc++
}
if c=='|'{
return'|'
}
if c=='@'&&len(buffer)==1{
return new_section
}
if loc<len(buffer){
l:=loc
loc++
return ccode[buffer[l]]
}
}
return 0
}



/*:160*/



/*161:*/


//line goweave.w:1629

/* copies \TEX/ code in comments */
func copy_comment(
is_long_comment bool,
bal int,/* brace balance */
tok_mem[]interface{})(int,[]interface{}){
for{
if loc>=len(buffer){
if is_long_comment{
if!get_line(){
err_print("! Input ended in mid-comment")

loc= 1
goto done
}
}else{
if bal> 1{
err_print("! Missing } in comment")

}
goto done
}
}
c:=buffer[loc]
loc++
if c=='|'{
return bal,tok_mem
}
if is_long_comment{


/*162:*/


//line goweave.w:1685

if c=='*'&&loc<len(buffer)&&buffer[loc]=='/'{
loc++
if bal> 1{
err_print("! Missing } in comment")

}
goto done
}



/*:162*/


//line goweave.w:1658

}
if phase==2{
if c> 0177{
tok_mem= append(tok_mem,quoted_char)
}
tok_mem= append(tok_mem,c)
}


/*163:*/


//line goweave.w:1695

if c=='@'{
l:=loc
loc++
if l<len(buffer)&&buffer[l]!='@'{
err_print("! Illegal use of @ in comment")

loc-= 2
if phase==2{
tok_mem[len(tok_mem)-1]= ' '
}
goto done
}
}else if c=='\\'&&loc<len(buffer)&&buffer[loc]!='@'{
if phase==2{
tok_mem= append(tok_mem,buffer[loc])
}
loc++
}



/*:163*/


//line goweave.w:1666

if c=='{'{
bal++
}else if c=='}'{
if bal> 1{
bal--
}else{
err_print("! Extra } in comment")

if phase==2{
tok_mem= tok_mem[:len(tok_mem)-1]
}
}
}
}
done:


/*164:*/


//line goweave.w:1717

if phase==2{
for bal--;bal>=0;bal--{
tok_mem= append(tok_mem,'}')
}
}
return 0,tok_mem



/*:164*/


//line goweave.w:1682

}



/*:161*/



/*169:*/


//line goweave.w:2068

/* symbolic printout of a category */
func print_cat(c int32){
fmt.Printf("%s",cat_name[c])
}



/*:169*/



/*178:*/


//line goweave.w:2206

func call(fs[]reducing){
for i:=len(fs)-1;i>=0;i--{
fs[i]()
}
}



/*:178*/



/*179:*/


//line goweave.w:2220

func one(ss[]scrap,c rune)([]scrap,reducing,bool){
m:="found"
if(tracing&4)==4{
fmt.Printf("%*cLooking for %q...\n",shift,' ',cat_name[c])
shift+= 5
defer func(){shift-= 5;fmt.Printf("%*c%q is %s\n",shift,' ',cat_name[c],m)}()
f:=cat_name[c]
fmt.Printf("%*c",shift,' ')


/*304:*/


//line goweave.w:4839

{
if(tracing&2)==2{
fmt.Printf("%s:",f)
for i,v:=range scrap_info{
if i==len(scrap_info)-len(ss){
fmt.Print(" *")
}else{
fmt.Print(" ")
}
if v.mathness%4==yes_math{
fmt.Print("+")
}else if v.mathness%4==no_math{
fmt.Print("-")
}
print_cat(v.cat)
if v.mathness/4==yes_math{
fmt.Print("+")
}else if v.mathness/4==no_math{
fmt.Print("-")
}
}
fmt.Println()
}
}



/*:304*/


//line goweave.w:2229

}
if len(ss)==0{
return ss,empty,false
}
if ss[0].cat==c{
return ss[1:],empty,true
}
switch c{
case ConstDecl:

/*193:*/


//line goweave.w:2591

if s,f1,ok:=one(ss,const_token);ok{
if s,f2,ok:=one(s,ConstSpec);ok{
return s,func(){
f2()
f1()
reduce(ss,2,ConstDecl,0,break_space,1,force)
},true
}else if s,f2,ok:=one(s,lpar);ok{
tok_mem:=append([]interface{}{},0,1)
s,f3,t,ok:=optional(s,2,pair{ConstSpec,true})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
}
if s,f4,ok:=one(s,rpar);ok{
tok_mem= append(tok_mem,2+len(f3),force)
return s,func(){
f4()
call(f3)
f2()
f1()
reduce(ss,3+len(f3),ConstDecl,tok_mem...)
},true
}
}
}



/*:193*/


//line goweave.w:2238

case TypeDecl:

/*195:*/


//line goweave.w:2639

if s,f1,ok:=one(ss,type_token);ok{
if s,f2,ok:=one(s,TypeSpec);ok{
return s,func(){
f2()
f1()
reduce(ss,2,TypeDecl,0,break_space,1,force)
},true
}else if s,f2,ok:=one(s,lpar);ok{
tok_mem:=append([]interface{}{},0,1)
s,f3,t,ok:=optional(s,2,pair{cat:TypeSpec,mand:true})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
}
if s,f4,ok:=one(s,rpar);ok{
tok_mem= append(tok_mem,2+len(f3),force)
return s,func(){
f4()
call(f3)
f2()
f1()
reduce(ss,3+len(f3),TypeDecl,tok_mem...)
},true
}
}
}



/*:195*/


//line goweave.w:2239

case VarDecl:

/*197:*/


//line goweave.w:2691

if s,f1,ok:=one(ss,var_token);ok{
if s,f2,ok:=one(s,VarSpec);ok{
return s,func(){
f2()
f1()
reduce(ss,2,VarDecl,0,break_space,1)
},true
}else if s,f2,ok:=one(s,lpar);ok{
tok_mem:=append([]interface{}{},0,1)
s,f3,t,ok:=optional(s,2,pair{cat:VarSpec,mand:true})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
}
if s,f4,ok:=one(s,rpar);ok{
tok_mem= append(tok_mem,2+len(f3),force)
return s,func(){
f4()
call(f3)
f2()
f1()
reduce(ss,3+len(f3),VarDecl,tok_mem...)
},true
}
}
}



/*:197*/


//line goweave.w:2240

case FunctionDecl:

/*201:*/


//line goweave.w:2796

if s,f1,ok:=sequence(ss,func_token,identifier,Signature);ok{
if s,f2,ok:=sequence(s,Block,semi);ok{
return s,func(){
call(f2)
call(f1)
make_underlined(ss[1].trans)
reduce(ss,5,FunctionDecl,0,break_space,1,2,3,4,big_force)
},true
}else if s,f2,ok:=one(s,semi);ok{
return s,func(){
f2()
call(f1)
make_underlined(ss[1].trans)
reduce(ss,4,FunctionDecl,0,break_space,1,2,3,big_force)
},true
}
}



/*:201*/


//line goweave.w:2241

case MethodDecl:

/*203:*/


//line goweave.w:2829

if s,f1,ok:=sequence(ss,func_token,Receiver,identifier,Signature);ok{
if s,f2,ok:=one(s,Block);ok{
return s,func(){
f2()
call(f1)
make_underlined(ss[2].trans)
reduce(ss,5,MethodDecl,0,break_space,1,break_space,2,3,4)
},true
}else{
return s,func(){
call(f1)
make_underlined(ss[2].trans)
reduce(ss,4,MethodDecl,0,break_space,1,break_space,2,3)
},true
}
}



/*:203*/


//line goweave.w:2242

case Receiver:

/*205:*/


//line goweave.w:2862

if s,f1,ok:=one(ss,lpar);ok{
if s,f2,ok:=one(s,identifier);ok{
if s,f3,ok:=sequence(s,asterisk,identifier,rpar);ok{
return s,func(){
call(f3)
f2()
f1()
reduce(ss,5,Receiver,0,1,2,3,4)
},true
}else if s,f,ok:=sequence(s,identifier,rpar);ok{
return s,func(){
call(f)
reduce(ss,4,Receiver,0,1,2,3)
},true
}else if s,f,ok:=one(s,rpar);ok{
return s,func(){
f()
reduce(ss,3,Receiver,0,1,2)
},true
}
}else if s,f,ok:=sequence(s,asterisk,identifier,rpar);ok{
return s,func(){
call(f)
reduce(ss,4,Receiver,0,1,2,3)
},true
}
}



/*:205*/


//line goweave.w:2243

case ConstSpec:

/*206:*/


//line goweave.w:2891

if s,f1,ok:=one(ss,IdentifierList);ok{
if s,f2,ok:=sequence(s,Type,eq,ExpressionList);ok{
if s,f3,ok:=one(s,semi);ok{
return s,func(){
f3()
call(f2)
f1()
reduce(ss,5,ConstSpec,0,break_space,1,break_space,2,break_space,3,4,force)
},true
}else if _,_,ok:=any(s,rpar,rbrace);ok{
return s,func(){
call(f2)
f1()
reduce(ss,4,ConstSpec,0,break_space,1,break_space,2,break_space,3,force)
},true
}
}else if s,f2,ok:=sequence(s,eq,ExpressionList);ok{
if s,f3,ok:=one(s,semi);ok{
return s,func(){
f3()
call(f2)
f1()
reduce(ss,4,ConstSpec,0,break_space,1,break_space,2,3,force)
},true
}else if _,_,ok:=any(s,rpar,rbrace);ok{
return s,func(){
call(f2)
f1()
reduce(ss,3,ConstSpec,0,break_space,1,break_space,2,force)
},true
}
}
}else if s,f,ok:=one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,ConstSpec,0,force)
},true
}



/*:206*/


//line goweave.w:2244

case TypeSpec:

/*207:*/


//line goweave.w:2931

if s,f1,ok:=sequence(ss,identifier,Type);ok{
if s,f2,ok:=one(s,semi);ok{
return s,func(){
f2()
call(f1)
make_underlined(ss[0].trans)
make_reserved(ss[0].trans)
reduce(ss,3,TypeSpec,0,break_space,1,2,force)
},true
}else if _,_,ok:=any(s,rpar,rbrace);ok{
return s,func(){
call(f1)
make_underlined(ss[0].trans)
make_reserved(ss[0].trans)
reduce(ss,2,TypeSpec,0,break_space,1,force)
},true
}
}else if s,f,ok:=one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,TypeSpec,0,force)
},true
}



/*:207*/


//line goweave.w:2245

case VarSpec:

/*208:*/


//line goweave.w:2956

if s,f1,ok:=one(ss,IdentifierList);ok{
if s,f2,ok:=one(s,Type);ok{
if s,f3,ok:=sequence(s,eq,ExpressionList);ok{
if s,f4,ok:=one(s,semi);ok{
return s,func(){
f4()
call(f3)
f2()
f1()
reduce(ss,5,VarSpec,0,break_space,1,2,3,4,force)
},true
}else if _,_,ok:=any(s,rpar,rbrace);ok{
return s,func(){
call(f3)
f2()
f1()
reduce(ss,4,VarSpec,0,break_space,1,2,3,force)
},true
}
}else if s,f3,ok:=one(s,semi);ok{
return s,func(){
f3()
f2()
f1()
reduce(ss,3,VarSpec,0,break_space,1,2,force)
},true
}else if _,_,ok:=any(s,rpar,rbrace);ok{
return s,func(){
f2()
f1()
reduce(ss,2,VarSpec,0,break_space,1,force)
},true
}
}else if s,f2,ok:=sequence(s,eq,ExpressionList);ok{
if s,f3,ok:=one(s,semi);ok{
return s,func(){
f3()
call(f2)
f1()
reduce(ss,4,VarSpec,0,1,2,3,force)
},true
}else if _,_,ok:=any(s,rpar,rbrace);ok{
return s,func(){
call(f2)
f1()
reduce(ss,3,VarSpec,0,1,2,force)
},true
}
}
}else if s,f,ok:=one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,VarSpec,0,force)
},true
}



/*:208*/


//line goweave.w:2246

case ImportSpec:

/*210:*/


//line goweave.w:3044

if s,f1,ok:=sequence(ss,identifier,str);ok{
if s,f2,ok:=one(s,semi);ok{
return s,func(){
f2()
call(f1)
make_reserved(ss[0].trans)
reduce(ss,3,ImportSpec,0,break_space,1,2,force)
},true
}else if _,_,ok:=any(s,rpar,rbrace);ok{
return s,func(){
call(f1)
make_reserved(ss[0].trans)
reduce(ss,2,ImportSpec,0,break_space,1,force)
},true
}
}else if s,f1,ok:=sequence(ss,dot,str);ok{
if s,f2,ok:=one(s,semi);ok{
return s,func(){
f2()
call(f1)
reduce(ss,3,ImportSpec,0,break_space,1,2,force)
},true
}else if _,_,ok:=any(s,rpar,rbrace);ok{
return s,func(){
call(f1)
reduce(ss,2,ImportSpec,0,break_space,1,force)
},true
}
}else if s,f1,ok:=one(ss,str);ok{
if s,f2,ok:=one(s,semi);ok{
return s,func(){
f2()
f1()
underline_import(ss[0].trans)
reduce(ss,2,ImportSpec,0,1,force)
},true
}else if _,_,ok:=any(s,rpar,rbrace);ok{
return s,func(){
f1()
underline_import(ss[0].trans)
reduce(ss,1,ImportSpec,0,force)
},true
}
}else if s,f,ok:=one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,ImportSpec,0,force)
},true
}



/*:210*/


//line goweave.w:2247

case FieldDecl:

/*211:*/


//line goweave.w:3095

if s,f1,ok:=sequence(ss,IdentifierList,Type);ok{
tok_mem:=append([]interface{}{},0,break_space,1)
c:=2
s,f2,ok:=one(s,str)
if ok{
tok_mem= append(tok_mem,break_space,2)
c++
}
if s,f3,ok:=one(s,semi);ok{
tok_mem= append(tok_mem,c,force)
c++
return s,func(){
f3()
f2()
call(f1)
reduce(ss,c,FieldDecl,tok_mem...)
},true
}else if _,_,ok:=any(s,rpar,rbrace);ok{
tok_mem= append(tok_mem,force)
return s,func(){
f2()
call(f1)
reduce(ss,c,FieldDecl,tok_mem...)
},true
}
}else if s,f1,ok:=one(ss,AnonymousField);ok{
tok_mem:=append([]interface{}{},0)
c:=1
s,f2,ok:=one(s,str)
if ok{
tok_mem= append(tok_mem,break_space,1)
c++
}
tok_mem= append(tok_mem,force)
if s,f3,ok:=one(s,semi);ok{
c++
return s,func(){
f3()
f2()
f1()
reduce(ss,c,FieldDecl,tok_mem...)
},true
}else if _,_,ok:=any(s,rpar,rbrace);ok{
return s,func(){
f2()
f1()
reduce(ss,c,FieldDecl,tok_mem...)
},true
}
}else if s,f,ok:=one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,FieldDecl,0,force)
},true
}



/*:211*/


//line goweave.w:2248

case AnonymousField:

/*212:*/


//line goweave.w:3152

if s,f,ok:=sequence(ss,asterisk,Type);ok{
return s,func(){
call(f)
reduce(ss,2,AnonymousField,0,1)
},true
}else if s,f,ok:=one(ss,Type);ok{
return s,func(){
f()
reduce(ss,1,AnonymousField,0)
},true
}



/*:212*/


//line goweave.w:2249

case Type:

/*213:*/


//line goweave.w:3165

if s,f,ok:=any(ss,
ArrayType,
StructType,
PointerType,
FunctionType,
InterfaceType,
SliceType,
MapType,
ChannelType,
QualifiedIdent);ok{
return s,func(){
f()
reduce(ss,1,Type,0)
},true
}



/*:213*/


//line goweave.w:2250

case ArrayType:

/*214:*/


//line goweave.w:3182

if s,f,ok:=sequence(ss,lbracket,Expression,rbracket,Type);ok{
return s,func(){
call(f)
reduce(ss,4,ArrayType,0,1,2,3)
},true
}



/*:214*/


//line goweave.w:2251

case StructType:

/*215:*/


//line goweave.w:3190

if s,f1,ok:=sequence(ss,struct_token,lbrace);ok{
tok_mem:=append([]interface{}{},0,1)
s,f2,t,ok:=optional(s,2,pair{cat:FieldDecl,mand:true})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
}
if s,f3,ok:=one(s,rbrace);ok{
tok_mem= append(tok_mem,2+len(f2))
return s,func(){
f3()
call(f2)
call(f1)
reduce(ss,3+len(f2),StructType,tok_mem...)
},true
}
}



/*:215*/


//line goweave.w:2252

case PointerType:

/*217:*/


//line goweave.w:3239

if s,f,ok:=sequence(ss,asterisk,Type);ok{
return s,func(){
call(f)
reduce(ss,2,PointerType,0,1)
},true
}



/*:217*/


//line goweave.w:2253

case Signature:

/*218:*/


//line goweave.w:3247

if s,f1,ok:=one(ss,Parameters);ok{
if s,f2,ok:=any(s,Type,Parameters);ok{
return s,func(){
f2()
f1()
reduce(ss,2,Signature,0,break_space,1)
},true
}else{
return s,func(){
f1()
reduce(ss,1,Signature,0)
},true
}
}else if s,f,ok:=one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,Signature,0,force)
},true
}



/*:218*/


//line goweave.w:2254

case Parameters:

/*219:*/


//line goweave.w:3268

if s,f1,ok:=one(ss,lpar);ok{
tok_mem:=append([]interface{}{},0)
s,f2,t,ok:=optional(s,1,pair{cat:ParameterList,mand:true},pair{cat:comma,mand:false})
if ok{
tok_mem= append(tok_mem,t)
}
if s,f3,ok:=one(s,rpar);ok{
tok_mem= append(tok_mem,1+len(f2))
return s,func(){
f3()
call(f2)
f1()
reduce(ss,2+len(f2),Parameters,tok_mem...)
},true
}
}else if s,f,ok:=one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,Signature,0,force)
},true
}



/*:219*/


//line goweave.w:2255

case ParameterList:

/*220:*/


//line goweave.w:3291

if s,f1,ok:=one(ss,ParameterDecl);ok{
tok_mem:=append([]interface{}{},0)
s,f2,t,ok:=optional(s,1,pair{cat:comma,mand:true},pair{cat:ParameterDecl,mand:true})
if ok{
tok_mem= append(tok_mem,t)
}
return s,func(){
call(f2)
f1()
reduce(ss,1+len(f2),ParameterList,tok_mem...)
},true
}



/*:220*/


//line goweave.w:2256

case ParameterDecl:

/*221:*/


//line goweave.w:3305

if s,f,ok:=sequence(ss,IdentifierList,dot_dot_dot,Type);ok{
return s,func(){
call(f)
reduce(ss,3,ParameterDecl,0,"\\,",1,2)
},true
}else if s,f,ok:=sequence(ss,IdentifierList,Type);ok{
return s,func(){
call(f)
reduce(ss,2,ParameterDecl,0,break_space,1)
},true
}else if s,f,ok:=sequence(ss,dot_dot_dot,Type);ok{
return s,func(){
call(f)
reduce(ss,2,ParameterDecl,0,"\\,",1)
},true
}else if s,f,ok:=one(ss,Type);ok{
return s,func(){
f()
reduce(ss,1,ParameterDecl,0)
},true
}



/*:221*/


//line goweave.w:2257

case InterfaceType:

/*222:*/


//line goweave.w:3328

if s,f1,ok:=sequence(ss,interface_token,lbrace);ok{
tok_mem:=append([]interface{}{},0,1)
s,f2,t,ok:=optional(s,2,pair{cat:MethodSpec,mand:true})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
}
if s,f3,ok:=one(s,rbrace);ok{
tok_mem= append(tok_mem,2+len(f2))
return s,func(){
f3()
call(f2)
call(f1)
reduce(ss,3+len(f2),InterfaceType,tok_mem...)
},true
}
}



/*:222*/


//line goweave.w:2258

case MethodSpec:

/*223:*/


//line goweave.w:3346

if s,f1,ok:=sequence(ss,identifier,Signature);ok{
if s,f2,ok:=one(s,semi);ok{
return s,func(){
f2()
call(f1)
reduce(ss,3,MethodSpec,0,1,2,force)
},true
}else if _,_,ok:=any(s,rpar,rbrace);ok{
return s,func(){
call(f1)
reduce(ss,2,MethodSpec,0,1,force)
},true
}
}else if s,f1,ok:=sequence(ss,Type);ok{
if s,f2,ok:=one(s,semi);ok{
return s,func(){
f2()
call(f1)
reduce(ss,2,MethodSpec,0,1,force)
},true
}else if _,_,ok:=any(s,rpar,rbrace);ok{
return s,func(){
call(f1)
reduce(ss,1,MethodSpec,0,force)
},true
}
}else if s,f,ok:=one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,MethodSpec,0,force)
},true
}



/*:223*/


//line goweave.w:2259

case SliceType:

/*224:*/


//line goweave.w:3380

if s,f,ok:=sequence(ss,lbracket,rbracket,Type);ok{
return s,func(){
call(f)
reduce(ss,3,SliceType,0,1,2)
},true
}



/*:224*/


//line goweave.w:2260

case MapType:

/*225:*/


//line goweave.w:3388

if s,f,ok:=sequence(ss,map_token,lbracket,Type,rbracket,Type);ok{
return s,func(){
call(f)
reduce(ss,5,MapType,0,1,2,3,4)
},true
}



/*:225*/


//line goweave.w:2261

case ChannelType:

/*226:*/


//line goweave.w:3396

if s,f,ok:=sequence(ss,direct,chan_token,Type);ok{
return s,func(){
call(f)
reduce(ss,3,ChannelType,0,1,break_space,2)
},true
}else if s,f1,ok:=one(ss,chan_token);ok{
if s,f2,ok:=sequence(s,direct,Type);ok{
return s,func(){
call(f2)
f1()
reduce(ss,3,ChannelType,0,1,2)
},true
}else if s,f2,ok:=one(s,Type);ok{
return s,func(){
f2()
f1()
reduce(ss,2,ChannelType,0,break_space,1)
},true
}
}



/*:226*/


//line goweave.w:2262

case IdentifierList:

/*227:*/


//line goweave.w:3418

if s,f1,ok:=one(ss,identifier);ok{
tok_mem:=append([]interface{}{},0)
s,f2,t,ok:=optional(s,1,pair{cat:comma,mand:true},pair{cat:identifier,mand:true})
if ok{
tok_mem= append(tok_mem,t)
}
return s,func(){
call(f2)
f1()
reduce(ss,1+len(f2),IdentifierList,tok_mem...)
},true
}



/*:227*/


//line goweave.w:2263

case ExpressionList:

/*228:*/


//line goweave.w:3432

if s,f1,ok:=one(ss,Expression);ok{
tok_mem:=append([]interface{}{},0)
s,f2,t,ok:=optional(s,1,pair{cat:comma,mand:true},pair{cat:Expression,mand:true})
if ok{
tok_mem= append(tok_mem,t)
}
return s,func(){
call(f2)
f1()
reduce(ss,1+len(f2),ExpressionList,tok_mem...)
},true
}



/*:228*/


//line goweave.w:2264

case Expression:

/*229:*/


//line goweave.w:3446

if s,f1,ok:=one(ss,UnaryExpr);ok{
tok_mem:=append([]interface{}{},0)
s,f2,t,ok:=optional(s,1,pair{binary_op,true},pair{UnaryExpr,true});
if ok{
tok_mem= append(tok_mem,t)
}
return s,func(){
call(f2)
f1()
reduce(ss,1+len(f2),Expression,tok_mem...)
},true
}



/*:229*/


//line goweave.w:2265

case UnaryExpr:

/*230:*/


//line goweave.w:3460

if s,f,ok:=one(ss,PrimaryExpr);ok{
return s,func(){
f()
reduce(ss,1,UnaryExpr,0)
},true
}else if s,f,ok:=sequence(ss,unary_op,UnaryExpr);ok{
return s,func(){
call(f)
reduce(ss,2,UnaryExpr,0,1)
},true
}



/*:230*/


//line goweave.w:2266

case binary_op:

/*231:*/


//line goweave.w:3473

if s,f,ok:=any(ss,rel_op,add_op,mul_op,asterisk);ok{
return s,func(){
f()
reduce(ss,1,binary_op,0)
},true
}



/*:231*/


//line goweave.w:2267

case PrimaryExpr:

/*232:*/


//line goweave.w:3481

if s,f1,ok:=any(ss,BuiltinCall,Conversion,Operand);ok{
tok_mem:=append([]interface{}{},0)
s,f2,t,ok:=optional(s,1,
pair{cat:Selector,mand:false},
pair{cat:Index,mand:false},
pair{cat:Slice,mand:false},
pair{cat:TypeAssertion,mand:false},
pair{cat:Call,mand:false});
if ok{
tok_mem= append(tok_mem,t)
}
return s,func(){
call(f2)
f1()
reduce(ss,1+len(f2),PrimaryExpr,tok_mem...)
},true

}



/*:232*/


//line goweave.w:2268

case Operand:

/*233:*/


//line goweave.w:3501

if s,f,ok:=any(ss,
CompositeLit,
FunctionLit,
MethodExpr,
str,
constant,
QualifiedIdent);ok{
return s,func(){
f()
reduce(ss,1,Operand,0)
},true
}else if s,f,ok:=sequence(ss,lpar,Expression,rpar);ok{
return s,func(){
call(f)
reduce(ss,3,Operand,0,1,2)
},true
}



/*:233*/


//line goweave.w:2269

case CompositeLit:

/*234:*/


//line goweave.w:3520

if s,f,ok:=sequence(ss,LiteralType,LiteralValue);ok{
return s,func(){
call(f)
reduce(ss,2,CompositeLit,0,1)
},true
}



/*:234*/


//line goweave.w:2270

case LiteralType:

/*235:*/


//line goweave.w:3528

if s,f,ok:=one(ss,Type);ok{
return s,func(){
f()
reduce(ss,1,LiteralType,0)
},true
}else if s,f,ok:=sequence(ss,lbracket,dot_dot_dot,rbracket,Type);ok{
return s,func(){
call(f)
reduce(ss,4,LiteralType,0,1,2,3)
},true
}



/*:235*/


//line goweave.w:2271

case LiteralValue:

/*236:*/


//line goweave.w:3541

if s,f1,ok:=one(ss,lbrace);ok{
tok_mem:=append([]interface{}{},0)
s,f2,t,ok:=optional(s,1,pair{cat:ElementList,mand:true},pair{cat:comma,mand:false})
if ok{
tok_mem= append(tok_mem,t)
}
if s,f3,ok:=one(s,rbrace);ok{
tok_mem= append(tok_mem,1+len(f2))
return s,func(){
f3()
call(f2)
f1()
reduce(ss,2+len(f2),LiteralValue,tok_mem...)
},true
}
}



/*:236*/


//line goweave.w:2272

case ElementList:

/*237:*/


//line goweave.w:3559

if s,f1,ok:=one(ss,Element);ok{
tok_mem:=append([]interface{}{},0)
s,f2,t,ok:=optional(s,1,pair{cat:comma,mand:true},pair{cat:Element,mand:true})
if ok{
tok_mem= append(tok_mem,t)
}
return s,func(){
call(f2)
f1()
reduce(ss,1+len(f2),ElementList,tok_mem...)
},true
}



/*:237*/


//line goweave.w:2273

case Element:

/*238:*/


//line goweave.w:3573

if s,f1,ok:=any(ss,identifier,Expression);ok{
if s,f2,ok:=one(s,colon);ok{
if s,f3,ok:=any(s,Expression,LiteralValue);ok{
return s,func(){
f3()
f2()
f1()
reduce(ss,3,Element,0,1,break_space,2)
},true
}
}
}
if s,f1,ok:=any(ss,Expression,LiteralValue);ok{
return s,func(){
f1()
reduce(ss,1,Element,0)
},true
}



/*:238*/


//line goweave.w:2274

case FunctionLit:

/*239:*/


//line goweave.w:3593

if s,f,ok:=sequence(ss,FunctionType,Block);ok{
return s,func(){
call(f)
reduce(ss,2,FunctionLit,0,1)
},true
}



/*:239*/


//line goweave.w:2275

case FunctionType:

/*240:*/


//line goweave.w:3601

if s,f,ok:=sequence(ss,func_token,Signature);ok{
return s,func(){
call(f)
reduce(ss,2,FunctionType,0,1)
},true
}



/*:240*/


//line goweave.w:2276

case Block:

/*241:*/


//line goweave.w:3609

if s,f1,ok:=one(ss,lbrace);ok{
tok_mem:=append([]interface{}{},0)
s,f2,t,ok:=optional(s,1,pair{cat:Statement,mand:true})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
}
if s,f3,ok:=one(s,rbrace);ok{
tok_mem= append(tok_mem,1+len(f2))
return s,func(){
f3()
call(f2)
f1()
reduce(ss,2+len(f2),Block,tok_mem...)
},true
}
}



/*:241*/


//line goweave.w:2277

case Statement:

/*243:*/


//line goweave.w:3635

if s,f,ok:=any(ss,
ImportDecl,
ConstDecl,
VarDecl,
TypeDecl,
LabeledStmt);ok{
return s,func(){
f()
reduce(ss,1,Statement,0)
},true
}else if s,f1,ok:=any(ss,
GoStmt,
ReturnStmt,
BreakStmt,
ContinueStmt,
GotoStmt,
fallthrough_token,
Block,
IfStmt,
ExprSwitchStmt,
TypeSwitchStmt,
SelectStmt,
ForStmt,
DeferStmt,
SimpleStmt);ok{
if s,f2,ok:=one(s,semi);ok{
return s,func(){
f2()
f1()
reduce(ss,2,Statement,0,1,force)
},true
}else if _,_,ok:=any(s,rpar,rbrace);ok{
return s,func(){
f1()
reduce(ss,1,Statement,0,force)
},true
}
}else if s,f,ok:=one(ss,semi);ok{
return s,func(){
f()
reduce(ss,1,Statement,0,force)
},true
}else if s,f,ok:=one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,Statement,0,force)
},true
}



/*:243*/


//line goweave.w:2278

case LabeledStmt:

/*244:*/


//line goweave.w:3685

if s,f,ok:=sequence(ss,identifier,colon,Statement);ok{
return s,func(){
call(f)
reduce(ss,3,LabeledStmt,0,1,force,2)
},true
}




/*:244*/


//line goweave.w:2279

case SimpleStmt:

/*246:*/


//line goweave.w:3700

if s,f,ok:=any(ss,
SendStmt,
IncDecStmt,
Assignment,
ShortVarDecl,
Expression);ok{
return s,func(){
f()
reduce(ss,1,SimpleStmt,0)
},true
}



/*:246*/


//line goweave.w:2280

case GoStmt:

/*247:*/


//line goweave.w:3713

if s,f,ok:=sequence(ss,go_token,Expression);ok{
return s,func(){
call(f)
reduce(ss,2,GoStmt,0,break_space,1)
},true
}



/*:247*/


//line goweave.w:2281

case ReturnStmt:

/*249:*/


//line goweave.w:3731

if s,f,ok:=sequence(ss,return_token,ExpressionList);ok{
return s,func(){
call(f)
reduce(ss,2,ReturnStmt,0,break_space,1)
},true
}else if s,f,ok:=one(ss,return_token);ok{
return s,func(){
f()
reduce(ss,1,ReturnStmt,0)
},true
}



/*:249*/


//line goweave.w:2282

case BreakStmt:

/*251:*/


//line goweave.w:3757

if s,f1,ok:=one(ss,break_token);ok{
if s,f2,ok:=one(s,identifier);ok{
return s,func(){
f2()
f1()
reduce(ss,2,BreakStmt,0,break_space,1)
},true
}else{
return s,func(){
f1()
reduce(ss,1,BreakStmt,0)
},true
}
}



/*:251*/


//line goweave.w:2283

case ContinueStmt:

/*253:*/


//line goweave.w:3794

if s,f,ok:=sequence(ss,continue_token,identifier);ok{
return s,func(){
call(f)
reduce(ss,2,ContinueStmt,0,break_space,1)
},true
}else if s,f,ok:=one(ss,continue_token);ok{
return s,func(){
f()
reduce(ss,1,ContinueStmt,0)
},true
}



/*:253*/


//line goweave.w:2284

case GotoStmt:

/*255:*/


//line goweave.w:3827

if s,f,ok:=sequence(ss,goto_token,identifier);ok{
return s,func(){
call(f)
reduce(ss,2,GotoStmt,0,break_space,1)
},true
}



/*:255*/


//line goweave.w:2285

case IfStmt:

/*257:*/


//line goweave.w:3841

if s,f1,ok:=one(ss,if_token);ok{
tok_mem:=append([]interface{}{},0)
c:=1
var f2[]reducing
f3,f4:=empty,empty
if s,f2,ok= sequence(s,SimpleStmt,semi,Expression,Block);ok{
tok_mem= append(tok_mem,break_space,c)
if len(scrap_info[c+1].trans)!=0{
tok_mem= append(tok_mem,c+1)
}else{
tok_mem= append(tok_mem,';')
}
tok_mem= append(tok_mem,break_space,c+2,break_space,c+3)
c+= 4
}else if s,f2,ok= sequence(s,SimpleStmt,semi,QualifiedIdent,Block);ok{
tok_mem= append(tok_mem,break_space,c)
if len(scrap_info[c+1].trans)!=0{
tok_mem= append(tok_mem,c+1)
}else{
tok_mem= append(tok_mem,';')
}
tok_mem= append(tok_mem,break_space,c+2,break_space,c+3)
c+= 4
}else if s,f2,ok= sequence(s,Expression,Block);ok{
tok_mem= append(tok_mem,break_space,c,break_space,c+1)
c+= 2
}else if s,f2,ok= sequence(s,QualifiedIdent,Block);ok{
tok_mem= append(tok_mem,break_space,c,break_space,c+1)
c+= 2
}else{
break
}
if s,f3,ok= one(s,else_token);ok{
if s,f4,ok= any(s,IfStmt,Block);ok{
tok_mem= append(tok_mem,break_space,c,break_space,c+1)
c+= 2
}else{
break
}
}
return s,func(){
f4()
f3()
call(f2)
f1()
reduce(ss,c,IfStmt,tok_mem...)
},true
}
break
if s,f1,ok:=one(ss,if_token);ok{
tok_mem:=append([]interface{}{},0)
c:=1
var f2,f3[]reducing
f4,f5:=empty,empty
s,f2,ok= sequence(s,SimpleStmt,semi)
if ok{
tok_mem= append(tok_mem,break_space,c)
if len(scrap_info[c+1].trans)!=0{
tok_mem= append(tok_mem,c+1)
}else{
tok_mem= append(tok_mem,';')
}
c+= 2
}
s,f3,ok= sequence(s,Expression,Block)
if ok{
tok_mem= append(tok_mem,break_space,c,break_space,c+1)
c+= 2
s,f4,ok= one(s,else_token)
if ok{
if s,f5,ok= any(s,IfStmt,Block);ok{
tok_mem= append(tok_mem,break_space,c,break_space,c+1)
c+= 2
}else{
break
}
}
return s,func(){
f5()
f4()
call(f3)
call(f2)
f1()
reduce(ss,c,IfStmt,tok_mem...)
},true
}
}




/*:257*/


//line goweave.w:2286

case ExprSwitchStmt:

/*259:*/


//line goweave.w:3953

if s,f1,ok:=one(ss,switch_token);ok{
tok_mem:=append([]interface{}{},0)
c:=1
var f2[]reducing
f3,f4:=empty,empty
if s,f2,ok= sequence(s,SimpleStmt,semi);ok{
tok_mem= append(tok_mem,break_space,c,c+1)
if len(scrap_info[c+1].trans)!=0{
tok_mem= append(tok_mem,break_space,c+1)
}else{
tok_mem= append(tok_mem,';')
}
c+= 2
}
if s,f3,ok= one(s,Expression);ok{
tok_mem= append(tok_mem,break_space,c,break_space)
c++
}
if s,f4,ok= one(s,lbrace);ok{
tok_mem= append(tok_mem,c)
c++
s,f5,t,ok:=optional(s,c,pair{cat:ExprCaseClause,mand:false})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
c+= len(f5)
}
if s,f6,ok:=one(s,rbrace);ok{
tok_mem= append(tok_mem,c)
c++
return s,func(){
f6()
call(f5)
f4()
f3()
call(f2)
f1()
reduce(ss,c,ExprSwitchStmt,tok_mem...)
},true
}
}
}




/*:259*/


//line goweave.w:2287

case ExprCaseClause:

/*260:*/


//line goweave.w:3997

if s,f1,ok:=sequence(ss,case_token,ExpressionList,colon);ok{
tok_mem:=append([]interface{}{},0,break_space,1,2)
s,f2,t,ok:=optional(s,3,pair{cat:Statement,mand:true})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
}
return s,func(){
call(f2)
call(f1)
reduce(ss,3+len(f2),ExprCaseClause,tok_mem...)
},true
}else if s,f1,ok:=sequence(ss,default_token,colon);ok{
tok_mem:=append([]interface{}{},0,1,force)
s,f2,t,ok:=optional(s,2,pair{cat:Statement,mand:true})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
}
return s,func(){
call(f2)
call(f1)
reduce(ss,2+len(f2),ExprCaseClause,tok_mem...)
},true
}else if s,f,ok:=one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,ExprCaseClause,0,force)
},true
}



/*:260*/


//line goweave.w:2288

case TypeSwitchStmt:

/*261:*/


//line goweave.w:4027

if s,f1,ok:=one(ss,switch_token);ok{
tok_mem:=append([]interface{}{},0)
c:=1
var f2[]reducing
if s,f2,ok= sequence(s,SimpleStmt,semi);ok{
tok_mem= append(tok_mem,break_space,c,c+1)
if len(scrap_info[c+1].trans)!=0{
tok_mem= append(tok_mem,break_space,c+1)
}else{
tok_mem= append(tok_mem,';')
}
c+= 2
}
if s,f3,ok:=sequence(s,TypeSwitchGuard,lbrace);ok{
tok_mem= append(tok_mem,break_space,c,break_space,c+1)
c+= 2
s,f4,t,ok:=optional(s,c,pair{cat:TypeCaseClause,mand:true})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
c+= len(f4)
}
if s,f5,ok:=one(s,rbrace);ok{
tok_mem= append(tok_mem,c)
c++
return s,func(){
f5()
call(f4)
call(f3)
call(f2)
f1()
reduce(ss,c,TypeSwitchStmt,tok_mem...)
},true
}
}
}



/*:261*/


//line goweave.w:2289

case TypeSwitchGuard:

/*262:*/


//line goweave.w:4064

if s,f,ok:=sequence(ss,identifier,col_eq,PrimaryExpr,dot,lpar,type_token,rpar);ok{
return s,func(){
call(f)
reduce(ss,7,TypeSwitchGuard,0,1,2,3,4,5,6)
},true
}else if s,f,ok:=sequence(ss,PrimaryExpr,dot,lpar,type_token,rpar);ok{
return s,func(){
call(f)
reduce(ss,5,TypeSwitchGuard,0,1,2,3,4)
},true
}



/*:262*/


//line goweave.w:2290

case TypeCaseClause:

/*263:*/


//line goweave.w:4077

if s,f1,ok:=sequence(ss,TypeSwitchCase,colon);ok{
tok_mem:=append([]interface{}{},0,1,force)
s,f2,t,ok:=optional(s,2,pair{cat:Statement,mand:true})
if ok{
tok_mem= append(tok_mem,indent,t,outdent)
}
return s,func(){
call(f2)
call(f1)
reduce(ss,2+len(f2),TypeCaseClause,tok_mem...)
},true
}else if s,f,ok:=one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,TypeCaseClause,0,force)
},true
}



/*:263*/


//line goweave.w:2291

case TypeSwitchCase:

/*264:*/


//line goweave.w:4096

if s,f1,ok:=sequence(ss,case_token);ok{
tok_mem:=append([]interface{}{},0)
if s,f2,ok:=any(s,Type,constant);ok{
tok_mem= append(tok_mem,break_space,1)
s,f3,t,ok:=optional(s,2,pair{cat:comma,mand:true},pair{cat:Type,mand:true})
if ok{
tok_mem= append(tok_mem,t)
}
return s,func(){
call(f3)
f2()
call(f1)
reduce(ss,2+len(f3),TypeSwitchCase,tok_mem...)
},true
}
}else if s,f,ok:=one(ss,default_token);ok{
return s,func(){
f()
reduce(ss,1,TypeSwitchCase,0)
},true
}else if s,f,ok:=one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,TypeSwitchCase,0,force)
},true
}



/*:264*/


//line goweave.w:2292

case SelectStmt:

/*266:*/


//line goweave.w:4163

if s,f1,ok:=sequence(ss,select_token,lbrace);ok{
tok_mem:=append([]interface{}{},0,1)
s,f2,t,ok:=optional(s,2,pair{cat:CommClause,mand:false})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
}
if s,f3,ok:=one(s,rbrace);ok{
tok_mem= append(tok_mem,2+len(f2))
return s,func(){
f3()
call(f2)
call(f1)
reduce(ss,3+len(f2),SelectStmt,tok_mem...)
},true
}
}



/*:266*/


//line goweave.w:2293

case CommClause:

/*267:*/


//line goweave.w:4181

if s,f1,ok:=sequence(ss,CommCase,colon);ok{
tok_mem:=append([]interface{}{},0,1,force)
s,f2,t,ok:=optional(s,2,pair{cat:Statement,mand:true})
if ok{
tok_mem= append(tok_mem,indent,t,outdent)
}
return s,func(){
call(f2)
call(f1)
reduce(ss,2+len(f2),CommClause,tok_mem...)
},true
}



/*:267*/


//line goweave.w:2294

case CommCase:

/*268:*/


//line goweave.w:4195

if s,f1,ok:=one(ss,case_token);ok{
if s,f2,ok:=any(s,SendStmt,RecvStmt);ok{
return s,func(){
f2()
f1()
reduce(ss,2,CommCase,0,break_space,1)
},true
}
}else if s,f,ok:=one(ss,default_token);ok{
return s,func(){
f()
reduce(ss,1,CommCase,0)
},true
}else if s,f,ok:=one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,CommCase,0,force)
},true
}



/*:268*/


//line goweave.w:2295

case RecvStmt:

/*269:*/


//line goweave.w:4216

if s,f1,ok:=one(ss,ExpressionList);ok{
if s,f2,ok:=any(s,eq,col_eq);ok{
if s,f3,ok:=one(s,Expression);ok{
return s,func(){
f3()
f2()
f1()
reduce(ss,3,RecvStmt,0,1,2)
},true
}
}
}else if s,f,ok:=one(s,Expression);ok{
return s,func(){
f()
reduce(ss,1,RecvStmt,0)
},true
}



/*:269*/


//line goweave.w:2296

case SendStmt:

/*270:*/


//line goweave.w:4235

if s,f,ok:=sequence(ss,Expression,direct,Expression);ok{
return s,func(){
call(f)
reduce(ss,3,SendStmt,0,1,2)
},true
}



/*:270*/


//line goweave.w:2297

case ForStmt:

/*273:*/


//line goweave.w:4277

if s,f1,ok:=one(ss,for_token);ok{
if s,f2,ok:=sequence(s,Expression,Block);ok{
return s,func(){
call(f2)
f1()
reduce(ss,3,ForStmt,0,break_space,1,break_space,2)
},true
}else if s,f2,ok:=sequence(s,ForClause,Block);ok{
return s,func(){
call(f2)
f1()
reduce(ss,3,ForStmt,0,break_space,1,break_space,2)
},true
}else if s,f2,ok:=sequence(s,RangeClause,Block);ok{
return s,func(){
call(f2)
f1()
reduce(ss,3,ForStmt,0,break_space,1,break_space,2)
},true
}else if s,f2,ok:=one(s,Block);ok{
return s,func(){
f2()
f1()
reduce(ss,2,ForStmt,0,1)
},true
}
}



/*:273*/


//line goweave.w:2298

case ForClause:

/*274:*/


//line goweave.w:4306

var tok_mem[]interface{}
c:=0
s,f1,ok:=one(ss,SimpleStmt)
if ok{
tok_mem= append(tok_mem,c)
c++
}
f2:=empty
s,f2,ok= one(s,semi)
if ok{
if len(scrap_info[c].trans)==0{
tok_mem= append(tok_mem,c)
}else{
tok_mem= append(tok_mem,';')
}
c++
f3:=empty
if s,f3,ok= one(s,Expression);ok{
tok_mem= append(tok_mem,break_space,c)
c++
}
if s,f4,ok:=one(s,semi);ok{
if len(scrap_info[c].trans)==0{
tok_mem= append(tok_mem,c)
}else{
tok_mem= append(tok_mem,';')
}
c++
f5:=empty
if s,f5,ok= one(s,SimpleStmt);ok{
tok_mem= append(tok_mem,break_space,c)
c++
}
return s,func(){
f5()
f4()
f3()
f2()
f1()
reduce(ss,c,ForClause,tok_mem...)
},true
}
}



/*:274*/


//line goweave.w:2299

case RangeClause:

/*275:*/


//line goweave.w:4351

if s,f1,ok:=one(ss,ExpressionList);ok{
if s,f2,ok:=any(s,eq,col_eq);ok{
if s,f3,ok:=sequence(s,range_token,Expression);ok{
return s,func(){
call(f3)
f2()
f1()
reduce(ss,4,RangeClause,0,1,2,break_space,3)
},true
}
}
}



/*:275*/


//line goweave.w:2300

case DeferStmt:

/*277:*/


//line goweave.w:4400

if s,f,ok:=sequence(ss,defer_token,Expression);ok{
return s,func(){
call(f)
reduce(ss,2,DeferStmt,0,break_space,1)
},true
}



/*:277*/


//line goweave.w:2301

case IncDecStmt:

/*279:*/


//line goweave.w:4420

if s,f1,ok:=one(ss,Expression);ok{
if s,f2,ok:=any(s,plus_plus,minus_minus);ok{
return s,func(){
f2()
f1()
reduce(ss,2,IncDecStmt,0,1)
},true
}
}



/*:279*/


//line goweave.w:2302

case Assignment:

/*281:*/


//line goweave.w:4440

if s,f,ok:=sequence(ss,ExpressionList,assign_op,ExpressionList);ok{
return s,func(){
call(f)
reduce(ss,3,Assignment,0,1,2)
},true
}



/*:281*/


//line goweave.w:2303

case assign_op:

/*283:*/


//line goweave.w:4502

if s,f,ok:=sequence(ss,binary_op,eq);ok{
return s,func(){
call(f)
reduce(ss,2,assign_op,math_rel,'{',0,'}','{',1,'}','}')
},true
}else if s,f,ok:=one(ss,eq);ok{
return s,func(){
f()
reduce(ss,1,assign_op,0)
},true
}



/*:283*/


//line goweave.w:2304

case ShortVarDecl:

/*284:*/


//line goweave.w:4515

if s,f,ok:=sequence(ss,IdentifierList,col_eq,ExpressionList);ok{
return s,func(){
call(f)
reduce(ss,3,ShortVarDecl,0,1,2)
},true
}



/*:284*/


//line goweave.w:2305

case QualifiedIdent:

/*286:*/


//line goweave.w:4541

if s,f1,ok:=one(ss,identifier);ok{
if s,f2,ok:=sequence(s,dot,identifier);ok{
return s,func(){
call(f2)
f1()
reduce(ss,3,QualifiedIdent,0,1,2)
//make\_reserved(ss[0],ss[0].cat)
},true
}else{
return s,func(){
f1()
reduce(ss,1,QualifiedIdent,0)
},true
}
}



/*:286*/


//line goweave.w:2306

case MethodExpr:

/*287:*/


//line goweave.w:4558

if s,f,ok:=sequence(ss,ReceiverType,dot,identifier);ok{
return s,func(){
call(f)
reduce(ss,3,MethodExpr,0,1,2)
},true
}



/*:287*/


//line goweave.w:2307

case ReceiverType:

/*288:*/


//line goweave.w:4566

if s,f,ok:=one(ss,Type);ok{
return s,func(){
f()
reduce(ss,1,ReceiverType,0)
},true
}else if s,f,ok:=sequence(ss,lpar,asterisk,Type,rpar);ok{
return s,func(){
call(f)
reduce(ss,4,ReceiverType,0,1,2,3)
},true
}



/*:288*/


//line goweave.w:2308

case Conversion:

/*289:*/


//line goweave.w:4579

if s,f,ok:=sequence(ss,Type,lpar,Expression,rpar);ok{
return s,func(){
call(f)
reduce(ss,4,Conversion,0,1,2,3)
},true
}



/*:289*/


//line goweave.w:2309

case BuiltinCall:

/*290:*/


//line goweave.w:4587

if s,f1,ok:=sequence(ss,identifier,lpar);ok{
tok_mem:=append([]interface{}{},0,1)
s,f2,t,ok:=optional(s,2,pair{cat:BuiltinArgs,mand:true},pair{cat:comma,mand:false})
if ok{
tok_mem= append(tok_mem,t)
}
if s,f3,ok:=one(s,rpar);ok{
tok_mem= append(tok_mem,2+len(f2))
return s,func(){
f3()
call(f2)
call(f1)
reduce(ss,3+len(f2),BuiltinCall,tok_mem...)
},true
}
}



/*:290*/


//line goweave.w:2310

case BuiltinArgs:

/*291:*/


//line goweave.w:4605

if s,f1,ok:=one(ss,Type);ok{
tok_mem:=append([]interface{}{},0)
s,f2,t,ok:=optional(s,1,pair{cat:comma,mand:true},pair{cat:ExpressionList,mand:true})
if ok{
tok_mem= append(tok_mem,t)
}
return s,func(){
call(f2)
f1()
reduce(ss,1+len(f2),BuiltinArgs,tok_mem...)
},true
}else if s,f,ok:=one(s,ExpressionList);ok{
return s,func(){
f()
reduce(ss,1,BuiltinArgs,0)
},true
}



/*:291*/


//line goweave.w:2311

case Selector:

/*292:*/


//line goweave.w:4624

if s,f,ok:=sequence(ss,dot,identifier);ok{
return s,func(){
call(f)
reduce(ss,2,Selector,0,1)
},true
}



/*:292*/


//line goweave.w:2312

case Index:

/*293:*/


//line goweave.w:4632

if s,f,ok:=sequence(ss,lbracket,Expression,rbracket);ok{
return s,func(){
call(f)
reduce(ss,3,Index,0,1,2)
},true
}



/*:293*/


//line goweave.w:2313

case Slice:

/*294:*/


//line goweave.w:4640

if s,f1,ok:=one(ss,lbracket);ok{
tok_mem:=append([]interface{}{},0)
s,f2,t1,ok:=optional(s,1,pair{cat:Expression,mand:false})
if ok{
tok_mem= append(tok_mem,t1)
}
if s,f3,ok:=one(s,colon);ok{
tok_mem= append(tok_mem,1+len(t1))
s,f4,t2,ok:=optional(s,1+len(t1)+1,pair{cat:Expression,mand:false})
if ok{
tok_mem= append(tok_mem,t2)
}
if s,f5,ok:=one(s,rbracket);ok{
tok_mem= append(tok_mem,2+len(t1)+len(t2))
return s,func(){
f5()
call(f4)
f3()
call(f2)
f1()
reduce(ss,3+len(t1)+len(t2),Slice,tok_mem...)
},true
}
}
}



/*:294*/


//line goweave.w:2314

case TypeAssertion:

/*295:*/


//line goweave.w:4667

if s,f,ok:=sequence(ss,dot,lpar,Type,rpar);ok{
return s,func(){
call(f)
reduce(ss,4,TypeAssertion,0,1,2,3)
},true
}



/*:295*/


//line goweave.w:2315

case Call:

/*296:*/


//line goweave.w:4675

if s,f1,ok:=one(ss,lpar);ok{
tok_mem:=append([]interface{}{},0)
c:=1
s,f2,ok:=one(s,ExpressionList)
f3:=empty
if ok{
tok_mem= append(tok_mem,c)
c++
s,f3,ok= one(s,dot_dot_dot)
if ok{
tok_mem= append(tok_mem,c)
c++
}
}
if s,f4,ok:=one(s,rpar);ok{
tok_mem= append(tok_mem,c)
c++
return s,func(){
f4()
f3()
f2()
f1()
reduce(ss,c,Call,tok_mem...)
},true
}
}



/*:296*/


//line goweave.w:2316

case unary_op:

/*297:*/


//line goweave.w:4703

if s,f,ok:=any(ss,asterisk,direct,add_op,mul_op);ok{
return s,func(){
f()
reduce(ss,1,unary_op,0)
},true
}



/*:297*/


//line goweave.w:2317

}
m= "not found"
return ss,empty,false
}



/*:179*/



/*180:*/


//line goweave.w:2330

func sequence(ss[]scrap,cats...rune)([]scrap,[]reducing,bool){
var fs[]reducing
s:=ss
for _,v:=range cats{
f:=empty
ok:=false
if s,f,ok= one(s,v);!ok{
return ss,nil,false
}
fs= append(fs,f)
}
return s,fs,true
}



/*:180*/



/*181:*/


//line goweave.w:2351

func any(s[]scrap,cats...rune)([]scrap,reducing,bool){
for _,v:=range cats{
if s,f,ok:=one(s,v);ok{
return s,f,ok
}
}
return s,empty,false
}




/*:181*/



/*183:*/


//line goweave.w:2378

func optional(ss[]scrap,g int,cats...pair)([]scrap,[]reducing,[]int,bool){
var trans[]int
var funcs[]reducing
ok:=false
for len(ss)> 0{
var t[]int
var fs[]reducing
s:=ss
exit:=false
for _,v:=range cats{
f:=empty
if s,f,ok= one(s,v.cat);ok{
t= append(t,g)
fs= append(fs,f)
g++
}else if v.mand{
exit= true
break
}
}
if exit||len(fs)==0{
break
}
funcs= append(funcs,fs...)
trans= append(trans,t...)
ss= s
}
ok= true
if len(funcs)==0{
ok= false
}
return ss,funcs,trans,ok
}




/*:183*/



/*185:*/


//line goweave.w:2464

func find_first_ident(p[]interface{})[]interface{}{
for i,j:=range p{
switch r:=j.(type){
case res_token:
if name_dir[r].ilk!=Type{
break
}
return p[i:i+1]
case id_token:
return p[i:i+1]
case list_token:
if q:=find_first_ident(r);q!=nil{
return q
}
case inner_list_token:
if q:=find_first_ident(r);q!=nil{
return q
}
case rune:/* char, section_token, fallthru: move on to next token */
if r==inserted{
return nil/* ignore inserts */
}
}
}
return nil
}



/*:185*/



/*186:*/


//line goweave.w:2496

/* make the first identifier in scrap_info[p].trans like c */
func make_reserved(p[]interface{}){
tok_ptr:=find_first_ident(p)
if tok_ptr==nil{
return/* this should not happen */
}
tok_ptr[0]= res_token(tok_ptr[0].(id_token))
}



/*:186*/



/*187:*/


//line goweave.w:2515

/* underline the entry for the first identifier in scrap_info[p].trans */
func make_underlined(p[]interface{}){
tok_ptr:=find_first_ident(p)
if tok_ptr==nil{
return/* this happens, for example, in case found: */
}
xref_switch= def_flag
underline_xref(tok_ptr[0].(id_token))
}



/*:187*/



/*188:*/


//line goweave.w:2531

func underline_xref(p id_token){
q:=name_dir[p].xref/* pointer to cross-reference being examined */
if flags['x']==false{
return
}
m:=section_count+xref_switch/* cross-reference value to be installed */
for q!=0{
n:=xmem[q].num/* cross-reference value being examined */
if n==m{
return
}else if m==n+def_flag{
xmem[q].num= m
return
}else if n>=def_flag&&n<m{
break
}
q= xmem[q].xlink
}


/*189:*/


//line goweave.w:2560

append_xref(0)/* this number doesn't matter */
xmem[len(xmem)-1].xlink= name_dir[p].xref
r:=int32(len(xmem)-1)/* temporary pointer for permuting cross-references */
name_dir[p].xref= r
for xmem[r].xlink!=q{
xmem[r].num= xmem[xmem[r].xlink].num
r= xmem[r].xlink
}
xmem[r].num= m/* everything from q on is left undisturbed */



/*:189*/


//line goweave.w:2550

}



/*:188*/



/*209:*/


//line goweave.w:3018

func underline_import(s[]interface{}){
var id[]rune
var i int
for i= 0;i<len(s);i++{
if c,ok:=s[i].(rune);ok&&c=='"'{
i++
break
}
}
for;i<len(s);i++{
if c,ok:=s[i].(rune);!ok||c=='/'{
id= nil
}else if c=='"'{
break
}else{
id= append(id,c)
}
}
if len(id)==0{
return
}
xref_switch= def_flag
underline_xref(id_token(id_lookup(id,normal)))
}



/*:209*/



/*300:*/


//line goweave.w:4763

func reduce(ss[]scrap,k int,c rune,s...interface{}){
var trans[]interface{}
cur_mathness:=maybe_math
init_mathness:=maybe_math

for _,t:=range s{
switch v:=t.(type){
case rune:
if v==' '||(v>=big_cancel&&v<=big_force)/* non-math token */{
if cur_mathness==maybe_math{
init_mathness= no_math
}else if cur_mathness==yes_math{
trans= append(trans,"{}$")
}
cur_mathness= no_math
}else{
if cur_mathness==maybe_math{
init_mathness= yes_math
}else if cur_mathness==no_math{
trans= append(trans,"${}")
}
cur_mathness= yes_math
}
trans= append(trans,v)
case int:
s:=ss[v]


/*298:*/


//line goweave.w:4713

switch s.mathness%4{/* left boundary */
case no_math:
if cur_mathness==maybe_math{
init_mathness= no_math
}else if cur_mathness==yes_math{
trans= append(trans,"{}$")
}
cur_mathness= s.mathness/4/* right boundary */
case yes_math:
if cur_mathness==maybe_math{
init_mathness= yes_math
}else if cur_mathness==no_math{
trans= append(trans,"${}")
}
cur_mathness= s.mathness/4/* right boundary */
case maybe_math:/* no changes */
}
trans= append(trans,s.trans...)



/*:298*/


//line goweave.w:4790

case[]int:
for _,v:=range v{
if v==-1{
continue
}
s:=ss[v]


/*298:*/


//line goweave.w:4713

switch s.mathness%4{/* left boundary */
case no_math:
if cur_mathness==maybe_math{
init_mathness= no_math
}else if cur_mathness==yes_math{
trans= append(trans,"{}$")
}
cur_mathness= s.mathness/4/* right boundary */
case yes_math:
if cur_mathness==maybe_math{
init_mathness= yes_math
}else if cur_mathness==no_math{
trans= append(trans,"${}")
}
cur_mathness= s.mathness/4/* right boundary */
case maybe_math:/* no changes */
}
trans= append(trans,s.trans...)



/*:298*/


//line goweave.w:4797

}
case string:
trans= append(trans,v)
default:
panic(fmt.Sprintf("Invalid type of translation: %T(%v)",v,v))
}
}
if init_mathness==maybe_math&&cur_mathness!=maybe_math{
init_mathness= cur_mathness
}
ss[0]= scrap{cat:c,trans:trans,mathness:4*cur_mathness+init_mathness,}
if k> 1{
copy(ss[1:],ss[k:])
ss= ss[:len(ss)-k+1]
scrap_info= scrap_info[:len(scrap_info)-k+1]
}
f:=fmt.Sprintf("reduce %q %v",cat_name[c],k)


/*304:*/


//line goweave.w:4839

{
if(tracing&2)==2{
fmt.Printf("%s:",f)
for i,v:=range scrap_info{
if i==len(scrap_info)-len(ss){
fmt.Print(" *")
}else{
fmt.Print(" ")
}
if v.mathness%4==yes_math{
fmt.Print("+")
}else if v.mathness%4==no_math{
fmt.Print("-")
}
print_cat(v.cat)
if v.mathness/4==yes_math{
fmt.Print("+")
}else if v.mathness/4==no_math{
fmt.Print("-")
}
}
fmt.Println()
}
}



/*:304*/


//line goweave.w:4815

if(tracing&8)==8{
fmt.Printf("translation of %s: %v\n",cat_name[c],trans)
}
}



/*:300*/



/*305:*/


//line goweave.w:4871

/* converts a sequence of scraps */
func translate()[]interface{}{
pp:=0


/*308:*/


//line goweave.w:4915

if(tracing&2)==2{
fmt.Printf("\nTracing after l. %d:\n",line[include_depth])
mark_harmless()

}



/*:308*/


//line goweave.w:4875



/*309:*/


//line goweave.w:4922

for i:=1;i<len(scrap_info);{
if scrap_info[i].cat==insert{
reduce(scrap_info[i-1:],2,scrap_info[i-1].cat,0,1)
continue
}
i++
}

if len(scrap_info)> 1&&scrap_info[0].cat==insert&&scrap_info[1].cat!=zero{
reduce(scrap_info,2,scrap_info[1].cat,0,1)
}




/*:309*/


//line goweave.w:4876



/*302:*/


//line goweave.w:4823

for{
if pp>=len(scrap_info){
break
}


/*184:*/


//line goweave.w:2419
{
_,f,ok:=func(ss[]scrap)([]scrap,reducing,bool){
switch ss[0].cat{
case package_token:


/*191:*/


//line goweave.w:2577

if s,f,ok:=sequence(ss,package_token,identifier);ok{
return s,func(){
call(f)
reduce(ss,2,PackageClause,0,break_space,1,big_force)
},true
}



/*:191*/


//line goweave.w:2423

case import_token:


/*199:*/


//line goweave.w:2745

if s,f1,ok:=one(ss,import_token);ok{
if s,f2,ok:=one(s,ImportSpec);ok{
return s,func(){
f2()
f1()
reduce(ss,2,ImportDecl,0,break_space,1,force)
},true
}else if s,f2,ok:=one(s,lpar);ok{
tok_mem:=append([]interface{}{},0,1)
s,f3,t,ok:=optional(s,2,pair{cat:ImportSpec,mand:true})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
}
if s,f4,ok:=one(s,rpar);ok{
tok_mem= append(tok_mem,2+len(f3),force)
return s,func(){
f4()
call(f3)
f2()
f1()
reduce(ss,3+len(f3),ImportDecl,tok_mem...)
},true
}
}
}



/*:199*/


//line goweave.w:2425

case struct_token:


/*215:*/


//line goweave.w:3190

if s,f1,ok:=sequence(ss,struct_token,lbrace);ok{
tok_mem:=append([]interface{}{},0,1)
s,f2,t,ok:=optional(s,2,pair{cat:FieldDecl,mand:true})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
}
if s,f3,ok:=one(s,rbrace);ok{
tok_mem= append(tok_mem,2+len(f2))
return s,func(){
f3()
call(f2)
call(f1)
reduce(ss,3+len(f2),StructType,tok_mem...)
},true
}
}



/*:215*/


//line goweave.w:2427

case interface_token:


/*222:*/


//line goweave.w:3328

if s,f1,ok:=sequence(ss,interface_token,lbrace);ok{
tok_mem:=append([]interface{}{},0,1)
s,f2,t,ok:=optional(s,2,pair{cat:MethodSpec,mand:true})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
}
if s,f3,ok:=one(s,rbrace);ok{
tok_mem= append(tok_mem,2+len(f2))
return s,func(){
f3()
call(f2)
call(f1)
reduce(ss,3+len(f2),InterfaceType,tok_mem...)
},true
}
}



/*:222*/


//line goweave.w:2429

case func_token:


/*201:*/


//line goweave.w:2796

if s,f1,ok:=sequence(ss,func_token,identifier,Signature);ok{
if s,f2,ok:=sequence(s,Block,semi);ok{
return s,func(){
call(f2)
call(f1)
make_underlined(ss[1].trans)
reduce(ss,5,FunctionDecl,0,break_space,1,2,3,4,big_force)
},true
}else if s,f2,ok:=one(s,semi);ok{
return s,func(){
f2()
call(f1)
make_underlined(ss[1].trans)
reduce(ss,4,FunctionDecl,0,break_space,1,2,3,big_force)
},true
}
}



/*:201*/


//line goweave.w:2431



/*203:*/


//line goweave.w:2829

if s,f1,ok:=sequence(ss,func_token,Receiver,identifier,Signature);ok{
if s,f2,ok:=one(s,Block);ok{
return s,func(){
f2()
call(f1)
make_underlined(ss[2].trans)
reduce(ss,5,MethodDecl,0,break_space,1,break_space,2,3,4)
},true
}else{
return s,func(){
call(f1)
make_underlined(ss[2].trans)
reduce(ss,4,MethodDecl,0,break_space,1,break_space,2,3)
},true
}
}



/*:203*/


//line goweave.w:2432



/*240:*/


//line goweave.w:3601

if s,f,ok:=sequence(ss,func_token,Signature);ok{
return s,func(){
call(f)
reduce(ss,2,FunctionType,0,1)
},true
}



/*:240*/


//line goweave.w:2433

default:


/*210:*/


//line goweave.w:3044

if s,f1,ok:=sequence(ss,identifier,str);ok{
if s,f2,ok:=one(s,semi);ok{
return s,func(){
f2()
call(f1)
make_reserved(ss[0].trans)
reduce(ss,3,ImportSpec,0,break_space,1,2,force)
},true
}else if _,_,ok:=any(s,rpar,rbrace);ok{
return s,func(){
call(f1)
make_reserved(ss[0].trans)
reduce(ss,2,ImportSpec,0,break_space,1,force)
},true
}
}else if s,f1,ok:=sequence(ss,dot,str);ok{
if s,f2,ok:=one(s,semi);ok{
return s,func(){
f2()
call(f1)
reduce(ss,3,ImportSpec,0,break_space,1,2,force)
},true
}else if _,_,ok:=any(s,rpar,rbrace);ok{
return s,func(){
call(f1)
reduce(ss,2,ImportSpec,0,break_space,1,force)
},true
}
}else if s,f1,ok:=one(ss,str);ok{
if s,f2,ok:=one(s,semi);ok{
return s,func(){
f2()
f1()
underline_import(ss[0].trans)
reduce(ss,2,ImportSpec,0,1,force)
},true
}else if _,_,ok:=any(s,rpar,rbrace);ok{
return s,func(){
f1()
underline_import(ss[0].trans)
reduce(ss,1,ImportSpec,0,force)
},true
}
}else if s,f,ok:=one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,ImportSpec,0,force)
},true
}



/*:210*/


//line goweave.w:2435



/*243:*/


//line goweave.w:3635

if s,f,ok:=any(ss,
ImportDecl,
ConstDecl,
VarDecl,
TypeDecl,
LabeledStmt);ok{
return s,func(){
f()
reduce(ss,1,Statement,0)
},true
}else if s,f1,ok:=any(ss,
GoStmt,
ReturnStmt,
BreakStmt,
ContinueStmt,
GotoStmt,
fallthrough_token,
Block,
IfStmt,
ExprSwitchStmt,
TypeSwitchStmt,
SelectStmt,
ForStmt,
DeferStmt,
SimpleStmt);ok{
if s,f2,ok:=one(s,semi);ok{
return s,func(){
f2()
f1()
reduce(ss,2,Statement,0,1,force)
},true
}else if _,_,ok:=any(s,rpar,rbrace);ok{
return s,func(){
f1()
reduce(ss,1,Statement,0,force)
},true
}
}else if s,f,ok:=one(ss,semi);ok{
return s,func(){
f()
reduce(ss,1,Statement,0,force)
},true
}else if s,f,ok:=one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,Statement,0,force)
},true
}



/*:243*/


//line goweave.w:2436



/*206:*/


//line goweave.w:2891

if s,f1,ok:=one(ss,IdentifierList);ok{
if s,f2,ok:=sequence(s,Type,eq,ExpressionList);ok{
if s,f3,ok:=one(s,semi);ok{
return s,func(){
f3()
call(f2)
f1()
reduce(ss,5,ConstSpec,0,break_space,1,break_space,2,break_space,3,4,force)
},true
}else if _,_,ok:=any(s,rpar,rbrace);ok{
return s,func(){
call(f2)
f1()
reduce(ss,4,ConstSpec,0,break_space,1,break_space,2,break_space,3,force)
},true
}
}else if s,f2,ok:=sequence(s,eq,ExpressionList);ok{
if s,f3,ok:=one(s,semi);ok{
return s,func(){
f3()
call(f2)
f1()
reduce(ss,4,ConstSpec,0,break_space,1,break_space,2,3,force)
},true
}else if _,_,ok:=any(s,rpar,rbrace);ok{
return s,func(){
call(f2)
f1()
reduce(ss,3,ConstSpec,0,break_space,1,break_space,2,force)
},true
}
}
}else if s,f,ok:=one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,ConstSpec,0,force)
},true
}



/*:206*/


//line goweave.w:2437



/*208:*/


//line goweave.w:2956

if s,f1,ok:=one(ss,IdentifierList);ok{
if s,f2,ok:=one(s,Type);ok{
if s,f3,ok:=sequence(s,eq,ExpressionList);ok{
if s,f4,ok:=one(s,semi);ok{
return s,func(){
f4()
call(f3)
f2()
f1()
reduce(ss,5,VarSpec,0,break_space,1,2,3,4,force)
},true
}else if _,_,ok:=any(s,rpar,rbrace);ok{
return s,func(){
call(f3)
f2()
f1()
reduce(ss,4,VarSpec,0,break_space,1,2,3,force)
},true
}
}else if s,f3,ok:=one(s,semi);ok{
return s,func(){
f3()
f2()
f1()
reduce(ss,3,VarSpec,0,break_space,1,2,force)
},true
}else if _,_,ok:=any(s,rpar,rbrace);ok{
return s,func(){
f2()
f1()
reduce(ss,2,VarSpec,0,break_space,1,force)
},true
}
}else if s,f2,ok:=sequence(s,eq,ExpressionList);ok{
if s,f3,ok:=one(s,semi);ok{
return s,func(){
f3()
call(f2)
f1()
reduce(ss,4,VarSpec,0,1,2,3,force)
},true
}else if _,_,ok:=any(s,rpar,rbrace);ok{
return s,func(){
call(f2)
f1()
reduce(ss,3,VarSpec,0,1,2,force)
},true
}
}
}else if s,f,ok:=one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,VarSpec,0,force)
},true
}



/*:208*/


//line goweave.w:2438



/*207:*/


//line goweave.w:2931

if s,f1,ok:=sequence(ss,identifier,Type);ok{
if s,f2,ok:=one(s,semi);ok{
return s,func(){
f2()
call(f1)
make_underlined(ss[0].trans)
make_reserved(ss[0].trans)
reduce(ss,3,TypeSpec,0,break_space,1,2,force)
},true
}else if _,_,ok:=any(s,rpar,rbrace);ok{
return s,func(){
call(f1)
make_underlined(ss[0].trans)
make_reserved(ss[0].trans)
reduce(ss,2,TypeSpec,0,break_space,1,force)
},true
}
}else if s,f,ok:=one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,TypeSpec,0,force)
},true
}



/*:207*/


//line goweave.w:2439



/*211:*/


//line goweave.w:3095

if s,f1,ok:=sequence(ss,IdentifierList,Type);ok{
tok_mem:=append([]interface{}{},0,break_space,1)
c:=2
s,f2,ok:=one(s,str)
if ok{
tok_mem= append(tok_mem,break_space,2)
c++
}
if s,f3,ok:=one(s,semi);ok{
tok_mem= append(tok_mem,c,force)
c++
return s,func(){
f3()
f2()
call(f1)
reduce(ss,c,FieldDecl,tok_mem...)
},true
}else if _,_,ok:=any(s,rpar,rbrace);ok{
tok_mem= append(tok_mem,force)
return s,func(){
f2()
call(f1)
reduce(ss,c,FieldDecl,tok_mem...)
},true
}
}else if s,f1,ok:=one(ss,AnonymousField);ok{
tok_mem:=append([]interface{}{},0)
c:=1
s,f2,ok:=one(s,str)
if ok{
tok_mem= append(tok_mem,break_space,1)
c++
}
tok_mem= append(tok_mem,force)
if s,f3,ok:=one(s,semi);ok{
c++
return s,func(){
f3()
f2()
f1()
reduce(ss,c,FieldDecl,tok_mem...)
},true
}else if _,_,ok:=any(s,rpar,rbrace);ok{
return s,func(){
f2()
f1()
reduce(ss,c,FieldDecl,tok_mem...)
},true
}
}else if s,f,ok:=one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,FieldDecl,0,force)
},true
}



/*:211*/


//line goweave.w:2440



/*260:*/


//line goweave.w:3997

if s,f1,ok:=sequence(ss,case_token,ExpressionList,colon);ok{
tok_mem:=append([]interface{}{},0,break_space,1,2)
s,f2,t,ok:=optional(s,3,pair{cat:Statement,mand:true})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
}
return s,func(){
call(f2)
call(f1)
reduce(ss,3+len(f2),ExprCaseClause,tok_mem...)
},true
}else if s,f1,ok:=sequence(ss,default_token,colon);ok{
tok_mem:=append([]interface{}{},0,1,force)
s,f2,t,ok:=optional(s,2,pair{cat:Statement,mand:true})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
}
return s,func(){
call(f2)
call(f1)
reduce(ss,2+len(f2),ExprCaseClause,tok_mem...)
},true
}else if s,f,ok:=one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,ExprCaseClause,0,force)
},true
}



/*:260*/


//line goweave.w:2441



/*263:*/


//line goweave.w:4077

if s,f1,ok:=sequence(ss,TypeSwitchCase,colon);ok{
tok_mem:=append([]interface{}{},0,1,force)
s,f2,t,ok:=optional(s,2,pair{cat:Statement,mand:true})
if ok{
tok_mem= append(tok_mem,indent,t,outdent)
}
return s,func(){
call(f2)
call(f1)
reduce(ss,2+len(f2),TypeCaseClause,tok_mem...)
},true
}else if s,f,ok:=one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,TypeCaseClause,0,force)
},true
}



/*:263*/


//line goweave.w:2442



/*267:*/


//line goweave.w:4181

if s,f1,ok:=sequence(ss,CommCase,colon);ok{
tok_mem:=append([]interface{}{},0,1,force)
s,f2,t,ok:=optional(s,2,pair{cat:Statement,mand:true})
if ok{
tok_mem= append(tok_mem,indent,t,outdent)
}
return s,func(){
call(f2)
call(f1)
reduce(ss,2+len(f2),CommClause,tok_mem...)
},true
}



/*:267*/


//line goweave.w:2443

}
return ss,empty,false
}(scrap_info[pp:])
if ok{
f()
}
pp++/* if no match was found, we move to the right */
}



/*:184*/


//line goweave.w:4828

}



/*:302*/


//line goweave.w:4877



/*306:*/


//line goweave.w:4886
{


/*307:*/


//line goweave.w:4904

if len(scrap_info)> 0&&(tracing&1)==1{
fmt.Printf("\nIrreducible scrap sequence in section %d:",section_count)

mark_harmless()
for i,_:=range scrap_info{
fmt.Printf(" ")
print_cat(scrap_info[i].cat)
}
}



/*:307*/


//line goweave.w:4887

var tok_mem[]interface{}
for i,v:=range scrap_info{
if i!=0{
tok_mem= append(tok_mem,' ')
}
if v.mathness%4==yes_math{
tok_mem= append(tok_mem,'$')
}
tok_mem= append(tok_mem,v.trans...)
if v.mathness/4==yes_math{
tok_mem= append(tok_mem,'$')
}
}
return tok_mem
}



/*:306*/


//line goweave.w:4878

}



/*:305*/



/*310:*/


//line goweave.w:4951

/* creates scraps from \GO/ tokens */
func Go_parse(spec_ctrl rune){
for next_control<format_code||next_control==spec_ctrl{


/*312:*/


//line goweave.w:4972

switch(next_control){
case section_name:
app_scrap(section_scrap,maybe_math,section_token(cur_section))
case str,constant,verbatim:


/*314:*/


//line goweave.w:5141

count:=-1
var tok_mem[]interface{}
if next_control==constant{
tok_mem= append(tok_mem,"\\T{")

}else if next_control==str{
count= 20
tok_mem= append(tok_mem,"\\.{")

}else{
tok_mem= append(tok_mem,"\\vb{")
}

for i:=0;i<len(id);{
if count==0{/* insert a discretionary break in a long string */
tok_mem= append(tok_mem,"}\\)\\.{")
count= 20

}
switch id[i]{
case' ','\\','#','%','$','^','{','}','~','&','_':
tok_mem= append(tok_mem,'\\')











case'@':
if i+1<len(id)&&id[i+1]=='@'{
i++
}else{
err_print("! Double @ should be used in strings")
}

}
tok_mem= append(tok_mem,id[i])
i++
count--
}
tok_mem= append(tok_mem,'}')
app_scrap(next_control,maybe_math,tok_mem...)



/*:314*/


//line goweave.w:4977

case identifier:
app_id(id)
case TeX_string:


/*315:*/


//line goweave.w:5191

tok_mem:=append([]interface{}{},"\\hbox{")
for i:=0;i<len(id);{
if id[i]=='@'{
i++
}
tok_mem= append(tok_mem,id[i])
i++
}
tok_mem= append(tok_mem,'}')
app_scrap(insert,no_math,tok_mem...)



/*:315*/


//line goweave.w:4981

case'/':
app_scrap(mul_op,yes_math,next_control)
next_control= mul_op
case'.':
app_scrap(dot,yes_math,next_control)
next_control= dot
case'_':
app_scrap(identifier,maybe_math,"\\_")
next_control= identifier
case'<':
app_scrap(rel_op,yes_math,"\\langle")
next_control= rel_op
case'>':
app_scrap(rel_op,yes_math,"\\rangle")
next_control= rel_op
case'=':
app_scrap(eq,yes_math,"\\K")
next_control= eq

case'|':
app_scrap(add_op,yes_math,"\\OR")
next_control= add_op

case'^':
app_scrap(add_op,yes_math,"\\CF")
next_control= add_op

case'%':
app_scrap(mul_op,yes_math,"\\MOD")
next_control= mul_op

case'!':
app_scrap(unary_op,yes_math,"\\R")
next_control= unary_op

case'+','-':
app_scrap(add_op,yes_math,next_control)
next_control= add_op
case'*':
app_scrap(asterisk,yes_math,next_control)
next_control= asterisk
case'&':
app_scrap(mul_op,yes_math,"\\AND")
next_control= mul_op

case ignore,xref_roman,xref_wildcard,xref_typewriter,noop:
break
case'(':
app_scrap(lpar,maybe_math,next_control)
next_control= lpar
case')':
app_scrap(rpar,maybe_math,next_control)
next_control= rpar
case'[':
app_scrap(lbracket,maybe_math,next_control)
next_control= lbracket
case']':
app_scrap(rbracket,maybe_math,next_control)
next_control= rbracket
case'{':
app_scrap(lbrace,yes_math,"\\{")
next_control= lbrace

case'}':
app_scrap(rbrace,yes_math,"\\}")
next_control= rbrace

case',':
app_scrap(comma,yes_math,next_control,opt,'9',)
next_control= comma
case';':
app_scrap(semi,maybe_math,next_control)
next_control= semi
case':':
app_scrap(colon,no_math,next_control)
next_control= colon


/*313:*/


//line goweave.w:5091

case not_eq:
app_scrap(rel_op,yes_math,"\\I")

case lt_eq:
app_scrap(rel_op,yes_math,"\\Z")

case gt_eq:
app_scrap(rel_op,yes_math,"\\G")

case eq_eq:
app_scrap(rel_op,yes_math,"\\E")

case and_and:
app_scrap(binary_op,yes_math,"\\W")

case or_or:
app_scrap(binary_op,yes_math,"\\V")

case plus_plus:
app_scrap(plus_plus,yes_math,"\\PP")

case minus_minus:
app_scrap(minus_minus,yes_math,"\\MM")

case gt_gt:
app_scrap(mul_op,yes_math,"\\GG")

case lt_lt:
app_scrap(mul_op,yes_math,"\\LL")

case dot_dot_dot:
app_scrap(dot_dot_dot,yes_math,"\\ldots")


case col_eq:
app_scrap(col_eq,yes_math,":\\K")

case direct:
app_scrap(direct,yes_math,"\\leftarrow")

case and_not:
app_scrap(mul_op,yes_math,"\\AND\\CF")





/*:313*/


//line goweave.w:5058

case thin_space:
app_scrap(insert,maybe_math,"\\,")
next_control= thin_space

case math_break:
app_scrap(insert,maybe_math,opt,"0")
next_control= insert
case line_break:
app_scrap(insert,no_math,force)
next_control= insert
case big_line_break:
app_scrap(insert,no_math,big_force)
next_control= insert
case no_line_break:
app_scrap(insert,no_math,big_cancel,noop,break_space,noop,big_cancel)
next_control= insert
case pseudo_semi:
next_control= semi
app_scrap(semi,maybe_math)
case join:
app_scrap(insert,no_math,"\\J")
next_control= insert

default:
app_scrap(insert,maybe_math,inserted,next_control)
next_control= insert
}



/*:312*/


//line goweave.w:4955

next_control= get_next()
if next_control=='|'||next_control==begin_comment||
next_control==begin_short_comment{
return
}
}
}



/*:310*/



/*311:*/


//line goweave.w:4967

func app_scrap(c int32,b int32,t...interface{}){
scrap_info= append(scrap_info,scrap{cat:c,trans:t,mathness:5*b,})
}



/*:311*/



/*316:*/


//line goweave.w:5206

func app_id(id[]rune)id_token{
p:=id_lookup(id,normal)
if name_dir[p].ilk<=custom{/* not a reserved word */
a1:=identifier
a2:=maybe_math
if name_dir[p].ilk==custom{
a2= yes_math
}
app_scrap(a1,a2,id_token(p))
}else{
if name_dir[p].ilk==binary_op||
name_dir[p].ilk==rel_op||
name_dir[p].ilk==add_op||
name_dir[p].ilk==mul_op{
app_scrap(name_dir[p].ilk,yes_math,res_token(p))
}else{
app_scrap(name_dir[p].ilk,maybe_math,res_token(p))
}
}
return id_token(p)
}



/*:316*/



/*317:*/


//line goweave.w:5234

func Go_translate()[]interface{}{
save_scraps:=scrap_info/* holds original value of scrap_info */
scrap_info= nil
Go_parse(section_name)/* get the scraps together */
if next_control!='|'{
err_print("! Missing '|' after Go text")

}
app_scrap(semi,no_math)
app_scrap(insert,maybe_math,cancel)
/* place a cancel token as a final ``comment'' */
p:=translate()/* make the translation */
scrap_info= save_scraps/* scrap the scraps */
return p
}



/*:317*/



/*318:*/


//line goweave.w:5261

/* makes scraps from \GO/ tokens and comments */
func outer_parse(){
for next_control<format_code{
var tok_mem[]interface{}
if next_control!=begin_comment&&next_control!=begin_short_comment{
Go_parse(ignore)
}else{
is_long_comment:=(next_control==begin_comment)
tok_mem= append(tok_mem,inserted)
// checking if a comment is placed at start of the line
s:=true
for i:=0;i<loc-2;i++{
if!unicode.IsSpace(buffer[i]){
s= false
break
}
}
if s{
tok_mem= append(tok_mem,force)
}
if is_long_comment{
tok_mem= append(tok_mem,"\\C{")

}else{
tok_mem= append(tok_mem,"\\SHC{")
}

var bal int
bal,tok_mem= copy_comment(is_long_comment,1,tok_mem)/* brace level in comment */
next_control= ignore
for bal> 0{
p:=tok_mem
tok_mem= nil
q:=Go_translate()/* partial comments */
tok_mem= append(tok_mem,list_token(p))
if flags['e']{
tok_mem= append(tok_mem,"\\PB{")

}
tok_mem= append(tok_mem,inner_list_token(q))
if flags['e']{
tok_mem= append(tok_mem,'}')
}
if next_control=='|'{
bal,tok_mem= copy_comment(is_long_comment,bal,tok_mem)
next_control= ignore
}else{
bal= 0/* an error has been reported */
}
}
// checking if the comment is a last entity in the line
for loc<len(buffer)&&unicode.IsSpace(buffer[loc]){
loc++
}
if loc>=len(buffer){
tok_mem= append(tok_mem,force)
}
app_scrap(insert,no_math,tok_mem...)
/* the full comment becomes a scrap */
}
}
}



/*:318*/



/*323:*/


//line goweave.w:5376
func init_stack(){
stack= make([]output_state,0,100)
cur_state.mode_field= outer
}



/*:323*/



/*325:*/


//line goweave.w:5389

/* suspends the current level */
func push_level(tokens[]interface{}){
stack= append(stack,output_state{tok_field:cur_state.tok_field,mode_field:cur_state.mode_field,})
cur_state.tok_field= tokens
}



/*:325*/



/*326:*/


//line goweave.w:5399

func pop_level()bool{
if len(stack)==0{
return false
}
p:=len(stack)-1
cur_state= stack[p]
stack= stack[:p]
return true
}



/*:326*/



/*329:*/


//line goweave.w:5428

/* returns the next token of output */
func get_output()rune{
restart:
for len(cur_state.tok_field)==0{
if!pop_level(){
return-1
}
}
val:=cur_state.tok_field[0]
cur_state.tok_field= cur_state.tok_field[1:]
switch tok:=val.(type){
case id_token:
cur_name= int32(tok)
return identifier
case res_token:
cur_name= int32(tok)
return res_word
case section_token:
cur_name= int32(tok)
return section_code
case inner_list_token:
push_level(tok)
cur_state.mode_field= inner
goto restart
case list_token:
push_level(tok)
goto restart
case rune:
return tok
case[]interface{}:
push_level(tok)
goto restart
case string:
var tok_mem[]interface{}
for _,v:=range tok{
tok_mem= append(tok_mem,v)
}
push_level(tok_mem)
goto restart
}
panic(fmt.Sprintf("Invalid type of scrap: %T(%v)",val,val))
}



/*:329*/



/*330:*/


//line goweave.w:5487

/* outputs the current token list */
func output_Go(){
save_next_control:=next_control/* values to be restored */
next_control= ignore
p:=Go_translate()/* translation of the \GO/ text */
if flags['e']{
out_str("\\PB{")
make_output(inner_list_token(p))
out('}')

}else{
make_output(inner_list_token(p))/* output the list */
}
next_control= save_next_control/* restore next_control to original state */
}



/*:330*/



/*331:*/


//line goweave.w:5506

/* outputs the equivalents of tokens */
func make_output(p interface{}){
var c int/* count of indent and outdent tokens */
tok_mem:=append([]interface{}{},p,end_translation)/* append a sentinel */
push_level(tok_mem)
tok_mem= nil
var b rune
for{
a:=get_output()/* current output byte */
reswitch:
switch a{
case end_translation:
return
case identifier,res_word:


/*332:*/


//line goweave.w:5567

out('\\')
if a==identifier{
if name_dir[cur_name].ilk==custom&&!doing_format{


/*333:*/


//line goweave.w:5600

for _,v:=range name_dir[cur_name].name{
if v=='_'{
out('x')
}else{
out(v)
}
}
break



/*:333*/


//line goweave.w:5571

}else if is_tiny(cur_name){
out('|')

}else{
delim:='.'
for _,v:=range name_dir[cur_name].name{
if unicode.IsLower(v){/* not entirely uppercase */
delim= '\\'
break
}
}
out(delim)
}


}else{
out('&')/* a==res_word */
}

if is_tiny(cur_name){
if name_dir[cur_name].name[0]=='_'{
out('\\')
}
out(name_dir[cur_name].name[0])
}else{
out_name(cur_name,true)
}



/*:332*/


//line goweave.w:5521

case section_code:


/*337:*/


//line goweave.w:5705
{
out_str("\\X")

cur_xref= name_dir[cur_name].xref
if xmem[cur_xref].num==file_flag{
an_output= true
cur_xref= xmem[cur_xref].xlink
}else{
an_output= false
}
if xmem[cur_xref].num>=def_flag{
out_str(section_str(xmem[cur_xref].num-def_flag))
if phase==3{
cur_xref= xmem[cur_xref].xlink
for xmem[cur_xref].num>=def_flag{
out_str(", ")
out_str(section_str(xmem[cur_xref].num-def_flag))
cur_xref= xmem[cur_xref].xlink
}
}
}else{
out('0')/* output the section number, or zero if it was undefined */
}
out(':')
if an_output{
out_str("\\.{")

}


/*338:*/


//line goweave.w:5740

scratch:=sprint_section_name(cur_name)
cur_section_name:=cur_name
for i:=0;i<len(scratch);{
b= scratch[i]
i++
if b=='@'{


/*339:*/


//line goweave.w:5783

ii:=i
i++
if ii<len(scratch)&&scratch[ii]!='@'{
fmt.Print("\n! Illegal control code in section name: <")

print_section_name(cur_section_name)
fmt.Print("> ")
mark_error()
}



/*:339*/


//line goweave.w:5747

}
if an_output{
switch b{
case' ','\\','#','%','$','^','{','}','~','&','_':
out('\\')
fallthrough











default:out(b)
}
}else if b!='|'{
out(b)
}else{
var buf[]rune


/*340:*/


//line goweave.w:5800

var delim rune
for{
if i>=len(scratch){
fmt.Print("\n! Go text in section name didn't end: <")

print_section_name(cur_section_name)
fmt.Print("> ")
mark_error()
break
}
b= scratch[i]
i++
if b=='@'||b=='\\'&&delim!=0{


/*341:*/


//line goweave.w:5831
{
buf= append(buf,b)
buf= append(buf,scratch[i])
i++
}



/*:341*/


//line goweave.w:5814

}else{
if b=='\''||b=='"'{
if delim==0{
delim= b
}else if delim==b{
delim= 0
}
}
if b!='|'||delim!=0{
buf= append(buf,b)
}else{
break
}
}
}



/*:340*/


//line goweave.w:5771

save_buf:=buffer
save_loc:=loc
buf= append(buf,'|')
buffer= buf
loc= 0
output_Go()
loc= save_loc
buffer= save_buf
}
}



/*:338*/


//line goweave.w:5733

if an_output{
out_str(" }")
}
out_str("\\X")
}



/*:337*/


//line goweave.w:5523

case math_rel:
out_str("\\MRL{")

case noop,inserted:
break
case cancel,big_cancel:
c= 0
b= a
for{
a= get_output()
if a==inserted{
continue
}
if a<indent&&!(b==big_cancel&&a==' ')||a> big_force{
break
}
if a==indent{
c++
}else if a==outdent{
c--
}else if a==opt{
a= get_output()
}
}


/*336:*/


//line goweave.w:5688

for;c> 0;c--{
out_str("\\1")

}
for;c<0;c++{
out_str("\\2")

}



/*:336*/


//line goweave.w:5548

goto reswitch
case indent,outdent,opt,backup,break_space,
force,big_force:


/*334:*/


//line goweave.w:5613

if a<break_space{
if cur_state.mode_field==outer{
out('\\')
out(a-cancel+'0')





if a==opt{
b= get_output()/* opt is followed by a digit */
if b!='0'||flags['f']==false{
out(b)
}else{
out_str("{-1}")/* flags['f'] encourages more \.{@\v} breaks */
}
}
}else if a==opt{
b= get_output()/* ignore digit following opt */
}
}else{


/*335:*/


//line goweave.w:5644
{
b= a
save_mode:=cur_state.mode_field/* value of cur_state.mode_field before a sequence of breaks */
c= 0
for{
a= get_output()
if a==inserted{
continue
}
if a==cancel||a==big_cancel{


/*336:*/


//line goweave.w:5688

for;c> 0;c--{
out_str("\\1")

}
for;c<0;c++{
out_str("\\2")

}



/*:336*/


//line goweave.w:5654

goto reswitch/* cancel overrides everything */
}
if a!=' '&&a<indent||a==backup||a> big_force{
if save_mode==outer{
if out_ptr> 3&&compare_runes(out_buf[out_ptr-3:out_ptr+1],[]rune("\\Y\\B"))==0{
goto reswitch
}


/*336:*/


//line goweave.w:5688

for;c> 0;c--{
out_str("\\1")

}
for;c<0;c++{
out_str("\\2")

}



/*:336*/


//line goweave.w:5662

out('\\')
out(b-cancel+'0')



if a!=end_translation{
finish_line()
}
}else if a!=end_translation&&cur_state.mode_field==inner{
out(' ')
}
goto reswitch
}
if a==indent{
c++
}else if a==outdent{
c--
}else if a==opt{
a= get_output()
}else if a> b{
b= a/* if a==' ' we have a<b */
}
}
}



/*:335*/


//line goweave.w:5635

}



/*:334*/


//line goweave.w:5552

case quoted_char:
out(cur_state.tok_field[0].(rune))
cur_state.tok_field= cur_state.tok_field[1:]
default:
out(a)/* otherwise a is an ordinary character */
}
}
}



/*:331*/



/*342:*/


//line goweave.w:5844

func phase_two(){
reset_input()
if show_progress(){
fmt.Print("\nWriting the output file...")

}
section_count= 0
format_visible= true
copy_limbo()
finish_line()
flush_buffer(0,false,false)/* insert a blank line, it looks nice */
for!input_has_ended{


/*345:*/


//line goweave.w:5892
{
section_count++


/*346:*/


//line goweave.w:5910

if loc-1>=len(buffer)||buffer[loc-1]!='*'{
out_str("\\M")

}else{
for loc<len(buffer)&&buffer[loc]==' '{
loc++
}
if loc<len(buffer)&&buffer[loc]=='*'{/* ``top'' level */
sec_depth= -1
loc++
}else{
for sec_depth= 0;loc<len(buffer)&&unicode.IsDigit(buffer[loc]);loc++{
sec_depth= sec_depth*10+buffer[loc]-'0'
}
}
for loc<len(buffer)&&buffer[loc]==' '{
loc++/* remove spaces before group title */
}
group_found= true
out_str("\\N")

{
s:=fmt.Sprintf("{%d}",sec_depth+1)
out_str(s)
}
if show_progress(){
fmt.Printf("*%d",section_count)
}
os.Stdout.Sync()/* print a progress report */
}
out_str("{")
out_str(section_str(section_count))
out_str("}")



/*:346*/


//line goweave.w:5894

save_position()


/*347:*/


//line goweave.w:5948

for{
next_control= copy_TeX()
switch next_control{
case'|':
init_stack()
output_Go()
case'@':
out('@')
case TeX_string,noop,xref_roman,xref_wildcard,xref_typewriter,section_name:
loc-= 2
next_control= get_next()/* skip to \.{@>} */
if next_control==TeX_string{
err_print("! TeX string should be in Go text only")

}
case thin_space,math_break,ord,
line_break,big_line_break,no_line_break,join,
pseudo_semi:
err_print("! You can't do that in TeX text")

}
if next_control>=format_code{
break
}
}



/*:347*/


//line goweave.w:5896



/*348:*/


//line goweave.w:5978

space_checked= false
for next_control<=format_code{/* format_code or definition */
init_stack()


/*350:*/


//line goweave.w:6026
{
doing_format= true
if buffer[loc-1]=='s'||buffer[loc-1]=='S'{
format_visible= false
}
if!space_checked{
emit_space_if_needed()
save_position()
}
tok_mem:=append([]interface{}{},"\\F")/* this will produce `\&{format }' */

next_control= get_next()
if next_control==identifier{
tok_mem= append(tok_mem,id_token(id_lookup(id,normal)),' ',break_space)/* this is syntactically separate from what follows */
next_control= get_next()
if next_control==identifier{
tok_mem= append(tok_mem,id_token(id_lookup(id,normal)))
app_scrap(Expression,maybe_math,tok_mem...)
app_scrap(semi,maybe_math)
next_control= get_next()
}
}
if len(scrap_info)!=2{
err_print("! Improper format definition")

}
}



/*:350*/


//line goweave.w:5982

outer_parse()
finish_Go(format_visible)
format_visible= true
doing_format= false
}



/*:348*/


//line goweave.w:5897



/*352:*/


//line goweave.w:6061

this_section= -1
if next_control<=section_name{
emit_space_if_needed()
init_stack()
if next_control==begin_code{
next_control= get_next()
}else{
this_section= cur_section


/*353:*/


//line goweave.w:6083

for{
next_control= get_next()
if next_control!='+'{
break
}
}/* allow optional `\.{+=}' */
var tok_mem[]interface{}
if next_control!='='&&next_control!=eq_eq{
err_print("! You need an = sign after the section name")

}else{
next_control= get_next()
}
if out_ptr> 1&&out_buf[out_ptr]=='Y'&&out_buf[out_ptr-1]=='\\'{
tok_mem= append(tok_mem,backup)
}
/* the section name will be flush left */

tok_mem= append(tok_mem,section_token(this_section))
cur_xref= name_dir[this_section].xref
if xmem[cur_xref].num==file_flag{
cur_xref= xmem[cur_xref].xlink
}
tok_mem= append(tok_mem,"${}")
if xmem[cur_xref].num!=section_count+def_flag{
tok_mem= append(tok_mem,"\\mathrel+")/*section name is multiply defined*/
this_section= -1/*so we won't give cross-reference info here*/
}
tok_mem= append(tok_mem,"\\E","{}$",force)/* output an equivalence sign */

app_scrap(dead,no_math,tok_mem...)
/* this forces a line break unless `\.{@+}' follows */



/*:353*/


//line goweave.w:6071

}
for next_control<=section_name{
outer_parse()


/*354:*/


//line goweave.w:6117

if next_control<section_name{
err_print("! You can't do that in Go text")

next_control= get_next()
}else if next_control==section_name{
app_scrap(section_scrap,maybe_math,section_token(cur_section))
next_control= get_next()
}



/*:354*/


//line goweave.w:6075

}
finish_Go(true)
}



/*:352*/


//line goweave.w:5898



/*355:*/


//line goweave.w:6130

if this_section!=-1{
cur_xref= name_dir[this_section].xref
if xmem[cur_xref].num==file_flag{
an_output= true
cur_xref= xmem[cur_xref].xlink
}else{
an_output= false
}
if xmem[cur_xref].num> def_flag{
cur_xref= xmem[cur_xref].xlink/* bypass current section number */
}
footnote(def_flag)
footnote(cite_flag)
footnote(0)
}



/*:355*/


//line goweave.w:5899



/*358:*/


//line goweave.w:6208

out_str("\\fi")
finish_line()

flush_buffer(0,false,false)/* insert a blank line, it looks nice */



/*:358*/


//line goweave.w:5900

}



/*:345*/


//line goweave.w:5857

}
}



/*:342*/



/*343:*/


//line goweave.w:5869

func save_position(){
save_line= out_line
save_place= out_ptr
}

func emit_space_if_needed(){
if save_line!=out_line||save_place!=out_ptr{
out_str("\\Y")
}
space_checked= true

}



/*:343*/



/*349:*/


//line goweave.w:5999

/* finishes a definition or a \GO/ part */
func finish_Go(visible bool/* visible is nonzero if we should produce \TEX/ output */){
if visible{
out_str("\\B")
app_scrap(insert,no_math,force)
p:=translate()/* translation of the scraps */

scrap_info= nil
make_output(list_token(p))/* output the list */
if out_ptr> 1{
if out_buf[out_ptr-1]=='\\'{



if out_buf[out_ptr]=='6'{
out_ptr-= 2
}else if out_buf[out_ptr]=='7'{
out_buf[out_ptr]= 'Y'
}
}
}
out_str("\\par")
finish_line()
}
}



/*:349*/



/*356:*/


//line goweave.w:6159

/* outputs section cross-references */
func footnote(flag int32){
if xmem[cur_xref].num<=flag{
return
}
finish_line()
out('\\')



switch flag{
case 0:
out('U')
case cite_flag:
out('Q')
default:
out('A')
}


/*357:*/


//line goweave.w:6186

q:=cur_xref/* cross-reference pointer variable */
if xmem[xmem[q].xlink].num> flag{
out('s')/* plural */
}
for{
out_str(section_str(xmem[cur_xref].num-flag))
cur_xref= xmem[cur_xref].xlink/* point to the next cross-reference to output */
if xmem[cur_xref].num<=flag{
break
}
if xmem[xmem[cur_xref].xlink].num> flag{
out_str(", ")/* not the last */
}else{
out_str("\\ET")/* the last */

if cur_xref!=xmem[q].xlink{
out('s')/* the last of more than two */
}
}
}



/*:357*/


//line goweave.w:6178

out('.')
}



/*:356*/



/*359:*/


//line goweave.w:6222

func phase_three(){
if!flags['x']{
finish_line()
out_str("\\end")

finish_line()
}else{
phase= 3
if show_progress(){
fmt.Print("\nWriting the index...")

}
finish_line()
if f,err:=os.OpenFile(idx_file_name,os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0666);
err!=nil{
fatal("! Cannot open index file ",idx_file_name)

}else{
idx_file= f
}
if change_exists&&flags['c']{


/*361:*/


//line goweave.w:6291
{
/* remember that the index is already marked as changed */
var k_section int32= 0/* runs through the sections */
for k_section++;!changed_section[k_section];k_section++{}
out_str("\\ch ")

out_str(section_str(k_section))
for k_section<section_count{
for k_section++;!changed_section[k_section];k_section++{}
out_str(", ")
out_str(section_str(k_section))
}
out('.')
}



/*:361*/


//line goweave.w:6244

finish_line()
finish_line()
}
out_str("\\inx")
finish_line()

active_file= idx_file/* change active file to the index file */


/*363:*/


//line goweave.w:6322
{
for c:=0;c<=255;c++{
bucket[c]= -1
}
for _,next_name:=range hash{
for next_name!=-1{
cur_name= next_name
next_name= name_dir[cur_name].llink
if name_dir[cur_name].xref!=0{
c:=name_dir[cur_name].name[0]
if unicode.IsUpper(c){
c= unicode.ToLower(c)
}
blink[cur_name]= bucket[c]
bucket[c]= cur_name
}
}
}
}



/*:363*/


//line goweave.w:6252



/*372:*/


//line goweave.w:6429

sort_ptr= 0
scrap_info= append(scrap_info,scrap{})
unbucket(1)
for sort_ptr> 0{
cur_depth= scrap_info[sort_ptr].cat
if blink[scrap_info[sort_ptr].head]==-1||cur_depth==infinity{


/*374:*/


//line goweave.w:6467
{
cur_name= scrap_info[sort_ptr].head
for{
out_str("\\I")



/*375:*/


//line goweave.w:6482

switch name_dir[cur_name].ilk{
case normal:
if is_tiny(cur_name){
out_str("\\|")

}else{
lowcase:=false
for _,v:=range name_dir[cur_name].name{
if unicode.IsLower(v){
lowcase= true
break
}
}
if!lowcase{
out_str("\\.")

}else{
out_str("\\\\")

}
}
case wildcard:
out_str("\\9")
out_name(cur_name,false)
goto name_done

case typewriter:
out_str("\\.")

fallthrough
case roman:
out_name(cur_name,false)
goto name_done
case custom:{
out_str("$\\")
for _,v:=range name_dir[cur_name].name{
if v=='_'{
out('x')
}else{
out(v)
}
}
out('$')
goto name_done
}
default:
out_str("\\&")

}
out_name(cur_name,true)
name_done:

/*:375*/


//line goweave.w:6472



/*377:*/


//line goweave.w:6538



/*379:*/


//line goweave.w:6567

this_xref= name_dir[cur_name].xref
cur_xref= 0
for{
next_xref= xmem[this_xref].xlink
xmem[this_xref].xlink= cur_xref
cur_xref= this_xref
this_xref= next_xref
if this_xref==0{
break
}
}



/*:379*/


//line goweave.w:6539

for{
out_str(", ")
cur_val= xmem[cur_xref].num
if cur_val<def_flag{
out_str(section_str(cur_val))
}else{
out_str("\\[")
out_str(section_str(cur_val-def_flag))
out(']')
}

cur_xref= xmem[cur_xref].xlink
if cur_xref==0{
break
}
}
out('.')
finish_line()



/*:377*/


//line goweave.w:6473

cur_name= blink[cur_name]
if cur_name==-1{
break
}
}
sort_ptr--
}



/*:374*/


//line goweave.w:6436

}else{


/*373:*/


//line goweave.w:6442
{
next_name:=scrap_info[sort_ptr].head
for{
var c rune
cur_name= next_name
next_name= blink[cur_name]
cur_byte= cur_depth
if cur_byte>=int32(len(name_dir[cur_name].name)){
c= 0/* hit end of the name */
}else{
c= name_dir[cur_name].name[cur_byte]
if unicode.IsUpper(c){
c= unicode.ToLower(c)
}
}
blink[cur_name]= bucket[c]
bucket[c]= cur_name
if next_name==-1{
break
}
}
sort_ptr--
unbucket(cur_depth+1)
}



/*:373*/


//line goweave.w:6438

}
}



/*:372*/


//line goweave.w:6253

finish_line()
active_file.Close()/* finished with idx_file */
active_file= tex_file/* switch back to tex_file for a tic */
out_str("\\fin")
finish_line()

if f,err:=os.OpenFile(scn_file_name,os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0666);
err!=nil{
fatal("! Cannot open section file ",scn_file_name)

}else{
scn_file= f
}
active_file= scn_file/* change active file to section listing file */


/*381:*/


//line goweave.w:6600
section_print(name_root)



/*:381*/


//line goweave.w:6268

finish_line()
active_file.Close()/* finished with scn_file */
active_file= tex_file
if group_found{
out_str("\\con")

}else{
out_str("\\end")

}
finish_line()
active_file.Close()
}
if show_happiness(){
fmt.Print("\nDone.")
}
check_complete()/* was all of the change file used? */
}



/*:359*/



/*371:*/


//line goweave.w:6405

/* empties buckets having depth d */
func unbucket(d int32){
/* index into bucket; cannot be a simple char because of sign
		comparison below*/
for c:=100+128;c>=0;c--{
if bucket[collate[c]]!=-1{

sort_ptr++
scrap_info= append(scrap_info,scrap{})
if sort_ptr> max_sort_ptr{
max_sort_ptr= sort_ptr
}
if c==0{
scrap_info[sort_ptr].cat= infinity
}else{
scrap_info[sort_ptr].cat= d
}
scrap_info[sort_ptr].head= bucket[collate[c]]
bucket[collate[c]]= -1
}
}
}



/*:371*/



/*380:*/


//line goweave.w:6584

/* print all section names in subtree p */
func section_print(p int32){
if p!=-1{
section_print(name_dir[p].llink)
out_str("\\I")

init_stack()
make_output(section_token(p))
footnote(cite_flag)
footnote(0)/* cur_xref was set by make_output */
finish_line()
section_print(name_dir[p].rlink)
}
}



/*:380*/



/*382:*/


//line goweave.w:6605

func print_stats(){
fmt.Print("\nMemory usage statistics:\n")

fmt.Printf("%v names\n",len(name_dir))
fmt.Println("Sorting:")
fmt.Printf("%v levels\n",max_sort_ptr)
}



/*:382*/


