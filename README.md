# be water, my friend

Research language

## Water

Water is a programming language with no defined syntax.
It could have any language, your choice. You can adapt it
dinamically to solve a very specific or general problem.

How we do that?

Water has a simple interpreter/evaluator in few hundred lines 
but the syntax is not defined. Internally, water is a 
lisp/scheme descendent domain specific language to work with 
parser generators and then you can use a lisp-like language 
to bootstrap another one. The interpreter expose an API to 
create grammars on the fly, making it possible to switch 
between languages dinamically.

Eg.:

```scheme
;;; The initial language is scheme-based

;;; The builtin parser, that parses this very
;;; code could be changed as expressions are
;;; evaluated.

(require "./js.wt")
exports = {
	"dumb": function() { console.log("dumb lib"); }
}

(require "./go.wt")
func main() {
	print("hello world")
}
```

The only requirement is the parser generating a water s-expr

		