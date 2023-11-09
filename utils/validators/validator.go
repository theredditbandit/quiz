package validators

// returns true if schema of provided file matches supported csv schema
func IsValid(csvfile string) bool {
	return true // TODO:  finish implementation , use polymorphism so validator can automagically detect which filetype is supplied and use appropriate validation
}
