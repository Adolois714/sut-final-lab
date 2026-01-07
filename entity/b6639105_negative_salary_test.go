package entity_test

import(
	. "github.com/onsi/gomega"
	"testing"
	"test/entity"
)

func TestEmployeesSalary(t *testing.T){
	g := NewGomegaWithT(t)

	employee := entity.Employees{
		Name: "Nopwi",
		Salary: 3000000,
		EmployeeCode: "AD-1234",
	}

	ok , err := employee.Validate()

	g.Expect(ok).To(BeFalse())
	g.Expect(err.Error()).To(ContainSubstring("Salary must be between 15000 and 200000"))
}