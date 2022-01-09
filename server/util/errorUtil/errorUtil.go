package errorUtil

import "fmt"

func Check(err error) error {
	if nil != err {
		fmt.Printf("error: %v", err)
		return err
	}
	return nil
}
