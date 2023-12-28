todo create -l | --list (list) // creates a new list 
 
todo delete -l | --list (list) // deletes a list
 
todo add -l | --list (list) (item) // add item to specified list
 
todo remove -l | --list (list) (item) // remove item from specified list

todo check -l | --list (list) (item) // check an item as done on the specified list

todo uncheck -l | --list (list) (item) // uncheck an item as done on the specified list

todo show  //show all lists

todo show -l | --list (list) //show a specified list
 
subcommands: create, delete, add, remove, check, uncheck, show
positional args: item
flags: -l, 

subcommand flag positional
subcommand flag positional positional
subcommand
