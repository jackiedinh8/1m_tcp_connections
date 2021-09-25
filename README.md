# Go, libevent, state-threads 1m tcp connections implementations

Many attempts has been made to achieve this goal in the past, especially using Go and Erlang. It is easy to write a server holding one million tcp connections (c1m), however doing nothing has no practical meaning, therefore we will examine echo server under c10k load (it should be ‘good performant’ server) with different implementations (Go, C and libevent).

Check out this [article](https://medium.com/@jackiedinh8/1m-tcp-connections-in-c-511da0b1a283).
