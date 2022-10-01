package msql

type Rows interface {
	Close() error

	//example of scan:
	//type user struct{}
	//
	//Scan(&us)
	//
	//us example:
	//1.us=[]*user{}
	//2.us=[]user{}
	//3.us=user{}
	//4.us=[]string{}
	//5.us=""
	Scan(any) error
}
