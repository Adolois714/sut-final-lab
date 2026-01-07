package entity_test

import(
	. "github.com/onsi/gomega"
	"testing"
	"test/entity"
)

//positive
func TestEmployeesValid(t *testing.T){
	g := NewGomegaWithT(t)

	employee := entity.Employees{
		Name: "Nopwi",
		Salary: 20000,
		EmployeeCode: "AD-1234",
	}

	ok , err := employee.Validate()

	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
}