package Student_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/toto3456-t/test-final/Entity"
)

func TestStudentValidation(t *testing.T) {
	g := NewGomegaWithT(t)

	studentinfo := Entity.Student{
		Name:        "test",
		Price:       10,
		Stock:       0,
		Description: "this is a test description",
	}

	t.Run("Data is corrected", func(t *testing.T) {
		obj := studentinfo
		err := obj.Validation()

		g.Expect(err).To(BeNil())
	})
	
	t.Run("Invalid Name data", func(r *testing.T){
		n := studentinfo
		n.Name = ""
		err := n.Validation()

		g.Expect(err).ToNot(BeNil())
		g.Expect(err.Error()).To(Equal("Name is required-"))
	})

	
}
