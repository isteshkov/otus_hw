package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	// Place your code here.
	for i := 0; i <= len(stages)-1; i++ {
		in = wrapStage(in, done, stages[i])
	}

	return in
}

func wrapStage(in In, done In, stage Stage) Out {
	inStream := make(Bi)
	go func() {
		defer close(inStream)
		for i := range in {
			select {
			case <-done:
				return
			default:
				inStream <- i
			}
		}
	}()
	return stage(inStream)
}
