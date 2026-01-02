# Chronoplay integeration doc

## Flow

1. Chronoplay-backend-service will initiate an event that will hit this service which will tell to setup a game with userIds provided.
2. Then the users with those ids only will be able to play game
3. When game ends then need to push an event for chronoplay-backend-service with gameid and results


## Game structure

1. Game id for this game will be different than game id in chronoplay.
2. Will store gameid received from chronoplay in here as well which will be sent in response.





## Questions

1. How will i continue this logic that if user is logged in with my service then it will be logged in into another service with same data (without using userids) => kind of google signin
2. How will i setup this event system