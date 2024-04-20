# This Repository is for add miscellaneous functions

## The final intention is to abstract certain usual blocks of code

Hi, I'm using this repo to make a package to miscellaneous functions
(Based in the things I learn in the course of Go that I'm Watching)
So if you have some suggestion, please make a Issue request, I'm thankful to those
that want to help.

### My First Contribution is the Error Handling Function

This is a **simple function that opens 3 possibilities to handle the error given:**
*First is panic the program (giving true to the errors.PanOrPrint, and nil to errors.Receive)
*Second is to print in the stdout without panic (giving false to the errors.PanOrPrint, and nil to errors.Receive)
*Third is to receive the error mensage in a string variable, giving the address in errors.Receive

> Here's a example
>
> `func main(){
> var variableForError string
>   err := fmt.Errorf("Error Reached") //Here we are simulating a error
>   misc.NewErrors(false, &variableForError, err).AlertError() //Here we are calling the package with the function creating and giving the method
> fmt.Println(variableForError) //Printing in stdout like a string
> }`
>
> The function seems with the same effect compared with a err != nil if statement, but I enjoyed to make this package
> I hope that I gonna learn much more and produce more effiecient programs and packages (It's gonna be 3 days that I'm learning the Go Language)

## Special Thanks

> My special thanks is for the course I'm taking in youtube (I'm gonna leave here the link to the playlist, It's very awesome)
> And have the repo from the course too that I'm leaving here
> My thanks to my friend and my cousin that inspire me to embark in this journey to the programming world
> Links: [Course Repo](https://github.com/vkorbes/aprendago)
> Links: [Youtube Playlist (Is in Portuguese)](https://youtube.com/playlist?list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&si=7umDgPm5IzZPqW8h)
