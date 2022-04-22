# is-it-up

It's a server checker allowing for the use of seeing if the designated IP and Port are a server that is up and running


---Useage---


To run, 

download the .go file and run (in a Terminal or Command Prompt,)

"go run serverchecker.go 129.189.124.21 80"

In this case, 129.189.124.21 is the server and 80 is the (HTTP) Port

---Troubleshooting---


When running the script, it will either respond saying the server is down or that it is up. 


---Future Updates---

- Error responses (IE, No specified IP Addrress / Port / Non-Existent Server)
- Way to use UDP and other forms of connectivitiy other than TCP.

