Why Go
** Motivation behind it
Infrastructure chnages in deploying application have change alot in recent time such as multicore processors in computer, container and distributed systems used in cloud computing. This infrastructure are more distributed, dynamic and have more capacity.
Before now application are built to execute tasks one at a time which means a task must complete before another one could start, but the advancement in Technology and infrastructure have allow for different task to run parallel at the same time in a single application. e.g Plating youtube video and also commenting, like a comment etc, uploading a file and also editting the name of another file at the same time, two or more people editing the same google spreadsheet at the same time. this has made application faster and user friendly. This process or concept is called multi-threading, it allow diff task to run on diff thread or say in parallel without affecting each other
Another thing is a case where we have two or more user trying to buy the last available pizza at the same time, the code or application must be able to handle all the request and make sure that no two user order the same pizza. this concept is called Concurrency (that is dealing with a lot of things at once with out conflict when dealing with the same data)

Go is built with all these beautify feature. it was design to be able to run on multiple core and with support for concurrency. The main Use case of go is for building application that have high performance and can run on scaled distributed systems, it is resource efficent. go combine the simple and readable syntax of dynamically typed language like python and efficiency and safety of low-level, statically typed language like c++, it ia a compile language. it is most sed for building server side application such as microservices, web application, database systems etc. it used for building docker, KB, hashicorp Vault, cockroach DB etc.
