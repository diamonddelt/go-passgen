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

# Usage 1 (results output to CLI)
* passgen -length=12 -number=3
* $ Password #1 is: oBXOR8_D5nnyOM0A
* $ Password #2 is: MSiFR8Gx1yLOJc5V
* $ Password #3 is: LmYGE03tL2aPLPsL

# Usage 2 (results written to file)
* cd into desired directory: 'C:/Users/<USER>/Desktop'
* $ passgen -length=12 -number=3 -file=true
* $ "Password file generated at: C:\Users\<USER>\Desktop\passgen-output.txt"