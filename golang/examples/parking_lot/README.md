Controller layer -> Service Layer -> Repository layer

Controller
- EntryController
1. when a vehicle enters the parking it has to 
check for available parking spot in every floor
2. if available occupy the spot 
3. generate the ticket which has entry time, slot num and floor num

-ExitController
1. when a vehicle leaves the parking, it has to unallocate the spot
2.process the payment for spot
3.generate the receipt

-AdminController
1. handle entries for,
adding new floors
adding spots to the floor
delete a floor
delete spots
