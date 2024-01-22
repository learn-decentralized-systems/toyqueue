#   Toy queue

This micro-library contains [][]byte-centric primitives (interfaces,
job queues). The default approach in Golang is something like
"protobuf for serialization, channels for job queues". Protobuf
implies request-response RPC and full parsing, channels imply 
item-by-item push/pop. That is not really that handy if you have 
real-time streams of tiny records; imagine network packet processing 
or database ops. Assuming your data serialization format is 
well-specified and you can parse things lazily, what will work 
better for you is probably this:
````
type Records [][]byte
````
This approach encourages batching; for example, `RecordQueue` 
consumer takes all the records at once. Also, [][]byte works well
with `writev()` and `net.Buffers`.