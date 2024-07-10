package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In //alias for out
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	for _, stage := range stages {
		ch := make(Bi)
		go createGForStage(stage, in, ch, done)
		in = ch
	}

	return in
}

func createGForStage(stage Stage, out Out, ch Bi, done In) {
	defer close(ch)
	for {
		select {
		case valueFromOut, ok := <-out:
			if !ok {
				return
			}
			resCh := stage(func(done In, valueFromOut interface{}) Out {
				toNextStage := make(Bi)
				go func() {
					defer close(toNextStage)
					select {
					case <-done: // сигнал на выход
						return
					default:
						toNextStage <- valueFromOut
					}
				}()
				return toNextStage
			}(done, valueFromOut))

			for res := range resCh {
				ch <- res
			}
		// сигнал на выход
		case <-done:
			return
		}
	}
}
