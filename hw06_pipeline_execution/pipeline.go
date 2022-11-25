package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func pipeline(in In, done In, Bi chan interface{}) {
	defer close(Bi)
	for {
		select {
		case v, ok := <-in: // читаем из канала in
			if !ok { // если на входе пусто то и делать нечего
				return
			}
			Bi <- v // в промежуточный канал
		case <-done: // прерываем канал по done
			return
		}
	}
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := in
	for _, stage := range stages { // для каждого стейджа
		makeBi := make(Bi) // создадим промежуточный канал
		// if ind == 0 { // если первый раз запустили то нет out и надо передать in
		//	out = in
		//}

		// in ---------|
		//  ^         out
		//  |          |
		//  ----------<-

		go pipeline(out, done, makeBi) // первый раз в out будет in

		out = stage(makeBi) // делаем полезную работу в стейдже
	}

	return out
}
