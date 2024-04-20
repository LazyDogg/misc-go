package lib

import "fmt"

type errors struct {
	PanOrPrint bool //True for Panic or False for Just Print without Panic
	Receive    *string //The Variable to receive the message
	err        error //The error to handle
}

func NewErrors(pop bool, receive *string, err error) *errors {
	//This function initates a pointer to a errors struct
	return &errors{pop, receive, err}
}

func (err *errors) AlertError() {
	//The function to handle the error, based on the struct created
	if (*err).err != nil {
		if (*err).PanOrPrint {
			panic(fmt.Sprintf("\nError -> %v\n", (*err).err))
		} else if (*err).Receive == nil && !(*err).PanOrPrint {
			fmt.Printf("\nError -> %v\n", (*err).err)
		} else {
			temp := fmt.Sprintf("\nError -> %v\n", (*err).err)
			(*err).Receive = &temp
		}
	}
}