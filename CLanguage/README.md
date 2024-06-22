# C Notes

## True or False

Standard header **stdbool.h** defines true and false as macros (since C99).

* The macro **true** expands to the integer constant 1.
* The macro **false** expands to the integer constant 0.

For the operators == and !=, according to standard.

_The == (equal to) and != (not equal to) operators are analogous
to the relational operators except for their lower precedence.
Each of the operators yields 1 if the specified relation is
true and 0 if it is false. The result has type int. For any pair
of operands, exactly one of the relations is true._
