#!/bin/bash

echo "bad example 00"
echo -e "####\n...#\n....\n....\n" > data.txt
go run . "data.txt"
echo "----------------"

echo "bad example 01"
echo -e "...#\n..#.\n.#..\n#...\n" > data.txt
go run . "data.txt"
echo "----------------"

echo "bad example 02"
echo -e "...#\n...#\n#...\n#...\n" > data.txt
go run . "data.txt"
echo "----------------"

echo "bad example 03"
echo -e "....\n....\n....\n....\n" > data.txt
go run . "data.txt"
echo "----------------"

echo "bad example 04"
echo -e "..##\n....\n....\n##..\n" > data.txt
go run . "data.txt"
echo "----------------"

echo "bad format"
echo -e "...#\n...#\n...#\n...#\n....\n....\n....\n####\n\n\n.###\n...#\n....\n....\n\n....\n..##\n.##.\n....\n" > data.txt
go run . "data.txt"
echo "----------------"

echo "good example 00"
echo -e "....\n.##.\n.##.\n....\n" > data.txt
go run . "data.txt"
echo "----------------"

echo "good example 01"
echo -e "...#\n...#\n...#\n...#\n\n....\n....\n....\n####\n\n.###\n...#\n....\n....\n\n....\n..##\n.##.\n....\n" > data.txt
go run . "data.txt"
echo "----------------"

echo "good example 02"
echo -e "...#\n...#\n...#\n...#\n\n....\n....\n....\n####\n\n.###\n...#\n....\n....\n\n....\n..##\n.##.\n....\n\n....\n.##.\n.##.\n....\n\n....\n....\n##..\n.##.\n\n##..\n.#..\n.#..\n....\n\n....\n###.\n.#..\n....\n" > data.txt
go run . "data.txt"
echo "----------------"

echo "good example 03"
echo -e "....
.##.
.##.
....

...#
...#
...#
...#

....
..##
.##.
....

....
.##.
.##.
....

....
..#.
.##.
.#..

.###
...#
....
....

##..
.#..
.#..
....

....
..##
.##.
....

##..
.#..
.#..
....

.#..
.##.
..#.
....

....
###.
.#..
....
" > data.txt
go run . "data.txt"
echo "----------------"

echo "hard example"
echo -e "....
.##.
.##.
....

.#..
.##.
.#..
....

....
..##
.##.
....

....
.##.
.##.
....

....
..#.
.##.
.#..

.###
...#
....
....

##..
.#..
.#..
....

....
.##.
.##.
....

....
..##
.##.
....

##..
.#..
.#..
....

.#..
.##.
..#.
....

....
###.
.#..
....
"> data.txt
go run . "data.txt"
echo "----------------"