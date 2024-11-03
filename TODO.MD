GOT TO IMPLEMENT FULL MODEL INTO DB 
     -> COME UP WITH IT FIRST 
index 
    - just a new game thing
    - current games?? - not for now

new game 
    - config and rules / invite linnk gen
    - timer (clue/position guess)
    - enter username + colour for each player
    - points
    - add kick function to lobby

game page / url based games but hard to guess 
    - add start new game to game page
    - leaderboard 
    - turns 
    - spinner 
    - polling
    - game state ?? HATEOAS?
    - svg spinner while waiting
    - svg spinner while typing
    - line position js thing
    - next prompt voting??? 

would be nice 
    - some kind of easy to memorize phrase that 
        gets hashed into an id 

//in a lobby you can 
//edit settings 
//view current players 
//vote to start game


notes - 
host loop 
    - index 
    - createGame -> /game?id=123
        . lobby 
        . votes to ready + has kick ability
    - start game 
        . gameplay loop 
    - back to home 

joiner loop 
    - index 
    - join game with code -> /game?id=123
    - votes to ready 
    - start game 
        . gameplay loop
