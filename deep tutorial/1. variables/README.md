# Go variable

## **Agenda**

you will learn

- Variable Declaration
- Redeclaration and shadowing
- Visibility
- Naming conventions

## Conclusion

1.  Variable declaration

    ```go
    var [name] [type]
    var [name] [type] = [value]
    [name] := [value]
    ```

2. Cannot redeclare variables, but can shadow the
3. All variable must be used
4. Visibility
    1. lower case first letter for package scope
    2. upper case first letter to export
    3. no private scope 
5. Naming conventions
    1. Pascal or CamelCase

        Capitalizer acronyms (HTTP, URL)

    2. As short as reasonable

        longer name for longer lives

6. Type conversions

    ```go
    [destination_type]([var_name])
    // import strconv and use intellisense to find function inside of it
    ```

    use strconv package for string