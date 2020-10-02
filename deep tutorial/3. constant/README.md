# Constant

## **Agenda**

you will learn

- Naming convention
- Typed constants
- Untyped constants
- Enummerated constants
- Enummerated expression

## Summary

1. Immutable, but can be shadowed
2. Replaced by compiler at compile time
    1. Value must be calculable at compile time
3. Named like variables
    1. PascalCase for exported constants
    2. camelCase for internal constants
4. Typed constant work like immutable variables
    1. Can interoperate only with the same type
5. Untyped constant work like literal
    1. Can interoperate with similar types int + {uint, int, int16, int32, ...}
6. Enumerated constant
    1. Special symbol iota allows related constants to be create easily
    2. iota start at 0 in each const block and increment by one
    3. Watch out of constant values that match zero value for variable
7. Enumerated expression
    1. Operation that can be determined at compile time are allowd
        1. Arithmetric 
        2. Bitwise operations
        3. Bitshifting
