% This file is part of GOWEB.
% This program by Alexander Sychev
% is based on a program CWEAVE by Silvio Levy and Donald E. Knuth
% It is distributed WITHOUT ANY WARRANTY, express or implied.
% Version 0.1 --- April 2012

% Copyright (C) 2012 Alexander Sychev

% Permission is granted to make and distribute verbatim copies of this
% document provided that the copyright notice and this permission notice
% are preserved on all copies.

% Permission is granted to copy and distribute modified versions of this
% document under the conditions for verbatim copying, provided that the
% entire resulting derived work is given a different name and distributed
% under the terms of a permission notice identical to this one.

% Here is TeX material that gets inserted after \input cwebmac
\def\hang{\hangindent 3em\indent\ignorespaces}
\def\pb{$\.|\ldots\.|$} % C brackets (|...|)
\def\v{\char'174} % vertical (|) in typewriter font
\def\dleft{[\![} \def\dright{]\!]} % double brackets
\mathchardef\RA="3221 % right arrow
\mathchardef\BA="3224 % double arrow
\def\({} % ) kludge for alphabetizing certain section names
\def\TeXxstring{\\{\TEX/\_string}}
\def\skipxTeX{\\{skip\_\TEX/}}
\def\copyxTeX{\\{copy\_\TEX/}}

\def\title{GOWEAVE (Version 0.1)}
\def\topofcontents{\null\vfill
	\centerline{\titlefont The {\ttitlefont GOWEAVE} processor}
	\vskip 15pt
	\centerline{(Version 0.1)}
	\vfill}
\def\botofcontents{\vfill
\noindent
Copyright \copyright\ 2012 Alexander Sychev
\bigskip\noindent
Permission is granted to make and distribute verbatim copies of this
document provided that the copyright notice and this permission notice
are preserved on all copies.

\smallskip\noindent
Permission is granted to copy and distribute modified versions of this
document under the conditions for verbatim copying, provided that the
entire resulting derived work is given a different name and distributed
under the terms of a permission notice identical to this one.
}
\pageno=\contentspagenumber \advance\pageno by 1
\let\maybe=\iftrue
@s not_eq normal @q unreserve a C++ keyword @>

@** Introduction.
This is the \.{GOWEAVE} program by Alexander Syhcev
based on \.{CWEAVE} by Silvio Levy and Donald E. Knuth.

The ``banner line'' defined here should be changed whenever \.{GOWEAVE}
is modified.

@<Constants@>=
banner = "This is GOWEAVE (Version 0.1)\n"

@ @c
package main

import (
@<Import packages@>@/
)

const (
@<Constants@>@/
)

@<Typedef declarations@>@/
@<Global variables@>@/

@ \.{GOWEAVE} has a fairly straightforward outline.  It operates in
three phases: First it inputs the source file and stores cross-reference
data, then it inputs the source once again and produces the \TEX/ output
file, finally it sorts and outputs the index.

Please read the documentation for \.{common}, the set of routines common
to \.{GOTANGLE} and \.{GOWEAVE}, before proceeding further.

@ @c
func main () {
	flags['x']=true
	flags['f']=true
	flags['e']=true /* controlled by command-line options */
	common_init()
	@<Set initial values@>
	if show_banner() {
		fmt.Print(banner) /* print a ``banner line'' */
	}
	@<Store all the reserved words@>
	phase_one() /* read all the user's text and store the cross-references */
	phase_two() /* read all the text again and translate it to \TEX/ form */
	phase_three() /* output the cross-reference index */
	os.Exit(wrap_up()) /* and exit gracefully */
}

@ The following parameters were sufficient in the original \.{WEAVE} to
handle \TEX/, so they should be sufficient for most applications of \.{GOWEAVE}.

@<Constants@>=
max_names = 4000 /* number of identifiers, strings, section names;
	must be less than 10240; used in |"common.w"| */
line_length = 80 /* lines of \TEX/ output have at most this many characters;
	should be less than 256 */
max_scraps = 2000 /* number of tokens in \GO/ texts being parsed */
stack_size = 400 /* number of simultaneous output levels */

@ The next few sections contain stuff from the file |"common.w"| that must
be included in both |"gotangle.w"| and |"goweave.w"|. 

@i common.w

@* Data structures exclusive to {\tt GOWEAVE}.
As explained in \.{common.w}, the field of a |name_info| structure
that contains the |rlink| of a section name is used for a completely
different purpose in the case of identifiers. It is then called the
|ilk| of the identifier, and it is used to
distinguish between various types of identifiers, as follows:

\yskip\hang |normal| and |func_template| identifiers are part of the
\GO/ program that will  appear in italic type (or in typewriter type
if all uppercase).

\yskip\hang |custom| identifiers are part of the \GO/ program that
will be typeset in special ways.

\yskip\hang |roman| identifiers are index entries that appear after
\.{@@\^} in the \.{CWEB} file.

\yskip\hang |wildcard| identifiers are index entries that appear after
\.{@@:} in the \.{CWEB} file.

\yskip\hang |typewriter| identifiers are index entries that appear after
\.{@@.} in the \.{CWEB} file.

\yskip\hang |alfop|, \dots, |template_like|
identifiers are \GO/ reserved words whose |ilk|
explains how they are to be treated when \GO/ code is being
formatted.

@<More elements of |name_info| structure@>=
	ilk		int32 /* used by identifiers in \.{GOWEAVE} only */

@ @<Constants@>=
normal = 0 /* ordinary identifiers have |normal| ilk */
roman = 1 /* normal index entries have |roman| ilk */
wildcard = 2 /* user-formatted index entries have |wildcard| ilk */
typewriter = 3 /* `typewriter type' entries have |typewriter| ilk */
func_template = 4 /* identifiers that can be followed by optional template */
custom = 5 /* identifiers with user-given control sequence */
alfop = 22 /* alphabetic operators like \&{and} or \&{not\_eq} */
else_like = 26 /* \&{else} */
public_like = 40 /* \&{public}, \&{private}, \&{protected} */
operator_like = 41 /* \&{operator} */
new_like = 42 /* \&{new} */
catch_like = 43 /* \&{catch} */
for_like = 45 /* \&{for}, \&{switch}, \&{while} */
do_like = 46 /* \&{do} */
if_like = 47 /* \&{if}, \&{ifdef}, \&{endif}, \&{pragma}, \dots */
delete_like = 48 /* \&{delete} */
raw_ubin = 49 /* `\.\&' or `\.*' when looking for \&{const} following */
const_like = 50 /* \&{const}, \&{volatile} */
raw_int = 51 /* \&{int}, \&{char}, \dots; also structure and class names  */
int_like = 52 /* same, when not followed by left parenthesis or \DC\ */
case_like = 53 /* \&{case}, \&{return}, \&{goto}, \&{break}, \&{continue} */
sizeof_like = 54 /* \&{sizeof} */
struct_like = 55 /* \&{struct}, \&{union}, \&{enum}, \&{class} */
typedef_like rune = 56 /* \&{typedef} */
define_like = 57 /* \&{define} */
template_like = 58 /* \&{template} */

@ We keep track of the current section number in |section_count|, which
is the total number of sections that have started.  Sections which have
been altered by a change file entry have their |changed_section| flag
turned on during the first phase.

@<Global...@>=
var change_exists bool /* has any section changed? */

@ The other large memory area in \.{GOWEAVE} keeps the cross-reference data.
All uses of the name |p| are recorded in a linked list beginning at
|name_dir[p].xref|, which is an index in the |xmem| array. The elements of |xmem|
are structures consisting of an integer, |num|, and an index |xlink|
to another element of |xmem|.  If |x=name_dir[p].xref| is an index into |xmem|,
the value of |xmem[x].num| is either a section number where |p| is used,
or |cite_flag| plus a section number where |p| is mentioned,
or |def_flag| plus a section number where |p| is defined;
and |xmem[x].xlink| points to the next such cross-reference for |p|,
if any. This list of cross-references is in decreasing order by
section number. The linked list ends at |0|.

The global variable |xref_switch| is set either to |def_flag| or to zero,
depending on whether the next cross-reference to an identifier is to be
underlined or not in the index. This switch is set to |def_flag| when
\.{@@!} or \.{@@d} is scanned, and it is cleared to zero when
the next identifier or index entry cross-reference has been made.
Similarly, the global variable |section_xref_switch| is either
|def_flag| or |cite_flag| or zero, depending
on whether a section name is being defined, cited or used in \GO/ text.

@<Type...@>=
type xref_info struct{
	num int32 /* section number plus zero or |def_flag| */
	xlink int32 /* index of the previous cross-reference */
}

@ @<Global...@>=
var xmem[]xref_info /* contains cross-reference information */
var xref_switch int32
var section_xref_switch int32 /* either zero or |def_flag| */

@ A section that is used for multi-file output (with the \.{@@(} feature)
has a special first cross-reference whose |num| field is |file_flag|.

@ @<Constants@>=
cite_flag = 10240 /* must be strictly larger than |max_sections| */
file_flag = 3*cite_flag
def_flag = 2*cite_flag

@ @<More elements of |name...@>=
xref int32 /* info corresponding to names */

@ @<Set init...@>=
xmem=append(xmem, xref_info{})
xref_switch=0
section_xref_switch=0

@ A new cross-reference for an identifier is formed by calling |new_xref|,
which discards duplicate entries and ignores non-underlined references
to one-letter identifiers or \GO/'s reserved words.

If the user has sent the |flags['x']==false| flag (the \.{-x} option of the command line),
it is unnecessary to keep track of cross-references for identifiers.
If one were careful, one could probably make more changes around section
100 to avoid a lot of identifier looking up.

@c
func append_xref(c int32) {
	xmem = append(xmem, xref_info{})
	xmem[len(xmem)-1].num=c
	xmem[len(xmem)-1].xlink=0
}

func is_tiny(p int32) bool {
	return p<int32(len(name_dir)) && len(name_dir[p].name) == 1
}

/* tells if uses of a name are to be indexed */ 
func unindexed(p int32) bool {
	return p<res_wd_end && name_dir[p].ilk>=custom
}

@ @c
func new_xref(p int32){
	if flags['x']==false {
		return
	}
	if (unindexed(p)|| is_tiny(p)) && xref_switch==0 {
		return
	}
	m:=section_count+xref_switch
	xref_switch=0
	q:=name_dir[p].xref /* pointer to previous cross-reference */
	if q >= 0 {
		n:=xmem[q].num /* new and previous cross-reference value */
		if n==m || n==m+def_flag {
			return
		} else if m==n+def_flag {
			xmem[q].num=m
			return
		}
	}
	append_xref(m)
	xmem[len(xmem)-1].xlink=int32(q)
	name_dir[p].xref=int32(len(xmem)-1)
}

@ The cross-reference lists for section names are slightly different.
Suppose that a section name is defined in sections $m_1$, \dots,
$m_k$, cited in sections $n_1$, \dots, $n_l$, and used in sections
$p_1$, \dots, $p_j$.  Then its list will contain $m_1+|def_flag|$,
\dots, $m_k+|def_flag|$, $n_1+|cite_flag|$, \dots,
$n_l+|cite_flag|$, $p_1$, \dots, $p_j$, in this order.

Although this method of storage takes quadratic time with respect to
the length of the list, under foreseeable uses of \.{GOWEAVE} this inefficiency
is insignificant.

@c
func new_section_xref(p int32) {
	var r int32 = 0 /* pointers to previous cross-references */
	q:=name_dir[p].xref
	
	if q>=0 {
		for q>=0 && q<int32(len(xmem)) && xmem[q].num>section_xref_switch {
			r=q
			q=xmem[q].xlink
		}
	}
	if r > 0 && r<int32(len(xmem)) && xmem[r].num==section_count+section_xref_switch {
		return /* don't duplicate entries */
	}
	append_xref(section_count+section_xref_switch)
	xmem[len(xmem)-1].xlink=q
	section_xref_switch=0
	if r==0 {
		name_dir[p].xref=int32(len(xmem)-1)
	} else {
		xmem[r].xlink=int32(len(xmem)-1)
	}
}

@ The cross-reference list for a section name may also begin with
|file_flag|. Here's how that flag gets put~in.

@c
func set_file_flag(p int32) {
	q:=name_dir[p].xref
	if xmem[q].num==file_flag {
		return
	}
	append_xref(file_flag)
	xmem[len(xmem)-1].xlink = q
	name_dir[p].xref = int32(len(xmem)-1)
}

@ A third large area of memory is used for sixteen-bit `tokens', which appear
in short lists similar to the strings of characters in |byte_mem|. Token lists
are used to contain the result of \GO/ code translated into \TEX/ form;
further details about them will be explained later.

@ @<Global...@>=
var tok_mem []rune /* tokens */
var tok_start []int32 /* directory into |tok_mem| */
var max_tok_ptr int /* largest length of |tok_mem| */
var max_text_ptr int /* largest length of |tok_start| */

@ @<Set init...@>=
tok_start=append(tok_start, 0)
max_tok_ptr=1
max_text_ptr=1

@ Here are the three procedures needed to complete |id_lookup|:
@c
func names_match(
	p int32, /* points to the proposed match */
	id []rune,
	t int32 /* desired ilk */ ) bool {
	if len(name_dir[p].name)!=len(id) {
		return false
	}
	if name_dir[p].ilk!=t && !(t==normal && name_dir[p].ilk>typewriter) {
		return false
	}
	return compare_runes(id,name_dir[p].name) == 0
}

func init_p(p int32, t int32) {
	name_dir[p].ilk=t
	name_dir[p].xref=0
}

func init_node(p int32){
	name_dir[p].xref=0
}

@ We have to get \GO/'s
reserved words into the hash table, and the simplest way to do this is
to insert them every time \.{GOWEAVE} is run.  Fortunately there are relatively
few reserved words. (Some of these are not strictly ``reserved,'' but
are defined in header files of the ISO Standard \GO/ Library.)
@^reserved words@>

@<Store all the reserved words@>=
id_lookup([]rune("and"),alfop)
id_lookup([]rune("and_eq"),alfop)
id_lookup([]rune("asm"),sizeof_like)
id_lookup([]rune("auto"),int_like)
id_lookup([]rune("bitand"),alfop)
id_lookup([]rune("bitor"),alfop)
id_lookup([]rune("bool"),raw_int)
id_lookup([]rune("break"),case_like)
id_lookup([]rune("case"),case_like)
id_lookup([]rune("catch"),catch_like)
id_lookup([]rune("char"),raw_int)
id_lookup([]rune("class"),struct_like)
id_lookup([]rune("clock_t"),raw_int)
id_lookup([]rune("compl"),alfop)
id_lookup([]rune("const"),const_like)
id_lookup([]rune("const_cast"),raw_int)
id_lookup([]rune("continue"),case_like)
id_lookup([]rune("default"),case_like)
id_lookup([]rune("define"),define_like)
id_lookup([]rune("defined"),sizeof_like)
id_lookup([]rune("delete"),delete_like)
id_lookup([]rune("div_t"),raw_int)
id_lookup([]rune("do"),do_like)
id_lookup([]rune("double"),raw_int)
id_lookup([]rune("dynamic_cast"),raw_int)
id_lookup([]rune("elif"),if_like)
id_lookup([]rune("else"),else_like)
id_lookup([]rune("endif"),if_like)
id_lookup([]rune("enum"),struct_like)
id_lookup([]rune("error"),if_like)
id_lookup([]rune("explicit"),int_like)
id_lookup([]rune("export"),int_like)
id_lookup([]rune("extern"),int_like)
id_lookup([]rune("FILE"),raw_int)
id_lookup([]rune("float"),raw_int)
id_lookup([]rune("for"),for_like)
id_lookup([]rune("fpos_t"),raw_int)
id_lookup([]rune("friend"),int_like)
id_lookup([]rune("goto"),case_like)
id_lookup([]rune("if"),if_like)
id_lookup([]rune("ifdef"),if_like)
id_lookup([]rune("ifndef"),if_like)
id_lookup([]rune("include"),if_like)
id_lookup([]rune("inline"),int_like)
id_lookup([]rune("int"),raw_int)
id_lookup([]rune("jmp_buf"),raw_int)
id_lookup([]rune("ldiv_t"),raw_int)
id_lookup([]rune("line"),if_like)
id_lookup([]rune("long"),raw_int)
id_lookup([]rune("mutable"),int_like)
id_lookup([]rune("namespace"),struct_like)
id_lookup([]rune("new"),new_like)
id_lookup([]rune("not"),alfop)
id_lookup([]rune("not_eq"),alfop)
id_lookup([]rune("NULL"),custom)
id_lookup([]rune("offsetof"),raw_int)
id_lookup([]rune("operator"),operator_like)
id_lookup([]rune("or"),alfop)
id_lookup([]rune("or_eq"),alfop)
id_lookup([]rune("pragma"),if_like)
id_lookup([]rune("private"),public_like)
id_lookup([]rune("protected"),public_like)
id_lookup([]rune("ptrdiff_t"),raw_int)
id_lookup([]rune("public"),public_like)
id_lookup([]rune("register"),int_like)
id_lookup([]rune("reinterpret_cast"),raw_int)
id_lookup([]rune("return"),case_like)
id_lookup([]rune("short"),raw_int)
id_lookup([]rune("sig_atomic_t"),raw_int)
id_lookup([]rune("signed"),raw_int)
id_lookup([]rune("size_t"),raw_int)
id_lookup([]rune("sizeof"),sizeof_like)
id_lookup([]rune("static"),int_like)
id_lookup([]rune("static_cast"),raw_int)
id_lookup([]rune("struct"),struct_like)
id_lookup([]rune("switch"),for_like)
id_lookup([]rune("template"),template_like)
id_lookup([]rune("this"),custom)
id_lookup([]rune("throw"),case_like)
id_lookup([]rune("time_t"),raw_int)
id_lookup([]rune("try"),else_like)
id_lookup([]rune("typedef"),typedef_like)
id_lookup([]rune("typeid"),raw_int)
id_lookup([]rune("typename"),struct_like)
id_lookup([]rune("undef"),if_like)
id_lookup([]rune("union"),struct_like)
id_lookup([]rune("unsigned"),raw_int)
id_lookup([]rune("using"),int_like)
id_lookup([]rune("va_dcl"),decl) /* Berkeley's variable-arg-list convention */
id_lookup([]rune("va_list"),raw_int) /* ditto */
id_lookup([]rune("virtual"),int_like)
id_lookup([]rune("void"),raw_int)
id_lookup([]rune("volatile"),const_like)
id_lookup([]rune("wchar_t"),raw_int)
id_lookup([]rune("while"),for_like)
id_lookup([]rune("xor"),alfop)
id_lookup([]rune("xor_eq"),alfop)
res_wd_end=int32(len(name_dir))
id_lookup([]rune("TeX"),custom)
id_lookup([]rune("make_pair"),func_template)

@* Lexical scanning.
Let us now consider the subroutines that read the \.{CWEB} source file
and break it into meaningful units. There are four such procedures:
One simply skips to the next `\.{@@\ }' or `\.{@@*}' that begins a
section; another passes over the \TEX/ text at the beginning of a
section; the third passes over the \TEX/ text in a \GO/ comment;
and the last, which is the most interesting, gets the next token of
a \GO/ text.  They all use the pointers |limit| and |loc| into
the line of input currently being studied.

@ Control codes in \.{CWEB}, which begin with `\.{@@}', are converted
into a numeric code designed to simplify \.{GOWEAVE}'s logic; for example,
larger numbers are given to the control codes that denote more significant
milestones, and the code of |new_section| should be the largest of
all. Some of these numeric control codes take the place of |char|
control codes that will not otherwise appear in the output of the
scanning routines.
@^ASCII code dependencies@>

@ @<Constants@>=
ignore rune = 00 /* control code of no interest to \.{GOWEAVE} */
verbatim rune = 02 /* takes the place of extended ASCII \.{\char2} */
begin_short_comment rune = 03 /* short comment */
begin_comment rune = '\t' /* tab marks will not appear */
underline rune = '\n' /* this code will be intercepted without confusion */
noop rune = 0177 /* takes the place of ASCII delete */
xref_roman rune = 0203 /* control code for `\.{@@\^}' */
xref_wildcard rune = 0204 /* control code for `\.{@@:}' */
xref_typewriter rune = 0205 /* control code for `\.{@@.}' */
TeX_string rune = 0206 /* control code for `\.{@@t}' */
ord rune = 0207 /* control code for `\.{@@'}' */
join rune = 0210 /* control code for `\.{@@\&}' */
thin_space rune = 0211 /* control code for `\.{@@,}' */
math_break rune = 0212 /* control code for `\.{@@\v}' */
line_break rune = 0213 /* control code for `\.{@@/}' */
big_line_break rune = 0214 /* control code for `\.{@@\#}' */
no_line_break rune = 0215 /* control code for `\.{@@+}' */
pseudo_semi rune = 0216 /* control code for `\.{@@;}' */
macro_arg_open rune = 0220 /* control code for `\.{@@[}' */
macro_arg_close rune = 0221 /* control code for `\.{@@]}' */
trace rune = 0222 /* control code for `\.{@@0}', `\.{@@1}' and `\.{@@2}' */
output_defs_code rune = 0224 /* control code for `\.{@@h}' */
format_code rune = 0225 /* control code for `\.{@@f}' and `\.{@@s}' */
definition rune = 0226 /* control code for `\.{@@d}' */
begin_code rune = 0227 /* control code for `\.{@@c}' */
section_name rune = 0230 /* control code for `\.{@@<}' */
new_section rune = 0231 /* control code for `\.{@@\ }' and `\.{@@*}' */

@ @f TeX_string TeX

@ Control codes are converted to \.{GOWEAVE}'s internal
representation by means of the table |ccode|.

@<Global...@>=
var ccode[256] rune /* meaning of a char following \.{@@} */

@ @<Set ini...@>=
{
	for c:=0; c<256; c++ {
		ccode[c]=ignore
	}
}
ccode[' ']=new_section
ccode['\t']=new_section
ccode['\n']=new_section
ccode['\v']=new_section
ccode['\r']=new_section
ccode['\f']=new_section
ccode['*']=new_section
ccode['@@']='@@' /* `quoted' at sign */
ccode['=']=verbatim
ccode['d']=definition
ccode['D']=definition
ccode['f']=format_code
ccode['F']=format_code
ccode['s']=format_code
ccode['S']=format_code
ccode['c']=begin_code
ccode['C']=begin_code
ccode['p']=begin_code
ccode['P']=begin_code
ccode['t']=TeX_string
ccode['T']=TeX_string
ccode['q']=noop
ccode['Q']=noop
ccode['h']=output_defs_code
ccode['H']=output_defs_code
ccode['&']=join
ccode['<']=section_name
ccode['(']=section_name
ccode['!']=underline
ccode['^']=xref_roman
ccode[':']=xref_wildcard
ccode['.']=xref_typewriter
ccode[',']=thin_space
ccode['|']=math_break
ccode['/']=line_break
ccode['#']=big_line_break
ccode['+']=no_line_break
ccode[';']=pseudo_semi
ccode['[']=macro_arg_open
ccode[']']=macro_arg_close
ccode['\'']=ord
@<Special control codes for debugging@>

@ Users can write
\.{@@2}, \.{@@1}, and \.{@@0} to turn tracing fully on, partly on,
and off, respectively.

@<Special control codes...@>=
ccode['0']=trace
ccode['1']=trace
ccode['2']=trace

@ The |skip_limbo| routine is used on the first pass to skip through
portions of the input that are not in any sections, i.e., that precede
the first section. After this procedure has been called, the value of
|input_has_ended| will tell whether or not a section has actually been found.

There's a complication that we will postpone until later: If the \.{@@s}
operation appears in limbo, we want to use it to adjust the default
interpretation of identifiers.

@ @c
func skip_limbo() {
	for true {
		if loc>=len(buffer) && !get_line() {
			return
		}
		for loc < len(buffer) && buffer[loc]!='@@' {
			loc++ /* look for '@@', then skip two chars */
		}
		l := loc
		loc++
		if l <len(buffer) { 
			c:=ccode[buffer[loc]]
			loc++
			if c==new_section {
				return
			}
			if c==noop {
				skip_restricted()
			} else if c==format_code {
				@<Process simple format in limbo@>
			}
		}
	}
}

@ The |skip_TeX| routine is used on the first pass to skip through
the \TEX/ code at the beginning of a section. It returns the next
control code or `\.{\v}' found in the input. A |new_section| is
assumed to exist at the very end of the file.

@f skip_TeX TeX

@c
/* skip past pure \TEX/ code */
func skip_TeX() rune {
	for true {
		if loc>=len(buffer) && !get_line() {
			return new_section
		}
		for loc < len(buffer) && buffer[loc]!='@@' && buffer[loc]!='|' {
			loc++
		}
		l := loc
		loc++
		if l < len(buffer) && buffer[l] =='|' {
			return '|' 
		}
		if loc<len(buffer) {
			l := loc
			loc++
			return ccode[buffer[l]]
		}
	}
	return 0
}

@*1 Inputting the next token.
As stated above, \.{GOWEAVE}'s most interesting lexical scanning routine is the
|get_next| function that inputs the next token of \GO/ input. However,
|get_next| is not especially complicated.

The result of |get_next| is either a |rune| code for some special character,
or it is a special code representing a pair of characters (e.g., `\.{!=}'),
or it is the numeric value computed by the |ccode|
table, or it is one of the following special codes:

\yskip\hang |identifier|: In this case the global variable |id| will contain
an identifier, as required by the |id_lookup| routine.

\yskip\hang |str|: The string will have been copied into the array
|section_text|; |id| are set as above (now it is a slice of |section_text|).

\yskip\hang |constant|: The constant is copied into |section_text|, with
slight modifications; |id| is set.

\yskip\noindent Furthermore, some of the control codes cause
|get_next| to take additional actions:

\yskip\hang |xref_roman|, |xref_wildcard|, |xref_typewriter|, |TeX_string|,
|verbatim|: The values of |id| will have been set to
the slice of the buffer.

\yskip\hang |section_name|: In this case the global variable |cur_section| will
point to the |byte_start| entry for the section name that has just been scanned.
The value of |cur_section_char| will be |'('| if the section name was
preceded by \.{@@(} instead of \.{@@<}.

\yskip\noindent If |get_next| sees `\.{@@!}'
it sets |xref_switch| to |def_flag| and goes on to the next token.

@ @<Constants@>=
constant = 0200 /* \GO/ constant */
str = 0201 /* \GO/ string */
identifier = 0202 /* \GO/ identifier or reserved word */

@ @<Global...@>=
var cur_section int32 /* name of section just scanned */
var cur_section_char rune /* the character just before that name */


@ As one might expect, |get_next| consists mostly of a big switch
that branches to the various special cases that can arise.
\GO/ allows underscores to appear in identifiers, and some \GO/
compilers even allow the dollar sign.

@ @c
/* produces the next input token */
func get_next() rune { 
	for true {
		@<Check if we're at the end of a preprocessor command@>
		if loc>=len(buffer) && !get_line() {
			return new_section 
		}
		@+c:=buffer[loc] /* the current character */
		loc++
		nc:=ignore
		if loc < len(buffer) {
			nc = buffer[loc]
		}
		if unicode.IsDigit(c) || c=='.' {
			@<Get a constant@>
		} else if c=='\'' || c=='"' || c=='L' && 
			(nc=='\'' || nc=='"') || c=='<' && sharp_include_line {
			@<Get a string@>
		} else if unicode.IsLetter(c) || c=='_' || c=='$' {
			@<Get an identifier@>
		} else if c=='@@' {
			@<Get control code and possible section name@>
		} else if unicode.IsSpace(c) {
			continue /* ignore spaces and tabs */
		}
		if c=='#' && loc==1 {
			@<Raise preprocessor flag@>
		}
mistake: 
		@<Compress two-symbol operator@>
		return c
	}
	return 0
}

@ Because preprocessor commands do not fit in with the rest of the syntax
of \GO/,
we have to deal with them separately.  One solution is to enclose such
commands between special markers.  Thus, when a \.\# is seen as the
first character of a line, |get_next| returns a special code
|left_preproc| and raises a flag |preprocessing|.

We can use the same internal code number for |left_preproc| as we do
for |ord|, since |get_next| changes |ord| into a string.

@ @<Constants@>=
left_preproc = ord /* begins a preprocessor command */
right_preproc = 0217 /* ends a preprocessor command */

@ @<Glob...@>= 
var preprocessing bool=false /* are we scanning a preprocessor command? */

@ @<Raise prep...@>= {
	preprocessing=true
	@<Check if next token is |include|@>
	return left_preproc
}

@ An additional complication is the freakish use of \.< and \.> to delimit
a file name in lines that start with \.{\#include}.  We must treat this file
name as a string.

@<Glob...@>=
var sharp_include_line bool=false /* are we scanning a |#include| line? */

@ @<Check if next token is |include|@>=
for len(buffer[loc:])>=7 && unicode.IsSpace(buffer[loc]) {
	loc++
}
if len(buffer[loc:])>=7 && compare_runes(buffer[loc:loc+7], []rune("include"))==0 {
	sharp_include_line=true
}

@ When we get to the end of a preprocessor line,
we lower the flag and send a code |right_preproc|, unless
the last character was a \.\\.

@<Check if we're at...@>=
	for loc==len(buffer)-1 && preprocessing && buffer[loc]=='\\' {
		if !get_line() {
			return new_section /* still in preprocessor mode */
		}
	}
	if loc>=len(buffer) && preprocessing {
		preprocessing=false
		sharp_include_line=false
		return right_preproc
	}

@ The following code assigns values to the combinations \.{++},
\.{--}, \.{>=}, \.{<=}, \.{==}, \.{<<}, \.{>>}, \.{!=}, \.{\v\v}, and
\.{\&\&}, \.{...}.
The compound assignment operators (e.g., \.{+=}) are
treated as separate tokens.

@<Compress tw...@>=
switch(c) {
	case '/': 
		if nc=='*' {
			l := loc
			loc++
			if l <=len(buffer) {
				return begin_comment
			}
		} else if nc=='/' { 
			l := loc
			loc++
			if l <=len(buffer) {
				return begin_short_comment
			}
		}
	case '+': 
		if nc=='+' {
			l := loc
			loc++
			if l <=len(buffer) {
				return plus_plus
			}
		}	
	case '-': 
		if nc=='-' {
			l := loc
			loc++
			if l <=len(buffer) {
				return minus_minus
			}
		}
	case '.': 
		if nc=='.' && loc+1<len(buffer) && buffer[loc+1]=='.' {
			loc++
			l := loc
			loc++
			if l <=len(buffer) {
				return dot_dot_dot
			}
		}
	case '=': 
		if nc=='=' {
			l := loc
			loc++
			if l <=len(buffer) {
				return eq_eq
			}
		}
	case '>': 
		if nc=='=' {
			l := loc
			loc++
			if l <=len(buffer) {
				return gt_eq
			}
		} else if nc=='>' {
			l := loc
			loc++
			if l <=len(buffer) {
				return gt_gt
			}
		}
	case '<': 
		if nc=='=' {
			l := loc
			loc++
			if l <=len(buffer) {
				return lt_eq
			}
		} else if nc=='<' {
			l := loc
			loc++
			if l <=len(buffer) {
				return lt_lt
			}
		}
	case '&': 
		if nc=='&' {
			l := loc
			loc++
			if l <=len(buffer) {
				return and_and
			}
		}
	case '|': 
		if nc=='|' {
			l := loc
			loc++
			if l <=len(buffer) {
				return or_or
			}
		}
	case '!':
		if nc=='=' {
			l := loc
			loc++
			if l <=len(buffer) {
				return not_eq
			}
		}
}

@ @<Get an identifier@>= {
	loc--
	id_first:=loc
	for loc < len(buffer) && 
		(unicode.IsLetter(buffer[loc]) || 
		unicode.IsDigit(buffer[loc]) || 
		buffer[loc]=='_' || 
		buffer[loc]=='$') {
		loc++
	}
	id = buffer[id_first:loc]
	return identifier
}

@ Different conventions are followed by \TEX/ and \GO/ to express octal
and hexadecimal numbers; it is reasonable to stick to each convention
within its realm.  Thus the \GO/ part of a \.{CWEB} file has octals
introduced by \.0 and hexadecimals by \.{0x}, but \.{GOWEAVE} will print
with \TEX/ macros that the user can redefine to fit the context.
In order to simplify such macros, we replace some of the characters.

Notice that in this section and the next, |id|
is a slice of the array |section_text|, not of |buffer|.

@<Get a constant@>= {
	id = nil
	is_dec := false
	if loc < len(buffer) && buffer[loc-1]=='0' {
		if buffer[loc]=='x' || buffer[loc]=='X' /* hex constant */ {
			id = append(id,'^')
			loc++
			for loc < len(buffer) && xisxdigit(buffer[loc]) {
				id = append(id, buffer[loc])
				loc++
			}
		} else if unicode.IsDigit(buffer[loc]) /* octal constant */{
			id = append(id,'~')
			for loc < len(buffer) && unicode.IsDigit(buffer[loc]) {
				id = append(id, buffer[loc])
				loc++
			}
		} else {
			is_dec = true/* decimal constant */
		}
	} else {
		is_dec =  true
	}
	if is_dec { /* decimal constant */
		if loc < len(buffer) && buffer[loc-1]=='.' && !unicode.IsDigit(buffer[loc]) {
			goto mistake /* not a constant */
		}
		id = append(id, buffer[loc-1])
		for loc < len(buffer) && (unicode.IsDigit(buffer[loc]) || buffer[loc]=='.') {
			id = append(id, buffer[loc])
			loc++
		}
		if loc < len(buffer) && (buffer[loc]=='e' || buffer[loc]=='E') { /* float constant */
			id = append(id, '_')
			loc++
			if loc < len(buffer) && (buffer[loc]=='+' || buffer[loc]=='-') {
				id = append(id, buffer[loc])
				loc++
			}
			for loc < len(buffer) && unicode.IsDigit(buffer[loc]) {
				id = append(id, buffer[loc])
				loc++
			}
		}
	}
	for loc < len(buffer) && 
		(buffer[loc]=='u' || 
		buffer[loc]=='U' || 
		buffer[loc]=='l' || 
		buffer[loc]=='L' || 
		buffer[loc]=='f' || 
		buffer[loc]=='F') {
		id = append(id, '$')
		id = append(id, unicode.ToUpper(buffer[loc]))
		loc++
	}
	return constant
}

@ \GO/ strings and character constants, delimited by double and single
quotes, respectively, can contain newlines or instances of their own
delimiters if they are protected by a backslash.

@<Get a string@>= {
	delim := c /* what started the string */
	section_text = section_text[0:0]

	if delim=='\'' && buffer[loc-2]=='@@' {
		section_text = append(section_text, '@@')
		section_text = append(section_text, '@@')
	}
	section_text = append(section_text, delim)
	if loc < len(buffer) && delim=='L' { /* wide character constant */
		delim=buffer[loc]
		loc++
		section_text = append(section_text, delim)
	}
	if delim=='<' {
		 delim='>' /* for file names in |#include| lines */
	}
	for true {
		if loc>=len(buffer) {
			if buffer[len(buffer)-1]!='\\' {
				err_print("! String didn't end")
				loc=len(buffer)
				break
@.String didn't end@>
			}
			if !get_line()  {
				err_print("! Input ended in middle of string")
				loc=0
				break;
@.Input ended in middle of string@>
			}
		}
		l := loc
		loc++
		if c=buffer[l]; c==delim {
			section_text = append(section_text, c)
			break
		}
		if c=='\\' { 
			if loc>=len(buffer) {
				continue
			} 
			section_text = append(section_text, '\\')
			c=buffer[loc]
			loc++
		}
		section_text = append(section_text, c)
	}
	id = section_text
	return str 
}

@ After an \.{@@} sign has been scanned, the next character tells us
whether there is more work to do.

@<Get control code and possible section name@>= {
	c=nc
	loc++	
	switch ccode[c] {
@.Use @@l in limbo...@>
		case underline: 
			xref_switch=def_flag
			continue
		case trace: 
			tracing=c-'0'
			continue
		case xref_roman, xref_wildcard, xref_typewriter, noop, TeX_string: 
			c=ccode[c]
			skip_restricted()
			return c
		case section_name:
			@<Scan the section name and make |cur_section| point to it@>
		case verbatim: 
			@<Scan a verbatim string@>
		case ord: 
			@<Get a string@>
		default: 
			return ccode[c]
	}
}

@ The occurrence of a section name sets |xref_switch| to zero,
because the section name might (for example) follow \&{int}.

@<Scan the section name...@>= {
	section_text = section_text[0:0]
	@<Put section name into |section_text|@>
	if len(section_text)>3 && 
		compare_runes(section_text[len(section_text)-3:],[]rune("..."))==0 {	
		cur_section=section_lookup(section_text[0:len(section_text)-3],
										true) /* 1 means is a prefix */
	} else {
		cur_section=section_lookup(section_text, false)
	}
	xref_switch=0
	return section_name
}

@ Section names are placed into the |section_text| array with consecutive spaces,
tabs, and carriage-returns replaced by single spaces. There will be no
spaces at the beginning or the end. (We set |section_text[0]=' '| to facilitate
this, since the |section_lookup| routine uses |section_text[1]| as the first
character of the name.)


@ @<Put section name...@>=
for true {
	if loc>=len(buffer) {
		if !get_line() {
			err_print("! Input ended in section name")
			@.Input ended in section name@>
			loc=1
			break
		}
		if len(section_text) > 0 {
			section_text=append(section_text, ' ')
		}
	}
	c=buffer[loc]
	@<If end of name or erroneous nesting, |break|@>
	loc++
	if unicode.IsSpace(c) {
		c=' '
		if len(section_text) > 0 && section_text[len(section_text)-1]==' ' {
			 section_text = section_text[:len(section_text)-1]
		}
	}
	section_text=append(section_text, c)
}

@ @<If end of name...@>=
if c=='@@' {
	if loc+1 >= len(buffer) {
		err_print("! Section name didn't end")
		break
		@.Section name didn't end@>
	}
	c=buffer[loc+1]
	if (c=='>') {
		loc+=2
		break
	}
	cc := ignore
	if c < int32(len(ccode)) {
		cc = ccode[c]
	}
	if cc==new_section {
		err_print("! Section name didn't end")
		break
		@.Section name didn't end@>
	}
	if cc==section_name {
		err_print("! Nesting of section names not allowed")
		break
		@.Nesting of section names...@>
	}
	section_text = append(section_text, '@@')
	loc++ /* now |c==buffer[loc]| again */
}

@ This function skips over a restricted context at relatively high speed.

@ @c
func skip_restricted() {
	id_first:=loc
false_alarm:
	for loc < len(buffer) && buffer[loc]!='@@' {
		loc++
	}
	id = buffer[id_first:loc]
	loc++
	if loc>=len(buffer) {
		err_print("! Control text didn't end")
		loc=len(buffer)
@.Control text didn't end@>
	} else {
		if buffer[loc]=='@@'&&loc<=len(buffer) {
			loc++
			goto false_alarm
		}
		l := loc
		loc++
		if buffer[l]!='>' {
			err_print("! Control codes are forbidden in control text")
@.Control codes are forbidden...@>
		}
	}
}

@ At the present point in the program we
have |*(loc-1)==verbatim|; we set |id| to the string itself.
We also set |loc| to the position just after the ending delimiter.

@<Scan a verbatim string@>= {
	id_first:=loc
	loc++
	for loc + 1 < len(buffer) && ( buffer[loc]!='@@' || buffer[loc+1]!='>') {
		loc++
	}
	if loc>=len(buffer) {
		err_print("! Verbatim string didn't end")
@.Verbatim string didn't end@>
	}
	id = buffer[id_first:loc]
	loc+=2
	return verbatim
}

@** Phase one processing.
We now have accumulated enough subroutines to make it possible to carry out
\.{GOWEAVE}'s first pass over the source file. If everything works right,
both phase one and phase two of \.{GOWEAVE} will assign the same numbers to
sections, and these numbers will agree with what \.{GOTANGLE} does.

The global variable |next_control| often contains the most recent output of
|get_next|; in interesting cases, this will be the control code that
ended a section or part of a section.

@<Global...@>=
var next_control rune /* control code waiting to be acting upon */

@ The overall processing strategy in phase one has the following
straightforward outline.

@ @c
func phase_one() {
	phase=1
	reset_input()
	section_count=0
	skip_limbo()
	change_exists=false
	for !input_has_ended {
		@<Store cross-reference data for the current section@>
	}
	changed_section[section_count]=change_exists
		/* the index changes if anything does */
	phase=2 /* prepare for second phase */
	@<Print error messages about unused or undefined section names@>
}

@ @<Store cross-reference data...@>=
{
	section_count++
	changed_section[section_count]=changing
		 /* it will become 1 if any line changes */
	if buffer[loc-1]=='*' && show_progress() {
		fmt.Printf("*%d",section_count)
		os.Stdout.Sync() /* print a progress report */
	}
	@<Store cross-references in the \TEX/ part of a section@>
	@<Store cross-references in the definition part of a section@>
	@<Store cross-references in the \GO/ part of a section@>
	if changed_section[section_count] {
		change_exists=true
	}
}

@ The |C_xref| subroutine stores references to identifiers in
\GO/ text material beginning with the current value of |next_control|
and continuing until |next_control| is `\.\{' or `\.{\v}', or until the next
``milestone'' is passed (i.e., |next_control>=format_code|). If
|next_control>=format_code| when |C_xref| is called, nothing will happen;
but if |next_control=='|'| upon entry, the procedure assumes that this is
the `\.{\v}' preceding \GO/ text that is to be processed.

The parameter |spec_ctrl| is used to change this behavior. In most cases
|C_xref| is called with |spec_ctrl==ignore|, which triggers the default
processing described above. If |spec_ctrl==section_name|, section names will
be gobbled. This is used when \GO/ text in the \TEX/ part or inside comments
is parsed: It allows for section names to appear in \pb, but these
strings will not be entered into the cross reference lists since they are not
definitions of section names.

The program uses the fact that our internal code numbers satisfy
the relations |xref_roman==identifier+roman| and |xref_wildcard==identifier
+wildcard| and |xref_typewriter==identifier+typewriter|,
as well as |normal==0|.

@ @c
/* makes cross-references for \GO/ identifiers */
func C_xref(spec_ctrl rune) {
	for next_control<format_code || next_control==spec_ctrl {
		if next_control>=identifier && next_control<=xref_typewriter {
			if next_control>identifier {
				@<Replace |"@@@@"| by |"@@"| @>
			}
			p:=id_lookup(id,next_control-identifier)
			/* a referenced name */
			new_xref(p)
		}
		if next_control==section_name {
			section_xref_switch=cite_flag
			new_section_xref(cur_section)
		}
		next_control=get_next()
		if next_control=='|' || next_control==begin_comment ||
				next_control==begin_short_comment {
			return
		}
	}
}

@ The |outer_xref| subroutine is like |C_xref| except that it begins
with |next_control!='|'| and ends with |next_control>=format_code|. Thus, it
handles \GO/ text with embedded comments.

@ @c
/* extension of |C_xref| */
func outer_xref() {
	for next_control<format_code {
		if next_control!=begin_comment && next_control!=begin_short_comment {
			C_xref(ignore)
		} else {
			is_long_comment:=(next_control==begin_comment)
			bal:=copy_comment(is_long_comment,1)/* brace level in comment */
			next_control='|'
			for bal>0 {
				C_xref(section_name) /* do not reference section names in comments */
				if next_control=='|' {
					 bal=copy_comment(is_long_comment,bal)
				} else { 
					bal=0 /* an error message will occur in phase two */
				}
			}
		}
	}
}

@ In the \TEX/ part of a section, cross-reference entries are made only for
the identifiers in \GO/ texts enclosed in \pb, or for control texts
enclosed in \.{@@\^}$\,\ldots\,$\.{@@>} or \.{@@.}$\,\ldots\,$\.{@@>}
or \.{@@:}$\,\ldots\,$\.{@@>}.

@<Store cross-references in the \T...@>=
for true {
	next_control=skip_TeX()
	switch next_control {
		case underline:
			xref_switch=def_flag
			continue
		case trace: 
			tracing=buffer[loc-1]-'0'
			continue
		case '|': 
			C_xref(section_name)
			break
		case xref_roman, xref_wildcard, xref_typewriter, noop, section_name:
			loc-=2
			next_control=get_next() /* scan to \.{@@>} */
			if next_control>=xref_roman && next_control<=xref_typewriter {
				@<Replace |"@@@@"| by |"@@"| @>
				new_xref(id_lookup(id,next_control-identifier))
			}
			break
	}
	if next_control>=format_code {
		break
	}
}

@ @<Replace |"@@@@"| by |"@@"| @>=
{
	i:=0
	j:=0
	for i<len(id) {
		if id[i]=='@@' {
			i++
		}
		id[j] = id[i]
		j++
		i++
	}
	for j<i {
		id[j]=' ' /* clean up in case of error message display */
		j++
	}
}

@ During the definition and \GO/ parts of a section, cross-references
are made for all identifiers except reserved words. However, the right
identifier in a format definition is not referenced, and the left
identifier is referenced only if it has been explicitly
underlined (preceded by \.{@@!}).
The \TEX/ code in comments is, of course, ignored, except for
\GO/ portions enclosed in \pb; the text of a section name is skipped
entirely, even if it contains \pb\ constructions.

The variables |lhs| and |rhs| point to the respective identifiers involved
in a format definition.

@<Global...@>=
var lhs int32
var rhs int32 /* pointers to |byte_start| for format identifiers */
var res_wd_end int32

@ When we get to the following code we have |next_control>=format_code|.

@<Store cross-references in the d...@>=
for next_control<=definition { /* |format_code| or |definition| */
	if next_control==definition {
		xref_switch=def_flag /* implied \.{@@!} */
		next_control=get_next()
	} else {
		@<Process a format definition@>
	}
	outer_xref()
}

@ Error messages for improper format definitions will be issued in phase
two. Our job in phase one is to define the |ilk| of a properly formatted
identifier, and to remove cross-references to identifiers that we now
discover should be unindexed.

@<Process a form...@>= {
	next_control=get_next()
	if next_control==identifier {
		lhs=id_lookup(id,normal)
		name_dir[lhs].ilk=normal
		if xref_switch != 0 {
			new_xref(lhs)
		}
		next_control=get_next()
		if next_control==identifier {
			rhs=id_lookup(id,normal)
			name_dir[lhs].ilk=name_dir[rhs].ilk
			if unindexed(lhs) { 
				/* retain only underlined entries */
				var r int32 = 0
				for q:=name_dir[lhs].xref;q>=0;q=xmem[q].xlink {
					if xmem[q].num<def_flag {
						if r != 0 {
							xmem[r].xlink=xmem[q].xlink
						} else {
							name_dir[lhs].xref=xmem[q].xlink
						}
					} else {
						r=q
					}
				}
			}
			next_control=get_next()
		}
	}
}

@ A much simpler processing of format definitions occurs when the
definition is found in limbo.

@<Process simple format in limbo@>=
{
	if get_next()!=identifier {
		err_print("! Missing left identifier of @@s");
@.Missing left identifier...@>
	} else {
		lhs=id_lookup(id,normal)
		if get_next()!=identifier {
			err_print("! Missing right identifier of @@s");
@.Missing right identifier...@>
		} else {
			rhs=id_lookup(id,normal)
			name_dir[lhs].ilk=name_dir[rhs].ilk
		}
	}
}

@ Finally, when the \TEX/ and definition parts have been treated, we have
|next_control>=begin_code|.

@<Store cross-references in the \GO/...@>=
if next_control<=section_name {  /* |begin_code| or |section_name| */
	if next_control==begin_code {
		section_xref_switch=0
	} else {
		section_xref_switch=def_flag
		if cur_section_char=='(' && cur_section!=-1 {
			set_file_flag(cur_section)	
		}
	}
	for true {
		if next_control==section_name && cur_section!=-1 {
			new_section_xref(cur_section)
		}
		next_control=get_next()
		outer_xref()
		if next_control>section_name {
			break
		}
	}
}

@ After phase one has looked at everything, we want to check that each
section name was both defined and used.  The variable |cur_xref| will point
to cross-references for the current section name of interest.

@<Global...@>=
var cur_xref int32; /* temporary cross-reference pointer */
var an_output bool /* did |file_flag| precede |cur_xref|? */

@ The following recursive procedure
walks through the tree of section names and prints out anomalies.
@^recursion@>

@ @c
/* print anomalies in subtree |p| */
func section_check(p int32) {
	if p != -1 {
		section_check(name_dir[p].llink)
		cur_xref=name_dir[p].xref
		if xmem[cur_xref].num==file_flag {
			an_output=true
			cur_xref=xmem[cur_xref].xlink
		} else {
			an_output=false
		}
		if xmem[cur_xref].num<def_flag {
			fmt.Print("\n! Never defined: <")
			print_section_name(p)
			fmt.Print(">")
			mark_harmless()
@.Never defined: <section name>@>
		}
		for cur_xref != 0 && xmem[cur_xref].num >=cite_flag {
			cur_xref=xmem[cur_xref].xlink
		}
		if cur_xref==0 && !an_output {
			fmt.Print("\n! Never used: <")
			print_section_name(p)
			fmt.Print(">")
			mark_harmless()
@.Never used: <section name>@>
		}
		section_check(name_dir[p].rlink)
	}
}

@ @<Print error messages about un...@>=section_check(name_root)

@* Low-level output routines.
The \TEX/ output is supposed to appear in lines at most |line_length|
characters long, so we place it into an output buffer. During the output
process, |out_line| will hold the current line number of the line about to
be output.

@<Global...@>=
var out_buf[line_length + 1] rune/* assembled characters */
var out_ptr int32 /* just after last character in |out_buf| */
var out_buf_end int32 = line_length /* end of |out_buf| */
var out_line int /* number of next line to be output */

@ The |flush_buffer| routine empties the buffer up to a given breakpoint,
and moves any remaining characters to the beginning of the next line.
If the |per_cent| parameter is 1 a |'%'| is appended to the line
that is being output; in this case the breakpoint |b| should be strictly
less than |out_buf_end|. If the |per_cent| parameter is |0|,
trailing blanks are suppressed.
The characters emptied from the buffer form a new line of output;
if the |carryover| parameter is true, a |"%"| in that line will be
carried over to the next line (so that \TEX/ will ignore the completion
of commented-out text).

@c
/* outputs from |out_buf+1| to |b|,where |b<=out_ptr| */
func flush_buffer(b int32, per_cent bool,carryover bool) {
	j:=b /* pointer into |out_buf| */
	if !per_cent { /* remove trailing blanks */
		for j>0 && out_buf[j]==' ' {
			j--
		}
	}
	fmt.Fprint(active_file, string(out_buf[1:j+1]))
	if per_cent {
		fmt.Fprint(active_file, "%")
	}
	fmt.Fprint(active_file, "\n")
	out_line++
	if carryover {
		for j>0 {
			jj:=j
			j--
			if out_buf[jj]=='%' && (j==0 || out_buf[j]!='\\') {
				out_buf[b]='%'
				b--
				break
			}
		}
	}
	if b<out_ptr {
		copy(out_buf[1:],out_buf[b+1:])
	}
	out_ptr-=b
}

@ When we are copying \TEX/ source material, we retain line breaks
that occur in the input, except that an empty line is not
output when the \TEX/ source line was nonempty. For example, a line
of the \TEX/ file that contains only an index cross-reference entry
will not be copied. The |finish_line| routine is called just before
|get_line| inputs a new line, and just after a line break token has
been emitted during the output of translated \GO/ text.

@c
/* do this at the end of a line */
func finish_line() {
	if out_ptr>0 {
		flush_buffer(out_ptr,false,false)
	} else {
		for _, v := range buffer {
			if !unicode.IsSpace(v) {
				return
			}
		}
		flush_buffer(0,false,false)
	}
}

@ In particular, the |finish_line| procedure is called near the very
beginning of phase two. We initialize the output variables in a slightly
tricky way so that the first line of the output file will be
`\.{\\input cwebmac}'.

@<Set init...@>=
out_ptr=1
out_line=1
active_file=tex_file
out_buf[out_ptr]='c'
fmt.Fprint(active_file,"\\input cwebma") 

@ When we wish to append one character |c| to the output buffer, we write
`|out(c)|'; this will cause the buffer to be emptied if it was already
full.  If we want to append more than one character at once, we say
|out_str(s)|, where |s| is a string containing the characters.

A line break will occur at a space or after a single-nonletter
\TEX/ control sequence.

@ @c
func out(c rune) {
	if out_ptr>=out_buf_end {
		break_out()
	}
	out_ptr++
	out_buf[out_ptr]=c
}

@ @c
/* output characters from |s| to end of string */
func out_str(s string) {
	for _, v := range s {
		out(v)
	}
}

@ The |break_out| routine is called just before the output buffer is about
to overflow. To make this routine a little faster, we initialize position
0 of the output buffer to `\.\\'; this character isn't really output.

@<Set init...@>=
out_buf[0]='\\'

@ A long line is broken at a blank space or just before a backslash that isn't
preceded by another backslash. In the latter case, a |'%'| is output at
the break.

@ @c
/* finds a way to break the output line */
func break_out() {
	k:=out_ptr /* pointer into |out_buf| */
	for true {
		if k==0 {
			@<Print warning message, break the line, |return|@>
		}
		if out_buf[k]==' ' {
			flush_buffer(k,false,true) 
			return
		}
		kk := k
		k--
		if out_buf[kk]=='\\' && out_buf[k]!='\\' { /* we've decreased |k| */
			flush_buffer(k,true,true)
			return
		}
	}
}

@ We get to this section only in the unusual case that the entire output line
consists of a string of backslashes followed by a string of nonblank
non-backslashes. In such cases it is almost always safe to break the
line by putting a |'%'| just before the last character.

@<Print warning message...@>=
{
	fmt.Printf("\n! Line had to be broken (output l. %d):\n",out_line)
@.Line had to be broken@>
	fmt.Fprint(os.Stdout, string(out_buf[1:out_ptr]))
	fmt.Println()
	mark_harmless()
	flush_buffer(out_ptr-1,true,true)
	return
}

@ Here is a macro that outputs a section number in decimal notation.
The number to be converted by |out_section| is known to be less than
|def_flag|, so it cannot have more than five decimal digits.  If
the section is changed, we output `\.{\\*}' just after the number.

@c
func out_section(n int32) {
	out_str(fmt.Sprintf("%d",n))
	if changed_section[n] {
		out_str ("\\*")
@.\\*@>
	}
}

@ The |out_name| procedure is used to output an identifier or index
entry, enclosing it in braces.

@c
func out_name(p int32, quote_xalpha bool) {
	out('{')
	for _, v := range name_dir[p].name {
		if (v=='_' || v=='$') && quote_xalpha {
			out('\\')
		}
@.\\\$@>
@.\\\_@>
		out(v)
	}
	out('}')
}

@* Routines that copy \TEX/ material.
During phase two, we use subroutines |copy_limbo|, |copy_TeX|, and
|copy_comment| in place of the analogous |skip_limbo|, |skip_TeX|, and
|skip_comment| that were used in phase one. (Well, |copy_comment|
was actually written in such a way that it functions as |skip_comment|
in phase one.)

The |copy_limbo| routine, for example, takes \TEX/ material that is not
part of any section and transcribes it almost verbatim to the output file.
The use of `\.{@@}' signs is severely restricted in such material:
`\.{@@@@}' pairs are replaced by singletons; `\.{@@l}' and `\.{@@q}' and
`\.{@@s}' are interpreted.

@c
func copy_limbo() {
	for true {
		if loc >= len(buffer) {
			finish_line()
			if !get_line() {
				return
			}
		}
		for ;loc < len(buffer) && buffer[loc]!='@@'; loc++ {
			out(buffer[loc])
		}
		l := loc
		loc++
		if l<len(buffer) {
			c:=buffer[loc]
			loc++
			if ccode[c]==new_section {
				break
			}
			switch ccode[c] {
				case '@@': 
					out('@@')
				case noop: 
					skip_restricted()
				case format_code: 
					if get_next()==identifier {
						get_next()
					}
					if loc>=len(buffer) {
						get_line() /* avoid blank lines in output */
					}
					/* the operands of \.{@@s} are ignored on this pass */
				default: 
					err_print("! Double @@ should be used in limbo")
@.Double @@ should be used...@>
					out('@@')
			}
		}
	}
}

@ The |copy_TeX| routine processes the \TEX/ code at the beginning of a
section; for example, the words you are now reading were copied in this
way. It returns the next control code or `\.{\v}' found in the input.
We don't copy spaces or tab marks into the beginning of a line. This
makes the test for empty lines in |finish_line| work.

@ @f copy_TeX TeX
@c
func copy_TeX() rune {
	for true {
		if loc>=len(buffer) {
			finish_line()
			if !get_line() {
				return new_section
			}
		}
		c := buffer[loc]
		loc++
		for c!='|' && c!='@@'{
			out(c)
			if out_ptr==1 && unicode.IsSpace(c) {
				out_ptr--
			}
			if loc == len(buffer) {
				break
			} 	
			c = buffer[loc]
			loc++
		}
		if c=='|' {
			return '|'
		}
		if c =='@@' && len(buffer)==1 {
			return new_section
		}
		if loc<len(buffer) {
			l := loc
			loc++
			return ccode[buffer[l]]
		}
	}
	return 0
}

@ The |copy_comment| function issues a warning if more braces are opened than
closed, and in the case of a more serious error it supplies enough
braces to keep \TEX/ from complaining about unbalanced braces.
Instead of copying the \TEX/ material
into the output buffer, this function copies it into the token memory
(in phase two only).
The function |app_tok(t)| is used to append token |t| to the current
token list..

@c
func app_tok(c rune){
	tok_mem = append(tok_mem, c)
}

@ @c
/* copies \TEX/ code in comments */
func copy_comment(
	is_long_comment bool, /* is this a traditional \GO/ comment? */
	bal int /* brace balance */) int {
	for true {
		if loc>=len(buffer) {
			if is_long_comment {
				if !get_line() {
					err_print("! Input ended in mid-comment")
@.Input ended in mid-comment@>
					loc=1 
					goto done
				}
			} else {
				if bal>1 {
					err_print("! Missing } in comment")
@.Missing \} in comment@>
				}
				goto done
			}
		}
		c:=buffer[loc]
		loc++
		if c=='|' {
			return bal
		}
		if is_long_comment {
			@<Check for end of comment@>
		}
		if phase==2 {
			if c>0177 {
				app_tok(quoted_char)
			}
			app_tok(c)
		}
		@<Copy special things when |c=='@@', '\\'|@>
		if c=='{' {
			bal++
		} else if c=='}' {
			if bal>1 {
				bal--
			} else {
				err_print("! Extra } in comment")
@.Extra \} in comment@>
				if phase==2 {
					tok_mem = tok_mem[:len(tok_mem)-1]
				}
			}
		}
	}
done:
	@<Clear |bal| and |return|@>
}

@ @<Check for end of comment@>=
if c=='*' && loc < len(buffer) && buffer[loc]=='/' {
	loc++
	if bal>1 { 
		err_print("! Missing } in comment")
@.Missing \} in comment@>
	}
	goto done
}

@ @<Copy special things when |c=='@@'...@>=
if c=='@@' {
	l := loc
	loc++
	if l < len(buffer) && buffer[l]!='@@' {
		err_print("! Illegal use of @@ in comment")
@.Illegal use of @@...@>
		loc-=2 
		if phase==2 {
			tok_mem[len(tok_mem)-1]=' '
		}
		goto done
	}
} else if c=='\\' && loc < len(buffer) && buffer[loc]!='@@' {
	if phase==2 {
		app_tok(buffer[loc])
	}
	loc++
}

@ We output
enough right braces to keep \TEX/ happy.

@<Clear |bal|...@>=
if phase==2 {
	for bal--; bal>=0; bal-- {
		app_tok('}')
	}
}
return 0

@** Parsing.
The most intricate part of \.{GOWEAVE} is its mechanism for converting
\GO/-like code into \TEX/ code, and we might as well plunge into this
aspect of the program now. A ``bottom up'' approach is used to parse the
\GO/-like material, since \.{GOWEAVE} must deal with fragmentary
constructions whose overall ``part of speech'' is not known.

At the lowest level, the input is represented as a sequence of entities
that we shall call {\it scraps}, where each scrap of information consists
of two parts, its {\it category} and its {\it translation}. The category
is essentially a syntactic class, and the translation is a token list that
represents \TEX/ code. Rules of syntax and semantics tell us how to
combine adjacent scraps into larger ones, and if we are lucky an entire
\GO/ text that starts out as hundreds of small scraps will join
together into one gigantic scrap whose translation is the desired \TEX/
code. If we are unlucky, we will be left with several scraps that don't
combine; their translations will simply be output, one by one.

The combination rules are given as context-sensitive productions that are
applied from left to right. Suppose that we are currently working on the
sequence of scraps $s_1\,s_2\ldots s_n$. We try first to find the longest
production that applies to an initial substring $s_1\,s_2\ldots\,$; but if
no such productions exist, we try to find the longest production
applicable to the next substring $s_2\,s_3\ldots\,$; and if that fails, we
try to match $s_3\,s_4\ldots\,$, etc.

A production applies if the category codes have a given pattern. For
example, one of the productions (see rule~3) is
$$\hbox{|exp| }\left\{\matrix{\hbox{|binop|}\cr\hbox{|ubinop|}}\right\}
\hbox{ |exp| }\RA\hbox{ |exp|}$$
and it means that three consecutive scraps whose respective categories are
|exp|, |binop| (or |ubinop|),
and |exp| are converted to one scrap whose category
is |exp|.  The translations of the original
scraps are simply concatenated.  The case of
$$\hbox{|exp| |comma| |exp| $\RA$ |exp|} \hskip4emE_1C\,\\{opt}9\,E_2$$
(rule 4) is only slightly more complicated:
Here the resulting |exp| translation
consists not only of the three original translations, but also of the
tokens |opt| and 9 between the translations of the
|comma| and the following |exp|.
In the \TEX/ file, this will specify an optional line break after the
comma, with penalty 90.

At each opportunity the longest possible production is applied.  For
example, if the current sequence of scraps is |int_like| |cast|
|lbrace|, rule 31 is applied; but if the sequence is |int_like| |cast|
followed by anything other than |lbrace|, rule 32 takes effect.

Translation rules such as `$E_1C\,\\{opt}9\,E_2$' above use subscripts
to distinguish between translations of scraps whose categories have the
same initial letter; these subscripts are assigned from left to right.

@ Here is a list of the category codes that scraps can have.
(A few others, like |int_like|, have already been defined; the
|cat_name| array contains a complete list.)

@<Constants@>=
exp rune = 1 /* denotes an expression, including perhaps a single identifier */
unop rune = 2 /* denotes a unary operator */
binop rune = 3 /* denotes a binary operator */
ubinop rune = 4
	/* denotes an operator that can be unary or binary, depending on context */
cast rune = 5 /* denotes a cast */
question rune = 6 /* denotes a question mark and possibly the expressions flanking it */
lbrace rune = 7 /* denotes a left brace */
rbrace rune = 8 /* denotes a right brace */
decl_head rune = 9 /* denotes an incomplete declaration */
comma rune = 10 /* denotes a comma */
lpar rune = 11 /* denotes a left parenthesis or left bracket */
rpar rune = 12 /* denotes a right parenthesis or right bracket */
prelangle rune = 13 /* denotes `$<$' before we know what it is */
prerangle rune = 14 /* denotes `$>$' before we know what it is */
langle rune = 15 /* denotes `$<$' when it's used as angle bracket in a template */
colcol rune = 18 /* denotes `::' */
base rune = 19 /* denotes a colon that introduces a base specifier */
decl rune = 20 /* denotes a complete declaration */
struct_head rune = 21 /* denotes the beginning of a structure specifier */
stmt rune = 23 /* denotes a complete statement */
function rune = 24 /* denotes a complete function */
fn_decl rune = 25 /* denotes a function declarator */
semi rune = 27 /* denotes a semicolon */
colon rune = 28 /* denotes a colon */
tag rune = 29 /* denotes a statement label */
if_head rune = 30 /* denotes the beginning of a compound conditional */
else_head rune = 31 /* denotes a prefix for a compound statement */
if_clause rune = 32 /* pending \.{if} together with a condition */
lproc rune = 35 /* begins a preprocessor command */
rproc rune = 36 /* ends a preprocessor command */
insert rune = 37 /* a scrap that gets combined with its neighbor */
section_scrap rune = 38 /* section name */
dead rune = 39 /* scrap that won't combine */
ftemplate rune = 59 /* \\{make\_pair} */
new_exp rune = 60 /* \&{new} and a following type identifier */
begin_arg rune = 61 /* \.{@@[} */
end_arg rune = 62 /* \.{@@]} */

@ @<Glo...@>=
var cat_name[256]string

@ @<Set in...@>=
	for cat_index:=0;cat_index<255;cat_index++ {
		cat_name[cat_index] = "UNKNOWN"
	}
@.UNKNOWN@>
		cat_name[exp]="exp"
		cat_name[unop]="unop"
		cat_name[binop]="binop"
		cat_name[ubinop]="ubinop"
		cat_name[cast]="cast"
		cat_name[question]="?"
		cat_name[lbrace]="{"@q}@>
		cat_name[rbrace]=@q{@>"}"
		cat_name[decl_head]="decl_head"
		cat_name[comma]=","
		cat_name[lpar]="("
		cat_name[rpar]=")"
		cat_name[prelangle]="<"
		cat_name[prerangle]=">"
		cat_name[langle]="\\<"
		cat_name[colcol]="::"
		cat_name[base]="\\:"
		cat_name[decl]="decl"
		cat_name[struct_head]="struct_head"
		cat_name[alfop]="alfop"
		cat_name[stmt]="stmt"
		cat_name[function]="function"
		cat_name[fn_decl]="fn_decl"
		cat_name[else_like]="else_like"
		cat_name[semi]=";"
		cat_name[colon]=":"
		cat_name[tag]="tag"
		cat_name[if_head]="if_head"
		cat_name[else_head]="else_head"
		cat_name[if_clause]="if()"
		cat_name[lproc]="#{"@q}@>
		cat_name[rproc]=@q{@>"#}"
		cat_name[insert]="insert"
		cat_name[section_scrap]="section"
		cat_name[dead]="@@d"
		cat_name[public_like]="public"
		cat_name[operator_like]="operator"
		cat_name[new_like]="new"
		cat_name[catch_like]="catch"
		cat_name[for_like]="for"
		cat_name[do_like]="do"
		cat_name[if_like]="if"
		cat_name[delete_like]="delete"
		cat_name[raw_ubin]="ubinop?"
		cat_name[const_like]="const"
		cat_name[raw_int]="raw"
		cat_name[int_like]="int"
		cat_name[case_like]="case"
		cat_name[sizeof_like]="sizeof"
		cat_name[struct_like]="struct"
		cat_name[typedef_like]="typedef"
		cat_name[define_like]="define"
		cat_name[template_like]="template"
		cat_name[ftemplate]="ftemplate"
		cat_name[new_exp]="new_exp"
		cat_name[begin_arg]="@@["@q]@>
		cat_name[end_arg]=@q[@>"@@]"
		cat_name[0]="zero"

@ This code allows \.{GOWEAVE} to display its parsing steps.

@c
/* symbolic printout of a category */
func print_cat(c int32) {
	fmt.Print(cat_name[c])
}

@ The token lists for translated \TEX/ output contain some special control
symbols as well as ordinary characters. These control symbols are
interpreted by \.{GOWEAVE} before they are written to the output file.

\yskip\hang |break_space| denotes an optional line break or an en space;

\yskip\hang |force| denotes a line break;

\yskip\hang |big_force| denotes a line break with additional vertical space;

\yskip\hang |preproc_line| denotes that the line will be printed flush left;

\yskip\hang |opt| denotes an optional line break (with the continuation
line indented two ems with respect to the normal starting position)---this
code is followed by an integer |n|, and the break will occur with penalty
$10n$;

\yskip\hang |backup| denotes a backspace of one em;

\yskip\hang |cancel| obliterates any |break_space|, |opt|, |force|, or
|big_force| tokens that immediately precede or follow it and also cancels any
|backup| tokens that follow it;

\yskip\hang |indent| causes future lines to be indented one more em;

\yskip\hang |outdent| causes future lines to be indented one less em.

\yskip\noindent All of these tokens are removed from the \TEX/ output that
comes from \GO/ text between \pb\ signs; |break_space| and |force| and
|big_force| become single spaces in this mode. The translation of other
\GO/ texts results in \TEX/ control sequences \.{\\1}, \.{\\2},
\.{\\3}, \.{\\4}, \.{\\5}, \.{\\6}, \.{\\7}, \.{\\8}
corresponding respectively to
|indent|, |outdent|, |opt|, |backup|, |break_space|, |force|,
|big_force| and |preproc_line|.
However, a sequence of consecutive `\.\ ', |break_space|,
|force|, and/or |big_force| tokens is first replaced by a single token
(the maximum of the given ones).

The token |math_rel| will be translated into
\.{\\MRL\{}, and it will get a matching \.\} later.
Other control sequences in the \TEX/ output will be
`\.{\\\\\{}$\,\ldots\,$\.\}'
surrounding identifiers, `\.{\\\&\{}$\,\ldots\,$\.\}' surrounding
reserved words, `\.{\\.\{}$\,\ldots\,$\.\}' surrounding strings,
`\.{\\C\{}$\,\ldots\,$\.\}$\,$|force|' surrounding comments, and
`\.{\\X$n$:}$\,\ldots\,$\.{\\X}' surrounding section names, where
|n| is the section number.

@<Constants@>=
math_rel = 0206
big_cancel = 0210 /* like |cancel|, also overrides spaces */
cancel = 0211 /* overrides |backup|, |break_space|, |force|, |big_force| */
indent = 0212 /* one more tab (\.{\\1}) */
outdent = 0213 /* one less tab (\.{\\2}) */
opt = 0214 /* optional break in mid-statement (\.{\\3}) */
backup = 0215 /* stick out one unit to the left (\.{\\4}) */
break_space = 0216 /* optional break between statements (\.{\\5}) */
force = 0217 /* forced break between statements (\.{\\6}) */
big_force = 0220 /* forced break with additional space (\.{\\7}) */
preproc_line = 0221 /* begin line without indentation (\.{\\8}) */
@^high-bit character handling@>
quoted_char = 0222
				/* introduces a character token in the range |0200|--|0377| */
end_translation = 0223 /* special sentinel token at end of list */
inserted = 0224 /* sentinel to mark translations of inserts */
qualifier = 0225 /* introduces an explicit namespace qualifier */

@ The raw input is converted into scraps according to the following table,
which gives category codes followed by the translations.
\def\stars {\.{**}}%
The symbol `\stars' stands for `\.{\\\&\{{\rm identifier}\}}',
i.e., the identifier itself treated as a reserved word.
The right-hand column is the so-called |mathness|, which is explained
further below.

An identifier |c| of length 1 is translated as \.{\\\v c} instead of
as \.{\\\\\{c\}}. An identifier \.{CAPS} in all caps is translated as
\.{\\.\{CAPS\}} instead of as \.{\\\\\{CAPS\}}. An identifier that has
become a reserved word via |typedef| is translated with \.{\\\&} replacing
\.{\\\\} and |raw_int| replacing |exp|.

A string of length greater than 20 is broken into pieces of size at most~20
with discretionary breaks in between.

\yskip\halign{\quad#\hfil&\quad#\hfil&\quad\hfil#\hfil\cr
\.{!=}&|binop|: \.{\\I}&yes\cr
\.{<=}&|binop|: \.{\\Z}&yes\cr
\.{>=}&|binop|: \.{\\G}&yes\cr
\.{==}&|binop|: \.{\\E}&yes\cr
\.{\&\&}&|binop|: \.{\\W}&yes\cr
\.{\v\v}&|binop|: \.{\\V}&yes\cr
\.{++}&|unop|: \.{\\PP}&yes\cr
\.{--}&|unop|: \.{\\MM}&yes\cr
\.{->}&|binop|: \.{\\MG}&yes\cr
\.{>>}&|binop|: \.{\\GG}&yes\cr
\.{<<}&|binop|: \.{\\LL}&yes\cr
\.{::}&|colcol|: \.{\\DC}&maybe\cr
\.{.*}&|binop|: \.{\\PA}&yes\cr
\.{->*}&|binop|: \.{\\MGA}&yes\cr
\.{...}&|raw_int|: \.{\\,\\ldots\\,}&yes\cr
\."string\."&|exp|: \.{\\.\{}string with special characters quoted\.\}&maybe\cr
\.{@@=}string\.{@@>}&|exp|: \.{\\vb\{}string with special characters
	quoted\.\}&maybe\cr
\.{@@'7'}&|exp|: \.{\\.\{@@'7'\}}&maybe\cr
\.{077} or \.{\\77}&|exp|: \.{\\T\{\\\~77\}}&maybe\cr
\.{0x7f}&|exp|: \.{\\T\{\\\^7f\}}&maybe\cr
\.{77}&|exp|: \.{\\T\{77\}}&maybe\cr
\.{77L}&|exp|: \.{\\T\{77\\\$L\}}&maybe\cr
\.{0.1E5}&|exp|: \.{\\T\{0.1\\\_5\}}&maybe\cr
\.+&|ubinop|: \.+&yes\cr
\.-&|ubinop|: \.-&yes\cr
\.*&|raw_ubin|: \.*&yes\cr
\./&|binop|: \./&yes\cr
\.<&|prelangle|: \.{\\langle}&yes\cr
\.=&|binop|: \.{\\K}&yes\cr
\.>&|prerangle|: \.{\\rangle}&yes\cr
\..&|binop|: \..&yes\cr
\.{\v}&|binop|: \.{\\OR}&yes\cr
\.\^&|binop|: \.{\\XOR}&yes\cr
\.\%&|binop|: \.{\\MOD}&yes\cr
\.?&|question|: \.{\\?}&yes\cr
\.!&|unop|: \.{\\R}&yes\cr
\.\~&|unop|: \.{\\CM}&yes\cr
\.\&&|raw_ubin|: \.{\\AND}&yes\cr
\.(&|lpar|: \.(&maybe\cr
\.[&|lpar|: \.[&maybe\cr
\.)&|rpar|: \.)&maybe\cr
\.]&|rpar|: \.]&maybe\cr
\.\{&|lbrace|: \.\{&yes\cr
\.\}&|lbrace|: \.\}&yes\cr
\.,&|comma|: \.,&yes\cr
\.;&|semi|: \.;&maybe\cr
\.:&|colon|: \.:&no\cr
\.\# (within line)&|ubinop|: \.{\\\#}&yes\cr
\.\# (at beginning)&|lproc|:  |force| |preproc_line| \.{\\\#}&no\cr
end of \.\# line&|rproc|:  |force|&no\cr
identifier&|exp|: \.{\\\\\{}identifier with underlines and
					   dollar signs quoted\.\}&maybe\cr
\.{and}&|alfop|: \stars&yes\cr
\.{and\_eq}&|alfop|: \stars&yes\cr
\.{asm}&|sizeof_like|: \stars&maybe\cr
\.{auto}&|int_like|: \stars&maybe\cr
\.{bitand}&|alfop|: \stars&yes\cr
\.{bitor}&|alfop|: \stars&yes\cr
\.{bool}&|raw_int|: \stars&maybe\cr
\.{break}&|case_like|: \stars&maybe\cr
\.{case}&|case_like|: \stars&maybe\cr
\.{catch}&|catch_like|: \stars&maybe\cr
\.{char}&|raw_int|: \stars&maybe\cr
\.{class}&|struct_like|: \stars&maybe\cr
\.{clock\_t}&|raw_int|: \stars&maybe\cr
\.{compl}&|alfop|: \stars&yes\cr
\.{const}&|const_like|: \stars&maybe\cr
\.{const\_cast}&|raw_int|: \stars&maybe\cr
\.{continue}&|case_like|: \stars&maybe\cr
\.{default}&|case_like|: \stars&maybe\cr
\.{define}&|define_like|: \stars&maybe\cr
\.{defined}&|sizeof_like|: \stars&maybe\cr
\.{delete}&|delete_like|: \stars&maybe\cr
\.{div\_t}&|raw_int|: \stars&maybe\cr
\.{do}&|do_like|: \stars&maybe\cr
\.{double}&|raw_int|: \stars&maybe\cr
\.{dynamic\_cast}&|raw_int|: \stars&maybe\cr
\.{elif}&|if_like|: \stars&maybe\cr
\.{else}&|else_like|: \stars&maybe\cr
\.{endif}&|if_like|: \stars&maybe\cr
\.{enum}&|struct_like|: \stars&maybe\cr
\.{error}&|if_like|: \stars&maybe\cr
\.{explicit}&|int_like|: \stars&maybe\cr
\.{export}&|int_like|: \stars&maybe\cr
\.{extern}&|int_like|: \stars&maybe\cr
\.{FILE}&|raw_int|: \stars&maybe\cr
\.{float}&|raw_int|: \stars&maybe\cr
\.{for}&|for_like|: \stars&maybe\cr
\.{fpos\_t}&|raw_int|: \stars&maybe\cr
\.{friend}&|int_like|: \stars&maybe\cr
\.{goto}&|case_like|: \stars&maybe\cr
\.{if}&|if_like|: \stars&maybe\cr
\.{ifdef}&|if_like|: \stars&maybe\cr
\.{ifndef}&|if_like|: \stars&maybe\cr
\.{include}&|if_like|: \stars&maybe\cr
\.{inline}&|int_like|: \stars&maybe\cr
\.{int}&|raw_int|: \stars&maybe\cr
\.{jmp\_buf}&|raw_int|: \stars&maybe\cr
\.{ldiv\_t}&|raw_int|: \stars&maybe\cr
\.{line}&|if_like|: \stars&maybe\cr
\.{long}&|raw_int|: \stars&maybe\cr
\.{make\_pair}&|ftemplate|: \.{\\\\\{make\\\_pair\}}&maybe\cr
\.{mutable}&|int_like|: \stars&maybe\cr
\.{namespace}&|struct_like|: \stars&maybe\cr
\.{new}&|new_like|: \stars&maybe\cr
\.{not}&|alfop|: \stars&yes\cr
\.{not\_eq}&|alfop|: \stars&yes\cr
\.{NULL}&|exp|: \.{\\NULL}&yes\cr
\.{offsetof}&|raw_int|: \stars&maybe\cr
\.{operator}&|operator_like|: \stars&maybe\cr
\.{or}&|alfop|: \stars&yes\cr
\.{or\_eq}&|alfop|: \stars&yes\cr
\.{pragma}&|if_like|: \stars&maybe\cr
\.{private}&|public_like|: \stars&maybe\cr
\.{protected}&|public_like|: \stars&maybe\cr
\.{ptrdiff\_t}&|raw_int|: \stars&maybe\cr
\.{public}&|public_like|: \stars&maybe\cr
\.{register}&|int_like|: \stars&maybe\cr
\.{reinterpret\_cast}&|raw_int|: \stars&maybe\cr
\.{return}&|case_like|: \stars&maybe\cr
\.{short}&|raw_int|: \stars&maybe\cr
\.{sig\_atomic\_t}&|raw_int|: \stars&maybe\cr
\.{signed}&|raw_int|: \stars&maybe\cr
\.{size\_t}&|raw_int|: \stars&maybe\cr
\.{sizeof}&|sizeof_like|: \stars&maybe\cr
\.{static}&|int_like|: \stars&maybe\cr
\.{static\_cast}&|raw_int|: \stars&maybe\cr
\.{struct}&|struct_like|: \stars&maybe\cr
\.{switch}&|for_like|: \stars&maybe\cr
\.{template}&|template_like|: \stars&maybe\cr
\.{TeX}&|exp|: \.{\\TeX}&yes\cr
\.{this}&|exp|: \.{\\this}&yes\cr
\.{throw}&|case_like|: \stars&maybe\cr
\.{time\_t}&|raw_int|: \stars&maybe\cr
\.{try}&|else_like|: \stars&maybe\cr
\.{typedef}&|typedef_like|: \stars&maybe\cr
\.{typeid}&|raw_int|: \stars&maybe\cr
\.{typename}&|struct_like|: \stars&maybe\cr
\.{undef}&|if_like|: \stars&maybe\cr
\.{union}&|struct_like|: \stars&maybe\cr
\.{unsigned}&|raw_int|: \stars&maybe\cr
\.{using}&|int_like|: \stars&maybe\cr
\.{va\_dcl}&|decl|: \stars&maybe\cr
\.{va\_list}&|raw_int|: \stars&maybe\cr
\.{virtual}&|int_like|: \stars&maybe\cr
\.{void}&|raw_int|: \stars&maybe\cr
\.{volatile}&|const_like|: \stars&maybe\cr
\.{wchar\_t}&|raw_int|: \stars&maybe\cr
\.{while}&|for_like|: \stars&maybe\cr
\.{xor}&|alfop|: \stars&yes\cr
\.{xor\_eq}&|alfop|: \stars&yes\cr
\.{@@,}&|insert|: \.{\\,}&maybe\cr
\.{@@\v}&|insert|:  |opt| \.0&maybe\cr
\.{@@/}&|insert|:  |force|&no\cr
\.{@@\#}&|insert|:  |big_force|&no\cr
\.{@@+}&|insert|:  |big_cancel| \.{\{\}} |break_space|
	\.{\{\}} |big_cancel|&no\cr
\.{@@;}&|semi|: &maybe\cr
\.{@@[@q]@>}&|begin_arg|: &maybe\cr
\.{@q[@>@@]}&|end_arg|: &maybe\cr
\.{@@\&}&|insert|: \.{\\J}&maybe\cr
\.{@@h}&|insert|: |force| \.{\\ATH} |force|&no\cr
\.{@@<}\thinspace section name\thinspace\.{@@>}&|section_scrap|:
 \.{\\X}$n$\.:translated section name\.{\\X}&maybe\cr
\.{@@(@q)@>}\thinspace section name\thinspace\.{@@>}&|section_scrap|:
 \.{\\X}$n$\.{:\\.\{}section name with special characters
			quoted\.{\ \}\\X}&maybe\cr
\.{/*}comment\.{*/}&|insert|: |cancel|
			\.{\\C\{}translated comment\.\} |force|&no\cr
\.{//}comment&|insert|: |cancel|
			\.{\\SHC\{}translated comment\.\} |force|&no\cr
}

\smallskip
The construction \.{@@t}\thinspace stuff\/\thinspace\.{@@>} contributes
\.{\\hbox\{}\thinspace  stuff\/\thinspace\.\} to the following scrap.

@i prod.w

@* Implementing the productions.
More specifically, a scrap is a structure consisting of a category
|cat| and a |trans|, which points to the translation in
|tok_start|.  When \GO/ text is to be processed with the grammar above,
we form an array |scrap_info| containing the initial scraps.
Our production rules have the nice property that the right-hand side is never
longer than the left-hand side. Therefore it is convenient to use sequential
allocation for the current sequence of scraps. Five pointers are used to
manage the parsing:

\yskip\hang |pp| is a pointer into |scrap_info|.  We will try to match
the category codes |scrap_info[pp].cat,@,@,scrap_info[pp+1].cat|$,\,\,\ldots\,$
to the left-hand sides of productions.

\yskip\hang |scrap_base|, |lo_ptr|, |hi_ptr|, and |scrap_ptr| are such that
the current sequence of scraps appears in positions |scrap_base| through
|lo_ptr| and |hi_ptr| through |scrap_ptr|, inclusive, in the |cat| and
|trans| arrays. Scraps located between |scrap_base| and |lo_ptr| have
been examined, while those in positions |>=hi_ptr| have not yet been
looked at by the parsing process.

\yskip\noindent Initially |scrap_ptr| is set to the position of the final
scrap to be parsed, and it doesn't change its value. The parsing process
makes sure that |lo_ptr>=pp+3|, since productions have as many as four terms,
by moving scraps from |hi_ptr| to |lo_ptr|. If there are
fewer than |pp+3| scraps left, the positions up to |pp+3| are filled with
blanks that will not match in any productions. Parsing stops when
|pp==lo_ptr+1| and |hi_ptr==scrap_ptr+1|.

Since the |scrap| structure will later be used for other purposes, we
declare its second element as a union.

@<Type...@>=
type trans struct {
	Trans int32
	@<Rest of |trans| struct@>
}

type scrap struct {
	cat int32
	mathness int32
	trans_plus trans
}

@ @<Global...@>=
var scrap_info [max_scraps]scrap /* memory array for scraps */
var pp int32 /* current position for reducing productions */
var scrap_base int32 /* beginning of the current scrap sequence */
var scrap_ptr int32 /* ending of the current scrap sequence */
var lo_ptr int32 /* last scrap that has been examined */
var hi_ptr int32 /* first scrap that has not been examined */
var max_scr_ptr int32 /* largest value assumed by |scrap_ptr| */

@ @<Set init...@>=
scrap_base=1
max_scr_ptr=0
scrap_ptr=0

@ Token lists in |@!tok_mem| are composed of the following kinds of
items for \TEX/ output.

\yskip\item{$\bullet$}Character codes and special codes like |force| and
|math_rel| represent themselves;

\item{$\bullet$}|id_flag+p| represents \.{\\\\\{{\rm identifier $p$}\}};

\item{$\bullet$}|res_flag+p| represents \.{\\\&\{{\rm identifier $p$}\}};

\item{$\bullet$}|section_flag+p| represents section name |p|;

\item{$\bullet$}|tok_flag+p| represents token list number |p|;

\item{$\bullet$}|inner_tok_flag+p| represents token list number |p|, to be
translated without line-break controls.

@<Constants@>=
id_flag rune = 10240 /* signifies an identifier */
res_flag rune = 2*id_flag /* signifies a reserved word */
section_flag rune = 3*id_flag /* signifies a section name */
tok_flag rune = 4*id_flag /* signifies a token list */
inner_tok_flag rune = 5*id_flag /* signifies a token list in `\pb' */

@ @<Print token |r|...@>=
switch (r) {
	case math_rel: 
		printf("\\mathrel{"@q}@>)
	case big_cancel: 
		printf("[ccancel]")
	case cancel: 
		printf("[cancel]")
	case indent: 
		printf("[indent]")
	case outdent: 
		printf("[outdent]")
	case backup: 
		printf("[backup]")
	case opt: 
		printf("[opt]")
	case break_space: 
		printf("[break]")
	case force: 
		printf("[force]")
	case big_force: 
		printf("[fforce]")
	case preproc_line: 
		printf("[preproc]")
	case quoted_char: 
		j++
		printf("[%o]",(unsigned)*j)
	case end_translation: 
		printf("[quit]")
	case inserted: 
		printf("[inserted]")
	default: 
		putxchar(r)
}


@ The production rules listed above are embedded directly into \.{GOWEAVE},
since it is easier to do this than to write an interpretive system
that would handle production systems in general. Several macros are defined
here so that the program for each production is fairly short.

All of our productions conform to the general notion that some |k|
consecutive scraps starting at some position |j| are to be replaced by a
single scrap of some category |c| whose translation is composed from the
translations of the disappearing scraps. After this production has been
applied, the production pointer |pp| should change by an amount |d|. Such
a production can be represented by the quadruple |(j,k,c,d)|. For example,
the production `|exp@,comma@,exp| $\RA$ |exp|' would be represented by
`|(pp,3,exp,-2)|'; in this case the pointer |pp| should decrease by 2
after the production has been applied, because some productions with
|exp| in their second or third positions might now match,
but no productions have
|exp| in the fourth position of their left-hand sides. Note that
the value of |d| is determined by the whole collection of productions, not
by an individual one.
The determination of |d| has been
done by hand in each case, based on the full set of productions but not on
the grammar of \GO/ or on the rules for constructing the initial
scraps.

We also attach a serial number to each production, so that additional
information is available when debugging. For example, the program below
contains the statement `|reduce(pp,3,exp,-2,4)|' when it implements
the production just mentioned.

Before calling |reduce|, the program should have appended the tokens of
the new translation to the |tok_mem| array. We commonly want to append
copies of several existing translations, and macros are defined to
simplify these common cases. For example, \\{app2}|(pp)| will append the
translations of two consecutive scraps, |scrap_info[pp].trans_plus.Trans| 
and |scrap_info[pp+1].trans_plus.Trans|, to
the current token list. If the entire new translation is formed in this
way, we write `|squash(j,k,c,d,n)|' instead of `|reduce(j,k,c,d,n)|'. For
example, `|squash(pp,3,exp,-2,3)|' is an abbreviation for `\\{app3}|(pp);
reduce(pp,3,exp,-2,3)|'.

A couple more words of explanation:
Both |big_app| and |app| append a token (while |big_app1| to |big_app4|
append the specified number of scrap translations) to the current token list.
The difference between |big_app| and |app| is simply that |big_app|
checks whether there can be a conflict between math and non-math
tokens, and intercalates a `\.{\$}' token if necessary.  When in
doubt what to use, use |big_app|.

The |mathness| is an attribute of scraps that says whether they are
to be printed in a math mode context or not.  It is separate from the
``part of speech'' (the |cat|) because to make each |cat| have
a fixed |mathness| (as in the original \.{WEAVE}) would multiply the
number of necessary production rules.

The low two bits (i.e. |mathness % 4|) control the left boundary.
(We need two bits because we allow cases |yes_math|, |no_math| and
|maybe_math|, which can go either way.)
The next two bits (i.e. |mathness / 4|) control the right boundary.
If we combine two scraps and the right boundary of the first has
a different mathness from the left boundary of the second, we
insert a \.{\$} in between.  Similarly, if at printing time some
irreducible scrap has a |yes_math| boundary the scrap gets preceded
or followed by a \.{\$}. The left boundary is |maybe_math| if and
only if the right boundary is.

The code below is an exact translation of the production rules into
\GO/, using such macros, and the reader should have no difficulty
understanding the format by comparing the code with the symbolic
productions as they were listed earlier.

@<Constants@>=
no_math rune = 2 /* should be in horizontal mode */
yes_math rune = 1 /* should be in math mode */
maybe_math rune = 0 /* works in either horizontal or math mode */

@ @c 
func big_app2(a rune) {
	big_app1(a)
	big_app1(a+1)
}

func big_app3(a rune) {
	big_app2(a)
	big_app1(a+2)
}

func big_app4(a rune) {
	big_app3(a)
	big_app1(a+3)
}

func app(a rune) {
	tok_mem = append(tok_mem, a)
}

func app1(a int32) {
	tok_mem = append(tok_mem, tok_flag+scrap_info[a].trans_plus.Trans)
}

@ @<Global...@>=
var cur_mathness int32
var init_mathness int32

@ @c
func app_str(s string) {
	for _, v := range s {
		app_tok(v)
	}
}

func big_app(a rune) {
	if a==' ' || (a>=big_cancel && a<=big_force) /* non-math token */ {
		if cur_mathness==maybe_math { 
			init_mathness=no_math
		} else if cur_mathness==yes_math { 
			app_str("{}$") 
		}
		cur_mathness=no_math
	} else {
		if cur_mathness==maybe_math { 
			init_mathness=yes_math
		} else if cur_mathness==no_math { 
			app_str("${}") 
		}
		cur_mathness=yes_math
	}
	app(a)
}

func big_app1(a int32) {
	switch scrap_info[a].mathness % 4 { /* left boundary */
	case no_math:
		if cur_mathness==maybe_math {
			init_mathness=no_math
		} else if (cur_mathness==yes_math) { 
			app_str("{}$") 
		}
		cur_mathness=scrap_info[a].mathness / 4 /* right boundary */
	case yes_math:
		if cur_mathness==maybe_math { 
			init_mathness=yes_math 
		} else if cur_mathness==no_math {
			app_str("${}")
		}
		cur_mathness=scrap_info[a].mathness / 4 /* right boundary */
	case maybe_math: /* no changes */
	}
	app(tok_flag+scrap_info[a].trans_plus.Trans)
}

@ Let us consider the big switch for productions now, before looking
at its context. We want to design the program so that this switch
works, so we might as well not keep ourselves in suspense about exactly what
code needs to be provided with a proper environment.

@<Match a production at |pp|, or increase |pp| if there is no match@>= {
	/* not a production with left side length 1 */	
	if scrap_info[pp+1].cat==end_arg && 
		scrap_info[pp].cat!=public_like && 
		scrap_info[pp].cat!=semi && 
		scrap_info[pp].cat!=prelangle && 
		scrap_info[pp].cat!=prerangle && 
		scrap_info[pp].cat!=template_like && 
		scrap_info[pp].cat!=new_like && 
		scrap_info[pp].cat!=new_exp && 
		scrap_info[pp].cat!=ftemplate && 
		scrap_info[pp].cat!=raw_ubin && 
		scrap_info[pp].cat!=const_like && 
		scrap_info[pp].cat!=raw_int && 
		scrap_info[pp].cat!=operator_like {
		if scrap_info[pp].cat==begin_arg { 
			squash(pp,2,exp,-2,124) 
		} else { 
			squash(pp,2,end_arg,-1,125) 
		}
	} else if (scrap_info[pp+1].cat==insert) { 
		squash(pp,2,scrap_info[pp].cat,-2,0)
	} else if (scrap_info[pp+2].cat==insert) { 
		squash(pp+1,2,scrap_info[pp+1].cat,-1,0)
	} else if (scrap_info[pp+3].cat==insert) { 
		squash(pp+2,2,scrap_info[pp+2].cat,0,0)
	} else {
		switch (scrap_info[pp].cat) {
			case exp: @<Cases for |exp|@>
			case lpar: @<Cases for |lpar|@>
			case unop: @<Cases for |unop|@>
			case ubinop: @<Cases for |ubinop|@>
			case binop: @<Cases for |binop|@>
			case cast: @<Cases for |cast|@>
			case sizeof_like: @<Cases for |sizeof_like|@>
			case int_like: @<Cases for |int_like|@>
			case public_like: @<Cases for |public_like|@>
			case colcol: @<Cases for |colcol|@>
			case decl_head: @<Cases for |decl_head|@>
			case decl: @<Cases for |decl|@>
			case base: @<Cases for |base|@>
			case struct_like: @<Cases for |struct_like|@>
			case struct_head: @<Cases for |struct_head|@>
			case fn_decl: @<Cases for |fn_decl|@>
			case function: @<Cases for |function|@>
			case lbrace: @<Cases for |lbrace|@>
			case if_like: @<Cases for |if_like|@>
			case else_like: @<Cases for |else_like|@>
			case else_head: @<Cases for |else_head|@>
			case if_clause: @<Cases for |if_clause|@>
			case if_head: @<Cases for |if_head|@>
			case do_like: @<Cases for |do_like|@>
			case case_like: @<Cases for |case_like|@>
			case catch_like: @<Cases for |catch_like|@>
			case tag: @<Cases for |tag|@>
			case stmt: @<Cases for |stmt|@>
			case semi: @<Cases for |semi|@>
			case lproc: @<Cases for |lproc|@>
			case section_scrap: @<Cases for |section_scrap|@>
			case insert: @<Cases for |insert|@>
			case prelangle: @<Cases for |prelangle|@>
			case prerangle: @<Cases for |prerangle|@>
			case langle: @<Cases for |langle|@>
			case template_like: @<Cases for |template_like|@>
			case new_like: @<Cases for |new_like|@>
			case new_exp: @<Cases for |new_exp|@>
			case ftemplate: @<Cases for |ftemplate|@>
			case for_like: @<Cases for |for_like|@>
			case raw_ubin: @<Cases for |raw_ubin|@>
			case const_like: @<Cases for |const_like|@>
			case raw_int: @<Cases for |raw_int|@>
			case operator_like: @<Cases for |operator_like|@>
			case typedef_like: @<Cases for |typedef_like|@>
			case delete_like: @<Cases for |delete_like|@>
			case question: @<Cases for |question|@>
		}
	}
	pp++ /* if no match was found, we move to the right */
}

@ In \GO/, new specifier names can be defined via |typedef|, and we want
to make the parser recognize future occurrences of the identifier thus
defined as specifiers.  This is done by the procedure |make_reserved|,
which changes the |ilk| of the relevant identifier.

We first need a procedure to recursively seek the first
identifier in a token list, because the identifier might
be enclosed in parentheses, as when one defines a function
returning a pointer.

If the first identifier found is a keyword like `\&{case}', we
return the special value |case_found|; this prevents underlining
of identifiers in case labels.

If the first identifier is the keyword `\&{operator}', we give up;
users who want to index definitions of overloaded \CPLUSPLUS/ operators
should say, for example, `\.{@@!@@\^\\\&\{operator\} \$+\{=\}\$@@>}' (or,
more properly alphebetized,
`\.{@@!@@:operator+=\}\{\\\&\{operator\} \$+\{=\}\$@@>}').

@<Constants@>=
no_ident_found int32 = -3 /* distinct from any identifier token */
case_found int32 = -2 /* likewise */
operator_found int32 = -1 /* likewise */

@ @c
func find_first_ident(p int32) int32 {
	for j:=tok_start[p]; j<tok_start[p+1]; j++ {
		r:=tok_mem[j]%id_flag /* remainder of token after the flag has been stripped off */
		switch tok_mem[j]/id_flag {
			case 2: /* |res_flag| */
				if name_dir[r].ilk==case_like {
					return case_found
				}
				if name_dir[r].ilk==operator_like { 
					return operator_found
				}
				if name_dir[r].ilk!=raw_int {
					break
				}
				fallthrough
			case 1: 
				return j
			case 4, 5: /* |tok_flag| or |inner_tok_flag| */
				if q:=find_first_ident(r); q!=no_ident_found {
					return q
				}
				fallthrough
			default:  /* char, |section_flag|, fall thru: move on to next token */
				if tok_mem[j]==inserted {
					return no_ident_found /* ignore inserts */
				} else if tok_mem[j]==qualifier { 
					j++ /* bypass namespace qualifier */
				}
		}
	}
	return no_ident_found
}

@ The scraps currently being parsed must be inspected for any
occurrence of the identifier that we're making reserved; hence
the |for| loop below.

@c
/* make the first identifier in |scrap_info[p].trans_plus.Trans| like |int| */
func make_reserved(p int32) {
	tok_loc:=find_first_ident(scrap_info[p].trans_plus.Trans)/* pointer to |tok_value| */
	if tok_loc<=operator_found {
		return /* this should not happen */
	}
	tok_value:=tok_mem[tok_loc] /* the name of this identifier, plus its flag*/
	for p<=scrap_ptr {
		if scrap_info[p].cat==exp {
			if tok_mem[tok_start[scrap_info[p].trans_plus.Trans]]==tok_value {
				scrap_info[p].cat=raw_int
				tok_mem[tok_start[scrap_info[p].trans_plus.Trans]]=tok_value%id_flag+res_flag
			}
		}
		if p==lo_ptr {
			p=hi_ptr 
		} else {
			p++
		}
	}
	name_dir[tok_value%id_flag].ilk=raw_int
	tok_mem[tok_loc]=tok_value%id_flag+res_flag
}

@ In the following situations we want to mark the occurrence of
an identifier as a definition: when |make_reserved| is just about to be
used; after a specifier, as in |char **argv|;
before a colon, as in \\{found}:; and in the declaration of a function,
as in \\{main}()$\{\ldots;\}$.  This is accomplished by the invocation
of |make_underlined| at appropriate times.  Notice that, in the declaration
of a function, we find out that the identifier is being defined only after
it has been swallowed up by an |exp|.

@c
/* underline the entry for the first identifier in |scrap_info[p].trans_plus.Trans| */
func make_underlined(p int32) {
	var tok_loc int32/* where the first identifier appears */
	if tok_loc=find_first_ident(scrap_info[p].trans_plus.Trans); tok_loc<=operator_found {
		return /* this happens, for example, in |case found:| */
	}
	xref_switch=def_flag
	underline_xref(tok_mem[tok_loc]%id_flag)
}

@ We cannot use |new_xref| to underline a cross-reference at this point
because this would just make a new cross-reference at the end of the list.
We actually have to search through the list for the existing
cross-reference.

@ @c
func underline_xref(p int32) {
	q:=name_dir[p].xref /* pointer to cross-reference being examined */
	if flags['x']==false {
		return
	}
	m:=section_count+xref_switch /* cross-reference value to be installed */
	for q != 0 {
		n:=xmem[q].num /* cross-reference value being examined */
		if n==m {
			return
		} else if m==n+def_flag {
			xmem[q].num=m
			return
		} else if n>=def_flag && n<m {
			break
		}
		q=xmem[q].xlink
	}
	@<Insert new cross-reference at |q|, not at beginning of list@>
}

@ We get to this section only when the identifier is one letter long,
so it didn't get a non-underlined entry during phase one.  But it may
have got some explicitly underlined entries in later sections, so in order
to preserve the numerical order of the entries in the index, we have
to insert the new cross-reference not at the beginning of the list
(namely, at |name_dir[p].xref|), but rather right before |q|.

@<Insert new cross-reference at |q|...@>=
	append_xref(0) /* this number doesn't matter */
	xmem[len(xmem)-1].xlink=name_dir[p].xref
	r:=int32(len(xmem)-1) /* temporary pointer for permuting cross-references */
	name_dir[p].xref=r
	for xmem[r].xlink!=q {
		xmem[r].num=xmem[xmem[r].xlink].num
		r=xmem[r].xlink
	}
	xmem[r].num=m /* everything from |q| on is left undisturbed */

@ Now comes the code that tries to match each production starting
with a particular type of scrap. Whenever a match is discovered,
the |squash| or |reduce| macro will cause the appropriate action
to be performed, followed by |goto found|.

@<Cases for |exp|@>=
if (scrap_info[pp+1].cat==lbrace || 
	scrap_info[pp+1].cat==int_like || 
	scrap_info[pp+1].cat==decl) {
	make_underlined(pp)
	big_app1(pp)
	big_app(indent)
	app(indent)
	reduce(pp,1,fn_decl,0,1)
} else if scrap_info[pp+1].cat==unop { 
		squash(pp,2,exp,-2,2) 
} else if (scrap_info[pp+1].cat==binop || 
		scrap_info[pp+1].cat==ubinop) && 
		scrap_info[pp+2].cat==exp {
	squash(pp,3,exp,-2,3)
} else if scrap_info[pp+1].cat==comma && 
		scrap_info[pp+2].cat==exp {
	big_app2(pp)
	app(opt)
	app('9')
	big_app1(pp+2)
	reduce(pp,3,exp,-2,4)
} else if scrap_info[pp+1].cat==lpar && 
		scrap_info[pp+2].cat==rpar && 
		scrap_info[pp+3].cat==colon {
	squash(pp+3,1,base,0,5)
} else if scrap_info[pp+1].cat==cast && 
		scrap_info[pp+2].cat==colon {
	squash(pp+2,1,base,0,5)
} else if scrap_info[pp+1].cat==semi {
		squash(pp,2,stmt,-1,6)
} else if scrap_info[pp+1].cat==colon {
	make_underlined (pp)
	squash(pp,2,tag,-1,7)
} else if scrap_info[pp+1].cat==rbrace {
		squash(pp,1,stmt,-1,8)
} else if scrap_info[pp+1].cat==lpar && 
		scrap_info[pp+2].cat==rpar && 
		(scrap_info[pp+3].cat==const_like || 
		scrap_info[pp+3].cat==case_like) {
	big_app1(pp+2)
	big_app(' ')
	big_app1(pp+3)
	reduce(pp+2,2,rpar,0,9)
} else if scrap_info[pp+1].cat==cast && 
		(scrap_info[pp+2].cat==const_like || 
		scrap_info[pp+2].cat==case_like) {
	big_app1(pp+1)
	big_app(' ')
	big_app1(pp+2)
	reduce(pp+1,2,cast,0,9)
} else if scrap_info[pp+1].cat==exp || 
		scrap_info[pp+1].cat==cast {
	squash(pp,2,exp,-2,10)
}

@ @<Cases for |lpar|@>=
if (scrap_info[pp+1].cat==exp ||
	scrap_info[pp+1].cat==ubinop) && 
	scrap_info[pp+2].cat==rpar {
	squash(pp,3,exp,-2,11)
} else if scrap_info[pp+1].cat==rpar {
	big_app1(pp)
	app('\\')
	app(',')
	big_app1(pp+1)
@.\\,@>
	reduce(pp,2,exp,-2,12)
} else if (scrap_info[pp+1].cat==decl_head || 
		scrap_info[pp+1].cat==int_like || 
		scrap_info[pp+1].cat==cast) && 
		scrap_info[pp+2].cat==rpar {
	squash(pp,3,cast,-2,13)
} else if (scrap_info[pp+1].cat==decl_head || 
		scrap_info[pp+1].cat==int_like || 
		scrap_info[pp+1].cat==exp) && 
		scrap_info[pp+2].cat==comma {
	big_app3(pp)
	app(opt)
	app('9')
	reduce(pp,3,lpar,-1,14)
} else if scrap_info[pp+1].cat==stmt || 
		scrap_info[pp+1].cat==decl {
	big_app2(pp)
	big_app(' ')
	reduce(pp,2,lpar,-1,15)
}

@ @<Cases for |unop|@>=
if scrap_info[pp+1].cat==exp || 
	scrap_info[pp+1].cat==int_like {
	squash(pp,2,exp,-2,16)
}

@ @<Cases for |ubinop|@>=
if scrap_info[pp+1].cat==cast && 
	scrap_info[pp+2].cat==rpar {
	big_app('{')
	big_app1(pp)
	big_app('}')
	big_app1(pp+1)
	reduce(pp,2,cast,-2,17)
} else if scrap_info[pp+1].cat==exp || 
		scrap_info[pp+1].cat==int_like {
	big_app('{')
	big_app1(pp)
	big_app('}')
	big_app1(pp+1)
	reduce(pp,2,scrap_info[pp+1].cat,-2,18)
} else if scrap_info[pp+1].cat==binop {
	big_app(math_rel)
	big_app1(pp)
	big_app('{')
	big_app1(pp+1)
	big_app('}')
	big_app('}')
	reduce(pp,2,binop,-1,19)
}

@ @<Cases for |binop|@>=
if scrap_info[pp+1].cat==binop {
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

@ @<Cases for |cast|@>=
if scrap_info[pp+1].cat==lpar {
	squash(pp,2,lpar,-1,21)
} else if scrap_info[pp+1].cat==exp {
	big_app1(pp)
	big_app(' ')
	big_app1(pp+1)
	reduce(pp,2,exp,-2,21)
} else if scrap_info[pp+1].cat==semi {
	squash(pp,1,exp,-2,22)
}

@ @<Cases for |sizeof_like|@>=
if scrap_info[pp+1].cat==cast {
	squash(pp,2,exp,-2,23)
} else if scrap_info[pp+1].cat==exp {
	big_app1(pp)
	big_app(' ')
	big_app1(pp+1)
	reduce(pp,2,exp,-2,24)
}

@ @<Cases for |int_like|@>=
if scrap_info[pp+1].cat==int_like || 
	scrap_info[pp+1].cat==struct_like {
	big_app1(pp)
	big_app(' ')
	big_app1(pp+1)
	reduce(pp,2,scrap_info[pp+1].cat,-2,25)
} else if scrap_info[pp+1].cat==exp && 
		(scrap_info[pp+2].cat==raw_int ||
		scrap_info[pp+2].cat==struct_like) {
	squash(pp,2,int_like,-2,26)
} else if scrap_info[pp+1].cat==exp || 
		scrap_info[pp+1].cat==ubinop || 
		scrap_info[pp+1].cat==colon {
	big_app1(pp)
	big_app(' ')
	reduce(pp,1,decl_head,-1,27)
} else if scrap_info[pp+1].cat==semi || 
		scrap_info[pp+1].cat==binop {
	squash(pp,1,decl_head,0,28)
}

@ @<Cases for |public_like|@>=
if scrap_info[pp+1].cat==colon {
	squash(pp,2,tag,-1,29)
} else {
	squash(pp,1,int_like,-2,30)
}

@ @<Cases for |colcol|@>=
if scrap_info[pp+1].cat==exp || 
	scrap_info[pp+1].cat==int_like {
	app(qualifier)
	squash(pp,2,scrap_info[pp+1].cat,-2,31)
}@+else if scrap_info[pp+1].cat==colcol {
		squash(pp,2,colcol,-1,32)
}

@ @<Cases for |decl_head|@>=
if scrap_info[pp+1].cat==comma {
	big_app2(pp)
	big_app(' ')
	reduce(pp,2,decl_head,-1,33)
} else if scrap_info[pp+1].cat==ubinop {
	big_app1(pp)
	big_app('{')
	big_app1(pp+1)
	big_app('}')
	reduce(pp,2,decl_head,-1,34)
} else if scrap_info[pp+1].cat==exp && 
		scrap_info[pp+2].cat!=lpar && 
		scrap_info[pp+2].cat!=exp && 
		scrap_info[pp+2].cat!=cast {
	make_underlined(pp+1)
	squash(pp,2,decl_head,-1,35)
} else if (scrap_info[pp+1].cat==binop ||
		scrap_info[pp+1].cat==colon) && 
		scrap_info[pp+2].cat==exp && 
		(scrap_info[pp+3].cat==comma ||
		scrap_info[pp+3].cat==semi || 
		scrap_info[pp+3].cat==rpar) {
	squash(pp,3,decl_head,-1,36)
} else if scrap_info[pp+1].cat==cast {
		squash(pp,2,decl_head,-1,37)
} else if scrap_info[pp+1].cat==lbrace || 
		scrap_info[pp+1].cat==int_like || 
		scrap_info[pp+1].cat==decl {
	big_app1(pp)
	big_app(indent)
	app(indent)
	reduce(pp,1,fn_decl,0,38)
} else if scrap_info[pp+1].cat==semi {
	squash(pp,2,decl,-1,39)
}

@ @<Cases for |decl|@>=
if scrap_info[pp+1].cat==decl {
	big_app1(pp)
	big_app(force)
	big_app1(pp+1)
	reduce(pp,2,decl,-1,40)
} else if scrap_info[pp+1].cat==stmt || 
		scrap_info[pp+1].cat==function {
	big_app1(pp)
	big_app(big_force)
	big_app1(pp+1)
	reduce(pp,2,scrap_info[pp+1].cat,-1,41)
}

@ @<Cases for |base|@>=
if scrap_info[pp+1].cat==int_like || 
	scrap_info[pp+1].cat==exp {
	if scrap_info[pp+2].cat==comma {
		big_app1(pp)
		big_app(' ')
		big_app2(pp+1)
		app(opt)
		app('9')
		reduce(pp,3,base,0,42)
	} else if scrap_info[pp+2].cat==lbrace {
		big_app1(pp)
		big_app(' ')
		big_app1(pp+1)
		big_app(' ')
		big_app1(pp+2);
		reduce(pp,3,lbrace,-2,43)
	}
}

@ @<Cases for |struct_like|@>=
if scrap_info[pp+1].cat==lbrace {
	big_app1(pp)
	big_app(' ')
	big_app1(pp+1)
	reduce(pp,2,struct_head,0,44)
} else if scrap_info[pp+1].cat==exp ||
		scrap_info[pp+1].cat==int_like {
	if scrap_info[pp+2].cat==lbrace || 
		scrap_info[pp+2].cat==semi {
		make_underlined(pp+1)
		make_reserved(pp+1)
		big_app1(pp)
		big_app(' ')
		big_app1(pp+1)
		if scrap_info[pp+2].cat==semi {
			reduce(pp,2,decl_head,0,45)
		} else {
			big_app(' ')
			big_app1(pp+2)
			reduce(pp,3,struct_head,0,46)
		}
	} else if scrap_info[pp+2].cat==colon {
		squash(pp+2,1,base,2,47)
	} else if scrap_info[pp+2].cat!=base {
		big_app1(pp)
		big_app(' ')
		big_app1(pp+1)
		reduce(pp,2,int_like,-2,48)
	}
}

@ @<Cases for |struct_head|@>=
if (scrap_info[pp+1].cat==decl || 
	scrap_info[pp+1].cat==stmt || 
	scrap_info[pp+1].cat==function) && 
	scrap_info[pp+2].cat==rbrace {
	big_app1(pp)
	big_app(indent)
	big_app(force)
	big_app1(pp+1)
	big_app(outdent); big_app(force)
	big_app1(pp+2)
	reduce(pp,3,int_like,-2,49)
} else if scrap_info[pp+1].cat==rbrace {
	big_app1(pp)
	app_str("\\,")
	big_app1(pp+1)
@.\\,@>
	reduce(pp,2,int_like,-2,50)
}

@ @<Cases for |fn_decl|@>=
if scrap_info[pp+1].cat==decl {
	big_app1(pp)
	big_app(force)
	big_app1(pp+1)
	reduce(pp,2,fn_decl,0,51)
} else if scrap_info[pp+1].cat==stmt {
	big_app1(pp)
	app(outdent)
	app(outdent)
	big_app(force)
	big_app1(pp+1)
	reduce(pp,2,function,-1,52)
}

@ @<Cases for |function|@>=
if scrap_info[pp+1].cat==function || 
	scrap_info[pp+1].cat==decl || 
	scrap_info[pp+1].cat==stmt {
	big_app1(pp)
	big_app(big_force)
	big_app1(pp+1)
	reduce(pp,2,scrap_info[pp+1].cat,-1,53)
}

@ @<Cases for |lbrace|@>=
if scrap_info[pp+1].cat==rbrace {
	big_app1(pp)
	app('\\')
	app(',')
	big_app1(pp+1)
@.\\,@>
	reduce(pp,2,stmt,-1,54)
} else if (scrap_info[pp+1].cat==stmt ||
		scrap_info[pp+1].cat==decl ||
		scrap_info[pp+1].cat==function) && 
		scrap_info[pp+2].cat==rbrace {
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
} else if scrap_info[pp+1].cat==exp {
	if scrap_info[pp+2].cat==rbrace {
		squash(pp,3,exp,-2,56)
	} else if scrap_info[pp+2].cat==comma && 
				scrap_info[pp+3].cat==rbrace {
		squash(pp,4,exp,-2,56)
	}
}

@ @<Cases for |if_like|@>=
if scrap_info[pp+1].cat==exp {
	big_app1(pp)
	big_app(' ')
	big_app1(pp+1)
	reduce(pp,2,if_clause,0,57);
}

@ @<Cases for |else_like|@>=
if scrap_info[pp+1].cat==colon {
	squash(pp+1,1,base,1,58)
} else if scrap_info[pp+1].cat==lbrace {
		squash(pp,1,else_head,0,59)
} else if scrap_info[pp+1].cat==stmt {
	big_app(force)
	big_app1(pp)
	big_app(indent)
	big_app(break_space)
	big_app1(pp+1)
	big_app(outdent)
	big_app(force)
	reduce(pp,2,stmt,-1,60)
}

@ @<Cases for |else_head|@>=
if scrap_info[pp+1].cat==stmt || 
	scrap_info[pp+1].cat==exp {
	big_app(force)
	big_app1(pp)
	big_app(break_space)
	app(noop);
	big_app(cancel)
	big_app1(pp+1)
	big_app(force)
	reduce(pp,2,stmt,-1,61)
}

@ @<Cases for |if_clause|@>=
if scrap_info[pp+1].cat==lbrace {
	squash(pp,1,if_head,0,62)
} else if scrap_info[pp+1].cat==stmt {
	if scrap_info[pp+2].cat==else_like {
		big_app(force)
		big_app1(pp)
		big_app(indent)
		big_app(break_space)
		big_app1(pp+1)
		big_app(outdent)
		big_app(force)
		big_app1(pp+2)
		if scrap_info[pp+3].cat==if_like {
			big_app(' ')
			big_app1(pp+3)
			reduce(pp,4,if_like,0,63)
		}@+else {
			reduce(pp,3,else_like,0,64)
		}
	} else {
		squash(pp,1,else_like,0,65)
	}
}

@ @<Cases for |if_head|@>=
if scrap_info[pp+1].cat==stmt || 
	scrap_info[pp+1].cat==exp {
	if scrap_info[pp+2].cat==else_like {
		big_app(force)
		big_app1(pp)
		big_app(break_space)
		app(noop)
		big_app(cancel)
		big_app1(pp+1)
		big_app(force)
		big_app1(pp+2)
		if scrap_info[pp+3].cat==if_like {
			big_app(' ')
			big_app1(pp+3)
			reduce(pp,4,if_like,0,66)
		}@+else {
			reduce(pp,3,else_like,0,67)
		}
	} else {
		squash(pp,1,else_head,0,68)
	}
}

@ @<Cases for |do_like|@>=
if scrap_info[pp+1].cat==stmt && 
	scrap_info[pp+2].cat==else_like && 
	scrap_info[pp+3].cat==semi {
	big_app1(pp)
	big_app(break_space)
	app(noop)
	big_app(cancel)
	big_app1(pp+1)
	big_app(cancel)
	app(noop)
	big_app(break_space)
	big_app2(pp+2)
	reduce(pp,4,stmt,-1,69)
}

@ @<Cases for |case_like|@>=
if scrap_info[pp+1].cat==semi {
	squash(pp,2,stmt,-1,70)
} else if scrap_info[pp+1].cat==colon {
		squash(pp,2,tag,-1,71)
} else if scrap_info[pp+1].cat==exp {
	big_app1(pp)
	big_app(' ')
	big_app1(pp+1)
	reduce(pp,2,exp,-2,72)
}

@ @<Cases for |catch_like|@>=
if scrap_info[pp+1].cat==cast || 
	scrap_info[pp+1].cat==exp {
	big_app2(pp)
	big_app(indent)
	big_app(indent)
	reduce(pp,2,fn_decl,0,73)
}

@ @<Cases for |tag|@>=
if scrap_info[pp+1].cat==tag {
	big_app1(pp)
	big_app(break_space)
	big_app1(pp+1)
	reduce(pp,2,tag,-1,74)
} else if scrap_info[pp+1].cat==stmt ||
		scrap_info[pp+1].cat==decl ||
		scrap_info[pp+1].cat==function {
	big_app(force)
	big_app(backup)
	big_app1(pp)
	big_app(break_space)
	big_app1(pp+1)
	reduce(pp,2,scrap_info[pp+1].cat,-1,75)
}

@ The user can decide at run-time whether short statements should be
grouped together on the same line.

@<Cases for |stmt|@>=
if scrap_info[pp+1].cat==stmt ||
	scrap_info[pp+1].cat==decl ||
	scrap_info[pp+1].cat==function {
	big_app1(pp)
	if scrap_info[pp+1].cat==function {
		big_app(big_force)
	} else if scrap_info[pp+1].cat==decl {
		big_app(big_force)
	} else if flags['f'] {
		big_app(force)
	} else {
		big_app(break_space)
	}
	big_app1(pp+1)
	reduce(pp,2,scrap_info[pp+1].cat,-1,76)
}

@ @<Cases for |semi|@>=
big_app(' ')
big_app1(pp)
reduce(pp,1,stmt,-1,77)

@ @<Cases for |lproc|@>=
if scrap_info[pp+1].cat==define_like {
	make_underlined(pp+2)
}
if scrap_info[pp+1].cat==else_like || 
	scrap_info[pp+1].cat==if_like ||
	scrap_info[pp+1].cat==define_like {
	squash(pp,2,lproc,0,78)
} else if scrap_info[pp+1].cat==rproc {
	app(inserted)
	big_app2(pp)
	reduce(pp,2,insert,-1,79)
} else if scrap_info[pp+1].cat==exp || 
		scrap_info[pp+1].cat==function {
	if scrap_info[pp+2].cat==rproc {
		app(inserted)
		big_app1(pp)
		big_app(' ')
		big_app2(pp+1)
		reduce(pp,3,insert,-1,80)
	} else if scrap_info[pp+2].cat==exp && 
				scrap_info[pp+3].cat==rproc && 
				scrap_info[pp+1].cat==exp {
		app(inserted)
		big_app1(pp)
		big_app(' ')
		big_app1(pp+1)
		app_str(" \\5")
@.\\5@>
		big_app2(pp+2)
		reduce(pp,4,insert,-1,80)
	}
}

@ @<Cases for |section_scrap|@>=
if scrap_info[pp+1].cat==semi {
	big_app2(pp)
	big_app(force)
	reduce(pp,2,stmt,-2,81)
} else {
	squash(pp,1,exp,-2,82)
}

@ @<Cases for |insert|@>=
if scrap_info[pp+1].cat != 0 {
	squash(pp,2,scrap_info[pp+1].cat,0,83)
}

@ @<Cases for |prelangle|@>=
init_mathness=yes_math
cur_mathness=yes_math
app('<')
reduce(pp,1,binop,-2,84)

@ @<Cases for |prerangle|@>=
init_mathness=yes_math
cur_mathness=yes_math
app('>')
reduce(pp,1,binop,-2,85)

@ @<Cases for |langle|@>=
if scrap_info[pp+1].cat==prerangle {
	big_app1(pp)
	app('\\')
	app(',')
	big_app1(pp+1)
@.\\,@>
	reduce(pp,2,cast,-1,86)
} else if scrap_info[pp+1].cat==decl_head || 
		scrap_info[pp+1].cat==int_like || 
		scrap_info[pp+1].cat==exp {
	if scrap_info[pp+2].cat==prerangle {
		squash(pp,3,cast,-1,87)
	} else if scrap_info[pp+2].cat==comma {
		big_app3(pp)
		app(opt)
		app('9')
		reduce(pp,3,langle,0,88)
	}
}

@ @<Cases for |template_like|@>=
if scrap_info[pp+1].cat==exp && 
	scrap_info[pp+2].cat==prelangle {
	squash(pp+2,1,langle,2,89)
} else if scrap_info[pp+1].cat==exp || 
		scrap_info[pp+1].cat==raw_int {
	big_app1(pp)
	big_app(' ')
	big_app1(pp+1)
	reduce(pp,2,scrap_info[pp+1].cat,-2,90)
}@+ else {
	squash(pp,1,raw_int,0,91)
}

@ @<Cases for |new_like|@>=
if scrap_info[pp+1].cat==lpar && 
	scrap_info[pp+2].cat==exp && 
	scrap_info[pp+3].cat==rpar {
	squash(pp,4,new_like,0,92)
} else if scrap_info[pp+1].cat==cast {
	big_app1(pp)
	big_app(' ')
	big_app1(pp+1)
	reduce(pp,2,exp,-2,93)
} else if scrap_info[pp+1].cat!=lpar {
	squash(pp,1,new_exp,0,94)
}

@ @<Cases for |new_exp|@>=
if scrap_info[pp+1].cat==int_like || 
	scrap_info[pp+1].cat==const_like {
	big_app1(pp)
	big_app(' ')
	big_app1(pp+1)
	reduce(pp,2,new_exp,0,95)
} else if scrap_info[pp+1].cat==struct_like && 
		(scrap_info[pp+2].cat==exp || 
		scrap_info[pp+2].cat==int_like) {
	big_app1(pp)
	big_app(' ')
	big_app1(pp+1)
	big_app(' ')
	big_app1(pp+2)
	reduce(pp,3,new_exp,0,96)
} else if scrap_info[pp+1].cat==raw_ubin {
	big_app1(pp)
	big_app('{')
	big_app1(pp+1)
	big_app('}')
	reduce(pp,2,new_exp,0,97)
} else if scrap_info[pp+1].cat==lpar {
	squash(pp,1,exp,-2,98)
} else if (scrap_info[pp+1].cat==exp) {
	big_app1(pp)
	big_app(' ')
	reduce(pp,1,exp,-2,98)
} else if scrap_info[pp+1].cat!=raw_int && 
		scrap_info[pp+1].cat!=struct_like && 
		scrap_info[pp+1].cat!=colcol {
	squash(pp,1,exp,-2,99)
}

@ @<Cases for |ftemplate|@>=
if scrap_info[pp+1].cat==prelangle {
	squash(pp+1,1,langle,1,100)
} else {
	squash(pp,1,exp,-2,101) 
}

@ @<Cases for |for_like|@>=
if scrap_info[pp+1].cat==exp {
	big_app1(pp)
	big_app(' ')
	big_app1(pp+1)
	reduce(pp,2,else_like,-2,102)
}

@ @<Cases for |raw_ubin|@>=
if scrap_info[pp+1].cat==const_like {
	big_app2(pp)
	app_str("\\ ")
	reduce(pp,2,raw_ubin,0,103)
@.\\\ @>
} else {
	squash(pp,1,ubinop,-2,104)
}

@ @<Cases for |const_like|@>=
squash(pp,1,int_like,-2,105)

@ @<Cases for |raw_int|@>=
if scrap_info[pp+1].cat==prelangle { 
	squash(pp+1,1,langle,1,106)
} else if scrap_info[pp+1].cat==colcol {
	squash(pp,2,colcol,-1,107)
} else if scrap_info[pp+1].cat==cast {
	squash(pp,2,raw_int,0,108)
} else if scrap_info[pp+1].cat==lpar {
	squash(pp,1,exp,-2,109)
} else if scrap_info[pp+1].cat!=langle {
	squash(pp,1,int_like,-3,110)
}

@ @<Cases for |operator_like|@>=
if scrap_info[pp+1].cat==binop || 
		scrap_info[pp+1].cat==unop || 
		scrap_info[pp+1].cat==ubinop {
	if scrap_info[pp+2].cat==binop {
		break
	}
	big_app1(pp)
	big_app('{')
	big_app1(pp+1)
	big_app('}')
	reduce(pp,2,exp,-2,111)
} else if scrap_info[pp+1].cat==new_like || 
		scrap_info[pp+1].cat==delete_like {
	big_app1(pp)
	big_app(' ')
	big_app1(pp+1)
	reduce(pp,2,exp,-2,112);
} else if scrap_info[pp+1].cat==comma {
	squash(pp,2,exp,-2,113)
} else if scrap_info[pp+1].cat!=raw_ubin {
	squash(pp,1,new_exp,0,114)
}

@ @<Cases for |typedef_like|@>=
if (scrap_info[pp+1].cat==int_like || 
		scrap_info[pp+1].cat==cast) && 
		(scrap_info[pp+2].cat==comma || 
		scrap_info[pp+2].cat==semi) {
	squash(pp+1,1,exp,-1,115)
} else if scrap_info[pp+1].cat==int_like {
	big_app1(pp)
	big_app(' ')
	big_app1(pp+1)
	reduce(pp,2,typedef_like,0,116)
} else if scrap_info[pp+1].cat==exp && 
		scrap_info[pp+2].cat!=lpar && 
		scrap_info[pp+2].cat!=exp && 
		scrap_info[pp+2].cat!=cast {
	make_underlined(pp+1)
	make_reserved(pp+1)
	big_app1(pp)
	big_app(' ')
	big_app1(pp+1)
	reduce(pp,2,typedef_like,0,117)
} else if scrap_info[pp+1].cat==comma {
	big_app2(pp)
	big_app(' ')
	reduce(pp,2,typedef_like,0,118)
} else if scrap_info[pp+1].cat==semi {
	squash(pp,2,decl,-1,119)
} else if scrap_info[pp+1].cat==ubinop && 
		(scrap_info[pp+2].cat==ubinop || 
		scrap_info[pp+2].cat==cast) {
	big_app('{')
	big_app1(pp+1)
	big_app('}')
	big_app1(pp+2)
	reduce(pp+1,2,scrap_info[pp+2].cat,0,120)
}

@ @<Cases for |delete_like|@>=
if scrap_info[pp+1].cat==lpar && 
	scrap_info[pp+2].cat==rpar {
	big_app2(pp)
	app('\\')
	app(',')
	big_app1(pp+2)
@.\\,@>
	reduce(pp,3,delete_like,0,121)
} else if scrap_info[pp+1].cat==exp {
	big_app1(pp)
	big_app(' ')
	big_app1(pp+1)
	reduce(pp,2,exp,-2,122)
}

@ @<Cases for |question|@>=
if scrap_info[pp+1].cat==exp && 
	(scrap_info[pp+2].cat==colon || 
	scrap_info[pp+2].cat==base) {
	scrap_info[pp+2].mathness=5*yes_math /* this colon should be in math mode */
	squash(pp,3,binop,-2,123)
}

@ Now here's the |reduce| procedure used in our code for productions.

The `|freeze_text|' function is used to give official status to a token list.
Before saying |freeze_text|, items are appended to the current token list,
and we know that the eventual number of this token list will be the current
value of |len(tok_mem)|. But no list of that number really exists as yet,
because no ending point for the current list has been
stored in the |tok_start| array. After saying |freeze_text|, the
old current token list becomes legitimate, and the new
current token list is empty and ready to be appended to.

@c
func freeze_text() {
	tok_start = append(tok_start, int32(len(tok_mem)))
}

@ @c
func reduce(j int32, k int32, c rune, d int32, n int32) {
	scrap_info[j].cat=c
	scrap_info[j].trans_plus.Trans=int32(len(tok_start)-1)
	scrap_info[j].mathness=4*cur_mathness+init_mathness
	freeze_text()
	if k>1 {
		i:=j+k
		i1:=j+1
		for i<=lo_ptr {
			scrap_info[i1].cat=scrap_info[i].cat
			scrap_info[i1].trans_plus.Trans=scrap_info[i].trans_plus.Trans
			scrap_info[i1].mathness=scrap_info[i].mathness
			i++
			i1++
		}
		lo_ptr=lo_ptr-k+1
	}
	if pp+d<scrap_base {
		pp = scrap_base
	} else { 
		pp=pp+d
	}
	f := "reduce"
	@<Print a snapshot of the scrap list if debugging @>
	pp-- /* we next say |pp++| */
}

@ Here's the |squash| procedure, which
takes advantage of the simplification that occurs when |k==1|.

@c
func squash(j int32, k int32, c rune,d int32, n int32) {
	if k==1 {
		scrap_info[j].cat=c
		if pp+d<scrap_base {
			pp = scrap_base
		} else { 
			pp=pp+d
		}
		f := "squash"
		@<Print a snapshot...@>
		pp-- /* we next say |pp++| */
		return
	}
	for i:=j; i<j+k; i++ {
		big_app1(i)
	}
	reduce(j,k,c,d,n)
}

@ And here now is the code that applies productions as long as possible.
Before applying the production mechanism, we must make sure
it has good input (at least four scraps, the length of the lhs of the
longest rules), and that there is enough room in the memory arrays
to hold the appended tokens and texts.  Here we use a very
conservative test; it's more important to make sure the program
will still work if we change the production rules (within reason)
than to squeeze the last bit of space from the memory arrays.

@<Constants@>=
safe_tok_incr = 20
safe_text_incr = 10
safe_scrap_incr = 10

@ @<Reduce the scraps using the productions until no more rules apply@>=
for true {
	@<Make sure the entries |pp| through |pp+3| of |cat| are defined@>
	if pp>lo_ptr {
		break
	}
	init_mathness=maybe_math
	cur_mathness=maybe_math
	@<Match a production...@>
}

@ If we get to the end of the scrap list, category codes equal to zero are
stored, since zero does not match anything in a production.

@<Make sure the entries...@>=
if lo_ptr<pp+3 {
	for hi_ptr<=scrap_ptr && lo_ptr!=pp+3 {
		lo_ptr++
		scrap_info[lo_ptr].cat=scrap_info[hi_ptr].cat
		scrap_info[lo_ptr].mathness=scrap_info[hi_ptr].mathness
		scrap_info[lo_ptr].trans_plus.Trans=scrap_info[hi_ptr].trans_plus.Trans
		hi_ptr++
	}
	for i:=lo_ptr+1;i<=pp+3;i++ {
		scrap_info[i].cat=0
	}
}

@ If \.{GOWEAVE} is being run in debugging mode, the production numbers and
current stack categories will be printed out when |tracing| is set to 2;
a sequence of two or more irreducible scraps will be printed out when
|tracing| is set to 1.

@<Global...@>=
var tracing int32  /* can be used to show parsing details */

@ @<Print a snapsh...@>=
{ 
	if tracing==2 {
		fmt.Printf("\n%s %d:", f, n)
		for k:=scrap_base; k<=lo_ptr; k++ {
			if k==pp {
				fmt.Print("*") 
			} else {
				fmt.Print(" ")
			}
			if scrap_info[k].mathness %4 == yes_math {
				fmt.Print("+")
			} else if scrap_info[k].mathness %4 == no_math {
				fmt.Print("-")
			}
			print_cat(scrap_info[k].cat)
			if scrap_info[k].mathness /4 == yes_math {
				fmt.Print("+")
			} else if scrap_info[k].mathness /4 == no_math {
				fmt.Print("-")
			}
		}
		if hi_ptr<=scrap_ptr {
			 fmt.Print("...") /* indicate that more is coming */
		}
	}
}

@ The |translate| function assumes that scraps have been stored in
positions |scrap_base| through |scrap_ptr| of |cat| and |trans_plus.Trans|. It
applies productions as much as
possible. The result is a token list containing the translation of
the given sequence of scraps.

@c 
/* converts a sequence of scraps */
func translate() int32 {
	pp=scrap_base
	lo_ptr=pp-1
	hi_ptr=pp
	@<If tracing, print an indication of where we are@>
	@<Reduce the scraps...@>
	@<Combine the irreducible scraps that remain@>
}

@ If the initial sequence of scraps does not reduce to a single scrap,
we concatenate the translations of all remaining scraps, separated by
blank spaces, with dollar signs surrounding the translations of scraps
where appropriate.

@<Combine the irreducible...@>= {
	@<If semi-tracing, show the irreducible scraps@>
	for j:=scrap_base; j<=lo_ptr; j++ {
		if j!=scrap_base {
			app(' ')
		}
		if scrap_info[j].mathness % 4 == yes_math {
			app('$')
		}
		app1(j)
		if scrap_info[j].mathness / 4 == yes_math {
			app('$')
		}
	}
	freeze_text()
	return int32(len(tok_start)-2)
}

@ @<If semi-tracing, show the irreducible scraps@>=
if lo_ptr>scrap_base && tracing==1 {
	fmt.Printf("\nIrreducible scrap sequence in section %d:",section_count);
@.Irreducible scrap sequence...@>
	mark_harmless()
	for j:=scrap_base; j<=lo_ptr; j++ {
		fmt.Printf(" ")
		print_cat(scrap_info[j].cat)
	}
}

@ @<If tracing,...@>=
if tracing==2 {
	fmt.Printf("\nTracing after l. %d:\n",line[include_depth])
	mark_harmless()
@.Tracing after...@>
}

@* Initializing the scraps.
If we are going to use the powerful production mechanism just developed, we
must get the scraps set up in the first place, given a \GO/ text. A table
of the initial scraps corresponding to \GO/ tokens appeared above in the
section on parsing; our goal now is to implement that table. We shall do this
by implementing a subroutine called |C_parse| that is analogous to the
|C_xref| routine used during phase one.

Like |C_xref|, the |C_parse| procedure starts with the current
value of |next_control| and it uses the operation |next_control=get_next()|
repeatedly to read \GO/ text until encountering the next `\.{\v}' or
`\.{/*}', or until |next_control>=format_code|. The scraps corresponding to
what it reads are appended into the |cat| and |trans_plus.Trans| arrays, and |scrap_ptr|
is advanced.

@c
/* creates scraps from \GO/ tokens */
func C_parse(spec_ctrl rune) {
	for next_control<format_code || next_control==spec_ctrl {
		@<Append the scrap appropriate to |next_control|@>
		next_control=get_next()
		if next_control=='|' || next_control==begin_comment ||
				next_control==begin_short_comment {
			return
		}
	}
}

@ The following function is used to append a scrap whose tokens have just
been appended:

@c
func app_scrap(c int32, b int32) {
	scrap_ptr++
	scrap_info[scrap_ptr].cat=c
	scrap_info[scrap_ptr].trans_plus.Trans=int32(len(tok_start)-1)
	scrap_info[scrap_ptr].mathness=5*(b) /* no no, yes yes, or maybe maybe */
	freeze_text()
}

@ @<Append the scr...@>=
switch (next_control) {
	case section_name:
		app(section_flag+cur_section)
		app_scrap(section_scrap,maybe_math)
		app_scrap(exp,yes_math)
	case str, constant,verbatim:
		@<Append a string or constant@>
	case identifier: 
		app_cur_id(true)
	case TeX_string:
		@<Append a \TEX/ string, without forming a scrap@>
	case '/', '.':
		app(next_control)
		app_scrap(binop,yes_math)
	case '<': 
		app_str("\\langle")
		@+app_scrap(prelangle,yes_math)
@.\\langle@>
	case '>': 
		app_str("\\rangle")
		@+app_scrap(prerangle,yes_math)
@.\\rangle@>
	case '=': 
		app_str("\\K")
		app_scrap(binop,yes_math)
@.\\K@>
	case '|': 
		app_str("\\OR")
		app_scrap(binop,yes_math)
@.\\OR@>
	case '^': 
		app_str("\\XOR")
		app_scrap(binop,yes_math)
@.\\XOR@>
	case '%': 
		app_str("\\MOD")
		app_scrap(binop,yes_math)
@.\\MOD@>
	case '!': 
		app_str("\\R")
		app_scrap(unop,yes_math)
@.\\R@>
	case '~': 
		app_str("\\CM")
		app_scrap(unop,yes_math)
@.\\CM@>
	case '+', '-':
		app(next_control)
		app_scrap(ubinop,yes_math)
	case '*': 
		app(next_control)
		app_scrap(raw_ubin,yes_math)
	case '&': 
		app_str("\\AND")
		app_scrap(raw_ubin,yes_math)
@.\\AND@>
	case '?': 
		app_str("\\?")
		app_scrap(question,yes_math)
@.\\?@>
	case '#': 
		app_str("\\#")
		app_scrap(ubinop,yes_math)
@.\\\#@>
	case ignore, xref_roman, xref_wildcard, xref_typewriter, noop:
		@+break;
	case '(', '[': 
		app(next_control)
		app_scrap(lpar,maybe_math)
	case ')', ']': 
		app(next_control)
		app_scrap(rpar,maybe_math)
	case '{': 
		app_str("\\{"@q}@>)
		app_scrap(lbrace,yes_math)
@.\\\{@>@q}@>
	case '}': 
		app_str(@q{@>"\\}")
		app_scrap(rbrace,yes_math)
@q{@>@.\\\}@>
	case ',': 
		app(',')
		app_scrap(comma,yes_math)
	case ';': 
		app(';')
		app_scrap(semi,maybe_math)
	case ':': 
		app(':')
		app_scrap(colon,no_math)@/
	@t\4@>  @<Cases involving nonstandard characters@>
	case thin_space: 
		app_str("\\,")
		app_scrap(insert,maybe_math)
@.\\,@>
	case math_break: 
		app(opt)
		app_str("0")
		app_scrap(insert,maybe_math)
	case line_break: 
		app(force)
		app_scrap(insert,no_math)
	case left_preproc: 
		app(force)
		app(preproc_line)
		app_str("\\#")
		app_scrap(lproc,no_math)
@.\\\#@>
	case right_preproc: 
		app(force)
		app_scrap(rproc,no_math)
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
	case macro_arg_open: 
		app_scrap(begin_arg,maybe_math)
	case macro_arg_close: 
		app_scrap(end_arg,maybe_math)
	case join: 
		app_str("\\J")
		app_scrap(insert,no_math)
@.\\J@>
	case output_defs_code: 
		app(force)
		app_str("\\ATH")
		app(force)
		app_scrap(insert,no_math)
@.\\ATH@>
	default: 
		app(inserted)
		app(next_control)
		app_scrap(insert,maybe_math)
}

@ Some nonstandard characters may have entered \.{GOWEAVE} by means of
standard ones. They are converted to \TEX/ control sequences so that it is
possible to keep \.{GOWEAVE} from outputting unusual |rune| codes.

@<Cases involving nonstandard...@>=
case not_eq: 
	app_str("\\I")
	@+app_scrap(binop,yes_math)
@.\\I@>
case lt_eq: 
	app_str("\\Z")
	@+app_scrap(binop,yes_math)
@.\\Z@>
case gt_eq: 
	app_str("\\G")
	@+app_scrap(binop,yes_math)
@.\\G@>
case eq_eq: 
	app_str("\\E")
	@+app_scrap(binop,yes_math)
@.\\E@>
case and_and: 
	app_str("\\W")
	@+app_scrap(binop,yes_math)
@.\\W@>
case or_or: 
	app_str("\\V")
	@+app_scrap(binop,yes_math)
@.\\V@>
case plus_plus: 
	app_str("\\PP")
	@+app_scrap(unop,yes_math)
@.\\PP@>
case minus_minus: 
	app_str("\\MM")
	@+app_scrap(unop,yes_math)
@.\\MM@>
case gt_gt: 
	app_str("\\GG")
	@+app_scrap(binop,yes_math)
@.\\GG@>
case lt_lt: 
	app_str("\\LL")
	@+app_scrap(binop,yes_math)
@.\\LL@>
case dot_dot_dot: 
	app_str("\\,\\ldots\\,")
	@+app_scrap(raw_int,yes_math);
@.\\,@>
@.\\ldots@>

@ Many of the special characters in a string must be prefixed by `\.\\' so that
\TEX/ will print them properly.
@^special string characters@>

@<Append a string or...@>=
count:= -1 
if next_control==constant {
	app_str("\\T{"@q}@>)
@.\\T@>
} else if next_control==str {
	count=20
	app_str("\\.{"@q}@>)
@.\\.@>
} else {
	app_str("\\vb{"@q}@>)
}
@.\\vb@>
for i:=0; i < len(id); {
	if count==0 { /* insert a discretionary break in a long string */
		app_str(@q(@>@q{@>"}\\)\\.{"@q}@>)
		count=20
@q(@>@.\\)@>
	}
	switch (id[i]) {
		case ' ', '\\', '#', '%', '$','^', '{', '}', '~', '&', '_': 
			app('\\')
@.\\\ @>
@.\\\\@>
@.\\\#@>
@.\\\%@>
@.\\\$@>
@.\\\^@>
@.\\\{@>@q}@>
@q{@>@.\\\}@>
@.\\\~@>
@.\\\&@>
@.\\\_@>
		case '@@': 
			if i+1 < len(id) && id[i+1]=='@@' {
				i++
			} else { 
				err_print("! Double @@ should be used in strings")
			}
@.Double @@ should be used...@>
	}
	app_tok(id[i])
	i++
	count--
}
app(@q{@>'}')
app_scrap(exp,maybe_math)

@ We do not make the \TEX/ string into a scrap, because there is no
telling what the user will be putting into it; instead we leave it
open, to be picked up by the next scrap. If it comes at the end of a
section, it will be made into a scrap when |finish_C| is called.

There's a known bug here, in cases where an adjacent scrap is
|prelangle| or |prerangle|. Then the \TEX/ string can disappear
when the \.{\\langle} or \.{\\rangle} becomes \.{<} or \.{>}.
For example, if the user writes \.{\v x<@@ty@@>\v}, the \TEX/ string
\.{\\hbox\{y\}} eventually becomes part of an |insert| scrap, which is combined
with a |prelangle| scrap and eventually lost. The best way to work around
this bug is probably to enclose the \.{@@t...@@>} in \.{@@[...@@]} so that
the \TEX/ string is treated as an expression.
@^bug, known@>

@<Append a \TEX/ string, without forming a scrap@>=
app_str("\\hbox{"@q}@>)
for i:=0; i < len(id);{ 
	if id[i]=='@@' {
		i++
	}
	app_tok(id[i])
	i++
}
app(@q{@>'}')

@ The function |app_cur_id| appends the current identifier to the
token list; it also builds a new scrap if |scrapping==1|.

@ @c
func app_cur_id(scrapping bool) {
	p:=id_lookup(id,normal)
	if name_dir[p].ilk<=custom { /* not a reserved word */
		app(id_flag+p)
		if scrapping {
			a1 := exp
			if name_dir[p].ilk==func_template {
				a1 = ftemplate
			}
			a2 := maybe_math
			if name_dir[p].ilk==custom {
				a2 = yes_math
			} 
			app_scrap(a1, a2)
		}
@.\\NULL@>
	} else {
		app(res_flag+p)
		if scrapping {
			if name_dir[p].ilk==alfop {
				app_scrap(ubinop,yes_math)
			} else {
				app_scrap(name_dir[p].ilk,maybe_math)
			}
		}
	}
}

@ When the `\.{\v}' that introduces \GO/ text is sensed, a call on
|C_translate| will return a pointer to the \TEX/ translation of
that text. If scraps exist in |scrap_info|, they are
unaffected by this translation process.

@c
func C_translate() int32 {
	save_base:=scrap_base /* holds original value of |scrap_base| */
	scrap_base=scrap_ptr+1
	C_parse(section_name) /* get the scraps together */
	if next_control!='|' {
		err_print("! Missing '|' after C text")
@.Missing '|'...@>
	}
	app_tok(cancel)
	app_scrap(insert,maybe_math)
				/* place a |cancel| token as a final ``comment'' */
	p:=translate() /* make the translation */
	if scrap_ptr>max_scr_ptr {
		max_scr_ptr=scrap_ptr
	}
	scrap_ptr=scrap_base-1
	scrap_base=save_base /* scrap the scraps */
	return p
}

@ The |outer_parse| routine is to |C_parse| as |outer_xref|
is to |C_xref|: It constructs a sequence of scraps for \GO/ text
until |next_control>=format_code|. Thus, it takes care of embedded comments.

The token list created from within `\pb' brackets is output as an argument
to \.{\\PB}, if the user has invoked \.{GOWEAVE} with the \.{+e} flag.
Although \.{cwebmac} ignores \.{\\PB}, other macro packages
might use it to localize the special meaning of the macros that mark up
program text.

@c
/* makes scraps from \GO/ tokens and comments */
func outer_parse() {
	for next_control<format_code {
		if next_control!=begin_comment && next_control!=begin_short_comment {
			C_parse(ignore)
		} else {
			is_long_comment:=(next_control==begin_comment);
			app(cancel)
			app(inserted)
			if is_long_comment {
				app_str("\\C{"@q}@>)
@.\\C@>
			} else {
				app_str("\\SHC{"@q}@>)
			}
@.\\SHC@>
			bal:=copy_comment(is_long_comment,1)  /* brace level in comment */
			next_control=ignore
			for bal>0 {
				p:=int32(len(tok_start)-1)
				freeze_text()
				q:=C_translate()/* partial comments */
				app(tok_flag+p)
				if flags['e'] {
					app_str("\\PB{")
@.\\PB@>
				}
				app(inner_tok_flag+q)
				if flags['e'] {
					app_tok('}') 
				}
				if next_control=='|' {
					bal=copy_comment(is_long_comment,bal)
					next_control=ignore
				} else {
					bal=0 /* an error has been reported */
				}
			}
			app(force)
			app_scrap(insert,no_math)
				/* the full comment becomes a scrap */
		}
	}
}

@* Output of tokens.
So far our programs have only built up multi-layered token lists in
\.{GOWEAVE}'s internal memory; we have to figure out how to get them into
the desired final form. The job of converting token lists to characters in
the \TEX/ output file is not difficult, although it is an implicitly
recursive process. Four main considerations had to be kept in mind when
this part of \.{GOWEAVE} was designed.  (a) There are two modes of output:
|outer| mode, which translates tokens like |force| into line-breaking
control sequences, and |inner| mode, which ignores them except that blank
spaces take the place of line breaks. (b) The |cancel| instruction applies
to adjacent token or tokens that are output, and this cuts across levels
of recursion since `|cancel|' occurs at the beginning or end of a token
list on one level. (c) The \TEX/ output file will be semi-readable if line
breaks are inserted after the result of tokens like |break_space| and
|force|.  (d) The final line break should be suppressed, and there should
be no |force| token output immediately after `\.{\\Y\\B}'.

@ The output process uses a stack to keep track of what is going on at
different ``levels'' as the token lists are being written out. Entries on
this stack have three parts:

\yskip\hang |end_field| is the |tok_mem| location where the token list of a
particular level will end;

\yskip\hang |tok_field| is the |tok_mem| location from which the next token
on a particular level will be read;

\yskip\hang |mode_field| is the current mode, either |inner| or |outer|.

\yskip\noindent The current values of these quantities are referred to
quite frequently, so they are stored in a separate place instead of in the
|stack| array. We call the current values |cur_state.end_field|, |cur_state.tok_field|, and
|cur_state.mode_field|.

The global variable |stack_ptr| tells how many levels of output are
currently in progress. The end of output occurs when an |end_translation|
token is found, so the stack is never empty except when we first begin the
output process.

@c type mode int

@ @<Constants@>=
inner mode = 0 /* value of |mode| for \GO/ texts within \TEX/ texts */
outer mode = 1 /* value of |mode| for \GO/ texts in sections */

@ @<Typed...@>= 
type output_state struct {
	end_field int32/* ending location of token list */
	tok_field int32 /* present location within token list */
	mode_field mode /* interpretation of control tokens */
}
type stack_pointer int32

@ @c func init_stack() {
	stack_ptr=0
	cur_state.mode_field=outer
}

@ @<Global...@>=
var cur_state output_state /* |cur_state.end_field|, 
	|cur_state.tok_field|, |cur_state.mode_field| */
var stack[stack_size]output_state /* info for non-current levels */
var stack_ptr stack_pointer /* first unused location in the output state stack */
var stack_end stack_pointer=stack_size-1 /* end of |stack| */
var max_stack_ptr stack_pointer /* largest value assumed by |stack_ptr| */

@ @<Set init...@>=
max_stack_ptr=0

@ To insert token-list |p| into the output, the |push_level| subroutine
is called; it saves the old level of output and gets a new one going.
The value of |cur_state.mode_field| is not changed.

@c
 /* suspends the current level */
func push_level(p int32) {
	if stack_ptr==stack_end {
		overflow("stack")
	}
	if stack_ptr>0 { /* save current state */
		stack[stack_ptr].end_field=cur_state.end_field
		stack[stack_ptr].tok_field=cur_state.tok_field
		stack[stack_ptr].mode_field=cur_state.mode_field
	}
	stack_ptr++
	if stack_ptr>max_stack_ptr {
		max_stack_ptr=stack_ptr
	}
	cur_state.tok_field=tok_start[p]
	cur_state.end_field=tok_start[p+1]
}

@ Conversely, the |pop_level| routine restores the conditions that were in
force when the current level was begun. This subroutine will never be
called when |stack_ptr==1|.

@c
func pop_level() {
	stack_ptr--
	cur_state.end_field=stack[stack_ptr].end_field
	cur_state.tok_field=stack[stack_ptr].tok_field
	cur_state.mode_field=stack[stack_ptr].mode_field
}

@ The |get_output| function returns the next byte of output that is not a
reference to a token list. It returns the values |identifier| or |res_word|
or |section_code| if the next token is to be an identifier (typeset in
italics), a reserved word (typeset in boldface), or a section name (typeset
by a complex routine that might generate additional levels of output).
In these cases |cur_name| points to the identifier or section name in
question.

@<Global...@>=
var cur_name int32 = -1

@ @<Constants@>=
res_word = 0201 /* returned by |get_output| for reserved words */
section_code = 0200 /* returned by |get_output| for section names */

@ @c
/* returns the next token of output */
func get_output() rune {
restart: 
	for cur_state.tok_field==cur_state.end_field {
		pop_level()
	}
	idx:=cur_state.tok_field
	a:=tok_mem[idx]/* current item read from |tok_mem| */
	cur_state.tok_field++
	if a>=0400 {
		cur_name=a % id_flag
		switch a / id_flag {
			case 2: 
				return res_word /* |a==res_flag+cur_name| */
			case 3: 
				return section_code /* |a==section_flag+cur_name| */
			case 4: 
				push_level(a % id_flag)
				goto restart /* |a==tok_flag+cur_name| */
			case 5: 
				push_level(a % id_flag)
				cur_state.mode_field=inner
				goto restart
				/* |a==inner_tok_flag+cur_name| */
			default: 
				return identifier /* |a==id_flag+cur_name| */
		}
	}
	return a
}

@ The real work associated with token output is done by |make_output|.
This procedure appends an |end_translation| token to the current token list,
and then it repeatedly calls |get_output| and feeds characters to the output
buffer until reaching the |end_translation| sentinel. It is possible for
|make_output| to be called recursively, since a section name may include
embedded \GO/ text; however, the depth of recursion never exceeds one
level, since section names cannot be inside of section names.

A procedure called |output_C| does the scanning, translation, and
output of \GO/ text within `\pb' brackets, and this procedure uses
|make_output| to output the current token list. Thus, the recursive call
of |make_output| actually occurs when |make_output| calls |output_C|
while outputting the name of a section.
@^recursion@>

@c
/* outputs the current token list */
func output_C() {
	save_tok_ptr:=len(tok_mem)
	save_text_ptr:=len(tok_start)
	save_next_control:=next_control/* values to be restored */
	next_control=ignore
	p:=C_translate()/* translation of the \GO/ text */
	app(inner_tok_flag+p)
	if flags['e'] {
		out_str("\\PB{")
		make_output()
		out('}')
@.\\PB@>
	}@+else {
		make_output() /* output the list */
	}
	if len(tok_start)>max_text_ptr {
		max_text_ptr=len(tok_start)
	}
	if len(tok_mem)>max_tok_ptr {
		max_tok_ptr=len(tok_mem)
	}
	tok_start = tok_start[:save_text_ptr]
	tok_mem = tok_mem[:save_tok_ptr]/* forget the tokens */
	next_control=save_next_control /* restore |next_control| to original state */
}

@ Here is \.{GOWEAVE}'s major output handler.

@ @c
/* outputs the equivalents of tokens */
func make_output() {
	var c int /* count of |indent| and |outdent| tokens */
	app(end_translation) /* append a sentinel */
	freeze_text()
	push_level(int32(len(tok_start)-2))
	var b rune
	for true {
		a:=get_output()/* current output byte */
reswitch: 
		switch a {
			case end_translation: 
				return
			case identifier, res_word: 
				@<Output an identifier@>
			case section_code: 
				@<Output a section name@>
			case math_rel: 
				out_str("\\MRL{"@q}@>)
@.\\MRL@>
			case noop,inserted: 
				break
			case cancel, big_cancel: 
				c=0
				b=a
				for true {
					a=get_output()
					if a==inserted {
						continue
					}
					if a<indent && !(b==big_cancel && a==' ') || a>big_force {
						break
					}
					if a==indent { 
						c++ 
					} else if a==outdent {
						c--
					} else if a==opt {
						a=get_output()
					}
				}
				@<Output saved |indent| or |outdent| tokens@>
				goto reswitch
			case indent, outdent, opt, backup, break_space,
				force, big_force, preproc_line: 
					@<Output a control,
				look ahead in case of line breaks, possibly |goto reswitch|@>
			case quoted_char: 
				out(tok_mem[cur_state.tok_field])
				cur_state.tok_field++
			case qualifier:
			default: 
				out(a) /* otherwise |a| is an ordinary character */
		}
	}
}

@ An identifier of length one does not have to be enclosed in braces, and it
looks slightly better if set in a math-italic font instead of a (slightly
narrower) text-italic font. Thus we output `\.{\\\v}\.{a}' but
`\.{\\\\\{aa\}}'.

@<Output an identifier@>=
out('\\')
if a==identifier {
	if name_dir[cur_name].ilk==custom && !doing_format {
	@<Custom out@>
	} else if is_tiny(cur_name) { 
		out('|')
@.\\|@>
	} else { 
		delim:='.'
		for _, v := range name_dir[cur_name].name  {
			if unicode.IsLower(v) { /* not entirely uppercase */
				delim='\\'
				break
			}
		}
		out(delim)
	}
@.\\\\@>
@.\\.@>
}@+else if name_dir[cur_name].ilk==alfop {
	out('X')
	@<Custom out@>
}@+else {
	out('&') /* |a==res_word| */
}
@.\\\&@>
if is_tiny(cur_name) {
	if name_dir[cur_name].name[0]=='_' || name_dir[cur_name].name[0]=='$' {
		out('\\')
	}
	out(name_dir[cur_name].name[0]) 
} else {
	out_name(cur_name,true)
}

@ @<Custom out@>=
for _, v := range name_dir[cur_name].name {
	if v == '_' {
		out('x')
	} else if v == '$' {
		out('X')
	} else {
		out(v)
	}
}
break

@ The current mode does not affect the behavior of \.{GOWEAVE}'s output routine
except when we are outputting control tokens.

@<Output a control...@>=
if a<break_space || a==preproc_line {
	if cur_state.mode_field==outer {
		out('\\')
		out(a-cancel+'0')
@.\\1@>
@.\\2@>
@.\\3@>
@.\\4@>
@.\\8@>
		if a==opt {
			b=get_output(); /* |opt| is followed by a digit */
			if b!='0' || flags['f']==false { 
				out(b) 
			} else {
				out_str("{-1}") /* |flags['f']| encourages more \.{@@\v} breaks */
			}
		}
	} else if a==opt {
		 b=get_output() /* ignore digit following |opt| */
	}
} else {
@<Look ahead for strongest line break, |goto reswitch|@>
}

@ If several of the tokens |break_space|, |force|, |big_force| occur in a
row, possibly mixed with blank spaces (which are ignored),
the largest one is used. A line break also occurs in the output file,
except at the very end of the translation. The very first line break
is suppressed (i.e., a line break that follows `\.{\\Y\\B}').

@<Look ahead for st...@>= {
	b=a
	save_mode:=cur_state.mode_field /* value of |cur_state.mode_field| before a sequence of breaks */
	c=0
	for true {
		a=get_output()
		if a==inserted {
			continue
		}
		if a==cancel || a==big_cancel {
			@<Output saved |indent| or |outdent| tokens@>
			goto reswitch /* |cancel| overrides everything */
		}
		if a!=' ' && a<indent || a==backup || a>big_force {
			if save_mode==outer {
				if out_ptr>3 && compare_runes(out_buf[out_ptr-3:out_ptr + 1],[]rune("\\Y\\B"))==0 {
					goto reswitch
				}
				@<Output saved |indent| or |outdent| tokens@>
				out('\\')
				out(b-cancel+'0')
@.\\5@>
@.\\6@>
@.\\7@>
				if a!=end_translation {
					finish_line()
				}
			} else if a!=end_translation && cur_state.mode_field==inner {
				out(' ')
			}
			goto reswitch
		}
		if a==indent {
			c++
		} else if a==outdent {
			c--
		} else if a==opt {
			a=get_output()
		} else if a>b {
			 b=a /* if |a==' '| we have |a<b| */
		}
	}
}

@ @<Output saved...@>=
	for ;c>0;c-- {
		out_str("\\1")
@.\\1@>
	}
	for ;c<0;c++ {
		out_str("\\2")
@.\\2@>
	}

@ The remaining part of |make_output| is somewhat more complicated. When we
output a section name, we may need to enter the parsing and translation
routines, since the name may contain \GO/ code embedded in
\pb\ constructions. This \GO/ code is placed at the end of the active
input buffer and the translation process uses the end of the active
|tok_mem| area.

@<Output a section name@>= {
	out_str("\\X")
@.\\X@>
	cur_xref=name_dir[cur_name].xref
	if xmem[cur_xref].num==file_flag {
		an_output=true
		cur_xref=xmem[cur_xref].xlink
	} else {
		an_output=false
	}
	if xmem[cur_xref].num>=def_flag {
		out_section(xmem[cur_xref].num-def_flag)
		if phase==3 {
			cur_xref=xmem[cur_xref].xlink
			for xmem[cur_xref].num>=def_flag {
				out_str(", ")
				out_section(xmem[cur_xref].num-def_flag)
				cur_xref=xmem[cur_xref].xlink
			}
		}
	} else {
		out('0') /* output the section number, or zero if it was undefined */
	}
	out(':')
	if an_output {
		out_str("\\.{"@q}@>)
@.\\.@>
	}
	@<Output the text of the section name@>
	if an_output {
		out_str(@q{@>" }") 
	}
	out_str("\\X")
}

@ @<Output the text...@>=
scratch:=sprint_section_name(cur_name) 
cur_section_name:=cur_name
for i := 0; i < len(scratch); {
	b=scratch[i]
	i++
	if b=='@@' {
		@<Skip next character, give error if not `\.{@@}'@>
	}
	if an_output {
		switch b {
			case  ' ','\\','#','%', '$', '^',
			'{','}','~','&','_':
				out('\\')
				fallthrough
@.\\\ @>
@.\\\\@>
@.\\\#@>
@.\\\%@>
@.\\\$@>
@.\\\^@>
@.\\\{@>@q}@>
@q{@>@.\\\}@>
@.\\\~@>
@.\\\&@>
@.\\\_@>
			default: out(b)
		}
	} else if b!='|' {
		out(b)
	} else {
		var buf []rune
		@<Copy the \GO/ text into the |buffer| array@> 
		save_buf:=buffer 
		save_loc:=loc
		buf=append(buf,'|')
		buffer = buf
		loc = 0
		output_C()
		loc=save_loc
		buffer=save_buf
	}
}

@ @<Skip next char...@>=
ii:= i
i++
if ii < len(scratch) && scratch[ii]!='@@' {
	fmt.Print("\n! Illegal control code in section name: <")
@.Illegal control code...@>
	print_section_name(cur_section_name)
	fmt.Print("> ")
	mark_error()
}

@ The \GO/ text enclosed in \pb\ should not contain `\.{\v}' characters,
except within strings. We put a `\.{\v}' at the front of the buffer, so that an
error message that displays the whole buffer will look a little bit sensible.
The variable |delim| is zero outside of strings, otherwise it
equals the delimiter that began the string being copied.

@<Copy the \GO/ text into...@>=
var delim rune
for true {
	if i>=len(scratch) {
		fmt.Print("\n! C text in section name didn't end: <");
@.C text...didn't end@>
		print_section_name(cur_section_name)
		fmt.Print("> ")
		mark_error()
		break
	}
	b=scratch[i]
	i++
	if b=='@@' || b=='\\' && delim!=0 {
		 @<Copy a quoted character into the |buf|@>
	} else {
		if b=='\'' || b=='"' {
			if delim==0 {
				delim=b
			} else if delim==b {
				delim=0
			}
		}
		if b!='|' || delim!=0 {
			buf=append(buf, b)
		} else {
			break
		}
	}
}

@ @<Copy a quoted char...@>= {
	buf = append(buf, b)
	buf = append(buf, scratch[i])
	i++
}

@** Phase two processing.
We have assembled enough pieces of the puzzle in order to be ready to specify
the processing in \.{GOWEAVE}'s main pass over the source file. Phase two
is analogous to phase one, except that more work is involved because we must
actually output the \TEX/ material instead of merely looking at the
\.{CWEB} specifications.

@ @c
func phase_two() {
	reset_input()
	if show_progress() {
		fmt.Print("\nWriting the output file...")
@.Writing the output file...@>
	}
	section_count=0
	format_visible=true
	copy_limbo()
	finish_line()
	flush_buffer(0,false,false) /* insert a blank line, it looks nice */
	for !input_has_ended {
		@<Translate the current section@>
	}
}

@ The output file will contain the control sequence \.{\\Y} between non-null
sections of a section, e.g., between the \TEX/ and definition parts if both
are nonempty. This puts a little white space between the parts when they are
printed. However, we don't want \.{\\Y} to occur between two definitions
within a single section. The variables |out_line| or |out_ptr| will
change if a section is non-null, so the following functions `|save_position|'
and `|emit_space_if_needed|' are able to handle the situation:

@c
func save_position() {
	save_line=out_line
	save_place=out_ptr
}

func emit_space_if_needed() {
	if save_line!=out_line || save_place!=out_ptr {
		out_str("\\Y")
	}
	space_checked=true
@.\\Y@>
}

@ @<Global...@>=
var save_line int /* former value of |out_line| */
var save_place int32 /* former value of |out_ptr| */
var sec_depth int32 /* the integer, if any, following \.{@@*} */
var space_checked bool /* have we done |emit_space_if_needed|? */
var format_visible bool /* should the next format declaration be output? */
var doing_format bool=false /* are we outputting a format declaration? */
var group_found bool=false /* has a starred section occurred? */

@ @<Translate the current section@>= {
	section_count++
	@<Output the code for the beginning of a new section@>
	save_position()
	@<Translate the \TEX/ part of the current section@>
	@<Translate the definition part of the current section@>
	@<Translate the \GO/ part of the current section@>
	@<Show cross-references to this section@>
	@<Output the code for the end of a section@>
}

@ Sections beginning with the \.{CWEB} control sequence `\.{@@\ }' start in the
output with the \TEX/ control sequence `\.{\\M}', followed by the section
number. Similarly, `\.{@@*}' sections lead to the control sequence `\.{\\N}'.
In this case there's an additional parameter, representing one plus the
specified depth, immediately after the \.{\\N}.
If the section has changed, we put \.{\\*} just after the section number.

@<Output the code for the beginning...@>=
if buffer[loc-1]!='*' {
	out_str("\\M")
@.\\M@>
} else {
	for loc < len(buffer) && buffer[loc] == ' ' {
		loc++
	}
	if loc < len(buffer) && buffer[loc]=='*' { /* ``top'' level */
		sec_depth = -1
		loc++
	} else {
		for sec_depth=0; loc < len(buffer) && unicode.IsDigit(buffer[loc]);loc++ {
			sec_depth = sec_depth*10 + buffer[loc] -'0'
		}
	}
	for loc < len(buffer) && buffer[loc] == ' ' {
		loc++ /* remove spaces before group title */
	}
	group_found=true 
	out_str("\\N")
@.\\N@>
	{
		@+s := fmt.Sprintf("{%d}",sec_depth+1)
		@+out_str(s)
	@+}
	if show_progress() {
		fmt.Printf("*%d",section_count)
	}
	os.Stdout.Sync() /* print a progress report */
}
out_str("{")
out_section(section_count)
out_str("}")

@ In the \TEX/ part of a section, we simply copy the source text, except that
index entries are not copied and \GO/ text within \pb\ is translated.

@<Translate the \T...@>= 
for true {
	next_control=copy_TeX()
	switch next_control {
		case '|': 
			init_stack()
			output_C()
		case '@@': 
			out('@@')
		case TeX_string, noop, xref_roman, xref_wildcard, xref_typewriter, section_name: 
			loc-=2
			next_control=get_next() /* skip to \.{@@>} */
			if next_control==TeX_string {
				err_print("! TeX string should be in C text only")
@.TeX string should be...@>
			}
		case thin_space,math_break,ord,
		line_break, big_line_break, no_line_break, join,
		pseudo_semi, macro_arg_open, macro_arg_close,
		output_defs_code:
				err_print("! You can't do that in TeX text")
@.You can't do that...@>
	}
	if next_control>=format_code {
		 break
	}
}

@ When we get to the following code we have |next_control>=format_code|, and
the token memory is in its initial empty state.

@<Translate the d...@>=
space_checked=false
for next_control<=definition { /* |format_code| or |definition| */
	init_stack() 
	if next_control==definition {
		 @<Start a macro definition@>
	} else {
		@<Start a format definition@>
	}
	outer_parse()
	finish_C(format_visible)
	format_visible=true
	doing_format=false
}

@ The |finish_C| procedure outputs the translation of the current
scraps, preceded by the control sequence `\.{\\B}' and followed by the
control sequence `\.{\\par}'. It also restores the token and scrap
memories to their initial empty state.

A |force| token is appended to the current scraps before translation
takes place, so that the translation will normally end with \.{\\6} or
\.{\\7} (the \TEX/ macros for |force| and |big_force|). This \.{\\6} or
\.{\\7} is replaced by the concluding \.{\\par} or by \.{\\Y\\par}.

@ @c
/* finishes a definition or a \GO/ part */
/* visible is nonzero if we should produce \TEX/ output */
func finish_C(visible bool) {
	if visible {
		out_str("\\B")
		app_tok(force)
		app_scrap(insert,no_math)
		p:=translate() /* translation of the scraps */
@.\\B@>
		app(tok_flag+p)
		make_output() /* output the list */
		if out_ptr>1 {
			if out_buf[out_ptr-1]=='\\' {
@.\\6@>
@.\\7@>
@.\\Y@>
				if out_buf[out_ptr]=='6' {
					out_ptr-=2
				} else if out_buf[out_ptr]=='7' {
					out_buf[out_ptr]='Y'
				}
			}
		}
		out_str("\\par")
		finish_line()
	}
	if len(tok_start)>max_text_ptr {
		max_text_ptr=len(tok_start)
	}
	if len(tok_mem)>max_tok_ptr {
		max_tok_ptr=len(tok_mem) 
	}
	if scrap_ptr>max_scr_ptr {
		max_scr_ptr=scrap_ptr
	}
	tok_mem=tok_mem[:0]
	tok_start=tok_start[:1]
	scrap_ptr=0
		/* forget the tokens and the scraps */
}

@ Keeping in line with the conventions of the \GO/ preprocessor (and
otherwise contrary to the rules of \.{CWEB}) we distinguish here
between the case that `\.(' immediately follows an identifier and the
case that the two are separated by a space.  In the latter case, and
if the identifier is not followed by `\.(' at all, the replacement
text starts immediately after the identifier.  In the former case,
it starts after we scan the matching `\.)'.

@<Start a macro...@>= {
	if save_line!=out_line || save_place!=out_ptr || space_checked {
		app(backup)
	}
	if !space_checked {
		emit_space_if_needed()
		save_position()
	}
	app_str("\\D") /* this will produce `\&{define }' */
@.\\D@>
	if next_control=get_next(); next_control!=identifier {
		err_print("! Improper macro definition")
@.Improper macro definition@>
	} else {
		app('$')
		app_cur_id(false)
		if loc < len(buffer) && buffer[loc]=='(' {
reswitch: 
			next_control=get_next()
			switch next_control {
				case '(', ',': 
					app(next_control)
					goto reswitch
				case identifier:
					app_cur_id(false)
					goto reswitch
				case ')': 
					app(next_control)
					next_control=get_next()
				default: 
					err_print("! Improper macro definition")
			}
		} else {
			next_control=get_next()
		}
		app_str("$ ")
		app(break_space)
		app_scrap(dead,no_math) /* scrap won't take part in the parsing */
	}
}

@ @<Start a format...@>= {
	doing_format=true
	if buffer[loc-1]=='s' || buffer[loc-1]=='S' {
		format_visible=false
	}
	if !space_checked {
		emit_space_if_needed()
		save_position()
	}
	app_str("\\F") /* this will produce `\&{format }' */
@.\\F@>
	next_control=get_next()
	if next_control==identifier {
		app(id_flag+id_lookup(id,normal))
		app(' ')
		app(break_space) /* this is syntactically separate from what follows */
		next_control=get_next()
		if next_control==identifier {
			app(id_flag+id_lookup(id,normal))
			app_scrap(exp,maybe_math)
			app_scrap(semi,maybe_math)
			next_control=get_next()
		}
	}
	if scrap_ptr!=2 {
		err_print("! Improper format definition")
@.Improper format definition@>
	}
}

@ Finally, when the \TEX/ and definition parts have been treated, we have
|next_control>=begin_code|. We will make the global variable |this_section|
point to the current section name, if it has a name.

@<Global...@>=
var this_section int32 /* the current section name, or zero */

@ @<Translate the \GO/...@>=
this_section=-1
if next_control<=section_name {
	emit_space_if_needed()
	init_stack()
	if next_control==begin_code {
		next_control=get_next()
	} else {
		this_section=cur_section
		@<Check that '=' or '==' follows this section name, and
			emit the scraps to start the section definition@>
	}
	for next_control<=section_name {
		outer_parse()
		@<Emit the scrap for a section name if present@>
	}
	finish_C(true)
}

@ The title of the section and an $\E$ or $\mathrel+\E$ are made
into a scrap that should not take part in the parsing.

@<Check that '='...@>=
for true {
	next_control=get_next()
	if next_control!='+' {
		break
	}
} /* allow optional `\.{+=}' */
if next_control!='=' && next_control!=eq_eq {
	err_print("! You need an = sign after the section name")
@.You need an = sign...@>
} else {
	next_control=get_next()
}
if out_ptr>1 && out_buf[out_ptr]=='Y' && out_buf[out_ptr-1]=='\\' {
	app(backup)
}
		/* the section name will be flush left */
@.\\Y@>
app(section_flag+this_section)
cur_xref=name_dir[this_section].xref
if xmem[cur_xref].num==file_flag {
	cur_xref=xmem[cur_xref].xlink
}
app_str("${}")
if xmem[cur_xref].num!=section_count+def_flag {
	app_str("\\mathrel+") /*section name is multiply defined*/
	this_section=-1 /*so we won't give cross-reference info here*/
}
app_str("\\E") /* output an equivalence sign */
@.\\E@>
app_str("{}$")
app(force)
app_scrap(dead,no_math)
				/* this forces a line break unless `\.{@@+}' follows */

@ @<Emit the scrap...@>=
if next_control<section_name {
	err_print("! You can't do that in C text")
@.You can't do that...@>
	next_control=get_next()
} else if next_control==section_name {
	app(section_flag+cur_section)
	app_scrap(section_scrap,maybe_math)
	next_control=get_next()
}

@ Cross references relating to a named section are given
after the section ends.

@<Show cross...@>=
if this_section!=-1 {
	cur_xref=name_dir[this_section].xref
	if xmem[cur_xref].num==file_flag {
		an_output=true 
		cur_xref=xmem[cur_xref].xlink
	} else {
		an_output=false
	}
	if xmem[cur_xref].num>def_flag {
		cur_xref=xmem[cur_xref].xlink /* bypass current section number */
	}
	footnote(def_flag)
	footnote(cite_flag)
	footnote(0)
}

@ The |footnote| procedure gives cross-reference information about
multiply defined section names (if the |flag| parameter is
|def_flag|), or about references to a section name
(if |flag==cite_flag|), or to its uses (if |flag==0|). It assumes that
|cur_xref| points to the first cross-reference entry of interest, and it
leaves |cur_xref| pointing to the first element not printed.  Typical outputs:
`\.{\\A101.}'; `\.{\\Us 370\\ET1009.}';
`\.{\\As 8, 27\\*\\ETs64.}'.

Note that the output of \.{GOWEAVE} is not English-specific; users may
supply new definitions for the macros \.{\\A}, \.{\\As}, etc.

@ @c
/* outputs section cross-references */
func footnote(flag int32) {
	if xmem[cur_xref].num<=flag {
		return
	}
	finish_line()
	out('\\')
@.\\A@>
@.\\Q@>
@.\\U@>
	switch flag {
		case 0:
			out('U')
		case cite_flag:
			out('Q')
		default:
			out('A')
	}
	@<Output all the section numbers on the reference list |cur_xref|@>
	out('.')
}

@ The following code distinguishes three cases, according as the number
of cross-references is one, two, or more than two. Variable |q| points
to the first cross-reference, and the last link is a zero.

@<Output all the section numbers...@>=
q:=cur_xref /* cross-reference pointer variable */
if xmem[xmem[q].xlink].num>flag {
	out('s') /* plural */
}
for true {
	out_section(xmem[cur_xref].num-flag)
	cur_xref=xmem[cur_xref].xlink /* point to the next cross-reference to output */
	if xmem[cur_xref].num<=flag {
		break
	}
	if xmem[xmem[cur_xref].xlink].num>flag {
		out_str(", ") /* not the last */
	} else {
		out_str("\\ET") /* the last */
@.\\ET@>
		if cur_xref != xmem[q].xlink {
			 out('s') /* the last of more than two */
		}
	}
}

@ @<Output the code for the end of a section@>=
out_str("\\fi")
finish_line()
@.\\fi@>
flush_buffer(0,false,false) /* insert a blank line, it looks nice */

@** Phase three processing.
We are nearly finished! \.{GOWEAVE}'s only remaining task is to write out the
index, after sorting the identifiers and index entries.

If the user has set the |flags['x']==0| flag (the \.{-x} option on the command line),
just finish off the page, omitting the index, section name list, and table of
contents.

@ @c
func phase_three() {
	if !flags['x'] {
		finish_line()
		out_str("\\end")
@.\\end@>
		finish_line()
	} else {
		phase=3
		if show_progress() {
			 fmt.Print("\nWriting the index...")
@.Writing the index...@>
		}
		finish_line()
		if f, err := os.OpenFile(idx_file_name, 
			os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666); err != nil {
			fatal("! Cannot open index file ",idx_file_name)
@.Cannot open index file@>
		} else {
			idx_file = f
		}
		if change_exists {
			@<Tell about changed sections@>
			finish_line()
			finish_line()
		}
		out_str("\\inx")
		finish_line()
@.\\inx@>
		active_file=idx_file /* change active file to the index file */ 
		@<Do the first pass of sorting@>
		@<Sort and output the index@>
		finish_line()
		active_file.Close() /* finished with |idx_file| */
		active_file=tex_file /* switch back to |tex_file| for a tic */
		out_str("\\fin")
		finish_line()
@.\\fin@>
		if f, err := os.OpenFile(scn_file_name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666); 
			err != nil {
			fatal("! Cannot open section file ",scn_file_name);
@.Cannot open section file@>
		} else {
			scn_file = f
		}
		active_file=scn_file /* change active file to section listing file */
		@<Output all the section names@>
		finish_line()
		active_file.Close() /* finished with |scn_file| */
		active_file=tex_file
		if group_found {
			out_str("\\con")
@.\\con@>
		} @+else {
			out_str("\\end")
@.\\end@>
		}
		finish_line()
		active_file.Close()
	}
	if show_happiness() {
		fmt.Print("\nDone.")
	}
	check_complete() /* was all of the change file used? */
}

@ Just before the index comes a list of all the changed sections, including
the index section itself.

@ @<Tell about changed sections@>= {
	/* remember that the index is already marked as changed */
	var k_section int32 = 0/* runs through the sections */
	for k_section++;!changed_section[k_section];k_section++ {}
	out_str("\\ch ")
@.\\ch@>
	out_section(k_section)
	for k_section<section_count {
		for k_section++;!changed_section[k_section];k_section++ {}
		out_str(", ")
		out_section(k_section)
	}
	out('.')
}

@ A left-to-right radix sorting method is used, since this makes it easy to
adjust the collating sequence and since the running time will be at worst
proportional to the total length of all entries in the index. We put the
identifiers into 102 different lists based on their first characters.
(Uppercase letters are put into the same list as the corresponding lowercase
letters, since we want to have `$t<\\{TeX}<\&{to}$'.) The
list for character |c| begins at location |bucket[c]| and continues through
the |blink| array.

@<Global...@>=
var bucket [256]int32
var blink [max_names]int32 /* links in the buckets */

@ To begin the sorting, we go through all the hash lists and put each entry
having a nonempty cross-reference list into the proper bucket.

@<Do the first pass...@>= {
for c:=0; c<=255; c++ {
	bucket[c]=-1
}
for _, next_name := range hash {
	for next_name != -1 {
		cur_name=next_name
		next_name=name_dir[cur_name].llink
		if name_dir[cur_name].xref!=0 {
			c:=name_dir[cur_name].name[0]
			if unicode.IsUpper(c) {
				c=unicode.ToLower(c)
			}
			blink[cur_name]=bucket[c]
			bucket[c]=cur_name
		}
	}
}
}

@ During the sorting phase we shall use the |cat| and |trans_plus.Trans| arrays from
\.{GOWEAVE}'s parsing algorithm and rename them |depth| and |head|. They now
represent a stack of identifier lists for all the index entries that have
not yet been output. The variable |sort_ptr| tells how many such lists are
present; the lists are output in reverse order (first |sort_ptr|, then
|sort_ptr-1|, etc.). The |j|th list starts at |head[j]|, and if the first
|k| characters of all entries on this list are known to be equal we have
|depth[j]==k|.

@ @<Rest of |trans| struct@>=
Head int32

@ @<Type...@>=
type sort_pointer int32

@ @<Constants@>=
max_sorts = max_scraps /* ditto */

@ @f sort_pointer int


@ @<Global...@>=
var cur_depth int32 /* depth of current buckets */
var cur_byte int32 /* index into |byte_mem| */
var cur_val int32 /* current cross-reference number */
var max_sort_ptr int32 /* largest value of |sort_ptr| */
var sort_ptr int32 /* ditto */

@ @<Set init...@>=
max_sort_ptr=0

@ The desired alphabetic order is specified by the |collate| array; namely,
$|collate|[0]<|collate|[1]<\cdots<|collate|[100]$.

@<Global...@>=
/* collation order */
var collate = [102+128] rune {
0, ' ',001,002,003,004,005,006,007,010,011,012,013,014,015,016,017,
020,021,022,023,024,025,026,027,030,031,032,033,034,035,036,037,
'!',042,'#','$','%','&','\'','(',')','*','+',',','-','.','/',':',
';','<','=','>','?','@@','[','\\',']','^','`','{','|','}','~','_',
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
@^high-bit character handling@>

@ We use the order $\hbox{null}<\.\ <\hbox{other characters}<{}$\.\_${}<
\.A=\.a<\cdots<\.Z=\.z<\.0<\cdots<\.9.$ Warning: The collation mapping
needs to be changed if ASCII code is not being used.
@^ASCII code dependencies@>
@^high-bit character handling@>

We initialize |collate| by copying a few characters at a time, because
some \GO/ compilers choke on long strings.

@ Procedure |unbucket| goes through the buckets and adds nonempty lists
to the stack, using the collating sequence specified in the |collate| array.
The parameter to |unbucket| tells the current depth in the buckets.
Any two sequences that agree in their first 255 character positions are
regarded as identical.

@<Constants@>=
infinity = -1  /* $\infty$ (approximately) */

@ @c
/* empties buckets having depth |d| */
func unbucket(d int32) {
	/* index into |bucket|; cannot be a simple |char| because of sign
		comparison below*/
	for c:=100+128; c>= 0; c-- {
		if bucket[collate[c]] != -1 {
@^high-bit character handling@>
			sort_ptr++
			if sort_ptr>max_sort_ptr {
				max_sort_ptr=sort_ptr
			}
			if c==0 {
				scrap_info[sort_ptr].cat=infinity
			} else {
				scrap_info[sort_ptr].cat=d
			}
			scrap_info[sort_ptr].trans_plus.Head=bucket[collate[c]]
			bucket[collate[c]]=-1
		}
	}
}

@ @<Sort and output...@>=
sort_ptr=0
unbucket(1)
for sort_ptr>0 {
	cur_depth=scrap_info[sort_ptr].cat
	if blink[scrap_info[sort_ptr].trans_plus.Head]==-1 || cur_depth==infinity {
		@<Output index entries for the list at |sort_ptr|@>
	} else {
		@<Split the list at |sort_ptr| into further lists@> 
	}
}

@ @<Split the list...@>= {
	next_name:=scrap_info[sort_ptr].trans_plus.Head
	for true {
		var c rune
		cur_name=next_name
		next_name=blink[cur_name]
		cur_byte=cur_depth
		if cur_byte>=int32(len(name_dir[cur_name].name)) {
			c=0 /* hit end of the name */
		} else {
			c=name_dir[cur_name].name[cur_byte]
			if unicode.IsUpper(c) {
				c=unicode.ToLower(c)
			}
		}
		blink[cur_name]=bucket[c]
		bucket[c]=cur_name
		if next_name == -1 {
			break
		}
	}
	sort_ptr--
	unbucket(cur_depth+1)
}

@ @<Output index...@>= {
	cur_name=scrap_info[sort_ptr].trans_plus.Head
	for true {
		out_str("\\I")
@.\\I@>
		@<Output the name at |cur_name|@>
		@<Output the cross-references at |cur_name|@>
		cur_name=blink[cur_name]
		if cur_name == -1 {
			break
		}
	}
	sort_ptr--
}

@ @<Output the name...@>=
switch name_dir[cur_name].ilk {
	case normal, func_template: 
		if is_tiny(cur_name) {
			out_str("\\|")
@.\\|@>
		} else {
			lowcase := false
			for _, v := range name_dir[cur_name].name {
				if unicode.IsLower(v) {
					lowcase = true
					break
				}
			}
			if !lowcase {
				out_str("\\.")
@.\\.@>
			} else {
				out_str("\\\\")
@.\\\\@>
			}
		}
	case wildcard: 
		out_str("\\9");
		out_name(cur_name,false)
		goto name_done
@.\\9@>
	case typewriter: 
		out_str("\\.");
@.\\.@>
		fallthrough 
	case roman:
		out_name(cur_name,false)
		goto name_done;
	case custom: {
		out_str("$\\")
		for _, v := range name_dir[cur_name].name {
			if v == '_' {
				out('x')
			} else if v == '$' {
				out('X')
			} else {
				out(v)
			}
		}
		out('$')
		goto name_done
	}
	default: 
		out_str("\\&")
@.\\\&@>
}
out_name(cur_name,true)
name_done:@

@ Section numbers that are to be underlined are enclosed in
`\.{\\[}$\,\ldots\,$\.]'.

@<Output the cross-references...@>=
@<Invert the cross-reference list at |cur_name|, making |cur_xref| the head@>
for true {
	out_str(", ")
	cur_val=xmem[cur_xref].num
	if cur_val<def_flag {
		out_section(cur_val)
	} else {
		out_str("\\[")
		out_section(cur_val-def_flag)
		out(']')
	}
@.\\[@>
	cur_xref=xmem[cur_xref].xlink
	if cur_xref==0 {
		break
	}
}
out('.')
finish_line()

@ List inversion is best thought of as popping elements off one stack and
pushing them onto another. In this case |cur_xref| will be the head of
the stack that we push things onto.
@<Global...@>=
var next_xref int32 
var this_xref int32
	/* pointer variables for rearranging a list */

@ @<Invert the cross-reference list at |cur_name|, making |cur_xref| the head@>=
this_xref=name_dir[cur_name].xref
cur_xref=0
for true {
	next_xref=xmem[this_xref].xlink
	xmem[this_xref].xlink=cur_xref
	cur_xref=this_xref
	this_xref=next_xref
	if this_xref==0 {
		break
	}
}

@ The following recursive procedure walks through the tree of section names and
prints them.
@^recursion@>

@ @c
/* print all section names in subtree |p| */
func section_print(p int32) {
	if p != -1{
		section_print(name_dir[p].llink)
		out_str("\\I")
@.\\I@>
		tok_mem=tok_mem[:0]
		tok_start=tok_start[:1]
		scrap_ptr=0
		init_stack()
		app(p+section_flag)
		make_output()
		footnote(cite_flag)
		footnote(0) /* |cur_xref| was set by |make_output| */
		finish_line() @/
		section_print(name_dir[p].rlink)
	}
}

@ @<Output all the section names@>=section_print(name_root)

@ Because on some systems the difference between two pointers is a |long|
rather than an |int|, we use \.{\%ld} to print these quantities.

@c
func print_stats() {
	fmt.Println("\nMemory usage statistics:\n");
@.Memory usage statistics:@>
	fmt.Println("%v names", len(name_dir))
	fmt.Println("Parsing:")
	fmt.Println("%v scraps", max_scr_ptr)
	fmt.Println("%v texts", max_text_ptr)
	fmt.Println("%v tokens", max_tok_ptr)
	fmt.Println("%v levels", max_stack_ptr)
	fmt.Println("Sorting:")
	fmt.Println("%v levels ",max_sort_ptr)
}

@ @<Print usage error message and quit@>=
{
	fatal("! Usage: goweave [options] webfile[.w] [{changefile[.ch]|-} [outfile[.tex]]]\n", "")
	@.Usage:@>
}

@ \.{GOWEAVE} specific creation of output file

@<Try to open output file@>=
if f, err := os.OpenFile(tex_file_name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666); 
	err != nil {
	fatal("! Cannot open output file ", tex_file_name)
} else {
	tex_file = f
}


@** Index.
If you have read and understood the code for Phase III above, you know what
is in this index and how it got here. All sections in which an identifier is
used are listed with that identifier, except that reserved words are
indexed only when they appear in format definitions, and the appearances
of identifiers in section names are not indexed. Underlined entries
correspond to where the identifier was declared. Error messages, control
sequences put into the output, and a few
other things like ``recursion'' are indexed here too.
