# data type in memory

bool

signed integer

+ sbyte
+ int
+ sort
+ long

unsigned integer

+ ulong
+ uint
+ ushort
+ strinb
+ byte
+ object
+ char
+ decimal (?)

float

+ float
+ double
+ decimal (?)


# Overflow

> CPU register overflow flag

> install microsoft calculator for android

## overflow operation

=
-
*
/
%
left shifting

% and other bitwise ops can't cause overflow


if both operands are < 50% the capacity of the data type,
Then no overflow is possible

1/n can be represented exactly if the only prime factors of n are 2 and 5
because 10 = 2 * 5

> IEEE 754/IEC 60559:1989

0/0 is NaN
1/0 is infinity

NaA + anything = NaN
NaN * anything = NaN
4 + infinity = infinity
4/infinity = 0
infinity*infinity = infinity


> Never use floating point numbers as map keys 

i> Don't compare floating points for equality
Alternatives to Floating Point Equality
 + Can you use integers
 + Can you use comparison

 // if (f1 == f2)
 if (f1 >= f2)
+ Can you check for nearness in value?

//if (f1 == f2)
epsilon = 1.0e-6
if abs(f1-f2) < epsilon


# Logical Operators

+ NOT
+ AND
+ OR
+ Exclusive OR
+ Inclusive OR

> If John eats the ice cream or he eats the cake, there will only be one thing left for me to eat.


> Truth table

NOT (x > 3) AND (NOT found)  => NOT(x>3) OR found => x<=3 OR found

Removing a Flag from a value

RegexOptions = oldValue & (~RegexOptions.IgnoreCase0)

