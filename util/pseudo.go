package util

import "log"

func CancelAll() error {
	for _, id := range orderLocalID {
		test := &ApiTest{
			"/v1/order/cancelByLocalID",
			"POST",
			map[string]interface{}{
				"orderLocalID": id,
			},
		}
		_, err := test.Send()
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}
