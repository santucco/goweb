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

% Here is TeX material that gets inserted after \input gowebmac
\def\hang{\hangindent 3em\indent\ignorespaces}
\def\pb{$\.|\ldots\.|$} % code brackets (|...|)
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
const banner = "This is GOWEAVE (Version 0.1)\n"

@ @c
package main

import (
@<Import packages@>@/
)

@<Typedef declarations@>@/

@<Constants@>@/


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
const (
	max_names = 4000 /* number of identifiers, strings, section names
		must be less than 10240; used in |"common.w"| */
	line_length = 80 /* lines of \TEX/ output have at most this many characters
		should be less than 256 */
)

@ The next few sections contain stuff from the file |"common.w"| that must
be included in both |"gotangle.w"| and |"goweave.w"|. 

@i common.w

@* Data structures exclusive to {\tt GOWEAVE}.
As explained in \.{common.w}, the field of a |name_info| structure
that contains the |rlink| of a section name is used for a completely
different purpose in the case of identifiers. It is then called the
|ilk| of the identifier, and it is used to
distinguish between various types of identifiers, as follows:

\yskip\hang |normal| identifiers are part of the
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

\yskip\hang |alfop|, \dots
identifiers are \GO/ reserved words whose |ilk|
explains how they are to be treated when \GO/ code is being
formatted.

@<More elements of |name_info| structure@>=
	ilk		int32 /* used by identifiers in \.{GOWEAVE} only */

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
const (
	cite_flag = 10240 /* must be strictly larger than |max_sections| */
	file_flag = 3*cite_flag
	def_flag = 2*cite_flag
)

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
// reserved words
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

// types
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

// constants
id_lookup([]rune("true"),Type)
id_lookup([]rune("false"),Type)
id_lookup([]rune("iota"),Expression)

// zero value
id_lookup([]rune("nil"),identifier)

// functions
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
res_wd_end=int32(len(name_dir))
id_lookup([]rune("TeX"),custom)

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
all. Some of these numeric control codes take the place of |rune|
control codes that will not otherwise appear in the output of the
scanning routines.
@^ASCII code dependencies@>

@ @<Constants@>=
const (
	ignore rune = 00 /* control code of no interest to \.{GOWEAVE} */
	verbatim rune = 02 /* takes the place of extended ASCII \.{\char2} */
	underline rune = '\n' /* this code will be intercepted without confusion */
	noop rune = 0177 /* takes the place of ASCII delete */
	xref_roman rune = 0213 /* control code for `\.{@@\^}' */
	xref_wildcard rune = 0214 /* control code for `\.{@@:}' */
	xref_typewriter rune = 0215 /* control code for `\.{@@.}' */
	TeX_string rune = 0216 /* control code for `\.{@@t}' */
	ord rune = 0217 /* control code for `\.{@@'}' */
	join rune = 0220 /* control code for `\.{@@\&}' */
	thin_space rune = 0221 /* control code for `\.{@@,}' */
	math_break rune = 0222 /* control code for `\.{@@\v}' */
	line_break rune = 0223 /* control code for `\.{@@/}' */
	big_line_break rune = 0224 /* control code for `\.{@@\#}' */
	no_line_break rune = 0225 /* control code for `\.{@@+}' */
	pseudo_semi rune = 0226 /* control code for `\.{@@;}' */
	trace rune = 0232 /* control code for `\.{@@0}', `\.{@@1}' and `\.{@@2}' */
	format_code rune = 0235 /* control code for `\.{@@f}' and `\.{@@s}' */
	begin_code rune = 0237 /* control code for `\.{@@c}' */
	section_name rune = 0240 /* control code for `\.{@@<}' */
	new_section rune = 0241 /* control code for `\.{@@\ }' and `\.{@@*}' */
)

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
ccode['\'']=ord
@<Special control codes for debugging@>

@ Users can write
\.{@@2}, \.{@@1}, and \.{@@0} to turn tracing fully on, partly on,
and off, respectively.

@<Special control codes...@>=
ccode['0']=trace
ccode['1']=trace
ccode['2']=trace
ccode['3']=trace
ccode['4']=trace
ccode['5']=trace
ccode['6']=trace
ccode['7']=trace

@ The |skip_limbo| routine is used on the first pass to skip through
portions of the input that are not in any sections, i.e., that precede
the first section. After this procedure has been called, the value of
|input_has_ended| will tell whether or not a section has actually been found.

There's a complication that we will postpone until later: If the \.{@@s}
operation appears in limbo, we want to use it to adjust the default
interpretation of identifiers.

@ @c
func skip_limbo() {
	for {
		if loc>=len(buffer) && !get_line() {
			return
		}
		for loc < len(buffer) && buffer[loc]!='@@' {
			loc++ /* look for '@@', then skip two symbols */
		}
		l := loc
		loc++
		if l<len(buffer) { 
			c:=new_section
			if loc < len(buffer) {
				c=ccode[buffer[loc]]
				loc++
			}
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
	for {
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
const (
	constant rune = 0210 /* \GO/ constant */
	str rune = 0211 /* \GO/ string */
	identifier rune = 0212 /* \GO/ identifier or reserved word */
)

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
	for {
		if loc>=len(buffer) {
			if next_control==identifier ||
				next_control==constant ||
				next_control==str ||
				next_control==break_token ||
				next_control==continue_token ||
				next_control==fallthrough_token ||
				next_control==return_token ||
				next_control==plus_plus ||
				next_control==minus_minus ||
				next_control==rpar ||
				next_control==rbracket ||
				next_control==rbrace {
				return pseudo_semi
			}
			if !get_line() {
				return new_section 
			}
		}
		@+c:=buffer[loc] /* the current character */
		loc++
		nc:=' '
		if loc < len(buffer) {
			nc = buffer[loc]
		}
		if unicode.IsDigit(c) || ( c=='.' && unicode.IsDigit(nc)) {
			@<Get a constant@>
		} else if c=='\'' || c=='"' || c=='`' {
			@<Get a string@>
		} else if unicode.IsLetter(c) || 
				c=='_' && (unicode.IsLetter(c) || unicode.IsDigit(c)) {
			@<Get an identifier@>
		} else if c=='@@' {
			@<Get control code and possible section name@>
		} else if unicode.IsSpace(c) {
			continue /* ignore spaces and tabs */
		}
mistake: 
		@<Compress two-symbol operator@>
		return c
	}
	return 0
}

@ The following code assigns values to the combinations \.{++},
\.{--}, \.{>=}, \.{<=}, \.{==}, \.{<<}, \.{>>}, \.{!=}, \.{\v\v}, and
\.{\&\&}, \.{...}.
The compound assignment operators (e.g., \.{+=}) are
treated as separate tokens.

@ @<Get an identifier@>= {
	loc--
	id_first:=loc
	for loc < len(buffer) && 
		(unicode.IsLetter(buffer[loc]) || 
		unicode.IsDigit(buffer[loc]) || 
		buffer[loc]=='_') {
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
		if loc < len(buffer) && buffer[loc]=='i' {
			id = append(id, '$')
			id = append(id, 'i')
			loc++
		}
	}
	return constant
}

@ \GO/ strings and character constants, delimited by double and single
quotes, respectively, can contain newlines or instances of their own
delimiters if they are protected by a backslash.

@<Get a string@>= {
	delim := c /* what started the string */
	section_text = section_text[0:0]

	if delim=='\'' && 
		loc-2<len(buffer) && loc-2>=0 && buffer[loc-2]=='@@' {
		section_text = append(section_text, '@@')
		section_text = append(section_text, '@@')
	}
	section_text = append(section_text, delim)
	if delim=='<' {
		 delim='>' /* for file names in |#include| lines */
	}
	for {
		if loc>=len(buffer) {
			if !get_line()  {
				err_print("! Input ended in middle of string")
				loc=0
				break
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
for {
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
	if loc-1 < len(buffer) && buffer[loc-1]=='*' && show_progress() {
		fmt.Printf("*%d",section_count)
		os.Stdout.Sync() /* print a progress report */
	}
	@<Store cross-references in the \TEX/ part of a section@>
	@<Store cross-references in the format definition part of a section@>
	@<Store cross-references in the \GO/ part of a section@>
	if changed_section[section_count] {
		change_exists=true
	}
}

@ The |Go_xref| subroutine stores references to identifiers in
\GO/ text material beginning with the current value of |next_control|
and continuing until |next_control| is `\.\{' or `\.{\v}', or until the next
``milestone'' is passed (i.e., |next_control>=format_code|). If
|next_control>=format_code| when |Go_xref| is called, nothing will happen;
but if |next_control=='|'| upon entry, the procedure assumes that this is
the `\.{\v}' preceding \GO/ text that is to be processed.

The parameter |spec_ctrl| is used to change this behavior. In most cases
|Go_xref| is called with |spec_ctrl==ignore|, which triggers the default
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
func Go_xref(spec_ctrl rune) {
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

@ The |outer_xref| subroutine is like |Go_xref| except that it begins
with |next_control!='|'| and ends with |next_control>=format_code|. Thus, it
handles \GO/ text with embedded comments.

@ @c
/* extension of |Go_xref| */
func outer_xref() {
	for next_control<format_code {
		if next_control!=begin_comment && next_control!=begin_short_comment {
			Go_xref(ignore)
		} else {
			is_long_comment:=(next_control==begin_comment)
			bal,res:=copy_comment(is_long_comment,1,nil)/* brace level in comment */
			next_control='|'
			for bal>0 {
				Go_xref(section_name) /* do not reference section names in comments */
				if next_control=='|' {
					 bal,res=copy_comment(is_long_comment,bal,res)
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
for {
	next_control=skip_TeX()
	switch next_control {
		case underline:
			xref_switch=def_flag
			continue
		case trace: 
			tracing=buffer[loc-1]-'0'
			continue
		case '|': 
			Go_xref(section_name)
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

@<Store cross-references in the format d...@>=
for next_control<=format_code { 
	@<Process a format definition@>
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
		err_print("! Missing left identifier of @@s")
@.Missing left identifier...@>
	} else {
		lhs=id_lookup(id,normal)
		if get_next()!=identifier {
			err_print("! Missing right identifier of @@s")
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
	for {
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
`\.{\\input gowebmac}'.

@<Set init...@>=
out_ptr=1
out_line=1
active_file=tex_file
out_buf[out_ptr]='c'
fmt.Fprint(active_file,"\\input gowebma") 

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
	for {
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
		if v=='_' && quote_xalpha {
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
	for {
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
			c:=' '
			if loc < len(buffer) {
				c=buffer[loc]
				loc++
			}
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
	for {
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

@ @c
/* copies \TEX/ code in comments */
func copy_comment(
	is_long_comment bool,
	bal int /* brace balance */,
	tok_mem []interface{} ) (int,[]interface{}) {
	for {
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
			return bal,tok_mem
		}
		if is_long_comment {
			@<Check for end of comment@>
		}
		if phase==2 {
			if c>0177 {
				tok_mem=append(tok_mem,quoted_char)
			}
			tok_mem=append(tok_mem,c)
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
		tok_mem=append(tok_mem,buffer[loc])
	}
	loc++
}

@ We output
enough right braces to keep \TEX/ happy.

@<Clear |bal|...@>=
if phase==2 {
	for bal--; bal>=0; bal-- {
		tok_mem=append(tok_mem,'}')
	}
}
return 0,tok_mem

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
$$\hbox{|exp| }\left\{\matrix{\hbox{|binary_op|}}\right\}
\hbox{ |exp| }\RA\hbox{ |exp|}$$
and it means that three consecutive scraps whose respective categories are
|exp|, |binary_op|
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
example, if the current sequence of scraps is |int| |cast|
|lbrace|, rule 31 is applied; but if the sequence is |int| |cast|
followed by anything other than |lbrace|, rule 32 takes effect.

Translation rules such as `$E_1C\,\\{opt}9\,E_2$' above use subscripts
to distinguish between translations of scraps whose categories have the
same initial letter; these subscripts are assigned from left to right.

@ Here is a list of the category codes that scraps can have.
(A few others, like |int|, have already been defined; the
|cat_name| array contains a complete list.)

@<Constants@>=
const (
	normal rune = iota /* ordinary identifiers have |normal| ilk */
	roman rune = iota /* normal index entries have |roman| ilk */
	wildcard rune = iota /* user-formatted index entries have |wildcard| ilk */
	typewriter rune = iota /* `typewriter type' entries have |typewriter| ilk */
	custom rune = iota /* identifiers with user-given control sequence */
)

const (
	zero rune = iota
	ArrayType rune = iota
	StructType rune = iota
	PointerType rune = iota
	InterfaceType rune = iota
	SliceType rune = iota
	MapType rune = iota
	ChannelType rune = iota
	FieldDecl rune = iota
	AnonymousField rune = iota
	Signature rune = iota
	Parameters rune = iota
	ParameterList rune = iota
	ParameterDecl rune = iota
	MethodSpec rune = iota
	Block rune = iota
	Statement rune = iota
	ConstDecl rune = iota
	TypeDecl rune = iota
	VarDecl rune = iota
	FunctionDecl rune = iota
	MethodDecl rune = iota
	ConstSpec rune = iota
	IdentifierList rune = iota
	ExpressionList rune = iota
	TypeSpec rune = iota
	VarSpec rune = iota
	ShortVarDecl rune = iota
	Receiver rune = iota
	Operand rune = iota
	QualifiedIdent rune = iota
	MethodExpr rune = iota
	CompositeLit rune = iota
	FunctionLit rune = iota
	FunctionType rune = iota
	LiteralType rune = iota
	LiteralValue rune = iota
	ElementList rune = iota
	Element rune = iota
	PrimaryExpr rune = iota
	Conversion rune = iota
	BuiltinCall rune = iota
	Selector rune = iota
	Index rune = iota
	Slice rune = iota
	TypeAssertion rune = iota
	Call rune = iota
	Expression rune = iota
	UnaryExpr rune = iota
	ReceiverType rune = iota
	LabeledStmt rune = iota
	SimpleStmt rune = iota
	GoStmt rune = iota
	ReturnStmt rune = iota
	BreakStmt rune = iota
	ContinueStmt rune = iota
	GotoStmt rune = iota
	fallthrough_token rune = iota
	IfStmt rune = iota
	SelectStmt rune = iota
	ForStmt rune = iota
	DeferStmt rune = iota
	SendStmt rune = iota
	IncDecStmt rune = iota
	Assignment rune = iota
	ExprSwitchStmt rune = iota
	ExprCaseClause rune = iota
	TypeSwitchStmt rune = iota
	TypeSwitchGuard rune = iota
	TypeCaseClause rune = iota
	TypeSwitchCase rune = iota
	ForClause rune = iota
	RangeClause rune = iota
	CommClause rune = iota
	CommCase rune = iota
	RecvStmt rune = iota
	BuiltinArgs rune = iota
	PackageClause rune = iota
	PackageName rune = iota
	ImportDecl rune = iota
	ImportSpec rune = iota
	Type rune = iota
	package_token rune = iota /* denotes \.{package}*/
	import_token rune = iota /* denotes \&{import} */
	type_token rune = iota /* \&{type} */
	interface_token rune = iota /* \&{interface} */
	const_token rune = iota /* \&{const} */
	go_token rune = iota /* \&{go} */
	return_token rune = iota /* \&{return} */
	break_token rune = iota /* \&{break} */
	continue_token rune = iota /* \&{continue} */
	goto_token rune = iota /* \&{goto} */
	if_token rune = iota /* \&{if} */
	switch_token rune = iota /* \&{switch} */
	select_token rune = iota /* \&{select} */
	case_token rune = iota /* \&{case} */
	default_token rune = iota /* \&{default} */
	for_token rune = iota /* \&{for}*/
	else_token rune = iota /* \&{else} */
	defer_token rune = iota /* denotes \.{defer} and \.{go} statements*/
	func_token rune = iota /* denotes a function declarator */
	struct_token rune = iota /* \&{struct} */
	var_token rune = iota /* \&{var} */
	range_token rune = iota /* \&{range} */
	map_token rune = iota /* \&{map} */
	chan_token rune = iota /* \&{cnah} */
	dot rune = iota /* \&{.} */
	eq rune = iota /* denotes an assign operator '=' */
	binary_op rune = iota /* "||" | "&&" | rel_op | add_op | mul_op  */
	rel_op rune = iota /* "==" | "!=" | "<" | "<=" | ">" | ">=" */
	add_op rune = iota	/* "+" | "-" | "|" | "^" . */
	mul_op rune = iota /*  "/" | "%" | "<<" | ">>" | "&" | "&^"  */
	unary_op rune = iota /* "+" | "-" | "!" | "^" | "*" | "&" | "<-" */
	asterisk rune = iota /* "*" */
	assign_op rune = iota

	lbrace rune = iota /* denotes a left brace */
	rbrace rune = iota /* denotes a right brace */
	comma rune = iota /* denotes a comma */
	lpar rune = iota /* denotes a left parenthesis */
	rpar rune = iota /* denotes a right parenthesis */
	lbracket rune = iota /* denotes a left bracket */
	rbracket rune = iota /* denotes a right bracket */

	semi rune = iota /* denotes a semicolon */
	colon rune = iota /* denotes a colon */
	insert rune = iota /* a scrap that gets combined with its neighbor */
	section_scrap rune = iota /* section name */
	dead rune = iota /* scrap that won't combine */
)

@ @<Glo...@>=
var cat_name[256]string

@ @<Set in...@>=
for cat_index:=0;cat_index<255;cat_index++ {
	cat_name[cat_index] = "UNKNOWN-" + fmt.Sprintf("%v", cat_index) 
}
@.UNKNOWN@>

cat_name[Type]="Type"
cat_name[ArrayType]="ArrayType"
cat_name[StructType]="StructType"
cat_name[PointerType]="PointerType"
cat_name[InterfaceType]="InterfaceType"
cat_name[SliceType]="SliceType"
cat_name[MapType]="MapType"
cat_name[ChannelType]="ChannelType"
cat_name[FieldDecl]="FieldDecl"
cat_name[AnonymousField]="AnonymousField"
cat_name[Signature]="Signature"
cat_name[Parameters]="Parameters"
cat_name[ParameterList]="ParameterList"
cat_name[ParameterDecl]="ParameterDecl"
cat_name[MethodSpec]="MethodSpec"
cat_name[Block]="Block"
cat_name[Statement]="Statement"
cat_name[ConstDecl]="ConstDecl"
cat_name[TypeDecl]="TypeDecl"
cat_name[VarDecl]="VarDecl"
cat_name[FunctionDecl]="FunctionDecl"
cat_name[MethodDecl]="MethodDecl"
cat_name[ConstSpec]="ConstSpec"
cat_name[IdentifierList]="IdentifierList"
cat_name[ExpressionList]="ExpressionList"
cat_name[TypeSpec]="TypeSpec"
cat_name[VarSpec]="VarSpec"
cat_name[ShortVarDecl]="ShortVarDecl"
cat_name[Receiver]="Receiver"
cat_name[Operand]="Operand"
cat_name[QualifiedIdent]="QualifiedIdent"
cat_name[MethodExpr]="MethodExpr"
cat_name[CompositeLit]="CompositeLit"
cat_name[FunctionLit]="FunctionLit"
cat_name[FunctionType]="FunctionType"
cat_name[LiteralType]="LiteralType"
cat_name[LiteralValue]="LiteralValue"
cat_name[ElementList]="ElementList"
cat_name[Element]="Element"
cat_name[PrimaryExpr]="PrimaryExpr"
cat_name[Conversion]="Conversion"
cat_name[BuiltinCall]="BuiltinCall"
cat_name[Selector]="Selector"
cat_name[Index]="Index"
cat_name[Slice]="Slice"
cat_name[TypeAssertion]="TypeAssertion"
cat_name[Call]="Call"
cat_name[Expression]="Expression"
cat_name[UnaryExpr]="UnaryExpr"
cat_name[ReceiverType]="ReceiverType"
cat_name[LabeledStmt]="LabeledStmt"
cat_name[SimpleStmt]="SimpleStmt"
cat_name[GoStmt]="GoStmt"
cat_name[ReturnStmt]="ReturnStmt"
cat_name[BreakStmt]="BreakStmt"
cat_name[ContinueStmt]="ContinueStmt"
cat_name[GotoStmt]="GotoStmt"
cat_name[fallthrough_token]="fallthrough_token"
cat_name[IfStmt]="IfStmt"
cat_name[SelectStmt]="SelectStmt"
cat_name[ForStmt]="ForStmt"
cat_name[DeferStmt]="DeferStmt"
cat_name[SendStmt]="SendStmt"
cat_name[IncDecStmt]="IncDecStmt"
cat_name[Assignment]="Assignment"
cat_name[ExprSwitchStmt]="ExprSwitchStmt"
cat_name[ExprCaseClause]="ExprCaseClause"
cat_name[TypeSwitchStmt]="TypeSwitchStmt"
cat_name[TypeSwitchGuard]="TypeSwitchGuard"
cat_name[TypeCaseClause]="TypeCaseClause"
cat_name[TypeSwitchCase]="TypeSwitchCase"
cat_name[ForClause]="ForClause"
cat_name[RangeClause]="RangeClause"
cat_name[CommClause]="CommClause"
cat_name[CommCase]="CommCase"
cat_name[RecvStmt]="RecvStmt"
cat_name[BuiltinArgs]="BuiltinArgs"
cat_name[PackageClause]="PackageClause"
cat_name[PackageName]="PackageName"
cat_name[ImportDecl]="ImportDecl"
cat_name[ImportSpec]="ImportSpec"

cat_name[package_token]="package"
cat_name[import_token]="import"
cat_name[type_token]="type"
cat_name[interface_token]="interface"
cat_name[const_token]="const"
cat_name[go_token]="go"
cat_name[return_token]="return"
cat_name[break_token]="break"
cat_name[continue_token]="continue"
cat_name[goto_token]="goto"
cat_name[if_token]="if"
cat_name[switch_token]= "switch"
cat_name[select_token]= "select"
cat_name[case_token]= "case"
cat_name[default_token]= "default"
cat_name[for_token]="for"
cat_name[else_token]="else"
cat_name[defer_token]="defer"
cat_name[func_token]="func"
cat_name[struct_token]="struct"
cat_name[var_token]="var"
cat_name[range_token]="range"
cat_name[map_token]="map"
cat_name[chan_token]="chan"

cat_name[dot]="'.'"

cat_name[eq]="'='"
cat_name[col_eq]="':='"
cat_name[binary_op]="binary_op"
cat_name[rel_op]="rel_op"
cat_name[add_op]="add_op"
cat_name[mul_op]="mul_op"
cat_name[unary_op]="unary_op"
cat_name[asterisk]="'*'"
cat_name[assign_op]="assign_op"

cat_name[lbrace]="'{'"@q}@>
cat_name[rbrace]=@q{@>"'}'"
cat_name[comma]="','"
cat_name[lpar]="'('"
cat_name[rpar]="')'"
cat_name[lbracket]="'['"
cat_name[rbracket]="']'"
cat_name[semi]="';'"
cat_name[colon]="':'"
cat_name[insert]="insert"
cat_name[section_scrap]="section_scrap"
cat_name[dead]="@@d"
cat_name[dot_dot_dot]="'...'"
cat_name[constant]="constant"
cat_name[str]="str"
cat_name[identifier]="identifier"
cat_name[0]="zero"
cat_name[direct]="'<-'"
cat_name[plus_plus]="'++'"
cat_name[minus_minus]="'--'"

@ This code allows \.{GOWEAVE} to display its parsing steps.

@c
/* symbolic printout of a category */
func print_cat(c int32) {
	fmt.Printf("%s(%v)", cat_name[c],c)
}

@ The token lists for translated \TEX/ output contain some special control
symbols as well as ordinary characters. These control symbols are
interpreted by \.{GOWEAVE} before they are written to the output file.

\yskip\hang |break_space| denotes an optional line break or an en space;

\yskip\hang |force| denotes a line break;

\yskip\hang |big_force| denotes a line break with additional vertical space;

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
|big_force|.
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
const (
	math_rel rune = 0244
	big_cancel rune = 0245 /* like |cancel|, also overrides spaces */
	cancel rune = 0246/* overrides |backup|, |break_space|, |force|, |big_force| */
	indent rune = 0247 /* one more tab (\.{\\1}) */
	outdent rune = 0250 /* one less tab (\.{\\2}) */
	opt rune = 0251 /* optional break in mid-statement (\.{\\3}) */
	backup rune = 0252 /* stick out one unit to the left (\.{\\4}) */
	break_space rune = 0253 /* optional break between statements (\.{\\5}) */
	force rune = 0254 /* forced break between statements (\.{\\6}) */
	big_force rune = 0255 /* forced break with additional space (\.{\\7}) */
	quoted_char rune = 0256 /* introduces a character token in the range |0200|--|0377| */
	end_translation rune = 0257 /* special sentinel token at end of list */
	inserted rune = 0260 /* sentinel to mark translations of inserts */
)

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
\.{\\\\} and |int_type| replacing |exp|.

A string of length greater than 20 is broken into pieces of size at most~20
with discretionary breaks in between.

\smallskip
The construction \.{@@t}\thinspace stuff\/\thinspace\.{@@>} contributes
\.{\\hbox\{}\thinspace  stuff\/\thinspace\.\} to the following scrap.

@i prod.w

@* Implementing the productions.
More specifically, a scrap is a structure consisting of a category
|cat| and a |trans|, which contains the translation.
When \GO/ text is to be processed with the grammar above,
we form an array |scrap_info| containing the initial scraps.

@ Here different types of token are defined
@<Type...@>=
type scrap struct {
	cat int32
	mathness int32
	trans []interface{}
	@<Rest of |scrap| struct@>
}

type id_token int 

type res_token int

type section_token int32

type list_token []interface{}

type inner_list_token []interface{}


@ @<Global...@>=
var scrap_info []scrap /* memory array for scraps */

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
//const (
//	id_flag rune = unicode.UpperLower /* signifies an identifier */
//	res_flag rune = 2*id_flag /* signifies a reserved word */
//	section_flag rune = 4*id_flag /* signifies a section name */
//	tok_flag rune = 6*id_flag /* signifies a token list */
//	inner_tok_flag rune = 8*id_flag /* signifies a token list in `\pb' */
//)






@ The production rules listed above are embedded directly into \.{GOWEAVE},
since it is easier to do this than to write an interpretive system
that would handle production systems in general. Several helper functions
are defined here so that the program for each production is fairly short.

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
copies of several existing translations, and few functions are defined to
simplify these common cases. For example, \\{app2}|(pp)| will append the
translations of two consecutive scraps, |scrap_info[pp].trans| 
and |scrap_info[pp+1].trans|, to
the current token list. If the entire new translation is formed in this
way, we write `|squash(j,k,c,d,n)|' instead of `|reduce(j,k,c,d,n)|'. For
example, `|squash(pp,3,exp,-2,3)|' is an abbreviation for `\\{app3}|(pp);
reduce(pp,3,exp,-2,3)|'.

A couple more words of explanation:
Both |big_app| and |app| append a token to the current token list.
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
const (
	maybe_math rune = iota /* works in either horizontal or math mode */
	yes_math rune = iota /* should be in math mode */
	no_math rune = iota /* should be in horizontal mode */
)

@ The function |isCat| checks if the specified index |i| is inside 
the |scrap_info| and a corresponding scrap has the specified category |cat|

@c
func isCat(pp int, cat int32) bool {
	if pp < 0 || pp >=len(scrap_info) {
		if (tracing & 4) == 4 {
			fmt.Fprintf(os.Stdout, "%v; is out of range of the scrap_info\n", pp)
		}
		return false
	}
	if (tracing & 4) == 4 {
		fmt.Fprintf(os.Stdout, "%v; looking for a category %q\n", pp, cat_name[cat])
	}
	if scrap_info[pp].cat==cat {
		if (tracing & 4) == 4 {
			fmt.Fprintf(os.Stdout, "%v; +category %q has been found\n", pp, cat_name[cat])
		}
		return true
	}
	@<Making copy of |scrap_info| and |rollback| function@>
	reduced_cat=-1
	switch cat {
		case ConstDecl: @<Cases for |ConstDecl|@>
		case TypeDecl: @<Cases for |TypeDecl|@>
		case VarDecl: @<Cases for |VarDecl|@>
		case FunctionDecl: @<Cases for |FunctionDecl|@>
		case MethodDecl: @<Cases for |MethodDecl|@>
		case Receiver: @<Cases for |Receiver|@>
		case ConstSpec: @<Cases for |ConstSpec|@> 
		case TypeSpec: @<Cases for |TypeSpec|@>
		case VarSpec: @<Cases for |VarSpec|@>
		case ImportSpec: @<Cases for |ImportSpec|@>
		case FieldDecl: @<Cases for |FieldDecl|@>
		case AnonymousField: @<Cases for |AnonymousField|@>
		case Type: @<Cases for |Type|@>
		case ArrayType: @<Cases for |ArrayType|@>
		case StructType: @<Cases for |StructType|@>
		case PointerType: @<Cases for |PointerType|@>
		case Signature: @<Cases for |Signature|@>
		case Parameters: @<Cases for |Parameters|@>
		case ParameterList: @<Cases for |ParameterList|@>
		case ParameterDecl: @<Cases for |ParameterDecl|@>
		case InterfaceType: @<Cases for |InterfaceType|@>
		case MethodSpec: @<Cases for |MethodSpec|@>
		case SliceType: @<Cases for |SliceType|@>
		case MapType: @<Cases for |MapType|@>
		case ChannelType: @<Cases for |ChannelType|@>
		case IdentifierList: @<Cases for |IdentifierList|@>
		case ExpressionList: @<Cases for |ExpressionList|@>
		case Expression: @<Cases for |Expression|@> 
		case UnaryExpr: @<Cases for |UnaryExpr|@>
		case binary_op: @<Cases for |binary_op|@>
		case PrimaryExpr: @<Cases for |PrimaryExpr|@>
		case Operand: @<Cases for |Operand|@>
		case CompositeLit: @<Cases for |CompositeLit|@>
		case LiteralType: @<Cases for |LiteralType|@>
		case LiteralValue: @<Cases for |LiteralValue|@>
		case ElementList: @<Cases for |ElementList|@>
		case Element: @<Cases for |Element|@>
		case FunctionLit: @<Cases for |FunctionLit|@>
		case FunctionType: @<Cases for |FunctionType|@>
		case Block: @<Cases for |Block|@>
		case Statement: @<Cases for |Statement|@>
		case LabeledStmt: @<Cases for |LabeledStmt|@>
		case SimpleStmt: @<Cases for |SimpleStmt|@>
		case GoStmt: @<Cases for |GoStmt|@>
		case ReturnStmt: @<Cases for |ReturnStmt|@>
		case BreakStmt: @<Cases for |BreakStmt|@>
		case ContinueStmt: @<Cases for |ContinueStmt|@>
		case GotoStmt: @<Cases for |GotoStmt|@>
		case IfStmt: @<Cases for |IfStmt|@>
		case ExprSwitchStmt: @<Cases for |ExprSwitchStmt|@>
		case ExprCaseClause: @<Cases for |ExprCaseClause|@>
		case TypeSwitchStmt: @<Cases for |TypeSwitchStmt|@>
		case TypeSwitchGuard: @<Cases for |TypeSwitchGuard|@>
		case TypeCaseClause: @<Cases for |TypeCaseClause|@>
		case TypeSwitchCase: @<Cases for |TypeSwitchCase|@>
		case SelectStmt: @<Cases for |SelectStmt|@>
		case CommClause: @<Cases for |CommClause|@>
		case CommCase: @<Cases for |CommCase|@>
		case RecvStmt: @<Cases for |RecvStmt|@>
		case SendStmt: @<Cases for |SendStmt|@>
		case ForStmt: @<Cases for |ForStmt|@>
		case ForClause: @<Cases for |ForClause|@>
		case RangeClause: @<Cases for |RangeClause|@>
		case DeferStmt: @<Cases for |DeferStmt|@>
		case IncDecStmt: @<Cases for |IncDecStmt|@>
		case Assignment: @<Cases for |Assignment|@>
		case assign_op: @<Cases for |assign_op|@> 
		case ShortVarDecl: @<Cases for |ShortVarDecl|@>
		case QualifiedIdent: @<Cases for |QualifiedIdent|@>
		case MethodExpr: @<Cases for |MethodExpr|@>
		case ReceiverType: @<Cases for |ReceiverType|@>
		case Conversion: @<Cases for |Conversion|@>
		case BuiltinCall: @<Cases for |BuiltinCall|@>
		case BuiltinArgs: @<Cases for |BuiltinArgs|@>
		case Selector: @<Cases for |Selector|@> 
		case Index: @<Cases for |Index|@>
		case Slice: @<Cases for |Slice|@>
		case TypeAssertion: @<Cases for |TypeAssertion|@>
		case Call: @<Cases for |Call|@>
		case unary_op: @<Cases for |unary_op|@> 
		default:
			if (tracing & 4) == 4 {
				fmt.Fprintf(os.Stdout, "%v; -category %q hasn't been found\n", pp, cat_name[cat])
			}
			rollback()
			return false
	}
	if reduced_cat==cat {
		if (tracing & 4) == 4 {
			fmt.Fprintf(os.Stdout, "%v; +category %q has been found\n", pp, cat_name[cat])
		}
	} else { 
		if (tracing & 4) == 4 {
			fmt.Fprintf(os.Stdout, "%v; -category %q hasn't been found\n", pp, cat_name[cat])
		}
		rollback()
	}
	return reduced_cat==cat
}

@ The function |isCats| checks if the specified index |pp| is inside 
the |scrap_info| and a corresponding scraps have the specified sequence of categories |cats|.
Some of the catigories |cats| can be optional.

@<Typedef declarations@>=
type cat_pair struct {
	cat int32
	mand bool
}

@ @c
func isCats(pp int, c *int, cats ...cat_pair) bool {
	*c=0
	res:=false
	exit:=false
	for !exit && pp<len(scrap_info) {
		r:=false
		for _,v:=range cats {
			if isCat(pp,v.cat) {
				r=true
				*c++
				pp++
			} else if v.mand {
				exit=true
				break
			}
		}
		if !res {
			res=r
		}
		if !r {
			exit=true
		}
 	}
	return res
}


@ The function |isNotCat| checks if the specified index |i| is outside 
the |scrap_info| or a corresponding scrap hasn't the specified category |cat|


@c
func isNotCat(i int, cat int32) bool {
	if i < 0 || i >=len(scrap_info) {
		return false
	}
	return scrap_info[i].cat != cat
}

@ @<Making copy of |scrap_info| and |rollback| function@>=
scraps_copy:=append([]scrap{},scrap_info[pp:]...)
reduced_copy:=reduced
reduced=false
rollback:=func(){
	if reduced {
		n:=pp
		scrap_info=scrap_info[:pp]
		scrap_info=append(scrap_info,scraps_copy...)
		f := "rollback"
		@<Print a snapshot of the scrap list if debugging@>
	}
	reduced=reduced_copy
}


@ Let us consider the big switch for productions now, before looking
at its context. We want to design the program so that this switch
works, so we might as well not keep ourselves in suspense about exactly what
code needs to be provided with a proper environment.

@ @<Match a production at |pp|, or increase |pp| if there is no match@>= {
	/* not a production with left side length 1 */	
	if isCat(pp+1,insert) { 
		reduce(pp,2,scrap_info[pp].cat,-2,0,pp,pp+1)
		pp--
	} else if isCat(pp+2,insert) { 
		reduce(pp+1,2,scrap_info[pp+1].cat,-1,0,pp+1,pp+2)
		pp--
	} else if isCat(pp+3,insert) { 
		reduce(pp+2,2,scrap_info[pp+2].cat,0,0,pp+2,pp+3)
		pp--
	} else {
		switch scrap_info[pp].cat {
			case insert: 
				@<Cases for |insert|@>
			case package_token:
				@<Cases for |PackageClause|@>
			case import_token:
				@<Cases for |ImportDecl|@>
			case struct_token:
				@<Cases for |StructType|@>
			case interface_token:
				@<Cases for |InterfaceType|@>
			case func_token:
				@<Cases for |FunctionDecl|@> 
				if reduced_cat == FunctionDecl {
					break
				}
				@<Cases for |MethodDecl|@>
				if reduced_cat == MethodDecl {
					break
				}
				@<Cases for |FunctionType|@>
			default:
				@<Cases for |Statement|@>	
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

@ @c
func find_first_ident(p []interface{}) []interface{} {
	for i, j:= range p {
		switch r := j.(type) {
			case res_token: /* |res_flag| */
				if name_dir[r].ilk==case_token {
					return nil
				}
				if name_dir[r].ilk!=Type {
					break
				}
				return p[i:i+1]
			case id_token: 
				return p[i:i+1]
			case list_token, inner_list_token: /* |tok_flag| or |inner_tok_flag| */
				if q:=find_first_ident(r.([]interface{})); q!=nil {
					return q
				}
			case rune:  /* char, |section_token|, fallthru: move on to next token */
				if r==inserted {
					return nil /* ignore inserts */
				}
		}
	}
	return nil
}

@ The scraps currently being parsed must be inspected for any
occurrence of the identifier that we're making reserved; hence
the |for| loop below.

@c
/* make the first identifier in |scrap_info[p].trans| like |c| */
func make_reserved(p int, c rune) {
	tok_ptr:=find_first_ident(scrap_info[p].trans)
	if tok_ptr==nil {
		return /* this should not happen */
	}
	name_dir[tok_ptr[0].(id_token)].ilk=c
	tok_ptr[0]=res_token(tok_ptr[0].(id_token))
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
/* underline the entry for the first identifier in |scrap_info[p].trans| */
func make_underlined(p int) {
	tok_ptr:=find_first_ident(scrap_info[p].trans)
	if tok_ptr==nil {
		return /* this happens, for example, in |case found:| */
	}
	xref_switch=def_flag
	underline_xref(tok_ptr[0].(id_token))
}

@ We cannot use |new_xref| to underline a cross-reference at this point
because this would just make a new cross-reference at the end of the list.
We actually have to search through the list for the existing
cross-reference.

@ @c
func underline_xref(p id_token) {
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
the |squash| or |reduce| funcs will cause the appropriate action
to be performed, followed by |goto found|.

@ @<Cases for |insert|@>=
if isNotCat(pp+1,zero) {
	reduce(pp,2,scrap_info[pp+1].cat,0,0,pp,pp+1)
}

@ @<Cases for |PackageClause|@>=
if isCat(pp,package_token)  && isCat(pp+1,identifier) {
	make_reserved(pp+1,PackageName)
	reduce(pp,2,PackageClause,1,1,pp,break_space,pp+1,big_force)
}

@ Test for |package|
@(tests/package.w@>=
@@
@@2
@@c
package main

@ @<Cases for |ConstDecl|@>= 
if isCat(pp,const_token) {
	if isCat(pp+1,ConstSpec) {
		reduce(pp,2,ConstDecl,0,2,pp,break_space,pp+1,big_force)
	} else if rollback(); isCat(pp+1,lpar) {
		c:=0
		isCats(pp+2,&c,cat_pair{cat:ConstSpec,mand:true},cat_pair{cat:semi,mand:false})	
		if isCat(pp+2+c,rpar) {
			tok_mem:=append([]interface{}{},pp,pp+1)
			for i:=0;i<c;i++ {
				if i==0 {
					tok_mem=append(tok_mem,force,indent)
				}
				if isCat(pp+2+i,ConstSpec) {
					tok_mem=append(tok_mem,pp+2+i,force)
				}
				if i==c-1 {
					tok_mem=append(tok_mem,outdent)
				}
			}
			tok_mem=append(tok_mem,pp+2+c,big_force)
			reduce(pp,3+c,ConstDecl,0,2,tok_mem...)
		}	
	}
}

@ Tests for |const|
@(tests/const.w@>=
@@
@@2
@@c
const Pi float64 = 3.14159265358979323846
@@
@@c
const zero = 0.0 
@@
@@c
const (
	size int64 = 1024
	eof        = -1
)
@@
@@c
const a, b, c = 3, 4, "foo"
@@
@@c
const u, v float32 = 0, 3

@ @<Cases for |TypeDecl|@>= 
if isCat(pp,type_token) {
	if isCat(pp+1,TypeSpec) {
		reduce(pp,2,TypeDecl,0,3,pp,break_space,pp+1,big_force)
	} else if rollback(); isCat(pp+1,lpar) {
		c:=0
		isCats(pp+2,&c,cat_pair{cat:TypeSpec,mand:true},cat_pair{cat:semi,mand:false})
		if isCat(pp+2+c,rpar) {
			tok_mem:=append([]interface{}{},pp,pp+1)
			for i:=0;i<c;i++ {
				if i==0 {
					tok_mem=append(tok_mem,force,indent)
				}
				if isCat(pp+2+i,TypeSpec) {
					tok_mem=append(tok_mem,pp+2+i,force)
				}
				if i==c-1 {
					tok_mem=append(tok_mem,outdent)
				}
			}
			tok_mem=append(tok_mem,pp+2+c,big_force)
			reduce(pp,3+c,TypeDecl,0,3,tok_mem...)
		}
	} 
}

@ Tests for |type|
@(tests/type.w@>=
@@
@@2
@@c
type IntArray [16]int
@@
@@c
type (
	Point struct{ x, y float64 }
	Polar Point
)
@@
@@c
type TreeNode struct {
	left, right *TreeNode
	value *Comparable
}
@@
@@c
type Block interface {
	BlockSize() int
	Encrypt(src, dst []byte)
	Decrypt(src, dst []byte)
}

@ @<Cases for |VarDecl|@>=
if isCat(pp,var_token) {
	if isCat(pp+1,VarSpec) {
		reduce(pp,2,VarDecl,0,4,pp,break_space,pp+1,big_force)
	} else if rollback(); isCat(pp+1,lpar) {
		c:=0
		isCats(pp+2,&c,cat_pair{cat:VarSpec,mand:true},cat_pair{cat:semi,mand:false}) 
		if isCat(pp+2+c,rpar) {
			tok_mem:=append([]interface{}{},pp,pp+1)
			for i:=0;i<c;i++ {
				if i==0 {
					tok_mem=append(tok_mem,force,indent)
				}
				if isCat(pp+2+i,VarSpec) {
					tok_mem=append(tok_mem,pp+2+i,force)
				}
				if i==c-1 {
					tok_mem=append(tok_mem,outdent)
				}
			}
			tok_mem=append(tok_mem,pp+2+c,big_force)
			reduce(pp,3+c,VarDecl,0,4,tok_mem...)
		}
	} 
}

@ Tests for |var|
@(tests/var.w@>=
@@
@@2
@@c
var i int
@@
@@c
var U, V, W float64
@@
@@c
var k = 0
@@
@@c
var x, y float32 = -1, -2
@@
@@c
var (
	i       int
	u, v, s = 2.0, 3.0, "bar"
)
@@
@@c
var re, im = complexSqrt(-1)
@@
@@c
var _, found = entries[name]

@ @<Cases for |ImportDecl|@>=
if isCat(pp,import_token) {
	@<Making copy...@>
	if isCat(pp+1,ImportSpec) {
		reduce(pp,2,ImportDecl,0,5,pp,break_space,pp+1,big_force)
	} else if rollback(); isCat(pp+1,lpar) {
		c:=0
		isCats(pp+2,&c,cat_pair{cat:ImportSpec,mand:true},cat_pair{cat:semi,mand:false})
		if isCat(pp+2+c,rpar) {
			tok_mem:=append([]interface{}{},pp,pp+1)
			for i:=0;i<c;i++ {
				if i==0 {
					tok_mem=append(tok_mem,force,indent)
				}
				if isCat(pp+2+i,ImportSpec) {
					tok_mem=append(tok_mem,pp+2+i,force)
				}
				if i==c-1 {
					tok_mem=append(tok_mem,outdent)
				}
			}
			tok_mem=append(tok_mem,pp+2+c,big_force)
			reduce(pp,3+c,ImportDecl,0,5,tok_mem...)	
		} 
	}
}

@ Tests for |import|
@(tests/import.w@>=
@@
@@2
@@c
import "im1" 
@@
@@c
import _ "im2"; /*im2*/
@@
@@c
import . "im3" //im3
@@
@@c
import IM "im4"
@@
@@c
import(
	"nim1" 
	. "nim2"; // nim2
	_ "nim3" /*nim3*/
	NIM "nim4"
)


@ @<Cases for |FunctionDecl|@>=
if isCat(pp,func_token) && isCat(pp+1,identifier) && isCat(pp+2,Signature){
	pp+=3
	@<Making copy...@>
	pp-=3
	if isCat(pp+3,Block) {
		reduce(pp,4,FunctionDecl,0,6,pp,break_space,pp+1,pp+2,pp+3)	
	} else {
		rollback()
		reduce(pp,3,FunctionDecl,0,6,pp,break_space,pp+1,pp+2)	
	}
}

@ Tests for |func|
@(tests/func.w@>=
@@
@@2
@@c
func min(x int, y int) int {
        if x < y {
                return x
        }
        return y
}
@@
@@c
func flushICache(begin, end uintptr)

@ @<Cases for |MethodDecl|@>=
if isCat(pp,func_token) && isCat(pp+1,Receiver) && isCat(pp+2,identifier) && isCat(pp+3,Signature) {
	pp+=3
	@<Making copy...@>
	pp-=3
	if isCat(pp+4,Block) {
		reduce(pp,5,MethodDecl,0,7,pp,break_space,pp+1,break_space,pp+2,pp+3,pp+4,force)
	} else {
		rollback()
		reduce(pp,4,MethodDecl,0,7,pp,break_space,pp+1,break_space,pp+2,pp+3)
	}
}

@ Tests for |method|
@(tests/method.w@>=
@@
@@2
@@c
func (p *Point) Length() float64 {
	return math.Sqrt(p.x * p.x + p.y * p.y)
}
@@
@@c
func (p *Point) Scale(factor float64) {
	p.x *= factor
	p.y *= factor
}


@ @<Cases for |Receiver|@>=
if isCat(pp,lpar) {
	if isCat(pp+1,identifier) {
		if isCat(pp+2,asterisk) && isCat(pp+3,identifier) && isCat(pp+4,rpar){
			reduce(pp,5,Receiver,0,8,pp,pp+1,pp+2,pp+3,pp+4)
		} else if rollback(); isCat(pp+2,identifier) && isCat(pp+3,rpar) {
			reduce(pp,4,Receiver,0,8,pp,pp+1,pp+2,pp+3)
		} else if rollback(); isCat(pp+2,rpar) {
			reduce(pp,3,Receiver,0,8,pp,pp+1,pp+2)
		}
	} else if rollback(); isCat(pp+1,asterisk) && isCat(pp+2,identifier) && isCat(pp+3,rpar) {
		reduce(pp,4,Receiver,0,8,pp,pp+1,pp+2,pp+3)
	}
}

@ @<Cases for |ConstSpec|@>= 
if isCat(pp,IdentifierList) {
	pp++
	@<Making copy...@>
	pp--
	if isCat(pp+1,Type) && isCat(pp+2,eq) && isCat(pp+3,ExpressionList) {
		reduce(pp,4,ConstSpec,0,9,pp,break_space,pp+1,break_space,pp+2,break_space,pp+3)
	} else if rollback(); isCat(pp+1,eq) && isCat(pp+2,ExpressionList) {
		reduce(pp,3,ConstSpec,0,9,pp,break_space,pp+1,break_space,pp+2)
	}
} else if rollback(); isCat(pp, section_scrap) {
	reduce(pp,1,ConstSpec,0,9,pp)
}

@ @<Cases for |TypeSpec|@>=
if isCat(pp,identifier) && isCat(pp+1,Type) {
	reduce(pp,2,TypeSpec,0,10,pp,break_space,pp+1)
} else if rollback(); isCat(pp, section_scrap) {
	reduce(pp,1,TypeSpec,0,10,pp)
}

@ @<Cases for |VarSpec|@>=
if isCat(pp,IdentifierList) {
	pp++
	@<Making copy...@>
	pp--
	if isCat(pp+1,Type) {
		if isCat(pp+2,eq) && isCat(pp+3,ExpressionList) {
			reduce(pp,4,VarSpec,0,11,pp,break_space,pp+1,pp+2,pp+3)
		} else {
			reduce(pp,2,VarSpec,0,11,pp,break_space,pp+1)
		}
	} else if rollback(); isCat(pp+1,eq) && isCat(pp+2,ExpressionList) {
		reduce(pp,3,VarSpec,0,11,pp,pp+1,pp+2)
	}
} else if rollback(); isCat(pp,section_scrap) {
	reduce(pp,1,VarSpec,0,11,pp)	
}

@ @<Cases for |ImportSpec|@>=
if isCat(pp,identifier) && isCat(pp+1,str) {
	c:=0
	isCats(pp+2,&c,cat_pair{cat:semi,mand:false})
	make_reserved(pp,PackageName)
	reduce(pp,2+c,ImportSpec,0,12,pp,break_space,pp+1)
} else if isCat(pp,dot) && isCat(pp+1,str) {
	c:=0
	isCats(pp+2,&c,cat_pair{cat:semi,mand:false})
	reduce(pp,2+c,ImportSpec,0,12,pp,break_space,pp+1)
} else if isCat(pp,str) {
	c:=0
	isCats(pp+1,&c,cat_pair{cat:semi,mand:false})
	reduce(pp,1+c,ImportSpec,0,12,pp)
} else if isCat(pp,section_scrap) {
	reduce(pp,1,ImportSpec,0,12,pp)
}

@ @<Cases for |FieldDecl|@>=
if isCat(pp,IdentifierList) && isCat(pp+1,Type) {
	tok_mem:=append([]interface{}{},pp,break_space,pp+1)
	p:=pp+2
	if isCat(p,str) {
		tok_mem=append(tok_mem,break_space,pp+2)
		p++
	}
	reduce(pp,p-pp,FieldDecl,0,13,tok_mem...)
} else if rollback(); isCat(pp,AnonymousField) {
	tok_mem:=append([]interface{}{},pp)
	p:=pp+1
	if isCat(p,str) {
		tok_mem=append(tok_mem,pp,break_space,pp+1,break_space,pp+1)
		p++
	}
	reduce(pp,p-pp,FieldDecl,0,13,tok_mem...)
} else if rollback(); isCat(pp,section_scrap) {
	reduce(pp,1,FieldDecl,0,13,pp)
}

@ @<Cases for |AnonymousField|@>=
if isCat(pp,asterisk) && isCat(pp+1,Type) {
	reduce(pp,2,AnonymousField,0,14,pp,pp+1)
} else if rollback(); isCat(pp,Type) {
	reduce(pp,1,AnonymousField,0,14,pp)
}

@ @<Cases for |Type|@>=
if  isCat(pp,ArrayType) || isCat(pp,StructType) || isCat(pp,PointerType) || 
	isCat(pp,FunctionType) || isCat(pp,InterfaceType) || isCat(pp,SliceType) || 
	isCat(pp,MapType) || isCat(pp,ChannelType) || isCat(pp,QualifiedIdent) {
	reduce(pp,1,Type,0,15,pp)
}

@ @<Cases for |ArrayType|@>=
if isCat(pp,lbracket) && isCat(pp+1,Expression) && isCat(pp+2,rbracket) && isCat(pp+3,Type) {
	reduce(pp,4,ArrayType,0,16,pp,pp+1,pp+2,pp+3)
}

@ @<Cases for |StructType|@>=
if isCat(pp,struct_token) && isCat(pp+1,lbrace) {
	c:=0
	isCats(pp+2,&c,cat_pair{cat:FieldDecl,mand:true},cat_pair{cat:semi,mand:false})
	if isCat(pp+2+c,rbrace) {
		tok_mem:=append([]interface{}{},pp,pp+1)
		for i:=0;i<c;i++ {
			if i==0 {
				tok_mem=append(tok_mem,force,indent)
			}
			if isCat(pp+2+i,FieldDecl) {
				tok_mem=append(tok_mem,pp+2+i,force)
			}
		}
		tok_mem=append(tok_mem,outdent,pp+2+c)
		reduce(pp,3+c,StructType,0,17,tok_mem...)
	}
}

@ Tests for |struct|
@(tests/struct.w@>=
@@
@@2
@@c
struct {}
@@
@@c
struct {
	x, y int
	u float32
	_ float32
	A *[]int
	F func()
}
@@
@@c
struct {
	T1
	*T2
	P.T3
	*P.T4
	x, y int
}
@@
@@c
struct {
	microsec  uint64 "field 1"
	serverIP6 uint64 "field 2"
	process   string "field 3"
}

@ @<Cases for |PointerType|@>=
if isCat(pp,asterisk) && isCat(pp+1,Type) {
	reduce(pp,2,PointerType,0,18,pp,pp+1)
}

@ @<Cases for |Signature|@>=
if isCat(pp,Parameters) {
	pp++
	@<Making copy...@>
	pp--
	if isCat(pp+1,Type) || isCat(pp+1,Parameters) {
		reduce(pp,2,Signature,0,19,pp,break_space,pp+1)
	} else {
		rollback()
		reduce(pp,1,Signature,0,19,pp)
	}
} else if rollback(); isCat(pp,section_scrap) {
	reduce(pp,1,Signature,0,19,pp)
}

@ @<Cases for |Parameters|@>=
if isCat(pp,lpar) {
	c:=0
	isCats(pp+1,&c,cat_pair{cat:ParameterList,mand:true},cat_pair{cat:comma,mand:false})
 	if isCat(pp+1+c,rpar) {
		tok_mem:=append([]interface{}{},pp)
		for i:=0;i<c;i++ {
			tok_mem=append(tok_mem,pp+1+i)
		}
		tok_mem=append(tok_mem,pp+1+c)
		reduce(pp,2+c,Parameters,0,20,tok_mem...)
	}
} else if rollback(); isCat(pp,section_scrap) {
	reduce(pp,1,Signature,0,20,pp)
}

@ @<Cases for |ParameterList|@>=
if isCat(pp,ParameterDecl) {
	c:=0
	isCats(pp+1,&c,cat_pair{cat:comma,mand:true},cat_pair{cat:ParameterDecl,mand:true})
	tok_mem:=append([]interface{}{},pp)	
	for i:=0;i<c;i++ {
		tok_mem=append(tok_mem,pp+1+i)
	}
	reduce(pp,1+c,ParameterList,0,21,tok_mem...)
}

@ @<Cases for |ParameterDecl|@>=
if isCat(pp,IdentifierList) && isCat(pp+1,dot_dot_dot) &&  isCat(pp+2,Type) {
	reduce(pp,3,ParameterDecl,0,22,pp,break_space,pp+1,pp+2)
} else if rollback(); isCat(pp,IdentifierList) && isCat(pp+1,Type) {
	reduce(pp,2,ParameterDecl,0,22,pp,break_space,pp+1)
} else if rollback(); isCat(pp,dot_dot_dot) &&  isCat(pp+1,Type) {
	reduce(pp,2,ParameterDecl,0,22,pp,pp+1)
} else if rollback(); isCat(pp,Type) {
	reduce(pp,1,ParameterDecl,0,22,pp)
}

break
p:=pp
var tok_mem []interface{}
if isCat(pp,IdentifierList) {
	tok_mem=append(tok_mem,pp,break_space)
	pp++
} else {
	rollback()
}
if isCat(pp,dot_dot_dot) {
	tok_mem=append(tok_mem,pp)
	pp++
}
if isCat(pp,Type) {
	tok_mem=append(tok_mem,pp)
	pp+=1
	reduce(p,pp-p,ParameterDecl,0,22,tok_mem...)
}
pp=p

@ @<Cases for |InterfaceType|@>=
if isCat(pp,interface_token) && isCat(pp+1,lbrace) {
	c:=0
	isCats(pp+2,&c,cat_pair{cat:MethodSpec,mand:true},cat_pair{cat:semi,mand:false})
	if isCat(pp+2+c,rbrace) {
		tok_mem:=append([]interface{}{},pp,pp+1,force,indent)
		for i:=0;i<c;i++ {
			if isCat(pp+2+i,MethodSpec) {
				tok_mem=append(tok_mem,pp+2+i,force)
			}
		}
		tok_mem=append(tok_mem,outdent,pp+2+c)
		reduce(pp,3+c,InterfaceType,0,23,tok_mem...)
	}
}

@ @<Cases for |MethodSpec|@>=
if isCat(pp,identifier) && isCat(pp+1,Signature) {
	reduce(pp,2,MethodSpec,0,24,pp,pp+1)
} else if rollback(); isCat(pp,Type) {
	reduce(pp,1,MethodSpec,0,24,pp)	
} else if rollback(); isCat(pp,section_scrap) {
	reduce(pp,1,MethodSpec,0,24,pp)	
}

@ @<Cases for |SliceType|@>=
if isCat(pp,lbracket) && isCat(pp+1,rbracket) && isCat(pp+2,Type) {
	reduce(pp,3,SliceType,0,25,pp,pp+1,pp+2)
}

@ @<Cases for |MapType|@>=
if isCat(pp,map_token) && isCat(pp+1,lbracket) && isCat(pp+2,Type) && isCat(pp+3,rbracket) && isCat(pp+4,Type) {
	reduce(pp,5,MapType,0,26,pp,pp+1,pp+2,pp+3,pp+4)
}

@ @<Cases for |ChannelType|@>=
if isCat(pp,direct) && isCat(pp+1,chan_token) && isCat(pp+2,Type) {
	reduce(pp,3,ChannelType,0,27,pp,pp+1,break_space,pp+2)
} else if rollback(); isCat(pp,chan_token) { 
	if isCat(pp+1,direct) && isCat(pp+2,Type) {
		reduce(pp,3,ChannelType,0,27,pp,pp+1,pp+2)
	} else if isCat(pp+1,Type) {
		reduce(pp,2,ChannelType,0,27,pp,break_space,pp+1)
	}
}

@ @<Cases for |IdentifierList|@>=
if isCat(pp,identifier) {
	c:=0
	isCats(pp+1,&c,cat_pair{cat:comma,mand:true},cat_pair{cat:identifier,mand:true}) 
	tok_mem:=append([]interface{}{},pp)
	for i:=0;i<c;i++ {
		tok_mem=append(tok_mem,pp+1+i)
	}
	reduce(pp,1+c,IdentifierList,0,28,tok_mem...)
} 

@ @<Cases for |ExpressionList|@>=
if isCat(pp,Expression) {
	c:=0
	isCats(pp+1,&c,cat_pair{cat:comma,mand:true},cat_pair{cat:Expression,mand:true})	
	tok_mem:=append([]interface{}{},pp)
	for i:=0;i<c;i++ {
		tok_mem=append(tok_mem,pp+1+i)
	}
	reduce(pp,1+c,ExpressionList,0,29,tok_mem...)
}

@ @<Cases for |Expression|@>= 
if isCat(pp,UnaryExpr) && isCat(pp+1,binary_op) && isCat(pp+2,UnaryExpr) {
	reduce(pp,3,Expression,0,30,pp,pp+1,pp+2)
} else if rollback(); isCat(pp,UnaryExpr) {
	reduce(pp,1,Expression,0,30,pp)
}

@ @<Cases for |UnaryExpr|@>=
if isCat(pp,unary_op) && isCat(pp+1,UnaryExpr) {
	reduce(pp,2,UnaryExpr,0,31,pp,pp+1)
} else if isCat(pp,PrimaryExpr) {
	reduce(pp,1,UnaryExpr,0,31,pp)
}

@ @<Cases for |binary_op|@>=
if isCat(pp,rel_op) || isCat(pp,add_op) || isCat(pp,mul_op) || isCat(pp,asterisk) {
	reduce(pp,1,binary_op,0,32,pp)
}

@ @<Cases for |PrimaryExpr|@>=
if isCat(pp,BuiltinCall) || isCat(pp,Conversion) || isCat(pp,Operand) {
	pp++
	@<Making copy...@>
	pp--
	if isCat(pp+1,Selector) || isCat(pp+1,Index) || isCat(pp+1,Slice) || isCat(pp+1,TypeAssertion) || isCat(pp+1,Call) {
		reduce(pp,2,PrimaryExpr,0,33,pp,pp+1)
	} else {
		rollback()
		reduce(pp,1,PrimaryExpr,0,33,pp)
	}
}

@ @<Cases for |Operand|@>=
if isCat(pp,str) || isCat(pp,constant) || isCat(pp,QualifiedIdent) || isCat(pp,CompositeLit) || isCat(pp,FunctionLit)  || isCat(pp,MethodExpr) {
	reduce(pp,1,Operand,0,34,pp)
} else if rollback(); isCat(pp,lpar) && isCat(pp+1,Expression) && isCat(pp+2,rpar) {
	reduce(pp,3,Operand,0,34,pp,pp+1,pp+2)
}

@ @<Cases for |CompositeLit|@>=
if isCat(pp,LiteralType) && isCat(pp+1,LiteralValue) {
	reduce(pp,2,CompositeLit,0,35,pp,break_space,pp+1)
}

@ @<Cases for |LiteralType|@>=
if isCat(pp,Type) {
	reduce(pp,1,LiteralType,0,36,pp)
} else if rollback(); isCat(pp,lbracket) && isCat(pp+1,dot_dot_dot) && isCat(pp+2,rbracket) && isCat(pp+3,Type) {
	reduce(pp,4,LiteralType,0,36,pp,pp+1,pp+2,pp+3)
}

@ @<Cases for |LiteralValue|@>=
if isCat(pp,lbrace) {
	c:=0
	isCats(pp+1,&c,cat_pair{cat:ElementList,mand:true},cat_pair{cat:comma,mand:true})
	if isCat(pp+1+c,rbrace) {
		tok_mem:=append([]interface{}{},pp)
		for i:=0;i<c;i++ {
			tok_mem=append(tok_mem,pp+1+i)
		}
		tok_mem=append(tok_mem,pp+1+c)
		reduce(pp,2+c,LiteralValue,0,37,tok_mem...)
	}
}

@ @<Cases for |ElementList|@>=
if isCat(pp,Element) {
	c:=0
	isCats(pp+1,&c,cat_pair{cat:comma,mand:true},cat_pair{cat:Element,mand:true})
	tok_mem:=append([]interface{}{},pp)
	for i:=0;i<c;i++ {
		tok_mem=append(tok_mem,pp+1+i)
	}
	reduce(pp,1+c,ElementList,0,38,tok_mem...)
}

@ @<Cases for |Element|@>=
if (isCat(pp,identifier) || isCat(pp,Expression)) && isCat(pp+1,colon) {
	pp+=2
	@<Making copy...@>
	pp-=2
	if isCat(pp+2,Expression) {
		reduce(pp,3,Element,0,39,pp,pp+1,break_space,pp+2)
	} else if rollback(); isCat(pp+2,LiteralValue) {
		reduce(pp,3,Element,0,39,pp,pp+1,break_space,pp+2)
	}
} else if isCat(pp,Expression) || isCat(pp,LiteralValue) {
	reduce(pp,1,Element,0,39,pp)
}

@ @<Cases for |FunctionLit|@>=
if isCat(pp,FunctionType) && isCat(pp+1,Block) {
	reduce(pp,2,FunctionLit,0,40,pp,pp+1)
}

@ @<Cases for |FunctionType|@>=
if isCat(pp,func_token) && isCat(pp+1,Signature) {
	reduce(pp,2,FunctionType,0,42,pp,pp+1)
}

@ @<Cases for |Block|@>=
if isCat(pp,lbrace) {
	c:=0
	isCats(pp+1,&c,cat_pair{cat:Statement,mand:true},cat_pair{cat:semi,mand:false})
	if isCat(pp+1+c,rbrace) {
		tok_mem:=append([]interface{}{},pp,big_force,indent)
		for i:=0;i<c;i++ {
			if isCat(pp+1+i,Statement) {
				tok_mem=append(tok_mem,pp+1+i,force)
			}
		}
		tok_mem=append(tok_mem,outdent,pp+1+c)
		reduce(pp,2+c,Block,0,43,tok_mem...)
	}
}

@ Tests for |block|
@(tests/block.w@>=
@@
@@2
@@c
{
	a:=b
}

@ @<Cases for |Statement|@>=
if isCat(pp,ConstDecl) || isCat(pp,VarDecl) || isCat(pp,TypeDecl) ||
	isCat(pp,LabeledStmt) || isCat(pp,SimpleStmt) ||
	isCat(pp,GoStmt) || isCat(pp,ReturnStmt) || isCat(pp,BreakStmt) || isCat(pp,ContinueStmt) || 
	isCat(pp,GotoStmt) || isCat(pp,fallthrough_token) || isCat(pp,Block) || isCat(pp,IfStmt) || 
	isCat(pp,ExprSwitchStmt) || isCat(pp,TypeSwitchStmt) || isCat(pp,SelectStmt) || 
	isCat(pp,ForStmt) || isCat(pp,DeferStmt) {
	reduce(pp,1,Statement,0,44,pp)
}

@ @<Cases for |LabeledStmt|@>=
if isCat(pp,identifier) && isCat(pp+1,colon) && isCat(pp+2,Statement) {
	reduce(pp,3,LabeledStmt,0,45,pp,pp+1,break_space,pp+2)
}


@ Tests for |label|
@(tests/label.w@>=
@@
@@2
@@c
Error: log.Panic("error encountered")

@ @<Cases for |SimpleStmt|@>=
if isCat(pp,SendStmt) || isCat(pp,IncDecStmt) || isCat(pp,Assignment) || isCat(pp,ShortVarDecl) || isCat(pp,Expression) {
	reduce(pp,1,SimpleStmt,0,46,pp)
} 

@ @<Cases for |GoStmt|@>=
if isCat(pp,go_token) && isCat(pp+1,Expression) {
	reduce(pp,2,GoStmt,0,47,pp,break_space,pp+1)
}

@ Tests for |go|
@(tests/go.w@>=
@@
@@2
@@c
go Server()
@@
@@c
go func(ch chan<- bool) { for { sleep(10); ch <- true; }} (c)


@ @<Cases for |ReturnStmt|@>=
if isCat(pp,return_token) && isCat(pp+1,ExpressionList) {
	reduce(pp,2,ReturnStmt,0,48,pp,break_space,pp+1)
} else if rollback();  isCat(pp,return_token) {
	reduce(pp,1,ReturnStmt,0,48,pp)
}

@ Tests for |return|
@(tests/return.w@>=
@@
@@2
@@c
return
@@
@@c
return -7.0, -4.0
@@
@@c
return complexF1()


@ @<Cases for |BreakStmt|@>=
if isCat(pp,break_token) {
	if isCat(pp+1,identifier) {
		reduce(pp,2,BreakStmt,0,49,pp,break_space,pp+1)
	} else {
		reduce(pp,1,BreakStmt,0,49,pp)
	}
}

@ Tests for |break|
@(tests/break.w@>=
@@
@@2
@@c
for i < n {
	switch i {
	case 5:
	break
	}
}@@
@@c
L:
for i < n {
	switch i {
	case 5:
	break L
	}
}


@ @<Cases for |ContinueStmt|@>=
if isCat(pp,continue_token) && isCat(pp+1,identifier) {
	reduce(pp,2,ContinueStmt,0,50,pp,break_space,pp+1)
} else if rollback(); isCat(pp,continue_token) {
	reduce(pp,1,ContinueStmt,0,50,pp)
}

@ Tests for |continue|
@(tests/continue.w@>=
@@
@@2
@@c
for i < n {
	switch i {
	case 5:
	continue
	}
}
@@
@@c
L:
for i < n {
	switch i {
	case 5:
	continue L
	}
}

@ @<Cases for |GotoStmt|@>=
if isCat(pp,goto_token) && isCat(pp+1,identifier) {
	reduce(pp,2,GotoStmt,0,51,pp,break_space,pp+1)
}

@ Tests for |goto|
@(tests/goto.w@>=
@@
@@2
@@c
goto Label

@ @<Cases for |IfStmt|@>=
p:=pp
if isCat(pp,if_token) {
	tok_mem:=append([]interface{}{},pp,break_space)
	pp++
	@<Making copy...@>
	if isCat(pp,SimpleStmt) && isCat(pp+1,semi) {
		tok_mem=append(tok_mem,pp,pp+1,break_space)
		pp+=2
	} else {
		rollback()
	}
	if isCat(pp,Expression) && isCat(pp+1,Block) {
		tok_mem=append(tok_mem,pp,break_space,pp+1)
		pp+=2
		@<Making copy...@>
		if isCat(pp,else_token) && (isCat(pp+1,IfStmt) || isCat(pp+1,Block)) {
			tok_mem=append(tok_mem,break_space,pp,break_space,pp+1)
			pp+=2
		} else {
			rollback()
		}
		pp++
		reduce(p,pp-p,IfStmt,0,52,tok_mem...)
	} 
}
pp=p

@ Tests for |if|
@(tests/if.w@>=
@@
@@2
@@c
if x > max {
	x = max
}
@@
@@c
if x := f(); x < y {
	return x
} else if x > z {
	return z
} else {
	return y
}


@ @<Cases for |ExprSwitchStmt|@>=
p:=pp
if isCat(pp,switch_token) {
	tok_mem:=append([]interface{}{},pp)
	pp++
	{
		@<Making copy...@>
		if isCat(pp,SimpleStmt) && isCat(pp+1,semi) {
			tok_mem=append(tok_mem,break_space,pp,pp+1)
			pp+=2
		} else {
			rollback()
		}
	}
	{
		@<Making copy...@>
		if isCat(pp,Expression) {
			tok_mem=append(tok_mem,break_space,pp,break_space)
			pp++
		} else {
			rollback()
		}
	}
	if isCat(pp,lbrace) {
		c:=0
		isCats(pp+1,&c,cat_pair{cat:ExprCaseClause,mand:false})
		if isCat(pp+1+c,rbrace) {
			tok_mem=append(tok_mem,pp)
			for i:=0;i<c;i++ {
				if i==0 {
					tok_mem=append(tok_mem,force,indent)
				}
				tok_mem=append(tok_mem,pp+1+i,force)
				if i==c-1 {
					tok_mem=append(tok_mem,outdent)
				}
			}
			tok_mem=append(tok_mem,pp+1+c)
			pp+=2+c
			reduce(p,pp-p,ExprSwitchStmt,0,53,tok_mem...)
		}
	}
}
pp=p

@ @<Cases for |ExprCaseClause|@>=
if isCat(pp,case_token) && isCat(pp+1,ExpressionList) && isCat(pp+2,colon) {
	c:=0
	isCats(pp+3,&c,cat_pair{cat:Statement,mand:true},cat_pair{cat:semi,mand:false})
	tok_mem:=append([]interface{}{},pp,break_space,pp+1,pp+2,force)
	for i:=0;i<c;i++ {
		if i==0 {
			tok_mem=append(tok_mem,indent)
		}
		if isCat(pp+3+i,Statement) {
			tok_mem=append(tok_mem,pp+3+i,force)
		}
		if i==c-1 {
			tok_mem=append(tok_mem,outdent)
		}
	}
	reduce(pp,3+c,ExprCaseClause,0,54,tok_mem...)
} else if rollback(); isCat(pp,default_token) && isCat(pp+1,colon) {
	c:=0
	isCats(pp+2,&c,cat_pair{cat:Statement,mand:true},cat_pair{cat:semi,mand:false})
	tok_mem:=append([]interface{}{},pp,pp+1,force)
	for i:=0;i<c;i++ {
		if i==0 {
			tok_mem=append(tok_mem,indent)
		}
		if isCat(pp+2+i,Statement) {
			tok_mem=append(tok_mem,pp+2+i,force)
		}
		if i==c-1 {
			tok_mem=append(tok_mem,outdent)
		}
	}
	reduce(pp,2+c,ExprCaseClause,0,54,tok_mem...)
} else if rollback(); isCat(pp,section_scrap) {
	reduce(pp,1,ExprCaseClause,0,54,pp)
}

@ @<Cases for |TypeSwitchStmt|@>=
p:=pp
if isCat(pp,switch_token) {
	tok_mem:=append([]interface{}{},pp)
	if isCat(pp+1,SimpleStmt) && isCat(pp+2,semi) {
		tok_mem=append(tok_mem,break_space,pp+1,pp+2)
		pp+=2
	} else {
		rollback()
	}
	if isCat(pp+1,TypeSwitchGuard) && isCat(pp+2,lbrace) {
	 	c:=0
		isCats(pp+3,&c,cat_pair{cat:TypeCaseClause,mand:true})
		if isCat(pp+3+c,rbrace) {
			tok_mem=append(tok_mem,break_space,pp+1,break_space,pp+2)
			for i:=0;i<c;i++ {
				if i==0 {
					tok_mem=append(tok_mem,force,indent)
				}
				tok_mem=append(tok_mem,pp+3+i,force)
				if i==c-1 {
					tok_mem=append(tok_mem,outdent)
				}
			}
			tok_mem=append(tok_mem,pp+3+c)
			pp+=4+c
			reduce(p,pp-p,TypeSwitchStmt,0,55,tok_mem...)
		}
	}
}
pp=p

@ @<Cases for |TypeSwitchGuard|@>=
if isCat(pp,identifier) && isCat(pp+1,col_eq) && isCat(pp+2,PrimaryExpr) && isCat(pp+3,dot) && isCat(pp+4,lpar) && isCat(pp+5,type_token) && isCat(pp+6,rpar){
	reduce(pp,7,TypeSwitchGuard,0,56,pp,pp+1,pp+2,pp+3,pp+4,pp+5,pp+6)
} else if rollback(); isCat(pp,PrimaryExpr) && isCat(pp+1,dot) && isCat(pp+2,lpar) && isCat(pp+3,type_token) && isCat(pp+4,rpar) {
	reduce(pp,5,TypeSwitchGuard,0,56,pp,pp+1,pp+2,pp+3,pp+4)
}

@ @<Cases for |TypeCaseClause|@>=
if isCat(pp,TypeSwitchCase) && isCat(pp+1,colon) {
	c:=0
	isCats(pp+2,&c,cat_pair{cat:Statement,mand:true},cat_pair{cat:semi,mand:false})
	tok_mem:=append([]interface{}{},pp,pp+1,force)
	for i:=0;i<c;i++ {
		if i==0 {
			tok_mem=append(tok_mem,indent)
		}
		if isCat(pp+2+i,Statement) {
			tok_mem=append(tok_mem,pp+2+i,force)
		}
		if i==c-1 {
			tok_mem=append(tok_mem,outdent)
		}
	}
	reduce(pp,2+c,TypeCaseClause,0,57,tok_mem...)
} else if rollback(); isCat(pp,section_scrap) {
	reduce(pp,1,TypeCaseClause,0,57,pp)
}

@ @<Cases for |TypeSwitchCase|@>=
if isCat(pp,case_token) && isCat(pp+1,Type) {
	c:=0
	isCats(pp+2,&c,cat_pair{cat:comma,mand:true},cat_pair{cat:Type,mand:true})
	tok_mem:=append([]interface{}{},pp,break_space,pp+1)
	for i:=0;i<c;i++ {
		tok_mem=append(tok_mem,pp+2+i)
	}
	reduce(pp,2+c,TypeSwitchCase,0,58,tok_mem...)
} else if rollback(); isCat(pp,default_token) {
	reduce(pp,1,TypeSwitchCase,0,58,pp)
} else if rollback(); isCat(pp,section_scrap) {
	reduce(pp,1,TypeSwitchCase,0,58,pp)
}

@ Tests for |switch|
@(tests/switch.w@>=
@@
@@2
@@c
switch tag {
	default: s3()
	case 0, 1, 2, 3: s1()
	case 4, 5, 6, 7: s2()
}
@@
@@c
switch x := f(); {
	case x < 0: return -x
	default: return x
}
@@
@@c
switch {
	case x < y: f1()
	case x < z: f2()
	case x == 4: f3()
}
@@
@@c
switch i := x.(type) {
case nil:
	printString("x is nil")
case int:
	printInt(i)
case float64:
	printFloat64(i)
case func(int) float64:
	printFunction(i)
case bool, string:
	printString("type is bool or string")
default:
	printString("don't know the type")
}

@ @<Cases for |SelectStmt|@>=
if isCat(pp,select_token) && isCat(pp+1,lbrace){
	c:=0
	isCats(pp+2,&c,cat_pair{cat:CommClause,mand:false})
	if isCat(pp+2+c,rbrace) {
		tok_mem:=append([]interface{}{},pp,pp+1)
		for i:=0;i<c;i++ {
			if i==0 {
				tok_mem=append(tok_mem,force,indent)
			}
			if isCat(pp+2+i,CommClause) {
				tok_mem=append(tok_mem,pp+2+i)
			}
			if i==c-1 {
				tok_mem=append(tok_mem,outdent)
			}
		}
		tok_mem=append(tok_mem,pp+2+c)
		reduce(pp,3+c,SelectStmt,0,59,tok_mem...)
	}
}

@ @<Cases for |CommClause|@>=
if isCat(pp,CommCase) && isCat(pp+1,colon) {
	c:=0
	isCats(pp+2,&c,cat_pair{cat:Statement,mand:true},cat_pair{cat:semi,mand:false})
	tok_mem:=append([]interface{}{},pp,pp+1,force)
	for i:=0;i<c;i++ {
		if i==0 {
			tok_mem=append(tok_mem,indent)
		}
		if isCat(pp+2+i,Statement) {
			tok_mem=append(tok_mem,pp+2+i,force)
		} 
		if i==c-1 {
			tok_mem=append(tok_mem,outdent)
		}
	}
	reduce(pp,2+c,CommClause,0,60,tok_mem...)
}

@ @<Cases for |CommCase|@>=
if isCat(pp,case_token) {
	if isCat(pp+1,SendStmt) { 
		reduce(pp,2,CommCase,0,61,pp,break_space,pp+1)
	} else if rollback(); isCat(pp+1,RecvStmt) {
		reduce(pp,2,CommCase,0,61,pp,break_space,pp+1)
	}
} else if rollback(); isCat(pp,default_token) {
	reduce(pp,1,CommCase,0,61,pp)
} else if rollback(); isCat(pp,section_scrap) {
	reduce(pp,1,CommCase,0,61,pp)
}

@ @<Cases for |RecvStmt|@>=
if isCat(pp,ExpressionList) && (isCat(pp+1,eq) || isCat(pp+1,col_eq)) && isCat(pp+2,Expression) {
	reduce(pp,3,RecvStmt,0,62,pp,pp+1,pp+2)
} else if isCat(pp,Expression) {
	reduce(pp,1,RecvStmt,0,62,pp)
}

@ @<Cases for |SendStmt|@>=
if isCat(pp,Expression) && isCat(pp+1,direct) && isCat(pp+2,Expression) {
	reduce(pp,3,SendStmt,0,63,pp,pp+1,pp+2)
}

@ Tests for |send|
@(tests/send.w@>=
@@
@@2
@@c
ch <- 3

@ Tests for |select|
@(tests/select.w@>=
@@
@@2
@@c
select {
case i1 = <-c1:
	print("received ", i1, " from c1\n")
case c2 <- i2:
	print("sent ", i2, " to c2\n")
case i3, ok := (<-c3):  // same as: i3, ok := <-c3
	if ok {
		print("received ", i3, " from c3\n")
	} else {
		print("c3 is closed\n")
	}
default:
	print("no communication\n")
}
@@
@@2
@@c
select {
	case c <- 0:  // note: no statement, no fallthrough, no folding of cases
	case c <- 1:
}
@@
@@2
@@c
select {}

@ @<Cases for |ForStmt|@>=
if isCat(pp,for_token) {
	@<Making copy...@>
	if isCat(pp+1,Expression) && isCat(pp+2,Block) {
		reduce(pp,3,ForStmt,0,64,pp,break_space,pp+1,break_space,pp+2)
	} else if rollback(); isCat(pp+1,ForClause) && isCat(pp+2,Block) {
		reduce(pp,3,ForStmt,0,64,pp,break_space,pp+1,break_space,pp+2)
	} else if rollback(); isCat(pp+1,RangeClause) && isCat(pp+2,Block) {
		reduce(pp,3,ForStmt,0,64,pp,break_space,pp+1,break_space,pp+2)
	} else if rollback(); isCat(pp+1,Block) {
		reduce(pp,2,ForStmt,0,64,pp,pp+1)
	}   
}

@ @<Cases for |ForClause|@>=
p:=pp
var tok_mem []interface{}
if isCat(pp,SimpleStmt) {
	tok_mem=append(tok_mem,pp)
	pp++
} else {
	rollback()
}
if isCat(pp,semi) {
	tok_mem=append(tok_mem,pp)
	pp++
	@<Making copy...@>
	if isCat(pp,Expression) {
		tok_mem=append(tok_mem,break_space,pp)
		pp++
	} else {
		rollback()
	}
	if isCat(pp,semi) {
		tok_mem=append(tok_mem,pp)
		pp++
		@<Making copy...@>
		if isCat(pp,SimpleStmt) {
			tok_mem=append(tok_mem,break_space,pp)
			pp++
		} else {
			rollback()
		}
		reduce(p,pp-p,ForClause,0,65,tok_mem...)
	}
}
pp=p

@ @<Cases for |RangeClause|@>=
if isCat(pp,ExpressionList) && (isCat(pp+1,eq) || isCat(pp+1,col_eq)) && isCat(pp+2,range_token) && isCat(pp+3,Expression) {
	reduce(pp,4,RangeClause,0,66,pp,pp+1,pp+2,break_space,pp+3)
}

@ Tests for |for|
@(tests/for.w@>=
@@
@@2
@@c
for a < b {
	a *= 2
}
@@
@@c
for i := 0; i < 10; i++ {
	f(i)
}
@@
@@c
for i, _ := range testdata.a {
	f(i)
}
@@
@@c
for i, s := range a {
	g(i, s)
}


@ @<Cases for |DeferStmt|@>=
if isCat(pp,defer_token) && isCat(pp+1,Expression) {
	reduce(pp,2,DeferStmt,0,67,pp,break_space,pp+1)
}

@ Tests for |defer|
@(tests/defer.w@>=
@@
@@2
@@c
defer unlock(l) 
@@
@@c
defer func() {
                result++
        }()


@ @<Cases for |IncDecStmt|@>=
if isCat(pp,Expression) && (isCat(pp+1,plus_plus) || isCat(pp+1,minus_minus)) {
	reduce(pp,2,IncDecStmt,0,68,pp,pp+1)
}

@ Tests for |incdec|
@(tests/incdec.w@>=
@@
@@2
@@c
i++
@@
@@c
j--

@ @<Cases for |Assignment|@>=
if isCat(pp,ExpressionList) && isCat(pp+1,assign_op) && isCat(pp+2,ExpressionList) {
	reduce(pp,3,Assignment,0,69,pp,pp+1,pp+2)
}

@ Tests for assignments
@(tests/assign.w@>=
@@
@@2
@@c 
x = 1
@@
@@c
*p = f()
@@
@@c
a[i] = 23
@@
@@c
(k) = <-ch
@@
@@c
a[i] <<= 2
@@
@@c
i &^= 1<<n
@@
@@c
x, y = f()
@@
@@c
x, _ = f()
@@
@@c
a, b = b, a
@@
@@c
i, x[i] = 1, 2
@@
@@c
i = 0
@@
@@c
x[i], i = 2, 1
@@
@@c
x[0], x[0] = 1, 2
@@
@@c
x[1], x[3] = 4, 5 
@@
@@c
x[2], p.x = 6, 7
@@
@@c
i = 2
@@
@@c
x = []int{3, 5, 7}

@ @<Cases for |assign_op|@>=
if (isCat(pp,unary_op) || isCat(pp,mul_op) || isCat(pp,asterisk)) && isCat(pp+1,eq) {
	reduce(pp,2,assign_op,0,70,math_rel,'{',pp,'}','{',pp+1,'}','}')
} else if rollback(); isCat(pp,eq) {
	reduce(pp,1,assign_op,0,70,pp)
}

@ @<Cases for |ShortVarDecl|@>=
if isCat(pp,IdentifierList) && isCat(pp+1,col_eq) && isCat(pp+2,ExpressionList) {
	reduce(pp,3,ShortVarDecl,0,71,pp,pp+1,pp+2)
}

@ Tests for short var declarations
@(tests/shortvar.w@>=
@@
@@2
@@c
i, j := 0, 10
@@
@@c
f := func() int { return 7 }
@@
@@c
ch := make(chan int)
@@
@@c
r, w := os.Pipe(fd) 
@@
@@c
_, y, _ := coord(p)

@ @<Cases for |QualifiedIdent|@>=
if (isCat(pp,identifier) || isCat(pp,PackageName)) && isCat(pp+1,dot) && isCat(pp+2,identifier) {
	reduce(pp,3,QualifiedIdent,0,72,pp,pp+1,pp+2)
} else if rollback(); isCat(pp,identifier) {
	reduce(pp,1,QualifiedIdent,0,72,pp)
}

@ @<Cases for |MethodExpr|@>=
if isCat(pp,ReceiverType) && isCat(pp+1,dot) && isCat(pp+2,identifier) {
	reduce(pp,3,MethodExpr,0,73,pp,pp+1,pp+2)
}

@ @<Cases for |ReceiverType|@>=
if isCat(pp,Type) {
	reduce(pp,1,ReceiverType,0,74,pp)
} else if rollback(); isCat(pp,lpar) && isCat(pp+1,asterisk) && isCat(pp+2,Type) && isCat(pp+3,rpar) {
	reduce(pp,4,ReceiverType,0,74,pp,pp+1,pp+2,pp+3)
}

@ @<Cases for |Conversion|@>=
if isCat(pp,Type) && isCat(pp+1,lpar) && isCat(pp+2,Expression) && isCat(pp+3,rpar) {
	reduce(pp,4,Conversion,0,75,pp,pp+1,pp+2,pp+3)
}

@ @<Cases for |BuiltinCall|@>=
if isCat(pp,identifier) && isCat(pp+1,lpar) {
	c:=0
	isCats(pp+2,&c,cat_pair{cat:BuiltinArgs,mand:true},cat_pair{cat:comma,mand:false}) 
	if isCat(pp+2+c,rpar) {
		tok_mem:=append([]interface{}{},pp,pp+1)
		for i:=0;i<c;i++ {
			tok_mem=append(tok_mem,pp+2+i)
		}
		tok_mem=append(tok_mem,pp+2+c)
		reduce(pp,3+c,BuiltinCall,0,76,tok_mem...)
	}
}

@ @<Cases for |BuiltinArgs|@>=
if isCat(pp,Type) {
	c:=0
	isCats(pp+1,&c,cat_pair{cat:comma,mand:true},cat_pair{cat:ExpressionList,mand:true})
	tok_mem:=append([]interface{}{},pp)
	for i:=0;i<c;i++ {
		tok_mem=append(tok_mem,pp+1+i)
	}
	reduce(pp,1+c,BuiltinArgs,0,77,tok_mem...)	
} else if rollback(); isCat(pp,ExpressionList) {
	reduce(pp,1,BuiltinArgs,0,77,pp)	
}

@ @<Cases for |Selector|@>= 
if isCat(pp,dot) && isCat(pp+1,identifier) {
	reduce(pp,2,Selector,0,78,pp,pp+1)
}

@ @<Cases for |Index|@>=
if isCat(pp,lbracket) && isCat(pp+1,Expression) && isCat(pp+2,rbracket) {
	reduce(pp,3,Index,0,79,pp,pp+1,pp+2)
}

@ @<Cases for |Slice|@>=
if isCat(pp,lbracket) {
	c1:=0
	isCats(pp+1,&c1, cat_pair{cat:Expression,mand:false})
	if isCat(pp+1+c1,colon) {
		c2 := 0
		isCats(pp+2+c1,&c2, cat_pair{cat:Expression,mand:false})
		if isCat(pp+2+c1+c2,rbracket) {
			tok_mem:=append([]interface{}{},pp)
			for i:=0;i<c1;i++ {
				tok_mem=append(tok_mem,pp+1+i)
			}
			tok_mem=append(tok_mem,pp+1+c1)
			for i:=0;i<c2;i++ {
				tok_mem=append(tok_mem,pp+2+c1+i)
			}
			tok_mem=append(tok_mem,pp+2+c1+c2)
			reduce(pp,3+c1+c2,Slice,0,80,tok_mem...)
		}
	}
}

@ @<Cases for |TypeAssertion|@>=
if isCat(pp,dot) && isCat(pp+1,lpar) && isCat(pp+2,Type) && isCat(pp+3,rpar) {
	reduce(pp,4,TypeAssertion,0,81,pp,pp+1,pp+2,pp+3)
}

@ @<Cases for |Call|@>=
if isCat(pp,lpar) {
	c:=0
	isCats(pp+1,&c,cat_pair{cat:ExpressionList,mand:false}, cat_pair{cat:dot_dot_dot,mand:false})
	if isCat(pp+1+c,rpar) {
		tok_mem:=append([]interface{}{},pp)
		for i:=0;i<c;i++ {
			tok_mem=append(tok_mem,pp+1+i)
		}
		tok_mem=append(tok_mem,pp+1+c)
		reduce(pp,2+c,Call,0,82,tok_mem...)
	}	
}

@ @<Cases for |unary_op|@>=
if isCat(pp,asterisk) || isCat(pp,direct) || isCat(pp,add_op) {
	reduce(pp,1,unary_op,0,83,pp)
}

@ Now here's the |reduce| procedure used in our code for productions.

@c
func reduce(pp int, k int, c rune, d int, n int, s ...interface{}) {
	reduced=true
	reduced_cat=c
	var trans []interface{}
	cur_mathness:=maybe_math
	init_mathness:=maybe_math

	for _,t:=range s {
		switch v := t.(type) {
			case rune: 
				if v==' ' || (v>=big_cancel && v<=big_force) /* non-math token */ {
					if cur_mathness==maybe_math { 
						init_mathness=no_math
					} else if cur_mathness==yes_math { 
						trans=append(trans,"{}$") 
					}
					cur_mathness=no_math
				} else {
					if scrap_info[pp].mathness==maybe_math { 
						init_mathness=yes_math
					} else if scrap_info[pp].mathness==no_math { 
						trans=append(trans,"${}") 
					}
					cur_mathness=yes_math
				}
				trans=append(trans,v)
			case int: 
				switch scrap_info[v].mathness % 4 { /* left boundary */
					case no_math:
						if cur_mathness==maybe_math {
							init_mathness=no_math
						} else if cur_mathness==yes_math { 
							trans=append(trans,"{}$") 
						}
						cur_mathness=scrap_info[v].mathness / 4 /* right boundary */
					case yes_math:
						if cur_mathness==maybe_math { 
							init_mathness=yes_math 
						} else if cur_mathness==no_math {
							trans=append(trans,"${}")
						}
						cur_mathness=scrap_info[v].mathness / 4 /* right boundary */
					case maybe_math: /* no changes */
				}
				trans=append(trans,scrap_info[v].trans)
			default:
				panic(fmt.Sprintf( "Invalid type of translation: %T(%v)",v,v))
		}
	}
	if k==1 {
		scrap_info[pp].cat=c
	} else {
		if init_mathness==maybe_math && cur_mathness!=maybe_math {
			init_mathness=cur_mathness
		}
		scrap_info[pp] = scrap{cat: c, trans: trans, mathness: 4*cur_mathness+init_mathness,}
		copy(scrap_info[pp+1:len(scrap_info)-1],scrap_info[pp+k:])
		scrap_info = scrap_info[:len(scrap_info)-k+1]
	}
	f := "reduce"
	@<Print a snapshot of the scrap list if debugging@>
}

@ And here now is the code that applies productions as long as possible.
Before applying the production mechanism, we must make sure
it has good input (at least four scraps, the length of the lhs of the
longest rules), and that there is enough room in the memory arrays
to hold the appended tokens and texts.  Here we use a very
conservative test; it's more important to make sure the program
will still work if we change the production rules (within reason)
than to squeeze the last bit of space from the memory arrays.

@ A variable |reduced_cat| is a category was applied.
@<Global...@>=
var reduced_cat rune = -1

@ A variable |reduced| is a flag of reducing was made.
@<Global...@>=
var reduced bool = false

@ @<Reduce the scraps using the productions until no more rules apply@>=
for  {
	if pp>=len(scrap_info) {
		break
	}
	@<Match a production...@>
}

@ If \.{GOWEAVE} is being run in debugging mode, the production numbers and
current stack categories will be printed out when |tracing| is set to 2;
a sequence of two or more irreducible scraps will be printed out when
|tracing| is set to 1.

@<Global...@>=
var tracing int32  /* can be used to show parsing details */

@ @<Print a snapsh...@>=
{ 
	if (tracing & 2) == 2 {
		fmt.Printf("%s %d:", f, n)
		for k, v:=range scrap_info {
			if k==pp {
				fmt.Print("*") 
			} else {
				fmt.Print(" ")
			}
			if v.mathness %4 == yes_math {
				fmt.Print("+")
			} else if v.mathness %4 == no_math {
				fmt.Print("-")
			}
			print_cat(v.cat)
			if v.mathness /4 == yes_math {
				fmt.Print("+")
			} else if v.mathness /4 == no_math {
				fmt.Print("-")
			}
		}
		fmt.Println()
	}
}

@ The |translate| function assumes that scraps have been stored in
|scrap_info| of |cat| and |trans|. It
applies productions as much as
possible. The result is a token list containing the translation of
the given sequence of scraps.

@c 
/* converts a sequence of scraps */
func translate() []interface{} {
	pp:=0
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
	var tok_mem []interface{}
	for i,v:=range scrap_info {
		if i!=0 {
			tok_mem=append(tok_mem,' ')
		}
		if v.mathness % 4 == yes_math {
			tok_mem=append(tok_mem,'$')
		}
		tok_mem=append(tok_mem,v.trans...)
		if v.mathness / 4 == yes_math {
			tok_mem=append(tok_mem,'$')
		}
	}
	return tok_mem
}

@ @<If semi-tracing, show the irreducible scraps@>=
if len(scrap_info)>0 && tracing==1 {
	fmt.Printf("\nIrreducible scrap sequence in section %d:",section_count)
@.Irreducible scrap sequence...@>
	mark_harmless()
	for i,_:=range scrap_info {
		fmt.Printf(" ")
		print_cat(scrap_info[i].cat)
	}
}

@ @<If tracing,...@>=
if (tracing & 2) == 2 {
	fmt.Printf("\nTracing after l. %d:\n",line[include_depth])
	mark_harmless()
@.Tracing after...@>
}

@* Initializing the scraps.
If we are going to use the powerful production mechanism just developed, we
must get the scraps set up in the first place, given a \GO/ text. A table
of the initial scraps corresponding to \GO/ tokens appeared above in the
section on parsing; our goal now is to implement that table. We shall do this
by implementing a subroutine called |Go_parse| that is analogous to the
|Go_xref| routine used during phase one.

Like |Go_xref|, the |Go_parse| procedure starts with the current
value of |next_control| and it uses the operation |next_control=get_next()|
repeatedly to read \GO/ text until encountering the next `\.{\v}' or
`\.{/*}', or until |next_control>=format_code|. The scraps corresponding to
what it reads are appended into the |cat| and |trans| arrays, and |scrap_ptr|
is advanced.

@c
/* creates scraps from \GO/ tokens */
func Go_parse(spec_ctrl rune) {
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
func app_scrap(c int32, b int32, t ...interface{}) {
	scrap_info = append(scrap_info, scrap{cat:c, trans: t, mathness:5*b,})
}

@ @<Append the scr...@>=
switch (next_control) {
	case section_name:
		app_scrap(section_scrap,maybe_math,section_token(cur_section))
	case str,constant,verbatim:
		@<Append a string or constant@>
	case identifier: 
		app_cur_id()
	case TeX_string:
		@<Append a \TEX/ string, without forming a scrap@>
	case '/':
		app_scrap(mul_op,yes_math,next_control)
		next_control=mul_op
	case '.':
		app_scrap(dot,yes_math,next_control)
		next_control=dot
	case '_':
		app_scrap(identifier,maybe_math,"\\_")
		next_control=identifier
	case '<': 
		@+app_scrap(rel_op,yes_math,"\\langle")
		next_control=rel_op
	case '>': 
		@+app_scrap(rel_op,yes_math,"\\rangle")
		next_control=rel_op
	case '=': 
		app_scrap(eq,yes_math,"\\K")
		next_control=eq
@.\\K@>
	case '|': 
		app_scrap(add_op,yes_math,"\\OR")
		next_control=add_op
@.\\OR@>
	case '^': 
		app_scrap(add_op,yes_math,"\\XOR")
		next_control=add_op
@.\\XOR@>
	case '%': 
		app_scrap(mul_op,yes_math,"\\MOD")
		next_control=mul_op
@.\\MOD@>
	case '!': 
		app_scrap(unary_op,yes_math,"\\R")
		next_control=unary_op
@.\\R@>
	case '+', '-':
		app_scrap(add_op,yes_math,next_control)
		next_control=add_op
	case '*': 
		app_scrap(asterisk,yes_math,next_control)
		next_control=asterisk
	case '&': 
		app_scrap(mul_op,yes_math,"\\AND")
		next_control=mul_op
@.\\AND@>
	case ignore, xref_roman, xref_wildcard, xref_typewriter, noop:
		@+break
	case '(' : 
		app_scrap(lpar,maybe_math,next_control)
		next_control=lpar
	case ')' : 
		app_scrap(rpar,maybe_math,next_control)
		next_control=rpar
	case '[': 
		app_scrap(lbracket,maybe_math,next_control)
		next_control=lbracket
	case ']': 
		app_scrap(rbracket,maybe_math,next_control)
		next_control=rbracket
	case '{': 
		app_scrap(lbrace,yes_math,"\\{"@q}@>)
		next_control=lbrace
@.\\\{@>@q}@>
	case '}': 
		app_scrap(rbrace,yes_math,@q{@>"\\}")
		next_control=rbrace
@q{@>@.\\\}@>
	case ',': 
		app_scrap(comma,yes_math,next_control,opt,'9',)
		next_control=comma
	case ';': 
		app_scrap(semi,maybe_math,next_control)
		next_control=semi
	case ':': 
		app_scrap(colon,no_math,next_control)@/
		next_control=colon
	@t\4@>  @<Cases involving nonstandard characters@>
	case thin_space: 
		app_scrap(insert,maybe_math,"\\,")
		next_control=thin_space
@.\\,@>
	case math_break: 
		app_scrap(insert,maybe_math,opt,"0")
		next_control=insert
	case line_break: 
		app_scrap(insert,no_math,force)
		next_control=insert
	case big_line_break: 
		app_scrap(insert,no_math,big_force)
		next_control=insert
	case no_line_break: 
		app_scrap(insert,no_math,big_cancel,noop,break_space,noop,big_cancel)
		next_control=insert
	case pseudo_semi: 
		next_control=semi
		app_scrap(semi,maybe_math)
	case join: 
		app_scrap(insert,no_math,"\\J")
		next_control=insert
@.\\J@>
	default: 
		app_scrap(insert,maybe_math,inserted,next_control)
		next_control=insert
}

@ Some nonstandard characters may have entered \.{GOWEAVE} by means of
standard ones. They are converted to \TEX/ control sequences so that it is
possible to keep \.{GOWEAVE} from outputting unusual |rune| codes.

@<Cases involving nonstandard...@>=
case not_eq: 
	@+app_scrap(rel_op,yes_math,"\\I")
@.\\I@>
case lt_eq: 
	@+app_scrap(rel_op,yes_math,"\\Z")
@.\\Z@>
case gt_eq: 
	@+app_scrap(rel_op,yes_math,"\\G")
@.\\G@>
case eq_eq: 
	@+app_scrap(rel_op,yes_math,"\\E")
@.\\E@>
case and_and: 
	@+app_scrap(binary_op,yes_math,"\\W")
@.\\W@>
case or_or: 
	@+app_scrap(binary_op,yes_math,"\\V")
@.\\V@>
case plus_plus: 
	@+app_scrap(plus_plus,yes_math,"\\PP")
@.\\PP@>
case minus_minus: 
	@+app_scrap(minus_minus,yes_math,"\\MM")
@.\\MM@>
case gt_gt: 
	@+app_scrap(mul_op,yes_math,"\\GG")
@.\\GG@>
case lt_lt: 
	@+app_scrap(mul_op,yes_math,"\\LL")
@.\\LL@>
case dot_dot_dot: 
	@+app_scrap(dot_dot_dot,yes_math,"\\,\\ldots\\,")
@.\\,@>
@.\\ldots@>
case col_eq: 
	@+app_scrap(col_eq,yes_math,":\\K")
case direct:
	@+app_scrap(direct,yes_math,"\\leftarrow")
case and_not:
	@+app_scrap(mul_op,yes_math,"\\AND\\XOR")


@ Many of the special characters in a string must be prefixed by `\.\\' so that
\TEX/ will print them properly.
@^special string characters@>

@<Append a string or...@>=
count:= -1 
var tok_mem []interface{}
if next_control==constant {
	tok_mem=append(tok_mem,"\\T{"@q}@>)
@.\\T@>
} else if next_control==str {
	count=20
	tok_mem=append(tok_mem,"\\.{"@q}@>)
@.\\.@>
} else {
	tok_mem=append(tok_mem,"\\vb{"@q}@>)
}
@.\\vb@>
for i:=0; i < len(id); {
	if count==0 { /* insert a discretionary break in a long string */
		tok_mem=append(tok_mem,@q(@>@q{@>"}\\)\\.{"@q}@>)
		count=20
@q(@>@.\\)@>
	}
	switch (id[i]) {
		case ' ', '\\', '#', '%', '$','^', '{', '}', '~', '&', '_': 
			tok_mem=append(tok_mem,'\\')
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
	tok_mem=append(tok_mem,id[i])
	i++
	count--
}
tok_mem=append(tok_mem,@q{@>'}')
app_scrap(next_control,maybe_math,tok_mem...)

@ We do not make the \TEX/ string into a scrap, because there is no
telling what the user will be putting into it; instead we leave it
open, to be picked up by the next scrap. If it comes at the end of a
section, it will be made into a scrap when |finish_Go| is called.

@<Append a \TEX/ string, without forming a scrap@>=
tok_mem:=append([]interface{}{},"\\hbox{"@q}@>)
for i:=0; i < len(id);{ 
	if id[i]=='@@' {
		i++
	}
	tok_mem=append(tok_mem,id[i])
	i++
}
tok_mem=append(tok_mem,@q{@>'}')
app_scrap(insert,no_math,tok_mem...)

@ The function |app_cur_id| appends the current identifier to the
token list; it also builds a new scrap if |scrapping==true|.

@ @c
func app_cur_id() {
	p:=id_lookup(id,normal)
	if name_dir[p].ilk<=custom { /* not a reserved word */
		a1 := identifier
		a2 := maybe_math
		if name_dir[p].ilk==custom {
			a2 = yes_math
		} 
		app_scrap(a1, a2,id_token(p))
	} else {
		if name_dir[p].ilk==binary_op || 
			name_dir[p].ilk==rel_op || 
			name_dir[p].ilk==add_op ||
			name_dir[p].ilk==mul_op {
			app_scrap(name_dir[p].ilk,yes_math,res_token(p))
		} else {
			app_scrap(name_dir[p].ilk,maybe_math,res_token(p))
		}
	}
}

@ When the `\.{\v}' that introduces \GO/ text is sensed, a call on
|Go_translate| will return a pointer to the \TEX/ translation of
that text. If scraps exist in |scrap_info|, they are
unaffected by this translation process.

@c
func Go_translate() []interface{} {
	save_scraps:=scrap_info /* holds original value of |scrap_info| */
	scrap_info=nil
	Go_parse(section_name) /* get the scraps together */
	if next_control!='|' {
		err_print("! Missing '|' after Go text")
@.Missing '|'...@>
	}
	app_scrap(insert,maybe_math,cancel)
				/* place a |cancel| token as a final ``comment'' */
	p:=translate() /* make the translation */
	scrap_info=save_scraps /* scrap the scraps */
	return p
}

@ The |outer_parse| routine is to |Go_parse| as |outer_xref|
is to |Go_xref|: It constructs a sequence of scraps for \GO/ text
until |next_control>=format_code|. Thus, it takes care of embedded comments.

The token list created from within `\pb' brackets is output as an argument
to \.{\\PB}, if the user has invoked \.{GOWEAVE} with the \.{+e} flag.
Although \.{gowebmac} ignores \.{\\PB}, other macro packages
might use it to localize the special meaning of the macros that mark up
program text.

@c
/* makes scraps from \GO/ tokens and comments */
func outer_parse() {
	for next_control<format_code {
		var tok_mem []interface{}
		if next_control!=begin_comment && next_control!=begin_short_comment {
			Go_parse(ignore)
		} else {
			is_long_comment:=(next_control==begin_comment)
			tok_mem=append(tok_mem,cancel,inserted)
			if is_long_comment {
				tok_mem=append(tok_mem,"\\C{"@q}@>)
@.\\C@>
			} else {
				tok_mem=append(tok_mem,"\\SHC{"@q}@>)
			}
@.\\SHC@>
			bal,tok_mem:=copy_comment(is_long_comment,1,tok_mem)  /* brace level in comment */
			next_control=ignore
			for bal>0 {
				p:=tok_mem
				tok_mem=nil
				q:=Go_translate()/* partial comments */
				tok_mem=append(tok_mem,list_token(p))
				if flags['e'] {
					tok_mem=append(tok_mem,"\\PB{")
@.\\PB@>
				}
				tok_mem=append(tok_mem,inner_list_token(q))
				if flags['e'] {
					tok_mem=append(tok_mem,'}') 
				}
				if next_control=='|' {
					bal,tok_mem=copy_comment(is_long_comment,bal,tok_mem)
					next_control=ignore
				} else {
					bal=0 /* an error has been reported */
				}
			}
			tok_mem=append(tok_mem,force)
			app_scrap(insert,no_math,tok_mem...)
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
const (
	inner mode = 0 /* value of |mode| for \GO/ texts within \TEX/ texts */
	outer mode = 1 /* value of |mode| for \GO/ texts in sections */
)

@ @<Typed...@>= 
type output_state struct {
	tok_field []interface{} /* present location of token list */
	mode_field mode /* interpretation of control tokens */
}

@ @c func init_stack() {
	stack=make([]output_state, 0, 100)
	cur_state.mode_field=outer
}

@ @<Global...@>=
var cur_state output_state /* |cur_state.tok_field|, |cur_state.mode_field| */
var stack[]output_state /* info for non-current levels */

@ To insert token-list |p| into the output, the |push_level| subroutine
is called; it saves the old level of output and gets a new one going.
The value of |cur_state.mode_field| is not changed.

@c
 /* suspends the current level */
func push_level(tokens []interface{}) {
	stack = append(stack, output_state{tok_field:cur_state.tok_field, mode_field:cur_state.mode_field,})
	cur_state.tok_field=tokens
}

@ Conversely, the |pop_level| routine restores the conditions that were in
force when the current level was begun. This subroutine will never be
called when |stack_ptr==1|.

@c
func pop_level() bool {
	if len(stack) == 0 {
		return false
	}
	p := len(stack) - 1
	cur_state=stack[p]
	stack=stack[:p]	
	return true
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
const (
	res_word rune = 0242 /* returned by |get_output| for reserved words */
	section_code rune = 0243 /* returned by |get_output| for section names */
)

@ @c
/* returns the next token of output */
func get_output() rune {
restart: 
	for len(cur_state.tok_field)==0 {
		if !pop_level() {
			return -1
		} 
	}
	val:=cur_state.tok_field[0]
	cur_state.tok_field = cur_state.tok_field[1:]
	switch tok := val.(type) {
		case id_token:
			cur_name = int32(tok)
			return identifier/* |a==id_flag+cur_name| */
		case res_token: 
			cur_name = int32(tok)
			return res_word /* |a==res_flag+cur_name| */
		case section_token: 
			cur_name = int32(tok)
			return section_code /* |a==section_flag+cur_name| */
		case inner_list_token: 	/* |a==inner_tok_flag+cur_name| */
			cur_state.mode_field=inner
			push_level(tok)
			goto restart
		case list_token: /* |a==tok_flag+cur_name| */
			push_level(tok)
			goto restart
		case rune: 
			return tok
		case []interface{}:
			push_level(tok)
			goto restart
		case string:
			var tok_mem []interface{}
			for _, v := range tok {
				tok_mem=append(tok_mem,v)
			}			
			push_level(tok_mem)
			goto restart
	}
	panic(fmt.Sprintf( "Invalid type of scrap: %T(%v)", val,val))
}

@ The real work associated with token output is done by |make_output|.
This procedure appends an |end_translation| token to the current token list,
and then it repeatedly calls |get_output| and feeds characters to the output
buffer until reaching the |end_translation| sentinel. It is possible for
|make_output| to be called recursively, since a section name may include
embedded \GO/ text; however, the depth of recursion never exceeds one
level, since section names cannot be inside of section names.

A procedure called |output_Go| does the scanning, translation, and
output of \GO/ text within `\pb' brackets, and this procedure uses
|make_output| to output the current token list. Thus, the recursive call
of |make_output| actually occurs when |make_output| calls |output_Go|
while outputting the name of a section.
@^recursion@>

@c
/* outputs the current token list */
func output_Go() {
	save_next_control:=next_control/* values to be restored */
	next_control=ignore
	p:=Go_translate()/* translation of the \GO/ text */
	if flags['e'] {
		out_str("\\PB{")
		make_output(inner_list_token(p))
		out('}')
@.\\PB@>
	}@+else {
		make_output(inner_list_token(p)) /* output the list */
	}
	next_control=save_next_control /* restore |next_control| to original state */
}

@ Here is \.{GOWEAVE}'s major output handler.

@ @c
/* outputs the equivalents of tokens */
func make_output(p interface{}) {
	var c int /* count of |indent| and |outdent| tokens */
	tok_mem:=append([]interface{}{},p,end_translation) /* append a sentinel */
	push_level(tok_mem)
	tok_mem=nil
	var b rune
	for {
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
				for {
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
				force, big_force: 
					@<Output a control, look ahead in case of line breaks, possibly |goto reswitch|@>
			case quoted_char: 
				out(cur_state.tok_field[0].(rune))
				cur_state.tok_field = cur_state.tok_field[1:]
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
}@+else {
	out('&') /* |a==res_word| */
}
@.\\\&@>
if is_tiny(cur_name) {
	if name_dir[cur_name].name[0]=='_'  {
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
	} else {
		out(v)
	}
}
break

@ The current mode does not affect the behavior of \.{GOWEAVE}'s output routine
except when we are outputting control tokens.

@<Output a control...@>=
if a<break_space {
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
	for {
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
		output_Go()
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
for {
	if i>=len(scratch) {
		fmt.Print("\n! Go text in section name didn't end: <")
@.Go text...didn't end@>
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
if loc-1 >= len(buffer) || buffer[loc-1]!='*' {
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
for {
	next_control=copy_TeX()
	switch next_control {
		case '|': 
			init_stack()
			output_Go()
		case '@@': 
			out('@@')
		case TeX_string, noop, xref_roman, xref_wildcard, xref_typewriter, section_name: 
			loc-=2
			next_control=get_next() /* skip to \.{@@>} */
			if next_control==TeX_string {
				err_print("! TeX string should be in Go text only")
@.TeX string should be...@>
			}
		case thin_space,math_break,ord,
		line_break, big_line_break, no_line_break, join,
		pseudo_semi:
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
for next_control<=format_code { /* |format_code| or |definition| */
	init_stack() 
	@<Start a format definition@>
	outer_parse()
	finish_Go(format_visible)
	format_visible=true
	doing_format=false
}

@ The |finish_Go| procedure outputs the translation of the current
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
func finish_Go(visible bool) {
	if visible {
		out_str("\\B")
		app_scrap(insert,no_math,force)
		p:=translate() /* translation of the scraps */
@.\\B@>
		scrap_info=nil
		make_output(list_token(p)) /* output the list */
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
	tok_mem:=append([]interface{}{},"\\F") /* this will produce `\&{format }' */
@.\\F@>
	next_control=get_next()
	if next_control==identifier {
		tok_mem=append(tok_mem,id_token(id_lookup(id,normal)),' ',break_space) /* this is syntactically separate from what follows */
		next_control=get_next()
		if next_control==identifier {
			tok_mem=append(tok_mem,id_token(id_lookup(id,normal)))
			app_scrap(Expression,maybe_math,tok_mem...)
			next_control=get_next()
		}
	}
	if len(scrap_info)!=2 {
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
	finish_Go(true)
}

@ The title of the section and an $\E$ or $\mathrel+\E$ are made
into a scrap that should not take part in the parsing.

@<Check that '='...@>=
for {
	next_control=get_next()
	if next_control!='+' {
		break
	}
} /* allow optional `\.{+=}' */
var tok_mem []interface{}
if next_control!='=' && next_control!=eq_eq {
	err_print("! You need an = sign after the section name")
@.You need an = sign...@>
} else {
	next_control=get_next()
}
if out_ptr>1 && out_buf[out_ptr]=='Y' && out_buf[out_ptr-1]=='\\' {
	tok_mem=append(tok_mem,backup)
}
		/* the section name will be flush left */
@.\\Y@>
tok_mem=append(tok_mem,section_token(this_section))
cur_xref=name_dir[this_section].xref
if xmem[cur_xref].num==file_flag {
	cur_xref=xmem[cur_xref].xlink
}
tok_mem=append(tok_mem,"${}")
if xmem[cur_xref].num!=section_count+def_flag {
	tok_mem=append(tok_mem,"\\mathrel+") /*section name is multiply defined*/
	this_section=-1 /*so we won't give cross-reference info here*/
}
tok_mem=append(tok_mem,"\\E","{}$",force) /* output an equivalence sign */
@.\\E@>
app_scrap(dead,no_math,tok_mem...)
				/* this forces a line break unless `\.{@@+}' follows */

@ @<Emit the scrap...@>=
if next_control<section_name {
	err_print("! You can't do that in Go text")
@.You can't do that...@>
	next_control=get_next()
} else if next_control==section_name {
	app_scrap(section_scrap,maybe_math,section_token(cur_section))
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
for {
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
			fatal("! Cannot open section file ",scn_file_name)
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

@ During the sorting phase we shall use the |cat| and |trans| arrays from
\.{GOWEAVE}'s parsing algorithm and rename them |depth| and |head|. They now
represent a stack of identifier lists for all the index entries that have
not yet been output. The variable |sort_ptr| tells how many such lists are
present; the lists are output in reverse order (first |sort_ptr|, then
|sort_ptr-1|, etc.). The |j|th list starts at |head[j]|, and if the first
|k| characters of all entries on this list are known to be equal we have
|depth[j]==k|.

@ @<Rest of |scrap| struct@>=
head int32

@ @<Type...@>=
type sort_pointer int32

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
const infinity = -1  /* $\infty$ (approximately) */

@ @c
/* empties buckets having depth |d| */
func unbucket(d int32) {
	/* index into |bucket|; cannot be a simple |char| because of sign
		comparison below*/
	for c:=100+128; c>= 0; c-- {
		if bucket[collate[c]] != -1 {
@^high-bit character handling@>
			sort_ptr++
			scrap_info = append(scrap_info, scrap{})
			if sort_ptr>max_sort_ptr {
				max_sort_ptr=sort_ptr
			}
			if c==0 {
				scrap_info[sort_ptr].cat=infinity
			} else {
				scrap_info[sort_ptr].cat=d
			}
			scrap_info[sort_ptr].head=bucket[collate[c]]
			bucket[collate[c]]=-1
		}
	}
}

@ @<Sort and output...@>=
sort_ptr=0
scrap_info = append(scrap_info, scrap{})
unbucket(1)
for sort_ptr>0 {
	cur_depth=scrap_info[sort_ptr].cat
	if blink[scrap_info[sort_ptr].head]==-1 || cur_depth==infinity {
		@<Output index entries for the list at |sort_ptr|@>
	} else {
		@<Split the list at |sort_ptr| into further lists@> 
	}
}

@ @<Split the list...@>= {
	next_name:=scrap_info[sort_ptr].head
	for {
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
	cur_name=scrap_info[sort_ptr].head
	for {
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
	case normal: 
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
		out_str("\\9")
		out_name(cur_name,false)
		goto name_done
@.\\9@>
	case typewriter: 
		out_str("\\.")
@.\\.@>
		fallthrough 
	case roman:
		out_name(cur_name,false)
		goto name_done
	case custom: {
		out_str("$\\")
		for _, v := range name_dir[cur_name].name {
			if v == '_' {
				out('x')
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
for {
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
for {
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
		init_stack()
		make_output(section_token(p))
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
	fmt.Println("\nMemory usage statistics:\n")
@.Memory usage statistics:@>
	fmt.Println("%v names", len(name_dir))
	fmt.Println("Parsing:")
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
