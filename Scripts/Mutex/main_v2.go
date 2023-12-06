//package main
//
///*
//  - Create a Global Maps
//  - Run Two functions : Reader, Writer : there will be issue of dirty read-write functions.
//  - Mutex : helps to lock the shared resources.
//    a.RwMutex : Multiple readers can reads the data but anyone can access.
//
//Error Handling Major Design Pattern:
//1. Return in the functions as the last value.
//2. Handling error if error != nil then do the stuff.
//3. Wrapping Error : need to some additional information, then will be pull the first place we have added here.
//*/
//
///*
//s
//- Ways to Handle Error Handling
//1. Returning Error from the function calls.
//2. Panic and Cover : Panic is used to stop the program executions.
//3. Errors Package : Also write the UnitTest of the Code in golang : Errors.Is, Errors.NotIn -> OSExists
//*/
//func main()
//	fmt.Println("testing in the new data")
//	//fmt.Println("Start of main")
//	//doSomething()
//	//fmt.Println("End of main")
//	//fmt.Println("End of main")
//	//fmt.Println("End of main")
//}
//
////func doSomething() {
////	//defer func() {
////	//	if r := recover(); r != nil {
////	//		fmt.Println("Recovered from panic:", r)
////	//	}
////	//}()
////	//fmt.Println("Start of doSomething")
////	//panic("something went wrong")
////	//fmt.Println("End of doSomething") // This won't be executed due to panic
////}
