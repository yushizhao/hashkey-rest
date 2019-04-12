package util

import "log"

func CancelAll() {
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
}

func CheckOrder(id string) error {
	test := &ApiTest{
		"/v1/order/getOrderByLocalID",
		"POST",
		map[string]interface{}{
			"orderLocalID": id,
		},
	}
	_, err := test.Send()
	return err
}
