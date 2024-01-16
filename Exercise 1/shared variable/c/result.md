T3
C:
A thread is a peice of code that can run "simuntaneously" as other code, this can be done by pthread in c++. We first create two seperate threads, one that increment a global variable; one that decrement the same varable. They should both run at psuedo the same time.
The pthread join function holds the current threads and suspends the calling thread until the target thread terminates. Regarding of the history of incremeentation and decrementetion, we end up with a different number each time. 

GO:
The GOMAXPROCS variable limits the number of operating system threads that can execute user-level Go code simultaneously. There is no limit to the number of threads that can be blocked in system calls on behalf of Go code; those do not count against the GOMAXPROCS limit. The result becomes the same as with the c program.

T4
C:
Mutex = mutual exclution
Is used whenever we need to have multiple threads to share the same resource, and only the one who locked the rosource, can unlock.

Semaphore = global value of 0 or 1, anyone can have the key and access the resource

In this case mutex is the best, as the decrementation/incrementation needs to be done first, and before the other begins.