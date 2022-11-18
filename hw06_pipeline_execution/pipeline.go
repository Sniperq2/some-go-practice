package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

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

		go func(Out) {
			defer close(makeBi)
			for {
				select {
				case <-done: // прерываем канал по done
					return
				case v, ok := <-in: // читаем из канала in
					if !ok { // если на входе пусто то и делать нечего
						return
					}
					makeBi <- v // в промежуточный канал
				}
			}
		}(out) // первый раз в out будет in

		out = stage(makeBi) // делаем полезную работу в стейдже
	}

	return out
}
