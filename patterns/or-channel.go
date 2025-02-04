package patterns

func Or[T any](channels ...<-chan T) <-chan T {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0] // If only one channel, return it directly
	}

	orDone := make(chan T)

	go func() {
		defer close(orDone) // Ensure cleanup

		switch len(channels) {
		case 2:
			select {
			case res := <-channels[0]: // First response wins
				orDone <- res
			case res := <-channels[1]:
				orDone <- res
			}
		default:
			select {
			case res := <-channels[0]:
				orDone <- res
			case res := <-channels[1]:
				orDone <- res
			case res := <-channels[2]:
				orDone <- res
			case res := <-Or(append(channels[3:], orDone)...): // Recursively reduce problem size
				orDone <- res
			}
		}
	}()

	return orDone
}
