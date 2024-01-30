# Notes on "Writing an Interpreter in Go"

## Introduction

- Parse source code ➜ Build Abstract Syntax Tree (AST) ➜ Evaluate AST
- Key Components: lexer, parser, AST, internal object system, evaluator
- Language: Monkey
  - C-like syntax
  - Features:
    - Variables
    - Integers/Booleans
    - Arithmetic ops
    - Built-in functions
    - First-class & higher-order functions
    - Closures
    - Data types: String/Array/Hash

## Chapter 1: Lexing

### Lexical Analysis

- Lexer: reads source code, produces tokens
- Token: a pair of token type and literal value
- REPL: read-eval-print loop

## Chapter 2: Parsing

### Parsers

- Parser: reads tokens, produces AST
- "Code is data, data is code" - Lisp
- Syntactic analysis

### Parser Generators

- CFG: context-free grammar
- BNF: Backus-Naur Form

### Parser for Monkey Language

- Recursive descent parser
  -- top down operator precedence parser (Pratt parser)
- construct root node of AST and then descend
- expressions return values, statements don't
