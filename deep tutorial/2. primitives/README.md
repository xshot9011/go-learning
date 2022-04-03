# Primitives

## **Agenda**

you will learn

- Boolean type
- Numberic types
    - intergers
    - Floating point
    - Complex numbers
- Text

## Conclusion

1. Boolean type
    1. Value are true or false
    2. Not an alias for other types (e.g. int)
    3. Zero value is false
2. Numberic types
    1. Integers
        1. Signed integers
            1. int type has varying size, but min 32 bits
            2. 8 bit (int8) through 64 bit
        2. Unsigned integers
            1. 8 bit (byte and uint8) through 32 bit(uint32)
        3. Arithmetic operations

            additions, subtractions, muliplication, division, remainder

        4. Bitwise operations

            And, or, xor, and not
        5. Zero value is 0
        6. Cannot mix types in same family (uint16 + uint32 >> error)
    2. Floating point numbers
        1. Floating point numbers
            1. follow ieee754 standard
            2. zero value is 0
            3. 32 and 64 bit version
            4. literal style
                1. Decimal (3.14)
                2. Exponential (13e10)
                3. mixed (13.14e10)
    3. Complex number
        1. zero value is (0+0i)
        2. 64 and 128 bit version
        3. Built-in function

            complex - make complex number from two float

            real - get the real part as float

            imag - get imaginary part as float

    4. Text type
        1. String
            1. utf8
            2. immutable
            3. Can be concatenated with plus(+) operation
            4. Can be convert to []byte
        2. Rune
            1. utf32
            2. alias for int32
            3. Special method normally required to process