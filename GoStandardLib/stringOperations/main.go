package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	x := "The quick, brown fox jumps over! the lazy dog"
	fmt.Println("The string used:", x)

	fmt.Println("Length:", len(x))

	//Iterate over each character
	for _, ch := range x {
		fmt.Print(string(ch), ", ")
	}
	fmt.Println()

	fmt.Println("strings.Compare(\"dog\",\"dog\"):", strings.Compare("dog", "dog"))

	fmt.Println("strings.EqualFold(\"dog\",\"Dog\"):", strings.EqualFold("dog", "Dog"))

	x1 := strings.ToLower(x)
	fmt.Println(x1)

	x2 := strings.ToUpper(x)
	fmt.Println(x2)

	x3 := strings.Title(x)
	fmt.Println(x3)

	//string searches
	fmt.Println("strings.Contains(x, \"jump\")", strings.Contains(x, "jump"))
	fmt.Println("strings.ContainsAny(x, \"abcd\")", strings.ContainsAny(x, "abcd"))

	//Index
	fmt.Println("strings.Index(x, \"fox\")", strings.Index(x, "fox"))
	fmt.Println("strings.Index(x, \"cat\")", strings.Index(x, "cat"))
	fmt.Println("strings.IndexAny(x, \"abcd\")", strings.IndexAny(x, "abcd")) // c is in the 7th spot

	//Prefix or Suffix
	fmt.Println("strings.HasPrefix(x, \"The\")", strings.HasPrefix(x, "The"))
	fmt.Println("strings.HasPrefix(x, \"the\")", strings.HasPrefix(x, "the"))

	fmt.Println("strings.HasSuffix(x, \"Dog\")", strings.HasSuffix(x, "Dog"))
	fmt.Println("strings.HasSuffix(x, \"dog\")", strings.HasSuffix(x, "dog"))

	//Count
	fmt.Println("strings.Count(x, \"the\")", strings.Count(x, "the"))
	fmt.Println("strings.Count(x, \"he\")", strings.Count(x, "he"))

	//string manipulation
	sub1 := strings.Split(x, " ")
	fmt.Printf("%q\n", sub1)

	result := strings.Join(sub1, "-")
	fmt.Printf("%q\n", result)

	result2 := strings.Fields(x)
	fmt.Printf("%q\n", result2)

	//splitting using a function
	f := func(c rune) bool {
		return unicode.IsPunct(c)
	}
	result3 := strings.FieldsFunc(x, f)
	fmt.Printf("%q\n", result3)

	//use a replacer
	rep := strings.NewReplacer(",", "|", "!", "|")
	result4 := rep.Replace(x)
	fmt.Println(result4)

}
