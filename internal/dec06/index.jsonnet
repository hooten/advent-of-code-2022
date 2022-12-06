local input = importstr './input.txt';

local chars = std.filter(
  function(line)
    line != '\n',
  input,
);

//local mapped = std.mapWithIndex(
//  function(i, char)
//    char,
//  chars,
//);

chars