local input = importstr './input.txt';

local lines = std.filter(function(line) line != "", std.split(input, '\n'));

local legend = {
            A: 'Rock',
            B: 'Paper',
            C: 'Scissors',
            X: 'Rock',
            Y: 'Paper',
            Z: 'Scissors',
        };

local handPoints = {
            Rock: 1,
            Paper: 2,
            Scissors: 3,
};

local games = std.map(
    function(line)
        local hands = std.split(line, ' ');
        local elfHand = legend[hands[0]];
        local myHand = legend[hands[1]];
        [elfHand, myHand],
    lines
);

local play(game) = (
    local x = game[0];
    local y = game[1];
    if x == y then 3
    else if x == 'Rock' && y == 'Paper' then 6
    else if x == 'Rock' && y == 'Scissors' then 0
    else if x == 'Paper' && y == 'Scissors' then 6
    else if x == 'Paper' && y == 'Rock' then 0
    else if x == 'Scissors' && y == 'Rock' then 6
    else if x == 'Scissors' && y == 'Paper' then 0
);

local sum(arr) = std.foldl(
    function(acc, elem) acc + elem, arr, 0
);


local totals = std.map(
    function(game)
        play(game) + handPoints[game[1]],
    games
);

sum(totals)