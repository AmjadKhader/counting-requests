# Counting Requests - Moving window

A Golang HTTP server that on each request responds with a
counter of the total number of requests that it has received during the previous 60 seconds with using only the standard library.

The server continue to return the correct numbers after restarting it, by
persisting data to a file "counter_data.json".