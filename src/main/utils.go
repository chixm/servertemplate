package main

func stringMatches(var1, var2 string) bool {
	if var1 == `` || var2 == `` {
		logger.Info(`empty string of var1:` + var1 + ` var2:` + var2)
		return false
	}
	if var1 == var2 {
		return true
	}
	return false
}
