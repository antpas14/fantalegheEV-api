### FantalegheEV-api

This project provides an API for the fantalegheEV project, written using openAPI. It provides simply an enpoint and define
a return type. It generates some code that can be import in a Java or Go project at the moment, and there 
are two implementations at the moment


[Java](https://github.com/antpas14/fantalegheEV)


[Go](https://github.com/antpas14/fantalegheGO)

Refer to those projects on how to import it

More about fantalegheEV project below

### Fantasy football rules
This game is based on Italian Serie A football championship, on website <a href="http://leghe.fantacalcio.it">leghe.fantacalcio.it</a> (formerly fantaleghe). I have no relationship with them.

A league is formed by teams which roaster is composed by Serie A clubs' players.
A league day is associated to a Serie A championship day. For each league day a team picks up eleven player for the starter team and a variable number of players to compose the bench.
After the game championship is ended is possible to calculate league results.
A team scores N goals based on grades (1 to 10) given to players by different Italian sports newspapers (source is configurable). 
In addition to it, malus/bonus are applied to a player's grade if a certain event happens (a red card is usually a -1, a goal scored is usually a +3. This values are configurable by a league admin and can change).

A team scores goal based on the sum of all grades inclusive of malus and bonuses. Typically, scoring 66 will award a goal, 72 will award two and so on. 
However, this is configurable by each league. 

In a *Championship* type league a team plays with all others team in the league: since team goals are scored independently of adversary (not always true), 
team goals scored in a head-to-head match, luck in the drawing phase may have had an impact on the league ranking.

This is way I come up with an algorithm which can give an estimate on how lucky, or unlucky, teams in the league are.
### Algorithm

Algorithm calculate all possible outcomes for each league day then it returns an average value

An example would be:

    TeamA 3 - TeamB 3
    TeamC 1 - TeamD 0

After this league game, ranking is:

    TeamC      3 
    TeamA      1    
    TeamB      1 
    TeamD      0 

Using the algorithm, we will have a new ranking:

    TeamA and TeamB: 1 + 3 + 3 = 7/3 = 2.33
    TeamC: 3 + 0 + 0 = 3/3 = 1
    TeamD: 0 + 0 + 0 = 0/3 = 0

### Api 

Api needs a `league-name` string, which is the name of the league from [leghe.fantacalcio.it](leghe.fantacalcio.it) returns a list of Rank object. A Rank object has the following attributes

- name: team name
- points: points that were achieved in the competition so far
- evPoints: sum of average of the points achieved using all possible combinations 

### Final considerations

It's worth saying that this estimation algorithm is far from perfect. It doesn't take into consideration that a team 
cannot play with the same team in the same round-robin of games, or that the scoring is not solely based on a team
performance but also on the opponent team performance: a team may score an extra goal if they are in the same goal band but
difference between team score is greater than an established threshold, or a team score may get a malus based on the difference 
of midfielders in the lineup of the other team (known as midfield modifier)

Finally, this is done just for fun and took an opportunity to learn/improve on some languages and technologies over the years such as
Java, Golang, Python, Docker and others. And also done to express my frustration to friends in my fantasy football league :)

### License

This work is distributed under MIT license.
