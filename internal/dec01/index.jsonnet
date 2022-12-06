local input = importstr './input.txt';

local lines = std.split(input, '\n\n');

local caloriesByElf = std.map(
    function(line) (
        local calorieCounts = std.split(line, '\n');
        std.map(
            function(countStr)
                if countStr == '' then 0 else std.parseInt(countStr),
            calorieCounts
        )
    ),
    lines,
);

local sum(arr) = std.foldl(
    function(acc, elem) acc + elem, arr, 0
);

local list = std.reverse(std.sort(std.map(
    function(countArr)
        sum(countArr),
    caloriesByElf,
)));

sum([list[0], list[1], list[2]])