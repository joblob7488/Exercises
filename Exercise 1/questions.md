Exercise 1 - Theory questions
-----------------------------

### Concepts

What is the difference between *concurrency* and *parallelism*?
> The difference between concurrency and paralellism is that in paralellism everytihing is actually happening at the same time, while in concurrency, it just looks like it does. This is because it switching between very fast.

What is the difference between a *race condition* and a *data race*? 
> A race condition is when the outcome of a recource is dependent of the ordering in time of the things. When the result is dependent of the timing in time. Data race is when two threads access a shared resource without any syncronization.
 
*Very* roughly - what does a *scheduler* do, and how does it do it?
> A scheduler is the comonent of a operating system wich is responsible for decining and handling wich thread to run next. Can manage excecution of task to acheive concurrency.


### Engineering

Why would we use multiple threads? What kinds of problems do threads solve?
> Multiple threads is beneficial when we need somehing to happen at the same time, or to let seperate processes be totaly devided from each other. Running some task in paralell can help improve speed.

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?
> Fibers are somewhat like as threads, and are also sharing adress space, but fibers are not connected to hardware and cannot have preemtive scheduling. Threads is the only thing that has the preemptive scheduling. Fibers are not managed by the OS, but rather by a itself or another application, and thus becomes more lightweight than threads.

Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
>Easier as it opens up a lot of implementation possebilities, harder as syncronisation has to be considered between the threads.

What do you think is best - *shared variables* or *message passing*?
> shared variables!


