package megapack

import (
	"fmt"

	"github.com/magefile/mage/mg"
)

type Docker mg.Namespace

func (Docker) Test() {
	fmt.Println("docker")
}

func Test() {
	fmt.Println("test")
}
