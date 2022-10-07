
Docbuilder will be a program for assembling a single page document from several data inputs. In the first version, it will support markdown output. Later, it may support latex or possible docx.

Docbuilder aims to solve issues relating the reference and glossary management by enabling a writer to divide their document into sections folders, each with content, glossary terms and references for that section. Docbuilder will then create a merged document from the subdirectory inputs and be able to append a merged glossary and reference section.

Content files in each section subdirectory are written in simple markdown with a couple of tweaks. First, a link tag without a link, e.g. `[1]` or `[hi]` (not `[1](www.mylink.com)`) will cause docbuilder to create a local link to a reference or glossary term, which will be re-numbered in the final output. 

Reference files in each subdirectory define glossary terms and references. Glossary terms can link references in that section in the same manner that content can. Reference files consist of a number of directories to docbuilder.

## Reference Directives

Directives are either top level (two "-" character precede the directive) or second level (one "-" character precedes the directive). Each directive should be on its own line.

|Directive|Effect|
|---------|------|
|--title [newtitle]| An optional new title to give this section; defaults to foldername.|
|--mainsection|Treat this section as a main section|
|--subsectionof [mainsection]|will cause this folder to be placed a subsection of some previously defined mainsection|
|--terms|Enter term parsing mode |
|-[termname]|When in term parsing mode, next lines should be parsed as a term|
|--references|Enter reference parsing mode|
|-[ref]|When in reference parsing mode, next lines should be parsed as reference; content and terms use [ref] to link this reference.|
|--insertterms|Will cause all found terms across the entire document to be inserted in this section after the content (if any), effectively making the section the glossary section.
|--insertrefs|Will cause all found refs across the entire document to be inserted into this section after the content (if any), effectively making the section the references section.

## Example
`./section1/content.md`
```
This project will consist of well formatted documentation written in [Latex] and complying with typical software development standards. The requirements document will comply with IEEE Standards. [1]
```

`./section1/ref.txt`
```
-- mainsection
-- title Introduction
-- terms
- Latex
A software system for document preparation.[2]
- 1
https://ieeexplore.ieee.org/document/502838
- 2
https://archive.org/details/latex00lesl

--insertterms
--insertrefs

```

Output:

`./out.md`

----

# Introduction
This project will consist of well formatted documentation written in [Latex](#term-latex) and compying with typical software development standards. The requirements document will comply with IEEE standards. [1](#ref-01)

<div id="term-latex">Latex</div>

A software system for document preparation. [2](#ref-02)

<div id="ref-01">1</div>
https://ieeexplore.ieee.org/document/502838
<div id="ref-02">2</div>
https://archive.org/details/latex00lesl

---
---


# Settings

Settings are controlled by a `docbuilder.json` file in the directory where you run the docbuilder process.

|Name|Description|
|----|-----------|
|useDefaults|Causes docbuilder to use default values when a setting isn't defined.|
|Outpath|The directory where the output should be placed.|
|Outname|The name of the file that will be built.|
|FolderNames|An array of input folders, ordered in the way you want them laid out in the document.|
|ContentFileName|The name of the content files in each directory.|
|ReferenceFileName|The name of the reference files in each directory.|
|OutputType|The format of the output file. Acceptable is only 'md' for now.


