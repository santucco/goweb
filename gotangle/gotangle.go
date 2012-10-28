/*2:*/
//line gotangle.w:58

package main

import(
/*14:*/
//line common.w:125

"io"
"bytes"

/*:14*//*17:*/
//line common.w:167

"bufio"

/*:17*//*21:*/
//line common.w:207

"unicode"

/*:21*//*28:*/
//line common.w:361

"fmt"

/*:28*//*35:*/
//line common.w:488

"os"
"strings"

/*:35*/
//line gotangle.w:62

)

const(
/*1:*/
//line gotangle.w:55

banner= "This is GOTANGLE (Version 0.1)\n"

/*:1*//*5:*/
//line gotangle.w:95

max_texts= 2500

/*:5*//*106:*/
//line gotangle.w:197

strs= 02
join= 0177

/*:106*//*112:*/
//line gotangle.w:291

section_number= 0211
identifier= 0212

/*:112*//*118:*/
//line gotangle.w:390

normal= 0
num_or_id= 1
post_slash= 2
unbreakable= 3
verbatim= 4

/*:118*//*134:*/
//line gotangle.w:657

ignore rune= 0
ord rune= 0302
control_text rune= 0303
format_code rune= 0306
definition rune= 0307
begin_code rune= 0310
section_name rune= 0311
new_section rune= 0312

/*:134*//*140:*/
//line gotangle.w:797

constant= 03

/*:140*//*154:*/
//line gotangle.w:1131

macro= 0

/*:154*//*157:*/
//line gotangle.w:1173

line_number= 0214

/*:157*/
//line gotangle.w:66

)


/*99:*/
//line gotangle.w:110

type text struct{
token[]rune
text_link int32
}
type text_pointer int

/*:99*//*107:*/
//line gotangle.w:219

type output_state struct{
byte_field[]rune
name_field int32
repl_field int32
section_field int32
}
type stack_pointer int

/*:107*/
//line gotangle.w:70

/*100:*/
//line gotangle.w:117

var text_info[]text
var tok_mem[]rune

/*:100*//*104:*/
//line gotangle.w:170

var last_unnamed int32

/*:104*//*108:*/
//line gotangle.w:228

var cur_state output_state

var stack[]output_state

/*:108*//*113:*/
//line gotangle.w:295

var cur_val rune

/*:113*//*119:*/
//line gotangle.w:397

var out_state rune

/*:119*//*122:*/
//line gotangle.w:423

var output_files[]int32
var cur_section_name_char rune
var output_file_name string

/*:122*//*135:*/
//line gotangle.w:667

var ccode[256]rune

/*:135*//*138:*/
//line gotangle.w:750

var comment_continues bool= false

/*:138*//*141:*/
//line gotangle.w:800

var cur_section_name int32
var no_where bool

/*:141*//*155:*/
//line gotangle.w:1134

var cur_text int32
var next_control rune

/*:155*/
//line gotangle.w:71


/*:2*//*4:*/
//line gotangle.w:83

func main(){
common_init()
/*105:*/
//line gotangle.w:173

last_unnamed= 0
text_info= append(text_info,text{})
text_info[0].text_link= 0

/*:105*//*136:*/
//line gotangle.w:670
{
for c:= 0;c<len(ccode);c++{
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

/*:136*/
//line gotangle.w:86

if show_banner(){
fmt.Print(banner)
}
phase_one()
phase_two()
os.Exit(wrap_up())
}

/*:4*//*7:*/
//line common.w:50

const(
/*11:*/
//line common.w:95

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
//line common.w:433

max_sections= 2000



/*:32*//*43:*/
//line common.w:646

hash_size= 353

/*:43*//*57:*/
//line common.w:789

less= 0
equal= 1
greater= 2
prefix= 3
extension= 4

/*:57*//*66:*/
//line common.w:1008

bad_extension= 5

/*:66*//*68:*/
//line common.w:1070

spotless= 0
harmless_message= 1
error_message= 2
fatal_message= 3

/*:68*/
//line common.w:52

)



/*13:*/
//line common.w:119

var buffer[]rune
var loc int= 0
var section_text[]rune
var id[]rune

/*:13*//*18:*/
//line common.w:170

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
//line common.w:438

var section_count int32
var changed_section[max_sections]bool
var change_pending bool

var print_where bool= false

/*:33*//*41:*/
//line common.w:621

type name_info struct{
name[]rune
/*42:*/
//line common.w:635

llink int32

/*:42*//*51:*/
//line common.w:718

ispref bool
rlink int32


/*:51*//*101:*/
//line gotangle.w:124

equiv int32

/*:101*/
//line common.w:624

}
type name_index int
var name_dir[]name_info
var name_root int32

/*:41*//*44:*/
//line common.w:650

var hash[hash_size]int32
var h int32

/*:44*//*71:*/
//line common.w:1088

var history int= spotless

/*:71*//*87:*/
//line common.w:1261

var go_file_name string
var tex_file_name string
var idx_file_name string
var scn_file_name string
var flags[128]bool

/*:87*//*95:*/
//line common.w:1403

var go_file io.WriteCloser
var tex_file io.WriteCloser
var idx_file io.WriteCloser
var scn_file io.WriteCloser
var active_file io.WriteCloser

/*:95*/
//line common.w:57

/*8:*/
//line common.w:66
var phase int

/*:8*//*19:*/
//line common.w:189

var change_buffer[]rune

/*:19*/
//line common.w:58


/*:7*//*9:*/
//line common.w:72

func common_init(){
/*45:*/
//line common.w:654

for i,_:= range hash{
hash[i]= -1
}

/*:45*//*52:*/
//line common.w:723

name_root= -1

/*:52*/
//line common.w:74

/*88:*/
//line common.w:1272

flags['b']= true
flags['h']= true
flags['p']= true

/*:88*/
//line common.w:75

/*96:*/
//line common.w:1410

scan_args()
/*174:*/
//line gotangle.w:1520

var err error
if go_file,err= os.OpenFile(go_file_name,os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0666);err!=nil{
fatal("! Cannot open output file ",go_file_name)

}

/*:174*/
//line common.w:1412


/*:96*/
//line common.w:76

}


/*:9*//*15:*/
//line common.w:129


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
//line common.w:199

func prime_the_change_buffer(){
change_buffer= nil
/*22:*/
//line common.w:214

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
//line common.w:202

/*23:*/
//line common.w:241

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
//line common.w:203

/*24:*/
//line common.w:254

{
change_buffer= buffer
buffer= nil
}

/*:24*/
//line common.w:204

}

/*:20*//*25:*/
//line common.w:275

func if_section_start_make_pending(b bool){
for loc= 0;loc<len(buffer)&&unicode.IsSpace(buffer[loc]);loc++{}
if len(buffer)>=2&&buffer[0]=='@'&&(unicode.IsSpace(buffer[1])||buffer[1]=='*'){
change_pending= b
}
}

/*:25*//*26:*/
//line common.w:284

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
//line common.w:307


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
//line common.w:364

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
//line common.w:340

}
/*24:*/
//line common.w:254

{
change_buffer= buffer
buffer= nil
}

/*:24*/
//line common.w:342

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
//line common.w:384

func reset_input(){
loc= 0
file= file[:0]
/*31:*/
//line common.w:403

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
//line common.w:388

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
//line common.w:445

func get_line()bool{
restart:
if changing&&include_depth==change_depth{
/*38:*/
//line common.w:566
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
//line common.w:449

}
if!changing||include_depth> change_depth{
/*37:*/
//line common.w:536
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
//line common.w:452

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
//line common.w:492
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
//line common.w:473

}
return true
}

/*:34*//*39:*/
//line common.w:601

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
//line common.w:661


func id_lookup(
id[]rune,
t int32)int32{
/*47:*/
//line common.w:678

h:= id[0]
for i:= 1;i<len(id);i++{
h= (h+h+id[i])%hash_size
}


/*:47*/
//line common.w:666

/*48:*/
//line common.w:688

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
//line common.w:667

if p==-1{
/*50:*/
//line common.w:706

p= int32(len(name_dir)-1)
name_dir[p].name= append(name_dir[p].name,id...)
init_p(p,t)

/*:50*/
//line common.w:669

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
//line common.w:761

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
//line common.w:777

func print_prefix_name(p int32){
l:= name_dir[p].name[0]
fmt.Print(string(name_dir[p].name[1:]))
if int(l)<len(name_dir[p].name){
fmt.Print("...")
}
}

/*:56*//*58:*/
//line common.w:796


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
//line common.w:834


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
//line common.w:868

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
//line common.w:893


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
//line common.w:916

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
//line common.w:905

/*64:*/
//line common.w:948

if r==-1{
return add_section_name(par,c,name,ispref)
}

/*:64*/
//line common.w:906

/*65:*/
//line common.w:957

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
//line common.w:907

return-1
}

/*:62*//*67:*/
//line common.w:1011

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
//line common.w:1076

func mark_harmless(){
if history==spotless{
history= harmless_message
}
}

/*:69*//*70:*/
//line common.w:1083

func mark_error(){
history= error_message
}

/*:70*//*73:*/
//line common.w:1098


func err_print(s string){
var l int
if len(s)> 0&&s[0]=='!'{
fmt.Printf("\n%s",s)
}else{
fmt.Printf("%s",s)
}
if len(file)> 0&&file[0]!=nil{
/*74:*/
//line common.w:1123

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
//line common.w:1108

}
os.Stdout.Sync()
mark_error()
}

/*:73*//*76:*/
//line common.w:1169

func wrap_up()int{
fmt.Print("\n")
if show_stats(){
print_stats()
}
/*77:*/
//line common.w:1182

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
//line common.w:1175

if history> harmless_message{
return 1
}
return 0
}

/*:76*//*79:*/
//line common.w:1202

func fatal(s string,t string){
if len(s)!=0{
fmt.Print(s)
}
err_print(t)
history= fatal_message
os.Exit(wrap_up())
}

/*:79*//*80:*/
//line common.w:1214

func overflow(t string){
fmt.Printf("\n! Sorry, %s capacity exceeded",t)
fatal("","")
}


/*:80*//*81:*/
//line common.w:1226

func confusion(s string){
fatal("! This can't happen: ",s)

}

/*:81*//*83:*/
//line common.w:1241

func show_banner()bool{
return flags['b']
}

/*:83*//*84:*/
//line common.w:1246

func show_progress()bool{
return flags['p']
}

/*:84*//*85:*/
//line common.w:1251

func show_stats()bool{
return flags['s']
}

/*:85*//*86:*/
//line common.w:1256

func show_happiness()bool{
return flags['h']
}

/*:86*//*90:*/
//line common.w:1292

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
//line common.w:1389

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
//line common.w:1305

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
//line common.w:1340

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
//line common.w:1318

}else if!found_change{
/*92:*/
//line common.w:1356

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
//line common.w:1320

}else if!found_out{
/*93:*/
//line common.w:1370

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
//line common.w:1322

}else{
/*175:*/
//line gotangle.w:1527

{
fatal("! Usage: gotangle [options] webfile[.w] [{changefile[.ch]|-} [outfile[.go]]]\n","")

}

/*:175*/
//line common.w:1324

}
}
}
if!found_web{
/*175:*/
//line gotangle.w:1527

{
fatal("! Usage: gotangle [options] webfile[.w] [{changefile[.ch]|-} [outfile[.go]]]\n","")

}

/*:175*/
//line common.w:1329

}
}

/*:90*//*97:*/
//line common.w:1417

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

/*:97*//*102:*/
//line gotangle.w:130

func names_match(
p int32,
id[]rune,
t int32)bool{
if len(name_dir[p].name)!=len(id){
return false
}
return compare_runes(id,name_dir[p].name)==0
}

/*:102*//*103:*/
//line gotangle.w:146

func init_node(node int32){
name_dir[node].equiv= -1
}

func init_p(int32,int32){}

/*:103*//*110:*/
//line gotangle.w:250


func push_level(p int32){
stack= append(stack,cur_state)
cur_state.name_field= p
cur_state.repl_field= name_dir[p].equiv
cur_state.byte_field= text_info[cur_state.repl_field].token
cur_state.section_field= 0
}

/*:110*//*111:*/
//line gotangle.w:264


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

/*:111*//*115:*/
//line gotangle.w:302


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
a:= cur_state.byte_field[0]
cur_state.byte_field= cur_state.byte_field[1:]
if out_state==verbatim&&a!=strs&&a!=constant&&a!='\n'{
fmt.Fprintf(go_file,"%c",a)
}else if a<unicode.UpperLower{
out_char(a)
}else{
c:= cur_state.byte_field[0]
cur_state.byte_field= cur_state.byte_field[1:]
switch a%unicode.UpperLower{
case identifier:
cur_val= c
out_char(identifier)
case section_name:
/*116:*/
//line gotangle.w:349

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

/*:116*/
//line gotangle.w:332

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

/*:115*//*120:*/
//line gotangle.w:404


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

/*:120*//*125:*/
//line gotangle.w:445

func phase_two(){
line[include_depth]= 1
/*109:*/
//line gotangle.w:239

cur_state.name_field= 0
cur_state.repl_field= text_info[0].text_link
cur_state.byte_field= text_info[cur_state.repl_field].token
cur_state.section_field= 0
stack= append(stack,output_state{})

/*:109*/
//line gotangle.w:448

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
/*126:*/
//line gotangle.w:484

for an_output_file:= len(output_files);an_output_file> 0;{
an_output_file--
output_file_name= string(sprint_section_name(output_files[an_output_file]))
if f,err:= os.OpenFile(output_file_name,os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0666);err!=nil{
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

/*:126*/
//line gotangle.w:473

if show_happiness(){
fmt.Print("\nDone.")
}
}
}

/*:125*//*128:*/
//line gotangle.w:513

func out_char(cur_char rune){
switch(cur_char){
case'\n':
flush_buffer()
if out_state!=verbatim{
out_state= normal
}
/*130:*/
//line gotangle.w:605

case identifier:
if out_state==num_or_id{
fmt.Fprint(go_file," ")
}
fmt.Fprintf(go_file,"%s",string(name_dir[cur_val].name))
out_state= num_or_id

/*:130*/
//line gotangle.w:521

/*131:*/
//line gotangle.w:613

case section_number:
if cur_val> 0{
fmt.Fprintf(go_file,"/*%d:*/",cur_val)
}else if cur_val<0{
fmt.Fprintf(go_file,"/*:%d*/",-cur_val)
}

/*:131*/
//line gotangle.w:522

/*132:*/
//line gotangle.w:621

case line_number:
fmt.Fprint(go_file,"\n//line ")

line:= cur_val
cur_val= cur_state.byte_field[0]
cur_state.byte_field= cur_state.byte_field[1:]
for _,v:= range name_dir[cur_val].name{
if v=='\\'||v=='"'{
fmt.Fprint(go_file,"\\")
}
fmt.Fprintf(go_file,"%c",v)
}
fmt.Fprintf(go_file,":%d\n",line)

/*:132*/
//line gotangle.w:523

/*129:*/
//line gotangle.w:560

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


/*:129*/
//line gotangle.w:524

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

/*:128*//*137:*/
//line gotangle.w:709


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
c:= new_section
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

/*:137*//*139:*/
//line gotangle.w:753


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
c:= buffer[loc]
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

/*:139*//*143:*/
//line gotangle.w:807


func get_next()rune{
for true{
if loc>=len(buffer){
if!get_line(){
return new_section
}else if print_where&&!no_where{
print_where= false
/*158:*/
//line gotangle.w:1176

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
a:= id_lookup(id,0)
tok_mem= append(tok_mem,a)
}

/*:158*/
//line gotangle.w:816

}else{
return'\n'
}
}
c:= buffer[loc]
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
/*145:*/
//line gotangle.w:871

{
id_first:= loc-1
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

/*:145*/
//line gotangle.w:837

}else if c=='\''||c=='"'||c=='`'{
/*146:*/
//line gotangle.w:919

{
delim:= c
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
return strs
}

/*:146*/
//line gotangle.w:839

}else if unicode.IsLetter(c)||c=='_'{
/*144:*/
//line gotangle.w:855

{
loc--
id_first:= loc
for loc<len(buffer)&&
(unicode.IsLetter(buffer[loc])||
unicode.IsDigit(buffer[loc])||
buffer[loc]=='_'||
buffer[loc]=='$'){

loc++
}
id= buffer[id_first:loc]
return(identifier)
}

/*:144*/
//line gotangle.w:841

}else if c=='@'{
/*147:*/
//line gotangle.w:959

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
/*149:*/
//line gotangle.w:1019

{
section_text= section_text[0:0]
/*151:*/
//line gotangle.w:1041

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
/*152:*/
//line gotangle.w:1066

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

/*:152*/
//line gotangle.w:1055

loc++
if unicode.IsSpace(c){
c= ' '
if len(section_text)> 0&&section_text[len(section_text)-1]==' '{
section_text= section_text[:len(section_text)-1]
}
}
section_text= append(section_text,c)
}

/*:151*/
//line gotangle.w:1022

if len(section_text)> 3&&
compare_runes(section_text[len(section_text)-3:],[]rune("..."))==0{
cur_section_name= section_lookup(section_text[0:len(section_text)-3],
true)
}else{
cur_section_name= section_lookup(section_text,false)
}
if cur_section_name_char=='('{
/*123:*/
//line gotangle.w:429

{
an_output_file:= 0
for;an_output_file<len(output_files);an_output_file++{
if output_files[an_output_file]==cur_section_name{
break
}
}
if an_output_file==len(output_files){
output_files= append(output_files,cur_section_name)
}
}

/*:123*/
//line gotangle.w:1032

}
return section_name
}

/*:149*/
//line gotangle.w:976

case strs:
/*153:*/
//line gotangle.w:1100
{
id_first:= loc
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

/*:153*/
//line gotangle.w:978

case ord:
/*148:*/
//line gotangle.w:992

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

/*:148*/
//line gotangle.w:980

default:
return c
}
}

/*:147*/
//line gotangle.w:843

}else if unicode.IsSpace(c){
continue
}
mistake:
/*98:*/
//line common.w:1438

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
//line gotangle.w:102


/*:98*/
//line gotangle.w:848

return c
}
return 0
}


/*:143*//*156:*/
//line gotangle.w:1138


func scan_repl(t rune){
var a int32
if t==section_name{
/*158:*/
//line gotangle.w:1176

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
a:= id_lookup(id,0)
tok_mem= append(tok_mem,a)
}

/*:158*/
//line gotangle.w:1143

}
for true{
a= get_next()
switch a{
/*159:*/
//line gotangle.w:1194

case identifier:
a= id_lookup(id,0)
tok_mem= append(tok_mem,unicode.UpperLower+identifier)
tok_mem= append(tok_mem,a)
case section_name:
if t!=section_name{
goto done
}else{
/*160:*/
//line gotangle.w:1224
{
try_loc:= loc
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

/*:160*/
//line gotangle.w:1203

tok_mem= append(tok_mem,unicode.UpperLower+section_name)
a= cur_section_name
tok_mem= append(tok_mem,a)
/*158:*/
//line gotangle.w:1176

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
a:= id_lookup(id,0)
tok_mem= append(tok_mem,a)
}

/*:158*/
//line gotangle.w:1207

}
case constant,strs:
/*161:*/
//line gotangle.w:1243

tok_mem= append(tok_mem,a)
for i:= 0;i<len(id);{
if(id[i]=='@'){
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

/*:161*/
//line gotangle.w:1210

case ord:
/*162:*/
//line gotangle.w:1259

{
c:= id[0]
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

/*:162*/
//line gotangle.w:1212

case definition,format_code,begin_code:
if t!=section_name{
goto done
}else{
err_print("! @d, @f and @c are ignored in C text")
continue

}
case new_section:
goto done

/*:159*/
//line gotangle.w:1151

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

/*:156*//*164:*/
//line gotangle.w:1342

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
/*165:*/
//line gotangle.w:1383

for next_control<definition{

if next_control= skip_ahead();next_control==section_name{
loc-= 2
next_control= get_next()
}
}

/*:165*/
//line gotangle.w:1356

if next_control==definition{
/*166:*/
//line gotangle.w:1392

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

/*:166*/
//line gotangle.w:1358

continue
}
if next_control==begin_code{
p= -1
break
}
if next_control==section_name{
p= cur_section_name
/*167:*/
//line gotangle.w:1423


for next_control= get_next();next_control=='+';next_control= get_next(){}
if next_control!='='&&next_control!=eq_eq{
continue
}

/*:167*/
//line gotangle.w:1367

break
}
return
}
no_where= false
print_where= false
/*168:*/
//line gotangle.w:1430

/*169:*/
//line gotangle.w:1435

tok_mem= append(tok_mem,unicode.UpperLower+section_number)
tok_mem= append(tok_mem,section_count)

/*:169*/
//line gotangle.w:1431

scan_repl(section_name)
/*170:*/
//line gotangle.w:1439

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


/*:170*/
//line gotangle.w:1433


/*:168*/
//line gotangle.w:1374

}

/*:164*//*171:*/
//line gotangle.w:1456

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

/*:171*//*172:*/
//line gotangle.w:1472

func skip_limbo(){
for true{
if loc>=len(buffer)&&!get_line(){
return
}
for loc<len(buffer)&&buffer[loc]!='@'{
loc++
}
if loc++;loc<len(buffer){
c:= buffer[loc]
loc++
cc:= ignore
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

/*:172*//*173:*/
//line gotangle.w:1511

func print_stats(){
fmt.Print("\nMemory usage statistics:\n")
fmt.Printf("%v names\n",len(name_dir))
fmt.Printf("%v replacement texts\n",len(text_info))
}

/*:173*/
