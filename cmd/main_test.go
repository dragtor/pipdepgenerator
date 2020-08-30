package cmd

import "testing"

func TestFindDependenciesForProject(t *testing.T) {
	FindDependenciesForProject("/Users/shubham/agrostar/ticketservice/")
}

func TestGenerateRequirementTxtFile(t *testing.T) {
	GenerateRequirementTxtFile("/Users/shubham/agrostar/ticketservice/", "/Users/shubham/agrostar/ticketservice/autogen_Requirement.txt")
}
