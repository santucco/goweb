/*2:*/
//line goweave.w:65

package main

import(
/*14:*/
//line common.w:123

"io"
"bytes"

/*:14*//*17:*/
//line common.w:165

"bufio"

/*:17*//*21:*/
//line common.w:205

"unicode"

/*:21*//*28:*/
//line common.w:359

"fmt"

/*:28*//*35:*/
//line common.w:487

"os"
"strings"

/*:35*/
//line goweave.w:69

)

/*101:*/
//line goweave.w:181

type xref_info struct{
num int32
xlink int32
}

/*:101*//*190:*/
//line goweave.w:2171

type scrap struct{
cat int32
mathness int32
trans[]interface{}
/*390:*/
//line goweave.w:6348

head int32

/*:390*/
//line goweave.w:2176

}

type id_token int

type res_token int

type section_token int32

type list_token[]interface{}

type inner_list_token[]interface{}


/*:190*//*194:*/
//line goweave.w:2289

type pair struct{
cat int32
mand bool
}

type reducing func()

/*:194*//*342:*/
//line goweave.w:5359

type output_state struct{
tok_field[]interface{}
mode_field mode
}

/*:342*//*391:*/
//line goweave.w:6351

type sort_pointer int32

/*:391*/
//line goweave.w:72


/*1:*/
//line goweave.w:61

const banner= "This is GOWEAVE (Version 0.1)\n"

/*:1*//*5:*/
//line goweave.w:108

const(
max_names= 4000

line_length= 80

)

/*:5*//*104:*/
//line goweave.w:195

const(
cite_flag= 10240
file_flag= 3*cite_flag
def_flag= 2*cite_flag
)

/*:104*//*115:*/
//line goweave.w:436

const(
ignore rune= 00
verbatim rune= 02
underline rune= '\n'
noop rune= 0177
xref_roman rune= 0213
xref_wildcard rune= 0214
xref_typewriter rune= 0215
TeX_string rune= 0216
ord rune= 0217
join rune= 0220
thin_space rune= 0221
math_break rune= 0222
line_break rune= 0223
big_line_break rune= 0224
no_line_break rune= 0225
pseudo_semi rune= 0226
trace rune= 0232
format_code rune= 0235
begin_code rune= 0237
section_name rune= 0240
new_section rune= 0241
)

/*:115*//*124:*/
//line goweave.w:635

const(
constant rune= 0210
str rune= 0211
identifier rune= 0212
)

/*:124*//*182:*/
//line goweave.w:1779

const(
normal rune= iota
roman rune= iota
wildcard rune= iota
typewriter rune= iota
custom rune= iota
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
PackageName rune= iota
ImportDecl rune= iota
ImportSpec rune= iota
Type rune= iota
package_token rune= iota
import_token rune= iota
type_token rune= iota
interface_token rune= iota
const_token rune= iota
go_token rune= iota
return_token rune= iota
break_token rune= iota
continue_token rune= iota
goto_token rune= iota
if_token rune= iota
switch_token rune= iota
select_token rune= iota
case_token rune= iota
default_token rune= iota
for_token rune= iota
else_token rune= iota
defer_token rune= iota
func_token rune= iota
struct_token rune= iota
var_token rune= iota
range_token rune= iota
map_token rune= iota
chan_token rune= iota
dot rune= iota
eq rune= iota
binary_op rune= iota
rel_op rune= iota
add_op rune= iota
mul_op rune= iota
unary_op rune= iota
asterisk rune= iota
assign_op rune= iota

lbrace rune= iota
rbrace rune= iota
comma rune= iota
lpar rune= iota
rpar rune= iota
lbracket rune= iota
rbracket rune= iota

semi rune= iota
colon rune= iota
insert rune= iota
section_scrap rune= iota
dead rune= iota
)

/*:182*//*186:*/
//line goweave.w:2124

const(
math_rel rune= 0244
big_cancel rune= 0245
cancel rune= 0246
indent rune= 0247
outdent rune= 0250
opt rune= 0251
backup rune= 0252
break_space rune= 0253
force rune= 0254
big_force rune= 0255
quoted_char rune= 0256
end_translation rune= 0257
inserted rune= 0260
)

/*:186*//*193:*/
//line goweave.w:2279

const(
maybe_math rune= iota
yes_math rune= iota
no_math rune= iota
)

/*:193*//*341:*/
//line goweave.w:5353

const(
inner mode= 0
outer mode= 1
)

/*:341*//*348:*/
//line goweave.w:5412

const(
res_word rune= 0242
section_code rune= 0243
)

/*:348*//*397:*/
//line goweave.w:6404

const infinity= -1

/*:397*/
//line goweave.w:74



/*100:*/
//line goweave.w:157

var change_exists bool

/*:100*//*102:*/
//line goweave.w:187

var xmem[]xref_info
var xref_switch int32
var section_xref_switch int32

/*:102*//*117:*/
//line goweave.w:466

var ccode[256]rune

/*:117*//*125:*/
//line goweave.w:642

var cur_section int32
var cur_section_char rune


/*:125*//*140:*/
//line goweave.w:999

var next_control rune

/*:140*//*150:*/
//line goweave.w:1172

var lhs int32
var rhs int32
var res_wd_end int32

/*:150*//*155:*/
//line goweave.w:1271

var cur_xref int32;
var an_output bool

/*:155*//*159:*/
//line goweave.w:1321

var out_buf[line_length+1]rune
var out_ptr int32
var out_buf_end int32= line_length
var out_line int

/*:159*//*183:*/
//line goweave.w:1920

var cat_name[256]string

/*:183*//*191:*/
//line goweave.w:2190

var scrap_info[]scrap

/*:191*//*195:*/
//line goweave.w:2298

var shift= 0
var empty reducing= func(){}

/*:195*//*322:*/
//line goweave.w:4838

var tracing int32

/*:322*//*344:*/
//line goweave.w:5371

var cur_state output_state
var stack[]output_state

/*:344*//*347:*/
//line goweave.w:5409

var cur_name int32= -1

/*:347*//*366:*/
//line goweave.w:5876

var save_line int
var save_place int32
var sec_depth int32
var space_checked bool
var format_visible bool
var doing_format bool= false
var group_found bool= false

/*:366*//*374:*/
//line goweave.w:6053

var this_section int32

/*:374*//*387:*/
//line goweave.w:6312

var bucket[256]int32
var blink[max_names]int32

/*:387*//*393:*/
//line goweave.w:6357

var cur_depth int32
var cur_byte int32
var cur_val int32
var max_sort_ptr int32
var sort_ptr int32

/*:393*//*395:*/
//line goweave.w:6370


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


/*:395*//*405:*/
//line goweave.w:6565

var next_xref int32
var this_xref int32


/*:405*/
//line goweave.w:77


/*:2*//*4:*/
//line goweave.w:88

func main(){
flags['x']= true
flags['f']= true
flags['e']= true
common_init()
/*106:*/
//line goweave.w:205

xmem= append(xmem,xref_info{})
xref_switch= 0
section_xref_switch= 0

/*:106*//*118:*/
//line goweave.w:469

{
for c:= 0;c<256;c++{
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
ccode['@']= '@'
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
/*119:*/
//line goweave.w:516

ccode['0']= trace
ccode['1']= trace
ccode['2']= trace
ccode['3']= trace
ccode['4']= trace
ccode['5']= trace
ccode['6']= trace
ccode['7']= trace
ccode['8']= trace
ccode['9']= trace

/*:119*/
//line goweave.w:510


/*:118*//*162:*/
//line goweave.w:1398

out_ptr= 1
out_line= 1
active_file= tex_file
out_buf[out_ptr]= 'c'
fmt.Fprint(active_file,"\\input gowebma")

/*:162*//*166:*/
//line goweave.w:1436

out_buf[0]= '\\'

/*:166*//*184:*/
//line goweave.w:1923

for cat_index:= 0;cat_index<255;cat_index++{
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
cat_name[PackageName]= "PackageName"
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

/*:184*//*394:*/
//line goweave.w:6364

max_sort_ptr= 0

/*:394*/
//line goweave.w:94

if show_banner(){
fmt.Print(banner)
}
/*112:*/
//line goweave.w:341


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

/*:112*/
//line goweave.w:98

phase_one()
phase_two()
phase_three()
os.Exit(wrap_up())
}

/*:4*//*7:*/
//line common.w:50

const(
/*11:*/
//line common.w:93

and_and rune= 04
lt_lt rune= 020
gt_gt rune= 021
plus_plus rune= 0200
minus_minus rune= 0201
col_eq rune= 0207
not_eq rune= 032
lt_eq rune= 034
gt_eq rune= 035
eq_eq rune= 036
or_or rune= 037
dot_dot_dot rune= 0202
begin_comment rune= '\t'
and_not rune= 010
direct rune= 0203
begin_short_comment rune= 031

/*:11*//*32:*/
//line common.w:431

max_sections= 2000



/*:32*//*43:*/
//line common.w:645

hash_size= 353

/*:43*//*57:*/
//line common.w:791

less= 0
equal= 1
greater= 2
prefix= 3
extension= 4

/*:57*//*66:*/
//line common.w:1013

bad_extension= 5

/*:66*//*68:*/
//line common.w:1076

spotless= 0
harmless_message= 1
error_message= 2
fatal_message= 3

/*:68*/
//line common.w:52

)

/*13:*/
//line common.w:117

var buffer[]rune
var loc int= 0
var section_text[]rune
var id[]rune

/*:13*//*18:*/
//line common.w:168

var include_depth int
var file[]*bufio.Reader
var change_file*bufio.Reader
var file_name[]string

var change_file_name string= "/dev/null"
var alt_file_name string
var line[]int
var change_line int
var change_depth int
var input_has_ended bool
var changing bool

/*:18*//*33:*/
//line common.w:436

var section_count int32
var changed_section[max_sections]bool
var change_pending bool

var print_where bool= false

/*:33*//*41:*/
//line common.w:620

type name_info struct{
name[]rune
/*42:*/
//line common.w:634

llink int32

/*:42*//*51:*/
//line common.w:717

ispref bool
rlink int32


/*:51*//*99:*/
//line goweave.w:149

ilk int32

/*:99*//*105:*/
//line goweave.w:202

xref int32

/*:105*/
//line common.w:623

}
type name_index int
var name_dir[]name_info
var name_root int32

/*:41*//*44:*/
//line common.w:649

var hash[hash_size]int32
var h int32

/*:44*//*71:*/
//line common.w:1096

var history int= spotless

/*:71*//*87:*/
//line common.w:1274

var go_file_name string
var tex_file_name string
var idx_file_name string
var scn_file_name string
var flags[128]bool

/*:87*//*95:*/
//line common.w:1417

var go_file io.WriteCloser
var tex_file io.WriteCloser
var idx_file io.WriteCloser
var scn_file io.WriteCloser
var active_file io.WriteCloser

/*:95*/
//line common.w:55

/*8:*/
//line common.w:64
var phase int

/*:8*//*19:*/
//line common.w:187

var change_buffer[]rune

/*:19*/
//line common.w:56


/*:7*//*9:*/
//line common.w:70

func common_init(){
/*45:*/
//line common.w:653

for i,_:= range hash{
hash[i]= -1
}

/*:45*//*52:*/
//line common.w:722

name_root= -1

/*:52*/
//line common.w:72

/*88:*/
//line common.w:1285

flags['b']= true
flags['h']= true
flags['p']= true

/*:88*/
//line common.w:73

/*96:*/
//line common.w:1424

scan_args()
/*412:*/
//line goweave.w:6627

if f,err:= os.OpenFile(tex_file_name,os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0666);
err!=nil{
fatal("! Cannot open output file ",tex_file_name)
}else{
tex_file= f
}


/*:412*/
//line common.w:1426


/*:96*/
//line common.w:74

}


/*:9*//*15:*/
//line common.w:128


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

/*:15*//*20:*/
//line common.w:197

func prime_the_change_buffer(){
change_buffer= nil
/*22:*/
//line common.w:212

for true{
change_line++
if err:= input_ln(change_file);err!=nil{
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

/*:22*/
//line common.w:200

/*23:*/
//line common.w:239

for true{
change_line++
if err:= input_ln(change_file);err!=nil{
err_print("! Change file ended after @x")

return
}
if len(buffer)!=0{
break
}
}

/*:23*/
//line common.w:201

/*24:*/
//line common.w:252

{
change_buffer= buffer
buffer= nil
}

/*:24*/
//line common.w:202

}

/*:20*//*25:*/
//line common.w:273

func if_section_start_make_pending(b bool){
for loc= 0;loc<len(buffer)&&unicode.IsSpace(buffer[loc]);loc++{}
if len(buffer)>=2&&buffer[0]=='@'&&(unicode.IsSpace(buffer[1])||buffer[1]=='*'){
change_pending= b
}
}

/*:25*//*26:*/
//line common.w:282

func compare_runes(l[]rune,r[]rune)int{
i:= 0
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

/*:26*//*27:*/
//line common.w:305


func check_change(){
n:= 0
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
if err:= input_ln(change_file);err!=nil{
err_print("! Change file ended before @y")

change_buffer= nil
changing= false
return
}
if(len(buffer)> 1&&buffer[0]=='@'){
var xyz_code rune
if unicode.IsUpper(buffer[1]){
xyz_code= unicode.ToLower(buffer[1])
}else{
xyz_code= buffer[1]
}
/*29:*/
//line common.w:362

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

/*:29*/
//line common.w:338

}
/*24:*/
//line common.w:252

{
change_buffer= buffer
buffer= nil
}

/*:24*/
//line common.w:340

changing= false
line[include_depth]++
for input_ln(file[include_depth])!=nil{
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

/*:27*//*30:*/
//line common.w:382

func reset_input(){
loc= 0
file= file[:0]
/*31:*/
//line common.w:401

if wf,err:= os.Open(file_name[0]);err!=nil{
file_name[0]= alt_file_name
if wf,err= os.Open(file_name[0]);err!=nil{
fatal("! Cannot open input file ",file_name[0])

}else{
file= append(file,bufio.NewReader(wf))
}
}else{
file= append(file,bufio.NewReader(wf))
}
if cf,err:= os.Open(change_file_name);err!=nil{
fatal("! Cannot open change file ",change_file_name)

}else{
change_file= bufio.NewReader(cf)
}

/*:31*/
//line common.w:386

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

/*:30*//*34:*/
//line common.w:444

func get_line()bool{
restart:
if changing&&include_depth==change_depth{
/*38:*/
//line common.w:565
{
change_line++
if input_ln(change_file)!=nil{
err_print("! Change file ended without @z")

buffer= append(buffer,[]rune("@z")...)
}
if len(buffer)> 0{
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

/*:38*/
//line common.w:448

}
if!changing||include_depth> change_depth{
/*37:*/
//line common.w:535
{
line[include_depth]++
for input_ln(file[include_depth])!=nil{
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

/*:37*/
//line common.w:451

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

include_depth++
/*36:*/
//line common.w:491
{
l:= loc
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


if f,err:= os.Open(file_name[include_depth]);err==nil{
file= append(file,bufio.NewReader(f))
line= append(line,0)
print_where= true
goto restart
}
temp_file_name:= os.Getenv("GOWEBINPUTS")
if len(temp_file_name)!=0{

for _,fn:= range strings.Split(temp_file_name,":"){
file_name[include_depth]= fn+"/"+file_name[include_depth]
if f,err:= os.Open(file_name[include_depth]);err==nil{
file= append(file,bufio.NewReader(f))
line= append(line,0)
print_where= true
goto restart
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

/*:36*/
//line common.w:472

}
return true
}

/*:34*//*39:*/
//line common.w:600

func check_complete(){
if len(change_buffer)> 0{
buffer= change_buffer
change_buffer= nil
changing= true
change_depth= include_depth
loc= 0
err_print("! Change file entry did not match")

}
}

/*:39*//*46:*/
//line common.w:660


func id_lookup(
id[]rune,
t int32)int32{
/*47:*/
//line common.w:677

h:= id[0]
for i:= 1;i<len(id);i++{
h= (h+h+id[i])%hash_size
}


/*:47*/
//line common.w:665

/*48:*/
//line common.w:687

p:= hash[h]
for p!=-1&&!names_match(p,id,t){
p= name_dir[p].llink
}
if p==-1{
p:= int32(len(name_dir))
name_dir= append(name_dir,name_info{})
name_dir[p].llink= -1
init_node(p)
name_dir[p].llink= hash[h]
hash[h]= p
}

/*:48*/
//line common.w:666

if p==-1{
/*50:*/
//line common.w:705

p= int32(len(name_dir)-1)
name_dir[p].name= append(name_dir[p].name,id...)
init_p(p,t)

/*:50*/
//line common.w:668

}
return p
}

/*:46*//*54:*/
//line common.w:743

func print_section_name(p int32){
q:= p+1
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
fmt.Print("...")
}
}

/*:54*//*55:*/
//line common.w:762

func sprint_section_name(p int32)[]rune{
q:= p+1
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

/*:55*//*56:*/
//line common.w:779

func print_prefix_name(p int32){
l:= name_dir[p].name[0]
fmt.Print(string(name_dir[p].name[1:]))
if int(l)<len(name_dir[p].name){
fmt.Print("...")
}
}

/*:56*//*58:*/
//line common.w:799


func web_strcmp(
j[]rune,
k[]rune)int{
i:= 0
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

/*:58*//*60:*/
//line common.w:838


func add_section_name(
par int32,
c int,
name[]rune,
ispref bool)int32{
p:= int32(len(name_dir))
name_dir= append(name_dir,name_info{})
name_dir[p].llink= -1
init_node(p)
if ispref{
name_dir= append(name_dir,name_info{})
name_dir[p+1].llink= -1
init_node(p+1)
}
name_dir[p].ispref= ispref
name_dir[p].name= append(name_dir[p].name,int32(len(name)))
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

/*:60*//*61:*/
//line common.w:873

func extend_section_name(
p int32,
text[]rune,
ispref bool){
q:= p+1
for name_dir[q].llink!=-1{
q= name_dir[q].llink
}
np:= int32(len(name_dir))
name_dir[q].llink= np
name_dir= append(name_dir,name_info{})
name_dir[np].llink= -1
init_node(np)
name_dir[np].name= append(name_dir[np].name,int32(len(text)))
name_dir[np].name= append(name_dir[np].name,text...)
name_dir[np].ispref= ispref

}

/*:61*//*62:*/
//line common.w:898


func section_lookup(
name[]rune,
ispref bool)int32{
c:= less
p:= name_root
var q int32= -1
var r int32= -1
var par int32= -1
name_len:= len(name)
/*63:*/
//line common.w:921

for p!=-1{
c= web_strcmp(name,name_dir[p].name[1:])
if c==less||c==greater{
if r==-1{
par= p
}
if c==less{
p= name_dir[p].llink
}else{
p= name_dir[p].rlink
}
}else{
if r!=-1{
fmt.Printf("\n! Ambiguous prefix: matches <")

print_prefix_name(p)
fmt.Printf(">\n and <")
print_prefix_name(r)
err_print(">")
return 0
}
r= p
p= name_dir[p].llink
q= name_dir[r].rlink
}
if p==-1{
p= q
q= -1
}
}

/*:63*/
//line common.w:910

/*64:*/
//line common.w:953

if r==-1{
return add_section_name(par,c,name,ispref)
}

/*:64*/
//line common.w:911

/*65:*/
//line common.w:962

first,cmp:= section_name_cmp(name,r)
switch cmp{

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
default:
fmt.Printf("\n! Section name incompatible with <")

print_prefix_name(r)
fmt.Printf(">,\n which abbreviates <")
print_section_name(r)
err_print(">")
return r
}

/*:65*/
//line common.w:912

return-1
}

/*:62*//*67:*/
//line common.w:1017

func section_name_cmp(
name[]rune,
r int32)(int,int){
q:= r+1
var ispref bool
first:= 0
for true{
if name_dir[r].ispref{
ispref= true
q= name_dir[q].llink
}else{
ispref= false
q= -1
}
c:= web_strcmp(name,name_dir[r].name[1:])
switch c{
case equal:
if q==-1{
if ispref{
return first+len(name_dir[r].name[1:]),extension
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

/*:67*//*69:*/
//line common.w:1083

func mark_harmless(){
if history==spotless{
history= harmless_message
}
}

/*:69*//*70:*/
//line common.w:1091

func mark_error(){
history= error_message
}

/*:70*//*73:*/
//line common.w:1107


func err_print(s string){
var l int
if len(s)> 0&&s[0]=='!'{
fmt.Printf("\n%s",s)
}else{
fmt.Printf("%s",s)
}
if len(file)> 0&&file[0]!=nil{
/*74:*/
//line common.w:1132

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
for k:= 0;k<l;k++{
if buffer[k]=='\t'{
fmt.Print(" ")
}else{
fmt.Printf("%c",buffer[k])
}
}

fmt.Println()
fmt.Printf("%*c",l,' ')
}
fmt.Println(string(buffer[l:]))
if len(buffer)> 0&&buffer[len(buffer)-1]=='|'{
fmt.Print("|")
}
fmt.Print(" ")
}

/*:74*/
//line common.w:1117

}
os.Stdout.Sync()
mark_error()
}

/*:73*//*76:*/
//line common.w:1178

func wrap_up()int{
fmt.Print("\n")
if show_stats(){
print_stats()
}
/*77:*/
//line common.w:1191

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
}

/*:77*/
//line common.w:1184

if history> harmless_message{
return 1
}
return 0
}

/*:76*//*79:*/
//line common.w:1211

func fatal(s string,t string){
if len(s)!=0{
fmt.Print(s)
}
err_print(t)
history= fatal_message
os.Exit(wrap_up())
}

/*:79*//*80:*/
//line common.w:1223

func overflow(t string){
fmt.Printf("\n! Sorry, %s capacity exceeded",t)
fatal("","")
}


/*:80*//*81:*/
//line common.w:1235

func confusion(s string){
fatal("! This can't happen: ",s)

}

/*:81*//*83:*/
//line common.w:1251

func show_banner()bool{
return flags['b']
}

/*:83*//*84:*/
//line common.w:1257

func show_progress()bool{
return flags['p']
}

/*:84*//*85:*/
//line common.w:1263

func show_stats()bool{
return flags['s']
}

/*:85*//*86:*/
//line common.w:1269

func show_happiness()bool{
return flags['h']
}

/*:86*//*90:*/
//line common.w:1306

func scan_args(){
dot_pos:= -1
name_pos:= 0
found_web:= false
found_change:= false
found_out:= false

flag_change:= false

for i:= 1;i<len(os.Args);i++{
arg:= os.Args[i]
if(arg[0]=='-'||arg[0]=='+')&&len(arg)> 1{
/*94:*/
//line common.w:1403

{
if arg[0]=='-'{
flag_change= false
}else{
flag_change= true
}
for i:= 1;i<len(arg);i++{
flags[arg[i]]= flag_change
}
}

/*:94*/
//line common.w:1319

}else{
name_pos= 0
dot_pos= -1
for j:= 0;j<len(arg);j++{
if arg[j]=='.'{
dot_pos= j
}else if arg[j]=='/'{
dot_pos= -1
name_pos= j+1
}
}
if!found_web{
/*91:*/
//line common.w:1354

{
if dot_pos==-1{
file_name= append(file_name,fmt.Sprintf("%s.w",arg))
}else{
file_name= append(file_name,arg)
arg= arg[:dot_pos]
}
alt_file_name= fmt.Sprintf("%s.web",arg)
tex_file_name= fmt.Sprintf("%s.tex",arg[name_pos:])
idx_file_name= fmt.Sprintf("%s.idx",arg[name_pos:])
scn_file_name= fmt.Sprintf("%s.scn",arg[name_pos:])
go_file_name= fmt.Sprintf("%s.go",arg[name_pos:])
found_web= true
}

/*:91*/
//line common.w:1332

}else if!found_change{
/*92:*/
//line common.w:1370

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

/*:92*/
//line common.w:1334

}else if!found_out{
/*93:*/
//line common.w:1384

{
if dot_pos==-1{
tex_file_name= fmt.Sprintf("%s.tex",arg)
idx_file_name= fmt.Sprintf("%s.idx",arg)
scn_file_name= fmt.Sprintf("%s.scn",arg)
go_file_name= fmt.Sprintf("%s.go",arg)
}else{
tex_file_name= arg
go_file_name= arg
if flags['x']{
dot_pos= -1
idx_file_name= fmt.Sprintf("%s.idx",arg)
scn_file_name= fmt.Sprintf("%s.scn",arg)
}
}
found_out= true
}

/*:93*/
//line common.w:1336

}else{
/*411:*/
//line goweave.w:6619

{
fatal("! Usage: goweave [options] webfile[.w] [{changefile[.ch]|-} [outfile[.tex]]]\n","")

}

/*:411*/
//line common.w:1338

}
}
}
if!found_web{
/*411:*/
//line goweave.w:6619

{
fatal("! Usage: goweave [options] webfile[.w] [{changefile[.ch]|-} [outfile[.tex]]]\n","")

}

/*:411*/
//line common.w:1343

}
}

/*:90*//*97:*/
//line common.w:1431

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

/*:97*//*107:*/
//line goweave.w:219

func append_xref(c int32){
xmem= append(xmem,xref_info{})
xmem[len(xmem)-1].num= c
xmem[len(xmem)-1].xlink= 0
}

func is_tiny(p int32)bool{
return p<int32(len(name_dir))&&len(name_dir[p].name)==1
}


func unindexed(p int32)bool{
return p<res_wd_end&&name_dir[p].ilk>=custom
}

/*:107*//*108:*/
//line goweave.w:236

func new_xref(p int32){
if flags['x']==false{
return
}
if(unindexed(p)||is_tiny(p))&&xref_switch==0{
return
}
m:= section_count+xref_switch
xref_switch= 0
q:= name_dir[p].xref
if q>=0{
n:= xmem[q].num
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

/*:108*//*109:*/
//line goweave.w:272

func new_section_xref(p int32){
var r int32= 0
q:= name_dir[p].xref

if q>=0{
for q>=0&&q<int32(len(xmem))&&xmem[q].num> section_xref_switch{
r= q
q= xmem[q].xlink
}
}
if r> 0&&r<int32(len(xmem))&&xmem[r].num==section_count+section_xref_switch{
return
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

/*:109*//*110:*/
//line goweave.w:299

func set_file_flag(p int32){
q:= name_dir[p].xref
if xmem[q].num==file_flag{
return
}
append_xref(file_flag)
xmem[len(xmem)-1].xlink= q
name_dir[p].xref= int32(len(xmem)-1)
}

/*:110*//*111:*/
//line goweave.w:311

func names_match(
p int32,
id[]rune,
t int32)bool{
if len(name_dir[p].name)!=len(id){
return false
}
if name_dir[p].ilk!=t&&!(t==normal&&name_dir[p].ilk> typewriter){
return false
}
return compare_runes(id,name_dir[p].name)==0
}

func init_p(p int32,t int32){
name_dir[p].ilk= t
name_dir[p].xref= 0
}

func init_node(p int32){
name_dir[p].xref= 0
}

/*:111*//*121:*/
//line goweave.w:538

func skip_limbo(){
for{
if loc>=len(buffer)&&!get_line(){
return
}
for loc<len(buffer)&&buffer[loc]!='@'{
loc++
}
l:= loc
loc++
if l<len(buffer){
c:= new_section
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
/*153:*/
//line goweave.w:1225

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

/*:153*/
//line goweave.w:561

}
}
}
}

/*:121*//*122:*/
//line goweave.w:574


func skip_TeX()rune{
for{
if loc>=len(buffer)&&!get_line(){
return new_section
}
for loc<len(buffer)&&buffer[loc]!='@'&&buffer[loc]!='|'{
loc++
}
l:= loc
loc++
if l<len(buffer)&&buffer[l]=='|'{
return'|'
}
if loc<len(buffer){
l:= loc
loc++
return ccode[buffer[l]]
}
if l<len(buffer)&&buffer[l]=='@'{
return new_section
}
}
return 0
}

/*:122*//*127:*/
//line goweave.w:653


func get_next()rune{
for{
if loc>=len(buffer){

i:= len(scrap_info)-1
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
c:= buffer[loc]
loc++
nc:= ' '
if loc<len(buffer){
nc= buffer[loc]
}
if unicode.IsDigit(c)||(c=='.'&&unicode.IsDigit(nc)){
/*130:*/
//line goweave.w:735
{
id= nil
is_dec:= false
if loc<len(buffer)&&buffer[loc-1]=='0'{
if buffer[loc]=='x'||buffer[loc]=='X'{
id= append(id,'^')
loc++
for loc<len(buffer)&&xisxdigit(buffer[loc]){
id= append(id,buffer[loc])
loc++
}
}else if unicode.IsDigit(buffer[loc]){
id= append(id,'~')
for loc<len(buffer)&&unicode.IsDigit(buffer[loc]){
id= append(id,buffer[loc])
loc++
}
}else{
is_dec= true
}
}else{
is_dec= true
}
if is_dec{
if loc<len(buffer)&&buffer[loc-1]=='.'&&!unicode.IsDigit(buffer[loc]){
goto mistake
}
id= append(id,buffer[loc-1])
for loc<len(buffer)&&(unicode.IsDigit(buffer[loc])||buffer[loc]=='.'){
id= append(id,buffer[loc])
loc++
}
if loc<len(buffer)&&(buffer[loc]=='e'||buffer[loc]=='E'){
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

/*:130*/
//line goweave.w:688

}else if c=='\''||c=='"'||c=='`'{
/*131:*/
//line goweave.w:792
{
delim:= c
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

}
}
l:= loc
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

/*:131*/
//line goweave.w:690

}else if unicode.IsLetter(c)||
c=='_'&&(unicode.IsLetter(c)||unicode.IsDigit(c)){
/*129:*/
//line goweave.w:712
{
loc--
id_first:= loc
for loc<len(buffer)&&
(unicode.IsLetter(buffer[loc])||
unicode.IsDigit(buffer[loc])||
buffer[loc]=='_'){
loc++
}
id= buffer[id_first:loc]
return identifier
}

/*:129*/
//line goweave.w:693

}else if c=='@'{
/*132:*/
//line goweave.w:834
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
/*133:*/
//line goweave.w:863
{
section_text= section_text[0:0]
cur_section_char= nc
/*135:*/
//line goweave.w:885

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
/*136:*/
//line goweave.w:910

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
cc:= ignore
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
loc++
}

/*:136*/
//line goweave.w:899

loc++
if unicode.IsSpace(c){
c= ' '
if len(section_text)> 0&&section_text[len(section_text)-1]==' '{
section_text= section_text[:len(section_text)-1]
}
}
section_text= append(section_text,c)
}

/*:135*/
//line goweave.w:866

if len(section_text)> 3&&
compare_runes(section_text[len(section_text)-3:],[]rune("..."))==0{
cur_section= section_lookup(section_text[0:len(section_text)-3],
true)
}else{
cur_section= section_lookup(section_text,false)
}
xref_switch= 0
return section_name
}

/*:133*/
//line goweave.w:850

case verbatim:
/*139:*/
//line goweave.w:974
{
id_first:= loc
loc++
for loc+1<len(buffer)&&(buffer[loc]!='@'||buffer[loc+1]!='>'){
loc++
}
if loc>=len(buffer){
err_print("! Verbatim string didn't end")

}
id= buffer[id_first:loc]
loc+= 2
return verbatim
}

/*:139*/
//line goweave.w:852

case ord:
/*131:*/
//line goweave.w:792
{
delim:= c
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

}
}
l:= loc
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

/*:131*/
//line goweave.w:854

default:
return ccode[c]
}
}

/*:132*/
//line goweave.w:695

}else if unicode.IsSpace(c){
continue
}
mistake:
/*98:*/
//line common.w:1452

switch c{
case'/':
if nc=='*'{
l:= loc
loc++
if l<=len(buffer){
return begin_comment
}
}else if nc=='/'{
l:= loc
loc++
if l<=len(buffer){
return begin_short_comment
}
}
case'+':
if nc=='+'{
l:= loc
loc++
if l<=len(buffer){
return plus_plus
}
}
case'-':
if nc=='-'{
l:= loc
loc++
if l<=len(buffer){
return minus_minus
}
}
case'.':
if nc=='.'&&loc+1<len(buffer)&&buffer[loc+1]=='.'{
loc++
l:= loc
loc++
if l<=len(buffer){
return dot_dot_dot
}
}
case'=':
if nc=='='{
l:= loc
loc++
if l<=len(buffer){
return eq_eq
}
}
case'>':
if nc=='='{
l:= loc
loc++
if l<=len(buffer){
return gt_eq
}
}else if nc=='>'{
l:= loc
loc++
if l<=len(buffer){
return gt_gt
}
}
case'<':
if nc=='<'{
l:= loc
loc++
if l<=len(buffer){
return lt_lt
}
}else if nc=='-'{
l:= loc
loc++
if l<=len(buffer){
return direct
}
}else if nc=='='{
l:= loc
loc++
if l<=len(buffer){
return lt_eq
}
}
case'&':
if nc=='&'{
l:= loc
loc++
if l<=len(buffer){
return and_and
}
}else if nc=='^'{
l:= loc
loc++
if l<=len(buffer){
return and_not
}
}

case'|':
if nc=='|'{
l:= loc
loc++
if l<=len(buffer){
return or_or
}
}
case'!':
if nc=='='{
l:= loc
loc++
if l<=len(buffer){
return not_eq
}
}
case':':
if nc=='='{
l:= loc
loc++
if l<=len(buffer){
return col_eq
}
}
}
//line goweave.w:120

/*:98*/
//line goweave.w:700

return c
}
return 0
}

/*:127*//*138:*/
//line goweave.w:943

func skip_restricted(){
id_first:= loc
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
l:= loc
loc++
if buffer[l]!='>'{
err_print("! Control codes are forbidden in control text")

}
}
}

/*:138*//*142:*/
//line goweave.w:1006

func phase_one(){
phase= 1
reset_input()
section_count= 0
skip_limbo()
change_exists= false
for!input_has_ended{
/*143:*/
//line goweave.w:1022

{
section_count++
changed_section[section_count]= changing

if loc-1<len(buffer)&&buffer[loc-1]=='*'&&show_progress(){
fmt.Printf("*%d",section_count)
os.Stdout.Sync()
}
/*148:*/
//line goweave.w:1117

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
next_control= get_next()
if next_control>=xref_roman&&next_control<=xref_typewriter{
/*149:*/
//line goweave.w:1142

{
i:= 0
j:= 0
for i<len(id){
if id[i]=='@'{
i++
}
id[j]= id[i]
j++
i++
}
for j<i{
id[j]= ' '
j++
}
}

/*:149*/
//line goweave.w:1133

new_xref(id_lookup(id,next_control-identifier))
}
}
if next_control>=format_code{
break
}
}

/*:148*/
//line goweave.w:1031

/*151:*/
//line goweave.w:1179

for next_control<=format_code{
/*152:*/
//line goweave.w:1190
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

var r int32= 0
for q:= name_dir[lhs].xref;q>=0;q= xmem[q].xlink{
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

/*:152*/
//line goweave.w:1181

outer_xref()
}

/*:151*/
//line goweave.w:1032

/*154:*/
//line goweave.w:1245

if next_control<=section_name{
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

/*:154*/
//line goweave.w:1033

if changed_section[section_count]{
change_exists= true
}
}

/*:143*/
//line goweave.w:1014

}
changed_section[section_count]= change_exists

phase= 2
/*158:*/
//line goweave.w:1313
section_check(name_root)

/*:158*/
//line goweave.w:1019

}

/*:142*//*145:*/
//line goweave.w:1061


func Go_xref(spec_ctrl rune){
for next_control<format_code||next_control==spec_ctrl{
if next_control>=identifier&&next_control<=xref_typewriter{
if next_control> identifier{
/*149:*/
//line goweave.w:1142

{
i:= 0
j:= 0
for i<len(id){
if id[i]=='@'{
i++
}
id[j]= id[i]
j++
i++
}
for j<i{
id[j]= ' '
j++
}
}

/*:149*/
//line goweave.w:1067

}
p:= id_lookup(id,next_control-identifier)

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

/*:145*//*147:*/
//line goweave.w:1090


func outer_xref(){
for next_control<format_code{
if next_control!=begin_comment&&next_control!=begin_short_comment{
Go_xref(ignore)
}else{
is_long_comment:= (next_control==begin_comment)
bal,res:= copy_comment(is_long_comment,1,nil)
next_control= '|'
for bal> 0{
Go_xref(section_name)
if next_control=='|'{
bal,res= copy_comment(is_long_comment,bal,res)
}else{
bal= 0
}
}
}
}
}

/*:147*//*157:*/
//line goweave.w:1280


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

/*:157*//*160:*/
//line goweave.w:1338


func flush_buffer(b int32,per_cent bool,carryover bool){
j:= b
if!per_cent{
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
jj:= j
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

/*:160*//*161:*/
//line goweave.w:1378


func finish_line(){
if out_ptr> 0{
flush_buffer(out_ptr,false,false)
}else{
for _,v:= range buffer{
if!unicode.IsSpace(v){
return
}
}
flush_buffer(0,false,false)
}
}

/*:161*//*164:*/
//line goweave.w:1414

func out(c rune){
if out_ptr>=out_buf_end{
break_out()
}
out_ptr++
out_buf[out_ptr]= c
}

/*:164*//*165:*/
//line goweave.w:1424


func out_str(s string){
for _,v:= range s{
out(v)
}
}

/*:165*//*168:*/
//line goweave.w:1444


func break_out(){
k:= out_ptr
for{
if k==0{
/*169:*/
//line goweave.w:1470

{
fmt.Printf("\n! Line had to be broken (output l. %d):\n",out_line)

fmt.Fprint(os.Stdout,string(out_buf[1:out_ptr]))
fmt.Println()
mark_harmless()
flush_buffer(out_ptr-1,true,true)
return
}

/*:169*/
//line goweave.w:1450

}
if out_buf[k]==' '{
flush_buffer(k,false,true)
return
}
kk:= k
k--
if out_buf[kk]=='\\'&&out_buf[k]!='\\'{
flush_buffer(k,true,true)
return
}
}
}

/*:168*//*170:*/
//line goweave.w:1486

func section_str(n int32)string{
s:= fmt.Sprintf("%d",n)
if changed_section[n]{
s= "\\*"

}
return s
}

/*:170*//*171:*/
//line goweave.w:1499

func out_name(p int32,quote_xalpha bool){
out('{')
for _,v:= range name_dir[p].name{
if v=='_'&&quote_xalpha{
out('\\')
}


out(v)
}
out('}')
}

/*:171*//*172:*/
//line goweave.w:1526

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
l:= loc
loc++
if l<len(buffer){
c:= ' '
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
get_line()
}

default:
err_print("! Double @ should be used in limbo")

out('@')
}
}
}
}

/*:172*//*175:*/
//line goweave.w:1580

func copy_TeX()rune{
for{
if loc>=len(buffer){
finish_line()
if!get_line(){
return new_section
}
}
c:= buffer[loc]
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
l:= loc
loc++
return ccode[buffer[l]]
}
}
return 0
}

/*:175*//*177:*/
//line goweave.w:1625


func copy_comment(
is_long_comment bool,
bal int,
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
c:= buffer[loc]
loc++
if c=='|'{
return bal,tok_mem
}
if is_long_comment{
/*178:*/
//line goweave.w:1681

if c=='*'&&loc<len(buffer)&&buffer[loc]=='/'{
loc++
if bal> 1{
err_print("! Missing } in comment")

}
goto done
}

/*:178*/
//line goweave.w:1654

}
if phase==2{
if c> 0177{
tok_mem= append(tok_mem,quoted_char)
}
tok_mem= append(tok_mem,c)
}
/*179:*/
//line goweave.w:1691

if c=='@'{
l:= loc
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

/*:179*/
//line goweave.w:1662

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
/*180:*/
//line goweave.w:1714

if phase==2{
for bal--;bal>=0;bal--{
tok_mem= append(tok_mem,'}')
}
}
return 0,tok_mem

/*:180*/
//line goweave.w:1678

}

/*:177*//*185:*/
//line goweave.w:2071


func print_cat(c int32){
fmt.Printf("%s(%v)",cat_name[c],c)
}

/*:185*//*196:*/
//line goweave.w:2303

func call(fs[]reducing){
for i:= len(fs)-1;i>=0;i--{
fs[i]()
}
}

/*:196*//*197:*/
//line goweave.w:2311

func one(ss[]scrap,c rune)([]scrap,reducing,bool){
m:= "found"
if(tracing&4)==4{
fmt.Printf("%*cLooking for %q...\n",shift,' ',cat_name[c])
shift+= 5
defer func(){shift-= 5;fmt.Printf("%*c%q is %s\n",shift,' ',cat_name[c],m)}()
f:= cat_name[c]
fmt.Printf("%*c",shift,' ')
/*323:*/
//line goweave.w:4841

{
if(tracing&2)==2{
fmt.Printf("%s:",f)
for i,v:= range scrap_info{
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

/*:323*/
//line goweave.w:2320

}
if len(ss)==0{
return ss,empty,false
}
if ss[0].cat==c{
return ss[1:],empty,true
}
switch c{
case ConstDecl:/*214:*/
//line goweave.w:2681

if s,f1,ok:= one(ss,const_token);ok{
if s,f2,ok:= one(s,ConstSpec);ok{
return s,func(){
f2()
f1()
reduce(ss,2,ConstDecl,0,break_space,1,force)
},true
}else if s,f2,ok:= one(s,lpar);ok{
tok_mem:= append([]interface{}{},0,1)
s,f3,t,ok:= optional(s,2,pair{ConstSpec,true})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
}
if s,f4,ok:= one(s,rpar);ok{
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

/*:214*/
//line goweave.w:2329

case TypeDecl:/*216:*/
//line goweave.w:2730

if s,f1,ok:= one(ss,type_token);ok{
if s,f2,ok:= one(s,TypeSpec);ok{
return s,func(){
f2()
f1()
reduce(ss,2,TypeDecl,0,break_space,1,force)
},true
}else if s,f2,ok:= one(s,lpar);ok{
tok_mem:= append([]interface{}{},0,1)
s,f3,t,ok:= optional(s,2,pair{cat:TypeSpec,mand:true})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
}
if s,f4,ok:= one(s,rpar);ok{
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

/*:216*/
//line goweave.w:2330

case VarDecl:/*218:*/
//line goweave.w:2783

if s,f1,ok:= one(ss,var_token);ok{
if s,f2,ok:= one(s,VarSpec);ok{
return s,func(){
f2()
f1()
reduce(ss,2,VarDecl,0,break_space,1)
},true
}else if s,f2,ok:= one(s,lpar);ok{
tok_mem:= append([]interface{}{},0,1)
s,f3,t,ok:= optional(s,2,pair{cat:VarSpec,mand:true})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
}
if s,f4,ok:= one(s,rpar);ok{
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

/*:218*/
//line goweave.w:2331

case FunctionDecl:/*222:*/
//line goweave.w:2890

if s,f1,ok:= sequence(ss,func_token,identifier,Signature);ok{
if s,f2,ok:= sequence(s,Block,semi);ok{
return s,func(){
call(f2)
call(f1)
make_underlined(1)
reduce(ss,5,FunctionDecl,0,break_space,1,2,3,4,big_force)
},true
}else if s,f2,ok:= one(s,semi);ok{
return s,func(){
f2()
call(f1)
make_underlined(1)
reduce(ss,4,FunctionDecl,0,break_space,1,2,3,big_force)
},true
}
}

/*:222*/
//line goweave.w:2332

case MethodDecl:/*224:*/
//line goweave.w:2924

if s,f1,ok:= sequence(ss,func_token,Receiver,identifier,Signature);ok{
if s,f2,ok:= one(s,Block);ok{
return s,func(){
f2()
call(f1)
make_underlined(2)
reduce(ss,5,MethodDecl,0,break_space,1,break_space,2,3,4)
},true
}else{
return s,func(){
call(f1)
make_underlined(2)
reduce(ss,4,MethodDecl,0,break_space,1,break_space,2,3)
},true
}
}

/*:224*/
//line goweave.w:2333

case Receiver:/*226:*/
//line goweave.w:2958

if s,f1,ok:= one(ss,lpar);ok{
if s,f2,ok:= one(s,identifier);ok{
if s,f3,ok:= sequence(s,asterisk,identifier,rpar);ok{
return s,func(){
call(f3)
f2()
f1()
reduce(ss,5,Receiver,0,1,2,3,4)
},true
}else if s,f,ok:= sequence(s,identifier,rpar);ok{
return s,func(){
call(f)
reduce(ss,4,Receiver,0,1,2,3)
},true
}else if s,f,ok:= one(s,rpar);ok{
return s,func(){
f()
reduce(ss,3,Receiver,0,1,2)
},true
}
}else if s,f,ok:= sequence(s,asterisk,identifier,rpar);ok{
return s,func(){
call(f)
reduce(ss,4,Receiver,0,1,2,3)
},true
}
}

/*:226*/
//line goweave.w:2334

case ConstSpec:/*227:*/
//line goweave.w:2987

if s,f1,ok:= one(ss,IdentifierList);ok{
if s,f2,ok:= sequence(s,Type,eq,ExpressionList);ok{
if s,f3,ok:= one(s,semi);ok{
return s,func(){
f3()
call(f2)
f1()
reduce(ss,5,ConstSpec,0,break_space,1,break_space,2,break_space,3,4,force)
},true
}else if _,_,ok:= any(s,rpar,rbrace);ok{
return s,func(){
call(f2)
f1()
reduce(ss,4,ConstSpec,0,break_space,1,break_space,2,break_space,3,force)
},true
}
}else if s,f2,ok:= sequence(s,eq,ExpressionList);ok{
if s,f3,ok:= one(s,semi);ok{
return s,func(){
f3()
call(f2)
f1()
reduce(ss,4,ConstSpec,0,break_space,1,break_space,2,3,force)
},true
}else if _,_,ok:= any(s,rpar,rbrace);ok{
return s,func(){
call(f2)
f1()
reduce(ss,3,ConstSpec,0,break_space,1,break_space,2,force)
},true
}
}
}else if s,f,ok:= one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,ConstSpec,0,force)
},true
}

/*:227*/
//line goweave.w:2335

case TypeSpec:/*228:*/
//line goweave.w:3027

if s,f1,ok:= sequence(ss,identifier,Type);ok{
if s,f2,ok:= one(s,semi);ok{
return s,func(){
f2()
call(f1)
reduce(ss,3,TypeSpec,0,break_space,1,2,force)
},true
}else if _,_,ok:= any(s,rpar,rbrace);ok{
return s,func(){
call(f1)
reduce(ss,2,TypeSpec,0,break_space,1,force)
},true
}
}else if s,f,ok:= one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,TypeSpec,0,force)
},true
}

/*:228*/
//line goweave.w:2336

case VarSpec:/*229:*/
//line goweave.w:3048

if s,f1,ok:= one(ss,IdentifierList);ok{
if s,f2,ok:= one(s,Type);ok{
if s,f3,ok:= sequence(s,eq,ExpressionList);ok{
if s,f4,ok:= one(s,semi);ok{
return s,func(){
f4()
call(f3)
f2()
f1()
reduce(ss,5,VarSpec,0,break_space,1,2,3,4,force)
},true
}else if _,_,ok:= any(s,rpar,rbrace);ok{
return s,func(){
call(f3)
f2()
f1()
reduce(ss,4,VarSpec,0,break_space,1,2,3,force)
},true
}
}else if s,f3,ok:= one(s,semi);ok{
return s,func(){
f3()
f2()
f1()
reduce(ss,3,VarSpec,0,break_space,1,2,force)
},true
}else if _,_,ok:= any(s,rpar,rbrace);ok{
return s,func(){
f2()
f1()
reduce(ss,2,VarSpec,0,break_space,1,force)
},true
}
}else if s,f2,ok:= sequence(s,eq,ExpressionList);ok{
if s,f3,ok:= one(s,semi);ok{
return s,func(){
f3()
call(f2)
f1()
reduce(ss,4,VarSpec,0,1,2,3,force)
},true
}else if _,_,ok:= any(s,rpar,rbrace);ok{
return s,func(){
call(f2)
f1()
reduce(ss,3,VarSpec,0,1,2,force)
},true
}
}
}else if s,f,ok:= one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,VarSpec,0,force)
},true
}

/*:229*/
//line goweave.w:2337

case ImportSpec:/*230:*/
//line goweave.w:3105

if s,f1,ok:= sequence(ss,identifier,str);ok{
if s,f2,ok:= one(s,semi);ok{
return s,func(){
f2()
call(f1)
make_reserved(ss[0],PackageName)
reduce(ss,3,ImportSpec,0,break_space,1,2,force)
},true
}else if _,_,ok:= any(s,rpar,rbrace);ok{
return s,func(){
call(f1)
make_reserved(ss[0],PackageName)
reduce(ss,2,ImportSpec,0,break_space,1,force)
},true
}
}else if s,f1,ok:= sequence(ss,dot,str);ok{
if s,f2,ok:= one(s,semi);ok{
return s,func(){
f2()
call(f1)
reduce(ss,3,ImportSpec,0,break_space,1,2,force)
},true
}else if _,_,ok:= any(s,rpar,rbrace);ok{
return s,func(){
call(f1)
reduce(ss,2,ImportSpec,0,break_space,1,force)
},true
}
}else if s,f1,ok:= one(ss,str);ok{
if s,f2,ok:= one(s,semi);ok{
return s,func(){
f2()
f1()
reduce(ss,2,ImportSpec,0,1,force)
},true
}else if _,_,ok:= any(s,rpar,rbrace);ok{
return s,func(){
f1()
reduce(ss,1,ImportSpec,0,force)
},true
}
}else if s,f,ok:= one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,ImportSpec,0,force)
},true
}

/*:230*/
//line goweave.w:2338

case FieldDecl:/*231:*/
//line goweave.w:3154

if s,f1,ok:= sequence(ss,IdentifierList,Type);ok{
tok_mem:= append([]interface{}{},0,break_space,1)
c:= 2
s,f2,ok:= one(s,str)
if ok{
tok_mem= append(tok_mem,break_space,2)
c++
}
if s,f3,ok:= one(s,semi);ok{
tok_mem= append(tok_mem,c,force)
c++
return s,func(){
f3()
f2()
call(f1)
reduce(ss,c,FieldDecl,tok_mem...)
},true
}else if _,_,ok:= any(s,rpar,rbrace);ok{
tok_mem= append(tok_mem,force)
return s,func(){
f2()
call(f1)
reduce(ss,c,FieldDecl,tok_mem...)
},true
}
}else if s,f1,ok:= one(ss,AnonymousField);ok{
tok_mem:= append([]interface{}{},0)
c:= 1
s,f2,ok:= one(s,str)
if ok{
tok_mem= append(tok_mem,break_space,1)
c++
}
tok_mem= append(tok_mem,force)
if s,f3,ok:= one(s,semi);ok{
c++
return s,func(){
f3()
f2()
f1()
reduce(ss,c,FieldDecl,tok_mem...)
},true
}else if _,_,ok:= any(s,rpar,rbrace);ok{
return s,func(){
f2()
f1()
reduce(ss,c,FieldDecl,tok_mem...)
},true
}
}else if s,f,ok:= one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,FieldDecl,0,force)
},true
}

/*:231*/
//line goweave.w:2339

case AnonymousField:/*232:*/
//line goweave.w:3211

if s,f,ok:= sequence(ss,asterisk,Type);ok{
return s,func(){
call(f)
reduce(ss,2,AnonymousField,0,1)
},true
}else if s,f,ok:= one(ss,Type);ok{
return s,func(){
f()
reduce(ss,1,AnonymousField,0)
},true
}

/*:232*/
//line goweave.w:2340

case Type:/*233:*/
//line goweave.w:3224

if s,f,ok:= any(ss,
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

/*:233*/
//line goweave.w:2341

case ArrayType:/*234:*/
//line goweave.w:3241

if s,f,ok:= sequence(ss,lbracket,Expression,rbracket,Type);ok{
return s,func(){
call(f)
reduce(ss,4,ArrayType,0,1,2,3)
},true
}

/*:234*/
//line goweave.w:2342

case StructType:/*235:*/
//line goweave.w:3249

if s,f1,ok:= sequence(ss,struct_token,lbrace);ok{
tok_mem:= append([]interface{}{},0,1)
s,f2,t,ok:= optional(s,2,pair{cat:FieldDecl,mand:true})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
}
if s,f3,ok:= one(s,rbrace);ok{
tok_mem= append(tok_mem,2+len(f2))
return s,func(){
f3()
call(f2)
call(f1)
reduce(ss,3+len(f2),StructType,tok_mem...)
},true
}
}

/*:235*/
//line goweave.w:2343

case PointerType:/*237:*/
//line goweave.w:3299

if s,f,ok:= sequence(ss,asterisk,Type);ok{
return s,func(){
call(f)
reduce(ss,2,PointerType,0,1)
},true
}

/*:237*/
//line goweave.w:2344

case Signature:/*238:*/
//line goweave.w:3307

if s,f1,ok:= one(ss,Parameters);ok{
if s,f2,ok:= any(s,Type,Parameters);ok{
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
}else if s,f,ok:= one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,Signature,0,force)
},true
}

/*:238*/
//line goweave.w:2345

case Parameters:/*239:*/
//line goweave.w:3328

if s,f1,ok:= one(ss,lpar);ok{
tok_mem:= append([]interface{}{},0)
s,f2,t,ok:= optional(s,1,pair{cat:ParameterList,mand:true},pair{cat:comma,mand:false})
if ok{
tok_mem= append(tok_mem,t)
}
if s,f3,ok:= one(s,rpar);ok{
tok_mem= append(tok_mem,1+len(f2))
return s,func(){
f3()
call(f2)
f1()
reduce(ss,2+len(f2),Parameters,tok_mem...)
},true
}
}else if s,f,ok:= one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,Signature,0,force)
},true
}

/*:239*/
//line goweave.w:2346

case ParameterList:/*240:*/
//line goweave.w:3351

if s,f1,ok:= one(ss,ParameterDecl);ok{
tok_mem:= append([]interface{}{},0)
s,f2,t,ok:= optional(s,1,pair{cat:comma,mand:true},pair{cat:ParameterDecl,mand:true})
if ok{
tok_mem= append(tok_mem,t)
}
return s,func(){
call(f2)
f1()
reduce(ss,1+len(f2),ParameterList,tok_mem...)
},true
}

/*:240*/
//line goweave.w:2347

case ParameterDecl:/*241:*/
//line goweave.w:3365

if s,f,ok:= sequence(ss,IdentifierList,dot_dot_dot,Type);ok{
return s,func(){
call(f)
reduce(ss,3,ParameterDecl,0,break_space,1,2)
},true
}else if s,f,ok:= sequence(ss,IdentifierList,Type);ok{
return s,func(){
call(f)
reduce(ss,2,ParameterDecl,0,break_space,1)
},true
}else if s,f,ok:= sequence(ss,dot_dot_dot,Type);ok{
return s,func(){
call(f)
reduce(ss,2,ParameterDecl,0,1)
},true
}else if s,f,ok:= one(ss,Type);ok{
return s,func(){
f()
reduce(ss,1,ParameterDecl,0)
},true
}

/*:241*/
//line goweave.w:2348

case InterfaceType:/*242:*/
//line goweave.w:3388

if s,f1,ok:= sequence(ss,interface_token,lbrace);ok{
tok_mem:= append([]interface{}{},0,1)
s,f2,t,ok:= optional(s,2,pair{cat:MethodSpec,mand:true})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
}
if s,f3,ok:= one(s,rbrace);ok{
tok_mem= append(tok_mem,2+len(f2))
return s,func(){
f3()
call(f2)
call(f1)
reduce(ss,3+len(f2),InterfaceType,tok_mem...)
},true
}
}

/*:242*/
//line goweave.w:2349

case MethodSpec:/*243:*/
//line goweave.w:3406

if s,f1,ok:= sequence(ss,identifier,Signature);ok{
if s,f2,ok:= one(s,semi);ok{
return s,func(){
f2()
call(f1)
reduce(ss,3,MethodSpec,0,1,2,force)
},true
}else if _,_,ok:= any(s,rpar,rbrace);ok{
return s,func(){
call(f1)
reduce(ss,2,MethodSpec,0,1,force)
},true
}
}else if s,f1,ok:= sequence(ss,Type);ok{
if s,f2,ok:= one(s,semi);ok{
return s,func(){
f2()
call(f1)
reduce(ss,2,MethodSpec,0,1,force)
},true
}else if _,_,ok:= any(s,rpar,rbrace);ok{
return s,func(){
call(f1)
reduce(ss,1,MethodSpec,0,force)
},true
}
}else if s,f,ok:= one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,MethodSpec,0,force)
},true
}

/*:243*/
//line goweave.w:2350

case SliceType:/*244:*/
//line goweave.w:3440

if s,f,ok:= sequence(ss,lbracket,rbracket,Type);ok{
return s,func(){
call(f)
reduce(ss,3,SliceType,0,1,2)
},true
}

/*:244*/
//line goweave.w:2351

case MapType:/*245:*/
//line goweave.w:3448

if s,f,ok:= sequence(ss,map_token,lbracket,Type,rbracket,Type);ok{
return s,func(){
call(f)
reduce(ss,5,MapType,0,1,2,3,4)
},true
}

/*:245*/
//line goweave.w:2352

case ChannelType:/*246:*/
//line goweave.w:3456

if s,f,ok:= sequence(ss,direct,chan_token,Type);ok{
return s,func(){
call(f)
reduce(ss,3,ChannelType,0,1,break_space,2)
},true
}else if s,f1,ok:= one(ss,chan_token);ok{
if s,f2,ok:= sequence(s,direct,Type);ok{
return s,func(){
call(f2)
f1()
reduce(ss,3,ChannelType,0,1,2)
},true
}else if s,f2,ok:= one(s,Type);ok{
return s,func(){
f2()
f1()
reduce(ss,2,ChannelType,0,break_space,1)
},true
}
}

/*:246*/
//line goweave.w:2353

case IdentifierList:/*247:*/
//line goweave.w:3478

if s,f1,ok:= one(ss,identifier);ok{
tok_mem:= append([]interface{}{},0)
s,f2,t,ok:= optional(s,1,pair{cat:comma,mand:true},pair{cat:identifier,mand:true})
if ok{
tok_mem= append(tok_mem,t)
}
return s,func(){
call(f2)
f1()
reduce(ss,1+len(f2),IdentifierList,tok_mem...)
},true
}

/*:247*/
//line goweave.w:2354

case ExpressionList:/*248:*/
//line goweave.w:3492

if s,f1,ok:= one(ss,Expression);ok{
tok_mem:= append([]interface{}{},0)
s,f2,t,ok:= optional(s,1,pair{cat:comma,mand:true},pair{cat:Expression,mand:true})
if ok{
tok_mem= append(tok_mem,t)
}
return s,func(){
call(f2)
f1()
reduce(ss,1+len(f2),ExpressionList,tok_mem...)
},true
}

/*:248*/
//line goweave.w:2355

case Expression:/*249:*/
//line goweave.w:3506

if s,f1,ok:= one(ss,UnaryExpr);ok{
tok_mem:= append([]interface{}{},0)
s,f2,t,ok:= optional(s,1,pair{binary_op,true},pair{UnaryExpr,true});
if ok{
tok_mem= append(tok_mem,t)
}
return s,func(){
call(f2)
f1()
reduce(ss,1+len(f2),Expression,tok_mem...)
},true
}

/*:249*/
//line goweave.w:2356

case UnaryExpr:/*250:*/
//line goweave.w:3520

if s,f,ok:= one(ss,PrimaryExpr);ok{
return s,func(){
f()
reduce(ss,1,UnaryExpr,0)
},true
}else if s,f,ok:= sequence(ss,unary_op,UnaryExpr);ok{
return s,func(){
call(f)
reduce(ss,2,UnaryExpr,0,1)
},true
}

/*:250*/
//line goweave.w:2357

case binary_op:/*251:*/
//line goweave.w:3533

if s,f,ok:= any(ss,rel_op,add_op,mul_op,asterisk);ok{
return s,func(){
f()
reduce(ss,1,binary_op,0)
},true
}

/*:251*/
//line goweave.w:2358

case PrimaryExpr:/*252:*/
//line goweave.w:3541

if s,f1,ok:= any(ss,BuiltinCall,Conversion,Operand);ok{
tok_mem:= append([]interface{}{},0)
s,f2,t,ok:= optional(s,1,pair{cat:Selector,mand:false},
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

/*:252*/
//line goweave.w:2359

case Operand:/*253:*/
//line goweave.w:3560

if s,f,ok:= any(ss,
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
}else if s,f,ok:= sequence(ss,lpar,Expression,rpar);ok{
return s,func(){
call(f)
reduce(ss,3,Operand,0,1,2)
},true
}

/*:253*/
//line goweave.w:2360

case CompositeLit:/*254:*/
//line goweave.w:3579

if s,f,ok:= sequence(ss,LiteralType,LiteralValue);ok{
return s,func(){
call(f)
reduce(ss,2,CompositeLit,0,1)
},true
}

/*:254*/
//line goweave.w:2361

case LiteralType:/*255:*/
//line goweave.w:3587

if s,f,ok:= one(ss,Type);ok{
return s,func(){
f()
reduce(ss,1,LiteralType,0)
},true
}else if s,f,ok:= sequence(ss,lbracket,dot_dot_dot,rbracket,Type);ok{
return s,func(){
call(f)
reduce(ss,4,LiteralType,0,1,2,3)
},true
}

/*:255*/
//line goweave.w:2362

case LiteralValue:/*256:*/
//line goweave.w:3600

if s,f1,ok:= one(ss,lbrace);ok{
tok_mem:= append([]interface{}{},0)
s,f2,t,ok:= optional(s,1,pair{cat:ElementList,mand:true},pair{cat:comma,mand:true})
if ok{
tok_mem= append(tok_mem,t)
}
if s,f3,ok:= one(s,rbrace);ok{
tok_mem= append(tok_mem,1+len(f2))
return s,func(){
f3()
call(f2)
f1()
reduce(ss,2+len(f2),LiteralValue,tok_mem...)
},true
}
}

/*:256*/
//line goweave.w:2363

case ElementList:/*257:*/
//line goweave.w:3618

if s,f1,ok:= one(ss,Element);ok{
tok_mem:= append([]interface{}{},0)
s,f2,t,ok:= optional(s,1,pair{cat:comma,mand:true},pair{cat:Element,mand:true})
if ok{
tok_mem= append(tok_mem,t)
}
return s,func(){
call(f2)
f1()
reduce(ss,1+len(f2),ElementList,tok_mem...)
},true
}

/*:257*/
//line goweave.w:2364

case Element:/*258:*/
//line goweave.w:3632

s,f1,ok:= any(ss,identifier,Expression)
f2:= empty
if ok{
s,f2,ok= one(s,colon)
}
if ok{
if s,f3,ok:= one(s,Expression);ok{
return s,func(){
f3()
f2()
f1()
reduce(ss,3,Element,0,1,break_space,2)
},true
}else if s,f3,ok:= one(s,LiteralValue);ok{
return s,func(){
f3()
f2()
f1()
reduce(ss,3,Element,0,1,break_space,2)
},true
}
}else if s,f3,ok:= any(s,Expression,LiteralValue);ok{
return s,func(){
f3()
f2()
f1()
reduce(ss,1,Element,0)
},true
}

/*:258*/
//line goweave.w:2365

case FunctionLit:/*259:*/
//line goweave.w:3663

if s,f,ok:= sequence(ss,FunctionType,Block);ok{
return s,func(){
call(f)
reduce(ss,2,FunctionLit,0,1)
},true
}

/*:259*/
//line goweave.w:2366

case FunctionType:/*260:*/
//line goweave.w:3671

if s,f,ok:= sequence(ss,func_token,Signature);ok{
return s,func(){
call(f)
reduce(ss,2,FunctionType,0,1)
},true
}

/*:260*/
//line goweave.w:2367

case Block:/*261:*/
//line goweave.w:3679

if s,f1,ok:= one(ss,lbrace);ok{
tok_mem:= append([]interface{}{},0)
s,f2,t,ok:= optional(s,1,pair{cat:Statement,mand:true})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
}
if s,f3,ok:= one(s,rbrace);ok{
tok_mem= append(tok_mem,1+len(f2))
return s,func(){
f3()
call(f2)
f1()
reduce(ss,2+len(f2),Block,tok_mem...)
},true
}
}

/*:261*/
//line goweave.w:2368

case Statement:/*263:*/
//line goweave.w:3706

if s,f,ok:= any(ss,
ImportDecl,
ConstDecl,
VarDecl,
TypeDecl,
LabeledStmt);ok{
return s,func(){
f()
reduce(ss,1,Statement,0)
},true
}else if s,f1,ok:= any(ss,
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
if s,f2,ok:= one(s,semi);ok{
return s,func(){
f2()
f1()
reduce(ss,2,Statement,0,1,force)
},true
}else if _,_,ok:= any(s,rpar,rbrace);ok{
return s,func(){
f1()
reduce(ss,1,Statement,0,force)
},true
}
}else if s,f,ok:= one(ss,semi);ok{
return s,func(){
f()
reduce(ss,1,Statement,0,force)
},true
}else if s,f,ok:= one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,Statement,0,force)
},true
}

/*:263*/
//line goweave.w:2369

case LabeledStmt:/*264:*/
//line goweave.w:3756

if s,f,ok:= sequence(ss,identifier,colon,Statement);ok{
return s,func(){
call(f)
reduce(ss,3,LabeledStmt,0,1,force,2)
},true
}


/*:264*/
//line goweave.w:2370

case SimpleStmt:/*266:*/
//line goweave.w:3772

if s,f,ok:= any(ss,
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

/*:266*/
//line goweave.w:2371

case GoStmt:/*267:*/
//line goweave.w:3785

if s,f,ok:= sequence(ss,go_token,Expression);ok{
return s,func(){
call(f)
reduce(ss,2,GoStmt,0,break_space,1)
},true
}

/*:267*/
//line goweave.w:2372

case ReturnStmt:/*269:*/
//line goweave.w:3804

if s,f,ok:= sequence(ss,return_token,ExpressionList);ok{
return s,func(){
call(f)
reduce(ss,2,ReturnStmt,0,break_space,1)
},true
}else if s,f,ok:= one(ss,return_token);ok{
return s,func(){
f()
reduce(ss,1,ReturnStmt,0)
},true
}

/*:269*/
//line goweave.w:2373

case BreakStmt:/*271:*/
//line goweave.w:3831

if s,f1,ok:= one(ss,break_token);ok{
if s,f2,ok:= one(s,identifier);ok{
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

/*:271*/
//line goweave.w:2374

case ContinueStmt:/*273:*/
//line goweave.w:3869

if s,f,ok:= sequence(ss,continue_token,identifier);ok{
return s,func(){
call(f)
reduce(ss,2,ContinueStmt,0,break_space,1)
},true
}else if s,f,ok:= one(ss,continue_token);ok{
return s,func(){
f()
reduce(ss,1,ContinueStmt,0)
},true
}

/*:273*/
//line goweave.w:2375

case GotoStmt:/*275:*/
//line goweave.w:3903

if s,f,ok:= sequence(ss,goto_token,identifier);ok{
return s,func(){
call(f)
reduce(ss,2,GotoStmt,0,break_space,1)
},true
}

/*:275*/
//line goweave.w:2376

case IfStmt:/*277:*/
//line goweave.w:3918

if s,f1,ok:= one(ss,if_token);ok{
tok_mem:= append([]interface{}{},0)
c:= 1
var f2,f3[]reducing
f4,f5:= empty,empty
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


/*:277*/
//line goweave.w:2377

case ExprSwitchStmt:/*279:*/
//line goweave.w:3977

if s,f1,ok:= one(ss,switch_token);ok{
tok_mem:= append([]interface{}{},0)
c:= 1
var f2[]reducing
f3,f4:= empty,empty
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
s,f5,t,ok:= optional(s,c,pair{cat:ExprCaseClause,mand:false})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
c+= len(f5)
}
if s,f6,ok:= one(s,rbrace);ok{
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


/*:279*/
//line goweave.w:2378

case ExprCaseClause:/*280:*/
//line goweave.w:4021

if s,f1,ok:= sequence(ss,case_token,ExpressionList,colon);ok{
tok_mem:= append([]interface{}{},0,break_space,1,2)
s,f2,t,ok:= optional(s,3,pair{cat:Statement,mand:true})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
}
return s,func(){
call(f2)
call(f1)
reduce(ss,3+len(f2),ExprCaseClause,tok_mem...)
},true
}else if s,f1,ok:= sequence(ss,default_token,colon);ok{
tok_mem:= append([]interface{}{},0,1,force)
s,f2,t,ok:= optional(s,2,pair{cat:Statement,mand:true})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
}
return s,func(){
call(f2)
call(f1)
reduce(ss,2+len(f2),ExprCaseClause,tok_mem...)
},true
}else if s,f,ok:= one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,ExprCaseClause,0,force)
},true
}

/*:280*/
//line goweave.w:2379

case TypeSwitchStmt:/*281:*/
//line goweave.w:4051

if s,f1,ok:= one(ss,switch_token);ok{
tok_mem:= append([]interface{}{},0)
c:= 1
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
if s,f3,ok:= sequence(s,TypeSwitchGuard,lbrace);ok{
tok_mem= append(tok_mem,break_space,c,break_space,c+1)
c+= 2
s,f4,t,ok:= optional(s,c,pair{cat:TypeCaseClause,mand:true})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
c+= len(f4)
}
if s,f5,ok:= one(s,rbrace);ok{
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

/*:281*/
//line goweave.w:2380

case TypeSwitchGuard:/*282:*/
//line goweave.w:4088

if s,f,ok:= sequence(ss,identifier,col_eq,PrimaryExpr,dot,lpar,type_token,rpar);ok{
return s,func(){
call(f)
reduce(ss,7,TypeSwitchGuard,0,1,2,3,4,5,6)
},true
}else if s,f,ok:= sequence(ss,PrimaryExpr,dot,lpar,type_token,rpar);ok{
return s,func(){
call(f)
reduce(ss,5,TypeSwitchGuard,0,1,2,3,4)
},true
}

/*:282*/
//line goweave.w:2381

case TypeCaseClause:/*283:*/
//line goweave.w:4101

if s,f1,ok:= sequence(ss,TypeSwitchCase,colon);ok{
tok_mem:= append([]interface{}{},0,1,force)
s,f2,t,ok:= optional(s,2,pair{cat:Statement,mand:true})
if ok{
tok_mem= append(tok_mem,indent,t,outdent)
}
return s,func(){
call(f2)
call(f1)
reduce(ss,2+len(f2),TypeCaseClause,tok_mem...)
},true
}else if s,f,ok:= one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,TypeCaseClause,0,force)
},true
}

/*:283*/
//line goweave.w:2382

case TypeSwitchCase:/*284:*/
//line goweave.w:4120

if s,f1,ok:= sequence(ss,case_token);ok{
tok_mem:= append([]interface{}{},0)
if s,f2,ok:= any(s,Type,constant);ok{
tok_mem= append(tok_mem,break_space,1)
s,f3,t,ok:= optional(s,2,pair{cat:comma,mand:true},pair{cat:Type,mand:true})
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
}else if s,f,ok:= one(ss,default_token);ok{
return s,func(){
f()
reduce(ss,1,TypeSwitchCase,0)
},true
}else if s,f,ok:= one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,TypeSwitchCase,0,force)
},true
}

/*:284*/
//line goweave.w:2383

case SelectStmt:/*286:*/
//line goweave.w:4188

if s,f1,ok:= sequence(ss,select_token,lbrace);ok{
tok_mem:= append([]interface{}{},0,1)
s,f2,t,ok:= optional(s,2,pair{cat:CommClause,mand:false})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
}
if s,f3,ok:= one(s,rbrace);ok{
tok_mem= append(tok_mem,2+len(f2))
return s,func(){
f3()
call(f2)
call(f1)
reduce(ss,3+len(f2),SelectStmt,tok_mem...)
},true
}
}

/*:286*/
//line goweave.w:2384

case CommClause:/*287:*/
//line goweave.w:4206

if s,f1,ok:= sequence(ss,CommCase,colon);ok{
tok_mem:= append([]interface{}{},0,1,force)
s,f2,t,ok:= optional(s,2,pair{cat:Statement,mand:true})
if ok{
tok_mem= append(tok_mem,indent,t,outdent)
}
return s,func(){
call(f2)
call(f1)
reduce(ss,2+len(f2),CommClause,tok_mem...)
},true
}

/*:287*/
//line goweave.w:2385

case CommCase:/*288:*/
//line goweave.w:4220

if s,f1,ok:= one(ss,case_token);ok{
if s,f2,ok:= any(s,SendStmt,RecvStmt);ok{
return s,func(){
f2()
f1()
reduce(ss,2,CommCase,0,break_space,1)
},true
}
}else if s,f,ok:= one(ss,default_token);ok{
return s,func(){
f()
reduce(ss,1,CommCase,0)
},true
}else if s,f,ok:= one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,CommCase,0,force)
},true
}

/*:288*/
//line goweave.w:2386

case RecvStmt:/*289:*/
//line goweave.w:4241

if s,f1,ok:= one(ss,ExpressionList);ok{
if s,f2,ok:= any(s,eq,col_eq);ok{
if s,f3,ok:= one(s,Expression);ok{
return s,func(){
f3()
f2()
f1()
reduce(ss,3,RecvStmt,0,1,2)
},true
}
}
}else if s,f,ok:= one(s,Expression);ok{
return s,func(){
f()
reduce(ss,1,RecvStmt,0)
},true
}

/*:289*/
//line goweave.w:2387

case SendStmt:/*290:*/
//line goweave.w:4260

if s,f,ok:= sequence(ss,Expression,direct,Expression);ok{
return s,func(){
call(f)
reduce(ss,3,SendStmt,0,1,2)
},true
}

/*:290*/
//line goweave.w:2388

case ForStmt:/*293:*/
//line goweave.w:4304

if s,f1,ok:= one(ss,for_token);ok{
if s,f2,ok:= sequence(s,Expression,Block);ok{
return s,func(){
call(f2)
f1()
reduce(ss,3,ForStmt,0,break_space,1,break_space,2)
},true
}else if s,f2,ok:= sequence(s,ForClause,Block);ok{
return s,func(){
call(f2)
f1()
reduce(ss,3,ForStmt,0,break_space,1,break_space,2)
},true
}else if s,f2,ok:= sequence(s,RangeClause,Block);ok{
return s,func(){
call(f2)
f1()
reduce(ss,3,ForStmt,0,break_space,1,break_space,2)
},true
}else if s,f2,ok:= one(s,Block);ok{
return s,func(){
f2()
f1()
reduce(ss,2,ForStmt,0,1)
},true
}
}

/*:293*/
//line goweave.w:2389

case ForClause:/*294:*/
//line goweave.w:4333

var tok_mem[]interface{}
c:= 0
s,f1,ok:= one(ss,SimpleStmt)
if ok{
tok_mem= append(tok_mem,c)
c++
}
f2:= empty
s,f2,ok= one(s,semi)
if ok{
if len(scrap_info[c].trans)==0{
tok_mem= append(tok_mem,c)
}else{
tok_mem= append(tok_mem,';')
}
c++
f3:= empty
if s,f3,ok= one(s,Expression);ok{
tok_mem= append(tok_mem,break_space,c)
c++
}
if s,f4,ok:= one(s,semi);ok{
if len(scrap_info[c].trans)==0{
tok_mem= append(tok_mem,c)
}else{
tok_mem= append(tok_mem,';')
}
c++
f5:= empty
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

/*:294*/
//line goweave.w:2390

case RangeClause:/*295:*/
//line goweave.w:4378

if s,f1,ok:= one(ss,ExpressionList);ok{
if s,f2,ok:= any(s,eq,col_eq);ok{
if s,f3,ok:= sequence(s,range_token,Expression);ok{
return s,func(){
call(f3)
f2()
f1()
reduce(ss,4,RangeClause,0,1,2,break_space,3)
},true
}
}
}

/*:295*/
//line goweave.w:2391

case DeferStmt:/*297:*/
//line goweave.w:4423

if s,f,ok:= sequence(ss,defer_token,Expression);ok{
return s,func(){
call(f)
reduce(ss,2,DeferStmt,0,break_space,1)
},true
}

/*:297*/
//line goweave.w:2392

case IncDecStmt:/*299:*/
//line goweave.w:4444

if s,f1,ok:= one(ss,Expression);ok{
if s,f2,ok:= any(s,plus_plus,minus_minus);ok{
return s,func(){
f2()
f1()
reduce(ss,2,IncDecStmt,0,1)
},true
}
}

/*:299*/
//line goweave.w:2393

case Assignment:/*301:*/
//line goweave.w:4465

if s,f,ok:= sequence(ss,ExpressionList,assign_op,ExpressionList);ok{
return s,func(){
call(f)
reduce(ss,3,Assignment,0,1,2)
},true
}

/*:301*/
//line goweave.w:2394

case assign_op:/*303:*/
//line goweave.w:4528

if s,f,ok:= sequence(ss,binary_op,eq);ok{
return s,func(){
call(f)
reduce(ss,2,assign_op,math_rel,'{',0,'}','{',1,'}','}')
},true
}else if s,f,ok:= one(ss,eq);ok{
return s,func(){
f()
reduce(ss,1,assign_op,0)
},true
}

/*:303*/
//line goweave.w:2395

case ShortVarDecl:/*304:*/
//line goweave.w:4541

if s,f,ok:= sequence(ss,IdentifierList,col_eq,ExpressionList);ok{
return s,func(){
call(f)
reduce(ss,3,ShortVarDecl,0,1,2)
},true
}

/*:304*/
//line goweave.w:2396

case QualifiedIdent:/*306:*/
//line goweave.w:4568

if s,f1,ok:= any(ss,identifier,PackageName);ok{
if s,f2,ok:= sequence(s,dot,identifier);ok{
return s,func(){
call(f2)
f1()
reduce(ss,3,QualifiedIdent,0,1,2)
},true
}else{
return s,func(){
f1()
reduce(ss,1,QualifiedIdent,0)
},true
}
}

/*:306*/
//line goweave.w:2397

case MethodExpr:/*307:*/
//line goweave.w:4584

if s,f,ok:= sequence(ss,ReceiverType,dot,identifier);ok{
return s,func(){
call(f)
reduce(ss,3,MethodExpr,0,1,2)
},true
}

/*:307*/
//line goweave.w:2398

case ReceiverType:/*308:*/
//line goweave.w:4592

if s,f,ok:= one(ss,Type);ok{
return s,func(){
f()
reduce(ss,1,ReceiverType,0)
},true
}else if s,f,ok:= sequence(ss,lpar,asterisk,Type,rpar);ok{
return s,func(){
call(f)
reduce(ss,4,ReceiverType,0,1,2,3)
},true
}

/*:308*/
//line goweave.w:2399

case Conversion:/*309:*/
//line goweave.w:4605

if s,f,ok:= sequence(ss,Type,lpar,Expression,rpar);ok{
return s,func(){
call(f)
reduce(ss,4,Conversion,0,1,2,3)
},true
}

/*:309*/
//line goweave.w:2400

case BuiltinCall:/*310:*/
//line goweave.w:4613

if s,f1,ok:= sequence(ss,identifier,lpar);ok{
tok_mem:= append([]interface{}{},0,1)
s,f2,t,ok:= optional(s,2,pair{cat:BuiltinArgs,mand:true},pair{cat:comma,mand:false})
if ok{
tok_mem= append(tok_mem,t)
}
if s,f3,ok:= one(s,rpar);ok{
tok_mem= append(tok_mem,2+len(f2))
return s,func(){
f3()
call(f2)
call(f1)
reduce(ss,3+len(f2),BuiltinCall,tok_mem...)
},true
}
}

/*:310*/
//line goweave.w:2401

case BuiltinArgs:/*311:*/
//line goweave.w:4631

if s,f1,ok:= one(ss,Type);ok{
tok_mem:= append([]interface{}{},0)
s,f2,t,ok:= optional(s,1,pair{cat:comma,mand:true},pair{cat:ExpressionList,mand:true})
if ok{
tok_mem= append(tok_mem,t)
}
return s,func(){
call(f2)
f1()
reduce(ss,1+len(f2),BuiltinArgs,tok_mem...)
},true
}else if s,f,ok:= one(s,ExpressionList);ok{
return s,func(){
f()
reduce(ss,1,BuiltinArgs,0)
},true
}

/*:311*/
//line goweave.w:2402

case Selector:/*312:*/
//line goweave.w:4650

if s,f,ok:= sequence(ss,dot,identifier);ok{
return s,func(){
call(f)
reduce(ss,2,Selector,0,1)
},true
}

/*:312*/
//line goweave.w:2403

case Index:/*313:*/
//line goweave.w:4658

if s,f,ok:= sequence(ss,lbracket,Expression,rbracket);ok{
return s,func(){
call(f)
reduce(ss,3,Index,0,1,2)
},true
}

/*:313*/
//line goweave.w:2404

case Slice:/*314:*/
//line goweave.w:4666

if s,f1,ok:= one(ss,lbracket);ok{
tok_mem:= append([]interface{}{},0)
s,f2,t1,ok:= optional(s,1,pair{cat:Expression,mand:false})
if ok{
tok_mem= append(tok_mem,t1)
}
if s,f3,ok:= one(s,colon);ok{
tok_mem= append(tok_mem,1+len(t1))
s,f4,t2,ok:= optional(s,1+len(t1)+1,pair{cat:Expression,mand:false})
if ok{
tok_mem= append(tok_mem,t2)
}
if s,f5,ok:= one(s,rbracket);ok{
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

/*:314*/
//line goweave.w:2405

case TypeAssertion:/*315:*/
//line goweave.w:4693

if s,f,ok:= sequence(ss,dot,lpar,Type,rpar);ok{
return s,func(){
call(f)
reduce(ss,4,TypeAssertion,0,1,2,3)
},true
}

/*:315*/
//line goweave.w:2406

case Call:/*316:*/
//line goweave.w:4701

if s,f1,ok:= one(ss,lpar);ok{
tok_mem:= append([]interface{}{},0)
c:= 1
s,f2,ok:= one(s,ExpressionList)
f3:= empty
if ok{
tok_mem= append(tok_mem,c)
c++
s,f3,ok= one(s,dot_dot_dot)
if ok{
tok_mem= append(tok_mem,c)
c++
}
}
if s,f4,ok:= one(s,rpar);ok{
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

/*:316*/
//line goweave.w:2407

case unary_op:/*317:*/
//line goweave.w:4729

if s,f,ok:= any(ss,asterisk,direct,add_op);ok{
return s,func(){
f()
reduce(ss,1,unary_op,0)
},true
}

/*:317*/
//line goweave.w:2408

}
m= "not found"
return ss,empty,false
}

/*:197*//*198:*/
//line goweave.w:2418

func sequence(ss[]scrap,cats...rune)([]scrap,[]reducing,bool){
var fs[]reducing
s:= ss
for _,v:= range cats{
f:= empty
ok:= false
if s,f,ok= one(s,v);!ok{
return ss,nil,false
}
fs= append(fs,f)
}
return s,fs,true
}

/*:198*//*199:*/
//line goweave.w:2436

func any(s[]scrap,cats...rune)([]scrap,reducing,bool){
for _,v:= range cats{
if s,f,ok:= one(s,v);ok{
return s,f,ok
}
}
return s,empty,false
}


/*:199*//*200:*/
//line goweave.w:2451

func optional(ss[]scrap,g int,cats...pair)([]scrap,[]reducing,[]int,bool){
var trans[]int
var funcs[]reducing
ok:= false
for len(ss)> 0{
var t[]int
var fs[]reducing
s:= ss
exit:= false
for _,v:= range cats{
f:= empty
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


/*:200*//*201:*/
//line goweave.w:2491

func isNotCat(i int,cat int32)bool{
if i<0||i>=len(scrap_info){
return false
}
return scrap_info[i].cat!=cat
}

/*:201*//*205:*/
//line goweave.w:2553

func find_first_ident(p[]interface{})[]interface{}{
for i,j:= range p{
switch r:= j.(type){
case res_token:
if name_dir[r].ilk==case_token{
return nil
}
if name_dir[r].ilk!=Type{
break
}
return p[i:i+1]
case id_token:
return p[i:i+1]
case list_token,inner_list_token:
if q:= find_first_ident(r.([]interface{}));q!=nil{
return q
}
case rune:
if r==inserted{
return nil
}
}
}
return nil
}

/*:205*//*206:*/
//line goweave.w:2584


func make_reserved(s scrap,c rune){
tok_ptr:= find_first_ident(s.trans)
if tok_ptr==nil{
return
}
name_dir[tok_ptr[0].(id_token)].ilk= c
tok_ptr[0]= res_token(tok_ptr[0].(id_token))
}

/*:206*//*207:*/
//line goweave.w:2604


func make_underlined(p int){
tok_ptr:= find_first_ident(scrap_info[p].trans)
if tok_ptr==nil{
return
}
xref_switch= def_flag
underline_xref(tok_ptr[0].(id_token))
}

/*:207*//*209:*/
//line goweave.w:2621

func underline_xref(p id_token){
q:= name_dir[p].xref
if flags['x']==false{
return
}
m:= section_count+xref_switch
for q!=0{
n:= xmem[q].num
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
/*210:*/
//line goweave.w:2650

append_xref(0)
xmem[len(xmem)-1].xlink= name_dir[p].xref
r:= int32(len(xmem)-1)
name_dir[p].xref= r
for xmem[r].xlink!=q{
xmem[r].num= xmem[xmem[r].xlink].num
r= xmem[r].xlink
}
xmem[r].num= m

/*:210*/
//line goweave.w:2640

}

/*:209*//*319:*/
//line goweave.w:4760

func reduce(ss[]scrap,k int,c rune,s...interface{}){
var trans[]interface{}
cur_mathness:= maybe_math
init_mathness:= maybe_math

for _,t:= range s{
switch v:= t.(type){
case rune:
if v==' '||(v>=big_cancel&&v<=big_force){
if cur_mathness==maybe_math{
init_mathness= no_math
}else if cur_mathness==yes_math{
trans= append(trans,"{}$")
}
cur_mathness= no_math
}else{
if ss[0].mathness==maybe_math{
init_mathness= yes_math
}else if ss[0].mathness==no_math{
trans= append(trans,"${}")
}
cur_mathness= yes_math
}
trans= append(trans,v)
case int:
s:= ss[v]
/*318:*/
//line goweave.w:4739

switch s.mathness%4{
case no_math:
if cur_mathness==maybe_math{
init_mathness= no_math
}else if cur_mathness==yes_math{
trans= append(trans,"{}$")
}
cur_mathness= s.mathness/4
case yes_math:
if cur_mathness==maybe_math{
init_mathness= yes_math
}else if cur_mathness==no_math{
trans= append(trans,"${}")
}
cur_mathness= s.mathness/4
case maybe_math:
}
trans= append(trans,s.trans...)

/*:318*/
//line goweave.w:4787

case[]int:
for _,v:= range v{
if v==-1{
continue
}
s:= ss[v]
/*318:*/
//line goweave.w:4739

switch s.mathness%4{
case no_math:
if cur_mathness==maybe_math{
init_mathness= no_math
}else if cur_mathness==yes_math{
trans= append(trans,"{}$")
}
cur_mathness= s.mathness/4
case yes_math:
if cur_mathness==maybe_math{
init_mathness= yes_math
}else if cur_mathness==no_math{
trans= append(trans,"${}")
}
cur_mathness= s.mathness/4
case maybe_math:
}
trans= append(trans,s.trans...)

/*:318*/
//line goweave.w:4794

}
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
f:= fmt.Sprintf("reduce %q %v",cat_name[c],k)
/*323:*/
//line goweave.w:4841

{
if(tracing&2)==2{
fmt.Printf("%s:",f)
for i,v:= range scrap_info{
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

/*:323*/
//line goweave.w:4810

if(tracing&8)==8{
fmt.Printf("translation of %s: %v\n",cat_name[c],trans)
}
}

/*:319*//*324:*/
//line goweave.w:4873


func translate()[]interface{}{
pp:= 0
/*327:*/
//line goweave.w:4917

if(tracing&2)==2{
fmt.Printf("\nTracing after l. %d:\n",line[include_depth])
mark_harmless()

}

/*:327*/
//line goweave.w:4877

/*328:*/
//line goweave.w:4924

for i:= 1;i<len(scrap_info);{
if scrap_info[i].cat==insert{
reduce(scrap_info[i-1:],2,scrap_info[i-1].cat,0,1)
continue
}
i++
}

if len(scrap_info)> 1&&scrap_info[0].cat==insert&&scrap_info[1].cat!=zero{
reduce(scrap_info,2,scrap_info[1].cat,0,1)
}


/*:328*/
//line goweave.w:4878

/*321:*/
//line goweave.w:4825

for{
if pp>=len(scrap_info){
break
}
/*203:*/
//line goweave.w:2504
{
_,f,ok:= func(ss[]scrap)([]scrap,reducing,bool){
switch ss[0].cat{
case package_token:
/*212:*/
//line goweave.w:2666

if s,f,ok:= sequence(ss,package_token,identifier);ok{
return s,func(){
call(f)
reduce(ss,2,PackageClause,0,break_space,1,big_force)
},true
}

/*:212*/
//line goweave.w:2508

case import_token:
/*220:*/
//line goweave.w:2838

if s,f1,ok:= one(ss,import_token);ok{
if s,f2,ok:= one(s,ImportSpec);ok{
return s,func(){
f2()
f1()
reduce(ss,2,ImportDecl,0,break_space,1,force)
},true
}else if s,f2,ok:= one(s,lpar);ok{
tok_mem:= append([]interface{}{},0,1)
s,f3,t,ok:= optional(s,2,pair{cat:ImportSpec,mand:true})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
}
if s,f4,ok:= one(s,rpar);ok{
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

/*:220*/
//line goweave.w:2510

case struct_token:
/*235:*/
//line goweave.w:3249

if s,f1,ok:= sequence(ss,struct_token,lbrace);ok{
tok_mem:= append([]interface{}{},0,1)
s,f2,t,ok:= optional(s,2,pair{cat:FieldDecl,mand:true})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
}
if s,f3,ok:= one(s,rbrace);ok{
tok_mem= append(tok_mem,2+len(f2))
return s,func(){
f3()
call(f2)
call(f1)
reduce(ss,3+len(f2),StructType,tok_mem...)
},true
}
}

/*:235*/
//line goweave.w:2512

case interface_token:
/*242:*/
//line goweave.w:3388

if s,f1,ok:= sequence(ss,interface_token,lbrace);ok{
tok_mem:= append([]interface{}{},0,1)
s,f2,t,ok:= optional(s,2,pair{cat:MethodSpec,mand:true})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
}
if s,f3,ok:= one(s,rbrace);ok{
tok_mem= append(tok_mem,2+len(f2))
return s,func(){
f3()
call(f2)
call(f1)
reduce(ss,3+len(f2),InterfaceType,tok_mem...)
},true
}
}

/*:242*/
//line goweave.w:2514

case func_token:
/*222:*/
//line goweave.w:2890

if s,f1,ok:= sequence(ss,func_token,identifier,Signature);ok{
if s,f2,ok:= sequence(s,Block,semi);ok{
return s,func(){
call(f2)
call(f1)
make_underlined(1)
reduce(ss,5,FunctionDecl,0,break_space,1,2,3,4,big_force)
},true
}else if s,f2,ok:= one(s,semi);ok{
return s,func(){
f2()
call(f1)
make_underlined(1)
reduce(ss,4,FunctionDecl,0,break_space,1,2,3,big_force)
},true
}
}

/*:222*/
//line goweave.w:2516

/*224:*/
//line goweave.w:2924

if s,f1,ok:= sequence(ss,func_token,Receiver,identifier,Signature);ok{
if s,f2,ok:= one(s,Block);ok{
return s,func(){
f2()
call(f1)
make_underlined(2)
reduce(ss,5,MethodDecl,0,break_space,1,break_space,2,3,4)
},true
}else{
return s,func(){
call(f1)
make_underlined(2)
reduce(ss,4,MethodDecl,0,break_space,1,break_space,2,3)
},true
}
}

/*:224*/
//line goweave.w:2517

/*260:*/
//line goweave.w:3671

if s,f,ok:= sequence(ss,func_token,Signature);ok{
return s,func(){
call(f)
reduce(ss,2,FunctionType,0,1)
},true
}

/*:260*/
//line goweave.w:2518

default:
/*263:*/
//line goweave.w:3706

if s,f,ok:= any(ss,
ImportDecl,
ConstDecl,
VarDecl,
TypeDecl,
LabeledStmt);ok{
return s,func(){
f()
reduce(ss,1,Statement,0)
},true
}else if s,f1,ok:= any(ss,
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
if s,f2,ok:= one(s,semi);ok{
return s,func(){
f2()
f1()
reduce(ss,2,Statement,0,1,force)
},true
}else if _,_,ok:= any(s,rpar,rbrace);ok{
return s,func(){
f1()
reduce(ss,1,Statement,0,force)
},true
}
}else if s,f,ok:= one(ss,semi);ok{
return s,func(){
f()
reduce(ss,1,Statement,0,force)
},true
}else if s,f,ok:= one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,Statement,0,force)
},true
}

/*:263*/
//line goweave.w:2520

/*230:*/
//line goweave.w:3105

if s,f1,ok:= sequence(ss,identifier,str);ok{
if s,f2,ok:= one(s,semi);ok{
return s,func(){
f2()
call(f1)
make_reserved(ss[0],PackageName)
reduce(ss,3,ImportSpec,0,break_space,1,2,force)
},true
}else if _,_,ok:= any(s,rpar,rbrace);ok{
return s,func(){
call(f1)
make_reserved(ss[0],PackageName)
reduce(ss,2,ImportSpec,0,break_space,1,force)
},true
}
}else if s,f1,ok:= sequence(ss,dot,str);ok{
if s,f2,ok:= one(s,semi);ok{
return s,func(){
f2()
call(f1)
reduce(ss,3,ImportSpec,0,break_space,1,2,force)
},true
}else if _,_,ok:= any(s,rpar,rbrace);ok{
return s,func(){
call(f1)
reduce(ss,2,ImportSpec,0,break_space,1,force)
},true
}
}else if s,f1,ok:= one(ss,str);ok{
if s,f2,ok:= one(s,semi);ok{
return s,func(){
f2()
f1()
reduce(ss,2,ImportSpec,0,1,force)
},true
}else if _,_,ok:= any(s,rpar,rbrace);ok{
return s,func(){
f1()
reduce(ss,1,ImportSpec,0,force)
},true
}
}else if s,f,ok:= one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,ImportSpec,0,force)
},true
}

/*:230*/
//line goweave.w:2521

/*227:*/
//line goweave.w:2987

if s,f1,ok:= one(ss,IdentifierList);ok{
if s,f2,ok:= sequence(s,Type,eq,ExpressionList);ok{
if s,f3,ok:= one(s,semi);ok{
return s,func(){
f3()
call(f2)
f1()
reduce(ss,5,ConstSpec,0,break_space,1,break_space,2,break_space,3,4,force)
},true
}else if _,_,ok:= any(s,rpar,rbrace);ok{
return s,func(){
call(f2)
f1()
reduce(ss,4,ConstSpec,0,break_space,1,break_space,2,break_space,3,force)
},true
}
}else if s,f2,ok:= sequence(s,eq,ExpressionList);ok{
if s,f3,ok:= one(s,semi);ok{
return s,func(){
f3()
call(f2)
f1()
reduce(ss,4,ConstSpec,0,break_space,1,break_space,2,3,force)
},true
}else if _,_,ok:= any(s,rpar,rbrace);ok{
return s,func(){
call(f2)
f1()
reduce(ss,3,ConstSpec,0,break_space,1,break_space,2,force)
},true
}
}
}else if s,f,ok:= one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,ConstSpec,0,force)
},true
}

/*:227*/
//line goweave.w:2522

/*229:*/
//line goweave.w:3048

if s,f1,ok:= one(ss,IdentifierList);ok{
if s,f2,ok:= one(s,Type);ok{
if s,f3,ok:= sequence(s,eq,ExpressionList);ok{
if s,f4,ok:= one(s,semi);ok{
return s,func(){
f4()
call(f3)
f2()
f1()
reduce(ss,5,VarSpec,0,break_space,1,2,3,4,force)
},true
}else if _,_,ok:= any(s,rpar,rbrace);ok{
return s,func(){
call(f3)
f2()
f1()
reduce(ss,4,VarSpec,0,break_space,1,2,3,force)
},true
}
}else if s,f3,ok:= one(s,semi);ok{
return s,func(){
f3()
f2()
f1()
reduce(ss,3,VarSpec,0,break_space,1,2,force)
},true
}else if _,_,ok:= any(s,rpar,rbrace);ok{
return s,func(){
f2()
f1()
reduce(ss,2,VarSpec,0,break_space,1,force)
},true
}
}else if s,f2,ok:= sequence(s,eq,ExpressionList);ok{
if s,f3,ok:= one(s,semi);ok{
return s,func(){
f3()
call(f2)
f1()
reduce(ss,4,VarSpec,0,1,2,3,force)
},true
}else if _,_,ok:= any(s,rpar,rbrace);ok{
return s,func(){
call(f2)
f1()
reduce(ss,3,VarSpec,0,1,2,force)
},true
}
}
}else if s,f,ok:= one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,VarSpec,0,force)
},true
}

/*:229*/
//line goweave.w:2523

/*228:*/
//line goweave.w:3027

if s,f1,ok:= sequence(ss,identifier,Type);ok{
if s,f2,ok:= one(s,semi);ok{
return s,func(){
f2()
call(f1)
reduce(ss,3,TypeSpec,0,break_space,1,2,force)
},true
}else if _,_,ok:= any(s,rpar,rbrace);ok{
return s,func(){
call(f1)
reduce(ss,2,TypeSpec,0,break_space,1,force)
},true
}
}else if s,f,ok:= one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,TypeSpec,0,force)
},true
}

/*:228*/
//line goweave.w:2524

/*231:*/
//line goweave.w:3154

if s,f1,ok:= sequence(ss,IdentifierList,Type);ok{
tok_mem:= append([]interface{}{},0,break_space,1)
c:= 2
s,f2,ok:= one(s,str)
if ok{
tok_mem= append(tok_mem,break_space,2)
c++
}
if s,f3,ok:= one(s,semi);ok{
tok_mem= append(tok_mem,c,force)
c++
return s,func(){
f3()
f2()
call(f1)
reduce(ss,c,FieldDecl,tok_mem...)
},true
}else if _,_,ok:= any(s,rpar,rbrace);ok{
tok_mem= append(tok_mem,force)
return s,func(){
f2()
call(f1)
reduce(ss,c,FieldDecl,tok_mem...)
},true
}
}else if s,f1,ok:= one(ss,AnonymousField);ok{
tok_mem:= append([]interface{}{},0)
c:= 1
s,f2,ok:= one(s,str)
if ok{
tok_mem= append(tok_mem,break_space,1)
c++
}
tok_mem= append(tok_mem,force)
if s,f3,ok:= one(s,semi);ok{
c++
return s,func(){
f3()
f2()
f1()
reduce(ss,c,FieldDecl,tok_mem...)
},true
}else if _,_,ok:= any(s,rpar,rbrace);ok{
return s,func(){
f2()
f1()
reduce(ss,c,FieldDecl,tok_mem...)
},true
}
}else if s,f,ok:= one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,FieldDecl,0,force)
},true
}

/*:231*/
//line goweave.w:2525

/*280:*/
//line goweave.w:4021

if s,f1,ok:= sequence(ss,case_token,ExpressionList,colon);ok{
tok_mem:= append([]interface{}{},0,break_space,1,2)
s,f2,t,ok:= optional(s,3,pair{cat:Statement,mand:true})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
}
return s,func(){
call(f2)
call(f1)
reduce(ss,3+len(f2),ExprCaseClause,tok_mem...)
},true
}else if s,f1,ok:= sequence(ss,default_token,colon);ok{
tok_mem:= append([]interface{}{},0,1,force)
s,f2,t,ok:= optional(s,2,pair{cat:Statement,mand:true})
if ok{
tok_mem= append(tok_mem,force,indent,t,outdent)
}
return s,func(){
call(f2)
call(f1)
reduce(ss,2+len(f2),ExprCaseClause,tok_mem...)
},true
}else if s,f,ok:= one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,ExprCaseClause,0,force)
},true
}

/*:280*/
//line goweave.w:2526

/*283:*/
//line goweave.w:4101

if s,f1,ok:= sequence(ss,TypeSwitchCase,colon);ok{
tok_mem:= append([]interface{}{},0,1,force)
s,f2,t,ok:= optional(s,2,pair{cat:Statement,mand:true})
if ok{
tok_mem= append(tok_mem,indent,t,outdent)
}
return s,func(){
call(f2)
call(f1)
reduce(ss,2+len(f2),TypeCaseClause,tok_mem...)
},true
}else if s,f,ok:= one(ss,section_scrap);ok{
return s,func(){
f()
reduce(ss,1,TypeCaseClause,0,force)
},true
}

/*:283*/
//line goweave.w:2527

/*287:*/
//line goweave.w:4206

if s,f1,ok:= sequence(ss,CommCase,colon);ok{
tok_mem:= append([]interface{}{},0,1,force)
s,f2,t,ok:= optional(s,2,pair{cat:Statement,mand:true})
if ok{
tok_mem= append(tok_mem,indent,t,outdent)
}
return s,func(){
call(f2)
call(f1)
reduce(ss,2+len(f2),CommClause,tok_mem...)
},true
}

/*:287*/
//line goweave.w:2528

}
return ss,empty,false
}(scrap_info[pp:])
if ok{
f()
}
pp++
}

/*:203*/
//line goweave.w:4830

}

/*:321*/
//line goweave.w:4879

/*325:*/
//line goweave.w:4888
{
/*326:*/
//line goweave.w:4906

if len(scrap_info)> 0&&tracing==1{
fmt.Printf("\nIrreducible scrap sequence in section %d:",section_count)

mark_harmless()
for i,_:= range scrap_info{
fmt.Printf(" ")
print_cat(scrap_info[i].cat)
}
}

/*:326*/
//line goweave.w:4889

var tok_mem[]interface{}
for i,v:= range scrap_info{
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

/*:325*/
//line goweave.w:4880

}

/*:324*//*329:*/
//line goweave.w:4953


func Go_parse(spec_ctrl rune){
for next_control<format_code||next_control==spec_ctrl{
/*331:*/
//line goweave.w:4974

switch(next_control){
case section_name:
app_scrap(section_scrap,maybe_math,section_token(cur_section))
case str,constant,verbatim:
/*333:*/
//line goweave.w:5142

count:= -1
var tok_mem[]interface{}
if next_control==constant{
tok_mem= append(tok_mem,"\\T{")

}else if next_control==str{
count= 20
tok_mem= append(tok_mem,"\\.{")

}else{
tok_mem= append(tok_mem,"\\vb{")
}

for i:= 0;i<len(id);{
if count==0{
tok_mem= append(tok_mem,"}\\)\\.{")
count= 20

}
switch(id[i]){
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

/*:333*/
//line goweave.w:4979

case identifier:
app_cur_id()
case TeX_string:
/*334:*/
//line goweave.w:5196

tok_mem:= append([]interface{}{},"\\hbox{")
for i:= 0;i<len(id);{
if id[i]=='@'{
i++
}
tok_mem= append(tok_mem,id[i])
i++
}
tok_mem= append(tok_mem,'}')
app_scrap(insert,no_math,tok_mem...)

/*:334*/
//line goweave.w:4983

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
/*332:*/
//line goweave.w:5093

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
app_scrap(dot_dot_dot,yes_math,"\\,\\ldots\\,")


case col_eq:
app_scrap(col_eq,yes_math,":\\K")

case direct:
app_scrap(direct,yes_math,"\\leftarrow")

case and_not:
app_scrap(mul_op,yes_math,"\\AND\\CF")


/*:332*/
//line goweave.w:5060

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

/*:331*/
//line goweave.w:4957

next_control= get_next()
if next_control=='|'||next_control==begin_comment||
next_control==begin_short_comment{
return
}
}
}

/*:329*//*330:*/
//line goweave.w:4969

func app_scrap(c int32,b int32,t...interface{}){
scrap_info= append(scrap_info,scrap{cat:c,trans:t,mathness:5*b,})
}

/*:330*//*336:*/
//line goweave.w:5212

func app_cur_id(){
p:= id_lookup(id,normal)
if name_dir[p].ilk<=custom{
a1:= identifier
a2:= maybe_math
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
}

/*:336*//*337:*/
//line goweave.w:5239

func Go_translate()[]interface{}{
save_scraps:= scrap_info
scrap_info= nil
Go_parse(section_name)
if next_control!='|'{
err_print("! Missing '|' after Go text")

}
app_scrap(semi,no_math)
app_scrap(insert,maybe_math,cancel)

p:= translate()
scrap_info= save_scraps
return p
}

/*:337*//*338:*/
//line goweave.w:5266


func outer_parse(){
for next_control<format_code{
var tok_mem[]interface{}
if next_control!=begin_comment&&next_control!=begin_short_comment{
Go_parse(ignore)
}else{
is_long_comment:= (next_control==begin_comment)
tok_mem= append(tok_mem,cancel,inserted)
if is_long_comment{
tok_mem= append(tok_mem,"\\C{")

}else{
tok_mem= append(tok_mem,"\\SHC{")
}

bal,tok_mem:= copy_comment(is_long_comment,1,tok_mem)
next_control= ignore
for bal> 0{
p:= tok_mem
tok_mem= nil
q:= Go_translate()
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
bal= 0
}
}
tok_mem= append(tok_mem,force)
app_scrap(insert,no_math,tok_mem...)

}
}
}

/*:338*//*340:*/
//line goweave.w:5351
type mode int

/*:340*//*343:*/
//line goweave.w:5366
func init_stack(){
stack= make([]output_state,0,100)
cur_state.mode_field= outer
}

/*:343*//*345:*/
//line goweave.w:5379


func push_level(tokens[]interface{}){
stack= append(stack,output_state{tok_field:cur_state.tok_field,mode_field:cur_state.mode_field,})
cur_state.tok_field= tokens
}

/*:345*//*346:*/
//line goweave.w:5390

func pop_level()bool{
if len(stack)==0{
return false
}
p:= len(stack)-1
cur_state= stack[p]
stack= stack[:p]
return true
}

/*:346*//*349:*/
//line goweave.w:5419


func get_output()rune{
restart:
for len(cur_state.tok_field)==0{
if!pop_level(){
return-1
}
}
val:= cur_state.tok_field[0]
cur_state.tok_field= cur_state.tok_field[1:]
switch tok:= val.(type){
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
for _,v:= range tok{
tok_mem= append(tok_mem,v)
}
push_level(tok_mem)
goto restart
}
panic(fmt.Sprintf("Invalid type of scrap: %T(%v)",val,val))
}

/*:349*//*350:*/
//line goweave.w:5478


func output_Go(){
save_next_control:= next_control
next_control= ignore
p:= Go_translate()
if flags['e']{
out_str("\\PB{")
make_output(inner_list_token(p))
out('}')

}else{
make_output(inner_list_token(p))
}
next_control= save_next_control
}

/*:350*//*352:*/
//line goweave.w:5498


func make_output(p interface{}){
var c int
tok_mem:= append([]interface{}{},p,end_translation)
push_level(tok_mem)
tok_mem= nil
var b rune
for{
a:= get_output()
reswitch:
switch a{
case end_translation:
return
case identifier,res_word:
/*353:*/
//line goweave.w:5559

out('\\')
if a==identifier{
if name_dir[cur_name].ilk==custom&&!doing_format{
/*354:*/
//line goweave.w:5592

for _,v:= range name_dir[cur_name].name{
if v=='_'{
out('x')
}else{
out(v)
}
}
break

/*:354*/
//line goweave.w:5563

}else if is_tiny(cur_name){
out('|')

}else{
delim:= '.'
for _,v:= range name_dir[cur_name].name{
if unicode.IsLower(v){
delim= '\\'
break
}
}
out(delim)
}


}else{
out('&')
}

if is_tiny(cur_name){
if name_dir[cur_name].name[0]=='_'{
out('\\')
}
out(name_dir[cur_name].name[0])
}else{
out_name(cur_name,true)
}

/*:353*/
//line goweave.w:5513

case section_code:
/*358:*/
//line goweave.w:5697
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
out('0')
}
out(':')
if an_output{
out_str("\\.{")

}
/*359:*/
//line goweave.w:5732

scratch:= sprint_section_name(cur_name)
cur_section_name:= cur_name
for i:= 0;i<len(scratch);{
b= scratch[i]
i++
if b=='@'{
/*360:*/
//line goweave.w:5775

ii:= i
i++
if ii<len(scratch)&&scratch[ii]!='@'{
fmt.Print("\n! Illegal control code in section name: <")

print_section_name(cur_section_name)
fmt.Print("> ")
mark_error()
}

/*:360*/
//line goweave.w:5739

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
/*361:*/
//line goweave.w:5792

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
/*362:*/
//line goweave.w:5823
{
buf= append(buf,b)
buf= append(buf,scratch[i])
i++
}

/*:362*/
//line goweave.w:5806

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

/*:361*/
//line goweave.w:5763

save_buf:= buffer
save_loc:= loc
buf= append(buf,'|')
buffer= buf
loc= 0
output_Go()
loc= save_loc
buffer= save_buf
}
}

/*:359*/
//line goweave.w:5725

if an_output{
out_str(" }")
}
out_str("\\X")
}

/*:358*/
//line goweave.w:5515

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
/*357:*/
//line goweave.w:5680

for;c> 0;c--{
out_str("\\1")

}
for;c<0;c++{
out_str("\\2")

}

/*:357*/
//line goweave.w:5540

goto reswitch
case indent,outdent,opt,backup,break_space,
force,big_force:
/*355:*/
//line goweave.w:5605

if a<break_space{
if cur_state.mode_field==outer{
out('\\')
out(a-cancel+'0')





if a==opt{
b= get_output()
if b!='0'||flags['f']==false{
out(b)
}else{
out_str("{-1}")
}
}
}else if a==opt{
b= get_output()
}
}else{
/*356:*/
//line goweave.w:5636
{
b= a
save_mode:= cur_state.mode_field
c= 0
for{
a= get_output()
if a==inserted{
continue
}
if a==cancel||a==big_cancel{
/*357:*/
//line goweave.w:5680

for;c> 0;c--{
out_str("\\1")

}
for;c<0;c++{
out_str("\\2")

}

/*:357*/
//line goweave.w:5646

goto reswitch
}
if a!=' '&&a<indent||a==backup||a> big_force{
if save_mode==outer{
if out_ptr> 3&&compare_runes(out_buf[out_ptr-3:out_ptr+1],[]rune("\\Y\\B"))==0{
goto reswitch
}
/*357:*/
//line goweave.w:5680

for;c> 0;c--{
out_str("\\1")

}
for;c<0;c++{
out_str("\\2")

}

/*:357*/
//line goweave.w:5654

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
b= a
}
}
}

/*:356*/
//line goweave.w:5627

}

/*:355*/
//line goweave.w:5544

case quoted_char:
out(cur_state.tok_field[0].(rune))
cur_state.tok_field= cur_state.tok_field[1:]
default:
out(a)
}
}
}

/*:352*//*364:*/
//line goweave.w:5837

func phase_two(){
reset_input()
if show_progress(){
fmt.Print("\nWriting the output file...")

}
section_count= 0
format_visible= true
copy_limbo()
finish_line()
flush_buffer(0,false,false)
for!input_has_ended{
/*367:*/
//line goweave.w:5885
{
section_count++
/*368:*/
//line goweave.w:5903

if loc-1>=len(buffer)||buffer[loc-1]!='*'{
out_str("\\M")

}else{
for loc<len(buffer)&&buffer[loc]==' '{
loc++
}
if loc<len(buffer)&&buffer[loc]=='*'{
sec_depth= -1
loc++
}else{
for sec_depth= 0;loc<len(buffer)&&unicode.IsDigit(buffer[loc]);loc++{
sec_depth= sec_depth*10+buffer[loc]-'0'
}
}
for loc<len(buffer)&&buffer[loc]==' '{
loc++
}
group_found= true
out_str("\\N")

{
s:= fmt.Sprintf("{%d}",sec_depth+1)
out_str(s)
}
if show_progress(){
fmt.Printf("*%d",section_count)
}
os.Stdout.Sync()
}
out_str("{")
out_str(section_str(section_count))
out_str("}")

/*:368*/
//line goweave.w:5887

save_position()
/*369:*/
//line goweave.w:5941

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
next_control= get_next()
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

/*:369*/
//line goweave.w:5889

/*370:*/
//line goweave.w:5971

space_checked= false
for next_control<=format_code{
init_stack()
/*373:*/
//line goweave.w:6021
{
doing_format= true
if buffer[loc-1]=='s'||buffer[loc-1]=='S'{
format_visible= false
}
if!space_checked{
emit_space_if_needed()
save_position()
}
tok_mem:= append([]interface{}{},"\\F")

next_control= get_next()
if next_control==identifier{
tok_mem= append(tok_mem,id_token(id_lookup(id,normal)),' ',break_space)
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

/*:373*/
//line goweave.w:5975

outer_parse()
finish_Go(format_visible)
format_visible= true
doing_format= false
}

/*:370*/
//line goweave.w:5890

/*375:*/
//line goweave.w:6056

this_section= -1
if next_control<=section_name{
emit_space_if_needed()
init_stack()
if next_control==begin_code{
next_control= get_next()
}else{
this_section= cur_section
/*376:*/
//line goweave.w:6078

for{
next_control= get_next()
if next_control!='+'{
break
}
}
var tok_mem[]interface{}
if next_control!='='&&next_control!=eq_eq{
err_print("! You need an = sign after the section name")

}else{
next_control= get_next()
}
if out_ptr> 1&&out_buf[out_ptr]=='Y'&&out_buf[out_ptr-1]=='\\'{
tok_mem= append(tok_mem,backup)
}


tok_mem= append(tok_mem,section_token(this_section))
cur_xref= name_dir[this_section].xref
if xmem[cur_xref].num==file_flag{
cur_xref= xmem[cur_xref].xlink
}
tok_mem= append(tok_mem,"${}")
if xmem[cur_xref].num!=section_count+def_flag{
tok_mem= append(tok_mem,"\\mathrel+")
this_section= -1
}
tok_mem= append(tok_mem,"\\E","{}$",force)

app_scrap(dead,no_math,tok_mem...)


/*:376*/
//line goweave.w:6066

}
for next_control<=section_name{
outer_parse()
/*377:*/
//line goweave.w:6112

if next_control<section_name{
err_print("! You can't do that in Go text")

next_control= get_next()
}else if next_control==section_name{
app_scrap(section_scrap,maybe_math,section_token(cur_section))
next_control= get_next()
}

/*:377*/
//line goweave.w:6070

}
finish_Go(true)
}

/*:375*/
//line goweave.w:5891

/*378:*/
//line goweave.w:6125

if this_section!=-1{
cur_xref= name_dir[this_section].xref
if xmem[cur_xref].num==file_flag{
an_output= true
cur_xref= xmem[cur_xref].xlink
}else{
an_output= false
}
if xmem[cur_xref].num> def_flag{
cur_xref= xmem[cur_xref].xlink
}
footnote(def_flag)
footnote(cite_flag)
footnote(0)
}

/*:378*/
//line goweave.w:5892

/*382:*/
//line goweave.w:6204

out_str("\\fi")
finish_line()

flush_buffer(0,false,false)

/*:382*/
//line goweave.w:5893

}

/*:367*/
//line goweave.w:5850

}
}

/*:364*//*365:*/
//line goweave.w:5862

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

/*:365*//*372:*/
//line goweave.w:5993



func finish_Go(visible bool){
if visible{
out_str("\\B")
app_scrap(insert,no_math,force)
p:= translate()

scrap_info= nil
make_output(list_token(p))
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

/*:372*//*380:*/
//line goweave.w:6155


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
/*381:*/
//line goweave.w:6182

q:= cur_xref
if xmem[xmem[q].xlink].num> flag{
out('s')
}
for{
out_str(section_str(xmem[cur_xref].num-flag))
cur_xref= xmem[cur_xref].xlink
if xmem[cur_xref].num<=flag{
break
}
if xmem[xmem[cur_xref].xlink].num> flag{
out_str(", ")
}else{
out_str("\\ET")

if cur_xref!=xmem[q].xlink{
out('s')
}
}
}

/*:381*/
//line goweave.w:6174

out('.')
}

/*:380*//*384:*/
//line goweave.w:6219

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
if f,err:= os.OpenFile(idx_file_name,
os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0666);err!=nil{
fatal("! Cannot open index file ",idx_file_name)

}else{
idx_file= f
}
if change_exists{
/*386:*/
//line goweave.w:6288
{

var k_section int32= 0
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

/*:386*/
//line goweave.w:6241

finish_line()
finish_line()
}
out_str("\\inx")
finish_line()

active_file= idx_file
/*388:*/
//line goweave.w:6319
{
for c:= 0;c<=255;c++{
bucket[c]= -1
}
for _,next_name:= range hash{
for next_name!=-1{
cur_name= next_name
next_name= name_dir[cur_name].llink
if name_dir[cur_name].xref!=0{
c:= name_dir[cur_name].name[0]
if unicode.IsUpper(c){
c= unicode.ToLower(c)
}
blink[cur_name]= bucket[c]
bucket[c]= cur_name
}
}
}
}

/*:388*/
//line goweave.w:6249

/*399:*/
//line goweave.w:6432

sort_ptr= 0
scrap_info= append(scrap_info,scrap{})
unbucket(1)
for sort_ptr> 0{
cur_depth= scrap_info[sort_ptr].cat
if blink[scrap_info[sort_ptr].head]==-1||cur_depth==infinity{
/*401:*/
//line goweave.w:6470
{
cur_name= scrap_info[sort_ptr].head
for{
out_str("\\I")

/*402:*/
//line goweave.w:6485

switch name_dir[cur_name].ilk{
case normal:
if is_tiny(cur_name){
out_str("\\|")

}else{
lowcase:= false
for _,v:= range name_dir[cur_name].name{
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
for _,v:= range name_dir[cur_name].name{
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
name_done:/*:402*/
//line goweave.w:6475

/*404:*/
//line goweave.w:6541

/*406:*/
//line goweave.w:6570

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

/*:406*/
//line goweave.w:6542

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

/*:404*/
//line goweave.w:6476

cur_name= blink[cur_name]
if cur_name==-1{
break
}
}
sort_ptr--
}

/*:401*/
//line goweave.w:6439

}else{
/*400:*/
//line goweave.w:6445
{
next_name:= scrap_info[sort_ptr].head
for{
var c rune
cur_name= next_name
next_name= blink[cur_name]
cur_byte= cur_depth
if cur_byte>=int32(len(name_dir[cur_name].name)){
c= 0
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

/*:400*/
//line goweave.w:6441

}
}

/*:399*/
//line goweave.w:6250

finish_line()
active_file.Close()
active_file= tex_file
out_str("\\fin")
finish_line()

if f,err:= os.OpenFile(scn_file_name,os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0666);
err!=nil{
fatal("! Cannot open section file ",scn_file_name)

}else{
scn_file= f
}
active_file= scn_file
/*409:*/
//line goweave.w:6604
section_print(name_root)

/*:409*/
//line goweave.w:6265

finish_line()
active_file.Close()
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
check_complete()
}

/*:384*//*398:*/
//line goweave.w:6408


func unbucket(d int32){


for c:= 100+128;c>=0;c--{
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

/*:398*//*408:*/
//line goweave.w:6588


func section_print(p int32){
if p!=-1{
section_print(name_dir[p].llink)
out_str("\\I")

init_stack()
make_output(section_token(p))
footnote(cite_flag)
footnote(0)
finish_line()
section_print(name_dir[p].rlink)
}
}

/*:408*//*410:*/
//line goweave.w:6609

func print_stats(){
fmt.Println("\nMemory usage statistics:\n")

fmt.Println("%v names",len(name_dir))
fmt.Println("Parsing:")
fmt.Println("Sorting:")
fmt.Println("%v levels ",max_sort_ptr)
}

/*:410*/
