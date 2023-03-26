// TODO: У меня завершение какое-то аварийное. Надо бы помягче

package timer

import (
	"context"
	"log"
	"time"
)

var (
	oneMinute = 1 * time.Minute
	tenMinute = 10 * time.Minute
)

type timer struct {
	start  *time.Timer
	finish *time.Timer
}

func (t *timer) Stop() {
	if t != nil {
		if t.start != nil && !t.start.Stop() {
			<-t.start.C
		}
		if t.finish != nil && !t.finish.Stop() {
			<-t.finish.C
		}
	}
}

/**
 * cur, next -- эмуляция порядкового номера минут.
 * Если выбрать диапазон [0; 9], который соответствует времени
 * ожидания оплаты, большая вероятность того, что новые заказы
 * будут перетирать старые прежде, чем старые будут обработаны.
 *
 * Диапазон [0; 59] стоит пересмотреть, но пока он может
 * эмулировать часовую работу и выглядит более простым для
 * понимания (чем каким-то образом рассчитанная верхняя граница).
 *
 * Задержка в минуту и такой большой буфер гарантируют
 * большую вероятность обработки всех заказов
 */

/**
 * в канал ToQueue записывает сервис createOrder
 * из канала ToComplete читает кто-то (воркер пул?)
 */

/**
 * подобие "управляющей горутины"
 * общий ресурс -- связки заказов для обработки
 *
 * Пока мы добавляем по одному товару в каналы за раз,
 * поэтому можно обойтись без буфера.
 * Здесь Order есть any
 * Мы оперируем указателями на заказы => статусы заказов
 * актуальны
 *
 * Заказы создаются постоянно. Включать для каждого счетчик --
 * слишком затратно. Поэтому мы собираем заказы в связки в
 * течение минуты. Таким образом, для отсчета 10и минут мы
 * привязываемся к времени, а не к самому заказу.
 *
 * Таймер работает на протяжении всей работы программы.
 * Даже при отсутствии заказов селект не блокируется, т.к.
 * таймеры врямя от времени оживляют его.
 */

func (t *timer) Start(ctx context.Context) {
	buckets := make(map[uint8][]Order)
	var cur, next uint8

	// оба таймера начали работу
	t.start = time.NewTimer(oneMinute)
	t.finish = time.NewTimer(tenMinute)

	// предотвратить срабатывание таймеров и разрядить каналы
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("something was wrong")
			return

		case order := <-ToQueue:
			if order != nil {
				buckets[next] = append(buckets[next], order)
			}

		case <-t.start.C:
			if next == 59 {
				next = 0
			} else {
				next++
			}

			t.start.Reset(oneMinute)

		case <-t.finish.C:
			t.addToComplete(ctx, buckets[cur])
			buckets[cur] = buckets[cur][:0]
			if cur == 59 {
				cur = 0
			} else {
				cur++
			}

			t.finish.Reset(oneMinute)
		}
	}
}

// в горутине сделать копию среза и уйти в фон
func (t *timer) addToComplete(ctx context.Context, orders []Order) {
	result := make([]Order, len(orders))
	copy(result, orders)

	go func(result []Order) {
		for _, order := range result {
			ToComplete <- order
		}
	}(result)
}

/// Что-то такое должен делать CreateOrder
// func addToQueue(ctx context.Context, order *Order) {
// 	toQueue <- order
// }

/// Что-то такое должен делать кто-то (воркер пул)
// func complete(ctx context.Context) {
// 	for order := range toComplete {
// 		if order.status == "awaiting payment" {
// 			doCancel(order)
// 		}
// 	}
// }

// func doCancel(order *Order) {
// 	// i do cancel
// }
