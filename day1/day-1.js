"use strict";

const fsPromises = require("fs").promises;

function getTotalCaloriesForNumberOfElves(input, size) {
  const elves = input.split(/\n\n/g).map(getTotalOfCalories).sort();

  return elves
    .slice(elves.length - size)
    .reduce((total, next) => total + next, 0);
}

function getTotalOfCalories(elfCalories) {
  const items = elfCalories.split(/\n/g);

  return items
    .map((i) => Number.parseInt(i))
    .reduce((total, next) => total + next, 0);
}

(async function () {
  const inputBuffer = await fsPromises.readFile("./input.txt");

  const input = inputBuffer.toString();

  const part1 = getTotalCaloriesForNumberOfElves(input, 1);
  const part2 = getTotalCaloriesForNumberOfElves(input, 3);

  console.log("Part 1:", part1);
  console.log("Part 2:", part2);
})();
