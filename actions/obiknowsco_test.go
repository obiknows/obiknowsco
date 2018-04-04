package actions

func (as *ActionSuite) Test_HomeHandler() {
	res := as.HTML("/").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Inertia EP")
}

func (as *ActionSuite) Test_WorkHandler() {
	res := as.HTML("/work").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Obi Knows")
}
