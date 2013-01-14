% This file is part of GOWEB.
% This program by Alexander Sychev 
% is based on a program CWEB by Silvio Levy and Donald E. Knuth
% It is distributed WITHOUT ANY WARRANTY, express or implied.
% Version 0.5 --- January 2013

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
\def\pb{$\.|\ldots\.|$} % C brackets (|...|)
\def\v{\char'174} % vertical (|) in typewriter font
\mathchardef\RA="3221 % right arrow
\mathchardef\BA="3224 % double arrow

\def\title{GOTANGLE (Version 0.5)}
\def\topofcontents{\null\vfill
  \centerline{\titlefont The {\ttitlefont GOTANGLE} processor}
  \vskip 15pt
  \centerline{(Version 0.5)}
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

@** Introduction.
This is the \.{GOTANGLE} program by Alexander Sychev,
based on \.{CTANGLE} by Silvio Levy and Donald E. Knuth.

The ``banner line'' defined here should be changed whenever \.{GOTANGLE}
is modified.

@<Constants@>=
banner = "This is GOTANGLE (Version 0.5)\n"

@
@c
package main @#

import (
@<Import packages@>
) @#

const (
@<Constants@>
) @#


@<Typedef declarations@> @#
@<Global variables@> @#

@ \.{GOTANGLE} has a fairly straightforward outline.  It operates in
two phases: first it reads the source file, saving the \GO/ code in
compressed form; then it shuffles and outputs the code.

@c
func main () {
	common_init()
	@<Set initial values@>
	if show_banner() {
		fmt.Print(banner) /* print a ``banner line'' */
	}
	phase_one() /* read all the user's text and compress it into |tok_mem| */
	phase_two() /* output the contents of the compressed tables */
	os.Exit(wrap_up()) /* and exit gracefully */
}

@ @<Constants@>=
max_texts = 2500 /* number of replacement texts, must be less than 10240 */

@ The next few sections contain stuff from the file \.{common.w} that must
be included in both \.{gotangle.w} and \.{goweave.w}. 

@i common.w


@** Data structures exclusive to {\tt GOTANGLE}.
A |text| is a structure containing a token into
|token|, and an integer |text_link|, which, as we shall see later, is used to connect
pieces of text that have the same name.  All the |text| are stored in
the array |text_info|.

@<Typed...@>=
type text struct {
	token []rune	/* pointer into |tok_mem| */
	text_link int32		/* relates replacement texts */
}

@ @<Glob...@>=
var text_info []text
var tok_mem []rune

@ If |p| is an index of a section name, |p.equiv| is an index of its
replacement text, an element of the array |text_info|.

@<More elements of |name...@>=
equiv int32 /* info corresponding to names */

@ Here's the procedure that decides whether a name |id| equals 
the identifier pointed to by |p|:

@c
func names_match(@t\1@>@/
	p int32, /* points to the proposed match */
	id []rune, /* the identifier*/
	t int32 @t\2@>@/) bool {
	if len(name_dir[p].name)!=len(id) {
		return false
	}
	return compare_runes(id, name_dir[p].name) == 0
}

@ The common lookup routine refers to separate routines |init_node| when the data structure grows.

@c
func init_node(node int32) {
    name_dir[node].equiv=-1
}

@ Actually \.{GOTANGLE} haven't got any specific code for initialization
a new identifier, so we declare an empty corresponding section.

@<Initialization of a new identifier@>=


@** Tokens.
Replacement texts, which represent \GO/ code in a compressed format,
appear in |tok_mem| as mentioned above. The codes in
these texts are called `tokens'..

If $p$ is an index of a replacement text, |p.token| contains code of that text.
If |text_info[p].text_link==0|, this is the replacement text for a macro,
otherwise it is the replacement text for a section. 
In the latter case |text_info[p].text_link| is either equal to
|max_texts|, which means that there is no further text for this section, or
|text_info[p].text_link| points to a continuation of this replacement text; such
links are created when several sections have \GO/ texts with the same
name, and they also tie together all the \GO/ texts of unnamed sections.
The replacement text pointer for the first unnamed section appears in
|text_info[0].text_link|, and the most recent such pointer is |last_unnamed|.

@<Glob...@>=
var last_unnamed int32/* most recent replacement text of unnamed section */

@ @<Set init...@>= 
last_unnamed=0
text_info = append(text_info, text{})
text_info[0].text_link=0

@ \.{GOTANGLE} operates with UTF8 encoding texts and represents a text in
4-bytes unicode code points internally. If the first byte of a token is 
less than |unicode.UpperLower|, this is usual character. Otherwise
if it is equal |unicode.UpperLower+0211|, the next rune is a section number;
if it is equal |unicode.UpperLower+0212|, the next rune is an identifier index;
if it is equal |unicode.UpperLower+0214|, the next element is an line number;
if it is equal |unicode.UpperLower+0311| and the next rune is equal 
|unicode.UpperLower+0215|, it is a macro definition, otherwise it is 
an index of section in which the current replacement text appears.

Some of the 7-bit codes will not be present, however, so we can
use them for special purposes. The following symbolic names are used:

\yskip \hang |join| denotes the concatenation of adjacent items with no
space or line breaks allowed between them (the \.{@@\&} operation of \.{GOWEB}).

\hang |strs| denotes the beginning or end of a string, verbatim
construction or numerical constant.

@<Constants@>=
strs = 02 /* takes the place of extended ASCII \.{\char2} */
join = 0177 /* takes the place of ASCII delete */

@** Stacks for output.  The output process uses a stack to keep track
of what is going on at different ``levels'' as the sections are being
written out.  Entries on this stack have five parts:

\hang |byte_field| is a slice of the next token on a particular level;

\hang |name_field| is an index of the name corresponding to a particular level;

\hang |repl_field| is an index of the replacement text currently being read
at a particular level;

\hang |section_field| is the section number, or zero if this is a macro.

\yskip\noindent The current values of these five quantities are referred to
quite frequently, so they are stored in a separate place instead of in
the |stack| array. We call the current values |cur_state.byte_field|,
|cur_state.name_field|, |cur_state.repl_field|, and |cur_state.section_field|.

@<Typed...@>=
type output_state struct {
	byte_field []rune	/* present location within replacement text */
	name_field int32	/* |byte_start| index for text being output */
	repl_field int32	/* |token| index for text being output */
	section_field int32	/* section number or zero if not a section */
} 

@ @<Global...@>=
var cur_state output_state  /* |cur_state.byte_field|, |cur_state.name_field|,
|cur_state.repl_field|, and |cur_state.section_field| */
var stack[]output_state /* info for non-current levels */

@ To get the output process started, we will perform the following
initialization steps. We may assume that |text_info[0].text_link| is nonzero,
since it points to the \GO/ text in the first unnamed section that generates
code; if there are no such sections, there is nothing to output, and an
error message will have been generated before we do any of the initialization.

@<Initialize the output stacks@>=
cur_state.name_field=0
cur_state.repl_field=text_info[0].text_link
cur_state.byte_field=text_info[cur_state.repl_field].token
cur_state.section_field=0
stack=append(stack, output_state{})

@ When the replacement text for name |p| is to be inserted into the output,
the following subroutine is called to save the old level of output and get
the new one going.

@c
/* suspends the current level */
func push_level(p int32) { 
	stack = append(stack, cur_state)
	cur_state.name_field=p 
	cur_state.repl_field=name_dir[p].equiv
	cur_state.byte_field=text_info[cur_state.repl_field].token
	cur_state.section_field=0
}

@ When we come to the end of a replacement text, the |pop_level| subroutine
does the right thing: It either moves to the continuation of this replacement
text or returns the state to the most recently stacked level.

@c
/* do this when |cur_state.byte_field| reaches end */
func pop_level()  {
	if text_info[cur_state.repl_field].text_link<max_texts { /* link to a continuation */
		cur_state.repl_field=text_info[cur_state.repl_field].text_link /* stay on the same level */
		cur_state.byte_field=text_info[cur_state.repl_field].token
		return
	}
	
	if len(stack)>0 {
		cur_state=stack[len(stack)-1]
		stack=stack[:len(stack)-1]
	}
}

@ The heart of the output procedure is the function |get_output|,
which produces the next token of output and sends it on to the lower-level
function |out_char|. The main purpose of |get_output| is to handle the
necessary stacking and unstacking. It sends the value |section_number|
if the next output begins or ends the replacement text of some section,
in which case |cur_val| is that section's number (if beginning) or the
negative of that value (if ending). (A section number of 0 indicates
not the beginning or ending of a section, but a \&{//line} command.)
And it sends the value |identifier|
if the next output is an identifier, in which case
|cur_val| points to that identifier name.

@<Constants@>=
section_number = 0211 /* code returned by |get_output| for section numbers */
identifier = 0212 /* code returned by |get_output| for identifiers */

@ @<Global...@>=
var cur_val rune /* additional information corresponding to output token */

@
@c
/* sends next token to |out_char| */
func get_output() {
restart: 
	if len(stack)==0 {
		return
	}
	if len(cur_state.byte_field) == 0 {
		cur_val=-cur_state.section_field /* cast needed because of sign extension */
		pop_level()
		if cur_val==0 {
			goto restart
		}
		out_char(section_number)
		return
	}	
	a:=cur_state.byte_field[0]
	cur_state.byte_field = cur_state.byte_field[1:]
	if out_state==verbatim && a!=strs && a!=constant && a!='\n' {
		fmt.Fprintf(go_file, "%c", a)
	} else if a<unicode.UpperLower { 
		out_char(a)
	} else {
		c:=cur_state.byte_field[0]
		cur_state.byte_field = cur_state.byte_field[1:]
		switch a%unicode.UpperLower { 
			case identifier: 
				cur_val=c
				out_char(identifier) 
			case section_name: 
				@<Expand section |c|, |goto restart|@>
			case line_number:
				cur_val=c
				out_char(line_number)
			case section_number:	
				cur_val=c
				if cur_val>0 {
					cur_state.section_field=cur_val 
				}
				out_char(section_number)
		}
	}
}

@ The user may have forgotten to give any \GO/ text for a section name,
or the \GO/ text may have been associated with a different name by mistake.

@<Expand section |c...@>=
{
	if name_dir[c].equiv!=-1 {
		push_level(c)
	} else if a!=0 {
		fmt.Printf("\n! Not present: <")
	print_section_name(c)
		err_print(">")
		@.Not present: <section name>@>
	}
	goto restart
}

@* Producing the output.
The |get_output| routine above handles most of the complexity of output
generation, but there are two further considerations that have a nontrivial
effect on \.{GOTANGLE}'s algorithms.

@ We want to make sure that the output has spaces and line breaks in
the right places (e.g., not in the middle of a string or a constant or an
identifier, not at a `\.{@@\&}' position
where quantities are being joined together).

The output process can be in one of following states:

\yskip\hang |num_or_id| means that the last item in the buffer is a number or
identifier, hence a blank space or line break must be inserted if the next
item is also a number or identifier.

\yskip\hang |unbreakable| means that the last item in the buffer was followed
by the \.{@@\&} operation that inhibits spaces between it and the next item.

\yskip\hang |verbatim| means we're copying only character tokens, and
that they are to be output exactly as stored.  This is the case during
strs, verbatim constructions and numerical constants.

\yskip\hang |post_slash| means we've just output a slash.

\yskip\hang |normal| means none of the above.


@<Constants@>=
normal = 0 /* non-unusual state */
num_or_id = 1 /* state associated with numbers and identifiers */
post_slash = 2 /* state following a \./ */
unbreakable = 3 /* state associated with \.{@@\&} */
verbatim = 4 /* state in the middle of a string */

@ @<Global...@>=
var out_state rune  /* current status of partial output */

@ Here is a routine that is invoked when we want to output the current line.
During the output process, |line[include_depth]| equals the number of the next line
to be output.

@c
/* writes one line to output file */
func flush_buffer() {
	fmt.Fprintln(go_file)
	if line[include_depth] % 100 == 0 && show_progress() {
		fmt.Print(".")
		if line[include_depth] % 500 == 0 {
			fmt.Printf("%d",line[include_depth])
		}
		os.Stdout.Sync() /* progress report */
	}
	line[include_depth]++
}

@ If a section name is introduced in at least one place by \.{@@(}
instead of \.{@@<}, we treat it as the name of a file.
All these special sections are saved on a stack, |output_files|.
We write them out after we've done the unnamed section.

@ @<Glob...@>=
var output_files []int32
var cur_section_name_char rune /* is it |'<'| or |'('| */
var output_file_name string/* name of the file */

@ @<If it's not there, add |cur_section_name| to the output file stack, or
complain we're out of room@>=
{
	an_output_file:=0
	for ; an_output_file<len(output_files); an_output_file++ {
		if output_files[an_output_file]==cur_section_name {
			break
		}
	}
	if an_output_file==len(output_files) {
		output_files = append(output_files, cur_section_name)
	}
}

@* The big output switch.  Here then is the routine that does the
output.

@c
func phase_two () {
	line[include_depth]=1
	@<Initialize the output stacks@>
	if text_info[0].text_link==0 && len(output_files) == 0 {
		fmt.Print("\n! No program text was specified.")
		mark_harmless()
		@.No program text...@>
	} else {
		if len(output_files) == 0 {
			if show_progress() {
				fmt.Printf("\nWriting the output file (%s):",go_file_name)
			}
		} else {
			if show_progress() {
				fmt.Printf("\nWriting the output files: (%s)",go_file_name)
				@.Writing the output...@>
				os.Stdout.Sync()
			}
			if text_info[0].text_link==0 { 
				goto writeloop
			}
		}
		for len(stack)>0 {
			get_output()
		}
		flush_buffer()
writeloop:
		@<Write all the named output files@>
		if show_happiness() {
			fmt.Print("\nDone.")
		}
	}
}

@ To write the named output files, we proceed as for the unnamed
section.
The only subtlety is that we have to open each one.

@<Write all the named output files@>=
for an_output_file:=len(output_files); an_output_file>0; {
	an_output_file--
	output_file_name = string(sprint_section_name(output_files[an_output_file]))
	if f, err := os.OpenFile(output_file_name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666);
		@t\1@>@/ err != nil @t\2@> {
		fatal("! Cannot open output file:",output_file_name)
		@.Cannot open output file@>
	} else {
		go_file.Close()
		go_file = f
	}
	fmt.Printf("\n(%s)",output_file_name)
	os.Stdout.Sync()
	line[include_depth]=1
	stack = append(stack, output_state{})
	cur_state.name_field=output_files[an_output_file]
	cur_state.repl_field=name_dir[cur_state.name_field].equiv
	cur_state.byte_field=text_info[cur_state.repl_field].token
	for len(stack) > 0 {
		get_output()
	}
	flush_buffer()
}

@ A many-way switch is used to send the output.  Note that this function
is not called if |out_state==verbatim|, except perhaps with arguments
|'\n'| (protect the newline), |string| (end the string), or |constant|
(end the constant).

@c
func out_char(cur_char rune) {
	switch cur_char {
		case '\n': 
			flush_buffer()
			if out_state!=verbatim {
				out_state=normal
			}
		@<Case of an identifier@>
		@<Case of a section number@>
		@<Case of a line number@>
		@<Cases like \.{!=}@>
		case '=', '>': 
			fmt.Fprintf(go_file, "%c ", cur_char)
			out_state=normal
		case join: 
			out_state=unbreakable
		case constant: 
			switch out_state {
				case verbatim: 
					out_state=num_or_id
				case num_or_id:
					fmt.Fprint(go_file," ")
					fallthrough
				default:
					out_state=verbatim
			}
		case strs: 
			if out_state==verbatim {
				out_state=normal
			} else {
				out_state=verbatim 
			}
		case '/': 
			fmt.Fprint(go_file,"/")
			out_state=post_slash
		case '*': 
			if out_state==post_slash {
				fmt.Fprint(go_file," ")
			}
			fallthrough
		default: 
			fmt.Fprintf(go_file, "%c", cur_char)
			out_state=normal
	}
}

@ @<Cases like \.{!=}@>=
case plus_plus: 
	fmt.Fprint(go_file,"++")
	out_state=normal
case minus_minus: 
	fmt.Fprint(go_file,"--")
	out_state=normal
case gt_gt: 
	fmt.Fprint(go_file,">>")
	out_state=normal
case eq_eq: 
	fmt.Fprint(go_file,"==")
	out_state=normal
case lt_lt: 
	fmt.Fprint(go_file,"<<")
	out_state=normal
case gt_eq: 
	fmt.Fprint(go_file,">=")
	out_state=normal
case lt_eq: 
	fmt.Fprint(go_file,"<=")
	out_state=normal
case not_eq: 
	fmt.Fprint(go_file,"!=")
	out_state=normal
case and_and: 
	fmt.Fprint(go_file,"&&")
	out_state=normal
case or_or: 
	fmt.Fprint(go_file,"||")
	out_state=normal
case dot_dot_dot: 
	fmt.Fprint(go_file,"...")
	out_state=normal
case direct:
	fmt.Fprint(go_file,"<-")
	out_state=normal
case and_not:
	fmt.Fprint(go_file,"&^")
	out_state=normal
case col_eq:
	fmt.Fprint(go_file,":=")
	out_state=normal


@ @<Case of an identifier@>=
case identifier:
	if out_state==num_or_id { 
		fmt.Fprint(go_file, " ")
	}
	fmt.Fprintf(go_file, "%s", string(name_dir[cur_val].name))
	out_state=num_or_id

@ @<Case of a sec...@>=
case section_number:
	if cur_val>0 {
		fmt.Fprintf(go_file,"/*%d:*/",cur_val)
	} else if cur_val<0 {
		fmt.Fprintf(go_file,"/*:%d*/",-cur_val)
	} 

@ @<Case of a line...@>=
case line_number:
	fmt.Fprint(go_file,"\n//line ")
	@:line}{\.{\#line}@>
	line:=cur_val
	cur_val=cur_state.byte_field[0]
	cur_state.byte_field = cur_state.byte_field[1:]
	for _, v := range name_dir[cur_val].name {
		if v=='\\' || v=='"' {
			fmt.Fprint(go_file, "\\")
		}
		fmt.Fprintf(go_file, "%c", v)
	}
	fmt.Fprintf(go_file, ":%d\n", line)

@** Introduction to the input phase.
We have now seen that \.{GOTANGLE} will be able to output the full
\GO/ program, if we can only get that program into the byte memory in
the proper format. The input process is something like the output process
in reverse, since we compress the text as we read it in and we expand it
as we write it out.

There are three main input routines. The most interesting is the one that gets
the next token of a \GO/ text; the other two are used to scan rapidly past
\TEX/ text in the \.{GOWEB} source code. One of the latter routines will jump to
the next token that starts with `\.{@@}', and the other skips to the end
of a \GO/ comment.

@ Control codes in \.{GOWEB} begin with `\.{@@}', and the next character
identifies the code. Some of these are of interest only to \.{GOWEAVE},
so \.{GOTANGLE} ignores them; the others are converted by \.{GOTANGLE} into
internal code numbers by the |ccode| table below. The ordering
of these internal code numbers has been chosen to simplify the program logic;
larger numbers are given to the control codes that denote more significant
milestones.

@<Constants@>=
ignore rune = 0 /* control code of no interest to \.{GOTANGLE} */
ord rune = 0302 /* control code for `\.{@@'}' */
control_text rune = 0303 /* control code for `\.{@@t}', `\.{@@\^}', etc. */
format_code rune = 0306 /* control code for `\.{@@f}' */
definition rune = 0307 /* control code for `\.{@@d}' */
begin_code rune = 0310 /* control code for `\.{@@c}' */
section_name rune = 0311 /* control code for `\.{@@<}' */
new_section rune = 0312 /* control code for `\.{@@\ }' and `\.{@@*}' */

@ @<Global...@>=
var ccode[256] rune/* meaning of a char following \.{@@} */

@ @<Set ini...@>= {
	for c:=0; c<len(ccode); c++ {
		ccode[c]=ignore
	}
	ccode[' ']=new_section
	ccode['\t']=new_section
	ccode['\n']=new_section
	ccode['\v']=new_section
	ccode['\r']=new_section
	ccode['\f']=new_section
	ccode['*']=new_section
	ccode['@@']='@@'
	ccode['=']=strs
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
	ccode['^']=control_text
	ccode[':']=control_text
	ccode['.']=control_text
	ccode['t']=control_text
	ccode['T']=control_text
	ccode['q']=control_text
	ccode['Q']=control_text
	ccode['&']=join
	ccode['<']=section_name
	ccode['(']=section_name
	ccode['\'']=ord
}

@ The |skip_ahead| procedure reads through the input at fairly high speed
until finding the next non-ignorable control code, which it returns.

@c
/* skip to next control code */
func skip_ahead() rune {
	for true {
	    if loc>=len(buffer) && !get_line() {
			return new_section 
		}
	    for loc<len(buffer) && buffer[loc]!='@@' {
			loc++
		}
	    if loc<len(buffer) {
			loc++
			c:=new_section
			if loc < len(buffer) && buffer[loc] < int32(len(ccode)) {
				c=ccode[buffer[loc]]
			}
			loc++
			if c!=ignore || (loc <= len(buffer) && buffer[loc-1]=='>') {
				return c
			}
		}
	}
	return 0
}

@ The |skip_comment| procedure reads through the input at somewhat high
speed in order to pass over comments, which \.{GOTANGLE} does not transmit
to the output. If the comment is introduced by \.{/*}, |skip_comment|
proceeds until finding the end-comment token \.{*/} or a newline; in the
latter case |skip_comment| will be called again by |get_next|, since the
comment is not finished.  This is done so that each newline in the
\GO/ part of a section is copied to the output; otherwise the \&{\#line}
commands inserted into the \GO/ file by the output routines become useless.
On the other hand, if the comment is introduced by \.{//} (i.e., if it
is a \GO/ ``short comment''), it always is simply delimited by the next
newline. The boolean argument |is_long_comment| distinguishes between
the two types of comments.

If |skip_comment| comes to the end of the section, it prints an error message.
No comment, long or short, is allowed to contain `\.{@@\ }' or `\.{@@*}'.

@<Global...@>=
var comment_continues bool=false /* are we scanning a comment? */

@
@c
 /* skips over comments */
func skip_comment(is_long_comment bool) bool {
	for true {
		if loc>=len(buffer) {
			if is_long_comment {
				if get_line() { 
					comment_continues=true
					return comment_continues
				} else {
					err_print("! Input ended in mid-comment")
					@.Input ended in mid-comment@>
					comment_continues=false
					return comment_continues
				}
			} else {
				comment_continues=false
				return comment_continues
			}
		}
		c:=buffer[loc]
		loc++
		if is_long_comment && c=='*' && loc < len(buffer) && buffer[loc]=='/' {
			loc++
			comment_continues=false
			return comment_continues
		}
		if c=='@@' {
			if buffer[loc] < int32(len(ccode)) && ccode[buffer[loc]]==new_section {
				err_print("! Section name ended in mid-comment")
				loc--
				@.Section name ended in mid-comment@>
				comment_continues=false
				return comment_continues
			} else { 
				loc++
			}
		}
	}
	return false
}

@* Inputting the next token.

@<Constants@>=
constant = 03

@ @<Global...@>=
var cur_section_name int32/* name of section just scanned */
var no_where bool /* suppress |print_where|? */

@ As one might expect, |get_next| consists mostly of a big switch
that branches to the various special cases that can arise.

@c
/* produces the next input token */
func get_next() rune {
	for true {
		if loc>=len(buffer) {
			if !get_line() {
				return new_section
			} else if print_where && !no_where {
				print_where=false
				@<Insert the line number into |tok_mem|@>
			} else {
				return '\n'
			}
		}
		c:=buffer[loc]
		var nc rune = ' '
		if loc + 1 < len(buffer) {
			nc=buffer[loc+1]
		}
		if comment_continues || (c=='/' && (nc=='*' || nc=='/')) {
			skip_comment(comment_continues||nc=='*')
			/* scan to end of comment or newline */
			if comment_continues {
				return '\n'
			} else {
				continue
			}
		}
		loc++
		if unicode.IsDigit(c) || c=='.' {
			@<Get a constant@>
		} else if c=='\'' || c=='"' || c=='`' {
			@<Get a string@>
		} else if unicode.IsLetter(c) || c=='_' {
			@<Get an identifier@>
		} else if c=='@@' {
			@<Get control code and possible section name@>
		} else if unicode.IsSpace(c) {
			continue /* ignore spaces and tabs*/
		} 
mistake: 
		@<Compress two-symbol operator@>
		return c
	}
	return 0
}


@ @<Get an identifier@>= 
{
	loc--
	id_first:=loc
	for loc < len(buffer) && @t\1@>@/
		(unicode.IsLetter(buffer[loc]) || @/
		unicode.IsDigit(buffer[loc]) || @/
		buffer[loc]=='_' || @/
		buffer[loc]=='$' @t\2@>) { 
		loc++
	}
	id = buffer[id_first: loc]
	return identifier
}

@ @<Get a constant@>= 
{
	id_first:=loc-1
	if buffer[id_first]=='.' && (loc >= len(buffer) || !unicode.IsDigit(buffer[loc])) {
		goto mistake /* not a constant */
	}
	if buffer[id_first]=='0' {
		if loc < len(buffer) && (buffer[loc]=='x' || buffer[loc]=='X') { /* hex constant */
			loc++
			for loc < len(buffer) && xisxdigit(buffer[loc]) {
				loc++
			}
			goto found
		}
	}
	for loc < len(buffer) && unicode.IsDigit(buffer[loc]) {
		 loc++
	}
	if loc < len(buffer) && buffer[loc]=='.' {
		loc++
		for loc < len(buffer) && unicode.IsDigit(buffer[loc]) {
			loc++
		}
	}
	if loc < len(buffer) && (buffer[loc]=='e' || buffer[loc]=='E') { /* float constant */
		loc++
		if loc < len(buffer) && (buffer[loc]=='+' || buffer[loc]=='-') {
			loc++
		}
		for loc < len(buffer) && unicode.IsDigit(buffer[loc]) {
			loc++
		}
	}
found: 
	for loc < len(buffer) && @t\1@>@/
		(buffer[loc]=='u' || buffer[loc]=='U' || @/
		buffer[loc]=='l' || buffer[loc]=='L' || @/
		buffer[loc]=='f' || buffer[loc]=='F' @t\2@>) {
		loc++
	}
	id = buffer[id_first: loc]
	return constant
}

@ \GO/ strs and character constants, delimited by double and single
quotes, respectively, can contain newlines or instances of their own
delimiters if they are protected by a backslash.

@<Get a string@>= 
{
	delim := c /* what started the string */
	section_text = section_text[0:0]
	section_text = append(section_text, delim)
	
	for true {
		if loc>=len(buffer) {
			if !get_line() {
				err_print("! Input ended in middle of string")
				loc=0
				break
				@.Input ended in middle of string@>
			} else {
				section_text = append(section_text, '\n')
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
	return strs
}

@ After an \.{@@} sign has been scanned, the next character tells us
whether there is more work to do.

@<Get control code and possible section name@>= 
{
	c=ccode[nc]
	loc++
	switch c  {
		case ignore: 
			continue
		case control_text: 
			for c=skip_ahead(); c =='@@'; c = skip_ahead() {}
			/* only \.{@@@@} and \.{@@>} are expected */
			if buffer[loc-1]!='>' {
				err_print("! Double @@ should be used in control text")
				@.Double @@ should be used...@>
			}
			continue
		case section_name:
			cur_section_name_char=buffer[loc-1]
			@<Scan the section name and make |cur_section_name| point to it@>
		case strs: 
			@<Scan a verbatim string@>
		case ord: 
			@<Scan an ASCII constant@>
		default: 
			return c
	}
}

@ After scanning a valid ASCII constant that follows
\.{@@'}, this code plows ahead until it finds the next single quote.
(Special care is taken if the quote is part of the constant.)
Anything after a valid ASCII constant is ignored;
thus, \.{@@'\\nopq'} gives the same result as \.{@@'\\n'}.

@<Scan an ASCII constant@>=
	if buffer[loc]=='\\' {
		loc++
		if buffer[loc]=='\'' {
			loc++
		}
	}
	for buffer[loc]!='\'' {
		if buffer[loc]=='@@' {
			if buffer[loc+1]!='@@' {
				err_print("! Double @@ should be used in ASCII constant")
				@.Double @@ should be used...@>
			} else {
				loc++
			}
		}
		loc++
		if loc>=len(buffer) {
			err_print("! String didn't end")
			loc=len(buffer)-1
			break
			@.String didn't end@>
		}
	}
	loc++
	return ord

@ @<Scan the section name...@>= 
{
	section_text = section_text[0:0]
	@<Put section name into |section_text|@>
	if len(section_text)>3 && 
		compare_runes(section_text[len(section_text)-3:],[]rune("..."))==0 {	
		cur_section_name=section_lookup(section_text[0:len(section_text)-3],
										true) /* 1 means is a prefix */
	} else {
		cur_section_name=section_lookup(section_text, false)
	}
	if cur_section_name_char=='(' {
    @<If it's not there, add |cur_section_name| to the output file stack, or
          complain we're out of room@>
	}
	return section_name
}

@ Section names are placed into the |section_text| array with consecutive spaces,
tabs, and carriage-returns replaced by single spaces. There will be no
spaces at the beginning or the end. 

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

@ @<If end of name or erroneous nesting,...@>=
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

@ At the present point in the program we
have |buffer[loc-1]==strs|; we set |id_first| to the beginning
of the string itself, and |loc| to the position just after the ending delimiter.

@<Scan a verbatim string@>= {
	id_first:=loc
	loc++
	for loc<len(buffer) && loc+1<len(buffer) && (buffer[loc]!='@@' || buffer[loc+1]!='>') {
		loc++
	}
	if loc>=len(buffer) {
		err_print("! Verbatim string didn't end")
	}
	@.Verbatim string didn't end@>
	id=buffer[id_first:loc]
	loc+=2
	return strs
}

@* Scanning a macro definition.
The rules for generating the replacement texts corresponding to macros and
\GO/ texts of a section are almost identical; the only differences are that

\yskip \item{a)}Section names are not allowed in macros;
in fact, the appearance of a section name terminates such macros and denotes
the name of the current section.

\item{b)}The symbols \.{@@d} and \.{@@f} and \.{@@c} are not allowed after
section names, while they terminate macro definitions.

\yskip Therefore there is a single procedure |scan_repl| whose parameter
|t| specifies either |macro| or |section_name|. After |scan_repl| has
acted, |cur_text| will point to the replacement text just generated, and
|next_control| will contain the control code that terminated the activity.

@<Constants@>=
macro = 0

@ @<Global...@>=
var cur_text int32  /* replacement text formed by |scan_repl| */
var next_control rune

@
@c
/* creates a replacement text */
func scan_repl(t rune) {
	var a int32 /* the current token */
	if t==section_name {
		@<Insert the line number into |tok_mem|@>
	}
	for true {
		a=get_next()
		switch a {
			@<In cases that |a| is a non-|char| token (|identifier|,
        	|section_name|, etc.), either process it and change |a| to a byte
        	that should be stored, or |continue| if |a| should be ignored,
        	or |goto done| if |a| signals the end of this replacement text@>
		case ')': 
			tok_mem = append(tok_mem, a)
			if t==macro {
				tok_mem = append(tok_mem,' ')
			}
		default: 
			tok_mem = append(tok_mem, a) /* store |a| in |tok_mem| */
		}
	}
done: 
	next_control=a
	cur_text=int32(len(text_info))
	text_info = append(text_info, text{})
	text_info[cur_text].token=tok_mem
	tok_mem = nil
}

@ Here is the code for the line number: a first element is equal
to |unicode.UpperLower| plus |line_number|; then the numeric line number; then a pointer to the
file name.

@<Constants@>=
line_number = 0214

@ @<Insert the line...@>=
tok_mem = append(tok_mem, unicode.UpperLower + line_number)

if changing {
	id = []rune(change_file_name)
} else {
	id = []rune(file_name[include_depth])
}
if changing {
	tok_mem = append(tok_mem, rune(change_line))
} else { 
	tok_mem = append(tok_mem, rune(line[include_depth]))
}
{
	a:=id_lookup(id,0)
	tok_mem = append(tok_mem, a)
}

@ @<In cases that |a| is...@>=
case identifier: 
	a=id_lookup(id,0)
	tok_mem = append(tok_mem, unicode.UpperLower+identifier)
	tok_mem = append(tok_mem, a)
case section_name: 
	if t!= section_name{
		goto done
	} else {
		@<Was an `@@' missed here?@>
		tok_mem = append(tok_mem, unicode.UpperLower + section_name)
		a=cur_section_name
		tok_mem = append(tok_mem, a)
		@<Insert the line number into |tok_mem|@>
	}
case constant, strs:
	@<Copy a string or verbatim construction or numerical constant@>
case ord:
	@<Copy an ASCII constant@>
case definition, format_code, begin_code: 
	if t!=section_name {
		goto done
	} else {
		err_print("! @@d, @@f and @@c are ignored in Go text")
		continue
		@.@@d, @@f and @@c are ignored in Go text@>
	}
case new_section: 
	goto done

@ @<Was an `@@'...@>= {
	try_loc:=loc
	for try_loc<len(buffer) && buffer[try_loc]==' '  {
		try_loc++
	}
	if try_loc<len(buffer) && buffer[try_loc]=='+' {
		try_loc++
	}
	for try_loc<len(buffer) && buffer[try_loc]==' ' {
		try_loc++
	}
	if try_loc<len(buffer) && buffer[try_loc]=='=' {
		err_print ("! Missing `@@ ' before a named section")
		@.Missing `@@ '...@>
	}
  /* user who isn't defining a section should put newline after the name,
     as explained in the manual */
}

@ @<Copy a string...@>=
	tok_mem = append(tok_mem, a) /* |string| or |constant| */
	for i := 0; i < len(id); { /* simplify \.{@@@@} pairs */
		if id[i]=='@@' {
			if id[i+1]=='@@' {
				i++
			} else {
				err_print("! Double @@ should be used in string")
				@.Double @@ should be used...@>
			}
		}
		tok_mem = append(tok_mem, id[i])
		i++
	}
	tok_mem = append(tok_mem, a)

@ @<Copy an ASCII constant@>= 
{
	c:=id[0]
	if c=='\\' {
		id = id[1:]
		c=id[0]
		if c>='0' && c<='7' {
			c-='0'
			if id[1]>='0' && id[1]<='7' {
				id = id[1:]
				c=8*c+id[0] - '0'
				if id[1]>='0' && id[1]<='7' && c<32 {
					id = id[1:]
					c=8*c+id[0]- '0'
				}
			}
		} else {
			switch c {
				case 't':
					c='\t'
				case 'n':
					c='\n'
				case 'b':
					c='\b'
				case 'f':
					c='\f'
				case 'v':
					c='\v'
				case 'r':
					c='\r'
				case 'a':
					c='\a'
				case '?':
					c='?'
				case 'x':
					if unicode.IsDigit(id[1]) {
						id = id[1:]
						c=id[0]-'0'
					} else if xisxdigit(id[1]) && 
				 		unicode.IsLower(id[1]){
						id = id[1:]
						c=unicode.ToUpper(id[0])-'A'+10
					}
					if unicode.IsDigit(id[1]) {
						id = id[1:]
						c=16*c+id[0]-'0'
					} else if xisxdigit(id[1]) &&  
				 		unicode.IsLower(id[1]) {
						id = id[1:]
						c=16*c+unicode.ToUpper(id[0])-'A'+10
					}
				case '\\':c='\\'
				case '\'':c='\''
				case '"':c='"'
				default: 
					err_print("! Unrecognized escape sequence")
					@.Unrecognized escape sequence@>
			}
		}
	}
  	@//* at this point |c| should have been converted to its ASCII code number */
	tok_mem = append(tok_mem, constant)
	if c>=100 {
		tok_mem = append(tok_mem, '0'+c/100)
	}
	if c>=10 {
		 tok_mem = append(tok_mem, '0'+(c/10)%10)
	}
	tok_mem = append(tok_mem, '0'+c%10)
	tok_mem = append(tok_mem, constant)
}

@* Scanning a section.
The |scan_section| procedure starts when `\.{@@\ }' or `\.{@@*}' has been
sensed in the input, and it proceeds until the end of that section.  It
uses |section_count| to keep track of the current section number; with luck,
\.{GOWEAVE} and \.{GOTANGLE} will both assign the same numbers to sections.


@ The body of |scan_section| is a loop where we look for control codes
that are significant to \.{GOTANGLE}: those
that delimit a definition, the \GO/ part of a module, or a new module.

@c
func scan_section() {
	var p int32 = 0 /* section name for the current section */
	var q int32 = 0 /* text for the current section */
	var a int32 = 0 /* token for left-hand side of definition */
	section_count++
	no_where=true
	if loc < len(buffer) && buffer[loc-1]=='*' && show_progress() { /* starred section */
		fmt.Printf("*%d",section_count)
		os.Stdout.Sync()
	}
	next_control=0
	for true {
		@<Skip ahead until |next_control| corresponds to \.{@@d}, \.{@@<},
      \.{@@\ } or the like@>
		if next_control == definition {  /* \.{@@d} */
			@<Scan a definition@>
			continue
		}
		if next_control == begin_code {  /* \.{@@c} or \.{@@p} */
			p=-1
			break
    	}
		if next_control == section_name { /* \.{@@<} or \.{@@(} */
			p=cur_section_name
			@<If section is not being defined, |continue| @>
			break
		}
		return /* \.{@@\ } or \.{@@*} */
	}
	no_where=false
	print_where=false
	@<Scan the \GO/ part of the current section@>
}

@ At the top of this loop, if |next_control==section_name|, the
section name has already been scanned (see |@<Get control code
and...@>|).  Thus, if we encounter |next_control==section_name| in the
skip-ahead process, we should likewise scan the section name, so later
processing will be the same in both cases.

@<Skip ahead until |next_control| ...@>=
for next_control<definition {
      /* |definition| is the lowest of the ``significant'' codes */
	if next_control=skip_ahead(); next_control==section_name{
		loc-=2 
		next_control=get_next()
	}
}

@ @<Scan a definition@>= 
{
	/*allow newline before definition */
	for next_control=get_next(); next_control=='\n'; next_control=get_next(){} 
	if next_control!=identifier {
		err_print("! Definition flushed, must start with identifier")
		@.Definition flushed...@>
		continue
	}
	a=id_lookup(id,0)
	tok_mem = append(tok_mem, unicode.UpperLower + identifier)
	tok_mem = append(tok_mem, a)
        /* append the lhs */
	if loc < len(buffer) && 
		buffer[loc]!='(' { /* identifier must be separated from replacement text */
		tok_mem = append(tok_mem, strs)
		tok_mem = append(tok_mem, ' ')
		tok_mem = append(tok_mem, strs)
	}
	scan_repl(macro) 
	text_info[cur_text].text_link=0 /* |text_link==0| characterizes a macro */
}

@ If the section name is not followed by \.{=} or \.{+=}, no \GO/
code is forthcoming: the section is being cited, not being
defined.  This use is illegal after the definition part of the
current section has started, except inside a comment, but
\.{GOTANGLE} does not enforce this rule; it simply ignores the offending
section name and everything following it, up to the next significant
control code.

@<If section is not being defined, |continue| @>=
/* allow optional \.{+=} */
for next_control=get_next();next_control=='+'; next_control=get_next(){} 
if next_control!='=' && next_control!=eq_eq {
	continue
}

@ @<Scan the \GO/...@>=
@<Insert the section number into |tok_mem|@>
scan_repl(section_name) /* now |cur_text| points to the replacement text */
@<Update the data structure so that the replacement text is accessible@>

@ @<Insert the section number...@>=
tok_mem = append(tok_mem, unicode.UpperLower+section_number)
tok_mem = append(tok_mem, section_count)

@ @<Update the data...@>=
if p==-1 { /* unnamed section, or bad section name */
	text_info[last_unnamed].text_link=cur_text
	last_unnamed=cur_text
} else if name_dir[p].equiv==-1 {
	name_dir[p].equiv=cur_text
  /* first section of this name */
} else {
	q=name_dir[p].equiv
	for text_info[q].text_link<max_texts {
    	q=text_info[q].text_link /* find end of list */
	}
	text_info[q].text_link=cur_text
}
text_info[cur_text].text_link=max_texts
  /* mark this replacement text as a nonmacro */

@
@c
func phase_one() {
	phase=1
	section_count=0
	reset_input()
	skip_limbo()
	for !input_has_ended {
		scan_section()
	}
	check_complete()
	phase=2
}

@ Only a small subset of the control codes is legal in limbo, so limbo
processing is straightforward.

@c
func skip_limbo() {
	for true {
		if loc>=len(buffer) && !get_line() {
			return
		}
		for loc<len(buffer) && buffer[loc]!='@@' {
			loc++
		}
		if loc++; loc<len(buffer) {
			c:=buffer[loc]
			loc++
			cc := ignore
			if c < int32(len(ccode)) {
				cc = ccode[c]
			}
			if cc==new_section {
				break
			}
			switch cc {
				case format_code, '@@': 
				case control_text: 
					if c=='q' || c=='Q' {
						for c=skip_ahead(); c=='@@'; c=skip_ahead(){}
						if buffer[loc-1]!='>' {
							err_print("! Double @@ should be used in control text")
							@.Double @@ should be used...@>
						}
						break
					}
					fallthrough
				default: 
					err_print("! Double @@ should be used in limbo")
					@.Double @@ should be used...@>
			}
		}
	}
}

@
@c
func print_stats() {
    fmt.Print("\nMemory usage statistics:\n")
    fmt.Printf("%v names\n", len(name_dir))
    fmt.Printf("%v replacement texts\n", len(text_info))
}

@ \.{GOTANGLE} specific creation of output file

@<Try to open output file@>=
var err error
if go_file, err = os.OpenFile(go_file_name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666); err != nil {
   	fatal("! Cannot open output file ", go_file_name)
	@.Cannot open output file@>
}

@ @<Print usage error message and quit@>=
{
	fatal("! Usage: gotangle [options] webfile[.w] [{changefile[.ch]|-} [outfile[.go]]]\n", "")
	@.Usage:@>
}

@** Index.
Here is a cross-reference table for \.{GOTANGLE}.
All sections in which an identifier is
used are listed with that identifier, except that reserved words are
indexed only when they appear in format definitions, and the appearances
of identifiers in section names are not indexed. Underlined entries
correspond to where the identifier was declared. Error messages and
a few other things are indexed here too.
