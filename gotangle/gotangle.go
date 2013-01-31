/*2:*/
//line gotangle.w:59

package main

import(
/*13:*/
//line gocommon.w:93

"io"
"bytes"

/*:13*//*16:*/
//line gocommon.w:135

"bufio"

/*:16*//*20:*/
//line gocommon.w:175

"unicode"

/*:20*//*27:*/
//line gocommon.w:330

"fmt"

/*:27*//*34:*/
//line gocommon.w:457

"os"
"strings"

/*:34*/
//line gotangle.w:63

)

const(
/*1:*/
//line gotangle.w:55

banner= "This is GOTANGLE (Version 0.5)\n"

/*:1*//*4:*/
//line gotangle.w:90

max_texts= 2500

/*:4*//*99:*/
//line gotangle.w:191

strs= 02
join= 0177

/*:99*//*105:*/
//line gotangle.w:284

section_number= 0211
identifier= 0212

/*:105*//*110:*/
//line gotangle.w:380

normal= 0
num_or_id= 1
post_slash= 2
unbreakable= 3
verbatim= 4

/*:110*//*124:*/
//line gotangle.w:648

ignore rune= 0
ord rune= 0302
control_text rune= 0303
format_code rune= 0306
definition rune= 0307
begin_code rune= 0310
section_name rune= 0311
new_section rune= 0312

/*:124*//*130:*/
//line gotangle.w:789

constant= 03

/*:130*//*143:*/
//line gotangle.w:1122

macro= 0

/*:143*//*146:*/
//line gotangle.w:1165

line_number= 0214

/*:146*/
//line gotangle.w:67

)


/*91:*/
//line gotangle.w:105

type text struct{
token[]rune
text_link int32
}

/*:91*//*100:*/
//line gotangle.w:213

type output_state struct{
byte_field[]rune
name_field int32
repl_field int32
section_field int32
}

/*:100*/
//line gotangle.w:71

/*92:*/
//line gotangle.w:111

var text_info[]text
var tok_mem[]rune

/*:92*//*97:*/
//line gotangle.w:164

var last_unnamed int32

/*:97*//*101:*/
//line gotangle.w:221

var cur_state output_state

var stack[]output_state

/*:101*//*106:*/
//line gotangle.w:288

var cur_val rune

/*:106*//*111:*/
//line gotangle.w:387

var out_state rune

/*:111*//*114:*/
//line gotangle.w:413

var output_files[]int32
var cur_section_name_char rune
var output_file_name string

/*:114*//*125:*/
//line gotangle.w:658

var ccode[256]rune

/*:125*//*128:*/
//line gotangle.w:741

var comment_continues bool= false

/*:128*//*131:*/
//line gotangle.w:792

var cur_section_name int32
var no_where bool

/*:131*//*144:*/
//line gotangle.w:1125

var cur_text int32
var next_control rune

/*:144*/
//line gotangle.w:72


/*:2*//*3:*/
//line gotangle.w:78

func main(){
common_init()
/*98:*/
//line gotangle.w:167

last_unnamed= 0
text_info= append(text_info,text{})
text_info[0].text_link= 0

/*:98*//*126:*/
//line gotangle.w:661
{
for c:=0;c<len(ccode);c++{
ccode[c]= ignore
}
ccode[' ']= new_section
ccode['\t']= new_section
ccode['\n']= new_section
ccode['\v']= new_section
ccode['\r']= new_section
ccode['\f']= new_section
ccode['*']= new_section
ccode['@']= '@'
ccode['=']= strs
ccode['d']= definition
ccode['D']= definition
ccode['f']= format_code
ccode['F']= format_code
ccode['s']= format_code
ccode['S']= format_code
ccode['c']= begin_code
ccode['C']= begin_code
ccode['p']= begin_code
ccode['P']= begin_code
ccode['^']= control_text
ccode[':']= control_text
ccode['.']= control_text
ccode['t']= control_text
ccode['T']= control_text
ccode['q']= control_text
ccode['Q']= control_text
ccode['&']= join
ccode['<']= section_name
ccode['(']= section_name
ccode['\'']= ord
}

/*:126*/
//line gotangle.w:81

if show_banner(){
fmt.Print(banner)
}
phase_one()
phase_two()
os.Exit(wrap_up())
}

/*:3*//*6:*/
//line gocommon.w:25

const(
/*10:*/
//line gocommon.w:63

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

/*:10*//*31:*/
//line gocommon.w:402

max_sections= 2000



/*:31*//*42:*/
//line gocommon.w:614

hash_size= 353

/*:42*//*55:*/
//line gocommon.w:757

less= 0
equal= 1
greater= 2
prefix= 3
extension= 4

/*:55*//*63:*/
//line gocommon.w:978

bad_extension= 5

/*:63*//*65:*/
//line gocommon.w:1041

spotless= 0
harmless_message= 1
error_message= 2
fatal_message= 3

/*:65*/
//line gocommon.w:27

)

/*12:*/
//line gocommon.w:87

var buffer[]rune
var loc int= 0
var section_text[]rune
var id[]rune

/*:12*//*17:*/
//line gocommon.w:138

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

/*:17*//*32:*/
//line gocommon.w:407

var section_count int32
var changed_section[max_sections]bool
var change_pending bool

var print_where bool= false

/*:32*//*40:*/
//line gocommon.w:589

type name_info struct{
name[]rune
/*41:*/
//line gocommon.w:603

llink int32

/*:41*//*50:*/
//line gocommon.w:684

ispref bool
rlink int32


/*:50*//*93:*/
//line gotangle.w:118

equiv int32

/*:93*/
//line gocommon.w:592

}
type name_index int
var name_dir[]name_info
var name_root int32

/*:40*//*43:*/
//line gocommon.w:617

var hash[hash_size]int32
var h int32

/*:43*//*68:*/
//line gocommon.w:1061

var history int= spotless

/*:68*//*80:*/
//line gocommon.w:1218

var go_file_name string
var tex_file_name string
var idx_file_name string
var scn_file_name string
var flags[128]bool

/*:80*//*87:*/
//line gocommon.w:1360

var go_file io.WriteCloser
var tex_file io.WriteCloser
var idx_file io.WriteCloser
var scn_file io.WriteCloser
var active_file io.WriteCloser

/*:87*/
//line gocommon.w:30

/*7:*/
//line gocommon.w:39
var phase int

/*:7*//*18:*/
//line gocommon.w:157

var change_buffer[]rune

/*:18*/
//line gocommon.w:31


/*:6*//*8:*/
//line gocommon.w:45

func common_init(){
/*44:*/
//line gocommon.w:621

for i,_:=range hash{
hash[i]= -1
}

/*:44*//*51:*/
//line gocommon.w:689

name_root= -1

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
/*163:*/
//line gotangle.w:1514

var err error
if go_file,err= os.OpenFile(go_file_name,os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0666);err!=nil{
fatal("! Cannot open output file ",go_file_name)

}

/*:163*/
//line gocommon.w:1369


/*:88*/
//line gocommon.w:49

}


/*:8*//*14:*/
//line gocommon.w:98


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

/*:14*//*19:*/
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

/*:19*//*24:*/
//line gocommon.w:243

func if_section_start_make_pending(b bool){
for loc= 0;loc<len(buffer)&&unicode.IsSpace(buffer[loc]);loc++{}
if len(buffer)>=2&&buffer[0]=='@'&&(unicode.IsSpace(buffer[1])||buffer[1]=='*'){
change_pending= b
}
}

/*:24*//*25:*/
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

/*:25*//*26:*/
//line gocommon.w:276


func check_change(){
n:=0
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

/*:26*//*29:*/
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

/*:29*//*33:*/
//line gocommon.w:415

func get_line()bool{
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

/*:37*/
//line gocommon.w:419

}
if!changing||include_depth> change_depth{
/*36:*/
//line gocommon.w:505
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

include_depth++
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
goto restart
}
temp_file_name:=os.Getenv("GOWEBINPUTS")
if len(temp_file_name)!=0{

for _,fn:=range strings.Split(temp_file_name,":"){
file_name[include_depth]= fn+"/"+file_name[include_depth]
if f,err:=os.Open(file_name[include_depth]);err==nil{
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

/*:35*/
//line gocommon.w:443

}
return true
}

/*:33*//*38:*/
//line gocommon.w:570

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

/*:38*//*45:*/
//line gocommon.w:628


func id_lookup(
id[]rune,
t int32)int32{
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
p:=int32(len(name_dir))
name_dir= append(name_dir,name_info{})
name_dir[p].llink= -1
init_node(p)
name_dir[p].llink= hash[h]
hash[h]= p
}

/*:47*/
//line gocommon.w:634

if p==-1{
/*49:*/
//line gocommon.w:672

p= int32(len(name_dir)-1)
name_dir[p].name= append(name_dir[p].name,id...)
/*96:*/
//line gotangle.w:145



/*:96*/
//line gocommon.w:675


/*:49*/
//line gocommon.w:636

}
return p
}

/*:45*//*52:*/
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
fmt.Print("...")
}
}

/*:52*//*53:*/
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

/*:53*//*54:*/
//line gocommon.w:745

func print_prefix_name(p int32){
l:=name_dir[p].name[0]
fmt.Print(string(name_dir[p].name[1:]))
if int(l)<len(name_dir[p].name){
fmt.Print("...")
}
}

/*:54*//*56:*/
//line gocommon.w:765


func web_strcmp(
j[]rune,
k[]rune)int{
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

/*:56*//*57:*/
//line gocommon.w:803


func add_section_name(
par int32,
c int,
name[]rune,
ispref bool)int32{
p:=int32(len(name_dir))
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

/*:57*//*58:*/
//line gocommon.w:838

func extend_section_name(
p int32,
text[]rune,
ispref bool){
q:=p+1
for name_dir[q].llink!=-1{
q= name_dir[q].llink
}
np:=int32(len(name_dir))
name_dir[q].llink= np
name_dir= append(name_dir,name_info{})
name_dir[np].llink= -1
init_node(np)
name_dir[np].name= append(name_dir[np].name,int32(len(text)))
name_dir[np].name= append(name_dir[np].name,text...)
name_dir[np].ispref= ispref

}

/*:58*//*59:*/
//line gocommon.w:863


func section_lookup(
name[]rune,
ispref bool)int32{
c:=less
p:=name_root
var q int32= -1
var r int32= -1
var par int32= -1
name_len:=len(name)
/*60:*/
//line gocommon.w:886

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

/*:60*/
//line gocommon.w:875

/*61:*/
//line gocommon.w:918

if r==-1{
return add_section_name(par,c,name,ispref)
}

/*:61*/
//line gocommon.w:876

/*62:*/
//line gocommon.w:927

first,cmp:=section_name_cmp(name,r)
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

/*:62*/
//line gocommon.w:877

return-1
}

/*:59*//*64:*/
//line gocommon.w:982

func section_name_cmp(
name[]rune,
r int32)(int,int){
q:=r+1
var ispref bool
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

/*:64*//*66:*/
//line gocommon.w:1048

func mark_harmless(){
if history==spotless{
history= harmless_message
}
}

/*:66*//*67:*/
//line gocommon.w:1056

func mark_error(){
history= error_message
}

/*:67*//*69:*/
//line gocommon.w:1071


func err_print(s string){
var l int
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

/*:70*/
//line gocommon.w:1081

}
os.Stdout.Sync()
mark_error()
}

/*:69*//*72:*/
//line gocommon.w:1142

func wrap_up()int{
fmt.Print("\n")
if show_stats(){
print_stats()
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
}

/*:73*/
//line gocommon.w:1148

if history> harmless_message{
return 1
}
return 0
}

/*:72*//*74:*/
//line gocommon.w:1175

func fatal(s string,t string){
if len(s)!=0{
fmt.Print(s)
}
err_print(t)
history= fatal_message
os.Exit(wrap_up())
}

/*:74*//*76:*/
//line gocommon.w:1195

func show_banner()bool{
return flags['b']
}

/*:76*//*77:*/
//line gocommon.w:1201

func show_progress()bool{
return flags['p']
}

/*:77*//*78:*/
//line gocommon.w:1207

func show_stats()bool{
return flags['s']
}

/*:78*//*79:*/
//line gocommon.w:1213

func show_happiness()bool{
return flags['h']
}

/*:79*//*82:*/
//line gocommon.w:1249

func scan_args(){
dot_pos:=-1
name_pos:=0
found_web:=false
found_change:=false
found_out:=false

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
arg= arg[:dot_pos]
}
alt_file_name= fmt.Sprintf("%s.web",arg)
tex_file_name= fmt.Sprintf("%s.tex",arg[name_pos:])
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
if flags['x']{
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
/*164:*/
//line gotangle.w:1521

{
fatal("! Usage: gotangle [options] webfile[.w] [{changefile[.ch]|-} [outfile[.go]]]\n","")

}

/*:164*/
//line gocommon.w:1281

}
}
}
if!found_web{
/*164:*/
//line gotangle.w:1521

{
fatal("! Usage: gotangle [options] webfile[.w] [{changefile[.ch]|-} [outfile[.go]]]\n","")

}

/*:164*/
//line gocommon.w:1286

}
}

/*:82*//*89:*/
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

/*:89*//*94:*/
//line gotangle.w:124

func names_match(
p int32,
id[]rune,
t int32)bool{
if len(name_dir[p].name)!=len(id){
return false
}
return compare_runes(id,name_dir[p].name)==0
}

/*:94*//*95:*/
//line gotangle.w:137

func init_node(node int32){
name_dir[node].equiv= -1
}

/*:95*//*103:*/
//line gotangle.w:243


func push_level(p int32){
stack= append(stack,cur_state)
cur_state.name_field= p
cur_state.repl_field= name_dir[p].equiv
cur_state.byte_field= text_info[cur_state.repl_field].token
cur_state.section_field= 0
}

/*:103*//*104:*/
//line gotangle.w:257


func pop_level(){
if text_info[cur_state.repl_field].text_link<max_texts{
cur_state.repl_field= text_info[cur_state.repl_field].text_link
cur_state.byte_field= text_info[cur_state.repl_field].token
return
}

if len(stack)> 0{
cur_state= stack[len(stack)-1]
stack= stack[:len(stack)-1]
}
}

/*:104*//*107:*/
//line gotangle.w:292


func get_output(){
restart:
if len(stack)==0{
return
}
if len(cur_state.byte_field)==0{
cur_val= -cur_state.section_field
pop_level()
if cur_val==0{
goto restart
}
out_char(section_number)
return
}
a:=cur_state.byte_field[0]
cur_state.byte_field= cur_state.byte_field[1:]
if out_state==verbatim&&a!=strs&&a!=constant&&a!='\n'{
fmt.Fprintf(go_file,"%c",a)
}else if a<unicode.UpperLower{
out_char(a)
}else{
c:=cur_state.byte_field[0]
cur_state.byte_field= cur_state.byte_field[1:]
switch a%unicode.UpperLower{
case identifier:
cur_val= c
out_char(identifier)
case section_name:
/*108:*/
//line gotangle.w:339

{
if name_dir[c].equiv!=-1{
push_level(c)
}else if a!=0{
fmt.Printf("\n! Not present: <")
print_section_name(c)
err_print(">")

}
goto restart
}

/*:108*/
//line gotangle.w:322

case line_number:
cur_val= c
out_char(line_number)
case section_number:
cur_val= c
if cur_val> 0{
cur_state.section_field= cur_val
}
out_char(section_number)
}
}
}

/*:107*//*112:*/
//line gotangle.w:394


func flush_buffer(){
fmt.Fprintln(go_file)
if line[include_depth]%100==0&&show_progress(){
fmt.Print(".")
if line[include_depth]%500==0{
fmt.Printf("%d",line[include_depth])
}
os.Stdout.Sync()
}
line[include_depth]++
}

/*:112*//*116:*/
//line gotangle.w:435

func phase_two(){
line[include_depth]= 1
/*102:*/
//line gotangle.w:232

cur_state.name_field= 0
cur_state.repl_field= text_info[0].text_link
cur_state.byte_field= text_info[cur_state.repl_field].token
cur_state.section_field= 0
stack= append(stack,output_state{})

/*:102*/
//line gotangle.w:438

if text_info[0].text_link==0&&len(output_files)==0{
fmt.Print("\n! No program text was specified.")
mark_harmless()

}else{
if len(output_files)==0{
if show_progress(){
fmt.Printf("\nWriting the output file (%s):",go_file_name)
}
}else{
if show_progress(){
fmt.Printf("\nWriting the output files: (%s)",go_file_name)

os.Stdout.Sync()
}
if text_info[0].text_link==0{
goto writeloop
}
}
for len(stack)> 0{
get_output()
}
flush_buffer()
writeloop:
/*117:*/
//line gotangle.w:474

for an_output_file:=len(output_files);an_output_file> 0;{
an_output_file--
output_file_name= string(sprint_section_name(output_files[an_output_file]))
if f,err:=os.OpenFile(output_file_name,os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0666);
err!=nil{
fatal("! Cannot open output file:",output_file_name)

}else{
go_file.Close()
go_file= f
}
fmt.Printf("\n(%s)",output_file_name)
os.Stdout.Sync()
line[include_depth]= 1
stack= append(stack,output_state{})
cur_state.name_field= output_files[an_output_file]
cur_state.repl_field= name_dir[cur_state.name_field].equiv
cur_state.byte_field= text_info[cur_state.repl_field].token
for len(stack)> 0{
get_output()
}
flush_buffer()
}

/*:117*/
//line gotangle.w:463

if show_happiness(){
fmt.Print("\nDone.")
}
}
}

/*:116*//*118:*/
//line gotangle.w:504

func out_char(cur_char rune){
switch cur_char{
case'\n':
flush_buffer()
if out_state!=verbatim{
out_state= normal
}
/*120:*/
//line gotangle.w:596

case identifier:
if out_state==num_or_id{
fmt.Fprint(go_file," ")
}
fmt.Fprintf(go_file,"%s",string(name_dir[cur_val].name))
out_state= num_or_id

/*:120*/
//line gotangle.w:512

/*121:*/
//line gotangle.w:604

case section_number:
if cur_val> 0{
fmt.Fprintf(go_file,"/*%d:*/",cur_val)
}else if cur_val<0{
fmt.Fprintf(go_file,"/*:%d*/",-cur_val)
}

/*:121*/
//line gotangle.w:513

/*122:*/
//line gotangle.w:612

case line_number:
fmt.Fprint(go_file,"\n//line ")

line:=cur_val
cur_val= cur_state.byte_field[0]
cur_state.byte_field= cur_state.byte_field[1:]
for _,v:=range name_dir[cur_val].name{
if v=='\\'||v=='"'{
fmt.Fprint(go_file,"\\")
}
fmt.Fprintf(go_file,"%c",v)
}
fmt.Fprintf(go_file,":%d\n",line)

/*:122*/
//line gotangle.w:514

/*119:*/
//line gotangle.w:551

case plus_plus:
fmt.Fprint(go_file,"++")
out_state= normal
case minus_minus:
fmt.Fprint(go_file,"--")
out_state= normal
case gt_gt:
fmt.Fprint(go_file,">>")
out_state= normal
case eq_eq:
fmt.Fprint(go_file,"==")
out_state= normal
case lt_lt:
fmt.Fprint(go_file,"<<")
out_state= normal
case gt_eq:
fmt.Fprint(go_file,">=")
out_state= normal
case lt_eq:
fmt.Fprint(go_file,"<=")
out_state= normal
case not_eq:
fmt.Fprint(go_file,"!=")
out_state= normal
case and_and:
fmt.Fprint(go_file,"&&")
out_state= normal
case or_or:
fmt.Fprint(go_file,"||")
out_state= normal
case dot_dot_dot:
fmt.Fprint(go_file,"...")
out_state= normal
case direct:
fmt.Fprint(go_file,"<-")
out_state= normal
case and_not:
fmt.Fprint(go_file,"&^")
out_state= normal
case col_eq:
fmt.Fprint(go_file,":=")
out_state= normal


/*:119*/
//line gotangle.w:515

case'=','>':
fmt.Fprintf(go_file,"%c ",cur_char)
out_state= normal
case join:
out_state= unbreakable
case constant:
switch out_state{
case verbatim:
out_state= num_or_id
case num_or_id:
fmt.Fprint(go_file," ")
fallthrough
default:
out_state= verbatim
}
case strs:
if out_state==verbatim{
out_state= normal
}else{
out_state= verbatim
}
case'/':
fmt.Fprint(go_file,"/")
out_state= post_slash
case'*':
if out_state==post_slash{
fmt.Fprint(go_file," ")
}
fallthrough
default:
fmt.Fprintf(go_file,"%c",cur_char)
out_state= normal
}
}

/*:118*//*127:*/
//line gotangle.w:700


func skip_ahead()rune{
for true{
if loc>=len(buffer)&&!get_line(){
return new_section
}
for loc<len(buffer)&&buffer[loc]!='@'{
loc++
}
if loc<len(buffer){
loc++
c:=new_section
if loc<len(buffer)&&buffer[loc]<int32(len(ccode)){
c= ccode[buffer[loc]]
}
loc++
if c!=ignore||(loc<=len(buffer)&&buffer[loc-1]=='>'){
return c
}
}
}
return 0
}

/*:127*//*129:*/
//line gotangle.w:745


func skip_comment(is_long_comment bool)bool{
for true{
if loc>=len(buffer){
if is_long_comment{
if get_line(){
comment_continues= true
return comment_continues
}else{
err_print("! Input ended in mid-comment")

comment_continues= false
return comment_continues
}
}else{
comment_continues= false
return comment_continues
}
}
c:=buffer[loc]
loc++
if is_long_comment&&c=='*'&&loc<len(buffer)&&buffer[loc]=='/'{
loc++
comment_continues= false
return comment_continues
}
if c=='@'{
if buffer[loc]<int32(len(ccode))&&ccode[buffer[loc]]==new_section{
err_print("! Section name ended in mid-comment")
loc--

comment_continues= false
return comment_continues
}else{
loc++
}
}
}
return false
}

/*:129*//*132:*/
//line gotangle.w:799


func get_next()rune{
for true{
if loc>=len(buffer){
if!get_line(){
return new_section
}else if print_where&&!no_where{
print_where= false
/*147:*/
//line gotangle.w:1168

tok_mem= append(tok_mem,unicode.UpperLower+line_number)

if changing{
id= []rune(change_file_name)
}else{
id= []rune(file_name[include_depth])
}
if changing{
tok_mem= append(tok_mem,rune(change_line))
}else{
tok_mem= append(tok_mem,rune(line[include_depth]))
}
{
a:=id_lookup(id,0)
tok_mem= append(tok_mem,a)
}

/*:147*/
//line gotangle.w:808

}else{
return'\n'
}
}
c:=buffer[loc]
var nc rune= ' '
if loc+1<len(buffer){
nc= buffer[loc+1]
}
if comment_continues||(c=='/'&&(nc=='*'||nc=='/')){
skip_comment(comment_continues||nc=='*')

if comment_continues{
return'\n'
}else{
continue
}
}
loc++
if unicode.IsDigit(c)||c=='.'{
/*134:*/
//line gotangle.w:862

{
id_first:=loc-1
if buffer[id_first]=='.'&&(loc>=len(buffer)||!unicode.IsDigit(buffer[loc])){
goto mistake
}
if buffer[id_first]=='0'{
if loc<len(buffer)&&(buffer[loc]=='x'||buffer[loc]=='X'){
loc++
for loc<len(buffer)&&xisxdigit(buffer[loc]){
loc++
}
goto found
}
}
for loc<len(buffer)&&unicode.IsDigit(buffer[loc]){
loc++
}
if loc<len(buffer)&&buffer[loc]=='.'{
loc++
for loc<len(buffer)&&unicode.IsDigit(buffer[loc]){
loc++
}
}
if loc<len(buffer)&&(buffer[loc]=='e'||buffer[loc]=='E'){
loc++
if loc<len(buffer)&&(buffer[loc]=='+'||buffer[loc]=='-'){
loc++
}
for loc<len(buffer)&&unicode.IsDigit(buffer[loc]){
loc++
}
}
found:
for loc<len(buffer)&&
(buffer[loc]=='u'||buffer[loc]=='U'||
buffer[loc]=='l'||buffer[loc]=='L'||
buffer[loc]=='f'||buffer[loc]=='F'){
loc++
}
id= buffer[id_first:loc]
return constant
}

/*:134*/
//line gotangle.w:829

}else if c=='\''||c=='"'||c=='`'{
/*135:*/
//line gotangle.w:910

{
delim:=c
section_text= section_text[0:0]
section_text= append(section_text,delim)

for true{
if loc>=len(buffer){
if!get_line(){
err_print("! Input ended in middle of string")
loc= 0
break

}else{
section_text= append(section_text,'\n')
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
return strs
}

/*:135*/
//line gotangle.w:831

}else if unicode.IsLetter(c)||c=='_'{
/*133:*/
//line gotangle.w:847

{
loc--
id_first:=loc
for loc<len(buffer)&&
(unicode.IsLetter(buffer[loc])||
unicode.IsDigit(buffer[loc])||
buffer[loc]=='_'||
buffer[loc]=='$'){
loc++
}
id= buffer[id_first:loc]
return identifier
}

/*:133*/
//line gotangle.w:833

}else if c=='@'{
/*136:*/
//line gotangle.w:950

{
c= ccode[nc]
loc++
switch c{
case ignore:
continue
case control_text:
for c= skip_ahead();c=='@';c= skip_ahead(){}

if buffer[loc-1]!='>'{
err_print("! Double @ should be used in control text")

}
continue
case section_name:
cur_section_name_char= buffer[loc-1]
/*138:*/
//line gotangle.w:1010

{
section_text= section_text[0:0]
/*140:*/
//line gotangle.w:1032

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
/*141:*/
//line gotangle.w:1057

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
loc++
}

/*:141*/
//line gotangle.w:1046

loc++
if unicode.IsSpace(c){
c= ' '
if len(section_text)> 0&&section_text[len(section_text)-1]==' '{
section_text= section_text[:len(section_text)-1]
}
}
section_text= append(section_text,c)
}

/*:140*/
//line gotangle.w:1013

if len(section_text)> 3&&
compare_runes(section_text[len(section_text)-3:],[]rune("..."))==0{
cur_section_name= section_lookup(section_text[0:len(section_text)-3],
true)
}else{
cur_section_name= section_lookup(section_text,false)
}
if cur_section_name_char=='('{
/*115:*/
//line gotangle.w:419

{
an_output_file:=0
for;an_output_file<len(output_files);an_output_file++{
if output_files[an_output_file]==cur_section_name{
break
}
}
if an_output_file==len(output_files){
output_files= append(output_files,cur_section_name)
}
}

/*:115*/
//line gotangle.w:1023

}
return section_name
}

/*:138*/
//line gotangle.w:967

case strs:
/*142:*/
//line gotangle.w:1091
{
id_first:=loc
loc++
for loc<len(buffer)&&loc+1<len(buffer)&&(buffer[loc]!='@'||buffer[loc+1]!='>'){
loc++
}
if loc>=len(buffer){
err_print("! Verbatim string didn't end")
}

id= buffer[id_first:loc]
loc+= 2
return strs
}

/*:142*/
//line gotangle.w:969

case ord:
/*137:*/
//line gotangle.w:983

if buffer[loc]=='\\'{
loc++
if buffer[loc]=='\''{
loc++
}
}
for buffer[loc]!='\''{
if buffer[loc]=='@'{
if buffer[loc+1]!='@'{
err_print("! Double @ should be used in ASCII constant")

}else{
loc++
}
}
loc++
if loc>=len(buffer){
err_print("! String didn't end")
loc= len(buffer)-1
break

}
}
loc++
return ord

/*:137*/
//line gotangle.w:971

default:
return c
}
}

/*:136*/
//line gotangle.w:835

}else if unicode.IsSpace(c){
continue
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
//line gotangle.w:97


/*:90*/
//line gotangle.w:840

return c
}
return 0
}


/*:132*//*145:*/
//line gotangle.w:1130


func scan_repl(t rune){
var a int32
if t==section_name{
/*147:*/
//line gotangle.w:1168

tok_mem= append(tok_mem,unicode.UpperLower+line_number)

if changing{
id= []rune(change_file_name)
}else{
id= []rune(file_name[include_depth])
}
if changing{
tok_mem= append(tok_mem,rune(change_line))
}else{
tok_mem= append(tok_mem,rune(line[include_depth]))
}
{
a:=id_lookup(id,0)
tok_mem= append(tok_mem,a)
}

/*:147*/
//line gotangle.w:1135

}
for true{
a= get_next()
switch a{
/*148:*/
//line gotangle.w:1186

case identifier:
a= id_lookup(id,0)
tok_mem= append(tok_mem,unicode.UpperLower+identifier)
tok_mem= append(tok_mem,a)
case section_name:
if t!=section_name{
goto done
}else{
/*149:*/
//line gotangle.w:1216
{
try_loc:=loc
for try_loc<len(buffer)&&buffer[try_loc]==' '{
try_loc++
}
if try_loc<len(buffer)&&buffer[try_loc]=='+'{
try_loc++
}
for try_loc<len(buffer)&&buffer[try_loc]==' '{
try_loc++
}
if try_loc<len(buffer)&&buffer[try_loc]=='='{
err_print("! Missing `@ ' before a named section")

}


}

/*:149*/
//line gotangle.w:1195

tok_mem= append(tok_mem,unicode.UpperLower+section_name)
a= cur_section_name
tok_mem= append(tok_mem,a)
/*147:*/
//line gotangle.w:1168

tok_mem= append(tok_mem,unicode.UpperLower+line_number)

if changing{
id= []rune(change_file_name)
}else{
id= []rune(file_name[include_depth])
}
if changing{
tok_mem= append(tok_mem,rune(change_line))
}else{
tok_mem= append(tok_mem,rune(line[include_depth]))
}
{
a:=id_lookup(id,0)
tok_mem= append(tok_mem,a)
}

/*:147*/
//line gotangle.w:1199

}
case constant,strs:
/*150:*/
//line gotangle.w:1235

tok_mem= append(tok_mem,a)
for i:=0;i<len(id);{
if id[i]=='@'{
if id[i+1]=='@'{
i++
}else{
err_print("! Double @ should be used in string")

}
}
tok_mem= append(tok_mem,id[i])
i++
}
tok_mem= append(tok_mem,a)

/*:150*/
//line gotangle.w:1202

case ord:
/*151:*/
//line gotangle.w:1251

{
c:=id[0]
if c=='\\'{
id= id[1:]
c= id[0]
if c>='0'&&c<='7'{
c-= '0'
if id[1]>='0'&&id[1]<='7'{
id= id[1:]
c= 8*c+id[0]-'0'
if id[1]>='0'&&id[1]<='7'&&c<32{
id= id[1:]
c= 8*c+id[0]-'0'
}
}
}else{
switch c{
case't':
c= '\t'
case'n':
c= '\n'
case'b':
c= '\b'
case'f':
c= '\f'
case'v':
c= '\v'
case'r':
c= '\r'
case'a':
c= '\a'
case'?':
c= '?'
case'x':
if unicode.IsDigit(id[1]){
id= id[1:]
c= id[0]-'0'
}else if xisxdigit(id[1])&&
unicode.IsLower(id[1]){
id= id[1:]
c= unicode.ToUpper(id[0])-'A'+10
}
if unicode.IsDigit(id[1]){
id= id[1:]
c= 16*c+id[0]-'0'
}else if xisxdigit(id[1])&&
unicode.IsLower(id[1]){
id= id[1:]
c= 16*c+unicode.ToUpper(id[0])-'A'+10
}
case'\\':c= '\\'
case'\'':c= '\''
case'"':c= '"'
default:
err_print("! Unrecognized escape sequence")

}
}
}

tok_mem= append(tok_mem,constant)
if c>=100{
tok_mem= append(tok_mem,'0'+c/100)
}
if c>=10{
tok_mem= append(tok_mem,'0'+(c/10)%10)
}
tok_mem= append(tok_mem,'0'+c%10)
tok_mem= append(tok_mem,constant)
}

/*:151*/
//line gotangle.w:1204

case definition,format_code,begin_code:
if t!=section_name{
goto done
}else{
err_print("! @d, @f and @c are ignored in Go text")
continue

}
case new_section:
goto done

/*:148*/
//line gotangle.w:1143

case')':
tok_mem= append(tok_mem,a)
if t==macro{
tok_mem= append(tok_mem,' ')
}
default:
tok_mem= append(tok_mem,a)
}
}
done:
next_control= a
cur_text= int32(len(text_info))
text_info= append(text_info,text{})
text_info[cur_text].token= tok_mem
tok_mem= nil
}

/*:145*//*153:*/
//line gotangle.w:1334

func scan_section(){
var p int32= 0
var q int32= 0
var a int32= 0
section_count++
no_where= true
if loc<len(buffer)&&buffer[loc-1]=='*'&&show_progress(){
fmt.Printf("*%d",section_count)
os.Stdout.Sync()
}
next_control= 0
for true{
/*154:*/
//line gotangle.w:1375

for next_control<definition{

if next_control= skip_ahead();next_control==section_name{
loc-= 2
next_control= get_next()
}
}

/*:154*/
//line gotangle.w:1348

if next_control==definition{
/*155:*/
//line gotangle.w:1384

{

for next_control= get_next();next_control=='\n';next_control= get_next(){}
if next_control!=identifier{
err_print("! Definition flushed, must start with identifier")

continue
}
a= id_lookup(id,0)
tok_mem= append(tok_mem,unicode.UpperLower+identifier)
tok_mem= append(tok_mem,a)

if loc<len(buffer)&&
buffer[loc]!='('{
tok_mem= append(tok_mem,strs)
tok_mem= append(tok_mem,' ')
tok_mem= append(tok_mem,strs)
}
scan_repl(macro)
text_info[cur_text].text_link= 0
}

/*:155*/
//line gotangle.w:1350

continue
}
if next_control==begin_code{
p= -1
break
}
if next_control==section_name{
p= cur_section_name
/*156:*/
//line gotangle.w:1415


for next_control= get_next();next_control=='+';next_control= get_next(){}
if next_control!='='&&next_control!=eq_eq{
continue
}

/*:156*/
//line gotangle.w:1359

break
}
return
}
no_where= false
print_where= false
/*157:*/
//line gotangle.w:1422

/*158:*/
//line gotangle.w:1427

tok_mem= append(tok_mem,unicode.UpperLower+section_number)
tok_mem= append(tok_mem,section_count)

/*:158*/
//line gotangle.w:1423

scan_repl(section_name)
/*159:*/
//line gotangle.w:1431

if p==-1{
text_info[last_unnamed].text_link= cur_text
last_unnamed= cur_text
}else if name_dir[p].equiv==-1{
name_dir[p].equiv= cur_text

}else{
q= name_dir[p].equiv
for text_info[q].text_link<max_texts{
q= text_info[q].text_link
}
text_info[q].text_link= cur_text
}
text_info[cur_text].text_link= max_texts


/*:159*/
//line gotangle.w:1425


/*:157*/
//line gotangle.w:1366

}

/*:153*//*160:*/
//line gotangle.w:1449

func phase_one(){
phase= 1
section_count= 0
reset_input()
skip_limbo()
for!input_has_ended{
scan_section()
}
check_complete()
phase= 2
}

/*:160*//*161:*/
//line gotangle.w:1465

func skip_limbo(){
for true{
if loc>=len(buffer)&&!get_line(){
return
}
for loc<len(buffer)&&buffer[loc]!='@'{
loc++
}
if loc++;loc<len(buffer){
c:=buffer[loc]
loc++
cc:=ignore
if c<int32(len(ccode)){
cc= ccode[c]
}
if cc==new_section{
break
}
switch cc{
case format_code,'@':
case control_text:
if c=='q'||c=='Q'{
for c= skip_ahead();c=='@';c= skip_ahead(){}
if buffer[loc-1]!='>'{
err_print("! Double @ should be used in control text")

}
break
}
fallthrough
default:
err_print("! Double @ should be used in limbo")

}
}
}
}

/*:161*//*162:*/
//line gotangle.w:1505

func print_stats(){
fmt.Print("\nMemory usage statistics:\n")
fmt.Printf("%v names\n",len(name_dir))
fmt.Printf("%v replacement texts\n",len(text_info))
}

/*:162*/
