Goroutines (managed by Go runtime) are designed for concurrency and allow multiple functions to be executed independently, making it easy to write concurrent code. Goroutines run in the same address space and share memory.

Traditional Threads are a lower-level operating system construct that can be used for both concurrency and parallelism. Threads can run in parallel on multiple CPU cores and may not share memory depending on the programming language and threading model used. Threads often follow preemptive scheduling, which means that the operating system can forcibly interrupt the execution of a thread and schedule another thread to run in its place.

Goroutines are created in user space, while OS threads are created in the Kernel space. This gives the Go scheduler complete control over the scheduling of goroutines.

Go scheduler

It's part of the Go runtime and is responsible to managing the execution of goroutines. It's an M:N scheduler, meaning it multiplexes M goroutines onto N OS threads. 
The primary job of the scheduler is to distribute goroutines over available threads and cores efficiently. Its responsibilities include:

Goroutine Management: Deciding which goroutine runs on which thread and when. Go Scheduler acts like an orchestrator, managing the goroutines.

Multiplexing: Go Scheduler is responsible for performing the M:N multiplexing, where it can efficiently run M goroutines on N OS threads.

Cooperative Scheduling & Preemption: Interrupting goroutines ensures that long-running goroutines don't starve others from CPU time. From Go 1.14, the Go scheduler preempts goroutines, rather than following cooperative scheduling.