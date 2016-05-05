TODO
---

HTTPPollingEar
HTTPRepeatingMouth
TwitterEar
TwitterMouth
TwilioMouth

EarMouthBrain
---

Basically an intuitive way to model a network of connections

Receiver
Transmitter

Ear implements Receiver
Mouth implements Transmitter
Brain implements both

Each of them should have a dynamic connections list
InConnection []Connection
OutConnection []Connection

Connection
  Should have an Input and Output

Connection should be an interface
There's a connection from the outside world to the Ear
and a connection from the Mouth to the outside world

So the connection that are in the program itself between the building blocks should be called
LocalConnection (and it's not a network connection, it's a shared memory based connection)

Connections can be closed and reopened
Connections can be removed from the list (which means they must be closed first)
New connections can be added dynamically
