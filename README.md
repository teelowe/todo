### a stupid todo CLI
#### this was an exercise attempting to use minimal 3rd party packages and to implement interfaces for mocking the data layer in tests
```
// creates a new list 
todo create -l (list) 
 
// deletes a list
todo delete -l (list) 
 
// add item(s) to specified list
todo add -l (list) -i (item(s)) 

// remove item(s) from specified list
todo remove -l (list) -i (item(s)) 

// check item(s) as "done" on the specified list
todo check -l (list) -i (item(s)) 

// uncheck item(s) as "done" on the specified list
todo uncheck -l (list) -i (item(s)) 

// show all lists
todo show  

// show specified list(s)
todo show -l (list(s)) 
```