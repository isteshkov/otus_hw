package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	// Place your code here.

	if len(stages) > 1 {
		out := WrapStage(ExecutePipeline(in, done, stages[:len(stages)-1]...), done, stages[len(stages)-1])
		if out != nil {
			return out
		}
	}

	if len(stages) == 1 {
		out := WrapStage(in, done, stages[0])
		if out != nil {
			return out
		}
	}

	return nil
}

func WrapStage(in In, done In, stage Stage) Out {
	inStream := make(Bi)
	go func() {
		defer close(inStream)
		for i := range in {
			select {
			case <-done:
				return
			case inStream <- i:
			}
		}
	}()
	return stage(inStream)
}
