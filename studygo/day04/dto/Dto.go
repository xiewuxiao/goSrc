package dto

//TestsuiteFirst xml最顶层
type TestsuiteFirst struct {
	NodeOrder       string      `xml:"node_order"`
	Details         string      `xml:"details"`
	TestsuiteSecond []Testsuite `xml:"testsuite"`
	Testcases       []Testcase  `xml:"testcase"`
}

//Testsuite xml第二层
type Testsuite struct {
	Nodeorder      string      `xml:"node_order"`
	Details        string      `xml:"details"`
	Testcases      []Testcase  `xml:"testcase"`
	Name           string      `xml:"name,attr"`
	Internalid     string      `xml:"internalid,attr"`
	TestsuiteThird []Testsuite `xml:"testsuite"`
}

// //Testsuite xml第三层
// type Testsuite struct {
// 	Nodeorder string     `xml:"node_order"`
// 	Details   string     `xml:"details"`
// 	Testcases []Testcase `xml:"testcase"`
// 	Name      string     `xml:"name,attr"`
// }

//Testcase 结构体为xml文件对应的实体类
type Testcase struct {
	NodeOrder             string `xml:"node_order"`
	Externalid            string `xml:"externalid"`
	Version               string `xml:"version"`
	Summary               string `xml:"summary"`
	Preconditions         string `xml:"preconditions"`
	ExecutionType         string `xml:"excution_type"`
	Importance            string `xml:"importance"`
	EstimatedExecDuration string `xml:"estimated_exec_duration"`
	Status                string `xml:"status"`
	Steps                 Steps  `xml:"steps"`
	Name                  string `xml:"name,attr"`
	Internalid            string `xml:"internalid,attr"`
}

//Steps testcase中的步骤
type Steps struct {
	Step []Step `xml:"step"`
}

//Step Steps里的元素
type Step struct {
	StepNumber      string `xml:"step_number"`
	Action          string `xml:"actions"`
	Expectedresults string `xml:"expectedresults"`
	ExecutionType   string `xml:"execution_type"`
}
