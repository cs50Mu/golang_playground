package main

import (
	"fmt"
	"sync"
)

// 启动2个线程，其中一个线程打印1，2，3，另一个线程打印 a，b， c，最终展示效果如下：
//   1 a 2 b 3 c

func altPrint() {
	letterChan := make(chan struct{})
	digitChan := make(chan struct{})

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i < 4; i++ {
			fmt.Printf("%v ", i)
			digitChan <- struct{}{}
			<-letterChan
		}
	}()

	go func() {
		defer wg.Done()
		for i := 'a'; i < 'd'; i++ {
			<-digitChan
			fmt.Printf("%c ", i)
			letterChan <- struct{}{}
		}
	}()

	wg.Wait()
}

// // 有一张订单表t_order，is_new_user表示新用户，用一条sql统计出当天分区中订单总金额、新用户订单数量、老用户订单数量
// select sum(amount), sum(case when is_new_user=1 then 1 else 0 end) as new_user_orders, sum(case when is_new_user=0 then 1 else 0 end) as old_user_orders from t_order;
