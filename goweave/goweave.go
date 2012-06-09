/*2:*/
//line goweave.w:64

package main

import(
/*14:*/
//line common.w:119

"io"
"bytes"

/*:14*//*17:*/
//line common.w:161

"bufio"

/*:17*//*21:*/
//line common.w:201

"unicode"

/*:21*//*28:*/
//line common.w:355

"fmt"

/*:28*//*35:*/
//line common.w:482

"os"
"strings"

/*:35*/
//line goweave.w:68

)

const(
/*1:*/
//line goweave.w:61

banner= "This is GOWEAVE (Version 0.1)\n"

/*:1*//*5:*/
//line goweave.w:106

max_names= 4000

line_length= 80

max_scraps= 2000
stack_size= 400

/*:5*//*99:*/
//line goweave.w:150

normal= 0
roman= 1
wildcard= 2
typewriter= 3
custom= 5
alfop= 22
else_like= 26
new_like= 42
for_like= 45
if_like= 47
raw_ubin= 49
const_like= 50
raw_int= 51
int_like= 52
case_like= 53
struct_like= 55
typedef_like rune= 56

/*:99*//*104:*/
//line goweave.w:212

cite_flag= 10240
file_flag= 3*cite_flag
def_flag= 2*cite_flag

/*:104*//*118:*/
//line goweave.w:479

ignore rune= 00
verbatim rune= 02
begin_short_comment rune= 03
begin_comment rune= '\t'
underline rune= '\n'
noop rune= 0177
xref_roman rune= 0203
xref_wildcard rune= 0204
xref_typewriter rune= 0205
TeX_string rune= 0206
ord rune= 0207
join rune= 0210
thin_space rune= 0211
math_break rune= 0212
line_break rune= 0213
big_line_break rune= 0214
no_line_break rune= 0215
pseudo_semi rune= 0216
trace rune= 0222
format_code rune= 0225
begin_code rune= 0227
section_name rune= 0230
new_section rune= 0231

/*:118*//*127:*/
//line goweave.w:667

constant= 0200
str= 0201
identifier= 0202

/*:127*//*184:*/
//line goweave.w:1905

exp rune= 1
unop rune= 2
binop rune= 3
ubinop rune= 4

cast rune= 5
lbrace rune= 7
rbrace rune= 8
decl_head rune= 9
comma rune= 10
lpar rune= 11
rpar rune= 12
prelangle rune= 13
prerangle rune= 14
langle rune= 15
base rune= 19
decl rune= 20
struct_head rune= 21
stmt rune= 23
function rune= 24
fn_decl rune= 25
semi rune= 27
colon rune= 28
tag rune= 29
if_head rune= 30
else_head rune= 31
if_clause rune= 32
insert rune= 37
section_scrap rune= 38
dead rune= 39
new_exp rune= 60

/*:184*//*188:*/
//line goweave.w:2045

math_rel= 0206
big_cancel= 0210
cancel= 0211
indent= 0212
outdent= 0213
opt= 0214
backup= 0215
break_space= 0216
force= 0217
big_force= 0220

quoted_char= 0222

end_translation= 0223
inserted= 0224
qualifier= 0225

/*:188*//*194:*/
//line goweave.w:2258

id_flag rune= unicode.UpperLower
res_flag rune= 2*id_flag
section_flag rune= 4*id_flag
tok_flag rune= 6*id_flag
inner_tok_flag rune= 8*id_flag

/*:194*//*195:*/
//line goweave.w:2335

no_math rune= 2
yes_math rune= 1
maybe_math rune= 0

/*:195*//*200:*/
//line goweave.w:2485

no_ident_found int32= -2
case_found int32= -1

/*:200*//*246:*/
//line goweave.w:3345

safe_tok_incr= 20
safe_text_incr= 10
safe_scrap_incr= 10

/*:246*//*267:*/
//line goweave.w:3909

inner mode= 0
outer mode= 1

/*:267*//*275:*/
//line goweave.w:3983

res_word= 0201
section_code= 0200

/*:275*//*319:*/
//line goweave.w:4942

max_sorts= max_scraps

/*:319*//*325:*/
//line goweave.w:4995

infinity= -1

/*:325*/
//line goweave.w:72

)

/*101:*/
//line goweave.w:198

type xref_info struct{
num int32
xlink int32
}

/*:101*//*191:*/
//line goweave.w:2215

type trans struct{
Trans int32
/*317:*/
//line goweave.w:4936

Head int32

/*:317*/
//line goweave.w:2218

}

type scrap struct{
cat int32
mathness int32
trans_plus trans
}

/*:191*//*268:*/
//line goweave.w:3913

type output_state struct{
end_field int32
tok_field int32
mode_field mode
}
type stack_pointer int32

/*:268*//*318:*/
//line goweave.w:4939

type sort_pointer int32

/*:318*/
//line goweave.w:75

/*100:*/
//line goweave.w:174

var change_exists bool

/*:100*//*102:*/
//line goweave.w:204

var xmem[]xref_info
var xref_switch int32
var section_xref_switch int32

/*:102*//*112:*/
//line goweave.w:329

var tok_mem[]rune
var tok_start[]int32
var max_tok_ptr int
var max_text_ptr int

/*:112*//*120:*/
//line goweave.w:509

var ccode[256]rune

/*:120*//*128:*/
//line goweave.w:672

var cur_section int32
var cur_section_char rune


/*:128*//*143:*/
//line goweave.w:1128

var next_control rune

/*:143*//*153:*/
//line goweave.w:1300

var lhs int32
var rhs int32
var res_wd_end int32

/*:153*//*158:*/
//line goweave.w:1399

var cur_xref int32;
var an_output bool

/*:158*//*162:*/
//line goweave.w:1448

var out_buf[line_length+1]rune
var out_ptr int32
var out_buf_end int32= line_length
var out_line int

/*:162*//*185:*/
//line goweave.w:1938

var cat_name[256]string

/*:185*//*192:*/
//line goweave.w:2227

var scrap_info[max_scraps]scrap
var pp int32
var scrap_base int32
var scrap_ptr int32
var lo_ptr int32
var hi_ptr int32
var max_scr_ptr int32

/*:192*//*197:*/
//line goweave.w:2364

var cur_mathness int32
var init_mathness int32

/*:197*//*249:*/
//line goweave.w:3383

var tracing int32

/*:249*//*270:*/
//line goweave.w:3926

var cur_state output_state

var stack[stack_size]output_state
var stack_ptr stack_pointer
var stack_end stack_pointer= stack_size-1
var max_stack_ptr stack_pointer

/*:270*//*274:*/
//line goweave.w:3980

var cur_name int32= -1

/*:274*//*293:*/
//line goweave.w:4449

var save_line int
var save_place int32
var sec_depth int32
var space_checked bool
var format_visible bool
var doing_format bool= false
var group_found bool= false

/*:293*//*301:*/
//line goweave.w:4641

var this_section int32

/*:301*//*314:*/
//line goweave.w:4900

var bucket[256]int32
var blink[max_names]int32

/*:314*//*321:*/
//line goweave.w:4948

var cur_depth int32
var cur_byte int32
var cur_val int32
var max_sort_ptr int32
var sort_ptr int32

/*:321*//*323:*/
//line goweave.w:4961


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


/*:323*//*333:*/
//line goweave.w:5155

var next_xref int32
var this_xref int32


/*:333*/
//line goweave.w:76


/*:2*//*4:*/
//line goweave.w:86

func main(){
flags['x']= true
flags['f']= true
flags['e']= true
common_init()
/*106:*/
//line goweave.w:220

xmem= append(xmem,xref_info{})
xref_switch= 0
section_xref_switch= 0

/*:106*//*113:*/
//line goweave.w:335

tok_start= append(tok_start,0)
max_tok_ptr= 1
max_text_ptr= 1

/*:113*//*121:*/
//line goweave.w:512

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
/*122:*/
//line goweave.w:559

ccode['0']= trace
ccode['1']= trace
ccode['2']= trace

/*:122*/
//line goweave.w:553


/*:121*//*165:*/
//line goweave.w:1525

out_ptr= 1
out_line= 1
active_file= tex_file
out_buf[out_ptr]= 'c'
fmt.Fprint(active_file,"\\input cwebma")

/*:165*//*169:*/
//line goweave.w:1561

out_buf[0]= '\\'

/*:169*//*186:*/
//line goweave.w:1941

for cat_index:= 0;cat_index<255;cat_index++{
cat_name[cat_index]= "UNKNOWN"
}

cat_name[exp]= "exp"
cat_name[unop]= "unop"
cat_name[binop]= "binop"
cat_name[ubinop]= "ubinop"
cat_name[cast]= "cast"
cat_name[lbrace]= "{"
cat_name[rbrace]= "}"
cat_name[decl_head]= "decl_head"
cat_name[comma]= ","
cat_name[lpar]= "("
cat_name[rpar]= ")"
cat_name[prelangle]= "<"
cat_name[prerangle]= ">"
cat_name[langle]= "\\<"
cat_name[base]= "\\:"
cat_name[decl]= "decl"
cat_name[struct_head]= "struct_head"
cat_name[alfop]= "alfop"
cat_name[stmt]= "stmt"
cat_name[function]= "function"
cat_name[fn_decl]= "fn_decl"
cat_name[else_like]= "else_like"
cat_name[semi]= ";"
cat_name[colon]= ":"
cat_name[tag]= "tag"
cat_name[if_head]= "if_head"
cat_name[else_head]= "else_head"
cat_name[if_clause]= "if()"
cat_name[insert]= "insert"
cat_name[section_scrap]= "section"
cat_name[dead]= "@d"
cat_name[new_like]= "new"
cat_name[for_like]= "for"
cat_name[if_like]= "if"
cat_name[raw_ubin]= "ubinop?"
cat_name[const_like]= "const"
cat_name[raw_int]= "raw"
cat_name[int_like]= "int"
cat_name[case_like]= "case"
cat_name[struct_like]= "struct"
cat_name[typedef_like]= "typedef"
cat_name[new_exp]= "new_exp"
cat_name[0]= "zero"

/*:186*//*193:*/
//line goweave.w:2236

scrap_base= 1
max_scr_ptr= 0
scrap_ptr= 0

/*:193*//*271:*/
//line goweave.w:3934

max_stack_ptr= 0

/*:271*//*322:*/
//line goweave.w:4955

max_sort_ptr= 0

/*:322*/
//line goweave.w:92

if show_banner(){
fmt.Print(banner)
}
/*115:*/
//line goweave.w:371

id_lookup([]rune("and"),alfop)
id_lookup([]rune("and_eq"),alfop)
id_lookup([]rune("bitand"),alfop)
id_lookup([]rune("bitor"),alfop)
id_lookup([]rune("compl"),alfop)
id_lookup([]rune("line"),if_like)
id_lookup([]rune("not"),alfop)
id_lookup([]rune("not_eq"),alfop)
id_lookup([]rune("or"),alfop)
id_lookup([]rune("or_eq"),alfop)
id_lookup([]rune("xor"),alfop)
id_lookup([]rune("xor_eq"),alfop)

id_lookup([]rune("break"),case_like)
id_lookup([]rune("case"),case_like)
id_lookup([]rune("chan"),raw_int)
id_lookup([]rune("const"),const_like)
id_lookup([]rune("continue"),case_like)
id_lookup([]rune("default"),case_like)
id_lookup([]rune("defer"),if_like)
id_lookup([]rune("else"),else_like)
id_lookup([]rune("fallthrough"),case_like)
id_lookup([]rune("for"),for_like)
id_lookup([]rune("func"),fn_decl)
id_lookup([]rune("go"),if_like)
id_lookup([]rune("goto"),case_like)
id_lookup([]rune("if"),if_like)
id_lookup([]rune("import"),if_like)
id_lookup([]rune("interface"),struct_like)
id_lookup([]rune("map"),raw_int)
id_lookup([]rune("package"),custom)
id_lookup([]rune("range"),for_like)
id_lookup([]rune("return"),case_like)
id_lookup([]rune("select"),for_like)
id_lookup([]rune("struct"),struct_like)
id_lookup([]rune("switch"),for_like)
id_lookup([]rune("type"),typedef_like)
id_lookup([]rune("var"),const_like)


id_lookup([]rune("bool"),raw_int)
id_lookup([]rune("byte"),raw_int)
id_lookup([]rune("complex64"),raw_int)
id_lookup([]rune("complex128"),raw_int)
id_lookup([]rune("error"),raw_int)
id_lookup([]rune("float32"),raw_int)
id_lookup([]rune("float64"),raw_int)
id_lookup([]rune("int"),raw_int)
id_lookup([]rune("int8"),raw_int)
id_lookup([]rune("int16"),raw_int)
id_lookup([]rune("int32"),raw_int)
id_lookup([]rune("int64"),raw_int)
id_lookup([]rune("rune"),raw_int)
id_lookup([]rune("string"),raw_int)
id_lookup([]rune("uint"),raw_int)
id_lookup([]rune("uint8"),raw_int)
id_lookup([]rune("uint16"),raw_int)
id_lookup([]rune("uint32"),raw_int)
id_lookup([]rune("uint64"),raw_int)
id_lookup([]rune("uintptr"),raw_int)


id_lookup([]rune("true"),raw_int)
id_lookup([]rune("false"),raw_int)
id_lookup([]rune("iota"),raw_int)


id_lookup([]rune("nil"),custom)


id_lookup([]rune("append"),new_like)
id_lookup([]rune("cap"),new_like)
id_lookup([]rune("close"),new_like)
id_lookup([]rune("complex"),new_like)
id_lookup([]rune("copy"),new_like)
id_lookup([]rune("delete"),new_like)
id_lookup([]rune("imag"),new_like)
id_lookup([]rune("len"),new_like)
id_lookup([]rune("make"),new_like)
id_lookup([]rune("new"),new_like)
id_lookup([]rune("panic"),new_like)
id_lookup([]rune("print"),new_like)
id_lookup([]rune("println"),new_like)
id_lookup([]rune("real"),new_like)
id_lookup([]rune("recover"),new_like)
res_wd_end= int32(len(name_dir))
id_lookup([]rune("TeX"),custom)

/*:115*/
//line goweave.w:96

phase_one()
phase_two()
phase_three()
os.Exit(wrap_up())
}

/*:4*//*7:*/
//line common.w:50

const(
/*11:*/
//line common.w:94

and_and= 04
lt_lt= 020
gt_gt= 021
plus_plus= 013
minus_minus= 01
not_eq= 032
lt_eq= 034
gt_eq= 035
eq_eq= 036
or_or= 037
dot_dot_dot= 016

/*:11*//*32:*/
//line common.w:427

max_sections= 2000



/*:32*//*43:*/
//line common.w:640

hash_size= 353

/*:43*//*57:*/
//line common.w:783

less= 0
equal= 1
greater= 2
prefix= 3
extension= 4

/*:57*//*66:*/
//line common.w:1002

bad_extension= 5

/*:66*//*68:*/
//line common.w:1064

spotless= 0
harmless_message= 1
error_message= 2
fatal_message= 3

/*:68*/
//line common.w:52

)


/*13:*/
//line common.w:113

var buffer[]rune
var loc int= 0
var section_text[]rune
var id[]rune

/*:13*//*18:*/
//line common.w:164

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
//line common.w:432

var section_count int32
var changed_section[max_sections]bool
var change_pending bool

var print_where bool= false

/*:33*//*41:*/
//line common.w:615

type name_info struct{
name[]rune
/*42:*/
//line common.w:629

llink int32

/*:42*//*51:*/
//line common.w:712

ispref bool
rlink int32


/*:51*//*98:*/
//line goweave.w:147

ilk int32

/*:98*//*105:*/
//line goweave.w:217

xref int32

/*:105*/
//line common.w:618

}
type name_index int
var name_dir[]name_info
var name_root int32

/*:41*//*44:*/
//line common.w:644

var hash[hash_size]int32
var h int32

/*:44*//*71:*/
//line common.w:1082

var history int= spotless

/*:71*//*87:*/
//line common.w:1255

var go_file_name string
var tex_file_name string
var idx_file_name string
var scn_file_name string
var flags[128]bool

/*:87*//*95:*/
//line common.w:1397

var go_file io.WriteCloser
var tex_file io.WriteCloser
var idx_file io.WriteCloser
var scn_file io.WriteCloser
var active_file io.WriteCloser

/*:95*/
//line common.w:56

/*8:*/
//line common.w:65
var phase int

/*:8*//*19:*/
//line common.w:183

var change_buffer[]rune

/*:19*/
//line common.w:57


/*:7*//*9:*/
//line common.w:71

func common_init(){
/*45:*/
//line common.w:648

for i,_:= range hash{
hash[i]= -1
}

/*:45*//*52:*/
//line common.w:717

name_root= -1

/*:52*/
//line common.w:73

/*88:*/
//line common.w:1266

flags['b']= true
flags['h']= true
flags['p']= true

/*:88*/
//line common.w:74

/*96:*/
//line common.w:1404

scan_args()
/*340:*/
//line goweave.w:5224

if f,err:= os.OpenFile(tex_file_name,os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0666);
err!=nil{
fatal("! Cannot open output file ",tex_file_name)
}else{
tex_file= f
}


/*:340*/
//line common.w:1406


/*:96*/
//line common.w:75

}


/*:9*//*15:*/
//line common.w:123


func input_ln(fp*bufio.Reader)error{
var prefix bool
var err error
var buf[]byte
var b[]byte
buffer= nil
for buf,prefix,err= fp.ReadLine();
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
//line common.w:193

func prime_the_change_buffer(){
change_buffer= nil
/*22:*/
//line common.w:208

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
//line common.w:196

/*23:*/
//line common.w:235

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
//line common.w:197

/*24:*/
//line common.w:248

{
change_buffer= buffer
buffer= nil
}

/*:24*/
//line common.w:198

}

/*:20*//*25:*/
//line common.w:269

func if_section_start_make_pending(b bool){
for loc= 0;loc<len(buffer)&&unicode.IsSpace(buffer[loc]);loc++{}
if len(buffer)>=2&&buffer[0]=='@'&&(unicode.IsSpace(buffer[1])||buffer[1]=='*'){
change_pending= b
}
}

/*:25*//*26:*/
//line common.w:278

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
//line common.w:301


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
//line common.w:358

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
//line common.w:334

}
/*24:*/
//line common.w:248

{
change_buffer= buffer
buffer= nil
}

/*:24*/
//line common.w:336

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
//line common.w:378

func reset_input(){
loc= 0
file= file[:0]
/*31:*/
//line common.w:397

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
//line common.w:382

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
//line common.w:439

func get_line()bool{
restart:
if changing&&include_depth==change_depth{
/*38:*/
//line common.w:560
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
//line common.w:443

}
if!changing||include_depth> change_depth{
/*37:*/
//line common.w:530
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
//line common.w:446

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
//line common.w:486
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
//line common.w:467

}
return true
}

/*:34*//*39:*/
//line common.w:595

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
//line common.w:655


func id_lookup(
id[]rune,
t int32)int32{
/*47:*/
//line common.w:672

h:= id[0]
for i:= 1;i<len(id);i++{
h= (h+h+id[i])%hash_size
}


/*:47*/
//line common.w:660

/*48:*/
//line common.w:682

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
//line common.w:661

if p==-1{
/*50:*/
//line common.w:700

p= int32(len(name_dir)-1)
name_dir[p].name= append(name_dir[p].name,id...)
init_p(p,t)

/*:50*/
//line common.w:663

}
return p
}

/*:46*//*54:*/
//line common.w:737

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
//line common.w:755

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
//line common.w:771

func print_prefix_name(p int32){
l:= name_dir[p].name[0]
fmt.Print(string(name_dir[p].name[1:]))
if int(l)<len(name_dir[p].name){
fmt.Print("...")
}
}

/*:56*//*58:*/
//line common.w:790


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
//line common.w:828


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
//line common.w:862

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
//line common.w:887


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
//line common.w:910

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
//line common.w:899

/*64:*/
//line common.w:942

if r==-1{
return add_section_name(par,c,name,ispref)
}

/*:64*/
//line common.w:900

/*65:*/
//line common.w:951

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
//line common.w:901

return-1
}

/*:62*//*67:*/
//line common.w:1005

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
//line common.w:1070

func mark_harmless(){
if history==spotless{
history= harmless_message
}
}

/*:69*//*70:*/
//line common.w:1077

func mark_error(){
history= error_message
}

/*:70*//*73:*/
//line common.w:1092


func err_print(s string){
var l int
if len(s)> 0&&s[0]=='!'{
fmt.Printf("\n%s",s)
}else{
fmt.Printf("%s",s)
}
if len(file)> 0&&file[0]!=nil{
/*74:*/
//line common.w:1117

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
//line common.w:1102

}
os.Stdout.Sync()
mark_error()
}

/*:73*//*76:*/
//line common.w:1163

func wrap_up()int{
fmt.Print("\n")
if show_stats(){
print_stats()
}
/*77:*/
//line common.w:1176

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
//line common.w:1169

if history> harmless_message{
return 1
}
return 0
}

/*:76*//*79:*/
//line common.w:1196

func fatal(s string,t string){
if len(s)!=0{
fmt.Print(s)
}
err_print(t)
history= fatal_message
os.Exit(wrap_up())
}

/*:79*//*80:*/
//line common.w:1208

func overflow(t string){
fmt.Printf("\n! Sorry, %s capacity exceeded",t)
fatal("","")
}


/*:80*//*81:*/
//line common.w:1220

func confusion(s string){
fatal("! This can't happen: ",s)

}

/*:81*//*83:*/
//line common.w:1235

func show_banner()bool{
return flags['b']
}

/*:83*//*84:*/
//line common.w:1240

func show_progress()bool{
return flags['p']
}

/*:84*//*85:*/
//line common.w:1245

func show_stats()bool{
return flags['s']
}

/*:85*//*86:*/
//line common.w:1250

func show_happiness()bool{
return flags['h']
}

/*:86*//*90:*/
//line common.w:1286

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
//line common.w:1383

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
//line common.w:1299

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
//line common.w:1334

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
//line common.w:1312

}else if!found_change{
/*92:*/
//line common.w:1350

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
//line common.w:1314

}else if!found_out{
/*93:*/
//line common.w:1364

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
//line common.w:1316

}else{
/*339:*/
//line goweave.w:5216

{
fatal("! Usage: goweave [options] webfile[.w] [{changefile[.ch]|-} [outfile[.tex]]]\n","")

}

/*:339*/
//line common.w:1318

}
}
}
if!found_web{
/*339:*/
//line goweave.w:5216

{
fatal("! Usage: goweave [options] webfile[.w] [{changefile[.ch]|-} [outfile[.tex]]]\n","")

}

/*:339*/
//line common.w:1323

}
}

/*:90*//*97:*/
//line common.w:1411

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
//line goweave.w:118

/*:97*//*107:*/
//line goweave.w:234

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
//line goweave.w:250

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
//line goweave.w:286

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
//line goweave.w:313

func set_file_flag(p int32){
q:= name_dir[p].xref
if xmem[q].num==file_flag{
return
}
append_xref(file_flag)
xmem[len(xmem)-1].xlink= q
name_dir[p].xref= int32(len(xmem)-1)
}

/*:110*//*114:*/
//line goweave.w:341

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

/*:114*//*124:*/
//line goweave.w:573

func skip_limbo(){
for true{
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
/*156:*/
//line goweave.w:1353

{
if get_next()!=identifier{
err_print("! Missing left identifier of @s");

}else{
lhs= id_lookup(id,normal)
if get_next()!=identifier{
err_print("! Missing right identifier of @s");

}else{
rhs= id_lookup(id,normal)
name_dir[lhs].ilk= name_dir[rhs].ilk
}
}
}

/*:156*/
//line goweave.w:596

}
}
}
}

/*:124*//*125:*/
//line goweave.w:609


func skip_TeX()rune{
for true{
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
}
return 0
}

/*:125*//*130:*/
//line goweave.w:682


func get_next()rune{
for true{
if loc>=len(buffer)&&!get_line(){
return new_section
}
c:= buffer[loc]
loc++
nc:= ' '
if loc<len(buffer){
nc= buffer[loc]
}
if unicode.IsDigit(c)||c=='.'{
/*133:*/
//line goweave.w:846
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
}
for loc<len(buffer)&&
(buffer[loc]=='u'||
buffer[loc]=='U'||
buffer[loc]=='l'||
buffer[loc]=='L'||
buffer[loc]=='f'||
buffer[loc]=='F'){
id= append(id,'$')
id= append(id,unicode.ToUpper(buffer[loc]))
loc++
}
return constant
}

/*:133*/
//line goweave.w:696

}else if c=='\''||c=='"'||c=='L'&&
(nc=='\''||nc=='"'){
/*134:*/
//line goweave.w:909
{
delim:= c
section_text= section_text[0:0]

if delim=='\''&&
loc-2<len(buffer)&&loc-2>=0&&buffer[loc-2]=='@'{
section_text= append(section_text,'@')
section_text= append(section_text,'@')
}
section_text= append(section_text,delim)
if loc<len(buffer)&&delim=='L'{
delim= buffer[loc]
loc++
section_text= append(section_text,delim)
}
if delim=='<'{
delim= '>'
}
for true{
if loc>=len(buffer){
if buffer[len(buffer)-1]!='\\'{
err_print("! String didn't end")
loc= len(buffer)
break

}
if!get_line(){
err_print("! Input ended in middle of string")
loc= 0
break;

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

/*:134*/
//line goweave.w:699

}else if unicode.IsLetter(c)||c=='_'{
/*132:*/
//line goweave.w:823
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

/*:132*/
//line goweave.w:701

}else if c=='@'{
/*135:*/
//line goweave.w:965
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
/*136:*/
//line goweave.w:994
{
section_text= section_text[0:0]
/*138:*/
//line goweave.w:1015

for true{
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
/*139:*/
//line goweave.w:1040

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

/*:139*/
//line goweave.w:1029

loc++
if unicode.IsSpace(c){
c= ' '
if len(section_text)> 0&&section_text[len(section_text)-1]==' '{
section_text= section_text[:len(section_text)-1]
}
}
section_text= append(section_text,c)
}

/*:138*/
//line goweave.w:996

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

/*:136*/
//line goweave.w:981

case verbatim:
/*142:*/
//line goweave.w:1103
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

/*:142*/
//line goweave.w:983

case ord:
/*134:*/
//line goweave.w:909
{
delim:= c
section_text= section_text[0:0]

if delim=='\''&&
loc-2<len(buffer)&&loc-2>=0&&buffer[loc-2]=='@'{
section_text= append(section_text,'@')
section_text= append(section_text,'@')
}
section_text= append(section_text,delim)
if loc<len(buffer)&&delim=='L'{
delim= buffer[loc]
loc++
section_text= append(section_text,delim)
}
if delim=='<'{
delim= '>'
}
for true{
if loc>=len(buffer){
if buffer[len(buffer)-1]!='\\'{
err_print("! String didn't end")
loc= len(buffer)
break

}
if!get_line(){
err_print("! Input ended in middle of string")
loc= 0
break;

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

/*:134*/
//line goweave.w:985

default:
return ccode[c]
}
}

/*:135*/
//line goweave.w:703

}else if unicode.IsSpace(c){
continue
}
mistake:
/*131:*/
//line goweave.w:720

switch(c){
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
if nc=='='{
l:= loc
loc++
if l<=len(buffer){
return lt_eq
}
}else if nc=='<'{
l:= loc
loc++
if l<=len(buffer){
return lt_lt
}
}
case'&':
if nc=='&'{
l:= loc
loc++
if l<=len(buffer){
return and_and
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
}

/*:131*/
//line goweave.w:708

return c
}
return 0
}

/*:130*//*141:*/
//line goweave.w:1072

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

/*:141*//*145:*/
//line goweave.w:1134

func phase_one(){
phase= 1
reset_input()
section_count= 0
skip_limbo()
change_exists= false
for!input_has_ended{
/*146:*/
//line goweave.w:1150

{
section_count++
changed_section[section_count]= changing

if loc-1<len(buffer)&&buffer[loc-1]=='*'&&show_progress(){
fmt.Printf("*%d",section_count)
os.Stdout.Sync()
}
/*151:*/
//line goweave.w:1243

for true{
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
break
case xref_roman,xref_wildcard,xref_typewriter,noop,section_name:
loc-= 2
next_control= get_next()
if next_control>=xref_roman&&next_control<=xref_typewriter{
/*152:*/
//line goweave.w:1270

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

/*:152*/
//line goweave.w:1260

new_xref(id_lookup(id,next_control-identifier))
}
break
}
if next_control>=format_code{
break
}
}

/*:151*/
//line goweave.w:1159

/*154:*/
//line goweave.w:1307

for next_control<=format_code{
/*155:*/
//line goweave.w:1318
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

/*:155*/
//line goweave.w:1309

outer_xref()
}

/*:154*/
//line goweave.w:1160

/*157:*/
//line goweave.w:1373

if next_control<=section_name{
if next_control==begin_code{
section_xref_switch= 0
}else{
section_xref_switch= def_flag
if cur_section_char=='('&&cur_section!=-1{
set_file_flag(cur_section)
}
}
for true{
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

/*:157*/
//line goweave.w:1161

if changed_section[section_count]{
change_exists= true
}
}

/*:146*/
//line goweave.w:1142

}
changed_section[section_count]= change_exists

phase= 2
/*161:*/
//line goweave.w:1440
section_check(name_root)

/*:161*/
//line goweave.w:1147

}

/*:145*//*148:*/
//line goweave.w:1188


func Go_xref(spec_ctrl rune){
for next_control<format_code||next_control==spec_ctrl{
if next_control>=identifier&&next_control<=xref_typewriter{
if next_control> identifier{
/*152:*/
//line goweave.w:1270

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

/*:152*/
//line goweave.w:1194

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

/*:148*//*150:*/
//line goweave.w:1216


func outer_xref(){
for next_control<format_code{
if next_control!=begin_comment&&next_control!=begin_short_comment{
Go_xref(ignore)
}else{
is_long_comment:= (next_control==begin_comment)
bal:= copy_comment(is_long_comment,1)
next_control= '|'
for bal> 0{
Go_xref(section_name)
if next_control=='|'{
bal= copy_comment(is_long_comment,bal)
}else{
bal= 0
}
}
}
}
}

/*:150*//*160:*/
//line goweave.w:1407


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

/*:160*//*163:*/
//line goweave.w:1465


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

/*:163*//*164:*/
//line goweave.w:1505


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

/*:164*//*167:*/
//line goweave.w:1540

func out(c rune){
if out_ptr>=out_buf_end{
break_out()
}
out_ptr++
out_buf[out_ptr]= c
}

/*:167*//*168:*/
//line goweave.w:1549


func out_str(s string){
for _,v:= range s{
out(v)
}
}

/*:168*//*171:*/
//line goweave.w:1568


func break_out(){
k:= out_ptr
for true{
if k==0{
/*172:*/
//line goweave.w:1594

{
fmt.Printf("\n! Line had to be broken (output l. %d):\n",out_line)

fmt.Fprint(os.Stdout,string(out_buf[1:out_ptr]))
fmt.Println()
mark_harmless()
flush_buffer(out_ptr-1,true,true)
return
}

/*:172*/
//line goweave.w:1574

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

/*:171*//*173:*/
//line goweave.w:1610

func out_section(n int32){
out_str(fmt.Sprintf("%d",n))
if changed_section[n]{
out_str("\\*")

}
}

/*:173*//*174:*/
//line goweave.w:1622

func out_name(p int32,quote_xalpha bool){
out('{')
for _,v:= range name_dir[p].name{
if(v=='_'||v=='$')&&quote_xalpha{
out('\\')
}


out(v)
}
out('}')
}

/*:174*//*175:*/
//line goweave.w:1649

func copy_limbo(){
for true{
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

/*:175*//*177:*/
//line goweave.w:1701

func copy_TeX()rune{
for true{
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

/*:177*//*178:*/
//line goweave.w:1747

func app_tok(c rune){
tok_mem= append(tok_mem,c)
}

/*:178*//*179:*/
//line goweave.w:1752


func copy_comment(
is_long_comment bool,
bal int)int{
for true{
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
return bal
}
if is_long_comment{
/*180:*/
//line goweave.w:1807

if c=='*'&&loc<len(buffer)&&buffer[loc]=='/'{
loc++
if bal> 1{
err_print("! Missing } in comment")

}
goto done
}

/*:180*/
//line goweave.w:1780

}
if phase==2{
if c> 0177{
app_tok(quoted_char)
}
app_tok(c)
}
/*181:*/
//line goweave.w:1817

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
app_tok(buffer[loc])
}
loc++
}

/*:181*/
//line goweave.w:1788

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
/*182:*/
//line goweave.w:1840

if phase==2{
for bal--;bal>=0;bal--{
app_tok('}')
}
}
return 0

/*:182*/
//line goweave.w:1804

}

/*:179*//*187:*/
//line goweave.w:1992


func print_cat(c int32){
fmt.Print(cat_name[c])
}

/*:187*//*196:*/
//line goweave.w:2340

func big_app2(a rune){
big_app1(a)
big_app1(a+1)
}

func big_app3(a rune){
big_app2(a)
big_app1(a+2)
}

func big_app4(a rune){
big_app3(a)
big_app1(a+3)
}

func app(a rune){
tok_mem= append(tok_mem,a)
}

func app1(a int32){
tok_mem= append(tok_mem,tok_flag+scrap_info[a].trans_plus.Trans)
}

/*:196*//*198:*/
//line goweave.w:2368

func app_str(s string){
for _,v:= range s{
app_tok(v)
}
}

func big_app(a rune){
if a==' '||(a>=big_cancel&&a<=big_force){
if cur_mathness==maybe_math{
init_mathness= no_math
}else if cur_mathness==yes_math{
app_str("{}$")
}
cur_mathness= no_math
}else{
if cur_mathness==maybe_math{
init_mathness= yes_math
}else if cur_mathness==no_math{
app_str("${}")
}
cur_mathness= yes_math
}
app(a)
}

func big_app1(a int32){
switch scrap_info[a].mathness%4{
case no_math:
if cur_mathness==maybe_math{
init_mathness= no_math
}else if(cur_mathness==yes_math){
app_str("{}$")
}
cur_mathness= scrap_info[a].mathness/4
case yes_math:
if cur_mathness==maybe_math{
init_mathness= yes_math
}else if cur_mathness==no_math{
app_str("${}")
}
cur_mathness= scrap_info[a].mathness/4
case maybe_math:
}
app(tok_flag+scrap_info[a].trans_plus.Trans)
}

/*:198*//*201:*/
//line goweave.w:2489

func find_first_ident(p int32)int32{
for j:= tok_start[p];j<tok_start[p+1];j++{
r:= tok_mem[j]%id_flag
switch tok_mem[j]/id_flag{
case 2:
if name_dir[r].ilk==case_like{
return case_found
}
if name_dir[r].ilk!=raw_int{
break
}
fallthrough
case 1:
return j
case 6,8:
if q:= find_first_ident(r);q!=no_ident_found{
return q
}
fallthrough
default:
if tok_mem[j]==inserted{
return no_ident_found
}else if tok_mem[j]==qualifier{
j++
}
}
}
return no_ident_found
}

/*:201*//*202:*/
//line goweave.w:2524


func make_reserved(p int32){
tok_loc:= find_first_ident(scrap_info[p].trans_plus.Trans)
if tok_loc<=case_found{
return
}
tok_value:= tok_mem[tok_loc]
for p<=scrap_ptr{
if scrap_info[p].cat==exp{
if tok_mem[tok_start[scrap_info[p].trans_plus.Trans]]==tok_value{
scrap_info[p].cat= raw_int
tok_mem[tok_start[scrap_info[p].trans_plus.Trans]]= tok_value%id_flag+res_flag
}
}
if p==lo_ptr{
p= hi_ptr
}else{
p++
}
}
name_dir[tok_value%id_flag].ilk= raw_int
tok_mem[tok_loc]= tok_value%id_flag+res_flag
}

/*:202*//*203:*/
//line goweave.w:2558


func make_underlined(p int32){
var tok_loc int32
if tok_loc= find_first_ident(scrap_info[p].trans_plus.Trans);tok_loc<=case_found{
return
}
xref_switch= def_flag
underline_xref(tok_mem[tok_loc]%id_flag)
}

/*:203*//*205:*/
//line goweave.w:2574

func underline_xref(p int32){
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
/*206:*/
//line goweave.w:2603

append_xref(0)
xmem[len(xmem)-1].xlink= name_dir[p].xref
r:= int32(len(xmem)-1)
name_dir[p].xref= r
for xmem[r].xlink!=q{
xmem[r].num= xmem[xmem[r].xlink].num
r= xmem[r].xlink
}
xmem[r].num= m

/*:206*/
//line goweave.w:2593

}

/*:205*//*243:*/
//line goweave.w:3280

func freeze_text(){
tok_start= append(tok_start,int32(len(tok_mem)))
}

/*:243*//*244:*/
//line goweave.w:3285

func reduce(j int32,k int32,c rune,d int32,n int32){
scrap_info[j].cat= c
scrap_info[j].trans_plus.Trans= int32(len(tok_start)-1)
scrap_info[j].mathness= 4*cur_mathness+init_mathness
freeze_text()
if k> 1{
i:= j+k
i1:= j+1
for i<=lo_ptr{
scrap_info[i1].cat= scrap_info[i].cat
scrap_info[i1].trans_plus.Trans= scrap_info[i].trans_plus.Trans
scrap_info[i1].mathness= scrap_info[i].mathness
i++
i1++
}
lo_ptr= lo_ptr-k+1
}
if pp+d<scrap_base{
pp= scrap_base
}else{
pp= pp+d
}
f:= "reduce"
/*250:*/
//line goweave.w:3386

{
if tracing==2{
fmt.Printf("\n%s %d:",f,n)
for k:= scrap_base;k<=lo_ptr;k++{
if k==pp{
fmt.Print("*")
}else{
fmt.Print(" ")
}
if scrap_info[k].mathness%4==yes_math{
fmt.Print("+")
}else if scrap_info[k].mathness%4==no_math{
fmt.Print("-")
}
print_cat(scrap_info[k].cat)
if scrap_info[k].mathness/4==yes_math{
fmt.Print("+")
}else if scrap_info[k].mathness/4==no_math{
fmt.Print("-")
}
}
if hi_ptr<=scrap_ptr{
fmt.Print("...")
}
}
}

/*:250*/
//line goweave.w:3309

pp--
}

/*:244*//*245:*/
//line goweave.w:3316

func squash(j int32,k int32,c rune,d int32,n int32){
if k==1{
scrap_info[j].cat= c
if pp+d<scrap_base{
pp= scrap_base
}else{
pp= pp+d
}
f:= "squash"
/*250:*/
//line goweave.w:3386

{
if tracing==2{
fmt.Printf("\n%s %d:",f,n)
for k:= scrap_base;k<=lo_ptr;k++{
if k==pp{
fmt.Print("*")
}else{
fmt.Print(" ")
}
if scrap_info[k].mathness%4==yes_math{
fmt.Print("+")
}else if scrap_info[k].mathness%4==no_math{
fmt.Print("-")
}
print_cat(scrap_info[k].cat)
if scrap_info[k].mathness/4==yes_math{
fmt.Print("+")
}else if scrap_info[k].mathness/4==no_math{
fmt.Print("-")
}
}
if hi_ptr<=scrap_ptr{
fmt.Print("...")
}
}
}

/*:250*/
//line goweave.w:3326

pp--
return
}
for i:= j;i<j+k;i++{
big_app1(i)
}
reduce(j,k,c,d,n)
}

/*:245*//*251:*/
//line goweave.w:3420


func translate()int32{
pp= scrap_base
lo_ptr= pp-1
hi_ptr= pp
/*254:*/
//line goweave.w:3465

if tracing==2{
fmt.Printf("\nTracing after l. %d:\n",line[include_depth])
mark_harmless()

}

/*:254*/
//line goweave.w:3426

/*247:*/
//line goweave.w:3350

for true{
/*248:*/
//line goweave.w:3364

if lo_ptr<pp+3{
for hi_ptr<=scrap_ptr&&lo_ptr!=pp+3{
lo_ptr++
scrap_info[lo_ptr].cat= scrap_info[hi_ptr].cat
scrap_info[lo_ptr].mathness= scrap_info[hi_ptr].mathness
scrap_info[lo_ptr].trans_plus.Trans= scrap_info[hi_ptr].trans_plus.Trans
hi_ptr++
}
for i:= lo_ptr+1;i<=pp+3;i++{
scrap_info[i].cat= 0
}
}

/*:248*/
//line goweave.w:3352

if pp> lo_ptr{
break
}
init_mathness= maybe_math
cur_mathness= maybe_math
/*199:*/
//line goweave.w:2420
{

if(scrap_info[pp+1].cat==insert){
squash(pp,2,scrap_info[pp].cat,-2,0)
}else if(scrap_info[pp+2].cat==insert){
squash(pp+1,2,scrap_info[pp+1].cat,-1,0)
}else if(scrap_info[pp+3].cat==insert){
squash(pp+2,2,scrap_info[pp+2].cat,0,0)
}else{
switch(scrap_info[pp].cat){
case exp:/*207:*/
//line goweave.w:2619

if(scrap_info[pp+1].cat==lbrace||
scrap_info[pp+1].cat==int_like||
scrap_info[pp+1].cat==decl){
make_underlined(pp)
big_app1(pp)
big_app(indent)
app(indent)
reduce(pp,1,fn_decl,0,1)
}else if scrap_info[pp+1].cat==unop{
squash(pp,2,exp,-2,2)
}else if(scrap_info[pp+1].cat==binop||
scrap_info[pp+1].cat==ubinop)&&
scrap_info[pp+2].cat==exp{
squash(pp,3,exp,-2,3)
}else if scrap_info[pp+1].cat==comma&&
scrap_info[pp+2].cat==exp{
big_app2(pp)
app(opt)
app('9')
big_app1(pp+2)
reduce(pp,3,exp,-2,4)
}else if scrap_info[pp+1].cat==lpar&&
scrap_info[pp+2].cat==rpar&&
scrap_info[pp+3].cat==colon{
squash(pp+3,1,base,0,5)
}else if scrap_info[pp+1].cat==cast&&
scrap_info[pp+2].cat==colon{
squash(pp+2,1,base,0,5)
}else if scrap_info[pp+1].cat==semi{
squash(pp,2,stmt,-1,6)
}else if scrap_info[pp+1].cat==colon{
make_underlined(pp)
squash(pp,2,tag,-1,7)
}else if scrap_info[pp+1].cat==rbrace{
squash(pp,1,stmt,-1,8)
}else if scrap_info[pp+1].cat==lpar&&
scrap_info[pp+2].cat==rpar&&
(scrap_info[pp+3].cat==const_like||
scrap_info[pp+3].cat==case_like){
big_app1(pp+2)
big_app(' ')
big_app1(pp+3)
reduce(pp+2,2,rpar,0,9)
}else if scrap_info[pp+1].cat==cast&&
(scrap_info[pp+2].cat==const_like||
scrap_info[pp+2].cat==case_like){
big_app1(pp+1)
big_app(' ')
big_app1(pp+2)
reduce(pp+1,2,cast,0,9)
}else if scrap_info[pp+1].cat==exp||
scrap_info[pp+1].cat==cast{
squash(pp,2,exp,-2,10)
}

/*:207*/
//line goweave.w:2430

case lpar:/*208:*/
//line goweave.w:2675

if(scrap_info[pp+1].cat==exp||
scrap_info[pp+1].cat==ubinop)&&
scrap_info[pp+2].cat==rpar{
squash(pp,3,exp,-2,11)
}else if scrap_info[pp+1].cat==rpar{
big_app1(pp)
app('\\')
app(',')
big_app1(pp+1)

reduce(pp,2,exp,-2,12)
}else if(scrap_info[pp+1].cat==decl_head||
scrap_info[pp+1].cat==int_like||
scrap_info[pp+1].cat==cast)&&
scrap_info[pp+2].cat==rpar{
squash(pp,3,cast,-2,13)
}else if(scrap_info[pp+1].cat==decl_head||
scrap_info[pp+1].cat==int_like||
scrap_info[pp+1].cat==exp)&&
scrap_info[pp+2].cat==comma{
big_app3(pp)
app(opt)
app('9')
reduce(pp,3,lpar,-1,14)
}else if scrap_info[pp+1].cat==stmt||
scrap_info[pp+1].cat==decl{
big_app2(pp)
big_app(' ')
reduce(pp,2,lpar,-1,15)
}

/*:208*/
//line goweave.w:2431

case unop:/*209:*/
//line goweave.w:2707

if scrap_info[pp+1].cat==exp||
scrap_info[pp+1].cat==int_like{
squash(pp,2,exp,-2,16)
}

/*:209*/
//line goweave.w:2432

case ubinop:/*210:*/
//line goweave.w:2713

if scrap_info[pp+1].cat==cast&&
scrap_info[pp+2].cat==rpar{
big_app('{')
big_app1(pp)
big_app('}')
big_app1(pp+1)
reduce(pp,2,cast,-2,17)
}else if scrap_info[pp+1].cat==exp||
scrap_info[pp+1].cat==int_like{
big_app('{')
big_app1(pp)
big_app('}')
big_app1(pp+1)
reduce(pp,2,scrap_info[pp+1].cat,-2,18)
}else if scrap_info[pp+1].cat==binop{
big_app(math_rel)
big_app1(pp)
big_app('{')
big_app1(pp+1)
big_app('}')
big_app('}')
reduce(pp,2,binop,-1,19)
}

/*:210*/
//line goweave.w:2433

case binop:/*211:*/
//line goweave.w:2738

if scrap_info[pp+1].cat==binop{
big_app(math_rel)
big_app('{')
big_app1(pp)
big_app('}')
big_app('{')
big_app1(pp+1)
big_app('}')
big_app('}')
reduce(pp,2,binop,-1,20)
}

/*:211*/
//line goweave.w:2434

case cast:/*212:*/
//line goweave.w:2751

if scrap_info[pp+1].cat==lpar{
squash(pp,2,lpar,-1,21)
}else if scrap_info[pp+1].cat==exp{
big_app1(pp)
big_app(' ')
big_app1(pp+1)
reduce(pp,2,exp,-2,21)
}else if scrap_info[pp+1].cat==semi{
squash(pp,1,exp,-2,22)
}

/*:212*/
//line goweave.w:2435

case int_like:/*213:*/
//line goweave.w:2763

if scrap_info[pp+1].cat==int_like||
scrap_info[pp+1].cat==struct_like{
big_app1(pp)
big_app(' ')
big_app1(pp+1)
reduce(pp,2,scrap_info[pp+1].cat,-2,25)
}else if scrap_info[pp+1].cat==exp&&
(scrap_info[pp+2].cat==raw_int||
scrap_info[pp+2].cat==struct_like){
squash(pp,2,int_like,-2,26)
}else if scrap_info[pp+1].cat==exp||
scrap_info[pp+1].cat==ubinop||
scrap_info[pp+1].cat==colon{
big_app1(pp)
big_app(' ')
reduce(pp,1,decl_head,-1,27)
}else if scrap_info[pp+1].cat==semi||
scrap_info[pp+1].cat==binop{
squash(pp,1,decl_head,0,28)
}

/*:213*/
//line goweave.w:2436

case decl_head:/*214:*/
//line goweave.w:2785

if scrap_info[pp+1].cat==comma{
big_app2(pp)
big_app(' ')
reduce(pp,2,decl_head,-1,33)
}else if scrap_info[pp+1].cat==ubinop{
big_app1(pp)
big_app('{')
big_app1(pp+1)
big_app('}')
reduce(pp,2,decl_head,-1,34)
}else if scrap_info[pp+1].cat==exp&&
scrap_info[pp+2].cat!=lpar&&
scrap_info[pp+2].cat!=exp&&
scrap_info[pp+2].cat!=cast{
make_underlined(pp+1)
squash(pp,2,decl_head,-1,35)
}else if(scrap_info[pp+1].cat==binop||
scrap_info[pp+1].cat==colon)&&
scrap_info[pp+2].cat==exp&&
(scrap_info[pp+3].cat==comma||
scrap_info[pp+3].cat==semi||
scrap_info[pp+3].cat==rpar){
squash(pp,3,decl_head,-1,36)
}else if scrap_info[pp+1].cat==cast{
squash(pp,2,decl_head,-1,37)
}else if scrap_info[pp+1].cat==lbrace||
scrap_info[pp+1].cat==int_like||
scrap_info[pp+1].cat==decl{
big_app1(pp)
big_app(indent)
app(indent)
reduce(pp,1,fn_decl,0,38)
}else if scrap_info[pp+1].cat==semi{
squash(pp,2,decl,-1,39)
}

/*:214*/
//line goweave.w:2437

case decl:/*215:*/
//line goweave.w:2822

if scrap_info[pp+1].cat==decl{
big_app1(pp)
big_app(force)
big_app1(pp+1)
reduce(pp,2,decl,-1,40)
}else if scrap_info[pp+1].cat==stmt||
scrap_info[pp+1].cat==function{
big_app1(pp)
big_app(big_force)
big_app1(pp+1)
reduce(pp,2,scrap_info[pp+1].cat,-1,41)
}

/*:215*/
//line goweave.w:2438

case base:/*216:*/
//line goweave.w:2836

if scrap_info[pp+1].cat==int_like||
scrap_info[pp+1].cat==exp{
if scrap_info[pp+2].cat==comma{
big_app1(pp)
big_app(' ')
big_app2(pp+1)
app(opt)
app('9')
reduce(pp,3,base,0,42)
}else if scrap_info[pp+2].cat==lbrace{
big_app1(pp)
big_app(' ')
big_app1(pp+1)
big_app(' ')
big_app1(pp+2);
reduce(pp,3,lbrace,-2,43)
}
}

/*:216*/
//line goweave.w:2439

case struct_like:/*217:*/
//line goweave.w:2856

if scrap_info[pp+1].cat==lbrace{
big_app1(pp)
big_app(' ')
big_app1(pp+1)
reduce(pp,2,struct_head,0,44)
}else if scrap_info[pp+1].cat==exp||
scrap_info[pp+1].cat==int_like{
if scrap_info[pp+2].cat==lbrace||
scrap_info[pp+2].cat==semi{
make_underlined(pp+1)
make_reserved(pp+1)
big_app1(pp)
big_app(' ')
big_app1(pp+1)
if scrap_info[pp+2].cat==semi{
reduce(pp,2,decl_head,0,45)
}else{
big_app(' ')
big_app1(pp+2)
reduce(pp,3,struct_head,0,46)
}
}else if scrap_info[pp+2].cat==colon{
squash(pp+2,1,base,2,47)
}else if scrap_info[pp+2].cat!=base{
big_app1(pp)
big_app(' ')
big_app1(pp+1)
reduce(pp,2,int_like,-2,48)
}
}

/*:217*/
//line goweave.w:2440

case struct_head:/*218:*/
//line goweave.w:2888

if(scrap_info[pp+1].cat==decl||
scrap_info[pp+1].cat==stmt||
scrap_info[pp+1].cat==function)&&
scrap_info[pp+2].cat==rbrace{
big_app1(pp)
big_app(indent)
big_app(force)
big_app1(pp+1)
big_app(outdent);big_app(force)
big_app1(pp+2)
reduce(pp,3,int_like,-2,49)
}else if scrap_info[pp+1].cat==rbrace{
big_app1(pp)
app_str("\\,")
big_app1(pp+1)

reduce(pp,2,int_like,-2,50)
}

/*:218*/
//line goweave.w:2441

case fn_decl:/*219:*/
//line goweave.w:2908

if scrap_info[pp+1].cat==decl{
big_app1(pp)
big_app(force)
big_app1(pp+1)
reduce(pp,2,fn_decl,0,51)
}else if scrap_info[pp+1].cat==stmt{
big_app1(pp)
app(outdent)
app(outdent)
big_app(force)
big_app1(pp+1)
reduce(pp,2,function,-1,52)
}

/*:219*/
//line goweave.w:2442

case function:/*220:*/
//line goweave.w:2923

if scrap_info[pp+1].cat==function||
scrap_info[pp+1].cat==decl||
scrap_info[pp+1].cat==stmt{
big_app1(pp)
big_app(big_force)
big_app1(pp+1)
reduce(pp,2,scrap_info[pp+1].cat,-1,53)
}

/*:220*/
//line goweave.w:2443

case lbrace:/*221:*/
//line goweave.w:2933

if scrap_info[pp+1].cat==rbrace{
big_app1(pp)
app('\\')
app(',')
big_app1(pp+1)

reduce(pp,2,stmt,-1,54)
}else if(scrap_info[pp+1].cat==stmt||
scrap_info[pp+1].cat==decl||
scrap_info[pp+1].cat==function)&&
scrap_info[pp+2].cat==rbrace{
big_app(force)
big_app1(pp)
big_app(indent)
big_app(force)
big_app1(pp+1)
big_app(force)
big_app(backup)
big_app1(pp+2)
big_app(outdent)
big_app(force)
reduce(pp,3,stmt,-1,55)
}else if scrap_info[pp+1].cat==exp{
if scrap_info[pp+2].cat==rbrace{
squash(pp,3,exp,-2,56)
}else if scrap_info[pp+2].cat==comma&&
scrap_info[pp+3].cat==rbrace{
squash(pp,4,exp,-2,56)
}
}

/*:221*/
//line goweave.w:2444

case if_like:/*222:*/
//line goweave.w:2965

if scrap_info[pp+1].cat==exp{
big_app1(pp)
big_app(' ')
big_app1(pp+1)
reduce(pp,2,if_clause,0,57);
}

/*:222*/
//line goweave.w:2445

case else_like:/*223:*/
//line goweave.w:2973

if scrap_info[pp+1].cat==colon{
squash(pp+1,1,base,1,58)
}else if scrap_info[pp+1].cat==lbrace{
squash(pp,1,else_head,0,59)
}else if scrap_info[pp+1].cat==stmt{
big_app(force)
big_app1(pp)
big_app(indent)
big_app(break_space)
big_app1(pp+1)
big_app(outdent)
big_app(force)
reduce(pp,2,stmt,-1,60)
}

/*:223*/
//line goweave.w:2446

case else_head:/*224:*/
//line goweave.w:2989

if scrap_info[pp+1].cat==stmt||
scrap_info[pp+1].cat==exp{
big_app(force)
big_app1(pp)
big_app(break_space)
app(noop);
big_app(cancel)
big_app1(pp+1)
big_app(force)
reduce(pp,2,stmt,-1,61)
}

/*:224*/
//line goweave.w:2447

case if_clause:/*225:*/
//line goweave.w:3002

if scrap_info[pp+1].cat==lbrace{
squash(pp,1,if_head,0,62)
}else if scrap_info[pp+1].cat==stmt{
if scrap_info[pp+2].cat==else_like{
big_app(force)
big_app1(pp)
big_app(indent)
big_app(break_space)
big_app1(pp+1)
big_app(outdent)
big_app(force)
big_app1(pp+2)
if scrap_info[pp+3].cat==if_like{
big_app(' ')
big_app1(pp+3)
reduce(pp,4,if_like,0,63)
}else{
reduce(pp,3,else_like,0,64)
}
}else{
squash(pp,1,else_like,0,65)
}
}

/*:225*/
//line goweave.w:2448

case if_head:/*226:*/
//line goweave.w:3027

if scrap_info[pp+1].cat==stmt||
scrap_info[pp+1].cat==exp{
if scrap_info[pp+2].cat==else_like{
big_app(force)
big_app1(pp)
big_app(break_space)
app(noop)
big_app(cancel)
big_app1(pp+1)
big_app(force)
big_app1(pp+2)
if scrap_info[pp+3].cat==if_like{
big_app(' ')
big_app1(pp+3)
reduce(pp,4,if_like,0,66)
}else{
reduce(pp,3,else_like,0,67)
}
}else{
squash(pp,1,else_head,0,68)
}
}

/*:226*/
//line goweave.w:2449

case case_like:/*227:*/
//line goweave.w:3051

if scrap_info[pp+1].cat==semi{
squash(pp,2,stmt,-1,70)
}else if scrap_info[pp+1].cat==colon{
squash(pp,2,tag,-1,71)
}else if scrap_info[pp+1].cat==exp{
big_app1(pp)
big_app(' ')
big_app1(pp+1)
reduce(pp,2,exp,-2,72)
}

/*:227*/
//line goweave.w:2450

case tag:/*228:*/
//line goweave.w:3063

if scrap_info[pp+1].cat==tag{
big_app1(pp)
big_app(break_space)
big_app1(pp+1)
reduce(pp,2,tag,-1,74)
}else if scrap_info[pp+1].cat==stmt||
scrap_info[pp+1].cat==decl||
scrap_info[pp+1].cat==function{
big_app(force)
big_app(backup)
big_app1(pp)
big_app(break_space)
big_app1(pp+1)
reduce(pp,2,scrap_info[pp+1].cat,-1,75)
}

/*:228*/
//line goweave.w:2451

case stmt:/*229:*/
//line goweave.w:3083

if scrap_info[pp+1].cat==stmt||
scrap_info[pp+1].cat==decl||
scrap_info[pp+1].cat==function{
big_app1(pp)
if scrap_info[pp+1].cat==function{
big_app(big_force)
}else if scrap_info[pp+1].cat==decl{
big_app(big_force)
}else if flags['f']{
big_app(force)
}else{
big_app(break_space)
}
big_app1(pp+1)
reduce(pp,2,scrap_info[pp+1].cat,-1,76)
}

/*:229*/
//line goweave.w:2452

case semi:/*230:*/
//line goweave.w:3101

big_app(' ')
big_app1(pp)
reduce(pp,1,stmt,-1,77)

/*:230*/
//line goweave.w:2453

case section_scrap:/*231:*/
//line goweave.w:3106

if scrap_info[pp+1].cat==semi{
big_app2(pp)
big_app(force)
reduce(pp,2,stmt,-2,81)
}else{
squash(pp,1,exp,-2,82)
}

/*:231*/
//line goweave.w:2454

case insert:/*232:*/
//line goweave.w:3115

if scrap_info[pp+1].cat!=0{
squash(pp,2,scrap_info[pp+1].cat,0,83)
}

/*:232*/
//line goweave.w:2455

case prelangle:/*233:*/
//line goweave.w:3120

init_mathness= yes_math
cur_mathness= yes_math
app('<')
reduce(pp,1,binop,-2,84)

/*:233*/
//line goweave.w:2456

case prerangle:/*234:*/
//line goweave.w:3126

init_mathness= yes_math
cur_mathness= yes_math
app('>')
reduce(pp,1,binop,-2,85)

/*:234*/
//line goweave.w:2457

case langle:/*235:*/
//line goweave.w:3132

if scrap_info[pp+1].cat==prerangle{
big_app1(pp)
app('\\')
app(',')
big_app1(pp+1)

reduce(pp,2,cast,-1,86)
}else if scrap_info[pp+1].cat==decl_head||
scrap_info[pp+1].cat==int_like||
scrap_info[pp+1].cat==exp{
if scrap_info[pp+2].cat==prerangle{
squash(pp,3,cast,-1,87)
}else if scrap_info[pp+2].cat==comma{
big_app3(pp)
app(opt)
app('9')
reduce(pp,3,langle,0,88)
}
}

/*:235*/
//line goweave.w:2458

case new_like:/*236:*/
//line goweave.w:3153

if scrap_info[pp+1].cat==lpar&&
scrap_info[pp+2].cat==exp&&
scrap_info[pp+3].cat==rpar{
squash(pp,4,new_like,0,92)
}else if scrap_info[pp+1].cat==cast{
big_app1(pp)
big_app(' ')
big_app1(pp+1)
reduce(pp,2,exp,-2,93)
}else if scrap_info[pp+1].cat!=lpar{
squash(pp,1,new_exp,0,94)
}

/*:236*/
//line goweave.w:2459

case new_exp:/*237:*/
//line goweave.w:3167

if scrap_info[pp+1].cat==int_like||
scrap_info[pp+1].cat==const_like{
big_app1(pp)
big_app(' ')
big_app1(pp+1)
reduce(pp,2,new_exp,0,95)
}else if scrap_info[pp+1].cat==struct_like&&
(scrap_info[pp+2].cat==exp||
scrap_info[pp+2].cat==int_like){
big_app1(pp)
big_app(' ')
big_app1(pp+1)
big_app(' ')
big_app1(pp+2)
reduce(pp,3,new_exp,0,96)
}else if scrap_info[pp+1].cat==raw_ubin{
big_app1(pp)
big_app('{')
big_app1(pp+1)
big_app('}')
reduce(pp,2,new_exp,0,97)
}else if scrap_info[pp+1].cat==lpar{
squash(pp,1,exp,-2,98)
}else if(scrap_info[pp+1].cat==exp){
big_app1(pp)
big_app(' ')
reduce(pp,1,exp,-2,98)
}else if scrap_info[pp+1].cat!=raw_int&&
scrap_info[pp+1].cat!=struct_like{
squash(pp,1,exp,-2,99)
}

/*:237*/
//line goweave.w:2460

case for_like:/*238:*/
//line goweave.w:3200

if scrap_info[pp+1].cat==exp{
big_app1(pp)
big_app(' ')
big_app1(pp+1)
reduce(pp,2,else_like,-2,102)
}

/*:238*/
//line goweave.w:2461

case raw_ubin:/*239:*/
//line goweave.w:3208

if scrap_info[pp+1].cat==const_like{
big_app2(pp)
app_str("\\ ")
reduce(pp,2,raw_ubin,0,103)

}else{
squash(pp,1,ubinop,-2,104)
}

/*:239*/
//line goweave.w:2462

case const_like:/*240:*/
//line goweave.w:3218

squash(pp,1,int_like,-2,105)

/*:240*/
//line goweave.w:2463

case raw_int:/*241:*/
//line goweave.w:3221

if scrap_info[pp+1].cat==prelangle{
squash(pp+1,1,langle,1,106)
}else if scrap_info[pp+1].cat==cast{
squash(pp,2,raw_int,0,108)
}else if scrap_info[pp+1].cat==lpar{
squash(pp,1,exp,-2,109)
}else if scrap_info[pp+1].cat!=langle{
squash(pp,1,int_like,-3,110)
}

/*:241*/
//line goweave.w:2464

case typedef_like:/*242:*/
//line goweave.w:3232

if(scrap_info[pp+1].cat==int_like||
scrap_info[pp+1].cat==cast)&&
(scrap_info[pp+2].cat==comma||
scrap_info[pp+2].cat==semi){
squash(pp+1,1,exp,-1,115)
}else if scrap_info[pp+1].cat==int_like{
big_app1(pp)
big_app(' ')
big_app1(pp+1)
reduce(pp,2,typedef_like,0,116)
}else if scrap_info[pp+1].cat==exp&&
scrap_info[pp+2].cat!=lpar&&
scrap_info[pp+2].cat!=exp&&
scrap_info[pp+2].cat!=cast{
make_underlined(pp+1)
make_reserved(pp+1)
big_app1(pp)
big_app(' ')
big_app1(pp+1)
reduce(pp,2,typedef_like,0,117)
}else if scrap_info[pp+1].cat==comma{
big_app2(pp)
big_app(' ')
reduce(pp,2,typedef_like,0,118)
}else if scrap_info[pp+1].cat==semi{
squash(pp,2,decl,-1,119)
}else if scrap_info[pp+1].cat==ubinop&&
(scrap_info[pp+2].cat==ubinop||
scrap_info[pp+2].cat==cast){
big_app('{')
big_app1(pp+1)
big_app('}')
big_app1(pp+2)
reduce(pp+1,2,scrap_info[pp+2].cat,0,120)
}

/*:242*/
//line goweave.w:2465

}
}
pp++
}

/*:199*/
//line goweave.w:3358

}

/*:247*/
//line goweave.w:3427

/*252:*/
//line goweave.w:3436
{
/*253:*/
//line goweave.w:3454

if lo_ptr> scrap_base&&tracing==1{
fmt.Printf("\nIrreducible scrap sequence in section %d:",section_count);

mark_harmless()
for j:= scrap_base;j<=lo_ptr;j++{
fmt.Printf(" ")
print_cat(scrap_info[j].cat)
}
}

/*:253*/
//line goweave.w:3437

for j:= scrap_base;j<=lo_ptr;j++{
if j!=scrap_base{
app(' ')
}
if scrap_info[j].mathness%4==yes_math{
app('$')
}
app1(j)
if scrap_info[j].mathness/4==yes_math{
app('$')
}
}
freeze_text()
return int32(len(tok_start)-2)
}

/*:252*/
//line goweave.w:3428

}

/*:251*//*255:*/
//line goweave.w:3487


func Go_parse(spec_ctrl rune){
for next_control<format_code||next_control==spec_ctrl{
/*257:*/
//line goweave.w:3512

switch(next_control){
case section_name:
app(section_flag+cur_section)
app_scrap(section_scrap,maybe_math)
app_scrap(exp,yes_math)
case str,constant,verbatim:
/*259:*/
//line goweave.w:3683

count:= -1
if next_control==constant{
app_str("\\T{")

}else if next_control==str{
count= 20
app_str("\\.{")

}else{
app_str("\\vb{")
}

for i:= 0;i<len(id);{
if count==0{
app_str("}\\)\\.{")
count= 20

}
switch(id[i]){
case' ','\\','#','%','$','^','{','}','~','&','_':
app('\\')











case'@':
if i+1<len(id)&&id[i+1]=='@'{
i++
}else{
err_print("! Double @ should be used in strings")
}

}
app_tok(id[i])
i++
count--
}
app('}')
app_scrap(exp,maybe_math)

/*:259*/
//line goweave.w:3519

case identifier:
app_cur_id(true)
case TeX_string:
/*260:*/
//line goweave.w:3746

app_str("\\hbox{")
for i:= 0;i<len(id);{
if id[i]=='@'{
i++
}
app_tok(id[i])
i++
}
app('}')

/*:260*/
//line goweave.w:3523

case'/','.':
app(next_control)
app_scrap(binop,yes_math)
case'<':
app_str("\\langle")
app_scrap(prelangle,yes_math)

case'>':
app_str("\\rangle")
app_scrap(prerangle,yes_math)

case'=':
app_str("\\K")
app_scrap(binop,yes_math)

case'|':
app_str("\\OR")
app_scrap(binop,yes_math)

case'^':
app_str("\\XOR")
app_scrap(binop,yes_math)

case'%':
app_str("\\MOD")
app_scrap(binop,yes_math)

case'!':
app_str("\\R")
app_scrap(unop,yes_math)

case'~':
app_str("\\CM")
app_scrap(unop,yes_math)

case'+','-':
app(next_control)
app_scrap(ubinop,yes_math)
case'*':
app(next_control)
app_scrap(raw_ubin,yes_math)
case'&':
app_str("\\AND")
app_scrap(raw_ubin,yes_math)

case ignore,xref_roman,xref_wildcard,xref_typewriter,noop:
break;
case'(','[':
app(next_control)
app_scrap(lpar,maybe_math)
case')',']':
app(next_control)
app_scrap(rpar,maybe_math)
case'{':
app_str("\\{")
app_scrap(lbrace,yes_math)

case'}':
app_str("\\}")
app_scrap(rbrace,yes_math)

case',':
app(',')
app_scrap(comma,yes_math)
case';':
app(';')
app_scrap(semi,maybe_math)
case':':
app(':')
app_scrap(colon,no_math)
/*258:*/
//line goweave.w:3632

case not_eq:
app_str("\\I")
app_scrap(binop,yes_math)

case lt_eq:
app_str("\\Z")
app_scrap(binop,yes_math)

case gt_eq:
app_str("\\G")
app_scrap(binop,yes_math)

case eq_eq:
app_str("\\E")
app_scrap(binop,yes_math)

case and_and:
app_str("\\W")
app_scrap(binop,yes_math)

case or_or:
app_str("\\V")
app_scrap(binop,yes_math)

case plus_plus:
app_str("\\PP")
app_scrap(unop,yes_math)

case minus_minus:
app_str("\\MM")
app_scrap(unop,yes_math)

case gt_gt:
app_str("\\GG")
app_scrap(binop,yes_math)

case lt_lt:
app_str("\\LL")
app_scrap(binop,yes_math)

case dot_dot_dot:
app_str("\\,\\ldots\\,")
app_scrap(raw_int,yes_math);



/*:258*/
//line goweave.w:3594

case thin_space:
app_str("\\,")
app_scrap(insert,maybe_math)

case math_break:
app(opt)
app_str("0")
app_scrap(insert,maybe_math)
case line_break:
app(force)
app_scrap(insert,no_math)
case big_line_break:
app(big_force)
app_scrap(insert,no_math)
case no_line_break:
app(big_cancel)
app(noop)
app(break_space)
app(noop)
app(big_cancel)
app_scrap(insert,no_math)
case pseudo_semi:
app_scrap(semi,maybe_math)
case join:
app_str("\\J")
app_scrap(insert,no_math)

default:
app(inserted)
app(next_control)
app_scrap(insert,maybe_math)
}

/*:257*/
//line goweave.w:3491

next_control= get_next()
if next_control=='|'||next_control==begin_comment||
next_control==begin_short_comment{
return
}
}
}

/*:255*//*256:*/
//line goweave.w:3503

func app_scrap(c int32,b int32){
scrap_ptr++
scrap_info[scrap_ptr].cat= c
scrap_info[scrap_ptr].trans_plus.Trans= int32(len(tok_start)-1)
scrap_info[scrap_ptr].mathness= 5*(b)
freeze_text()
}

/*:256*//*262:*/
//line goweave.w:3760

func app_cur_id(scrapping bool){
p:= id_lookup(id,normal)
if name_dir[p].ilk<=custom{
app(id_flag+p)
if scrapping{
a1:= exp
a2:= maybe_math
if name_dir[p].ilk==custom{
a2= yes_math
}
app_scrap(a1,a2)
}

}else{
app(res_flag+p)
if scrapping{
if name_dir[p].ilk==alfop{
app_scrap(ubinop,yes_math)
}else{
app_scrap(name_dir[p].ilk,maybe_math)
}
}
}
}

/*:262*//*263:*/
//line goweave.w:3791

func Go_translate()int32{
save_base:= scrap_base
scrap_base= scrap_ptr+1
Go_parse(section_name)
if next_control!='|'{
err_print("! Missing '|' after Go text")

}
app_tok(cancel)
app_scrap(insert,maybe_math)

p:= translate()
if scrap_ptr> max_scr_ptr{
max_scr_ptr= scrap_ptr
}
scrap_ptr= scrap_base-1
scrap_base= save_base
return p
}

/*:263*//*264:*/
//line goweave.w:3822


func outer_parse(){
for next_control<format_code{
if next_control!=begin_comment&&next_control!=begin_short_comment{
Go_parse(ignore)
}else{
is_long_comment:= (next_control==begin_comment);
app(cancel)
app(inserted)
if is_long_comment{
app_str("\\C{")

}else{
app_str("\\SHC{")
}

bal:= copy_comment(is_long_comment,1)
next_control= ignore
for bal> 0{
p:= int32(len(tok_start)-1)
freeze_text()
q:= Go_translate()
app(tok_flag+p)
if flags['e']{
app_str("\\PB{")

}
app(inner_tok_flag+q)
if flags['e']{
app_tok('}')
}
if next_control=='|'{
bal= copy_comment(is_long_comment,bal)
next_control= ignore
}else{
bal= 0
}
}
app(force)
app_scrap(insert,no_math)

}
}
}

/*:264*//*266:*/
//line goweave.w:3907
type mode int

/*:266*//*269:*/
//line goweave.w:3921
func init_stack(){
stack_ptr= 0
cur_state.mode_field= outer
}

/*:269*//*272:*/
//line goweave.w:3941


func push_level(p int32){
if stack_ptr==stack_end{
overflow("stack")
}
if stack_ptr> 0{
stack[stack_ptr].end_field= cur_state.end_field
stack[stack_ptr].tok_field= cur_state.tok_field
stack[stack_ptr].mode_field= cur_state.mode_field
}
stack_ptr++
if stack_ptr> max_stack_ptr{
max_stack_ptr= stack_ptr
}
cur_state.tok_field= tok_start[p]
cur_state.end_field= tok_start[p+1]
}

/*:272*//*273:*/
//line goweave.w:3964

func pop_level(){
stack_ptr--
cur_state.end_field= stack[stack_ptr].end_field
cur_state.tok_field= stack[stack_ptr].tok_field
cur_state.mode_field= stack[stack_ptr].mode_field
}

/*:273*//*276:*/
//line goweave.w:3987


func get_output()rune{
restart:
for cur_state.tok_field==cur_state.end_field{
pop_level()
}
idx:= cur_state.tok_field
a:= tok_mem[idx]
cur_state.tok_field++
if a>=0400{
cur_name= a%id_flag
switch a/id_flag{
case 2:
return res_word
case 4:
return section_code
case 6:
push_level(a%id_flag)
goto restart
case 8:
push_level(a%id_flag)
cur_state.mode_field= inner
goto restart

default:
return identifier
}
}
return a
}

/*:276*//*277:*/
//line goweave.w:4034


func output_Go(){
save_tok_ptr:= len(tok_mem)
save_text_ptr:= len(tok_start)
save_next_control:= next_control
next_control= ignore
p:= Go_translate()
app(inner_tok_flag+p)
if flags['e']{
out_str("\\PB{")
make_output()
out('}')

}else{
make_output()
}
if len(tok_start)> max_text_ptr{
max_text_ptr= len(tok_start)
}
if len(tok_mem)> max_tok_ptr{
max_tok_ptr= len(tok_mem)
}
tok_start= tok_start[:save_text_ptr]
tok_mem= tok_mem[:save_tok_ptr]
next_control= save_next_control
}

/*:277*//*279:*/
//line goweave.w:4064


func make_output(){
var c int
app(end_translation)
freeze_text()
push_level(int32(len(tok_start)-2))
var b rune
for true{
a:= get_output()
reswitch:
switch a{
case end_translation:
return
case identifier,res_word:
/*280:*/
//line goweave.w:4127

out('\\')
if a==identifier{
if name_dir[cur_name].ilk==custom&&!doing_format{
/*281:*/
//line goweave.w:4163

for _,v:= range name_dir[cur_name].name{
if v=='_'{
out('x')
}else if v=='$'{
out('X')
}else{
out(v)
}
}
break

/*:281*/
//line goweave.w:4131

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


}else if name_dir[cur_name].ilk==alfop{
out('X')
/*281:*/
//line goweave.w:4163

for _,v:= range name_dir[cur_name].name{
if v=='_'{
out('x')
}else if v=='$'{
out('X')
}else{
out(v)
}
}
break

/*:281*/
//line goweave.w:4149

}else{
out('&')
}

if is_tiny(cur_name){
if name_dir[cur_name].name[0]=='_'||name_dir[cur_name].name[0]=='$'{
out('\\')
}
out(name_dir[cur_name].name[0])
}else{
out_name(cur_name,true)
}

/*:280*/
//line goweave.w:4079

case section_code:
/*285:*/
//line goweave.w:4270
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
out_section(xmem[cur_xref].num-def_flag)
if phase==3{
cur_xref= xmem[cur_xref].xlink
for xmem[cur_xref].num>=def_flag{
out_str(", ")
out_section(xmem[cur_xref].num-def_flag)
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
/*286:*/
//line goweave.w:4305

scratch:= sprint_section_name(cur_name)
cur_section_name:= cur_name
for i:= 0;i<len(scratch);{
b= scratch[i]
i++
if b=='@'{
/*287:*/
//line goweave.w:4349

ii:= i
i++
if ii<len(scratch)&&scratch[ii]!='@'{
fmt.Print("\n! Illegal control code in section name: <")

print_section_name(cur_section_name)
fmt.Print("> ")
mark_error()
}

/*:287*/
//line goweave.w:4312

}
if an_output{
switch b{
case' ','\\','#','%','$','^',
'{','}','~','&','_':
out('\\')
fallthrough











default:out(b)
}
}else if b!='|'{
out(b)
}else{
var buf[]rune
/*288:*/
//line goweave.w:4366

var delim rune
for true{
if i>=len(scratch){
fmt.Print("\n! Go text in section name didn't end: <");

print_section_name(cur_section_name)
fmt.Print("> ")
mark_error()
break
}
b= scratch[i]
i++
if b=='@'||b=='\\'&&delim!=0{
/*289:*/
//line goweave.w:4397
{
buf= append(buf,b)
buf= append(buf,scratch[i])
i++
}

/*:289*/
//line goweave.w:4380

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

/*:288*/
//line goweave.w:4337

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

/*:286*/
//line goweave.w:4298

if an_output{
out_str(" }")
}
out_str("\\X")
}

/*:285*/
//line goweave.w:4081

case math_rel:
out_str("\\MRL{")

case noop,inserted:
break
case cancel,big_cancel:
c= 0
b= a
for true{
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
/*284:*/
//line goweave.w:4253

for;c> 0;c--{
out_str("\\1")

}
for;c<0;c++{
out_str("\\2")

}

/*:284*/
//line goweave.w:4106

goto reswitch
case indent,outdent,opt,backup,break_space,
force,big_force:
/*282:*/
//line goweave.w:4178

if a<break_space{
if cur_state.mode_field==outer{
out('\\')
out(a-cancel+'0')





if a==opt{
b= get_output();
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
/*283:*/
//line goweave.w:4209
{
b= a
save_mode:= cur_state.mode_field
c= 0
for true{
a= get_output()
if a==inserted{
continue
}
if a==cancel||a==big_cancel{
/*284:*/
//line goweave.w:4253

for;c> 0;c--{
out_str("\\1")

}
for;c<0;c++{
out_str("\\2")

}

/*:284*/
//line goweave.w:4219

goto reswitch
}
if a!=' '&&a<indent||a==backup||a> big_force{
if save_mode==outer{
if out_ptr> 3&&compare_runes(out_buf[out_ptr-3:out_ptr+1],[]rune("\\Y\\B"))==0{
goto reswitch
}
/*284:*/
//line goweave.w:4253

for;c> 0;c--{
out_str("\\1")

}
for;c<0;c++{
out_str("\\2")

}

/*:284*/
//line goweave.w:4227

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

/*:283*/
//line goweave.w:4200

}

/*:282*/
//line goweave.w:4111

case quoted_char:
out(tok_mem[cur_state.tok_field])
cur_state.tok_field++
case qualifier:
default:
out(a)
}
}
}

/*:279*//*291:*/
//line goweave.w:4410

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
/*294:*/
//line goweave.w:4458
{
section_count++
/*295:*/
//line goweave.w:4476

if loc-1<len(buffer)&&buffer[loc-1]!='*'{
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
out_section(section_count)
out_str("}")

/*:295*/
//line goweave.w:4460

save_position()
/*296:*/
//line goweave.w:4514

for true{
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

/*:296*/
//line goweave.w:4462

/*297:*/
//line goweave.w:4544

space_checked= false
for next_control<=format_code{
init_stack()
/*300:*/
//line goweave.w:4607
{
doing_format= true
if buffer[loc-1]=='s'||buffer[loc-1]=='S'{
format_visible= false
}
if!space_checked{
emit_space_if_needed()
save_position()
}
app_str("\\F")

next_control= get_next()
if next_control==identifier{
app(id_flag+id_lookup(id,normal))
app(' ')
app(break_space)
next_control= get_next()
if next_control==identifier{
app(id_flag+id_lookup(id,normal))
app_scrap(exp,maybe_math)
app_scrap(semi,maybe_math)
next_control= get_next()
}
}
if scrap_ptr!=2{
err_print("! Improper format definition")

}
}

/*:300*/
//line goweave.w:4548

outer_parse()
finish_Go(format_visible)
format_visible= true
doing_format= false
}

/*:297*/
//line goweave.w:4463

/*302:*/
//line goweave.w:4644

this_section= -1
if next_control<=section_name{
emit_space_if_needed()
init_stack()
if next_control==begin_code{
next_control= get_next()
}else{
this_section= cur_section
/*303:*/
//line goweave.w:4666

for true{
next_control= get_next()
if next_control!='+'{
break
}
}
if next_control!='='&&next_control!=eq_eq{
err_print("! You need an = sign after the section name")

}else{
next_control= get_next()
}
if out_ptr> 1&&out_buf[out_ptr]=='Y'&&out_buf[out_ptr-1]=='\\'{
app(backup)
}


app(section_flag+this_section)
cur_xref= name_dir[this_section].xref
if xmem[cur_xref].num==file_flag{
cur_xref= xmem[cur_xref].xlink
}
app_str("${}")
if xmem[cur_xref].num!=section_count+def_flag{
app_str("\\mathrel+")
this_section= -1
}
app_str("\\E")

app_str("{}$")
app(force)
app_scrap(dead,no_math)


/*:303*/
//line goweave.w:4654

}
for next_control<=section_name{
outer_parse()
/*304:*/
//line goweave.w:4701

if next_control<section_name{
err_print("! You can't do that in Go text")

next_control= get_next()
}else if next_control==section_name{
app(section_flag+cur_section)
app_scrap(section_scrap,maybe_math)
next_control= get_next()
}

/*:304*/
//line goweave.w:4658

}
finish_Go(true)
}

/*:302*/
//line goweave.w:4464

/*305:*/
//line goweave.w:4715

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

/*:305*/
//line goweave.w:4465

/*309:*/
//line goweave.w:4793

out_str("\\fi")
finish_line()

flush_buffer(0,false,false)

/*:309*/
//line goweave.w:4466

}

/*:294*/
//line goweave.w:4423

}
}

/*:291*//*292:*/
//line goweave.w:4435

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

/*:292*//*299:*/
//line goweave.w:4565



func finish_Go(visible bool){
if visible{
out_str("\\B")
app_tok(force)
app_scrap(insert,no_math)
p:= translate()

app(tok_flag+p)
make_output()
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
if len(tok_start)> max_text_ptr{
max_text_ptr= len(tok_start)
}
if len(tok_mem)> max_tok_ptr{
max_tok_ptr= len(tok_mem)
}
if scrap_ptr> max_scr_ptr{
max_scr_ptr= scrap_ptr
}
tok_mem= tok_mem[:0]
tok_start= tok_start[:1]
scrap_ptr= 0

}

/*:299*//*307:*/
//line goweave.w:4744


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
/*308:*/
//line goweave.w:4771

q:= cur_xref
if xmem[xmem[q].xlink].num> flag{
out('s')
}
for true{
out_section(xmem[cur_xref].num-flag)
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

/*:308*/
//line goweave.w:4763

out('.')
}

/*:307*//*311:*/
//line goweave.w:4807

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
/*313:*/
//line goweave.w:4876
{

var k_section int32= 0
for k_section++;!changed_section[k_section];k_section++{}
out_str("\\ch ")

out_section(k_section)
for k_section<section_count{
for k_section++;!changed_section[k_section];k_section++{}
out_str(", ")
out_section(k_section)
}
out('.')
}

/*:313*/
//line goweave.w:4829

finish_line()
finish_line()
}
out_str("\\inx")
finish_line()

active_file= idx_file
/*315:*/
//line goweave.w:4907
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

/*:315*/
//line goweave.w:4837

/*327:*/
//line goweave.w:5021

sort_ptr= 0
unbucket(1)
for sort_ptr> 0{
cur_depth= scrap_info[sort_ptr].cat
if blink[scrap_info[sort_ptr].trans_plus.Head]==-1||cur_depth==infinity{
/*329:*/
//line goweave.w:5058
{
cur_name= scrap_info[sort_ptr].trans_plus.Head
for true{
out_str("\\I")

/*330:*/
//line goweave.w:5073

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
out_str("\\9");
out_name(cur_name,false)
goto name_done

case typewriter:
out_str("\\.");

fallthrough
case roman:
out_name(cur_name,false)
goto name_done;
case custom:{
out_str("$\\")
for _,v:= range name_dir[cur_name].name{
if v=='_'{
out('x')
}else if v=='$'{
out('X')
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
name_done:/*:330*/
//line goweave.w:5063

/*332:*/
//line goweave.w:5131

/*334:*/
//line goweave.w:5160

this_xref= name_dir[cur_name].xref
cur_xref= 0
for true{
next_xref= xmem[this_xref].xlink
xmem[this_xref].xlink= cur_xref
cur_xref= this_xref
this_xref= next_xref
if this_xref==0{
break
}
}

/*:334*/
//line goweave.w:5132

for true{
out_str(", ")
cur_val= xmem[cur_xref].num
if cur_val<def_flag{
out_section(cur_val)
}else{
out_str("\\[")
out_section(cur_val-def_flag)
out(']')
}

cur_xref= xmem[cur_xref].xlink
if cur_xref==0{
break
}
}
out('.')
finish_line()

/*:332*/
//line goweave.w:5064

cur_name= blink[cur_name]
if cur_name==-1{
break
}
}
sort_ptr--
}

/*:329*/
//line goweave.w:5027

}else{
/*328:*/
//line goweave.w:5033
{
next_name:= scrap_info[sort_ptr].trans_plus.Head
for true{
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

/*:328*/
//line goweave.w:5029

}
}

/*:327*/
//line goweave.w:4838

finish_line()
active_file.Close()
active_file= tex_file
out_str("\\fin")
finish_line()

if f,err:= os.OpenFile(scn_file_name,os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0666);
err!=nil{
fatal("! Cannot open section file ",scn_file_name);

}else{
scn_file= f
}
active_file= scn_file
/*337:*/
//line goweave.w:5197
section_print(name_root)

/*:337*/
//line goweave.w:4853

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

/*:311*//*326:*/
//line goweave.w:4998


func unbucket(d int32){


for c:= 100+128;c>=0;c--{
if bucket[collate[c]]!=-1{

sort_ptr++
if sort_ptr> max_sort_ptr{
max_sort_ptr= sort_ptr
}
if c==0{
scrap_info[sort_ptr].cat= infinity
}else{
scrap_info[sort_ptr].cat= d
}
scrap_info[sort_ptr].trans_plus.Head= bucket[collate[c]]
bucket[collate[c]]= -1
}
}
}

/*:326*//*336:*/
//line goweave.w:5177


func section_print(p int32){
if p!=-1{
section_print(name_dir[p].llink)
out_str("\\I")

tok_mem= tok_mem[:0]
tok_start= tok_start[:1]
scrap_ptr= 0
init_stack()
app(p+section_flag)
make_output()
footnote(cite_flag)
footnote(0)
finish_line()
section_print(name_dir[p].rlink)
}
}

/*:336*//*338:*/
//line goweave.w:5202

func print_stats(){
fmt.Println("\nMemory usage statistics:\n");

fmt.Println("%v names",len(name_dir))
fmt.Println("Parsing:")
fmt.Println("%v scraps",max_scr_ptr)
fmt.Println("%v texts",max_text_ptr)
fmt.Println("%v tokens",max_tok_ptr)
fmt.Println("%v levels",max_stack_ptr)
fmt.Println("Sorting:")
fmt.Println("%v levels ",max_sort_ptr)
}

/*:338*/
