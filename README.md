# Passgen
## A Go-based CLI tool for generating cryptographically secure passwords
## Author: Ryan Rasti

# Installation
1. cd into the passgen directory
2. Run 'go build' to build the application
3. On Windows environments, this will generate a 'passgen.exe' file. Place the compiled .exe in your system PATH
4. Once the application is on your PATH, invoke the application as demonstrated in "Usage" below

# Flags
- -length = (int) the length of the password you wish to generate.
    - DEFAULT VALUE = 8
- -number = (int) the number of passwords you wish to generate.
    - DEFAULT VALUE = 1
- -file = (bool) true/false indicating whether or not to write the generated passwords to the current working directory.
    - DEFAULT VALUE = false
- -ext = (string) the extension of the file if the "-file" flag is set to true. Currently supports "txt" and "csv" file extensions
    - DEFAULT VALUE = txt
- -type = (string) a string indicating the type of password to generate. Currently supports "alphanumeric"/"a" or "numeric"/"n" as values
    - DEFAULT VALUE = alphanumeric

# Usage 1 (results output to CLI)
* passgen -length=12 -number=3
* $ Password #1 is: oBXOR8_D5nnyOM0A
* $ Password #2 is: MSiFR8Gx1yLOJc5V
* $ Password #3 is: LmYGE03tL2aPLPsL

# Usage 2 (results output to CLI)
* passgen -length=6 -type numeric
* $ Password #1 is: 739485

# Usage 3 (results written to file)
* cd into desired directory: 'C:/Users/USERNAME/Desktop'
* $ passgen -length=12 -number=3 -file=true
* $ "Password file generated at: C:\Users\USERNAME\Desktop\passgen-output.txt"

# Usage 4 (writes only numeric results to a csv file)
* cd into desired directory: 'C:/Users/USERNAME/Desktop'
* $ passgen -length=25 -number=30 -file=true -ext=csv -type=numeric
* $ "Password file generated at: C:\Users\USERNAME\Desktop\passgen-output.txt"