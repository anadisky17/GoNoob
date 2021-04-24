
package main
import (

  "fmt"
  "unicode/utf8"

)

func main() {
// Strings in Go are UTF-8 encoded by default. Strings are defined between double quotes "..." and not single quotes

//  String is Immutable can cannot be changes as In Go, a string is in effect a read-only slice of bytes.
// Therefore, str1[1]="t" throws an error

 str1:= "hello"

 var str2 string
 str2="world"

 var str3 string = "!"
/*.................*/



 fmt.Println(str1,str2,str3)

 str4:= str1+str2+str3 //Concaetante string

 fmt.Println(str4)
fmt.Println(str2)
 str2=str2+str1
fmt.Println(str2)
/*...............*/

// Here each char is uint8 type as each char is 8bit=1byte long
// UTF-8 char can be defined from 1 byte to 4 bytes (UTF-8 is variable encoding),ASCII char only take 1 byte
// But Non-ASCII take between 1-4 bytes
// Therefore In go all characters are then represented in int32

bookTitle:="The Road Trip To London"
for i:=0;i<len(bookTitle);i++{

      fmt.Printf("\n bytes value of each char=%v",bookTitle[i])
      fmt.Printf("\t %dth char of book title= %c",i,bookTitle[i])
      fmt.Printf("\n Type of each char=%T  ",bookTitle[i])

}

// Here each char is uint8 type why?-because LOOP TRAVERSES ONE BYTE AT A TIME NOT ONE CHAR AT A TIME.
// If a char is a  NON-ASCII (2-4 bytes) it will not print properly:

bookTitle2:="The Rõad Trip To Lõndõn"  // õ is Latin char having 2 bytes space
for i:=0;i<len(bookTitle2);i++{

      fmt.Printf("\n bytes value of each char=%v",bookTitle2[i])
      fmt.Printf("\t %dth char of book title= %c",i,bookTitle2[i])
      fmt.Printf("\tType of each char=%T  ",bookTitle2[i])

} // Wrong Output


// For this GO has  a RUNE TYPE. it is a int32 type and therefore can represent NON-ASCII char.
//Any valid UTF-8 character within a single quote (') is a rune and its type is int32

thisIsRune:='h'

fmt.Printf("\n Type is=%T \t byte value is=%v \t value is=%c",thisIsRune,thisIsRune,thisIsRune)


// use utf package to decode each rune literals
fmt.Println("\n \t Length of String with Rune:",utf8.RuneCountInString(bookTitle2))
//Iterate over a Rune in string,  use utf8.DecideRuneInString it retruns len of each char in bytes and value
 for i:=0;i<len(bookTitle2);{

      v,size:=utf8.DecodeRuneInString(bookTitle2[i:])

      fmt.Printf("\n byte length of each char=%d",size)
      fmt.Printf("\t %dth char of book title= %c",i,v)
      fmt.Printf("\tType of each char=%T  ",v)
      i=i+size
} // Right Output

fmt.Printf("\n\n")
// the Range for loop auto adjust the len in index for each char
for i,v:= range bookTitle2  {

      fmt.Printf("\n Char Number Value value of each char=%v",v)
      fmt.Printf("\t %dth char of book title= %c",i,v)
      fmt.Printf("\tType of each char=%T  ",v)


}
fmt.Printf("\n\n")

bookTitle2b:=[]rune(bookTitle2) //or convert into slice of runes
//A []rune conversion applied to a UTF-8-encoded string returns the sequence of Unicode
//code points that the string encodes


for i,v:= range bookTitle2b{

      fmt.Printf("\n Char Number Value value of each char=%v",v)
      fmt.Printf("\t %dth char of book title= %c",i,v)
      fmt.Printf("\tType of each char=%T  ",v)


} //Right Output




}
